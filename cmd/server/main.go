// isolo-news 后端服务入口
// 职责：启动 HTTP 服务 + 定时采集任务（不做立即采集）
// 立即采集请使用: go run ./cmd/seed/

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL 驱动注册
	"isolo-news/internal/casbin"
	"isolo-news/internal/config"
	"isolo-news/internal/ent"
	"isolo-news/internal/ent/migrate"
	"isolo-news/internal/handler"
	"isolo-news/internal/router"
	"isolo-news/internal/seeder"
	"isolo-news/internal/service"
	"isolo-news/internal/service/collector"
)

func main() {

	ctx := context.Background()

	// 加载配置（默认读取 configs/config.yaml，也支持 .env 覆盖）
	configPath := "configs/config.yaml"
	if envPath := os.Getenv("ISOLO_CONFIG"); envPath != "" {
		configPath = envPath
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 设置 Gin 模式
	gin.SetMode(cfg.Server.Mode)

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
		log.Printf("初始化种子数据: %v", err)
	}

	// 初始化 Casbin 权限引擎
	enforcer, err := casbin.InitCasbin(cfg)
	if err != nil {
		log.Fatalf("初始化 Casbin 失败: %v", err)
	}

	// 初始化游戏化服务（XP、等级、成就、任务统一管理）
	gameSvc := service.NewGameService(db)

	// 初始化业务服务
	articleSvc := service.NewArticleService(db, gameSvc)
	emailSvc := service.NewEmailService(&cfg.SMTP)
	authSvc := service.NewAuthService(db, cfg, emailSvc)
	userSvc := service.NewUserService(db)

	// 初始化采集系统（启动定时任务，不做立即采集）
	col := collector.NewCollector(db, &cfg.Collector)
	scheduler := collector.NewScheduler(col, db)
	if err := scheduler.Start(ctx); err != nil {
		log.Printf("启动采集调度器失败: %v", err)
	}

	// 初始化 HTTP 处理器
	articleHandler := handler.NewArticleHandler(articleSvc)
	authHandler := handler.NewAuthHandler(authSvc)
	userHandler := handler.NewUserHandler(userSvc)
	categoryHandler := handler.NewCategoryHandler()

	// 创建 Gin 引擎并注册路由
	r := gin.New()
	router.Setup(
		r,
		enforcer,
		cfg.Auth.JWTSecret,
		articleHandler,
		authHandler,
		userHandler,
		categoryHandler,
	)

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 启动服务前先探测端口是否可用
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := probePort(addr); err != nil {
		log.Fatalf("端口 %s 已被占用，无法启动服务。\n"+
			"  解决方法:\n"+
			"    1. 停止旧服务后再启动: go run ./cmd/server/\n"+
			"    2. 单独执行采集: go run ./cmd/seed/\n"+
			"    3. 强制停止旧进程: taskkill /F /IM server.exe", addr)
	}

	log.Printf("isolo-news 后端服务启动于 %s", addr)
	go func() {
		if err := r.Run(addr); err != nil {
			log.Fatalf("服务启动失败: %v", err)
		}
	}()

	// 等待关闭信号
	<-quit
	log.Println("正在关闭服务...")
	scheduler.Stop()
}

// probePort 探测端口是否可用
func probePort(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("端口 %s 不可用: %w", addr, err)
	}
	ln.Close()
	return nil
}
