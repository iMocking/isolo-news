<template>
  <div class="min-h-screen">
    <!-- Scanline overlay -->
    <div v-if="themeStore.currentTheme === 'nexus'" class="fixed inset-0 pointer-events-none z-40 opacity-[0.03]" style="background: repeating-linear-gradient(0deg, transparent, transparent 2px, rgba(255, 255, 255, 0.08) 2px, rgba(255, 255, 255, 0.08) 4px);"></div>

    <AppNavigation />

    <!-- Search Header -->
    <section class="relative pt-24 pb-8" style="background: var(--color-bg-primary);">
      <div class="max-w-7xl mx-auto px-6">
        <!-- HUD header decoration -->
        <div class="flex items-center gap-3 mb-4">
          <span class="w-3 h-3 rounded-sm" style="background: var(--color-primary); box-shadow: 0 0 8px rgba(0, 240, 255, 0.5);"></span>
          <span class="text-xs tracking-widest uppercase" style="font-family: var(--font-mono); color: var(--color-primary);">// SEARCH MODULE ACTIVE</span>
          <span class="flex-1 h-px" style="background: linear-gradient(90deg, var(--color-border), transparent);"></span>
        </div>
        <h1 class="text-2xl md:text-3xl font-bold mb-6" style="font-family: var(--font-display); color: var(--color-text-primary); word-break: keep-all;">搜索终端 // SEARCH TERMINAL</h1>

        <!-- Search Input -->
        <div class="relative max-w-2xl">
          <div class="flex items-center rounded-md overflow-hidden transition-all duration-150" style="background: var(--color-bg-card); border: 1px solid var(--color-border);">
            <Search class="w-5 h-5 ml-4 shrink-0" style="color: var(--color-primary);" />
            <input type="text" v-model="searchQuery" class="flex-1 px-4 py-3 text-base font-medium outline-none" style="background: transparent; color: var(--color-text-primary); font-family: var(--font-body); caret-color: var(--color-primary);" placeholder="输入搜索关键词..." />
            <span class="shrink-0 px-4 py-3 text-sm" style="font-family: var(--font-mono); color: var(--color-primary);">{{ searchResults.length }} 结果</span>
          </div>
          <!-- Blinking cursor decoration -->
          <div v-if="searchQuery" class="absolute right-[120px] top-3 w-0.5 h-5 animate-pulse" style="background: var(--color-primary);"></div>
        </div>

        <!-- Filter Chips -->
        <div class="flex flex-wrap items-center gap-3 mt-6">
          <div class="flex items-center gap-2 px-3 py-2 rounded-md text-xs font-medium cursor-pointer transition-all duration-150" style="background: var(--color-primary-50); border: 1px solid var(--color-primary); color: var(--color-primary); font-family: var(--font-mono);">
            <Clock class="w-3 h-3" />
            时间 (全部)
            <ChevronDown class="w-3 h-3" />
          </div>
          <div class="flex items-center gap-2 px-3 py-2 rounded-md text-xs font-medium cursor-pointer transition-all duration-150" style="background: var(--color-primary-50); border: 1px solid var(--color-primary); color: var(--color-primary); font-family: var(--font-mono);">
            <Layers class="w-3 h-3" />
            分类 (全部)
            <ChevronDown class="w-3 h-3" />
          </div>
          <div class="flex items-center gap-2 px-3 py-2 rounded-md text-xs font-medium cursor-pointer transition-all duration-150" style="background: var(--color-primary-50); border: 1px solid var(--color-primary); color: var(--color-primary); font-family: var(--font-mono);">
            <ArrowUpDown class="w-3 h-3" />
            排序 (相关度)
            <ChevronDown class="w-3 h-3" />
          </div>
        </div>
      </div>
    </section>

    <!-- Search Results + Sidebar -->
    <section class="max-w-7xl mx-auto px-6 py-8">
      <div class="grid gap-8" :style="{ gridTemplateColumns: '1fr 300px' }">
        <!-- Main Results -->
        <div>
          <div class="flex items-center gap-3 mb-5">
            <span class="w-1 h-5 rounded-sm" style="background: var(--color-primary);"></span>
            <h2 class="text-lg font-semibold truncate" style="font-family: var(--font-display); color: var(--color-text-primary);">扫描结果</h2>
            <span class="text-xs px-2 py-0.5 rounded" style="font-family: var(--font-mono); color: var(--color-primary); background: var(--color-primary-50);">{{ searchResults.length }} HITS</span>
          </div>

          <article 
            v-for="(result, index) in searchResults" 
            :key="result.id"
            class="block mb-4 rounded-lg overflow-hidden transition-all duration-160 cursor-pointer"
            :style="{
              background: 'var(--color-bg-card)',
              border: '1px solid var(--color-border-subtle)'
            }"
            @mouseenter="handleCardHover($event, result.category)"
            @mouseleave="handleCardLeave($event)"
            @click="router.push(`/article/${result.id}`)"
          >
            <div class="p-5">
              <div class="flex items-center gap-2 mb-2">
                <span class="px-2 py-0.5 rounded text-[10px] font-bold tracking-wider uppercase" :style="{
                  background: getCategoryColor(result.category),
                  color: 'var(--color-text-inverse)',
                  fontFamily: 'var(--font-mono)'
                }">{{ getCategoryName(result.category) }}</span>
                <span class="px-2 py-0.5 rounded text-[10px] font-bold" style="font-family: var(--font-mono); background: rgba(0, 255, 136, 0.15); color: var(--state-success);">+{{ result.xpReward }} XP</span>
              </div>
              <h3 class="text-base font-semibold mb-2" style="color: var(--color-text-primary); word-break: keep-all;">
                {{ result.title }}
              </h3>
              <p class="text-sm mb-3 line-clamp-2" style="color: var(--color-text-secondary); line-height: 1.6;">{{ result.summary }}</p>
              <div class="flex items-center gap-3 text-xs" style="color: var(--color-text-tertiary);">
                <span>来源: {{ result.source }}</span>
                <span class="shrink-0">{{ result.publishDate }}</span>
              </div>
              <!-- Relevance Bar -->
              <div class="flex items-center gap-3 mt-3">
                <div class="flex-1 h-1.5 rounded-full overflow-hidden" style="background: var(--color-bg-primary);">
                  <div class="h-full rounded-full" :style="{ width: `${result.relevance}%`, background: 'linear-gradient(90deg, var(--color-primary-dark), var(--color-primary))' }"></div>
                </div>
                <span class="text-[10px] font-bold shrink-0" style="font-family: var(--font-mono); color: var(--color-primary);">{{ result.relevance }}%</span>
              </div>
            </div>
          </article>
        </div>

        <!-- Sidebar -->
        <aside class="hidden lg:block">
          <!-- Search Suggestions -->
          <div class="rounded-lg p-5 mb-6" style="background: var(--color-bg-card); border: 1px solid var(--color-border);">
            <div class="flex items-center gap-2 mb-4">
              <Sparkles class="w-4 h-4" style="color: var(--color-primary);" />
              <h3 class="text-sm font-semibold tracking-wide" style="font-family: var(--font-display); color: var(--color-text-primary);">搜索建议</h3>
            </div>
            <div class="flex flex-wrap gap-2">
              <span 
                v-for="suggestion in searchSuggestions" 
                :key="suggestion"
                class="px-3 py-1.5 rounded-md text-xs font-medium cursor-pointer transition-all duration-150"
                :style="{
                  background: searchQuery === suggestion ? 'var(--color-primary-50)' : 'var(--color-bg-primary)',
                  border: searchQuery === suggestion ? '1px solid var(--color-border)' : '1px solid var(--color-border-subtle)',
                  color: searchQuery === suggestion ? 'var(--color-primary)' : 'var(--color-text-secondary)',
                  fontFamily: 'var(--font-mono)'
                }"
                @mouseenter="($event.target as HTMLElement).style.borderColor = 'var(--color-primary)'; ($event.target as HTMLElement).style.color = 'var(--color-primary)'"
                @mouseleave="($event.target as HTMLElement).style.borderColor = searchQuery === suggestion ? 'var(--color-border)' : 'var(--color-border-subtle)'; ($event.target as HTMLElement).style.color = searchQuery === suggestion ? 'var(--color-primary)' : 'var(--color-text-secondary)'"
                @click="searchQuery = suggestion"
              >{{ suggestion }}</span>
            </div>
          </div>

          <!-- Trending Searches -->
          <div class="rounded-lg p-5" style="background: var(--color-bg-card); border: 1px solid var(--color-border);">
            <div class="flex items-center gap-2 mb-4">
              <TrendingUp class="w-4 h-4" style="color: var(--color-secondary);" />
              <h3 class="text-sm font-semibold tracking-wide" style="font-family: var(--font-display); color: var(--color-text-primary);">热门搜索</h3>
            </div>
            <div class="space-y-3">
              <div 
                v-for="(item, index) in trendingSearches" 
                :key="index" 
                class="flex items-center gap-3 cursor-pointer group"
                @click="searchQuery = item"
              >
                <span class="text-xs font-bold shrink-0" :style="{
                  fontFamily: 'var(--font-mono)',
                  color: index < 2 ? 'var(--color-state-error)' : index === 2 ? 'var(--color-state-warning)' : 'var(--color-text-tertiary)'
                }">{{ String(index + 1).padStart(2, '0') }}</span>
                <span class="text-sm truncate group-hover:opacity-70 transition-opacity" style="color: var(--color-text-primary);">{{ item }}</span>
                <ArrowUpRight class="w-3 h-3 shrink-0 ml-auto" style="color: var(--color-text-tertiary);" />
              </div>
            </div>
          </div>
        </aside>
      </div>
    </section>

    <AppFooter />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useThemeStore } from '@/stores/themeStore'
