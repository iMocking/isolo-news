import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { ThemeName, ThemeConfig } from '@/types'
import { themes, nexusTheme } from '@/config/themes'

export const useThemeStore = defineStore('theme', () => {
  const currentTheme = ref<ThemeName>('nexus')

  const themeConfig = computed<ThemeConfig>(() => {
    return themes[currentTheme.value]
  })

  const setTheme = (themeName: ThemeName) => {
    currentTheme.value = themeName
    applyThemeToDOM(themes[themeName])
  }

  const applyThemeToDOM = (config: ThemeConfig) => {
    const root = document.documentElement
    
    // Apply colors
    Object.entries(config.colors).forEach(([key, value]) => {
      root.style.setProperty(`--color-${key}`, value)
    })

    // Apply fonts
    Object.entries(config.fonts).forEach(([key, value]) => {
      root.style.setProperty(`--font-${key}`, value)
    })

    // Apply radius
    Object.entries(config.radius).forEach(([key, value]) => {
      root.style.setProperty(`--radius-${key}`, value)
    })
  }

  // Initialize theme
  const initializeTheme = () => {
    applyThemeToDOM(nexusTheme)
  }

  return {
    currentTheme,
    themeConfig,
    setTheme,
    initializeTheme
  }
})