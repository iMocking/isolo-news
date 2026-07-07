// 种子数据管理包，同时被 cmd/server 和 cmd/seed 使用

package seeder

import (
	"context"
	"fmt"
	"log"
	"time"

	"isolo-news/internal/ent"
	"isolo-news/internal/ent/source"
)

// SeedData 初始化种子数据（分类、数据源、调度配置、成就、任务、管理员）
func SeedData(ctx context.Context, db *ent.Client) error {
	now := time.Now().Unix()

	// 检查是否已有分类数据
	count, err := db.Category.Query().Count(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("种子数据已存在，跳过")
		return nil
	}

	log.Println("初始化种子数据...")

	// 创建分类
	categories := []struct {
		slug        string
		name        string
		displayName string
		color       string
		icon        string
		sortOrder   int
	}{
		{"tech", "AI前沿", "AI前沿", "#00f0ff", "bot", 1},
		{"ai_robot", "AI机器人", "AI机器人", "#4ecdc4", "cpu", 2},
		{"ai_coding", "AI编程", "AI编程", "#3b9eff", "terminal", 3},
		{"hardware", "AI硬件", "AI硬件", "#0aaaff", "hard-drive", 4},
		{"anime", "二次元", "二次元", "#ffaa00", "tv", 5},
		{"godot", "Godot游戏", "Godot游戏", "#ff6b2b", "gamepad-2", 6},
	}

	for _, c := range categories {
		_, err := db.Category.Create().
			SetSlug(c.slug).
			SetName(c.name).
			SetDisplayName(c.displayName).
			SetColor(c.color).
			SetIcon(c.icon).
			SetSortOrder(c.sortOrder).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("创建分类 %s 失败: %w", c.slug, err)
		}
	}
	log.Printf("已创建 %d 个分类", len(categories))

	// 创建默认 RSS 数据源
	sources := []struct {
		name       string
		feedURL    string
		siteURL    string
		categoryID string
	}{
		// 英文源
		{"Hacker News", "https://hnrss.org/frontpage", "https://news.ycombinator.com", "tech"},
		{"TechCrunch AI", "https://techcrunch.com/category/artificial-intelligence/feed/", "https://techcrunch.com", "tech"},
		{"MIT Tech Review - AI", "https://www.technologyreview.com/topic/artificial-intelligence/feed/", "https://www.technologyreview.com", "tech"},
		{"IEEE Spectrum Robotics", "https://spectrum.ieee.org/feeds/topic/robotics/", "https://spectrum.ieee.org/robotics", "ai_robot"},
		{"Dev.to AI", "https://dev.to/feed/tag/ai", "https://dev.to", "ai_coding"},
		{"GitHub Trending", "https://github.com/trending.rss", "https://github.com/trending", "ai_coding"},
		{"Tom's Hardware", "https://www.tomshardware.com/feeds/all", "https://www.tomshardware.com", "hardware"},
		{"Anime News Network", "https://www.animenewsnetwork.com/newsroom/otakumagazine.xml", "https://www.animenewsnetwork.com", "anime"},
		{"Godot Blog", "https://godotengine.org/blog/rss.xml", "https://godotengine.org", "godot"},
		{"Reddit r/godot", "https://www.reddit.com/r/godot/.rss", "https://reddit.com/r/godot", "godot"},
		// 中文源
		{"36氪", "https://36kr.com/feed", "https://36kr.com", "tech"},
		{"Solidot", "https://www.solidot.org/index.rss", "https://www.solidot.org", "tech"},
		{"机器之心", "https://www.jiqizhixin.com/rss", "https://www.jiqizhixin.com", "tech"},
		{"掘金", "https://juejin.cn/rss", "https://juejin.cn", "ai_coding"},
		{"超能网", "https://www.expreview.com/feed", "https://www.expreview.com", "hardware"},
		{"机核网", "https://www.gcores.com/rss", "https://www.gcores.com", "anime"},
	}

	for _, s := range sources {
		_, err := db.Source.Create().
			SetName(s.name).
			SetFeedURL(s.feedURL).
			SetSiteURL(s.siteURL).
			SetCategoryID(s.categoryID).
			SetIsActive(true).
			SetFetchInterval(30).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("创建数据源 %s 失败: %w", s.name, err)
		}
	}
	log.Printf("已创建 %d 个数据源", len(sources))

	// 创建默认调度配置
	schedules := map[string]string{
		"tech":      "*/30 * * * *",
		"ai_robot":  "0 * * * *",
		"ai_coding": "*/30 * * * *",
		"hardware":  "0 */2 * * *",
		"anime":     "0 */3 * * *",
		"godot":     "0 */4 * * *",
	}
	for categoryID, cronExpr := range schedules {
		_, err := db.ScheduleSetting.Create().
			SetCategoryID(categoryID).
			SetCronExpression(cronExpr).
			SetIsActive(true).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("创建调度配置 %s 失败: %w", categoryID, err)
		}
	}
	log.Printf("已创建 %d 个调度配置", len(schedules))

	// 创建预设成就
	achievements := []struct {
		name        string
		description string
		icon        string
		condType    string
		condValue   int
	}{
		{"初次阅读", "完成第一篇文章阅读", "book", "read_articles", 1},
		{"连续登录7天", "连续登录7天", "calendar", "login_days", 7},
		{"科技爱好者", "阅读10篇科技资讯", "cpu", "read_articles", 10},
		{"评论之星", "发表100条评论", "message-square", "comments", 100},
		{"收藏家", "收藏50篇情报", "bookmark", "favorites", 50},
	}
	for _, a := range achievements {
		_, err := db.Achievement.Create().
			SetName(a.name).
			SetDescription(a.description).
			SetIcon(a.icon).
			SetConditionType(a.condType).
			SetConditionValue(a.condValue).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("创建成就 %s 失败: %w", a.name, err)
		}
	}
	log.Printf("已创建 %d 个成就", len(achievements))

	// 创建预设每日任务
	quests := []struct {
		title       string
		description string
		xpReward    int
		target      int
		questType   string
	}{
		{"阅读3篇资讯", "今日阅读任务", 300, 3, "read_articles"},
		{"完成每日签到", "每日签到奖励", 50, 1, "login"},
		{"评论一篇评测", "互动任务", 200, 1, "comment"},
	}
	for _, q := range quests {
		_, err := db.Quest.Create().
			SetTitle(q.title).
			SetDescription(q.description).
			SetXpReward(q.xpReward).
			SetTarget(q.target).
			SetQuestType(q.questType).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("创建任务 %s 失败: %w", q.title, err)
		}
	}
	log.Printf("已创建 %d 个任务", len(quests))

	// 创建示例管理员用户
	adminEmail := "admin@isolo.news"
	adminExists, err := db.User.Query().Count(ctx)
	if err == nil && adminExists == 0 {
		_, err := db.User.Create().
			SetEmail(adminEmail).
			SetPasswordHash("$2a$12$LJ3m4ys3Lk0TSwHnbfOMiOXPm1Qm0LqFQq0LqFQq0LqFQq0LqFQq").
			SetPlayerName("NEXUS管理员").
			SetLevel(99).
			SetXp(999999).
			SetMaxXp(999999).
			SetTitle("系统管理员 // 全知者").
			SetTags([]string{"管理员", "全知者", "维护者"}).
			SetEmailVerified(true).
			SetLastLoginAt(now).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("创建管理员失败: %w", err)
		}
		log.Printf("已创建管理员账号: %s", adminEmail)
	}

	log.Println("种子数据初始化完成")
	return nil
}

