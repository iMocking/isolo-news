import { createI18n } from 'vue-i18n'
import zhCN from './locales/zh-CN'
import enUS from './locales/en-US'

/** localStorage 键名，用于持久化语言偏好 */
const LOCALE_STORAGE_KEY = 'isolo-locale'

/** 支持的语言列表 */
export type SupportedLocale = 'zh-CN' | 'en-US'

/** 从 localStorage 或浏览器偏好检测语言 */
function detectLocale(): SupportedLocale {
  try {
    const saved = localStorage.getItem(LOCALE_STORAGE_KEY) as SupportedLocale | null
    if (saved === 'zh-CN' || saved === 'en-US') return saved
  } catch {
    // localStorage 不可用，静默忽略
  }

  // 检测浏览器语言
  const browserLang = navigator.language?.toLowerCase() || ''
  if (browserLang.startsWith('zh')) return 'zh-CN'
  // 默认中文
  return 'zh-CN'
}

const i18n = createI18n({
  legacy: false,          // 使用 Composition API 模式
  locale: detectLocale(), // 检测初始语言
  fallbackLocale: 'zh-CN',
  messages: {
    'zh-CN': zhCN,
    'en-US': enUS
  }
})

/**
 * 切换语言并持久化
 * @param locale 目标语言
 */
export function setLocale(locale: SupportedLocale): void {
  i18n.global.locale.value = locale
  try {
    localStorage.setItem(LOCALE_STORAGE_KEY, locale)
  } catch {
    // 持久化失败时静默忽略
  }
}

export default i18n
