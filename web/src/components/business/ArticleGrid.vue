<template>
  <section class="max-w-7xl mx-auto px-6 py-12">
    <!--
      Flexbox 双列布局：
      - 主内容区 flex-1 填满剩余空间，但最大 720px 防止卡片过宽
      - 侧边栏固定 280px，不收缩
      - justify-center 让整体居中
      - min-w-0 防止文本/图片溢出撑宽 flex 子项
    -->
    <div class="flex flex-col lg:flex-row gap-6 justify-center">
      <!-- Main Bento Grid -->
      <div class="min-w-0 flex-1 max-w-full lg:max-w-[800px]">
        <!-- Section Header -->
        <div class="flex items-center gap-3 mb-6">
          <span class="w-1 h-5 rounded-sm" :style="headerAccentStyle"></span>
          <h2 class="text-xl font-semibold truncate" :style="headerStyle">
            {{ sectionTitle }}
          </h2>
          <span class="text-xs px-2 py-0.5 rounded" :style="headerTagStyle">
            {{ sectionTag }}
          </span>
        </div>

        <!-- Bento Grid -->
        <div class="grid gap-4" :style="bentoGridStyle">
          <!-- 大卡片（精选文章） -->
          <article
            v-for="article in featuredArticles"
            :key="article.id"
            class="rounded-lg overflow-hidden transition-all duration-160 cursor-pointer"
            :style="largeCardStyle(article)"
            @mouseenter="handleCardHover(article)"
            @mouseleave="handleCardLeave(article)"
            @click="handleArticleClick(article.id)"
          >
            <div class="relative h-48 overflow-hidden" :style="imageContainerStyle">
              <img :src="article.imageUrl" :alt="article.title"
                class="w-full h-full object-cover opacity-70" />
              <div class="absolute top-3 left-3 flex items-center gap-2">
                <span class="px-2 py-0.5 rounded text-[10px] font-bold tracking-wider uppercase"
                  :style="categoryTagStyle(article)">
                  {{ getCategoryName(article.category) }}
                </span>
              </div>
              <div class="absolute top-3 right-3">
                <span class="px-2 py-0.5 rounded text-[10px] font-bold" :style="xpTagStyle">
                  +{{ article.xpReward }} XP
                </span>
              </div>
            </div>
            <div class="p-5">
              <div class="flex items-center gap-2 mb-2">
                <!-- 已读角标 -->
                <span v-if="isRead(article.id)" class="shrink-0 px-1.5 py-0.5 text-[10px] font-bold whitespace-nowrap" style="background: var(--color-primary); color: white; border-radius: var(--radius-sm); font-family: var(--font-mono);">已读</span>
                <h3 class="text-base font-semibold truncate" :style="cardTitleStyle">
                  {{ article.title }}
                </h3>
              </div>
              <p class="text-sm mb-3 line-clamp-2" :style="cardSummaryStyle">
                {{ article.summary }}
              </p>
              <div class="flex items-center gap-3 text-xs" :style="cardMetaStyle">
                <span class="truncate">{{ article.author }}</span>
                <span class="shrink-0">{{ formatTime(article.publishDate) }}</span>
              </div>
            </div>
          </article>

          <!-- 小卡片 -->
          <ArticleSmallCard
            v-for="article in smallArticles"
            :key="article.id"
            :article="article"
            :isHovered="hoveredCards.has(article.id)"
            @hover="handleCardHover"
            @leave="handleCardLeave"
            @click="handleArticleClick"
          />
        </div>

        <!-- 空数据提示 -->
        <EmptyState
          v-if="articles.length === 0"
          :type="emptyType || 'empty'"
          class="col-span-2"
        />
      </div>

      <!-- Sidebar (固定 280px，不收缩) -->
      <aside class="hidden lg:block w-[280px] shrink-0 pt-[52px]">
        <slot name="sidebar" />
      </aside>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useThemeStore } from '@/stores/themeStore'
