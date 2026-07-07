<template>
  <div>
    <!-- Article Hero -->
    <section class="w-full relative overflow-hidden" style="background: linear-gradient(180deg, var(--color-bg-secondary) 0%, var(--color-bg-card) 100%);">
      <div v-if="themeStore.currentTheme === 'nexus'" class="absolute inset-0 pointer-events-none opacity-5" style="background: repeating-linear-gradient(0deg, transparent, transparent 2px, var(--color-text-primary) 2px, var(--color-text-primary) 3px);"></div>

      <!-- 加载中状态 -->
      <div v-if="articleLoading && !currentArticle" class="max-w-7xl mx-auto px-6 pt-24 pb-10 relative z-10">
        <div class="flex flex-col items-center justify-center py-20">
          <div class="w-8 h-8 border-2 rounded-full animate-spin" style="border-color: var(--color-text-tertiary); border-top-color: var(--color-primary);"></div>
          <p class="mt-4 text-sm" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">{{ $t('common.loading') }}</p>
        </div>
      </div>

      <!-- 错误状态 -->
      <div v-else-if="loadError && !currentArticle" class="max-w-7xl mx-auto px-6 pt-24 pb-10 relative z-10">
        <div class="flex flex-col items-center justify-center py-20">
          <p class="text-sm mb-4" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">⚠ {{ $t('common.errors.networkError') }}</p>
          <router-link to="/article/list" class="inline-flex items-center gap-2 px-4 py-2 text-xs font-medium transition-all duration-150"
            :style="{ background: 'var(--color-primary50)', color: 'var(--color-primary)', border: '1px solid var(--color-primary)', borderRadius: 'var(--radius-sm)' }">
            <ArrowLeft class="w-3 h-3" />
            <span>{{ $t('articleDetail.backToList') }}</span>
          </router-link>
        </div>
      </div>

      <!-- 文章已加载 -->
      <div v-if="currentArticle" class="max-w-7xl mx-auto px-6 pt-24 pb-10 relative z-10">
        <!-- Breadcrumb -->
        <div class="flex items-center gap-2 mb-8">
          <router-link to="/" class="text-xs whitespace-nowrap transition-colors duration-150" style="color: var(--color-text-tertiary);" @mouseenter="($event.target as HTMLElement).style.color = 'var(--color-primary)'" @mouseleave="($event.target as HTMLElement).style.color = 'var(--color-text-tertiary)'">{{ $t('articleDetail.breadcrumb.home') }}</router-link>
          <span class="text-xs" style="color: var(--color-text-tertiary);">/</span>
          <router-link to="/article/list" class="text-xs whitespace-nowrap transition-colors duration-150" style="color: var(--color-text-tertiary);" @mouseenter="($event.target as HTMLElement).style.color = 'var(--color-primary)'" @mouseleave="($event.target as HTMLElement).style.color = 'var(--color-text-tertiary)'">{{ $t('articleDetail.breadcrumb.list') }}</router-link>
          <span class="text-xs" style="color: var(--color-text-tertiary);">/</span>
          <span class="text-xs whitespace-nowrap" style="color: var(--color-text-secondary);">{{ $t('articleDetail.breadcrumb.current') }}</span>
        </div>

        <!-- Category badge -->
        <div class="mb-6">
          <span class="inline-flex items-center px-3 py-1 text-xs font-semibold whitespace-nowrap" :style="{ background: getCategoryBg(currentArticle.category?.slug), color: getCategoryColor(currentArticle.category?.slug), border: `1px solid ${getCategoryBorder(currentArticle.category?.slug)}`, fontFamily: 'var(--font-display)', letterSpacing: '0.05em' }">{{ $t(`categories.${currentArticle.category?.slug || currentArticle.category}`) }}</span>
        </div>

        <!-- Article title -->
        <h1 class="mb-6" style="font-size: clamp(28px, 3.5vw, 48px); line-height: 1.2; font-weight: 600; color: var(--color-text-primary); word-break: keep-all;">
          {{ currentArticle.title }}
        </h1>

        <!-- Author info row -->
        <div class="flex flex-wrap items-center gap-4 mb-6">
          <div class="flex items-center gap-3">
            <div class="w-9 h-9 rounded-full flex items-center justify-center shrink-0" style="background: linear-gradient(135deg, var(--color-primary-200), var(--color-secondary-100)); border: 1px solid var(--color-border);">
              <span class="text-sm font-bold" style="color: var(--color-text-inverse);">{{ currentArticle.author.charAt(0).toUpperCase() }}</span>
            </div>
            <div class="flex items-center gap-2">
              <span class="text-sm font-medium whitespace-nowrap" style="color: var(--color-text-primary);">{{ currentArticle.author }}</span>
              <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">{{ currentArticle.publishDate }}</span>
            </div>
          </div>
          <div class="flex items-center gap-1">
            <Eye class="w-4 h-4" style="color: var(--color-text-tertiary);" />
            <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">{{ $t('articleDetail.readTime', { count: currentArticle.readTime }) }}</span>
          </div>
        </div>

        <!-- XP Progress Bar -->
        <div class="mb-8 p-4 max-w-2xl" style="background: var(--color-bg-elevated); border: 1px solid var(--color-border); border-radius: var(--radius-md);">
          <div class="flex items-center justify-between mb-2">
            <span class="text-xs font-medium whitespace-nowrap" style="color: var(--color-primary); font-family: var(--font-display);">{{ $t('articleDetail.xpProgress') }}</span>
            <span class="text-xs whitespace-nowrap" style="color: var(--color-text-secondary);">+{{ currentArticle.xpReward }} XP</span>
          </div>
          <div class="w-full h-2 rounded-full overflow-hidden" style="background: var(--color-bg-tertiary);">
            <div class="h-full rounded-full transition-all duration-300" :style="{ width: `${readProgress}%`, background: 'linear-gradient(90deg, var(--color-primary), var(--color-secondary))' }"></div>
          </div>
          <div class="flex items-center justify-between mt-2">
            <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">{{ $t('articleDetail.readProgress', { progress: readProgress }) }}</span>
            <span v-if="!xpClaimed" class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">{{ $t('articleDetail.readReward', { xp: currentArticle.xpReward }) }}</span>
            <span v-else class="text-xs font-semibold whitespace-nowrap" style="color: var(--color-primary);">+{{ xpGained || currentArticle.xpReward }} XP ✓</span>
          </div>
        </div>

        <!-- Tag pills -->
        <div class="flex flex-wrap gap-2">
          <span v-for="tag in articleTags" :key="tag" class="inline-flex items-center px-3 py-1 text-xs whitespace-nowrap" style="background: var(--color-bg-elevated); color: var(--color-text-secondary); border: 1px solid var(--color-border-subtle);">{{ tag }}</span>
        </div>
      </div>
    </section>

    <!-- Article Body + Sidebar -->
    <section v-if="currentArticle" class="max-w-7xl mx-auto px-6 py-12">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
        <!-- Main article content -->
        <article class="lg:col-span-2">
          <!-- Terminal frame container -->
          <div class="relative" style="background: var(--color-bg-card); border: 1px solid var(--color-border); border-radius: var(--radius-lg);">
            <!-- Terminal header bar -->
            <div class="flex items-center gap-2 px-4 py-3" style="background: var(--color-bg-elevated); border-bottom: 1px solid var(--color-border); border-radius: var(--radius-lg) var(--radius-lg) 0 0;">
              <div class="w-3 h-3 rounded-full shrink-0" style="background: var(--state-error);"></div>
              <div class="w-3 h-3 rounded-full shrink-0" style="background: var(--state-warning);"></div>
              <div class="w-3 h-3 rounded-full shrink-0" style="background: var(--state-success);"></div>
              <span class="text-xs ml-2 whitespace-nowrap" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">nexus://article/gpt5-game-dev.log</span>
            </div>

            <!-- Scanline overlay on terminal body -->
            <div v-if="themeStore.currentTheme === 'nexus'" class="absolute inset-0 pointer-events-none opacity-[0.03] z-10" style="background: repeating-linear-gradient(0deg, transparent, transparent 1px, var(--color-text-primary) 1px, var(--color-text-primary) 2px); border-radius: var(--radius-lg);"></div>

            <!-- Article text body - 动态渲染 API 返回的 content -->
            <div class="article-content relative z-10 p-6 md:p-10" style="font-size: 15px; color: var(--color-text-primary);">
              <div v-if="currentArticle.content" style="display:contents" v-html="currentArticle.content" />
              <div v-else class="py-8 text-center">
                <div class="inline-block w-5 h-5 border-2 rounded-full animate-spin" style="border-color: var(--color-text-tertiary); border-top-color: var(--color-primary);"></div>
                <p class="mt-3 text-xs" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">LOADING_CONTENT...</p>
              </div>
              <!-- EOF 仅在内容加载完成后显示 -->
              <div v-if="currentArticle.content" class="flex items-center gap-1 mt-6 pt-4" style="border-top: 1px solid var(--color-border-subtle);">
                <span class="text-xs" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">EOF</span>
              </div>
            </div>
          </div>

          <!-- Article Footer -->
          <div class="mt-8 space-y-6">
            <!-- Tags row -->
            <div class="flex flex-wrap gap-2">
              <span v-for="tag in currentArticle.tags" :key="tag" class="inline-flex items-center px-3 py-1 text-xs whitespace-nowrap" style="background: var(--color-primary-50); color: var(--color-primary); border: 1px solid var(--color-border);">#{{ tag }}</span>
            </div>

            <!-- Author bio card -->
            <div class="flex items-center gap-4 p-5" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);">
              <div class="w-12 h-12 rounded-full flex items-center justify-center shrink-0" style="background: linear-gradient(135deg, var(--color-primary-200), var(--color-secondary-100)); border: 1px solid var(--color-border);">
                <span class="text-lg font-bold" style="color: var(--color-text-inverse);">A</span>
              </div>
              <div class="min-w-0 flex-1">
                <p class="text-sm font-semibold truncate" style="color: var(--color-text-primary);">{{ currentArticle.author || $t('articleDetail.authorBio') }}</p>
                <p class="text-xs mt-1 line-clamp-2" style="color: var(--color-text-tertiary);">{{ $t('articleDetail.authorDesc') }}</p>
              </div>
            </div>

            <!-- Back button -->
            <router-link to="/article/list" class="inline-flex items-center justify-center gap-2 px-5 py-2.5 text-sm font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-primary-50); color: var(--color-primary); border: 1px solid var(--color-border);" @mouseenter="($event.target as HTMLElement).style.background = 'var(--color-primary-100)'; ($event.target as HTMLElement).style.transform = 'translateY(-1px)'" @mouseleave="($event.target as HTMLElement).style.background = 'var(--color-primary-50)'; ($event.target as HTMLElement).style.transform = 'translateY(0)'">
              <ArrowLeft class="w-4 h-4" />
              <span>{{ $t('articleDetail.backToList') }}</span>
            </router-link>
          </div>

          <ArticleComments :articleId="articleId" @update:comment-count="commentCount = $event" />
        </article>

        <ArticleSidebar :articleId="articleId" :currentArticle="currentArticle" :commentCount="commentCount" />
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import { useThemeStore } from '@/stores/themeStore'
import { useArticleStore } from '@/stores/articleStore'
import { useUserStore } from '@/stores/userStore'
import { fetchArticleById, recordRead } from '@/api/articles'
import { Eye, ArrowLeft } from 'lucide-vue-next'
import { useReadTracker } from '@/composables/useReadTracker'
import ArticleComments from './ArticleComments.vue'
import ArticleSidebar from './ArticleSidebar.vue'

