import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { UserState, Quest, Achievement } from '@/types'

export const useUserStore = defineStore('user', () => {
  const isLoggedIn = ref(false)
  const playerName = ref('NEXUS指挥官')
  const level = ref(42)
  const xp = ref(3200)
  const maxXp = ref(5000)
  const readArticles = ref(128)
  const loginDays = ref(7)
  const rank = ref(2847)
  const title = ref('赛博先锋 // 资讯猎人')
  const tags = ref(['赛博朋克', '硬件极客', 'AI观察者'])

  const achievements = ref<Achievement[]>([
    {
      id: 'ach-1',
      name: '初次阅读',
      description: '完成第一篇文章阅读',
      icon: 'book',
      unlockedAt: '2026-01-15'
    },
    {
      id: 'ach-2',
      name: '连续登录7天',
      description: '连续登录7天',
      icon: 'calendar',
      unlockedAt: '2026-07-06'
    },
    {
      id: 'ach-3',
      name: '科技爱好者',
      description: '阅读10篇科技资讯',
      icon: 'cpu',
      unlockedAt: '2026-03-20'
    }
  ])

  const dailyQuests = ref<Quest[]>([
    {
      id: 'quest-1',
      title: '阅读3篇科技资讯',
      description: '今日任务',
      xpReward: 300,
      progress: 1,
      target: 3,
      status: 'in_progress',
      icon: 'BookOpen'
    },
    {
      id: 'quest-2',
      title: '完成每日签到',
      description: '每日签到奖励',
      xpReward: 50,
      progress: 1,
      target: 1,
      status: 'completed',
      icon: 'Check'
    },
    {
      id: 'quest-3',
      title: '评论一篇游戏评测',
      description: '互动任务',
      xpReward: 200,
      progress: 0,
      target: 1,
      status: 'not_started',
      icon: 'MessageSquare'
    }
  ])

  const achievementCount = ref(42)
  const favoriteArticles = ref<string[]>([])

  const login = () => {
    isLoggedIn.value = true
  }

  const logout = () => {
    isLoggedIn.value = false
  }

  const toggleFavorite = (articleId: string) => {
    const index = favoriteArticles.value.indexOf(articleId)
    if (index > -1) {
      favoriteArticles.value.splice(index, 1)
    } else {
      favoriteArticles.value.push(articleId)
    }
  }

  const isFavorite = (articleId: string) => {
    return favoriteArticles.value.includes(articleId)
  }

  return {
    isLoggedIn,
    playerName,
    level,
    xp,
    maxXp,
    readArticles,
    loginDays,
    rank,
    title,
    tags,
    achievements,
    dailyQuests,
    achievementCount,
    favoriteArticles,
    login,
    logout,
    toggleFavorite,
    isFavorite
  }
})