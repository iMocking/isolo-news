// 鎴愬氨绯荤粺鍜屼换鍔¤繘搴︾鐞?
package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"isolo-news/internal/ent"
	"isolo-news/internal/ent/comment"
	"isolo-news/internal/ent/favorite"
	"isolo-news/internal/ent/quest"
	"isolo-news/internal/ent/user"
	"isolo-news/internal/ent/userachievement"
	"isolo-news/internal/ent/userquest"
)

// checkAndUnlockAchievements 检查并解锁所有符合条件的成就
func (s *GameService) checkAndUnlockAchievements(ctx context.Context, userID string) error {
	// 获取所有成就定义
	allAchs, err := s.db.Achievement.Query().All(ctx)
	if err != nil {
		return fmt.Errorf("鏌ヨ鎴愬氨瀹氫箟澶辫触: %w", err)
	}

	// 鑾峰彇鐢ㄦ埛宸茶В閿佺殑鎴愬氨
	userAchs, err := s.db.UserAchievement.Query().
		Where(userachievement.UserIDEQ(userID)).
		All(ctx)
	if err != nil {
		return fmt.Errorf("鏌ヨ鐢ㄦ埛鎴愬氨澶辫触: %w", err)
	}

	unlockedMap := make(map[string]bool)
	for _, ua := range userAchs {
		unlockedMap[ua.AchievementID] = true
	}

	// 鑾峰彇鐢ㄦ埛缁熻鏁版嵁
	u, err := s.db.User.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("鏌ヨ鐢ㄦ埛澶辫触: %w", err)
	}

	commentCount, _ := s.db.Comment.Query().
		Where(comment.UserIDEQ(userID)).
		Count(ctx)
	favCount, _ := s.db.Favorite.Query().
		Where(favorite.UserIDEQ(userID)).
		Count(ctx)

	now := time.Now().Unix()

	for _, ach := range allAchs {
		if unlockedMap[ach.ID] {
			continue
		}

		met := false
		switch ach.ConditionType {
		case "read_articles":
			met = u.ReadArticleCount >= ach.ConditionValue
		case "login_days":
			met = u.LoginDays >= ach.ConditionValue
		case "comments":
			met = commentCount >= ach.ConditionValue
		case "favorites":
			met = favCount >= ach.ConditionValue
		default:
			log.Printf("[Game] 鏈煡鎴愬氨鏉′欢绫诲瀷: %s", ach.ConditionType)
			continue
		}

		if met {
			_, err := s.db.UserAchievement.Create().
				SetUserID(userID).
				SetAchievementID(ach.ID).
				SetUnlockedAt(now).
				Save(ctx)
			if err != nil {
				log.Printf("[Game] 瑙ｉ攣鎴愬氨澶辫触 [%s]: %v", ach.Name, err)
				continue
			}
			log.Printf("[Game] 馃弳 鐢ㄦ埛 %s 瑙ｉ攣鎴愬氨: %s", u.PlayerName, ach.Name)
		}
	}

	return nil
}

// updateQuestProgress 鏇存柊鐢ㄦ埛姣忔棩浠诲姟杩涘害
func (s *GameService) updateQuestProgress(ctx context.Context, userID, questType string) error {
	// 获取当天日期时间戳
	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Unix()

	// 鏌ヨ鐢ㄦ埛褰撳ぉ鐨勮绫诲瀷浠诲姟杩涘害
	userQ, err := s.db.UserQuest.Query().
		Where(
			userquest.UserIDEQ(userID),
			userquest.DateAt(today),
		).
		All(ctx)
	if err != nil {
		return fmt.Errorf("鏌ヨ浠诲姟杩涘害澶辫触: %w", err)
	}

	// 鑾峰彇鎵€鏈夊尮閰嶇殑浠诲姟瀹氫箟
	quests, err := s.db.Quest.Query().
		Where(quest.QuestTypeEQ(questType)).
		All(ctx)
	if err != nil {
		return fmt.Errorf("鏌ヨ浠诲姟瀹氫箟澶辫触: %w", err)
	}

	// 寤虹珛 quest_id -> progress 鏄犲皠
	progressMap := make(map[string]*ent.UserQuest)
	for _, uq := range userQ {
		progressMap[uq.QuestID] = uq
	}

	for _, q := range quests {
		if pq, ok := progressMap[q.ID]; ok {
			// 已有进度记录，更新
			newProgress := pq.Progress + 1
			newStatus := "in_progress"
			if newProgress >= q.Target {
				newStatus = "completed"
				// 浠诲姟瀹屾垚鏃跺彂鏀鹃澶?XP
				_, _ = s.db.User.Update().
					Where(user.IDEQ(userID)).
					AddXp(q.XpReward).
					Save(ctx)
				log.Printf("[Game] 鉁?鐢ㄦ埛浠诲姟瀹屾垚: %s, 濂栧姳 %d XP", q.Title, q.XpReward)
			}

			_, err = s.db.UserQuest.Update().
				Where(userquest.IDEQ(pq.ID)).
				SetProgress(newProgress).
				SetStatus(newStatus).
				Save(ctx)
			if err != nil {
				log.Printf("[Game] 鏇存柊浠诲姟杩涘害澶辫触: %v", err)
			}
		} else {
			// 没有进度记录，新建
			status := "in_progress"
			if 1 >= q.Target {
				status = "completed"
				_, _ = s.db.User.Update().
					Where(user.IDEQ(userID)).
					AddXp(q.XpReward).
					Save(ctx)
				log.Printf("[Game] 鉁?鐢ㄦ埛浠诲姟瀹屾垚: %s, 濂栧姳 %d XP", q.Title, q.XpReward)
			}

			_, err = s.db.UserQuest.Create().
				SetUserID(userID).
				SetQuestID(q.ID).
				SetProgress(1).
				SetStatus(status).
				SetDateAt(today).
				Save(ctx)
			if err != nil {
				log.Printf("[Game] 鍒涘缓浠诲姟杩涘害澶辫触: %v", err)
			}
		}
	}

	return nil
}

// ProcessCommentReward 处理评论奖励：更新任务进度、检查成就
func (s *GameService) ProcessCommentReward(ctx context.Context, userID string) error {
	if err := s.updateQuestProgress(ctx, userID, "comment"); err != nil {
		return fmt.Errorf("鏇存柊浠诲姟杩涘害澶辫触: %w", err)
	}
	if err := s.checkAndUnlockAchievements(ctx, userID); err != nil {
		return fmt.Errorf("妫€鏌ユ垚灏卞け璐? %w", err)
	}
	return nil
}
