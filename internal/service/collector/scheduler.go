// 定时调度器

package collector

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"
	"isolo-news/internal/ent"
	"isolo-news/internal/ent/source"
)

// Scheduler 定时采集调度器
type Scheduler struct {
	cron      *cron.Cron
	collector *Collector
	db        *ent.Client
	jobs      map[string]cron.EntryID // category_id -> cron entry id
}

// NewScheduler 创建调度器
func NewScheduler(collector *Collector, db *ent.Client) *Scheduler {
	return &Scheduler{
		cron:      cron.New(),
		collector: collector,
		db:        db,
		jobs:      make(map[string]cron.EntryID),
	}
}

// Start 启动所有定时任务
func (s *Scheduler) Start(ctx context.Context) error {
	// 查询所有调度配置
	schedules, err := s.db.ScheduleSetting.Query().
		All(ctx)
	if err != nil {
		return err
	}

	for _, sc := range schedules {
		if !sc.IsActive {
			continue
		}
		s.addJob(sc.CategoryID, sc.CronExpression)
	}

	// 添加默认调度（如果数据库中没有配置）
	if len(schedules) == 0 {
		s.addDefaultJobs()
	}

	s.cron.Start()
	log.Printf("[调度器] 已启动 %d 个定时任务", len(s.cron.Entries()))
	return nil
}

// Stop 停止所有定时任务
func (s *Scheduler) Stop() {
	ctx := s.cron.Stop()
	<-ctx.Done()
	log.Println("[调度器] 已停止所有定时任务")
}

// addJob 添加单个调度任务
func (s *Scheduler) addJob(categoryID, cronExpr string) {
	entryID, err := s.cron.AddFunc(cronExpr, func() {
		s.collectByCategory(categoryID)
	})
	if err != nil {
		log.Printf("[调度器] 添加任务失败 [%s] %s: %v", categoryID, cronExpr, err)
		return
	}
	s.jobs[categoryID] = entryID
	log.Printf("[调度器] 添加采集任务: 分类=%s, 周期=%s", categoryID, cronExpr)
}

// collectByCategory 按分类采集
func (s *Scheduler) collectByCategory(categoryID string) {
	ctx := context.Background()
	log.Printf("[调度器] 触发采集: 分类=%s", categoryID)

	sources, err := s.db.Source.Query().
		Where(source.IsActive(true), source.CategoryIDEQ(categoryID)).
		All(ctx)
	if err != nil {
		log.Printf("[调度器] 查询数据源失败: %v", err)
		return
	}

	for _, src := range sources {
		count, err := s.collector.CollectFromSource(ctx, src)
		if err != nil {
			log.Printf("[调度器] 采集 [%s] 失败: %v", src.Name, err)
			continue
		}
		if count > 0 {
			log.Printf("[调度器] 采集 [%s] 完成: %d 篇新文章", src.Name, count)
		}
	}

	// 采集完成后更新分类总数
	if err := s.collector.UpdateCategoryTotals(ctx); err != nil {
		log.Printf("[调度器] 更新分类总数失败: %v", err)
	}
}

// RunNow 立即执行所有活跃数据源的一次性采集
func (s *Scheduler) RunNow(ctx context.Context) {
	log.Println("[调度器] 开始首次立即采集...")
	count, err := s.collector.CollectAll(ctx)
	if err != nil {
		log.Printf("[调度器] 首次立即采集失败: %v", err)
		return
	}
	total := 0
	for _, c := range count {
		total += c
	}
	log.Printf("[调度器] 首次立即采集完成, 共新增 %d 篇文章", total)
}

// addDefaultJobs 添加默认调度任务
func (s *Scheduler) addDefaultJobs() {
	defaults := map[string]string{
		"tech":     "*/30 * * * *",  // AI 科技：每30分钟
		"ai_robot": "0 * * * *",     // AI 机器人：每小时
		"ai_coding":"*/30 * * * *",  // AI 编程：每30分钟
		"hardware": "0 */2 * * *",   // AI 硬件：每2小时
		"anime":    "0 */3 * * *",   // 二次元：每3小时
		"godot":    "0 */4 * * *",   // Godot：每4小时
	}

	for categoryID, cronExpr := range defaults {
		s.addJob(categoryID, cronExpr)
	}
}

// Refresh 刷新调度任务（配置变更后调用）
func (s *Scheduler) Refresh(ctx context.Context) error {
	// 停止所有现有任务
	for _, entryID := range s.jobs {
		s.cron.Remove(entryID)
	}
	s.jobs = make(map[string]cron.EntryID)

	// 重新加载
	return s.Start(ctx)
}

// GetStatus 获取调度状态
func (s *Scheduler) GetStatus() map[string]interface{} {
	entries := s.cron.Entries()
	status := map[string]interface{}{
		"running": s.cron.Entries() != nil,
		"job_count": len(entries),
		"jobs": make([]map[string]interface{}, 0),
	}

	jobs := status["jobs"].([]map[string]interface{})
	for _, e := range entries {
		job := map[string]interface{}{
			"id":     e.ID,
			"next":   e.Next.Format(time.RFC3339),
			"prev":   e.Prev.Format(time.RFC3339),
			"schedule": e.Schedule,
		}
		jobs = append(jobs, job)
	}

	return status
}
