<template>
  <div class="rounded-lg p-5" :style="panelStyle">
    <!-- Header -->
    <div class="flex items-center gap-2 mb-5">
      <component :is="headerIcon" class="w-4 h-4" :style="iconStyle" />
      <h3
        class="text-sm font-semibold tracking-wide"
        :style="headerStyle"
      >
        {{ panelTitle }}
      </h3>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-2 gap-3 mb-5">
      <div
        v-for="stat in stats"
        :key="stat.label"
        class="rounded-md p-3 text-center"
        :style="statBoxStyle(stat)"
      >
        <div
          class="text-lg font-bold"
          :style="statValueStyle(stat)"
        >
          {{ stat.value }}
        </div>
        <div class="text-[10px] mt-1" :style="statLabelStyle">
          {{ stat.label }}
        </div>
      </div>
    </div>

    <!-- Level Progress -->
    <div class="mb-4">
      <div class="flex items-center justify-between mb-2">
        <span class="text-xs" :style="levelLabelStyle">Lv.{{ userLevel }}</span>
        <span class="text-[10px]" :style="xpProgressStyle">{{ currentXP }} / {{ maxXP }} XP</span>
      </div>
      <div class="h-1.5 rounded-full overflow-hidden" :style="progressBarStyle">
        <div class="h-full rounded-full" :style="progressFillStyle"></div>
      </div>
    </div>

    <!-- Scanline decoration -->
    <div class="h-px w-full" :style="scanlineStyle"></div>

    <!-- Player Info -->
    <div class="flex items-center gap-3 mt-4">
      <div
        class="w-10 h-10 rounded-full flex items-center justify-center text-sm font-bold shrink-0"
        :style="avatarStyle"
      >
        {{ avatarText }}
      </div>
      <div class="min-w-0 flex-1">
        <div class="text-sm font-medium truncate" :style="playerNameStyle">
          {{ playerName }}
        </div>
        <div class="text-[10px]" :style="playerTitleStyle">
          {{ playerTitle }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { useThemeStore } from '@/stores/themeStore'
import { useUserStore } from '@/stores/userStore'
import { BarChart3 } from 'lucide-vue-next'

const { t, tm } = useI18n()
const themeStore = useThemeStore()
const userStore = useUserStore()

const theme = computed(() => themeStore.currentTheme)

const panelTitle = computed(() => t(`home.player.title.${theme.value}`))

const headerIcon = computed(() => BarChart3)

const stats = computed(() => {
  const statNames = tm(`home.player.statsName.${theme.value}`) as unknown as string[]
  return [
    { label: statNames[0], value: theme.value === 'ironcore' ? userStore.readArticles : userStore.readArticles, colorType: 'primary' },
    { label: statNames[1], value: theme.value === 'ironcore' ? 'S+' : userStore.achievements.length, colorType: 'secondary' },
    { label: statNames[2], value: theme.value === 'ironcore' ? '87%' : userStore.loginDays, colorType: 'success' },
    { label: statNames[3], value: theme.value === 'ironcore' ? '2,450' : `#${userStore.rank}`, colorType: 'warning' }
  ]
})

const userLevel = computed(() => userStore.level)
const currentXP = computed(() => userStore.xp)
const maxXP = computed(() => 5000)
const playerName = computed(() => t(`profile.playerName.${theme.value}`))

const playerTitle = computed(() => t(`profile.playerTitle.${theme.value}`))

const avatarText = computed(() => {
  const texts = {
    nexus: 'NX',
    comiket: 'CO',
    ironcore: 'IR'
  }
  return texts[themeStore.currentTheme]
})

const panelStyle = computed(() => ({
  background: `var(--color-bgCard)`,
  border: `1px solid var(--color-border)`
}))

const iconStyle = computed(() => ({
  color: `var(--color-primary)`
}))

const headerStyle = computed(() => ({
  fontFamily: `var(--font-display)`,
  color: `var(--color-textPrimary)`
}))

const statBoxStyle = computed(() => (stat: any) => ({
  background: `var(--color-bgPrimary)`,
  border: `1px solid var(--color-borderSubtle)`
}))

const statValueStyle = computed(() => (stat: any) => {
  const colorMap = {
    primary: 'var(--color-primary)',
    secondary: 'var(--color-secondary)',
    success: 'var(--state-success)',
    warning: 'var(--state-warning)',
    info: 'var(--state-info)'
  }
  
  return {
    fontFamily: `var(--font-display)`,
    color: colorMap[stat.colorType as keyof typeof colorMap] || 'var(--color-primary)',
    fontVariantNumeric: 'tabular-nums'
  }
})

const statLabelStyle = computed(() => ({
  color: `var(--color-textTertiary)`
}))

const levelLabelStyle = computed(() => ({
  color: `var(--color-textSecondary)`
}))

const xpProgressStyle = computed(() => ({
  fontFamily: `var(--font-mono)`,
  color: `var(--color-textTertiary)`
}))

const progressBarStyle = computed(() => ({
  background: `var(--color-bgPrimary)`
}))

const progressFillStyle = computed(() => {
  const percentage = Math.min((currentXP.value / maxXP.value) * 100, 100)
  return {
    width: `${percentage}%`,
    background: `linear-gradient(90deg, var(--color-primaryDark), var(--color-primary))`
  }
})

const scanlineStyle = computed(() => ({
  background: `linear-gradient(90deg, var(--color-border), transparent)`
}))

const avatarStyle = computed(() => ({
  background: `linear-gradient(135deg, var(--color-primaryDark), var(--color-secondaryDark))`,
  color: `var(--color-textPrimary)`
}))

const playerNameStyle = computed(() => ({
  color: `var(--color-textPrimary)`
}))

const playerTitleStyle = computed(() => ({
  fontFamily: `var(--font-mono)`,
  color: `var(--color-textTertiary)`
}))
</script>