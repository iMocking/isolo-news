/**
 * 资讯状态管理
 * 优先从 API 获取数据，API 不可用时使用本地模拟数据
 */

import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { ArticleCategory } from '@/types'
import type { ArticleVO } from '@/api/articles'
import { fetchArticles, fetchFeaturedArticles, fetchCategories, fetchTrendingTopics } from '@/api/articles'

const FALLBACK_IMAGES = [
  '/images/image_0_yi19x4.jpg',
  '/images/image_1_yi19x4.jpg',
  '/images/image_2_yi19x4.jpg',
  '/images/image_3_yi19x4.jpg',
  '/images/image_5_yi19x4.jpg',
]

function stripHtml(html: string): string {
  if (!html) return ''
  const tmp = document.createElement('div')
  tmp.innerHTML = html
  let text = tmp.textContent || tmp.innerText || ''
  text = text.replace(/\{\{.*?\}\}/g, '')
  text = text.replace(/\$json\./g, '')
  text = text.replace(/&nbsp;/g, ' ')
  text = text.replace(/\s+/g, ' ')
  return text.trim()
}

function processArticle(article: ArticleVO, index: number): ArticleVO {
  const processed = { ...article }
  
  if (!processed.imageUrl || processed.imageUrl === '/' || processed.imageUrl === '' || processed.imageUrl.startsWith('http') === false && !processed.imageUrl.startsWith('/')) {
    processed.imageUrl = FALLBACK_IMAGES[index % FALLBACK_IMAGES.length]
  }
  
  if (processed.summary) {
    processed.summary = stripHtml(processed.summary)
  }
  
  if (processed.summary && processed.summary.length > 120) {
    processed.summary = processed.summary.substring(0, 120) + '...'
  }
  
  if (!processed.summary || processed.summary.length < 5) {
    processed.summary = article.title
  }
  
  return processed
}

const MOCK_TRENDING_TOPICS = [
  { id: 1, name: 'PS6首发', icon: 'Flame', colorType: 'error' },
  { id: 2, name: 'AI绘画革命', icon: 'Bot', colorType: 'primary' },
  { id: 3, name: '原神5.0', icon: 'Gamepad2', colorType: 'warning' },
  { id: 4, name: 'RTX5090评测', icon: 'Cpu', colorType: 'success' },
  { id: 5, name: '量子芯片突破', icon: 'Atom', colorType: 'secondary' },
]

