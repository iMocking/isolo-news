// 路由注册与中间件装配

package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"isolo-news/internal/handler"
	"isolo-news/internal/middleware"
)

// Setup 配置所有路由
func Setup(
	r *gin.Engine,
	enforcer *casbin.Enforcer,
	jwtSecret string,
	articleHandler *handler.ArticleHandler,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	categoryHandler *handler.CategoryHandler,
) {
	// 全局中间件
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "isolo-news-server"})
	})

	// API v1 路由组
	v1 := r.Group("/api/v1")

	// ===== 公开接口（无需认证） =====
	public := v1.Group("")
	public.Use(middleware.AuthOptional(jwtSecret), middleware.CasbinAuth(enforcer))
	{
		// 资讯
		public.GET("/articles", articleHandler.List)
		public.GET("/articles/featured", articleHandler.GetFeatured)
		public.GET("/articles/:id", articleHandler.GetByID)
		public.GET("/articles/:id/comments", articleHandler.GetComments)

		// 分类
		public.GET("/categories", categoryHandler.GetCategories)

		// 标签与话题
		public.GET("/tags/hot", articleHandler.GetHotTags)
		public.GET("/topics/trending", articleHandler.GetTrendingTopics)

		// 排行榜
		public.GET("/leaderboard/weekly", userHandler.GetLeaderboard)

		// 认证
		public.POST("/auth/login", authHandler.Login)
		public.POST("/auth/register", authHandler.Register)
		public.POST("/auth/email/code", authHandler.SendVerifyCode)
		public.POST("/auth/email/verify", authHandler.VerifyEmail)
		public.POST("/auth/refresh", authHandler.RefreshToken)
	}

	// ===== 用户接口（需登录） =====
	user := v1.Group("")
	user.Use(middleware.AuthRequired(jwtSecret), middleware.CasbinAuth(enforcer))
	{
		// 资讯交互
		user.POST("/articles/:id/comments", articleHandler.CreateComment)
		user.POST("/articles/:id/like", articleHandler.ToggleLike)
		user.POST("/articles/:id/favorite", articleHandler.ToggleFavorite)
		user.POST("/articles/:id/read", articleHandler.RecordRead)

		// 用户信息
		user.GET("/user/profile", userHandler.GetProfile)
		user.PUT("/user/profile", userHandler.UpdateProfile)
		user.GET("/user/favorites", userHandler.GetFavorites)
		user.GET("/user/achievements", userHandler.GetAchievements)
		user.GET("/user/quests", userHandler.GetQuests)
		user.GET("/user/stats", userHandler.GetStats)
	}
}
