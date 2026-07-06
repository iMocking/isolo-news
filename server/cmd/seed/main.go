// isolo-news 采集种子命令
// 职责：初始化种子数据 + 立即从所有活跃 RSS 源采集一次，然后退出
// 用法: go run ./cmd/seed/

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"isolo-news/server/internal/config"
	"isolo-news/server/internal/repository/ent"
	"isolo-news/server/internal/repository/ent/migrate"
	"isolo-news/server/internal/seeder"
	"isolo-news/server/internal/service/collector"
)

func main() {
	ctx := context.Background()

	// 加载配置
	configPath := "config.yaml"
	if envPath := os.Getenv("ISOLO_CONFIG"); envPath != "" {
		configPath = envPath
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 连接数据库
	db, err := ent.Open("postgres", cfg.Database.DSN())
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	defer db.Close()

	// 自动迁移数据库表结构
	if err := db.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}
	log.Println("数据库迁移完成")

	// 初始化种子数据（仅首次创建，已有则跳过）
	if err := seeder.SeedData(ctx, db); err != nil {
		log.Fatalf("初始化种子数据失败: %v", err)
	}

	// 兼容旧数据库：激活所有未激活的数据源
	seeder.ActivateSources(ctx, db)

	// 初始化采集器并立即执行一次全量采集
	col := collector.NewCollector(db, &cfg.Collector)
	log.Println("开始从 RSS 数据源采集数据（可能需要一些时间）...")
	results, err := col.CollectAll(ctx)
	if err != nil {
		log.Printf("采集过程出现错误: %v", err)
	}

	// 输出采集总结
	total := 0
	for name, count := range results {
		status := "成功"
		if count == 0 {
			status = "失败/无新增"
		}
		log.Printf("  [%s] %s -> %d 篇", name, status, count)
		total += count
	}

	fmt.Println()
	log.Printf("===================================")
	log.Printf("采集完成! 共新增 %d 篇文章", total)
	log.Printf("可使用 'go run ./cmd/server/' 启动 HTTP 服务")
	log.Printf("===================================")
}
