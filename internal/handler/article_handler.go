// 资讯相关 HTTP 处理器

package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"isolo-news/internal/middleware"
	"isolo-news/internal/dto"
	"isolo-news/internal/service"
	"isolo-news/pkg/response"
)

// ArticleHandler 资讯处理器
type ArticleHandler struct {
	svc *service.ArticleService
}

// NewArticleHandler 创建资讯处理器
func NewArticleHandler(svc *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{svc: svc}
}

// List 获取资讯列表
// GET /api/v1/articles
func (h *ArticleHandler) List(c *gin.Context) {
	var q dto.ArticleListQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		response.ParamError(c)
		return
	}
	if q.Page < 1 {
		q.Page = 1
	}
	if q.PageSize < 1 {
		q.PageSize = 20
	}

	result, err := h.svc.List(c.Request.Context(), &q)
	if err != nil {
		response.InternalError(c)
		return
	}

	response.Success(c, result)
}

// GetFeatured 获取精选资讯（首页用）
// GET /api/v1/articles/featured
func (h *ArticleHandler) GetFeatured(c *gin.Context) {
	featured, latest, err := h.svc.GetFeatured(c.Request.Context())
	if err != nil {
		response.InternalError(c)
		return
	}

	response.Success(c, gin.H{
		"featured": featured,
		"latest":   latest,
	})
}

// GetByID 获取文章详情
// GET /api/v1/articles/:id
func (h *ArticleHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.ParamError(c)
		return
	}

	// 增加浏览量
	_ = h.svc.IncrementView(c.Request.Context(), id)

	article, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.InternalError(c)
		return
	}
	if article == nil {
		response.Error(c, http.StatusNotFound, response.CodeArticleNotFound)
		return
	}

	response.Success(c, article)
}

// GetComments 获取文章评论
// GET /api/v1/articles/:id/comments
func (h *ArticleHandler) GetComments(c *gin.Context) {
	articleID := c.Param("id")
	if articleID == "" {
		response.ParamError(c)
		return
	}

	comments, err := h.svc.GetComments(c.Request.Context(), articleID)
	if err != nil {
		response.InternalError(c)
		return
	}

	response.Success(c, comments)
}

// CreateComment 创建评论
// POST /api/v1/articles/:id/comments
func (h *ArticleHandler) CreateComment(c *gin.Context) {
	articleID := c.Param("id")
	if articleID == "" {
		response.ParamError(c)
		return
	}

	userID, _ := c.Get(middleware.CtxKeyUserID)
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}

	var req struct {
		Content  string `json:"content" binding:"required,min=1,max=1000"`
		ParentID string `json:"parentId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c)
		return
	}

	comment, err := h.svc.CreateComment(
		c.Request.Context(),
		articleID,
		userID.(string),
		req.Content,
		req.ParentID,
	)
	if err != nil {
		response.InternalError(c)
		return
	}

	c.JSON(http.StatusCreated, response.NewResponse(response.CodeSuccess, comment))
}

// ToggleLike 切换点赞
// POST /api/v1/articles/:id/like
func (h *ArticleHandler) ToggleLike(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.ParamError(c)
		return
	}

	isLiked, err := h.svc.ToggleLike(c.Request.Context(), id)
	if err != nil {
		response.InternalError(c)
		return
	}

	response.Success(c, gin.H{"isLiked": isLiked, "articleId": id})
}

// ToggleFavorite 切换收藏
// POST /api/v1/articles/:id/favorite
func (h *ArticleHandler) ToggleFavorite(c *gin.Context) {
	articleID := c.Param("id")
	if articleID == "" {
		response.ParamError(c)
		return
	}

	userID, _ := c.Get(middleware.CtxKeyUserID)
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}

	isFavorited, err := h.svc.ToggleFavorite(c.Request.Context(), userID.(string), articleID)
	if err != nil {
		response.InternalError(c)
		return
	}

	response.Success(c, gin.H{"isFavorited": isFavorited, "articleId": articleID})
}

// RecordRead 记录阅读
// POST /api/v1/articles/:id/read
func (h *ArticleHandler) RecordRead(c *gin.Context) {
	articleID := c.Param("id")
	if articleID == "" {
		response.ParamError(c)
		return
	}

	userID, _ := c.Get(middleware.CtxKeyUserID)
	if userID == "" {
		response.Error(c, http.StatusUnauthorized, response.CodeTokenInvalid)
		return
	}

	xpGained, err := h.svc.RecordRead(c.Request.Context(), userID.(string), articleID)
	if err != nil {
		response.InternalError(c)
		return
	}

	response.Success(c, gin.H{"xpGained": xpGained})
}

// GetTrendingTopics 获取热门话题
// GET /api/v1/topics/trending
func (h *ArticleHandler) GetTrendingTopics(c *gin.Context) {
	topics := []dto.TrendingTopicVO{
		{ID: 1, Name: "PS6首发", Icon: "flame", ColorType: "error"},
		{ID: 2, Name: "AI绘画革命", Icon: "bot", ColorType: "primary"},
		{ID: 3, Name: "原神5.0", Icon: "gamepad-2", ColorType: "warning"},
		{ID: 4, Name: "RTX5090评测", Icon: "cpu", ColorType: "success"},
		{ID: 5, Name: "量子芯片突破", Icon: "atom", ColorType: "secondary"},
	}
	response.Success(c, topics)
}

// GetHotTags 获取热门标签
// GET /api/v1/tags/hot
func (h *ArticleHandler) GetHotTags(c *gin.Context) {
	tags := []gin.H{
		{"name": "AI", "count": 1203},
		{"name": "VR", "count": 856},
		{"name": "显卡", "count": 412},
		{"name": "新番", "count": 376},
		{"name": "机器人", "count": 298},
		{"name": "RPG", "count": 245},
		{"name": "芯片", "count": 189},
		{"name": "漫画", "count": 156},
	}
	response.Success(c, tags)
}

// parsePagination 辅助方法：解析分页参数
func parsePagination(c *gin.Context) (int, int) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "20")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 50 {
		pageSize = 20
	}

	return page, pageSize
}
