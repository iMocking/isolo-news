<template>
  <div>
    <!-- Scanline overlay for NEXUS theme -->
    <div v-if="themeStore.currentTheme === 'nexus'" class="fixed inset-0 pointer-events-none z-50" style="background: repeating-linear-gradient(0deg, transparent, transparent 2px, rgba(0,240,255,0.08) 2px, rgba(0,240,255,0.08) 4px);"></div>
    
    <!-- Page Header -->
    <section class="pt-24 pb-6 px-6">
      <div class="max-w-7xl mx-auto">
        <h1 class="text-2xl md:text-4xl tracking-wider" style="font-family: var(--font-display); color: var(--color-primary); word-break: keep-all;">{{ $t('articleList.title') }} // {{ $t('articleList.subtitle') }}</h1>
        <p class="mt-3 text-sm" style="font-family: var(--font-mono); color: var(--color-text-secondary);">{{ $t('articleList.total', { count: articleStore.total }) }}</p>
        
        <!-- Search bar -->
        <div class="mt-5 max-w-xl relative">
          <div class="flex items-center px-4 h-12" style="background: var(--color-bg-card); border: 1px solid var(--color-border); border-radius: var(--radius-md);">
            <span class="text-sm mr-3" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">&gt;_</span>
            <input 
              type="text" 
              v-model="searchQuery"
              :placeholder="$t('articleList.searchPlaceholder')" 
              class="flex-1 bg-transparent outline-none text-sm" 
              style="color: var(--color-text-primary); font-family: var(--font-mono); caret-color: var(--color-primary);"
            />
            <div class="w-0.5 h-5 animate-pulse" style="background: var(--color-primary);"></div>
          </div>
        </div>
      </div>
    </section>

    <!-- Category Filter Bar -->
    <section class="px-6 pb-4">
      <div class="max-w-7xl mx-auto">
        <div class="flex gap-2 overflow-x-auto no-scrollbar py-1">
          <button 
            v-for="cat in categories" 
            :key="cat.id"
            @click="currentCategory = cat.id"
            class="shrink-0 px-4 h-9 text-sm font-medium transition-all duration-150 whitespace-nowrap"
            :style="{
              background: currentCategory === cat.id ? 'var(--color-primary-200)' : 'transparent',
              border: currentCategory === cat.id ? '1px solid var(--color-primary)' : '1px solid var(--color-border-subtle)',
              color: currentCategory === cat.id ? 'var(--color-primary)' : 'var(--color-text-secondary)',
              borderRadius: 'var(--radius-md)',
              fontFamily: 'var(--font-mono)',
              boxShadow: currentCategory === cat.id ? '0 0 12px rgba(0,240,255,0.2)' : 'none'
            }"
          >
            {{ getCategoryName(cat.id) }} <span class="ml-1 text-xs" :style="{
              background: currentCategory === cat.id ? 'var(--color-primary)' : 'var(--color-bg-elevated)',
              color: currentCategory === cat.id ? '#ffffff' : 'var(--color-text-tertiary)',
              padding: '1px 6px',
              borderRadius: 'var(--radius-full)'
            }">{{ cat.count }}</span>
          </button>
        </div>
      </div>
    </section>

    <!-- Main Content Area: List + Sidebar -->
    <section class="px-6 pb-12">
      <div class="max-w-7xl mx-auto flex gap-6">
        <!-- Article List -->
        <div class="flex-1 min-w-0 space-y-3">
          <!-- 文章卡片循环 -->
          <template v-if="articleStore.articles.length > 0">
            <article 
              v-for="(article, index) in articleStore.articles" 
              :key="article.id"
              class="flex items-stretch gap-0 p-5 transition-all duration-150 hover:translate-y-[-1px] cursor-pointer"
              :style="{
                ...articleCardStyle,
                borderLeft: `3px solid ${getCategoryColor(article.category)}`
              }"
              @click="router.push(`/article/detail/${article.id}`)"
            >
              <div class="shrink-0 w-12 flex items-center justify-center mr-5">
                <span class="text-2xl font-bold" :style="{ fontFamily: 'var(--font-display)', color: getCategoryColor(article.category) }">
                  {{ String(index + 1).padStart(2, '0') }}
                </span>
              </div>
              <div class="flex-1 min-w-0">
                <div class="flex items-center gap-2 mb-2">
                <span class="px-2 py-0.5 text-xs font-medium whitespace-nowrap" :style="{
                  background: getCategoryBg(article.category),
                  color: getCategoryColor(article.category),
                  borderRadius: 'var(--radius-sm)',
                  fontFamily: 'var(--font-mono)'
                }">{{ getCategoryName(article.category) }}</span>
                <!-- 已读角标 -->
                <span v-if="isRead(article.id)" class="shrink-0 px-1.5 py-0.5 text-[10px] font-bold whitespace-nowrap" style="background: var(--color-primary); color: white; border-radius: var(--radius-sm); font-family: var(--font-mono);">{{ t('articleList.read') }}</span>
                <h2 class="text-base font-semibold truncate" style="color: var(--color-text-primary);">{{ article.title }}</h2>
                </div>
                <p class="text-sm line-clamp-2 mb-3" style="color: var(--color-text-secondary); line-height: 1.6;">{{ article.summary }}</p>
                <div class="flex items-center gap-3">
                  <div class="w-5 h-5 flex items-center justify-center" style="background: var(--color-bg-elevated); border-radius: var(--radius-full);">
                    <span class="text-[10px] font-bold" :style="{ color: getCategoryColor(article.category) }">{{ article.author.charAt(0).toUpperCase() }}</span>
                  </div>
                  <span class="text-xs" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">{{ article.author }} · {{ formatTime(article.publishDate) }} · {{ $t('articleList.comments', { count: article.commentCount || 0 }) }}</span>
                </div>
              </div>
              <div class="shrink-0 ml-5 flex flex-col items-center gap-3 justify-center">
                <div class="px-3 py-1.5 text-xs font-bold whitespace-nowrap" :style="{
                  background: getCategoryBg(article.category),
                  border: `1px solid ${getCategoryBgLight(article.category)}`,
                  color: getCategoryColor(article.category),
                  borderRadius: 'var(--radius-sm)',
                  fontFamily: 'var(--font-display)'
                }">+{{ article.xpReward }} XP</div>
                <div class="flex gap-1">
                  <span v-for="i in 5" :key="i" class="w-2 h-2" :style="{
                    background: i <= Math.min(5, Math.floor(article.xpReward / 100)) ? getCategoryColor(article.category) : 'var(--color-bg-elevated)',
                    borderRadius: 'var(--radius-full)'
                  }"></span>
                </div>
                <button class="px-3 py-1.5 text-xs font-medium whitespace-nowrap transition-all duration-150 hover:opacity-80" :style="{
                  background: getCategoryColor(article.category),
                  color: 'var(--color-text-inverse)',
                  borderRadius: 'var(--radius-sm)',
                  fontFamily: 'var(--font-mono)'
                }">{{ $t('articleList.startMission') }}</button>
              </div>
            </article>
          </template>

          <!-- 空数据/离线提示 -->
          <EmptyState
            v-else-if="!articleStore.loading"
            :type="articleStore.apiOffline ? 'offline' : (searchQuery ? 'search-empty' : 'empty')"
          />

          <!-- 分页组件（含每页条数选择） -->
          <ListPagination />
        </div>

        <ArticleFilterSidebar 
          :categories="categories" 
          :currentCategory="currentCategory" 
          @update:currentCategory="currentCategory = $event"
          :hotTags="hotTags" 
          :leaderboard="leaderboard" 
        />
      </div>
    </section>

  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useThemeStore } from '@/stores/themeStore'
