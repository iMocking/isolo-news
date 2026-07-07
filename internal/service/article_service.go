// 资讯业务逻辑

package service

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"isolo-news/internal/dto"
	"isolo-news/internal/ent"
	"isolo-news/internal/ent/article"
	"isolo-news/internal/ent/articleread"
	"isolo-news/internal/ent/favorite"
	"isolo-news/internal/ent/user"
)

// ArticleService 资讯服务
type ArticleService struct {
	db          *ent.Client
	game        *GameService
	categoryMap map[string]*ent.Category // slug -> 分类缓存，toVO 中使用
}

// NewArticleService 创建资讯服务
func NewArticleService(db *ent.Client, game *GameService) *ArticleService {
	svc := &ArticleService{db: db, game: game}
	// 初始化分类缓存
	svc.refreshCategoryMap(context.Background())
	return svc
}

// refreshCategoryMap 刷新分类缓存（采集完成后调用）
func (s *ArticleService) refreshCategoryMap(ctx context.Context) {
	cats, err := s.db.Category.Query().All(ctx)
	if err != nil {
		log.Printf("[ArticleService] 加载分类缓存失败: %v", err)
		return
	}
	m := make(map[string]*ent.Category, len(cats))
	for _, cat := range cats {
		m[cat.Slug] = cat
	}
	s.categoryMap = m
}

