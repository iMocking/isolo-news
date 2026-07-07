<template>
  <article
    class="p-4 transition-all duration-160 cursor-pointer"
    :style="cardStyle"
    @mouseenter="$emit('hover', article)"
    @mouseleave="$emit('leave', article)"
    @click="$emit('click', article.id)"
  >
    <div class="flex items-start gap-3">
      <div class="shrink-0 w-2 h-2 rounded-full mt-1.5" :style="dotStyle"></div>
      <div class="min-w-0 flex-1">
        <div class="flex items-center gap-2 mb-1">
          <span class="px-1.5 py-0.5 rounded text-[9px] font-bold" :style="categoryTagStyle">
            {{ categoryName }}
          </span>
          <!-- 已读角标 -->
          <span v-if="isRead(article.id)" class="shrink-0 px-1 py-0.5 text-[9px] font-bold whitespace-nowrap" style="background: var(--color-primary); color: white; border-radius: var(--radius-sm); font-family: var(--font-mono);">已读</span>
        </div>
        <h4 class="text-sm font-semibold mb-1 truncate" :style="titleStyle">
          {{ article.title }}
        </h4>
        <p class="text-xs line-clamp-2" :style="summaryStyle">
          {{ article.summary }}
        </p>
        <div class="flex items-center justify-between mt-2">
          <span class="text-[10px]" :style="timeStyle">{{ formatTime(article.publishDate) }}</span>
          <span class="text-[10px] font-bold" :style="xpStyle">+{{ article.xpReward }} XP</span>
        </div>
      </div>
    </div>
  </article>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import type { ArticleVO } from '@/api/articles'
import { useCardStyles } from '@/hooks/useCardStyles'
import { useReadTracker } from '@/composables/useReadTracker'

const props = defineProps<{
  article: ArticleVO
  isHovered: boolean
}>()

const { isRead } = useReadTracker()
const { t } = useI18n()

defineEmits<{
  hover: [article: ArticleVO]
  leave: [article: ArticleVO]
  click: [id: string]
}>()

const { getCardStyle } = useCardStyles()

const getCatSlug = (cat: any): string => typeof cat === 'string' ? cat : cat?.slug || ''

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

const slug = computed(() => getCatSlug(props.article.category))
const categoryName = computed(() => {
  const translated = t(`categories.${slug.value}`)
  const rawName = typeof props.article.category === 'string' ? '' : props.article.category?.name || ''
  return translated !== `categories.${slug.value}` ? translated : (rawName || slug.value)
})

const cardStyle = computed(() => {
  return getCardStyle('default', props.isHovered)
})

const dotStyle = computed(() => {
  const colors: Record<string, string> = {
    tech: 'var(--color-primary)', game: 'var(--color-secondary)',
    hardware: 'var(--state-info)', anime: 'var(--state-warning)',
    ai_robot: 'var(--color-primary)', ai_coding: 'var(--color-secondary)',
    godot: 'var(--state-success)'
  }
  return { background: colors[slug.value] || 'var(--color-primary)' }
})

const categoryStyles: Record<string, { bg: string; color: string }> = {
  tech: { bg: 'var(--color-primary50)', color: 'var(--color-primary)' },
  game: { bg: 'var(--color-secondary50)', color: 'var(--color-secondary)' },
  hardware: { bg: 'rgba(0, 170, 255, 0.1)', color: 'var(--state-info)' },
  anime: { bg: 'rgba(255, 170, 0, 0.1)', color: 'var(--state-warning)' },
  ai_robot: { bg: 'var(--color-primary50)', color: 'var(--color-primary)' },
  ai_coding: { bg: 'var(--color-secondary50)', color: 'var(--color-secondary)' },
  godot: { bg: 'rgba(34,197,94,0.1)', color: 'var(--state-success)' }
}

const categoryTagStyle = computed(() => {
  const style = categoryStyles[slug.value]
  return { fontFamily: 'var(--font-mono)', background: style?.bg || 'var(--color-primary50)', color: style?.color || 'var(--color-primary)' }
})

const titleStyle = computed(() => ({ color: 'var(--color-textPrimary)' }))
const summaryStyle = computed(() => ({ color: 'var(--color-textSecondary)' }))
const timeStyle = computed(() => ({ color: 'var(--color-textTertiary)' }))
const xpStyle = computed(() => ({ fontFamily: 'var(--font-mono)', color: 'var(--state-success)' }))
</script>
