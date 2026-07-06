<template>
  <section class="max-w-7xl mx-auto px-6 py-12">
    <div class="grid gap-6" :style="gridStyle">
      <!-- Main Bento Grid -->
      <div>
        <!-- Section Header -->
        <div class="flex items-center gap-3 mb-6">
          <span
            class="w-1 h-5 rounded-sm"
            :style="headerAccentStyle"
          ></span>
          <h2
            class="text-xl font-semibold truncate"
            :style="headerStyle"
          >
            {{ sectionTitle }}
          </h2>
          <span
            class="text-xs px-2 py-0.5 rounded"
            :style="headerTagStyle"
          >
            {{ sectionTag }}
          </span>
        </div>

        <!-- Bento Grid -->
        <div class="grid gap-4" :style="bentoGridStyle">
          <!-- Large Cards -->
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
              <img
                :src="article.imageUrl"
                :alt="article.title"
                class="w-full h-full object-cover opacity-70"
              />
              <div class="absolute top-3 left-3 flex items-center gap-2">
                <span
                  class="px-2 py-0.5 rounded text-[10px] font-bold tracking-wider uppercase"
                  :style="categoryTagStyle(article)"
                >
                  {{ article.category }}
                </span>
              </div>
              <div class="absolute top-3 right-3">
                <span
                  class="px-2 py-0.5 rounded text-[10px] font-bold"
                  :style="xpTagStyle"
                >
                  +{{ article.xpReward }} XP
                </span>
              </div>
            </div>
            <div class="p-5">
              <h3
                class="text-base font-semibold mb-2 truncate"
                :style="cardTitleStyle"
              >
                {{ article.title }}
              </h3>
              <p
                class="text-sm mb-3 line-clamp-2"
                :style="cardSummaryStyle"
              >
                {{ article.summary }}
              </p>
              <div class="flex items-center gap-3 text-xs" :style="cardMetaStyle">
                <span class="truncate">{{ article.author }}</span>
                <span class="shrink-0">{{ article.publishDate }}</span>
              </div>
            </div>
          </article>

          <!-- Small Cards -->
          <article
            v-for="article in smallArticles"
            :key="article.id"
            class="rounded-lg p-4 transition-all duration-160 cursor-pointer"
            :style="smallCardStyle(article)"
            @mouseenter="handleCardHover(article)"
            @mouseleave="handleCardLeave(article)"
            @click="handleArticleClick(article.id)"
          >
            <div class="flex items-start gap-3">
              <div
                class="shrink-0 w-2 h-2 rounded-full mt-1.5"
                :style="cardDotStyle(article)"
              ></div>
              <div class="min-w-0 flex-1">
                <div class="flex items-center gap-2 mb-1">
                  <span
                    class="px-1.5 py-0.5 rounded text-[9px] font-bold"
                    :style="smallCategoryTagStyle(article)"
                  >
                    {{ article.category }}
                  </span>
                </div>
                <h4
                  class="text-sm font-semibold mb-1 truncate"
                  :style="smallCardTitleStyle"
                >
                  {{ article.title }}
                </h4>
                <p
                  class="text-xs line-clamp-2"
                  :style="smallCardSummaryStyle"
                >
                  {{ article.summary }}
                </p>
                <div class="flex items-center justify-between mt-2">
                  <span class="text-[10px]" :style="smallCardTimeStyle">
                    {{ article.timeAgo }}
                  </span>
                  <span class="text-[10px] font-bold" :style="smallCardXPStyle">
                    +{{ article.xpReward }} XP
                  </span>
                </div>
              </div>
            </div>
          </article>
        </div>
      </div>

      <!-- Sidebar -->
      <aside class="hidden lg:block">
        <slot name="sidebar" />
      </aside>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useThemeStore } from '@/stores/themeStore'
import type { Article } from '@/types'

interface Props {
  articles: Article[]
}

const props = defineProps<Props>()
const router = useRouter()
const themeStore = useThemeStore()

const handleArticleClick = (articleId: string) => {
  router.push(`/article/${articleId}`)
}

const hoveredCards = ref<Set<string>>(new Set())

const sectionTitle = computed(() => {
  const titles = {
    nexus: '精选情报',
    comiket: '漫画面板网格',
    ironcore: 'INTEL FEED'
  }
  return titles[themeStore.currentTheme]
})

const sectionTag = computed(() => {
  const tags = {
    nexus: 'DAILY PICK',
    comiket: 'FEATURED',
    ironcore: 'CLASSIFIED & VERIFIED'
  }
  return tags[themeStore.currentTheme]
})