// List 获取资讯列表（分页+筛选+搜索）
func (s *ArticleService) List(ctx context.Context, q *dto.ArticleListQuery) (*dto.Paginated, error) {
	// 构建查询
	query := s.db.Article.Query()

	// 分类筛选
	if q.Category != "" && q.Category != "all" {
		query = query.Where(article.CategoryIDEQ(q.Category))
	}

	// 搜索
	if q.Search != "" {
		query = query.Where(
			article.Or(
				article.TitleContainsFold(q.Search),
				article.SummaryContainsFold(q.Search),
			),
		)
	}

	// 精选筛选
	if q.Featured {
		query = query.Where(article.IsFeatured(true))
	}

	// 排序
	switch q.Sort {
	case "hot":
		query = query.Order(ent.Desc(article.FieldViewCount))
	case "most_commented":
		query = query.Order(ent.Desc(article.FieldCommentCount))
	default:
		query = query.Order(ent.Desc(article.FieldPublishedAt))
	}

	// 时间范围筛选
	if q.TimeRange != "" && q.TimeRange != "all" {
		var since int64
		now := time.Now().Unix()
		switch q.TimeRange {
		case "24h":
			since = now - 86400
		case "week":
			since = now - 604800
		case "month":
			since = now - 2592000
		}
		if since > 0 {
			query = query.Where(article.PublishedAtGTE(since))
		}
	}

	// 计算总数
	total, err := query.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询文章总数失败: %w", err)
	}

	// 分页
	page := q.Page
	if page < 1 {
		page = 1
	}
	pageSize := q.PageSize
	if pageSize < 1 {
		pageSize = 20
	}
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	offset := (page - 1) * pageSize

	articles, err := query.
		Limit(pageSize).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询文章列表失败: %w", err)
	}

	// 转换为 VO
	items := make([]dto.ArticleVO, 0, len(articles))
	for _, a := range articles {
		items = append(items, s.toVO(a))
	}

	return &dto.Paginated{
		Items:      items,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// GetFeatured 获取精选资讯（首页用）
func (s *ArticleService) GetFeatured(ctx context.Context) ([]dto.ArticleVO, []dto.ArticleVO, error) {
	// 精选资讯（2篇大卡片）
	featured, err := s.db.Article.Query().
		Where(article.IsFeatured(true)).
		Order(ent.Desc(article.FieldPublishedAt)).
		Limit(2).
		All(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("查询精选资讯失败: %w", err)
	}

	// 最新资讯（4篇小卡片）
	latest, err := s.db.Article.Query().
		Order(ent.Desc(article.FieldPublishedAt)).
		Limit(4).
		All(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("查询最新资讯失败: %w", err)
	}

	featuredVOs := make([]dto.ArticleVO, 0, len(featured))
	for _, a := range featured {
		featuredVOs = append(featuredVOs, s.toVO(a))
	}

	latestVOs := make([]dto.ArticleVO, 0, len(latest))
	for _, a := range latest {
		latestVOs = append(latestVOs, s.toVO(a))
	}

	return featuredVOs, latestVOs, nil
}

// GetByID 获取文章详情
func (s *ArticleService) GetByID(ctx context.Context, id string) (*dto.ArticleVO, error) {
	a, err := s.db.Article.Query().
		Where(article.IDEQ(id)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("查询文章详情失败: %w", err)
	}

	vo := s.toVO(a)
	vo.Content = a.Content
	return &vo, nil
}

// IncrementView 增加浏览量
func (s *ArticleService) IncrementView(ctx context.Context, id string) error {
	_, err := s.db.Article.Update().
		Where(article.IDEQ(id)).
		AddViewCount(1).
		Save(ctx)
	return err
}

// ToggleLike 切换点赞状态
func (s *ArticleService) ToggleLike(ctx context.Context, articleID string) (bool, error) {
	// 简单实现：切换 like_count 的增减
	// 实际应用中应有 likes 表记录点赞关系
	return true, nil
}

// ToggleFavorite 切换收藏状态
func (s *ArticleService) ToggleFavorite(ctx context.Context, userID, articleID string) (bool, error) {
	// 检查是否已收藏
	exists, err := s.db.Favorite.Query().
		Where(favorite.UserIDEQ(userID), favorite.ArticleIDEQ(articleID)).
		Exist(ctx)
	if err != nil {
		return false, fmt.Errorf("查询收藏状态失败: %w", err)
	}

	if exists {
		// 取消收藏
		_, err = s.db.Favorite.Delete().
			Where(favorite.UserIDEQ(userID), favorite.ArticleIDEQ(articleID)).
			Exec(ctx)
		if err != nil {
			return false, fmt.Errorf("取消收藏失败: %w", err)
		}
		return false, nil
	}

	// 添加收藏
	_, err = s.db.Favorite.Create().
		SetUserID(userID).
		SetArticleID(articleID).
		Save(ctx)
	if err != nil {
		return false, fmt.Errorf("收藏失败: %w", err)
	}
	return true, nil
}

// RecordRead 记录阅读并增加 XP
func (s *ArticleService) RecordRead(ctx context.Context, userID, articleID string) (int, error) {
	// 检查是否已读过
	exists, err := s.db.ArticleRead.Query().
		Where(articleread.UserIDEQ(userID), articleread.ArticleIDEQ(articleID)).
		Exist(ctx)
	if err != nil {
		return 0, err
	}
	_ = exists // 暂时允许重复阅读获得 XP

	// 获取文章 XP 奖励
	a, err := s.db.Article.Query().
		Where(article.IDEQ(articleID)).
		Only(ctx)
	if err != nil {
		return 0, fmt.Errorf("文章不存在: %w", err)
	}

	// 给用户加 XP
	_, err = s.db.User.Update().
		Where(user.IDEQ(userID)).
		AddXp(a.XpReward).
		AddReadArticleCount(1).
		Save(ctx)
	if err != nil {
		return 0, fmt.Errorf("增加 XP 失败: %w", err)
	}

	return a.XpReward, nil
}

// toVO 将 Ent Article 转换为 ArticleVO
// 分类信息从缓存中查询，无需硬编码
func (s *ArticleService) toVO(a *ent.Article) dto.ArticleVO {
	// 从分类缓存中查找（兜底使用 slug 本身）
	cat, hasCat := s.categoryMap[a.CategoryID]
	catSlug := a.CategoryID
	catName := catSlug
	catColor := "#00f0ff"
	catIcon := ""
	if hasCat {
		catName = cat.DisplayName
		catColor = cat.Color
		catIcon = cat.Icon
	}

	return dto.ArticleVO{
		ID:      a.ID,
		Title:   a.Title,
		Summary: a.Summary,
		Category: dto.CategoryVO{
			Slug:        catSlug,
			Name:        catName,
			DisplayName: catName,
			Color:       catColor,
			Icon:        catIcon,
		},
		Author:       a.Author,
		PublishDate:  a.PublishedAt,
		ReadTime:     a.ReadTime,
		XPReward:     a.XpReward,
		ImageURL:     a.ImageURL,
		Tags:         a.Tags,
		IsFeatured:   a.IsFeatured,
		ViewCount:    a.ViewCount,
		CommentCount: a.CommentCount,
		LikeCount:    a.LikeCount,
		SourceName:   a.SourceName,
		CollectedAt:  a.CollectedAt,
	}
}