import { useArticleStore } from '@/stores/articleStore'
import { fetchHotTags, fetchTrendingTopics } from '@/api/articles'
import { fetchLeaderboard } from '@/api/user'
import ArticleFilterSidebar from './ArticleFilterSidebar.vue'
import { useCardStyles } from '@/hooks/useCardStyles'
import ListPagination from './ListPagination.vue'
import EmptyState from '@/components/base/EmptyState.vue'
import { useReadTracker } from '@/composables/useReadTracker'

const router = useRouter()
const route = useRoute()
const themeStore = useThemeStore()
const articleStore = useArticleStore()
const { getCardStyle } = useCardStyles()
const { t } = useI18n()

const searchQuery = ref('')
const currentCategory = ref<string>((route.query.category as string) || 'all')

/** ======== 状态持久化：从 sessionStorage 恢复 ======== */
const LIST_STATE_KEY = 'article_list_state'
function restoreListState() {
  try {
    const saved = sessionStorage.getItem(LIST_STATE_KEY)
    if (!saved) return
    const state = JSON.parse(saved)
    // URL query 参数优先于持久化状态
    if (!route.query.category && state.category) currentCategory.value = state.category
    if (state.search) searchQuery.value = state.search
    if (state.page) articleStore.currentPage = state.page
    if (state.pageSize) articleStore.pageSize = state.pageSize
  } catch { /* ignore */ }
}
restoreListState()

/** 保存当前筛选条件 + 分页状态到 sessionStorage */
function saveListState() {
  try {
    sessionStorage.setItem(LIST_STATE_KEY, JSON.stringify({
      category: currentCategory.value,
      search: searchQuery.value,
      page: articleStore.currentPage,
      pageSize: articleStore.pageSize,
    }))
  } catch { /* ignore */ }
}

/** 监听分页变化自动持久化 */
watch(() => articleStore.currentPage, saveListState)
watch(() => articleStore.pageSize, saveListState)

/** ======== 分类切换 → 重新从 API 加载 ======== */
watch(currentCategory, (cat) => {
  articleStore.currentCategory = cat
  articleStore.currentPage = 1
  saveListState()
  articleStore.loadArticles({ category: cat === 'all' ? undefined : cat, page: 1 })
})