const featuredArticles = computed(() => props.articles.slice(0, 2))
const smallArticles = computed(() => props.articles.slice(2, 6))

const gridStyle = computed(() => ({
  gridTemplateColumns: themeStore.currentTheme === 'ironcore' ? '1fr 300px' : '1fr 320px'
}))

const bentoGridStyle = computed(() => ({
  gridTemplateColumns: '1fr 1fr'
}))

const headerAccentStyle = computed(() => ({
  background: `var(--color-primary)`
}))

const headerStyle = computed(() => ({
  fontFamily: `var(--font-display)`,
  color: `var(--color-textPrimary)`
}))

const headerTagStyle = computed(() => ({
  fontFamily: `var(--font-mono)`,
  color: `var(--color-primary)`,
  background: `var(--color-primary50)`
}))

const largeCardStyle = computed(() => (article: Article) => {
  const isHovered = hoveredCards.value.has(article.id)
  
  return {
    background: `var(--color-bgCard)`,
    border: isHovered ? '1px solid var(--color-primary)' : '1px solid var(--color-borderSubtle)',
    gridColumn: 'span 1',
    boxShadow: isHovered ? '0 0 20px rgba(0, 240, 255, 0.08)' : 'none'
  }
})

const imageContainerStyle = computed(() => ({
  background: `var(--color-bgTertiary)`
}))

const categoryTagStyle = computed(() => (article: Article) => {
  const categoryColors = {
    tech: { bg: 'var(--color-primary)', color: 'var(--color-textInverse)' },
    game: { bg: 'var(--color-secondary)', color: 'white' },
    hardware: { bg: 'var(--color-primary50)', color: 'var(--color-primary)' },
    anime: { bg: 'var(--color-secondary50)', color: 'var(--color-secondary)' }
  }
  
  const colors = categoryColors[article.category as keyof typeof categoryColors]
  return {
    fontFamily: `var(--font-mono)`,
    background: colors?.bg || 'var(--color-primary)',
    color: colors?.color || 'var(--color-textInverse)'
  }
})

const xpTagStyle = computed(() => ({
  fontFamily: `var(--font-mono)`,
  background: 'rgba(0, 255, 136, 0.15)',
  color: 'var(--state-success)'
}))

const cardTitleStyle = computed(() => ({
  color: `var(--color-textPrimary)`
}))

const cardSummaryStyle = computed(() => ({
  color: `var(--color-textSecondary)`,
  lineHeight: 1.6
}))

const cardMetaStyle = computed(() => ({
  color: `var(--color-textTertiary)`
}))

const smallCardStyle = computed(() => (article: Article) => {
  const isHovered = hoveredCards.value.has(article.id)
  
  return {
    background: `var(--color-bgCard)`,
    border: isHovered ? '1px solid var(--color-border)' : '1px solid var(--color-borderSubtle)'
  }
})

const cardDotStyle = computed(() => (article: Article) => {
  const dotColors = {
    tech: 'var(--color-primary)',
    game: 'var(--color-secondary)',
    hardware: 'var(--state-info)',
    anime: 'var(--state-warning)'
  }
  
  return {
    background: dotColors[article.category as keyof typeof dotColors] || 'var(--color-primary)'
  }
})

const smallCategoryTagStyle = computed(() => (article: Article) => {
  const categoryStyles = {
    tech: { bg: 'var(--color-primary50)', color: 'var(--color-primary)' },
    game: { bg: 'var(--color-secondary50)', color: 'var(--color-secondary)' },
    hardware: { bg: 'rgba(0, 170, 255, 0.1)', color: 'var(--state-info)' },
    anime: { bg: 'rgba(255, 170, 0, 0.1)', color: 'var(--state-warning)' }
  }
  
  const style = categoryStyles[article.category as keyof typeof categoryStyles]
  return {
    fontFamily: `var(--font-mono)`,
    background: style?.bg || 'var(--color-primary50)',
    color: style?.color || 'var(--color-primary)'
  }
})

const smallCardTitleStyle = computed(() => ({
  color: `var(--color-textPrimary)`
}))

const smallCardSummaryStyle = computed(() => ({
  color: `var(--color-textSecondary)`
}))

const smallCardTimeStyle = computed(() => ({
  color: `var(--color-textTertiary)`
}))

const smallCardXPStyle = computed(() => ({
  fontFamily: `var(--font-mono)`,
  color: 'var(--state-success)'
}))

const handleCardHover = (article: Article) => {
  hoveredCards.value.add(article.id)
}

const handleCardLeave = (article: Article) => {
  hoveredCards.value.delete(article.id)
}
</script>