// ActivateSources 将已存在但未激活的数据源全量激活（兼容旧数据库升级）
func ActivateSources(ctx context.Context, db *ent.Client) {
	count, err := db.Source.Update().
		Where(source.IsActive(false)).
		SetIsActive(true).
		Save(ctx)
	if err != nil {
		log.Printf("激活数据源失败: %v", err)
		return
	}
	if count > 0 {
		log.Printf("已激活 %d 个旧数据源", count)
	}
}

// AddMissingSources 添加种子数据中尚未入库的数据源（兼容已有安装）
// 每次 seed 或 server 启动时调用，确保中文等新增源自动补充
func AddMissingSources(ctx context.Context, db *ent.Client) {
	newSources := []struct {
		name       string
		feedURL    string
		siteURL    string
		categoryID string
	}{
		// 中文源（与 SeedData 保持同步）
		{"36氪", "https://36kr.com/feed", "https://36kr.com", "tech"},
		{"Solidot", "https://www.solidot.org/index.rss", "https://www.solidot.org", "tech"},
		{"机器之心", "https://www.jiqizhixin.com/rss", "https://www.jiqizhixin.com", "tech"},
		{"掘金", "https://juejin.cn/rss", "https://juejin.cn", "ai_coding"},
		{"超能网", "https://www.expreview.com/feed", "https://www.expreview.com", "hardware"},
		{"机核网", "https://www.gcores.com/rss", "https://www.gcores.com", "anime"},
	}

	added := 0
	for _, s := range newSources {
		exists, err := db.Source.Query().
			Where(source.FeedURLEQ(s.feedURL)).
			Exist(ctx)
		if err != nil {
			log.Printf("[种子] 检查数据源 %s 失败: %v", s.name, err)
			continue
		}
		if exists {
			continue
		}

		_, err = db.Source.Create().
			SetName(s.name).
			SetFeedURL(s.feedURL).
			SetSiteURL(s.siteURL).
			SetCategoryID(s.categoryID).
			SetIsActive(true).
			SetFetchInterval(30).
			Save(ctx)
		if err != nil {
			log.Printf("[种子] 添加数据源 %s 失败: %v", s.name, err)
			continue
		}
		added++
		log.Printf("[种子] 新增数据源: %s (%s)", s.name, s.categoryID)
	}

	if added > 0 {
		log.Printf("[种子] 新增 %d 个数据源", added)
	}
}
