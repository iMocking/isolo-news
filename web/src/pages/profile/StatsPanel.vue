<template>
  <section class="max-w-5xl mx-auto px-6 py-12">
    <div class="flex items-center gap-3 mb-6">
      <span class="w-1 h-5 rounded-sm" :style="{ background: 'var(--color-primary)' }"></span>
      <h2 class="text-xl font-semibold tracking-wider"
          :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }">
        {{ $t('profile.panels.ability') }}
      </h2>
    </div>
    <div class="grid grid-cols-2 md:grid-cols-3 gap-4">
      <div v-for="stat in stats" :key="stat.label"
           class="rounded-lg p-5 transition-all duration-150 cursor-pointer hover:-translate-y-0.5"
           :style="statCardStyle">
        <div class="flex items-center gap-2.5 mb-3">
          <div class="w-8 h-8 rounded-md flex items-center justify-center"
               :style="{ background: stat.color === 'primary' ? 'var(--color-primary50)' : 'var(--color-secondary50)' }">
            <component :is="stat.icon" class="w-4 h-4"
                      :style="{ color: stat.color === 'primary' ? 'var(--color-primary)' : 'var(--color-secondary)' }" />
          </div>
          <span class="text-xs"
                :style="{ fontFamily: 'var(--font-mono)', color: 'var(--color-textTertiary)' }">
            {{ stat.label }}
          </span>
        </div>
        <span class="text-2xl md:text-3xl font-bold tabular-nums"
              :style="{
                fontFamily: 'var(--font-display)',
                color: stat.color === 'secondary' ? 'var(--color-secondary)' : 'var(--color-textPrimary)'
              }">
          {{ stat.value }}
        </span>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { BookOpen, Trophy, Flame, BarChart3, MessageSquare, Bookmark } from 'lucide-vue-next'
import { useUserStore } from '@/stores/userStore'
import { useCardStyles } from '@/hooks/useCardStyles'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const userStore = useUserStore()
const { getCardStyle } = useCardStyles()

interface StatItem {
  icon: any
  label: string
  value: string | number
  color: string
}

const stats = computed<StatItem[]>(() => [
  { icon: BookOpen, label: t('profile.stats.0'), value: userStore.readArticles, color: 'primary' },
  { icon: Trophy, label: t('profile.stats.1'), value: userStore.achievementCount, color: 'secondary' },
  { icon: Flame, label: t('profile.stats.2'), value: `${userStore.loginDays}${t('profile.unitDays')}`, color: 'primary' },
  { icon: BarChart3, label: t('profile.stats.3'), value: `#${userStore.rank}`, color: 'primary' },
  { icon: MessageSquare, label: t('profile.stats.4'), value: userStore.achievementCount, color: 'secondary' },
  { icon: Bookmark, label: t('profile.stats.5'), value: userStore.favoriteArticles.length, color: 'primary' }
])

const statCardStyle = computed(() => getCardStyle('stat', false))
</script>
