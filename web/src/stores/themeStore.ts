import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { ThemeName, ThemeConfig } from '@/types'
import { themes, nexusTheme } from '@/config/themes'

/** localStorage 中存储主题偏好的键名 */
const THEME_STORAGE_KEY = 'isolo-theme-preference'

/** 合法的主题名称列表，用于校验持久化数据 */
const VALID_THEME_NAMES: ThemeName[] = ['nexus', 'comiket', 'ironcore']

/**
 * 从 localStorage 读取持久化的主题名称
 * 若不存在或数据损坏则返回 null
 */
function loadThemeFromStorage(): ThemeName | null {
  try {
    const saved = localStorage.getItem(THEME_STORAGE_KEY)
    if (saved !== null && VALID_THEME_NAMES.includes(saved as ThemeName)) {
      return saved as ThemeName
    }
  } catch {
    // localStorage 不可用时（隐私模式、存储满等），静默忽略
  }
  return null
}

/**
 * 将主题名称持久化到 localStorage
 */
function saveThemeToStorage(themeName: ThemeName): void {
  try {
    localStorage.setItem(THEME_STORAGE_KEY, themeName)
  } catch {
    // localStorage 不可用时静默忽略，不阻塞渲染
  }
}

export const useThemeStore = defineStore('theme', () => {
  const currentTheme = ref<ThemeName>('nexus')

  const themeConfig = computed<ThemeConfig>(() => {
    return themes[currentTheme.value]
  })

  const setTheme = (themeName: ThemeName) => {
    currentTheme.value = themeName
    applyThemeToDOM(themes[themeName])
    // 同步持久化到 localStorage
    saveThemeToStorage(themeName)
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

  /**
   * 初始化主题
   * 优先级：localStorage 持久化主题 > NEXUS 默认主题
   */
  const initializeTheme = () => {
    const savedTheme = loadThemeFromStorage()
    if (savedTheme !== null) {
      currentTheme.value = savedTheme
      applyThemeToDOM(themes[savedTheme])
    } else {
      currentTheme.value = 'nexus'
      applyThemeToDOM(nexusTheme)
    }
  }

  return {
    currentTheme,
    themeConfig,
    setTheme,
    initializeTheme
  }
})