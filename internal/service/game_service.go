// 娓告垙鍖栫郴缁熶笟鍔￠€昏緫锛圶P銆佺瓑绾ф牳蹇冮€昏緫锛?// 鎴愬氨鍜屼换鍔￠€昏緫鍦?game_achievement.go 涓?
package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"isolo-news/internal/ent"
	"isolo-news/internal/ent/articleread"
	"isolo-news/internal/ent/user"
)

// 绛夌骇绯荤粺甯搁噺
const (
	baseMaxXP     = 2000 // 1 绾у崌 2 绾ф墍闇€缁忛獙
	levelUpGap    = 500  // 每升一级递增的经验值
	titleLevelGap = 10   // 每10级更换一次称号前缀
)

// 绉板彿妯℃澘锛氭寜绛夌骇鍖洪棿鑷姩鐢熸垚
var titleTemplates = map[int]string{
	1:  "璧涘崥瀛﹀緬",
	10: "璧勮鐚庝汉",
	20: "数据行者",
	30: "代码修士",
	40: "AI 先驱",
	50: "量子领航员",
}

// GameService 游戏化服务
type GameService struct {
	db *ent.Client
}

// NewGameService 创建游戏化服务
func NewGameService(db *ent.Client) *GameService {
	return &GameService{db: db}
}

// calculateDynamicXP 鏍规嵁闃呰鏃堕棿锛堝垎閽燂級鍔ㄦ€佽绠?XP 濂栧姳
func calculateDynamicXP(readTime int) int {
	switch {
	case readTime <= 0:
		return 0
	case readTime <= 1:
		return 50
	case readTime <= 3:
		return 100
	case readTime <= 8:
		return 200
	case readTime <= 15:
		return 300
	default:
		return 500
	}
}

// calculateMaxXP 计算指定等级所需的经验上限
func calculateMaxXP(level int) int {
	return baseMaxXP + (level-1)*levelUpGap
}

// generateTitle 鏍规嵁绛夌骇鑷姩鐢熸垚绉板彿
func generateTitle(level int) string {
	selected := "璧涘崥鍏堥攱"
	highest := 0
	for lvl, title := range titleTemplates {
		if level >= lvl && lvl > highest {
			selected = title
			highest = lvl
		}
	}
	rank := (level-1)/titleLevelGap + 1
	suffix := ""
	switch {
	case rank == 1:
		suffix = " // 鍒濆嚭鑼呭簮"
	case rank == 2:
		suffix = " // 灏忔湁鍚嶆皵"
	case rank == 3:
		suffix = " // 瀹炲姏骞插皢"
	case rank == 4:
		suffix = " // 璧勬繁涓撳"
	case rank >= 5:
		suffix = " // 浼犲鍏堥攱"
	}
	return selected + suffix
}

// ProcessReadReward 澶勭悊闃呰濂栧姳锛氳褰曢槄璇汇€佸彂鏀?XP銆佹鏌ュ崌绾с€佹洿鏂版垚灏卞拰浠诲姟
// 杩斿洖鍊? xpGained锛堟湰娆¤幏寰楃殑 XP锛? leveledUp锛堟槸鍚﹀崌绾э級, newLevel锛堟柊绛夌骇锛? error
func (s *GameService) ProcessReadReward(ctx context.Context, userID, articleID string) (int, bool, int, error) {
	// 检查是否已读过该文章（防止重复获得 XP）
	alreadyRead, err := s.db.ArticleRead.Query().
		Where(articleread.UserIDEQ(userID), articleread.ArticleIDEQ(articleID)).
		Exist(ctx)
	if err != nil {
		return 0, false, 0, fmt.Errorf("鏌ヨ闃呰璁板綍澶辫触: %w", err)
	}
	if alreadyRead {
		return 0, false, 0, nil
	}

	// 获取文章阅读时长
	a, err := s.db.Article.Get(ctx, articleID)
	if err != nil {
		return 0, false, 0, fmt.Errorf("文章不存在: %w", err)
	}

	// 鏍规嵁闃呰鏃堕暱鍔ㄦ€佽绠?XP
	xpGained := calculateDynamicXP(a.ReadTime)

	// 璁板綍闃呰
	_, err = s.db.ArticleRead.Create().
		SetUserID(userID).
		SetArticleID(articleID).
		SetReadAt(time.Now().Unix()).
		Save(ctx)
	if err != nil {
		return 0, false, 0, fmt.Errorf("璁板綍闃呰澶辫触: %w", err)
	}

	// 增加用户 XP 和阅读计数
	_, err = s.db.User.UpdateOneID(userID).
		AddXp(xpGained).
		AddReadArticleCount(1).
		Save(ctx)
	if err != nil {
		return 0, false, 0, fmt.Errorf("鏇存柊鐢ㄦ埛 XP 澶辫触: %w", err)
	}

	// 閲嶆柊鏌ヨ鐢ㄦ埛淇℃伅
	u, err := s.db.User.Get(ctx, userID)
	if err != nil {
		return xpGained, false, 0, fmt.Errorf("鏌ヨ鐢ㄦ埛淇℃伅澶辫触: %w", err)
	}

	// 检查是否升级
	leveledUp := false
	for u.Xp >= calculateMaxXP(u.Level) {
		u.Xp -= calculateMaxXP(u.Level)
		u.Level++
		leveledUp = true
		newTitle := generateTitle(u.Level)

		_, err = s.db.User.Update().
			Where(user.IDEQ(userID)).
			SetLevel(u.Level).
			SetXp(u.Xp).
			SetMaxXp(calculateMaxXP(u.Level)).
			SetTitle(newTitle).
			Save(ctx)
		if err != nil {
			return xpGained, leveledUp, u.Level, fmt.Errorf("鍗囩骇澶辫触: %w", err)
		}
	}

	// 鏈崌绾ф椂涔熸洿鏂?max_xp
	if !leveledUp {
		maxXP := calculateMaxXP(u.Level)
		_, _ = s.db.User.Update().
			Where(user.IDEQ(userID)).
			SetMaxXp(maxXP).
			Save(ctx)
	}

	// 检查成就（非关键路径，失败只记日志）
	if err := s.checkAndUnlockAchievements(ctx, userID); err != nil {
		log.Printf("[Game] 妫€鏌ユ垚灏卞け璐? %v", err)
	}

	// 鏇存柊浠诲姟杩涘害
	if err := s.updateQuestProgress(ctx, userID, "read_articles"); err != nil {
		log.Printf("[Game] 鏇存柊浠诲姟杩涘害澶辫触: %v", err)
	}

	return xpGained, leveledUp, u.Level, nil
}