/** 本地模拟数据（API 不可用时兜底） */
const MOCK_ARTICLES: ArticleVO[] = [
  {
    id: 'demo-1',
    title: 'GPT-5来了:AI将如何改变游戏开发',
    summary: '深度解析GPT-5在游戏叙事、NPC交互和程序化生成领域的革命性突破。',
    content: '<p>2026年7月，OpenAI正式发布GPT-5，这一里程碑事件迅速在游戏开发行业引发了连锁反应。从程序化内容生成到智能NPC对话系统，从自动Bug检测到美术资源优化，GPT-5展现出的多模态推理能力远超业内预期。</p><p>据内部测试数据显示，GPT-5在代码生成准确率上较前代提升了47%，在复杂逻辑推理任务上的表现已接近高级工程师水平。</p><h2 style="color: var(--color-primary);">程序化内容生成的新纪元</h2><p>游戏世界构建历来是开发成本最高的环节之一。GPT-5的3D空间理解能力使得程序化生成从"随机拼凑"进化到了"语义驱动"。引擎能够根据剧情需要自动生成符合叙事逻辑的地形、建筑群落甚至天气系统，同时保持视觉风格的高度一致性。</p><h2 style="color: var(--color-primary);">智能NPC与动态叙事</h2><p>GPT-5的上下文窗口扩展至百万级Token，使得NPC能够维持跨章节的长期记忆，并根据玩家的行为模式动态调整交互策略。更令人瞩目的是，GPT-5展现出了跨模态的叙事理解能力，能同步控制角色的面部表情、肢体动作和语音语调。</p>',
    category: { slug: 'tech', name: 'AI前沿', color: '#00f0ff' },
    author: '赛博观察者',
    publishDate: Date.now() / 1000 - 7200,
    readTime: 8,
    xpReward: 500,
    imageUrl: '/images/image_0_yi19x4.jpg',
    tags: ['AI', '游戏开发', 'GPT-5'],
    isFeatured: true,
    viewCount: 2847,
    commentCount: 23,
    likeCount: 100,
  },
  {
    id: 'demo-2',
    title: '《黑神话2》首曝:虚幻引擎6实机演示',
    summary: '游戏科学最新力作首次公开实机画面，采用虚幻引擎6打造。',
    content: '<p>游戏科学工作室在昨晚的线上发布会上正式公布了《黑神话：悟空2》的首支实机演示预告。与初代作品不同，本作采用了虚幻引擎6进行开发，画面效果达到了前所未有的高度。</p><p>预告片中展示了全新的"轮回"系统——玩家可以在不同时间线之间穿梭，每一次选择都会改变关卡的布局和敌人的配置。这种动态关卡设计在AAA级游戏中尚属首次。</p><p>据游戏科学创始人冯骥透露，本作的开发团队规模已扩大至300人，开发周期预计为4年。首个可玩Demo将在明年ChinaJoy上提供试玩。</p>',
    category: { slug: 'game', name: '游戏', color: '#ff6b2b' },
    author: '游戏前线',
    publishDate: Date.now() / 1000 - 86400,
    readTime: 10,
    xpReward: 500,
    imageUrl: '/images/image_1_yi19x4.jpg',
    tags: ['黑神话', '虚幻引擎6', '游戏'],
    isFeatured: true,
    viewCount: 5231,
    commentCount: 56,
    likeCount: 230,
  },
  {
    id: 'demo-3',
    title: 'AMD RX 9070 XT深度评测',
    summary: '新一代RDNA 5架构显卡实测，4K光追性能对比NVIDIA旗舰。',
    content: '<p>AMD今日正式发布了基于RDNA 5架构的RX 9070 XT显卡。我们的评测数据显示，在4K分辨率下打开光线追踪，RX 9070 XT在《赛博朋克2077》中平均帧率达到87fps，相比上代提升了63%。</p><p>在AI算力方面，RX 9070 XT配备了全新的Matrix Core 5.0单元，FP16算力达到180 TFLOPS，在Stable Diffusion XL图像生成任务中比RTX 5090快12%。</p><p>售价方面，RX 9070 XT建议零售价为5999元，将于下月15日正式开售。首批上市的包括华硕、技嘉、微星等厂商的非公版产品。</p>',
    category: { slug: 'hardware', name: 'AI硬件', color: '#0aaaff' },
    author: '硬件频道',
    publishDate: Date.now() / 1000 - 10800,
    readTime: 12,
    xpReward: 200,
    imageUrl: '/images/image_2_yi19x4.jpg',
    tags: ['AMD', '显卡', '硬件评测'],
    isFeatured: false,
    viewCount: 1890,
    commentCount: 34,
    likeCount: 89,
  },
  {
    id: 'demo-4',
    title: '7月新番前瞻：这些续作你不能错过',
    summary: '夏季动画档强势来袭，多部人气作品续作集结登场。',
    content: '<p>随着7月的到来，新一季的动画番剧也即将与观众见面。本季度可谓续作云集，多部人气作品的最新季都将在这个夏天与大家见面。</p><p>从热血战斗到日常治愈，从科幻悬疑到校园恋爱，本季度的番剧类型丰富多样，相信每位观众都能找到自己喜欢的作品。</p>',
    category: { slug: 'anime', name: '二次元', color: '#ffaa00' },
    author: '二次元频道',
    publishDate: Date.now() / 1000 - 14400,
    readTime: 6,
    xpReward: 150,
    imageUrl: '/images/image_3_yi19x4.jpg',
    tags: ['新番', '动画', '7月新番'],
    isFeatured: false,
    viewCount: 1560,
    commentCount: 28,
    likeCount: 78,
  },
  {
    id: 'demo-5',
    title: 'Godot 4.3正式发布：性能大幅提升',
    summary: '开源游戏引擎迎来重大更新，3D渲染性能提升40%。',
    content: '<p>Godot团队今日正式发布了4.3版本，这一版本带来了大量的性能优化和新功能。其中最引人注目的是3D渲染性能的大幅提升，在部分场景中帧率提升可达40%以上。</p><p>此外，新版本还加入了对更多着色器特性的支持，优化了物理引擎的稳定性，并改进了编辑器的用户体验。</p>',
    category: { slug: 'godot', name: 'Godot游戏', color: '#22c55e' },
    author: '游戏开发',
    publishDate: Date.now() / 1000 - 18000,
    readTime: 7,
    xpReward: 180,
    imageUrl: '/images/image_5_yi19x4.jpg',
    tags: ['Godot', '游戏引擎', '开源'],
    isFeatured: false,
    viewCount: 980,
    commentCount: 15,
    likeCount: 65,
  },
  {
    id: 'demo-6',
    title: 'AI编程助手横评：谁才是程序员的最佳搭档',
    summary: '对比5款主流AI编程工具，从代码质量、生成速度等多维度评测。',
    content: '<p>随着AI编程工具的快速发展，程序员们有了越来越多的选择。从GitHub Copilot到Cursor，从Claude Code到GPT-4o，每款工具都有其独特的优势。</p><p>我们从代码生成质量、理解上下文能力、调试效率、支持语言等多个维度对主流AI编程工具进行了全面对比测试。</p>',
    category: { slug: 'ai_coding', name: 'AI编程', color: '#ff2d95' },
    author: '技术编辑',
    publishDate: Date.now() / 1000 - 21600,
    readTime: 9,
    xpReward: 220,
    imageUrl: '/images/image_0_yi19x4.jpg',
    tags: ['AI编程', '工具评测', '开发效率'],
    isFeatured: false,
    viewCount: 2100,
    commentCount: 42,
    likeCount: 120,
  },
]

