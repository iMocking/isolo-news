<template>
  <section class="max-w-5xl mx-auto px-6 pb-12">
    <div class="flex items-center gap-3 mb-6">
      <span class="w-1 h-5 rounded-sm" :style="{ background: 'var(--color-secondary)' }"></span>
      <h2 class="text-xl font-semibold tracking-wider"
          :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }">
        {{ $t('profile.panels.achievements', { count: achievements.length }) }}
      </h2>
    </div>
    <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
      <div v-for="(achievement, index) in achievements" :key="index"
           class="rounded-lg p-5 transition-all duration-150" :style="achievementCardStyle(achievement)">
        <div class="w-12 h-12 mx-auto mb-3 rounded-lg flex items-center justify-center"
             :style="achievementIconBoxStyle(achievement)">
          <component :is="achievement.icon" class="w-6 h-6"
                    :style="{ color: achievement.unlocked ? 'var(--color-primary)' : 'var(--color-textTertiary)' }" />
        </div>
        <h3 class="text-sm font-semibold mb-1 text-center"
            :style="{ color: achievement.unlocked ? 'var(--color-textPrimary)' : 'var(--color-textTertiary)' }">
          {{ achievement.name }}
        </h3>
        <p class="text-xs text-center"
           :style="{ color: 'var(--color-textSecondary)', lineHeight: 1.5 }">
          {{ achievement.description }}
        </p>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useCardStyles } from '@/hooks/useCardStyles'

const props = defineProps<{
  achievements: Array<{
    name: string
    description: string
    icon: any
    unlocked: boolean
  }>
}>()

const { getCardStyle } = useCardStyles()

const achievementCardStyle = computed(() => (achievement: { unlocked: boolean }) => ({
  ...getCardStyle('default', false),
  opacity: achievement.unlocked ? 1 : 0.6
}))

const achievementIconBoxStyle = (achievement: { unlocked: boolean }) => ({
  background: achievement.unlocked ? 'var(--color-primary50)' : 'var(--color-bgTertiary)',
  border: `1px solid ${achievement.unlocked ? 'var(--color-border)' : 'var(--color-borderSubtle)'}`
})
</script>
