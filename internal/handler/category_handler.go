// 分类相关 HTTP 处理器

package handler

import (
	"github.com/gin-gonic/gin"
	"isolo-news/internal/dto"
	"isolo-news/internal/ent"
	"isolo-news/internal/ent/article"
	"isolo-news/internal/errcode"
	"isolo-news/internal/response"
)

// CategoryHandler 分类处理器
type CategoryHandler struct {
	db *ent.Client
}

// NewCategoryHandler 创建分类处理器
func NewCategoryHandler(db *ent.Client) *CategoryHandler {
	return &CategoryHandler{db: db}
}

// GetCategories 获取分类列表
// 从数据库读取所有分类，并实时计算每类的文章数
// GET /api/v1/categories
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	ctx := c.Request.Context()

	// 查询所有分类
	categories, err := h.db.Category.Query().
		Order(ent.Asc("sort_order")).
		All(ctx)
	if err != nil {
		response.Error(c, 500, errcode.CodeInternalError)
		return
	}

	// 转换为 VO，实时计算文章数
	result := make([]dto.CategoryVO, 0, len(categories)+1)

	// 统计全站文章总数
	totalCount, _ := h.db.Article.Query().Count(ctx)

	// 添加"全部"虚拟分类
	result = append(result, dto.CategoryVO{
		Slug:         "all",
		Name:         "全部",
		DisplayName:  "全部",
		Color:        "#00f0ff",
		ArticleCount: totalCount,
	})

	// 逐个统计每类文章数
	for _, cat := range categories {
		articleCount, _ := h.db.Article.Query().
			Where(article.CategoryIDEQ(cat.Slug)).
			Count(ctx)

		result = append(result, dto.CategoryVO{
			Slug:         cat.Slug,
			Name:         cat.Name,
			DisplayName:  cat.DisplayName,
			Color:        cat.Color,
			Icon:         cat.Icon,
			ArticleCount: articleCount,
		})
	}

	response.Success(c, result)
}