export const useArticleStore = defineStore('article', () => {
  // 状态
  const articles = ref<ArticleVO[]>([])
  const currentCategory = ref<string>('all')
  const searchQuery = ref('')
  const total = ref(0)
  const loading = ref(false)
  const categories = ref<Array<{ id: string; name: string; count: number }>>([])
  const trendingTopics = ref<Array<{ id: number; name: string; icon: string; colorType: string }>>([])

  /** 初始化加载 */
  async function initialize() {
    await Promise.all([
      loadArticles(),
      loadCategories(),
    ])
  }

  /** 加载文章列表 */
  async function loadArticles(params?: {
    page?: number
    page_size?: number
    category?: string
    search?: string
    sort?: string
  }) {
    loading.value = true
    try {
      const res = await fetchArticles({
        page: params?.page || 1,
        page_size: params?.page_size || 20,
        category: params?.category || (currentCategory.value === 'all' ? undefined : currentCategory.value),
        search: params?.search || searchQuery.value || undefined,
        sort: (params?.sort as 'latest' | 'hot' | 'most_commented') || 'latest',
      })
      if (res.code === 0 && res.data) {
        articles.value = res.data.items.map((article, index) => processArticle(article, index))
        total.value = res.data.total
      }
    } catch {
      console.warn('[articleStore] API不可用，使用模拟数据')
      articles.value = MOCK_ARTICLES
      total.value = MOCK_ARTICLES.length
    } finally {
      loading.value = false
    }
  }

  /** 加载分类 */
  async function loadCategories() {
    try {
      const res = await fetchCategories()
      if (res.code === 0 && res.data) {
        categories.value = res.data.map((c) => ({
          id: c.slug,
          name: c.displayName,
          count: c.articleCount,
        }))
      }
    } catch {
      categories.value = [
        { id: 'all', name: '全部', count: 2847 },
        { id: 'tech', name: 'AI前沿', count: 1203 },
        { id: 'ai_robot', name: 'AI机器人', count: 456 },
        { id: 'ai_coding', name: 'AI编程', count: 342 },
        { id: 'hardware', name: 'AI硬件', count: 412 },
        { id: 'anime', name: '二次元', count: 376 },
        { id: 'godot', name: 'Godot游戏', count: 158 },
      ]
    }
  }

  /** 加载热门话题 */
  async function loadTrendingTopics() {
    try {
      const res = await fetchTrendingTopics()
      if (res.code === 0 && res.data && res.data.length > 0) {
        trendingTopics.value = res.data.map((t) => ({
          id: t.id,
          name: t.name,
          icon: t.icon,
          colorType: t.colorType,
        }))
      } else {
        trendingTopics.value = MOCK_TRENDING_TOPICS
      }
    } catch {
      console.warn('[articleStore] 话题API不可用，使用模拟数据')
      trendingTopics.value = MOCK_TRENDING_TOPICS
    }
  }

  /** 过滤后的文章 */
  const filteredArticles = computed(() => {
    if (!searchQuery.value && currentCategory.value === 'all') return articles.value
    return articles.value.filter((article) => {
      const matchCategory = currentCategory.value === 'all' || article.category.slug === currentCategory.value
      const matchSearch = !searchQuery.value ||
        article.title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
        article.summary.toLowerCase().includes(searchQuery.value.toLowerCase())
      return matchCategory && matchSearch
    })
  })

  return {
    // 状态
    articles,
    currentCategory,
    searchQuery,
    total,
    loading,
    categories,
    trendingTopics,
    // 方法
    initialize,
    loadArticles,
    loadCategories,
    loadTrendingTopics,
    filteredArticles,
  }
})
