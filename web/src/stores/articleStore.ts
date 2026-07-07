/**
 * 资讯状态管理
 * 所有数据均从 API 获取
 */

import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { ArticleVO } from '@/api/articles'
import { fetchArticles, fetchCategories, fetchTrendingTopics } from '@/api/articles'

function stripHtml(html: string): string {
  if (!html) return ''
  const tmp = document.createElement('div')
  tmp.innerHTML = html
  let text = tmp.textContent || tmp.innerText || ''
  text = text.replace(/\{\{.*?\}\}/g, '')
  text = text.replace(/\$json\./g, '')
  text = text.replace(/&nbsp;/g, ' ')
  text = text.replace(/\s+/g, ' ')
  return text.trim()
}

function processArticle(article: ArticleVO): ArticleVO {
  const processed = { ...article }
  if (processed.summary) processed.summary = stripHtml(processed.summary)
  if (processed.summary && processed.summary.length > 120) {
    processed.summary = processed.summary.substring(0, 120) + '...'
  }
  if (!processed.summary || processed.summary.length < 5) {
    processed.summary = article.title
  }
  return processed
}

export const useArticleStore = defineStore('article', () => {
  // 状态
  const articles = ref<ArticleVO[]>([])
  const currentCategory = ref<string>('all')
  const searchQuery = ref('')
  const total = ref(0)
  const currentPage = ref(1)
  const pageSize = ref(20)
  const totalPages = ref(0)
  const loading = ref(false)
  const categories = ref<Array<{ id: string; name: string; count: number }>>([])
  const trendingTopics = ref<Array<{ id: number; name: string; icon: string; colorType: string }>>([])
  const apiOffline = ref(false)

  /** 初始化加载 */
  async function initialize() {
    await Promise.all([
      loadArticles(),
      loadCategories(),
    ])
  }

  /** 加载文章列表 */
  async function loadArticles(params?: {
    page?: number
    page_size?: number
    category?: string
    search?: string
    sort?: string
  }) {
    loading.value = true
    try {
      apiOffline.value = false
      // 使用显式 undefined 判断，让 params 中的值能正确覆盖 store 值
      const cat = params && params.category !== undefined
        ? (params.category || undefined)
        : (currentCategory.value === 'all' ? undefined : currentCategory.value)
      const q = params && params.search !== undefined
        ? (params.search || undefined)
        : (searchQuery.value || undefined)
      const pg = params?.page || currentPage.value
      const ps = params?.page_size || pageSize.value
      const res = await fetchArticles({
        page: pg,
        page_size: ps,
        category: cat,
        search: q,
        sort: (params?.sort as 'latest' | 'hot' | 'most_commented') || 'latest',
      })
      if (res.code === 0 && res.data) {
        articles.value = res.data.items.map((article) => processArticle(article))
        total.value = res.data.total
        currentPage.value = res.data.page
        pageSize.value = res.data.page_size
        totalPages.value = res.data.total_pages
      } else if (res.code !== 0) {
        apiOffline.value = true
        articles.value = []
        total.value = 0
      }
    } catch (e) {
      console.warn('[articleStore] 文章列表API异常:', e)
      apiOffline.value = true
      articles.value = []
      total.value = 0
    } finally {
      loading.value = false
    }
  }

  /** 加载分类 */
  async function loadCategories() {
    try {
      apiOffline.value = false
      const res = await fetchCategories()
      if (res.code === 0 && res.data) {
        categories.value = res.data.map((c) => ({
          id: c.slug,
          name: c.displayName,
          count: c.articleCount,
        }))
      } else if (res.code !== 0) {
        apiOffline.value = true
        categories.value = []
      }
    } catch (e) {
      console.warn('[articleStore] 分类API异常:', e)
      apiOffline.value = true
      categories.value = []
    }
  }

  /** 加载热门话题 */
  async function loadTrendingTopics() {
    try {
      apiOffline.value = false
      const res = await fetchTrendingTopics()
      if (res.code === 0 && res.data) {
        trendingTopics.value = res.data.map((t) => ({
          id: t.id,
          name: t.name,
          icon: t.icon,
          colorType: t.colorType,
        }))
      } else if (res.code !== 0) {
        apiOffline.value = true
        trendingTopics.value = []
      }
    } catch (e) {
      console.warn('[articleStore] 话题API异常:', e)
      apiOffline.value = true
      trendingTopics.value = []
    }
  }

  /** 重置 API 离线状态 */
  function resetApiState() {
    apiOffline.value = false
  }

  /** 过滤后的文章 */
    return {
    // 状态
    articles,
    currentCategory,
    searchQuery,
    total,
    currentPage,
    pageSize,
    totalPages,
    loading,
    categories,
    trendingTopics,
    apiOffline,
    // 方法
    initialize,
    loadArticles,
    loadCategories,
    loadTrendingTopics,
    resetApiState,
  }
})
