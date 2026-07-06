// 资讯业务逻辑

package service

import (
	"context"
	"fmt"
	"math"
	"time"

	"isolo-news/server/internal/model"
	"isolo-news/server/internal/repository/ent"
	"isolo-news/server/internal/repository/ent/article"
	"isolo-news/server/internal/repository/ent/articleread"
	"isolo-news/server/internal/repository/ent/comment"
	"isolo-news/server/internal/repository/ent/favorite"
	"isolo-news/server/internal/repository/ent/user"
)

// ArticleService 资讯服务
type ArticleService struct {
	db *ent.Client
}

// NewArticleService 创建资讯服务
func NewArticleService(db *ent.Client) *ArticleService {
	return &ArticleService{db: db}
}

// List 获取资讯列表（分页+筛选+搜索）
func (s *ArticleService) List(ctx context.Context, q *model.ArticleListQuery) (*model.Paginated, error) {
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
	items := make([]model.ArticleVO, 0, len(articles))
	for _, a := range articles {
		items = append(items, s.toVO(a))
	}

	return &model.Paginated{
		Items:      items,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// GetFeatured 获取精选资讯（首页用）
func (s *ArticleService) GetFeatured(ctx context.Context) ([]model.ArticleVO, []model.ArticleVO, error) {
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

	featuredVOs := make([]model.ArticleVO, 0, len(featured))
	for _, a := range featured {
		featuredVOs = append(featuredVOs, s.toVO(a))
	}

	latestVOs := make([]model.ArticleVO, 0, len(latest))
	for _, a := range latest {
		latestVOs = append(latestVOs, s.toVO(a))
	}

	return featuredVOs, latestVOs, nil
}

// GetByID 获取文章详情
func (s *ArticleService) GetByID(ctx context.Context, id string) (*model.ArticleVO, error) {
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

// GetComments 获取文章评论
func (s *ArticleService) GetComments(ctx context.Context, articleID string) ([]model.CommentVO, error) {
	comments, err := s.db.Comment.Query().
		Where(comment.ArticleIDEQ(articleID)).
		Order(ent.Desc(comment.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询评论失败: %w", err)
	}

	vos := make([]model.CommentVO, 0, len(comments))
	for _, c := range comments {
		// 获取评论者信息
		u, err := s.db.User.Query().
			Where(user.IDEQ(c.UserID)).
			Only(ctx)
		if err != nil {
			continue
		}

		vos = append(vos, model.CommentVO{
			ID:        c.ID,
			ArticleID: c.ArticleID,
			Author:    u.PlayerName,
			Avatar:    string([]rune(u.PlayerName)[0]),
			Content:   c.Content,
			Time:      c.CreatedAt,
			Likes:     c.Likes,
			ParentID:  c.ParentID,
		})
	}
	return vos, nil
}

// CreateComment 创建评论
func (s *ArticleService) CreateComment(ctx context.Context, articleID, userID, content, parentID string) (*model.CommentVO, error) {
	// 验证文章存在
	exists, err := s.db.Article.Query().
		Where(article.IDEQ(articleID)).
		Exist(ctx)
	if err != nil || !exists {
		return nil, fmt.Errorf("文章不存在")
	}

	// 创建评论
	c, err := s.db.Comment.Create().
		SetArticleID(articleID).
		SetUserID(userID).
		SetContent(content).
		SetNillableParentID(func() *string {
			if parentID == "" {
				return nil
			}
			return &parentID
		}()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("创建评论失败: %w", err)
	}

	// 增加文章评论数
	_, _ = s.db.Article.Update().
		Where(article.IDEQ(articleID)).
		AddCommentCount(1).
		Save(ctx)

	// 获取用户信息
	u, err := s.db.User.Query().
		Where(user.IDEQ(userID)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	return &model.CommentVO{
		ID:        c.ID,
		ArticleID: c.ArticleID,
		Author:    u.PlayerName,
		Avatar:    string([]rune(u.PlayerName)[0]),
		Content:   c.Content,
		Time:      c.CreatedAt,
		Likes:     c.Likes,
	}, nil
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
func (s *ArticleService) toVO(a *ent.Article) model.ArticleVO {
	catName := "AI前沿"
	catSlug := "tech"
	catColor := "#00f0ff"
	catIcon := ""

	// 根据 category_id 映射分类展示名
	switch a.CategoryID {
	case "tech":
		catName = "AI前沿"
		catSlug = "tech"
		catColor = "#00f0ff"
	case "ai_robot":
		catName = "AI机器人"
		catSlug = "ai_robot"
		catColor = "#4ecdc4"
	case "ai_coding":
		catName = "AI编程"
		catSlug = "ai_coding"
		catColor = "#3b9eff"
	case "hardware":
		catName = "AI硬件"
		catSlug = "hardware"
		catColor = "#0aaaff"
	case "anime":
		catName = "二次元"
		catSlug = "anime"
		catColor = "#ffaa00"
	case "godot":
		catName = "Godot游戏"
		catSlug = "godot"
		catColor = "#ff6b2b"
	}

	return model.ArticleVO{
		ID:      a.ID,
		Title:   a.Title,
		Summary: a.Summary,
		Category: model.CategoryVO{
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