import type { ArticleVO } from '@/api/articles'
import ArticleSmallCard from './ArticleSmallCard.vue'
import { useCardStyles } from '@/hooks/useCardStyles'
import EmptyState from '@/components/base/EmptyState.vue'
import { useReadTracker } from '@/composables/useReadTracker'

interface Props {
  articles: ArticleVO[]
  /** 空状态类型，articles 为空时显示 */
  emptyType?: 'empty' | 'offline' | 'search-empty'
}

const getCategoryName = (cat: any): string => {
  const slug = typeof cat === 'string' ? cat : cat?.slug || ''
  const translated = t(`categories.${slug}`)
  return translated !== `categories.${slug}` ? translated : (cat?.name || slug)
}

const getCatSlug = (cat: any): string =>
  typeof cat === 'string' ? cat : cat?.slug || ''

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

const { t } = useI18n()
const props = defineProps<Props>()
const router = useRouter()
const themeStore = useThemeStore()
const { getCardStyle } = useCardStyles()

const { isRead } = useReadTracker()

const handleArticleClick = (articleId: string) => {
  router.push(`/article/detail/${articleId}`)
}

const hoveredCards = ref<Set<string>>(new Set())

const theme = computed(() => themeStore.currentTheme)

const sectionTitle = computed(() => t(`home.grid.title.${theme.value}`))
const sectionTag = computed(() => t(`home.grid.tag.${theme.value}`))

const featuredArticles = computed(() => props.articles.slice(0, 2))
const smallArticles = computed(() => props.articles.slice(2, 6))

const bentoGridStyle = computed(() => ({
  gridTemplateColumns: 'minmax(0, 1fr) minmax(0, 1fr)'
}))

const headerAccentStyle = computed(() => ({ background: 'var(--color-primary)' }))
const headerStyle = computed(() => ({ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }))
const headerTagStyle = computed(() => ({ fontFamily: 'var(--font-mono)', color: 'var(--color-primary)', background: 'var(--color-primary50)' }))

const largeCardStyle = computed(() => (article: ArticleVO) => {
  const isHovered = hoveredCards.value.has(article.id)
  return {
    ...getCardStyle('article', isHovered),
    gridColumn: 'span 1'
  }
})

const imageContainerStyle = computed(() => ({ background: 'var(--color-bgTertiary)' }))

const categoryTagStyle = computed(() => (article: ArticleVO) => {
  const slug = getCatSlug(article.category)
  const colors: Record<string, { bg: string; color: string }> = {
    tech: { bg: 'var(--color-primary)', color: 'var(--color-textInverse)' },
    game: { bg: 'var(--color-secondary)', color: 'white' },
    hardware: { bg: 'var(--color-primary50)', color: 'var(--color-primary)' },
    anime: { bg: 'var(--color-secondary50)', color: 'var(--color-secondary)' },
    ai_robot: { bg: 'var(--color-primary)', color: 'var(--color-textInverse)' },
    ai_coding: { bg: 'var(--color-secondary)', color: 'white' },
    godot: { bg: 'rgba(34,197,94,0.15)', color: 'var(--state-success)' }
  }
  const c = colors[slug]
  return { fontFamily: 'var(--font-mono)', background: c?.bg || 'var(--color-primary)', color: c?.color || 'var(--color-textInverse)' }
})

const xpTagStyle = computed(() => ({ fontFamily: 'var(--font-mono)', background: 'rgba(0, 255, 136, 0.15)', color: 'var(--state-success)' }))
const cardTitleStyle = computed(() => ({ color: 'var(--color-textPrimary)' }))
const cardSummaryStyle = computed(() => ({ color: 'var(--color-textSecondary)', lineHeight: 1.6 }))
const cardMetaStyle = computed(() => ({ color: 'var(--color-textTertiary)' }))

const handleCardHover = (article: ArticleVO) => { hoveredCards.value.add(article.id) }
const handleCardLeave = (article: ArticleVO) => { hoveredCards.value.delete(article.id) }
</script>