import { useArticleStore } from '@/stores/articleStore'
import { Search, Clock, Layers, ArrowUpDown, ChevronDown, Sparkles, TrendingUp, ArrowUpRight } from 'lucide-vue-next'
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
import type { ArticleCategory } from '@/types'

const route = useRoute()
const router = useRouter()
const themeStore = useThemeStore()
const articleStore = useArticleStore()

const searchQuery = ref((route.query.q as string) || '')

watch(() => route.query.q, (newQ) => {
  searchQuery.value = (newQ as string) || ''
})

watch(searchQuery, (newQuery) => {
  if (newQuery) {
    router.replace({ path: '/search', query: { q: newQuery } })
  }
})

const searchSuggestions = ['AI游戏', 'AI绘画', 'VR游戏', 'NPC智能', '游戏引擎', 'ChatGPT']

const trendingSearches = ['PS6首发阵容', 'RTX5090评测', '原神5.0', '量子芯片', '苹果Vision Pro 2']

const searchResults = computed(() => {
  if (!searchQuery.value) return articleStore.articles.map((article, index) => ({
    ...article,
    source: article.author,
    relevance: Math.max(50, 100 - index * 10)
  }))
  const query = searchQuery.value.toLowerCase()
  return articleStore.articles
    .filter(item => 
      item.title.toLowerCase().includes(query) ||
      item.summary.toLowerCase().includes(query) ||
      item.tags.some(tag => tag.toLowerCase().includes(query))
    )
    .map(article => ({
      ...article,
      source: article.author,
      relevance: Math.floor(Math.random() * 30) + 70
    }))
})

const getCategoryColor = (category: ArticleCategory) => {
  const colors: Record<ArticleCategory, string> = {
    tech: 'var(--color-primary)',
    game: 'var(--color-secondary)',
    hardware: 'var(--color-state-info)',
    anime: 'var(--color-state-warning)'
  }
  return colors[category] || 'var(--color-primary)'
}

const getCategoryName = (category: ArticleCategory) => {
  const names: Record<ArticleCategory, string> = {
    tech: '科技',
    game: '游戏',
    hardware: '硬件',
    anime: '二次元'
  }
  return names[category] || '科技'
}

const handleCardHover = (event: MouseEvent, category: ArticleCategory) => {
  const target = event.currentTarget as HTMLElement
  target.style.borderColor = getCategoryColor(category)
  target.style.boxShadow = category === 'anime' ? '0 0 20px rgba(255, 45, 149, 0.08)' : '0 0 20px rgba(0, 240, 255, 0.08)'
}

const handleCardLeave = (event: MouseEvent) => {
  const target = event.currentTarget as HTMLElement
  target.style.borderColor = 'var(--color-border-subtle)'
  target.style.boxShadow = 'none'
}
</script>