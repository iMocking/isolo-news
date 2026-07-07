// 用户相关 HTTP 处理器

package handler

import (
	"github.com/gin-gonic/gin"
	"isolo-news/internal/middleware"
	"isolo-news/internal/service"
	"isolo-news/pkg/response"
)

// UserHandler 用户处理器
type UserHandler struct {
	svc *service.UserService
}

// NewUserHandler 创建用户处理器
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// GetProfile 获取当前用户信息
// GET /api/v1/user/profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.CtxKeyUserID)
	if userID == "" {
		response.Error(c, 401, response.CodeTokenInvalid)
		return
	}

	user, err := h.svc.GetProfile(c.Request.Context(), userID.(string))
	if err != nil {
		response.InternalError(c)
		return
	}
	if user == nil {
		response.Error(c, 404, response.CodeUserNotFound)
		return
	}

	response.Success(c, user)
}

// UpdateProfile 更新用户信息
// PUT /api/v1/user/profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, _ := c.Get(middleware.CtxKeyUserID)
	if userID == "" {
		response.Error(c, 401, response.CodeTokenInvalid)
		return
	}

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c)
		return
	}

	user, err := h.svc.UpdateProfile(c.Request.Context(), userID.(string), req)
	if err != nil {
		response.InternalError(c)
		return
	}

	response.Success(c, user)
}

// GetFavorites 获取收藏列表
// GET /api/v1/user/favorites
func (h *UserHandler) GetFavorites(c *gin.Context) {
	userID, _ := c.Get(middleware.CtxKeyUserID)
	if userID == "" {
		response.Error(c, 401, response.CodeTokenInvalid)
		return
	}

	favorites, err := h.svc.GetFavorites(c.Request.Context(), userID.(string))
	if err != nil {
		response.InternalError(c)
		return
	}

	response.Success(c, favorites)
}

// GetAchievements 获取成就列表
// GET /api/v1/user/achievements
func (h *UserHandler) GetAchievements(c *gin.Context) {
	userID, _ := c.Get(middleware.CtxKeyUserID)
	if userID == "" {
		response.Error(c, 401, response.CodeTokenInvalid)
		return
	}

	achievements, err := h.svc.GetAchievements(c.Request.Context(), userID.(string))
	if err != nil {
		response.InternalError(c)
		return
	}

	response.Success(c, achievements)
}

// GetQuests 获取每日任务
// GET /api/v1/user/quests
func (h *UserHandler) GetQuests(c *gin.Context) {
	userID, _ := c.Get(middleware.CtxKeyUserID)
	if userID == "" {
		response.Error(c, 401, response.CodeTokenInvalid)
		return
	}

	quests, err := h.svc.GetQuests(c.Request.Context(), userID.(string))
	if err != nil {
		response.InternalError(c)
		return
	}

	response.Success(c, quests)
}

// GetStats 获取用户统计数据
// GET /api/v1/user/stats
func (h *UserHandler) GetStats(c *gin.Context) {
	userID, _ := c.Get(middleware.CtxKeyUserID)
	if userID == "" {
		response.Error(c, 401, response.CodeTokenInvalid)
		return
	}

	stats, err := h.svc.GetStats(c.Request.Context(), userID.(string))
	if err != nil {
		response.InternalError(c)
		return
	}

	response.Success(c, stats)
}

// GetLeaderboard 获取周排行
// GET /api/v1/leaderboard/weekly
func (h *UserHandler) GetLeaderboard(c *gin.Context) {
	// 模拟数据
	leaderboard := []gin.H{
		{"title": "苹果Vision Pro 2深度评测", "views": "12.3K"},
		{"title": "RTX 5090对比测试", "views": "9.8K"},
		{"title": "GPT-5开发者体验报告", "views": "8.1K"},
		{"title": "夜行者之剑Boss攻略", "views": "6.5K"},
		{"title": "Switch 2首发游戏清单", "views": "5.2K"},
	}
	response.Success(c, leaderboard)
}