const route = useRoute()
const themeStore = useThemeStore()
const articleStore = useArticleStore()
const userStore = useUserStore()

const readProgress = ref(0)
const xpClaimed = ref(false)    // 是否已领取阅读 XP
const xpGained = ref(0)         // 本次阅读获得的 XP
const commentCount = ref(0)
const articleId = computed(() => route.params.id as string)
const articleLoading = ref(true) // 是否正在加载
const loadError = ref(false)     // 加载是否失败
const { markAsRead } = useReadTracker()

/** 当前文章：仅在 store 中找到对应 ID 时才返回，避免刷新生效时错误展示 loading */
const currentArticle = computed(() => {
  return articleStore.articles.find(a => a.id === articleId.value) || null
})

const articleTags = computed(() => currentArticle.value.tags || [])

const getCategoryColor = (category: string) => {
  const colors: Record<string, string> = {
    tech: 'var(--color-primary)',
    game: 'var(--color-secondary)',
    hardware: 'var(--state-info)',
    anime: 'var(--state-warning)',
    ai_robot: 'var(--color-primary)',
    ai_coding: 'var(--color-secondary)',
    godot: 'var(--state-success)'
  }
  return colors[category] || 'var(--color-primary)'
}

const getCategoryBg = (category: string) => {
  const colors: Record<string, string> = {
    tech: 'var(--color-primary-50)',
    game: 'var(--color-secondary-50)',
    hardware: 'rgba(0, 170, 255, 0.1)',
    anime: 'rgba(255, 170, 0, 0.1)',
    ai_robot: 'var(--color-primary-50)',
    ai_coding: 'var(--color-secondary-50)',
    godot: 'rgba(34, 197, 94, 0.1)'
  }
  return colors[category] || 'var(--color-primary-50)'
}

