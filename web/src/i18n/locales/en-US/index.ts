import ui from './ui'
import pages from './pages'

/**
 * 英文语言包
 * 合并 UI 文本与页面文本，供 i18n 实例使用
 */
export default {
  ...ui,
  ...pages,
}
