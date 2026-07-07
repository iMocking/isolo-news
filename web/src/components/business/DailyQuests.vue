<template>
  <section class="max-w-7xl mx-auto px-6 pb-16">
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

    <div class="grid gap-4 md:grid-cols-3">
      <div
        v-for="quest in quests"
        :key="quest.id"
        class="p-5 transition-all duration-160"
        :style="questCardStyle(quest)"
        @mouseenter="handleQuestHover(quest)"
        @mouseleave="handleQuestLeave(quest)"
      >
        <div class="flex items-center justify-between mb-3">
          <div class="flex items-center gap-2">
            <div
              class="w-8 h-8 rounded-md flex items-center justify-center"
              :style="questIconBoxStyle(quest)"
            >
              <component :is="quest.icon" class="w-4 h-4" :style="questIconStyle(quest)" />
            </div>
            <div>
              <h4
                class="text-sm font-semibold"
                :style="questTitleStyle(quest)"
              >
                {{ quest.title }}
              </h4>
              <span class="text-[10px]" :style="questProgressTextStyle(quest)">
                {{ quest.statusText }}
              </span>
            </div>
          </div>
          <span class="text-[10px] font-bold" :style="questXPStyle(quest)">
            +{{ quest.xpReward }} XP
          </span>
        </div>
        <div class="h-1.5 rounded-full overflow-hidden" :style="questProgressBarStyle">
          <div class="h-full rounded-full" :style="questProgressFillStyle(quest)"></div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useThemeStore } from '@/stores/themeStore'
import { useUserStore } from '@/stores/userStore'
import { BookOpen, Check, MessageSquare } from 'lucide-vue-next'
import type { Quest } from '@/types'
import { useCardStyles } from '@/hooks/useCardStyles'

const { t } = useI18n()
const themeStore = useThemeStore()
const userStore = useUserStore()
const { getCardStyle } = useCardStyles()

const hoveredQuests = ref<Set<string>>(new Set())

const theme = computed(() => themeStore.currentTheme)

const sectionTitle = computed(() => t(`home.quests.title.${theme.value}`))

const sectionTag = computed(() => t(`home.quests.tag.${theme.value}`))

const iconMap: Record<string, any> = {
  BookOpen,
  Check,
  MessageSquare
}

const quests = computed(() => {
  return userStore.dailyQuests.map(quest => ({
    ...quest,
    icon: iconMap[quest.icon as string] || BookOpen,
    statusText: quest.status === 'completed' ? t('home.quests.statusCompleted') : 
                quest.status === 'in_progress' ? t('home.quests.progress', { progress: quest.progress, target: quest.target }) : t('home.quests.statusNotStarted')
  }))
})

const headerAccentStyle = computed(() => ({
  background: `var(--color-secondary)`
}))

const headerStyle = computed(() => ({
  fontFamily: `var(--font-display)`,
  color: `var(--color-textPrimary)`
}))

const headerTagStyle = computed(() => ({
  fontFamily: `var(--font-mono)`,
  color: `var(--color-secondary)`,
  background: `var(--color-secondary50)`
}))

const questCardStyle = computed(() => (quest: Quest) => {
  const isHovered = hoveredQuests.value.has(quest.id)
  const isCompleted = quest.status === 'completed'
  
  if (isCompleted) {
    return {
      ...getCardStyle('default', false),
      border: '1px solid rgba(0, 255, 136, 0.15)',
      opacity: 0.7
    }
  } else if (quest.status === 'in_progress') {
    return getCardStyle('default', isHovered)
  } else {
    return getCardStyle('default', isHovered)
  }
})

const questIconBoxStyle = computed(() => (quest: Quest) => {
  const isCompleted = quest.status === 'completed'
  const isInProgress = quest.status === 'in_progress'
  
  if (isCompleted) {
    return {
      background: 'rgba(0, 255, 136, 0.1)',
      border: '1px solid rgba(0, 255, 136, 0.15)'
    }
  } else if (isInProgress) {
    return {
      background: `var(--color-primary50)`,
      border: `1px solid var(--color-border)`
    }
  } else {
    return {
      background: `var(--color-bgPrimary)`,
      border: `1px solid var(--color-borderSubtle)`
    }
  }
})

const questIconStyle = computed(() => (quest: Quest) => {
  const isCompleted = quest.status === 'completed'
  const isInProgress = quest.status === 'in_progress'
  
  if (isCompleted) {
    return { color: 'var(--state-success)' }
  } else if (isInProgress) {
    return { color: 'var(--color-primary)' }
  } else {
    return { color: 'var(--color-textTertiary)' }
  }
})

const questTitleStyle = computed(() => (quest: Quest) => {
  const isCompleted = quest.status === 'completed'
  
  return {
    color: isCompleted ? 'var(--color-textSecondary)' : 'var(--color-textPrimary)',
    textDecoration: isCompleted ? 'line-through' : 'none'
  }
})

const questProgressTextStyle = computed(() => (quest: Quest) => {
  const isCompleted = quest.status === 'completed'
  
  return {
    fontFamily: `var(--font-mono)`,
    color: isCompleted ? 'var(--state-success)' : 'var(--color-textTertiary)'
  }
})

const questXPStyle = computed(() => (quest: Quest) => {
  const isCompleted = quest.status === 'completed'
  
  return {
    fontFamily: `var(--font-mono)`,
    color: isCompleted ? 'var(--color-textTertiary)' : 'var(--state-success)'
  }
})

const questProgressBarStyle = computed(() => ({
  background: `var(--color-bgPrimary)`
}))

const questProgressFillStyle = computed(() => (quest: Quest) => {
  const percentage = (quest.progress / quest.target) * 100
  const isCompleted = quest.status === 'completed'
  
  return {
    width: `${percentage}%`,
    background: isCompleted ? 'var(--state-success)' : 
                quest.status === 'in_progress' ? 'var(--color-primary)' : 'var(--color-textTertiary)'
  }
})

const handleQuestHover = (quest: Quest) => {
  hoveredQuests.value.add(quest.id)
}

const handleQuestLeave = (quest: Quest) => {
  hoveredQuests.value.delete(quest.id)
}
</script>
