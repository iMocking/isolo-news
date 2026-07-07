// 鏂囩珷璇勮涓氬姟閫昏緫鍜岀儹闂ㄦ爣绛?
package service

import (
	"context"
	"fmt"

	"isolo-news/internal/dto"
	"isolo-news/internal/ent"
	"isolo-news/internal/ent/article"
	"isolo-news/internal/ent/comment"
	"isolo-news/internal/ent/user"
)

// GetHotTags 鑾峰彇鐑棬鏍囩锛堜粠鏁版嵁搴撹仛鍚堬級
func (s *ArticleService) GetHotTags(ctx context.Context) ([]map[string]interface{}, error) {
	articles, err := s.db.Article.Query().
		Order(ent.Desc(article.FieldViewCount)).
		Limit(100).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("鏌ヨ鏂囩珷澶辫触: %w", err)
	}

	tagCount := make(map[string]int)
	for _, a := range articles {
		for _, tag := range a.Tags {
			tagCount[tag]++
		}
	}

	// 鎺掑簭鍙栧墠 10
	type tagItem struct {
		name  string
		count int
	}
	sorted := make([]tagItem, 0, len(tagCount))
	for name, count := range tagCount {
		sorted = append(sorted, tagItem{name, count})
	}
	for i := 0; i < len(sorted) && i < 10; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j].count > sorted[i].count {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	limit := 10
	if len(sorted) < limit {
		limit = len(sorted)
	}

	result := make([]map[string]interface{}, 0, limit)
	for i := 0; i < limit; i++ {
		result = append(result, map[string]interface{}{
			"name":  sorted[i].name,
			"count": sorted[i].count,
		})
	}
	return result, nil
}

// GetComments 鑾峰彇鏂囩珷璇勮
func (s *ArticleService) GetComments(ctx context.Context, articleID string) ([]dto.CommentVO, error) {
	comments, err := s.db.Comment.Query().
		Where(comment.ArticleIDEQ(articleID)).
		Order(ent.Desc(comment.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("鏌ヨ璇勮澶辫触: %w", err)
	}

	vos := make([]dto.CommentVO, 0, len(comments))
	for _, c := range comments {
		u, err := s.db.User.Query().
			Where(user.IDEQ(c.UserID)).
			Only(ctx)
		if err != nil {
			continue
		}

		vos = append(vos, dto.CommentVO{
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

// CreateComment 鍒涘缓璇勮
func (s *ArticleService) CreateComment(ctx context.Context, articleID, userID, content, parentID string) (*dto.CommentVO, error) {
	// 楠岃瘉鏂囩珷瀛樺湪
	exists, err := s.db.Article.Query().
		Where(article.IDEQ(articleID)).
		Exist(ctx)
	if err != nil || !exists {
		return nil, fmt.Errorf("文章不存在")
	}

	// 鍒涘缓璇勮
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
		return nil, fmt.Errorf("鍒涘缓璇勮澶辫触: %w", err)
	}

	// 增加文章评论数
	_, _ = s.db.Article.UpdateOneID(articleID).
		AddCommentCount(1).
		Save(ctx)

	// 鑾峰彇鐢ㄦ埛淇℃伅
	u, err := s.db.User.Query().
		Where(user.IDEQ(userID)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("鏌ヨ鐢ㄦ埛澶辫触: %w", err)
	}

	// 澶勭悊璇勮濂栧姳锛堟垚灏便€佷换鍔★級
	if err := s.game.ProcessCommentReward(ctx, userID); err != nil {
		return nil, fmt.Errorf("澶勭悊璇勮濂栧姳澶辫触: %w", err)
	}

	return &dto.CommentVO{
		ID:        c.ID,
		ArticleID: c.ArticleID,
		Author:    u.PlayerName,
		Avatar:    string([]rune(u.PlayerName)[0]),
		Content:   c.Content,
		Time:      c.CreatedAt,
		Likes:     c.Likes,
	}, nil
}
