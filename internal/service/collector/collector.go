// 新闻采集器定义

package collector

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
	"isolo-news/internal/config"
	"isolo-news/internal/ent"
	"isolo-news/internal/ent/article"
	"isolo-news/internal/ent/category"
	"isolo-news/internal/ent/source"
	"isolo-news/internal/service/reader"
)

// Collector 新闻采集器
type Collector struct {
	db        *ent.Client
	cfg       *config.CollectorConfig
	fp        *gofeed.Parser
	rd        *reader.Reader // 全文提取器（RSS 内容不足时自动提取）
}

// NewCollector 创建采集器
func NewCollector(db *ent.Client, cfg *config.CollectorConfig) *Collector {
	fp := gofeed.NewParser()
	fp.UserAgent = cfg.UserAgent

	// 初始化全文提取器
	rd := reader.NewReader(reader.Config{
		Enabled:          cfg.Readability.Enabled,
		MinContentLength: cfg.Readability.MinContentLength,
		Timeout:          cfg.Readability.Timeout,
		Concurrency:      cfg.Readability.Concurrency,
		UserAgent:        cfg.UserAgent,
	})

	return &Collector{
		db:  db,
		cfg: cfg,
		fp:  fp,
		rd:  rd,
	}
}

// CollectFromSource 从单个数据源采集新闻
func (c *Collector) CollectFromSource(ctx context.Context, src *ent.Source) (int, error) {
	log.Printf("[采集器] 开始采集: %s (%s)", src.Name, src.FeedURL)

	// 设置超时上下文
	timeout := time.Duration(c.cfg.Timeout) * time.Second
	fetchCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// 解析 RSS Feed
	feed, err := c.fp.ParseURLWithContext(src.FeedURL, fetchCtx)
	if err != nil {
		// 标记采集失败
		now := time.Now().Unix()
		_, updateErr := c.db.Source.Update().
			Where(source.IDEQ(src.ID)).
			SetLastFailedAt(now).
			SetLastError(err.Error()).
			AddFailCount(1).
			Save(ctx)
		if updateErr != nil {
			log.Printf("[采集器] 更新失败状态出错: %v", updateErr)
		}
		return 0, fmt.Errorf("解析 RSS 失败 %s: %w", src.FeedURL, err)
	}

	if feed == nil {
		return 0, fmt.Errorf("RSS 源返回空内容: %s", src.FeedURL)
	}

	// 统计成功数
	successCount := 0

	// 遍历文章
	for _, item := range feed.Items {
		if item.Title == "" {
			continue
		}

		// 构建文章数据
		sourceURL := item.Link
		if sourceURL == "" {
			continue
		}

		// 去重检查
		exists, err := c.db.Article.Query().
			Where(article.SourceURLEQ(sourceURL)).
			Exist(ctx)
		if err != nil {
			log.Printf("[采集器] 去重检查出错: %v", err)
			continue
		}
		if exists {
			continue
		}

		// 提取摘要和内容（相互兜底）
		summary := item.Description
		content := item.Content
		if summary == "" {
			summary = content
		}
		if content == "" {
			content = summary
		}
		// 两个都为空时使用标题
		if summary == "" {
			summary = item.Title
			content = item.Title
		}

		// 截断过长的摘要（按 rune 截断，避免截断中文字符）
		if len(summary) > 500 {
			runes := []rune(summary)
			if len(runes) > 500 {
				summary = string(runes[:500])
			}
		}

		// 计算发布时间
		publishedAt := time.Now().Unix()
		if item.PublishedParsed != nil {
			publishedAt = item.PublishedParsed.Unix()
		} else if item.UpdatedParsed != nil {
			publishedAt = item.UpdatedParsed.Unix()
		}

		// 获取作者（优先使用 RSS 提供的作者）
		author := src.Name
		if item.Author != nil && item.Author.Name != "" {
			author = item.Author.Name
		}

		// 解析标签
		tags := make([]string, 0)
		for _, cat := range item.Categories {
			if cat != "" {
				tags = append(tags, cat)
			}
		}

		// 获取封面图
		imageURL := ""
		if item.Image != nil && item.Image.URL != "" {
			imageURL = item.Image.URL
		}

		// 如果 RSS 提供的内容过短，尝试 Readability 全文提取
		if c.rd.ShouldExtract(content) {
			result := c.rd.Extract(ctx, sourceURL)
			if result.Success {
				log.Printf("[采集器] Readability 提取成功 [%s]: %d 字",
					item.Title, len([]rune(result.Content)))
				content = result.Content
				summary = result.Summary
				// 用提取的信息补充缺失字段（RSS 未提供时才覆盖）
				if imageURL == "" && result.ImageURL != "" {
					imageURL = result.ImageURL
				}
				if result.Byline != "" && author == src.Name {
					author = result.Byline
				}
			}
		}

		// 估算阅读时间（按中文 300 字/分钟）
		contentLen := len([]rune(content))
		if contentLen == 0 {
			contentLen = len([]rune(summary))
		}
		readTime := contentLen / 300
		if readTime < 1 {
			readTime = 1
		}
		if readTime > 30 {
			readTime = 30
		}

		// 写入数据库
		_, err = c.db.Article.Create().
			SetTitle(item.Title).
			SetSummary(summary).
			SetContent(content).
			SetCategoryID(src.CategoryID).
			SetAuthor(author).
			SetPublishedAt(publishedAt).
			SetReadTime(readTime).
			SetXpReward(calculateXpReward(readTime)).
			SetImageURL(imageURL).
			SetTags(tags).
			SetSourceID(src.ID).
			SetSourceURL(sourceURL).
			SetSourceName(src.Name).
			SetCollectedAt(time.Now().Unix()).
			Save(ctx)
		if err != nil {
			log.Printf("[采集器] 写入文章失败 [%s]: %v", item.Title, err)
			continue
		}

		successCount++
	}

	// 更新采集成功状态
	now := time.Now().Unix()
	_, err = c.db.Source.Update().
		Where(source.IDEQ(src.ID)).
		SetLastFetchedAt(now).
		SetFailCount(0).
		SetLastError("").
		Save(ctx)
	if err != nil {
		log.Printf("[采集器] 更新成功状态出错: %v", err)
	}

	log.Printf("[采集器] 采集完成: %s, 新增 %d 篇", src.Name, successCount)
	return successCount, nil
}

