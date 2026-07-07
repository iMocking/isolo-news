<template>
  <aside class="lg:col-span-1 space-y-8 sticky top-24 self-start">
    <!-- Related Missions -->
    <div class="p-5" :style="sidebarPanelStyle">
      <h3 class="text-sm font-semibold mb-4 pb-3 whitespace-nowrap" style="color: var(--color-text-primary); font-family: var(--font-display); letter-spacing: 0.02em; border-bottom: 1px solid var(--color-border-subtle);">{{ $t('articleDetail.sidebar.related') }}</h3>
      <div class="space-y-4">
        <a v-for="(item, index) in relatedArticles" :key="index" href="#" class="block group">
          <div class="flex items-start gap-3">
            <div class="w-16 h-10 shrink-0 flex items-center justify-center overflow-hidden" :style="{ background: `linear-gradient(135deg, ${index % 2 === 0 ? 'var(--color-bg-tertiary)' : 'var(--color-bg-elevated)'}, ${index % 2 === 0 ? 'var(--color-bg-elevated)' : 'var(--color-bg-tertiary)'})` }">
              <span class="text-xs" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">{{ String(index + 1).padStart(2, '0') }}</span>
            </div>
            <div class="min-w-0 flex-1">
              <p class="text-sm font-medium line-clamp-2 group transition-colors duration-150" style="color: var(--color-text-primary);" @mouseenter="($event.target as HTMLElement).style.color = 'var(--color-primary)'" @mouseleave="($event.target as HTMLElement).style.color = 'var(--color-text-primary)'">{{ item.title }}</p>
              <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">{{ item.time }}</span>
            </div>
          </div>
        </a>
      </div>
    </div>

    <!-- Interaction Area — 横排纵向布局 -->
    <div class="p-5" :style="sidebarPanelStyle">
      <h3 class="text-sm font-semibold mb-4 pb-3 whitespace-nowrap" style="color: var(--color-text-primary); font-family: var(--font-display); letter-spacing: 0.02em; border-bottom: 1px solid var(--color-border-subtle);">{{ $t('articleDetail.sidebar.interactions') }}</h3>
      <div class="grid grid-cols-4 gap-2">
        <!-- 评论 -->
        <button class="flex flex-col items-center gap-1.5 py-3 px-1 rounded-lg transition-all duration-150 cursor-default" :style="interactionCardStyle">
          <MessageCircle class="w-[18px] h-[18px]" style="color: var(--color-text-tertiary);" />
          <span class="text-[11px] whitespace-nowrap" style="color: var(--color-text-secondary);">{{ $t('articleDetail.sidebar.comment') }}</span>
          <span class="text-xs font-bold" style="color: var(--color-text-primary); font-family: var(--font-mono);">{{ commentCount }}</span>
        </button>
        <!-- 点赞 -->
        <button class="flex flex-col items-center gap-1.5 py-3 px-1 rounded-lg transition-all duration-150" :style="likeButtonStyle" @click="handleLike">
          <Heart class="w-[18px] h-[18px]" :fill="isLiked ? 'currentColor' : 'none'" :style="{ color: isLiked ? 'white' : 'var(--color-text-tertiary)' }" />
          <span class="text-[11px] whitespace-nowrap" :style="{ color: isLiked ? 'white' : 'var(--color-text-secondary)' }">{{ $t('articleDetail.sidebar.like') }}</span>
          <span class="text-xs font-bold" :style="{ color: isLiked ? 'white' : 'var(--color-text-primary)', fontFamily: 'var(--font-mono)' }">{{ likeCount }}</span>
        </button>
        <!-- 收藏 -->
        <button class="flex flex-col items-center gap-1.5 py-3 px-1 rounded-lg transition-all duration-150" :style="favButtonStyle" @click="handleFavorite">
          <Bookmark class="w-[18px] h-[18px]" :fill="isFavorited ? 'currentColor' : 'none'" :style="{ color: isFavorited ? 'white' : 'var(--color-text-tertiary)' }" />
          <span class="text-[11px] whitespace-nowrap" :style="{ color: isFavorited ? 'white' : 'var(--color-text-secondary)' }">{{ $t('articleDetail.sidebar.favorite') }}</span>
        </button>
        <!-- 分享 -->
        <button class="flex flex-col items-center gap-1.5 py-3 px-1 rounded-lg transition-all duration-150" :style="shareButtonStyle" @click="handleShare">
          <Share2 class="w-[18px] h-[18px]" style="color: var(--color-text-tertiary);" />
          <span class="text-[11px] whitespace-nowrap" style="color: var(--color-text-secondary);">{{ $t('articleDetail.sidebar.share') }}</span>
        </button>
      </div>
    </div>

    <!-- XP Summary — 数据来自 userStore -->
    <div class="p-5" :style="xpPanelStyle">
      <h3 class="text-sm font-semibold mb-4 pb-3 whitespace-nowrap" style="color: var(--color-primary); font-family: var(--font-display); letter-spacing: 0.02em; border-bottom: 1px solid var(--color-border-subtle);">{{ $t('articleDetail.sidebar.xpOverview') }}</h3>
      <div class="space-y-3">
        <div class="flex items-center justify-between">
          <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">{{ $t('articleDetail.sidebar.xpThisArticle') }}</span>
          <span class="text-sm font-semibold whitespace-nowrap" style="color: var(--color-primary); font-family: var(--font-display);">+{{ currentArticle.xpReward || 0 }} XP</span>
        </div>
        <div class="flex items-center justify-between">
          <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">{{ $t('articleDetail.sidebar.currentLevel') }}</span>
          <span class="text-sm font-semibold whitespace-nowrap" style="color: var(--color-state-success); font-family: var(--font-display);">LV.{{ userStore.level }}</span>
        </div>
        <div class="pt-2" style="border-top: 1px solid var(--color-border-subtle);">
          <div class="flex items-center justify-between mb-1">
            <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">{{ $t('articleDetail.sidebar.xpToNextLevel', { level: userStore.level + 1 }) }}</span>
            <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">{{ userStore.xp }} / {{ userStore.maxXp }} XP</span>
          </div>
          <div class="w-full h-2 rounded-full overflow-hidden" style="background: var(--color-bg-tertiary);">
            <div class="h-full rounded-full" :style="{ width: `${xpProgressPercent}%`, background: 'linear-gradient(90deg, var(--color-primary), var(--color-primaryLight))' }"></div>
          </div>
        </div>
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useUserStore } from '@/stores/userStore'
import { Heart, Share2, MessageCircle, Bookmark } from 'lucide-vue-next'
import { toggleLike, toggleFavorite as toggleFavoriteApi } from '@/api/articles'
import { useCardStyles } from '@/hooks/useCardStyles'

