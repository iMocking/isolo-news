<template>
  <section class="max-w-5xl mx-auto px-6 pb-12">
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center gap-3">
        <span class="w-1 h-5 rounded-sm" :style="{ background: 'var(--state-success)' }"></span>
        <h2 class="text-xl font-semibold tracking-wider"
            :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }">
          {{ $t('profile.panels.favorites') }}
        </h2>
      </div>
      <span class="text-xs" :style="{ fontFamily: 'var(--font-mono)', color: 'var(--color-textTertiary)' }">
        {{ $t('profile.panels.favoriteCount', { count: articles.length }) }}
      </span>
    </div>

    <!-- 有收藏时 -->
    <div v-if="articles.length > 0" class="space-y-3">
      <div v-for="article in articles" :key="article.id"
           class="rounded-lg p-4 cursor-pointer transition-all duration-150 hover:-translate-y-0.5"
           :style="cardStyle"
           @click="router.push(`/article/${article.id}`)">
        <div class="flex items-start gap-3">
          <div class="min-w-0 flex-1">
            <div class="flex items-center gap-2 mb-2">
              <span class="px-2 py-0.5 rounded text-[10px] font-bold"
                    :style="{ background: 'var(--color-primary50)', color: 'var(--color-primary)', fontFamily: 'var(--font-mono)' }">
                {{ typeof article.category === 'string' ? article.category : article.category?.name }}
              </span>
              <span v-if="article.xpReward" class="text-[10px]"
                    :style="{ color: 'var(--state-success)', fontFamily: 'var(--font-mono)' }">
                +{{ article.xpReward }} XP
              </span>
            </div>
            <h3 class="text-sm font-semibold mb-1 truncate"
                :style="{ color: 'var(--color-textPrimary)' }">{{ article.title }}</h3>
            <p class="text-xs line-clamp-1"
               :style="{ color: 'var(--color-textSecondary)' }">{{ article.summary }}</p>
          </div>
          <ArrowRight class="w-4 h-4 shrink-0 mt-2" :style="{ color: 'var(--color-textTertiary)' }" />
        </div>
      </div>
    </div>

    <!-- 空状态 -->
    <EmptyState v-else type="empty" />
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { ArrowRight } from 'lucide-vue-next'
import EmptyState from '@/components/base/EmptyState.vue'
import { useCardStyles } from '@/hooks/useCardStyles'
import type { ArticleVO } from '@/api/articles'

const props = defineProps<{
  articles: ArticleVO[]
}>()

const router = useRouter()
const { getCardStyle } = useCardStyles()

const cardStyle = computed(() => getCardStyle('article', false))
</script>
