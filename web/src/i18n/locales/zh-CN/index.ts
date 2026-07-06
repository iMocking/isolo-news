import ui from './ui'
import pages from './pages'
import * as articleData from './articles'

/**
 * 中文语言包
 * 合并 UI 文本、页面文本与文章数据，供 i18n 实例使用
 */
export default {
  ...ui,
  ...pages,
  articleData
}
