// 分类相关 HTTP 处理器

package handler

import (
	"github.com/gin-gonic/gin"
	"isolo-news/internal/dto"
	"isolo-news/pkg/response"
)

// CategoryHandler 分类处理器
type CategoryHandler struct{}

// NewCategoryHandler 创建分类处理器
func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

// GetCategories 获取分类列表
// GET /api/v1/categories
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	// 返回预设分类列表（后续从数据库读取）
	categories := []dto.CategoryVO{
		{Slug: "all", Name: "全部", DisplayName: "全部", Color: "#00f0ff", ArticleCount: 2847},
		{Slug: "tech", Name: "AI前沿", DisplayName: "AI前沿", Color: "#00f0ff", ArticleCount: 1203},
		{Slug: "ai_robot", Name: "AI机器人", DisplayName: "AI机器人", Color: "#4ecdc4", ArticleCount: 456},
		{Slug: "ai_coding", Name: "AI编程", DisplayName: "AI编程", Color: "#3b9eff", ArticleCount: 342},
		{Slug: "hardware", Name: "AI硬件", DisplayName: "AI硬件", Color: "#0aaaff", ArticleCount: 412},
		{Slug: "anime", Name: "二次元", DisplayName: "二次元", Color: "#ffaa00", ArticleCount: 376},
		{Slug: "godot", Name: "Godot游戏", DisplayName: "Godot游戏", Color: "#ff6b2b", ArticleCount: 158},
	}

	response.Success(c, categories)
}