/** ======== 搜索输入（防抖 400ms）→ 重新从 API 加载 ======== */
let searchTimer: ReturnType<typeof setTimeout> | undefined
watch(searchQuery, (q) => {
  if (searchTimer) clearTimeout(searchTimer)
  searchTimer = setTimeout(() => {
    articleStore.currentPage = 1
    articleStore.searchQuery = q
    saveListState()
    articleStore.loadArticles({ search: q || undefined, page: 1 })
  }, 400)
})

const { isRead } = useReadTracker()
const articleCardStyle = computed(() => getCardStyle('article', false))

/** 格式化时间戳 */
const formatTime = (ts: number | string): string => {
  if (!ts) return ''
  if (typeof ts === 'string') return ts
  const now = Date.now() / 1000
  const diff = now - ts
  if (diff < 60) return t('common.justNow')
  if (diff < 3600) return t('common.minutes', { count: Math.floor(diff / 60) })
  if (diff < 86400) return t('common.hours', { count: Math.floor(diff / 3600) })
  if (diff < 2592000) return t('common.days', { count: Math.floor(diff / 86400) })
  return new Date(ts * 1000).toLocaleDateString()
}

/** 从 store 获取分类列表（去重「全部」） */
const categories = computed(() => {
  const cats = articleStore.categories
    .filter(c => c.id !== 'all')
    .map(c => ({ id: c.id, name: c.name, count: c.count }))
  return [{ id: 'all', name: t('articleList.categories.all'), count: articleStore.total || 0 }, ...cats]
})

const hotTags = ref<Array<{ name: string; bg: string; border: string; color: string }>>([])
const leaderboard = ref<Array<{ title: string; views: string }>>([])

const defaultTagStyles = {
  bg: 'var(--color-primary-50)',
  border: '1px solid var(--color-primary-100)',
  color: 'var(--color-primary)'
}

onMounted(async () => {
  // 进入页面时用恢复的状态加载文章（含搜索词和分类参数）
  articleStore.currentCategory = currentCategory.value
  const initialParams: Record<string, any> = { page: articleStore.currentPage, page_size: articleStore.pageSize }
  if (searchQuery.value) {
    initialParams.search = searchQuery.value
    articleStore.searchQuery = searchQuery.value
  }
  articleStore.loadArticles(initialParams)
  // 刷新分类列表（含实时计数）
  articleStore.loadCategories()

  // 加载热门标签
  try {
    const res = await fetchHotTags()
    if (res.code === 0 && res.data) {
      hotTags.value = res.data.map(t => ({
        name: t.name,
        ...defaultTagStyles
      }))
    }
  } catch {
    hotTags.value = []
  }

  // 加载排行榜
  try {
    const res = await fetchLeaderboard()
    if (res.code === 0 && res.data) {
      leaderboard.value = res.data
    }
  } catch {
    leaderboard.value = []
  }
})

/** 获取分类 slug */
const getCatSlug = (cat: any): string =>
  typeof cat === 'string' ? cat : cat?.slug || ''

/** 获取分类名称（通过 i18n 适配多语言） */
const getCategoryName = (cat: any): string => {
  const slug = typeof cat === 'string' ? cat : cat?.slug || ''
  // "全部"特例：后端无 all 分类，直接返回翻译
  if (slug === 'all') return t('articleList.categories.all')
  // 尝试从 i18n 获取翻译，key 路径：articleList.categories.{slug}
  const key = `articleList.categories.${slug}`
  const translated = t(key)
  // vue-i18n 无匹配时返回 key 本身，此时回退到 API 返回的 name
  return translated !== key ? translated : (cat?.name || slug)
}

const getCategoryColor = (cat: any) => {
  const slug = getCatSlug(cat)
  const colors: Record<string, string> = {
    tech: 'var(--color-primary)',
    game: 'var(--color-secondary)',
    hardware: 'var(--color-state-info)',
    anime: 'var(--color-state-warning)',
    ai_robot: 'var(--color-primary)',
    ai_coding: 'var(--color-secondary)',
    godot: 'var(--color-state-success)'
  }
  return colors[slug] || 'var(--color-primary)'
}

const getCategoryBg = (cat: any) => {
  const slug = getCatSlug(cat)
  const bgs: Record<string, string> = {
    tech: 'var(--color-primary-50)',
    game: 'var(--color-secondary-50)',
    hardware: 'rgba(0,170,255,0.08)',
    anime: 'rgba(255,170,0,0.08)',
    ai_robot: 'var(--color-primary-50)',
    ai_coding: 'var(--color-secondary-50)',
    godot: 'rgba(34,197,94,0.08)'
  }
  return bgs[slug] || 'var(--color-primary-50)'
}

const getCategoryBgLight = (cat: any) => {
  const slug = getCatSlug(cat)
  const bgs: Record<string, string> = {
    tech: 'var(--color-primary-100)',
    game: 'var(--color-secondary-100)',
    hardware: 'rgba(0,170,255,0.15)',
    anime: 'rgba(255,170,0,0.15)',
    ai_robot: 'var(--color-primary-100)',
    ai_coding: 'var(--color-secondary-100)',
    godot: 'rgba(34,197,94,0.15)'
  }
  return bgs[slug] || 'var(--color-primary-100)'
}
</script>