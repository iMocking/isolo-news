<template>
  <aside class="hidden lg:block w-72 shrink-0 space-y-6 sticky top-24 self-start">
    <!-- Filter Panel -->
    <div class="p-5" :style="sidebarPanelStyle">
      <h3 class="text-sm font-semibold mb-4" style="font-family: var(--font-display); color: var(--color-primary);">{{ t('articleList.filter.title') }}</h3>
      <div class="space-y-3">
        <div>
          <label class="block text-xs mb-1.5" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">{{ t('articleList.filter.timeRange') }}</label>
          <select class="w-full px-3 h-9 text-sm outline-none" style="background: var(--color-bg-elevated); border: 1px solid var(--color-border-subtle); color: var(--color-text-secondary); border-radius: var(--radius-sm); font-family: var(--font-mono);">
            <option>{{ t('articleList.filter.allTime') }}</option>
            <option>{{ t('articleList.filter.within24h') }}</option>
            <option>{{ t('articleList.filter.withinWeek') }}</option>
            <option>{{ t('articleList.filter.withinMonth') }}</option>
          </select>
        </div>
        <div>
          <label class="block text-xs mb-1.5" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">{{ t('articleList.filter.sortBy') }}</label>
          <select class="w-full px-3 h-9 text-sm outline-none" style="background: var(--color-bg-elevated); border: 1px solid var(--color-border-subtle); color: var(--color-text-secondary); border-radius: var(--radius-sm); font-family: var(--font-mono);">
            <option>{{ t('articleList.filter.latest') }}</option>
            <option>{{ t('articleList.filter.mostComments') }}</option>
            <option>{{ t('articleList.filter.hottest') }}</option>
          </select>
        </div>
        <div>
          <label class="block text-xs mb-1.5" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">{{ t('articleList.filter.source') }}</label>
          <select class="w-full px-3 h-9 text-sm outline-none" style="background: var(--color-bg-elevated); border: 1px solid var(--color-border-subtle); color: var(--color-text-secondary); border-radius: var(--radius-sm); font-family: var(--font-mono);">
            <option>{{ t('articleList.filter.allSources') }}</option>
            <option>{{ t('articleList.filter.official') }}</option>
            <option>{{ t('articleList.filter.community') }}</option>
            <option>{{ t('articleList.filter.industry') }}</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Hot Tags -->
    <div class="p-5" :style="sidebarPanelStyle">
      <h3 class="text-sm font-semibold mb-4" style="font-family: var(--font-display); color: var(--color-primary);">{{ t('articleList.sidebar.hotTags') }}</h3>
      <div class="flex flex-wrap gap-2">
        <span v-for="tag in hotTags" :key="tag.name" class="px-2.5 py-1 text-xs whitespace-nowrap transition-all duration-150" :style="{
          background: tag.bg,
          border: tag.border,
          color: tag.color,
          borderRadius: 'var(--radius-sm)',
          fontFamily: 'var(--font-mono)'
        }">{{ tag.name }}</span>
      </div>
    </div>

    <!-- Weekly Leaderboard -->
    <div class="p-5" :style="sidebarPanelStyle">
      <h3 class="text-sm font-semibold mb-4" style="font-family: var(--font-display); color: var(--color-primary);">{{ t('articleList.sidebar.weeklyRank') }}</h3>
      <div class="space-y-3">
        <div v-for="(item, index) in leaderboard" :key="index" class="flex items-center gap-3 group cursor-pointer">
          <span class="w-5 h-5 flex items-center justify-center text-[10px] font-bold shrink-0" :style="{
            background: getLeaderboardColor(index),
            color: index < 3 ? 'var(--color-text-inverse)' : 'var(--color-text-tertiary)',
            borderRadius: 'var(--radius-full)',
            fontFamily: 'var(--font-display)'
          }">{{ index + 1 }}</span>
          <p class="text-sm truncate flex-1 min-w-0 transition-colors duration-150" style="color: var(--color-text-primary);">{{ item.title }}</p>
          <span class="text-[10px] shrink-0" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">{{ item.views }}</span>
        </div>
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { useCardStyles } from '@/hooks/useCardStyles'

const { t } = useI18n()

defineProps<{
  categories: Array<{ id: string; name: string; count: number }>
  currentCategory: string
  hotTags: Array<{ name: string; bg: string; border: string; color: string }>
  leaderboard: Array<{ title: string; views: string }>
}>()

defineEmits<{
  'update:currentCategory': [value: string]
}>()

const { getCardStyle } = useCardStyles()

const sidebarPanelStyle = computed(() => getCardStyle('panel', false))

const getLeaderboardColor = (index: number) => {
  const colors = ['var(--color-primary)', 'var(--color-secondary)', 'var(--color-state-info)', 'var(--color-state-warning)', 'var(--color-bg-elevated)']
  return colors[index] || 'var(--color-bg-elevated)'
}
</script>