const props = defineProps<{
  articleId: string
  currentArticle: any
  commentCount: number
}>()

const userStore = useUserStore()
const { getCardStyle } = useCardStyles()

const isLiked = ref(false)
const likeCount = ref(0)
const isFavorited = ref(false)

const relatedArticles = ref<Array<{ title: string; time: string }>>([])

const sidebarPanelStyle = computed(() => getCardStyle('panel', false))
const xpPanelStyle = computed(() => getCardStyle('elevated', false))

/** 互动卡片基础样式 */
const interactionCardStyle = computed(() => ({
  background: 'var(--color-bg-elevated)',
  border: '1px solid var(--color-border-subtle)',
}))

const likeButtonStyle = computed(() => ({
  background: isLiked.value ? 'var(--color-primary)' : 'var(--color-bg-elevated)',
  border: `1px solid ${isLiked.value ? 'var(--color-primary)' : 'var(--color-border-subtle)'}`,
  color: isLiked.value ? 'white' : 'var(--color-text-secondary)',
}))

const favButtonStyle = computed(() => ({
  background: isFavorited.value ? 'var(--state-warning)' : 'var(--color-bg-elevated)',
  border: `1px solid ${isFavorited.value ? 'var(--state-warning)' : 'var(--color-border-subtle)'}`,
  color: isFavorited.value ? 'white' : 'var(--color-text-secondary)',
}))

const shareButtonStyle = computed(() => ({
  background: 'var(--color-bg-elevated)',
  border: '1px solid var(--color-border-subtle)',
}))

/** XP 进度百分比（防止除以 0） */
const xpProgressPercent = computed(() =>
  userStore.maxXp > 0 ? Math.min(100, Math.round((userStore.xp / userStore.maxXp) * 100)) : 0
)

const handleLike = async () => {
  if (!userStore.isLoggedIn) return
  try {
    const res = await toggleLike(props.articleId)
    if (res.code === 0) {
      isLiked.value = res.data.isLiked
      likeCount.value += res.data.isLiked ? 1 : -1
    }
  } catch {
    console.warn('点赞操作失败')
  }
}

const handleFavorite = async () => {
  if (!userStore.isLoggedIn) return
  try {
    const res = await toggleFavoriteApi(props.articleId)
    if (res.code === 0) {
      isFavorited.value = res.data.isFavorited
      userStore.toggleFavorite(props.articleId)
    }
  } catch {
    console.warn('收藏操作失败')
  }
}

/** 分享功能（优先使用 Web Share API） */
const handleShare = async () => {
  const url = window.location.href
  const title = props.currentArticle?.title || ''
  if (navigator.share) {
    try {
      await navigator.share({ title, url })
    } catch {
      // 用户取消分享不视为错误
    }
  } else {
    try {
      await navigator.clipboard.writeText(url)
      // 可考虑使用 toast 提示
    } catch {
      console.warn('复制链接失败')
    }
  }
}
</script>
