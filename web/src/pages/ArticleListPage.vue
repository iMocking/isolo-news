<template>
  <div class="min-h-screen">
    <!-- Scanline overlay for NEXUS theme -->
    <div v-if="themeStore.currentTheme === 'nexus'" class="fixed inset-0 pointer-events-none z-50 opacity-[0.03]" style="background: repeating-linear-gradient(0deg, transparent, transparent 2px, rgba(0,240,255,0.08) 2px, rgba(0,240,255,0.08) 4px);"></div>
    
    <AppNavigation />
    
    <!-- Page Header -->
    <section class="pt-24 pb-6 px-6">
      <div class="max-w-7xl mx-auto">
        <h1 class="text-2xl md:text-4xl tracking-wider" style="font-family: var(--font-display); color: var(--color-primary); word-break: keep-all;">任务列表 // MISSION SELECT</h1>
        <p class="mt-3 text-sm" style="font-family: var(--font-mono); color: var(--color-text-secondary);">共 2,847 条情报</p>
        
        <!-- Search bar -->
        <div class="mt-5 max-w-xl relative">
          <div class="flex items-center px-4 h-12" style="background: var(--color-bg-card); border: 1px solid var(--color-border); border-radius: var(--radius-md);">
            <span class="text-sm mr-3" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">&gt;_</span>
            <input 
              type="text" 
              v-model="searchQuery"
              placeholder="搜索情报关键词..." 
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
            {{ cat.name }} <span class="ml-1 text-xs" :style="{
              background: currentCategory === cat.id ? 'var(--color-primary)' : 'var(--color-bg-elevated)',
              color: currentCategory === cat.id ? 'var(--color-text-inverse)' : 'var(--color-text-tertiary)',
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
          <article 
            v-for="(article, index) in filteredArticles" 
            :key="article.id"
            class="flex items-stretch gap-0 p-5 transition-all duration-150 hover:translate-y-[-1px] cursor-pointer"
            :style="{
              background: 'var(--color-bg-card)',
              border: '1px solid var(--color-border-subtle)',
              borderLeft: `3px solid ${getCategoryColor(article.category)}`,
              borderRadius: 'var(--radius-md)'
            }"
            @click="router.push(`/article/${article.id}`)"
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
                <h2 class="text-base font-semibold truncate" style="color: var(--color-text-primary);">{{ article.title }}</h2>
              </div>
              <p class="text-sm line-clamp-2 mb-3" style="color: var(--color-text-secondary); line-height: 1.6;">{{ article.summary }}</p>
              <div class="flex items-center gap-3">
                <div class="w-5 h-5 flex items-center justify-center" style="background: var(--color-bg-elevated); border-radius: var(--radius-full);">
                  <span class="text-[10px] font-bold" :style="{ color: getCategoryColor(article.category) }">{{ article.author.charAt(0).toUpperCase() }}</span>
                </div>
                <span class="text-xs" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">{{ article.author }} · {{ formatTime(article.publishDate) }} · {{ article.commentCount || 0 }} 评论</span>
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
              }">开始任务</button>
            </div>
          </article>

          <!-- Pagination -->
          <div class="flex items-center justify-center gap-2 pt-6">
            <button class="px-3 py-2 text-xs font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); color: var(--color-text-tertiary); border-radius: var(--radius-sm); font-family: var(--font-mono);">PREV</button>
            <button class="w-9 h-9 flex items-center justify-center text-sm font-medium whitespace-nowrap" style="background: var(--color-primary); color: var(--color-text-inverse); border-radius: var(--radius-sm); font-family: var(--font-display);">1</button>
            <button class="w-9 h-9 flex items-center justify-center text-sm font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); color: var(--color-text-secondary); border-radius: var(--radius-sm); font-family: var(--font-mono);">2</button>
            <button class="w-9 h-9 flex items-center justify-center text-sm font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); color: var(--color-text-secondary); border-radius: var(--radius-sm); font-family: var(--font-mono);">3</button>
            <span class="text-sm px-1" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">...</span>
            <button class="w-9 h-9 flex items-center justify-center text-sm font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); color: var(--color-text-secondary); border-radius: var(--radius-sm); font-family: var(--font-mono);">42</button>
            <button class="px-4 py-2 text-xs font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-primary-50); border: 1px solid var(--color-primary); color: var(--color-primary); border-radius: var(--radius-sm); font-family: var(--font-display);">NEXT STAGE →</button>
          </div>
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

    <AppFooter />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useThemeStore } from '@/stores/themeStore'
import { useArticleStore } from '@/stores/articleStore'
import { fetchHotTags, fetchTrendingTopics } from '@/api/articles'
import { fetchLeaderboard } from '@/api/user'
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
import ArticleFilterSidebar from '@/components/business/ArticleFilterSidebar.vue'

const router = useRouter()
const themeStore = useThemeStore()
const articleStore = useArticleStore()

const searchQuery = ref('')
const currentCategory = ref<string>('all')

/** 格式化时间戳 */
const formatTime = (ts: number | string): string => {
  if (!ts) return ''
  if (typeof ts === 'string') return ts
  const now = Date.now() / 1000
  const diff = now - ts
  if (diff < 3600) return `${Math.floor(diff / 60)}分钟前`
  if (diff < 86400) return `${Math.floor(diff / 3600)}小时前`
  if (diff < 2592000) return `${Math.floor(diff / 86400)}天前`
  return new Date(ts * 1000).toLocaleDateString()
}

/** 从 store 获取分类列表 */
const categories = computed(() => {
  const cats = articleStore.categories.map(c => ({ id: c.id, name: c.name, count: c.count }))
  return [{ id: 'all', name: '全部', count: articleStore.total || 0 }, ...cats]
})

const hotTags = ref<Array<{ name: string; bg: string; border: string; color: string }>>([])
const leaderboard = ref<Array<{ title: string; views: string }>>([])

const defaultTagStyles = {
  bg: 'var(--color-primary-50)',
  border: '1px solid var(--color-primary-100)',
  color: 'var(--color-primary)'
}

onMounted(async () => {
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

/** 从 store 获取文章列表（自动分页过滤） */
const articles = computed(() => articleStore.articles)

const filteredArticles = computed(() => {
  return articles.value.filter(article => {
    const catSlug = typeof article.category === 'string' ? article.category : (article.category?.slug || '')
    const matchesCategory = currentCategory.value === 'all' || catSlug === currentCategory.value
    const title = article.title || ''
    const summary = article.summary || ''
    const matchesSearch = !searchQuery.value ||
      title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      summary.toLowerCase().includes(searchQuery.value.toLowerCase())
    return matchesCategory && matchesSearch
  })
})

/** 获取分类 slug */
const getCatSlug = (cat: any): string =>
  typeof cat === 'string' ? cat : cat?.slug || ''

/** 获取分类名称 */
const getCategoryName = (cat: any): string =>
  typeof cat === 'string'
    ? ({ tech: '科技', game: '游戏', hardware: '硬件', anime: '二次元', ai_robot: 'AI机器人', ai_coding: 'AI编程', godot: 'Godot' }[cat] || cat)
    : cat?.name || ''

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