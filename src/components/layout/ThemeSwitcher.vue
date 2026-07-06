<template>
  <div class="hidden md:flex items-center gap-1.5">
    <button
      v-for="themeOption in themeOptions"
      :key="themeOption.value"
      class="px-3 py-1 text-xs font-semibold rounded-full whitespace-nowrap transition-all duration-150"
      :style="getButtonStyle(themeOption.value)"
      @click="handleThemeChange(themeOption.value)"
    >
      {{ themeOption.label }}
    </button>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useThemeStore } from '@/stores/themeStore'
import type { ThemeName } from '@/types'

const themeStore = useThemeStore()

const themeOptions: Array<{ value: ThemeName; label: string }> = [
  { value: 'nexus', label: 'NEXUS' },
  { value: 'comiket', label: 'COMIKET' },
  { value: 'ironcore', label: 'IRONCORE' }
]

const getButtonStyle = (themeName: ThemeName) => {
  const isActive = themeStore.currentTheme === themeName
  const theme = themeStore.themeConfig
  
  if (isActive) {
    return {
      background: `var(--color-primary50)`,
      border: `1px solid var(--color-primary)`,
      color: `var(--color-primary)`,
      fontFamily: `var(--font-mono)`
    }
  } else {
    return {
      background: 'transparent',
      border: `1px solid var(--color-borderSubtle)`,
      color: `var(--color-textTertiary)`,
      fontFamily: `var(--font-mono)`
    }
  }
}

const handleThemeChange = (themeName: ThemeName) => {
  themeStore.setTheme(themeName)
}
</script>