// CheckAndUpgrade 手动检查用户是否需要升级（供登录或其他场景调用）
func (s *GameService) CheckAndUpgrade(ctx context.Context, userID string) (bool, int, error) {
	u, err := s.db.User.Get(ctx, userID)
	if err != nil {
		return false, 0, fmt.Errorf("鏌ヨ鐢ㄦ埛澶辫触: %w", err)
	}

	leveledUp := false
	for u.Xp >= calculateMaxXP(u.Level) {
		u.Xp -= calculateMaxXP(u.Level)
		u.Level++
		leveledUp = true
		newTitle := generateTitle(u.Level)

		_, err = s.db.User.Update().
			Where(user.IDEQ(userID)).
			SetLevel(u.Level).
			SetXp(u.Xp).
			SetMaxXp(calculateMaxXP(u.Level)).
			SetTitle(newTitle).
			Save(ctx)
		if err != nil {
			return leveledUp, u.Level, fmt.Errorf("鍗囩骇澶辫触: %w", err)
		}
	}

	if !leveledUp {
		_, _ = s.db.User.Update().
			Where(user.IDEQ(userID)).
			SetMaxXp(calculateMaxXP(u.Level)).
			Save(ctx)
	}

	u, _ = s.db.User.Get(ctx, userID)
	return leveledUp, u.Level, nil
}

// ProcessLoginReward 澶勭悊鐧诲綍濂栧姳锛氭洿鏂扮櫥褰曞ぉ鏁般€佺鍒颁换鍔″拰鎴愬氨
func (s *GameService) ProcessLoginReward(ctx context.Context, userID string) error {
	u, err := s.db.User.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("鏌ヨ鐢ㄦ埛澶辫触: %w", err)
	}

	// 璁＄畻杩炵画鐧诲綍澶╂暟
	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Unix()
	yesterday := today - 86400

	newLoginDays := u.LoginDays
	if u.LastLoginAt >= yesterday && u.LastLoginAt < today {
		newLoginDays++ // 鏄ㄥぉ鐧诲綍杩囷紝杩炵画澶╂暟+1
	} else if u.LastLoginAt < yesterday {
		newLoginDays = 1 // 瓒呰繃涓€澶╂病鐧诲綍锛岄噸缃?	}

	// 更新登录天数和最后登录时间
	_, err = s.db.User.UpdateOneID(userID).
		SetLoginDays(newLoginDays).
		SetLastLoginAt(now.Unix()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("鏇存柊鐧诲綍淇℃伅澶辫触: %w", err)
	}

	// 鏇存柊绛惧埌浠诲姟杩涘害
	if err := s.updateQuestProgress(ctx, userID, "login"); err != nil {
		log.Printf("[Game] 鏇存柊绛惧埌浠诲姟杩涘害澶辫触: %v", err)
	}

	// 妫€鏌ユ垚灏?	if err := s.checkAndUnlockAchievements(ctx, userID); err != nil {
		log.Printf("[Game] 妫€鏌ユ垚灏卞け璐? %v", err)
	}

	return nil
}