const getCategoryBorder = (category: string) => {
  const colors: Record<string, string> = {
    tech: 'var(--color-primary-200)',
    game: 'var(--color-secondary-200)',
    hardware: 'rgba(0, 170, 255, 0.3)',
    anime: 'rgba(255, 170, 0, 0.3)',
    ai_robot: 'var(--color-primary-200)',
    ai_coding: 'var(--color-secondary-200)',
    godot: 'rgba(34, 197, 94, 0.3)'
  }
  return colors[category] || 'var(--color-primary-200)'
}

/** 阅读进度达到 90% 时上报阅读+领取 XP */
async function claimReadXp() {
  if (xpClaimed.value || !currentArticle.value) return
  xpClaimed.value = true
  try {
    const readRes = await recordRead(articleId.value)
    if (readRes.code === 0 && readRes.data) {
      xpGained.value = readRes.data.xpGained
      markAsRead(articleId.value)
      await userStore.refreshUserData()
    }
  } catch {
    console.warn('阅读记录上报失败')
  }
}

onMounted(async () => {
  // 从 API 获取完整文章内容（含 content 字段）
  try {
    const res = await fetchArticleById(articleId.value)
    if (res.code === 0 && res.data) {
      const idx = articleStore.articles.findIndex(a => a.id === articleId.value)
      if (idx >= 0) {
        articleStore.articles[idx] = res.data
      } else {
        articleStore.articles.push(res.data)
      }
      loadError.value = false
    } else {
      loadError.value = true
    }
  } catch {
    loadError.value = true
    console.warn('文章详情加载失败')
  } finally {
    articleLoading.value = false
  }

  // 监听滚动位置 → 计算真实阅读进度
  const handleScroll = () => {
    if (xpClaimed.value) return
    const scrollTop = window.scrollY
    const docHeight = document.documentElement.scrollHeight - window.innerHeight
    if (docHeight <= 0) return
    const progress = Math.min(100, Math.round((scrollTop / docHeight) * 100))
    readProgress.value = progress
    if (progress >= 90) claimReadXp()
  }

  window.addEventListener('scroll', handleScroll, { passive: true })
  handleScroll()
})

onUnmounted(() => {
  // scroll listener 会在组件销毁时一并被移除
})
</script>


