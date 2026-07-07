// 用户业务逻辑

package service

import (
	"context"
	"fmt"

	"isolo-news/internal/dto"
	"isolo-news/internal/ent"
	"isolo-news/internal/ent/article"
	"isolo-news/internal/ent/favorite"
	"isolo-news/internal/ent/user"
	"isolo-news/internal/ent/userachievement"
	"isolo-news/internal/ent/userquest"
)

// UserService 用户服务
type UserService struct {
	db *ent.Client
}

// NewUserService 创建用户服务
func NewUserService(db *ent.Client) *UserService {
	return &UserService{db: db}
}

// GetProfile 获取用户信息
func (s *UserService) GetProfile(ctx context.Context, userID string) (*dto.UserVO, error) {
	u, err := s.db.User.Query().
		Where(user.IDEQ(userID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	// 获取成就数
	achCount, err := s.db.UserAchievement.Query().
		Where(userachievement.UserIDEQ(userID)).
		Count(ctx)
	if err != nil {
		achCount = 0
	}

	return &dto.UserVO{
		ID:               u.ID,
		Email:            u.Email,
		PlayerName:       u.PlayerName,
		Level:            u.Level,
		XP:               u.Xp,
		MaxXP:            u.MaxXp,
		Title:            u.Title,
		Tags:             u.Tags,
		ReadArticles:     u.ReadArticleCount,
		LoginDays:        u.LoginDays,
		AvatarURL:        u.AvatarURL,
		ThemePreference:  u.ThemePreference,
		AchievementCount: achCount,
	}, nil
}

// UpdateProfile 更新用户信息
func (s *UserService) UpdateProfile(ctx context.Context, userID string, updates map[string]interface{}) (*dto.UserVO, error) {
	updateQuery := s.db.User.Update().Where(user.IDEQ(userID))

	if name, ok := updates["player_name"]; ok {
		updateQuery = updateQuery.SetPlayerName(name.(string))
	}
	if title, ok := updates["title"]; ok {
		updateQuery = updateQuery.SetTitle(title.(string))
	}
	if tags, ok := updates["tags"]; ok {
		updateQuery = updateQuery.SetTags(tags.([]string))
	}
	if theme, ok := updates["theme_preference"]; ok {
		updateQuery = updateQuery.SetThemePreference(theme.(string))
	}
	if avatar, ok := updates["avatar_url"]; ok {
		updateQuery = updateQuery.SetAvatarURL(avatar.(string))
	}

	_, err := updateQuery.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("更新用户失败: %w", err)
	}

	return s.GetProfile(ctx, userID)
}

// GetFavorites 获取用户收藏列表
func (s *UserService) GetFavorites(ctx context.Context, userID string) ([]dto.ArticleVO, error) {
	favs, err := s.db.Favorite.Query().
		Where(favorite.UserIDEQ(userID)).
		Order(ent.Desc(favorite.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询收藏失败: %w", err)
	}

	vos := make([]dto.ArticleVO, 0, len(favs))
	for _, f := range favs {
		a, err := s.db.Article.Query().
			Where(article.IDEQ(f.ArticleID)).
			Only(ctx)
		if err != nil {
			continue
		}
		vos = append(vos, dto.ArticleVO{
			ID:        a.ID,
			Title:     a.Title,
			Summary:   a.Summary,
			Author:    a.Author,
			PublishDate: a.PublishedAt,
			ReadTime:  a.ReadTime,
			XPReward:  a.XpReward,
			ImageURL:  a.ImageURL,
			Tags:      a.Tags,
		})
	}
	return vos, nil
}

// GetAchievements 获取用户成就列表
func (s *UserService) GetAchievements(ctx context.Context, userID string) ([]dto.AchievementVO, error) {
	// 获取用户已解锁的成就
	userAchs, err := s.db.UserAchievement.Query().
		Where(userachievement.UserIDEQ(userID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询成就失败: %w", err)
	}

	unlockedMap := make(map[string]int64)
	for _, ua := range userAchs {
		unlockedMap[ua.AchievementID] = ua.UnlockedAt
	}

	// 获取所有成就定义
	allAchs, err := s.db.Achievement.Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询成就定义失败: %w", err)
	}

	vos := make([]dto.AchievementVO, 0, len(allAchs))
	for _, a := range allAchs {
		vo := dto.AchievementVO{
			ID:          a.ID,
			Name:        a.Name,
			Description: a.Description,
			Icon:        a.Icon,
		}
		if unlockedAt, ok := unlockedMap[a.ID]; ok {
			vo.UnlockedAt = unlockedAt
		}
		vos = append(vos, vo)
	}
	return vos, nil
}

// GetQuests 获取用户每日任务
func (s *UserService) GetQuests(ctx context.Context, userID string) ([]dto.QuestVO, error) {
	// 获取所有任务定义
	allQuests, err := s.db.Quest.Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询任务定义失败: %w", err)
	}

	// 获取用户今日任务进度
	userQuests, err := s.db.UserQuest.Query().
		Where(userquest.UserIDEQ(userID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询任务进度失败: %w", err)
	}

	progressMap := make(map[string]*ent.UserQuest)
	for _, uq := range userQuests {
		progressMap[uq.QuestID] = uq
	}

	vos := make([]dto.QuestVO, 0, len(allQuests))
	for _, q := range allQuests {
		vo := dto.QuestVO{
			ID:          q.ID,
			Title:       q.Title,
			Description: q.Description,
			XPReward:    q.XpReward,
			Target:      q.Target,
			Status:      "not_started",
			Progress:    0,
			QuestType:   q.QuestType,
		}

		if pq, ok := progressMap[q.ID]; ok {
			vo.Progress = pq.Progress
			vo.Status = pq.Status
		}

		vos = append(vos, vo)
	}
	return vos, nil
}

// GetStats 获取用户统计数据（用于 PlayerStats 组件）
func (s *UserService) GetStats(ctx context.Context, userID string) (map[string]interface{}, error) {
	u, err := s.db.User.Query().
		Where(user.IDEQ(userID)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	achCount, _ := s.db.UserAchievement.Query().
		Where(userachievement.UserIDEQ(userID)).
		Count(ctx)

	return map[string]interface{}{
		"readArticles":     u.ReadArticleCount,
		"achievementCount": achCount,
		"loginDays":        u.LoginDays,
		"level":            u.Level,
		"xp":               u.Xp,
		"maxXp":            u.MaxXp,
		"title":            u.Title,
		"playerName":       u.PlayerName,
	}, nil
}