// CollectAll 采集所有活跃数据源
func (c *Collector) CollectAll(ctx context.Context) (map[string]int, error) {
	sources, err := c.db.Source.Query().
		Where(source.IsActive(true)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询活跃数据源失败: %w", err)
	}

	results := make(map[string]int)
	for _, src := range sources {
		count, err := c.CollectFromSource(ctx, src)
		if err != nil {
			log.Printf("[采集器] 采集数据源 [%s] 失败: %v", src.Name, err)
			results[src.Name] = 0
			continue
		}
		results[src.Name] = count
	}

	// 采集完成后更新各分类的文章总数
	if err := c.UpdateCategoryTotals(ctx); err != nil {
		log.Printf("[采集器] 更新分类总数失败: %v", err)
	}

	return results, nil
}

// UpdateCategoryTotals 更新所有分类的文章总数到数据库
// 每次采集完成后调用，确保分类 total 值实时准确
func (c *Collector) UpdateCategoryTotals(ctx context.Context) error {
	categories, err := c.db.Category.Query().All(ctx)
	if err != nil {
		return fmt.Errorf("查询分类列表失败: %w", err)
	}

	for _, cat := range categories {
		count, err := c.db.Article.Query().
			Where(article.CategoryIDEQ(cat.Slug)).
			Count(ctx)
		if err != nil {
			log.Printf("[采集器] 统计分类 [%s] 文章数失败: %v", cat.Slug, err)
			continue
		}

		_, err = c.db.Category.Update().
			Where(category.IDEQ(cat.ID)).
			SetTotal(count).
			Save(ctx)
		if err != nil {
			log.Printf("[采集器] 更新分类 [%s] total 失败: %v", cat.Slug, err)
			continue
		}
	}

	log.Printf("[采集器] 分类总数更新完成")
	return nil
}

// calculateXpReward 根据阅读时间计算 XP 奖励
func calculateXpReward(readTime int) int {
	if readTime <= 3 {
		return 50
	} else if readTime <= 8 {
		return 100
	} else if readTime <= 15 {
		return 200
	} else {
		return 300
	}
}
