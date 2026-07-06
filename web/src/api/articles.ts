/**
 * 资讯相关 API
 */

import apiClient, { type ApiResponse, type PaginatedData } from './index'

/** 资讯视图对象（与后端 model.ArticleVO 对应） */
export interface ArticleVO {
  id: string
  title: string
  summary: string
  content?: string
  category: {
    slug: string
    name: string
    color: string
  }
  author: string
  publishDate: number           // UTC Unix 时间戳
  readTime: number
  xpReward: number
  imageUrl: string
  tags: string[]
  isFeatured: boolean
  viewCount: number
  commentCount: number
  likeCount: number
  timeAgo?: string
  sourceName?: string
}

/** 分类视图对象 */
export interface CategoryVO {
  slug: string
  name: string
  displayName: string
  color: string
  articleCount: number
}

/** 评论视图对象 */
export interface CommentVO {
  id: string
  articleId: string
  author: string
  avatar: string
  content: string
  time: number              // UTC Unix 时间戳
  likes: number
  parentId?: string
}

/** 查询参数 */
export interface ArticleQuery {
  page?: number
  page_size?: number
  category?: string
  tag?: string
  sort?: 'latest' | 'hot' | 'most_commented'
  time_range?: 'all' | '24h' | 'week' | 'month'
  search?: string
  featured?: boolean
}

/** 获取资讯列表 */
export async function fetchArticles(query: ArticleQuery = {}): Promise<ApiResponse<PaginatedData<ArticleVO>>> {
  const params: Record<string, string | number | boolean> = {}
  if (query.page) params.page = query.page
  if (query.page_size) params.page_size = query.page_size
  if (query.category && query.category !== 'all') params.category = query.category
  if (query.sort) params.sort = query.sort
  if (query.time_range) params.time_range = query.time_range
  if (query.search) params.search = query.search
  if (query.featured) params.featured = true

  const res = await apiClient.get('/articles', { params })
  return res.data
}

/** 获取精选资讯（首页用） */
export async function fetchFeaturedArticles(): Promise<ApiResponse<{
  featured: ArticleVO[]
  latest: ArticleVO[]
}>> {
  const res = await apiClient.get('/articles/featured')
  return res.data
}

/** 获取文章详情 */
export async function fetchArticleById(id: string): Promise<ApiResponse<ArticleVO>> {
  const res = await apiClient.get(`/articles/${id}`)
  return res.data
}

/** 获取评论列表 */
export async function fetchComments(articleId: string): Promise<ApiResponse<CommentVO[]>> {
  const res = await apiClient.get(`/articles/${articleId}/comments`)
  return res.data
}

/** 创建评论 */
export async function createComment(articleId: string, content: string, parentId?: string): Promise<ApiResponse<CommentVO>> {
  const res = await apiClient.post(`/articles/${articleId}/comments`, { content, parentId })
  return res.data
}

/** 切换点赞 */
export async function toggleLike(articleId: string): Promise<ApiResponse<{ isLiked: boolean; articleId: string }>> {
  const res = await apiClient.post(`/articles/${articleId}/like`)
  return res.data
}

/** 切换收藏 */
export async function toggleFavorite(articleId: string): Promise<ApiResponse<{ isFavorited: boolean; articleId: string }>> {
  const res = await apiClient.post(`/articles/${articleId}/favorite`)
  return res.data
}

/** 记录阅读 */
export async function recordRead(articleId: string): Promise<ApiResponse<{ xpGained: number }>> {
  const res = await apiClient.post(`/articles/${articleId}/read`)
  return res.data
}

/** 获取分类列表 */
export async function fetchCategories(): Promise<ApiResponse<CategoryVO[]>> {
  const res = await apiClient.get('/categories')
  return res.data
}

/** 获取热门标签 */
export async function fetchHotTags(): Promise<ApiResponse<Array<{ name: string; count: number }>>> {
  const res = await apiClient.get('/tags/hot')
  return res.data
}

/** 获取热门话题 */
export async function fetchTrendingTopics(): Promise<ApiResponse<Array<{ id: number; name: string; icon: string; colorType: string }>>> {
  const res = await apiClient.get('/topics/trending')
  return res.data
}
