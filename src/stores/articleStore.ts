import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Article, ArticleCategory } from '@/types'

export const useArticleStore = defineStore('article', () => {
  const articles = ref<Article[]>([
    {
      id: 'article-1',
      title: 'GPT-5来了:AI将如何改变游戏开发',
      summary: '深度解析GPT-5在游戏叙事、NPC交互和程序化生成领域的革命性突破,以及其对独立开发者生态的影响。',
      content: '详细内容...',
      category: 'tech',
      author: '赛博观察者',
      publishDate: '2026-07-06',
      readTime: 8,
      xpReward: 500,
      imageUrl: '/images/image_0_yi19x4.jpg',
      tags: ['AI', '游戏开发', 'GPT-5'],
      isFeatured: true,
      timeAgo: '2小时前'
    },
    {
      id: 'article-2',
      title: '《黑神话2》首曝:虚幻引擎6实机演示',
      summary: '游戏科学最新力作首次公开实机画面,采用虚幻引擎6打造,展现前所未有的东方神话视觉盛宴。',
      content: '详细内容...',
      category: 'game',
      author: '游戏前线',
      publishDate: '2026-07-05',
      readTime: 10,
      xpReward: 500,
      imageUrl: '/images/image_1_yi19x4.jpg',
      tags: ['黑神话', '虚幻引擎6', '游戏'],
      isFeatured: true,
      timeAgo: '1天前'
    },
    {
      id: 'article-3',
      title: 'AMD RX 9070 XT深度评测',
      summary: '新一代RDNA 5架构显卡实测,4K光追性能对比NVIDIA旗舰。',
      content: '详细内容...',
      category: 'hardware',
      author: '硬件频道',
      publishDate: '2026-07-06',
      readTime: 12,
      xpReward: 200,
      imageUrl: '/images/image_2_yi19x4.jpg',
      tags: ['AMD', '显卡', '硬件评测'],
      isFeatured: false,
      timeAgo: '3小时前'
    },
    {
      id: 'article-4',
      title: '《葬送的芙莉莲》第二季定档',
      summary: '人气奇幻动画续作确认秋季开播,原班制作团队回归。',
      content: '详细内容...',
      category: 'anime',
      author: '动漫情报',
      publishDate: '2026-07-06',
      readTime: 5,
      xpReward: 150,
      imageUrl: '/images/image_3_yi19x4.jpg',
      tags: ['芙莉莲', '动画', '二次元'],
      isFeatured: false,
      timeAgo: '5小时前'
    },
    {
      id: 'article-5',
      title: '苹果Vision Pro 2正式发布',
      summary: '第二代空间计算设备重量减半,售价下探至2499美元。',
      content: '详细内容...',
      category: 'tech',
      author: '科技快报',
      publishDate: '2026-07-06',
      readTime: 6,
      xpReward: 200,
      imageUrl: '/images/image_4_yi19x4.jpg',
      tags: ['苹果', 'VR', '科技'],
      isFeatured: false,
      timeAgo: '8小时前'
    },
    {
      id: 'article-6',
      title: 'Steam夏季特卖最佳折扣推荐',
      summary: '精选30款不可错过的折扣游戏,从独立佳作到3A大作。',
      content: '详细内容...',
      category: 'game',
      author: '游戏推荐',
      publishDate: '2026-07-06',
      readTime: 8,
      xpReward: 100,
      imageUrl: '/images/image_5_yi19x4.jpg',
      tags: ['Steam', '折扣', '游戏'],
      isFeatured: false,
      timeAgo: '12小时前'
    }
  ])

  const currentCategory = ref<ArticleCategory | 'all'>('all')
  const searchQuery = ref('')

  const trendingTopics = ref([
    { name: 'PS6首发', icon: 'flame', color: 'stateError' },
    { name: 'AI绘画革命', icon: 'bot', color: 'primary' },
    { name: '原神5.0', icon: 'gamepad-2', color: 'stateWarning' },
    { name: 'RTX5090评测', icon: 'cpu', color: 'stateSuccess' },
    { name: '量子芯片突破', icon: 'atom', color: 'secondary' }
  ])

  return {
    articles,
    currentCategory,
    searchQuery,
    trendingTopics
  }
})