<template>
  <aside class="lg:col-span-1 space-y-8">
    <!-- Related Missions -->
    <div class="p-5" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);">
      <h3 class="text-sm font-semibold mb-4 pb-3 whitespace-nowrap" style="color: var(--color-text-primary); font-family: var(--font-display); letter-spacing: 0.02em; border-bottom: 1px solid var(--color-border-subtle);">相关任务</h3>
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

    <!-- Interaction Area -->
    <div class="p-5" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);">
      <h3 class="text-sm font-semibold mb-4 pb-3 whitespace-nowrap" style="color: var(--color-text-primary); font-family: var(--font-display); letter-spacing: 0.02em; border-bottom: 1px solid var(--color-border-subtle);">互动区</h3>
      <div class="space-y-3">
        <!-- Comment count -->
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-2">
            <MessageCircle class="w-4 h-4" style="color: var(--color-text-tertiary);" />
            <span class="text-sm whitespace-nowrap" style="color: var(--color-text-secondary);">{{ commentCount }} 条评论</span>
          </div>
        </div>
        <!-- Like button -->
        <button 
          class="w-full flex items-center justify-center gap-2 px-4 py-2.5 text-sm font-medium whitespace-nowrap transition-all duration-150"
          :style="{
            background: isLiked ? 'var(--color-primary)' : 'var(--color-primary-50)',
            color: isLiked ? 'white' : 'var(--color-primary)',
            border: '1px solid var(--color-border)'
          }"
          @mouseenter="($event.target as HTMLElement).style.background = isLiked ? 'var(--color-primary-dark)' : 'var(--color-primary-100)'"
          @mouseleave="($event.target as HTMLElement).style.background = isLiked ? 'var(--color-primary)' : 'var(--color-primary-50)'"
          @click="handleLike"
        >
          <Heart class="w-4 h-4" :fill="isLiked ? 'currentColor' : 'none'" />
          <span>{{ isLiked ? '已点赞' : `+${likeCount} 点赞` }}</span>
        </button>
        <!-- Favorite button -->
        <button 
          class="w-full flex items-center justify-center gap-2 px-4 py-2.5 text-sm font-medium whitespace-nowrap transition-all duration-150"
          :style="{
            background: isFavorited ? 'var(--state-warning)' : 'var(--color-bg-elevated)',
            color: isFavorited ? 'white' : 'var(--color-text-secondary)',
            border: '1px solid var(--color-border-subtle)'
          }"
          @mouseenter="($event.target as HTMLElement).style.background = isFavorited ? '#e6a700' : 'var(--color-bg-tertiary)'"
          @mouseleave="($event.target as HTMLElement).style.background = isFavorited ? 'var(--state-warning)' : 'var(--color-bg-elevated)'"
          @click="handleFavorite"
        >
          <Bookmark class="w-4 h-4" :fill="isFavorited ? 'currentColor' : 'none'" />
          <span>{{ isFavorited ? '已收藏' : '收藏' }}</span>
        </button>
        <!-- Share button -->
        <button class="w-full flex items-center justify-center gap-2 px-4 py-2.5 text-sm font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-bg-elevated); color: var(--color-text-secondary); border: 1px solid var(--color-border-subtle);" @mouseenter="($event.target as HTMLElement).style.background = 'var(--color-bg-tertiary)'" @mouseleave="($event.target as HTMLElement).style.background = 'var(--color-bg-elevated)'">
          <Share2 class="w-4 h-4" />
          <span>分享</span>
        </button>
      </div>
    </div>

    <!-- XP Summary -->
    <div class="p-5" style="background: var(--color-bg-card); border: 1px solid var(--color-border); border-radius: var(--radius-md);">
      <h3 class="text-sm font-semibold mb-4 pb-3 whitespace-nowrap" style="color: var(--color-primary); font-family: var(--font-display); letter-spacing: 0.02em; border-bottom: 1px solid var(--color-border-subtle);">XP 总览</h3>
      <div class="space-y-3">
        <div class="flex items-center justify-between">
          <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">本篇奖励</span>
          <span class="text-sm font-semibold whitespace-nowrap" style="color: var(--color-primary); font-family: var(--font-display);">+500 XP</span>
        </div>
        <div class="flex items-center justify-between">
          <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">今日已获</span>
          <span class="text-sm font-semibold whitespace-nowrap" style="color: var(--color-state-success); font-family: var(--font-display);">+1,200 XP</span>
        </div>
        <div class="pt-2" style="border-top: 1px solid var(--color-border-subtle);">
          <div class="flex items-center justify-between mb-1">
            <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">距离 LV.13</span>
            <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">800 / 2,000 XP</span>
          </div>
          <div class="w-full h-2 rounded-full overflow-hidden" style="background: var(--color-bg-tertiary);">
            <div class="h-full rounded-full" style="width: 40%; background: linear-gradient(90deg, var(--color-primary), var(--color-primaryLight));"></div>
          </div>
        </div>
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useUserStore } from '@/stores/userStore'
import { Heart, Share2, MessageCircle, Bookmark } from 'lucide-vue-next'
import { toggleLike, toggleFavorite as toggleFavoriteApi } from '@/api/articles'

const props = defineProps<{
  articleId: string
  currentArticle: any
  commentCount: number
}>()

const userStore = useUserStore()

const isLiked = ref(false)
const likeCount = ref(0)
const isFavorited = ref(false)

const relatedArticles = ref<Array<{ title: string; time: string }>>([])

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
</script>
