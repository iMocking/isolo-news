// Casbin 权限引擎初始化

package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"isolo-news/server/internal/config"
)

// initCasbin 初始化 Casbin 权限引擎
func initCasbin(cfg *config.Config) (*casbin.Enforcer, error) {
	// 从文件加载模型
	m, err := model.NewModelFromFile(cfg.Casbin.ModelPath)
	if err != nil {
		return nil, fmt.Errorf("加载 Casbin 模型失败: %w", err)
	}

	// 使用内存适配器（开发阶段）
	e, err := casbin.NewEnforcer(m)
	if err != nil {
		return nil, fmt.Errorf("创建 Casbin 执行器失败: %w", err)
	}

	// 手动添加策略
	addPolicies(e)

	return e, nil
}

// addPolicies 添加 Casbin 策略
func addPolicies(e *casbin.Enforcer) {
	type policy struct {
		role string
		path string
		act  string
	}

	policies := []policy{
		{"admin", "/api/v1/*", "GET|POST|PUT|DELETE"},
		{"editor", "/api/v1/articles", "GET"},
		{"editor", "/api/v1/articles/*", "GET|POST|PUT"},
		{"editor", "/api/v1/categories", "GET|POST|PUT"},
		{"editor", "/api/v1/sources", "GET"},
		{"editor", "/api/v1/schedules", "GET"},
		{"user", "/api/v1/articles", "GET"},
		{"user", "/api/v1/articles/featured", "GET"},
		{"user", "/api/v1/articles/*", "GET"},
		{"user", "/api/v1/articles/*/comments", "GET|POST"},
		{"user", "/api/v1/articles/*/like", "POST"},
		{"user", "/api/v1/articles/*/favorite", "POST"},
		{"user", "/api/v1/articles/*/read", "POST"},
		{"user", "/api/v1/user/*", "GET|PUT"},
		{"guest", "/api/v1/articles", "GET"},
		{"guest", "/api/v1/articles/featured", "GET"},
		{"guest", "/api/v1/articles/*", "GET"},
		{"guest", "/api/v1/articles/*/comments", "GET"},
		{"guest", "/api/v1/categories", "GET"},
		{"guest", "/api/v1/categories/*", "GET"},
		{"guest", "/api/v1/tags/hot", "GET"},
		{"guest", "/api/v1/topics/trending", "GET"},
		{"guest", "/api/v1/leaderboard/*", "GET"},
		{"guest", "/api/v1/auth/*", "POST"},
	}

	for _, p := range policies {
		_, _ = e.AddPolicy(p.role, p.path, p.act)
	}

	_, _ = e.AddRoleForUser("user", "guest")
	_, _ = e.AddRoleForUser("editor", "user")
	_, _ = e.AddRoleForUser("admin", "editor")
}
