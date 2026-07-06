// 新闻采集器定义

package collector

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
	"isolo-news/server/internal/config"
	"isolo-news/server/internal/repository/ent"
	"isolo-news/server/internal/repository/ent/article"
	"isolo-news/server/internal/repository/ent/source"
)

// Collector 新闻采集器
type Collector struct {
	db        *ent.Client
	cfg       *config.CollectorConfig
	fp        *gofeed.Parser
}

// NewCollector 创建采集器
func NewCollector(db *ent.Client, cfg *config.CollectorConfig) *Collector {
	fp := gofeed.NewParser()
	fp.UserAgent = cfg.UserAgent

	return &Collector{
		db:  db,
		cfg: cfg,
		fp:  fp,
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
		// 截断过长的摘要
		if len(summary) > 500 {
			summary = summary[:500]
		}

		// 计算发布时间
		publishedAt := time.Now().Unix()
		if item.PublishedParsed != nil {
			publishedAt = item.PublishedParsed.Unix()
		} else if item.UpdatedParsed != nil {
			publishedAt = item.UpdatedParsed.Unix()
		}

		// 获取作者
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

		// 估算阅读时间（按中文 300 字/分钟）
		contentLen := len(item.Content)
		if contentLen == 0 {
			contentLen = len(summary)
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

	return results, nil
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
