/**
 * 用户状态管理
 * 对接后端 API，API 不可用时使用本地模拟数据
 */

import { defineStore } from 'pinia'
import { ref } from 'vue'
import { login, type LoginResponse, type UserVO } from '@/api/auth'
import type { AchievementVO, QuestVO } from '@/api/user'
import { fetchProfile, fetchAchievements, fetchQuests, fetchStats } from '@/api/user'

const DEFAULT_QUESTS: QuestVO[] = [
  {
    id: 'q1',
    title: '阅读3篇文章',
    description: '今日阅读3篇资讯文章',
    xpReward: 100,
    progress: 2,
    target: 3,
    status: 'in_progress',
    icon: 'BookOpen'
  },
  {
    id: 'q2',
    title: '发表评论',
    description: '在文章下发表1条评论',
    xpReward: 50,
    progress: 0,
    target: 1,
    status: 'not_started',
    icon: 'MessageSquare'
  },
  {
    id: 'q3',
    title: '每日签到',
    description: '完成每日签到领取奖励',
    xpReward: 30,
    progress: 1,
    target: 1,
    status: 'completed',
    icon: 'Check'
  }
]

export const useUserStore = defineStore('user', () => {
  // 状态
  const isLoggedIn = ref(false)
  const currentUser = ref<UserVO | null>(null)
  const playerName = ref('NEXUS指挥官')
  const level = ref(42)
  const xp = ref(3200)
  const maxXp = ref(5000)
  const readArticles = ref(128)
  const loginDays = ref(7)
  const rank = ref(2847)
  const title = ref('赛博先锋 // 资讯猎人')
  const tags = ref<string[]>(['赛博朋克', '硬件极客', 'AI观察者'])
  const achievements = ref<AchievementVO[]>([])
  const dailyQuests = ref<QuestVO[]>([...DEFAULT_QUESTS])
  const achievementCount = ref(42)
  const favoriteArticles = ref<string[]>([])

  /** 初始化用户状态（从本地 Token 恢复） */
  async function initialize() {
    const token = localStorage.getItem('access_token')
    if (token) {
      isLoggedIn.value = true
      await refreshUserData()
    }
  }

  /** 刷新用户数据 */
  async function refreshUserData() {
    try {
      const [profileRes, achRes, questRes, statsRes] = await Promise.all([
        fetchProfile(),
        fetchAchievements().catch(() => null),
        fetchQuests().catch(() => null),
        fetchStats().catch(() => null),
      ])

      if (profileRes.code === 0 && profileRes.data) {
        applyUserData(profileRes.data)
      }

      if (achRes && achRes.code === 0 && achRes.data) {
        achievements.value = achRes.data
      }

      if (questRes && questRes.code === 0 && questRes.data) {
        dailyQuests.value = questRes.data.map((q) => ({
          ...q,
          icon: q.icon || (q.status === 'completed' ? 'Check' : q.status === 'in_progress' ? 'BookOpen' : 'MessageSquare'),
          statusText: q.status === 'completed' ? '已完成' : `进度 ${q.progress}/${q.target}`,
        }))
      }

      if (statsRes && statsRes.code === 0 && statsRes.data) {
        // 补充统计数据
      }
    } catch {
      console.warn('[userStore] API不可用，使用本地数据')
    }
  }

  /** 应用用户数据到状态 */
  function applyUserData(data: UserVO) {
    currentUser.value = data
    playerName.value = data.playerName
    level.value = data.level
    xp.value = data.xp
    maxXp.value = data.maxXp
    title.value = data.title
    tags.value = data.tags
    readArticles.value = data.readArticles
    loginDays.value = data.loginDays
    achievementCount.value = data.achievementCount
  }

  /** 登录 */
  async function loginByEmail(email: string, password: string, rememberMe = false) {
    const res = await login({ email, password, rememberMe })
    if (res.code === 0 && res.data) {
      const { accessToken, refreshToken, user } = res.data
      localStorage.setItem('access_token', accessToken)
      localStorage.setItem('refresh_token', refreshToken)
      isLoggedIn.value = true
      applyUserData(user)
    }
    return res
  }

  /** 登出 */
  function logout() {
    isLoggedIn.value = false
    currentUser.value = null
    playerName.value = 'NEXUS指挥官'
    level.value = 42
    xp.value = 3200
    maxXp.value = 5000
    title.value = '赛博先锋 // 资讯猎人'
    tags.value = ['赛博朋克', '硬件极客', 'AI观察者']
    readArticles.value = 128
    loginDays.value = 7
    rank.value = 2847
    favoriteArticles.value = []
    localStorage.removeItem('access_token')
    localStorage.removeItem('refresh_token')
  }

  /** 切换收藏 */
  function toggleFavorite(articleId: string) {
    const index = favoriteArticles.value.indexOf(articleId)
    if (index > -1) {
      favoriteArticles.value.splice(index, 1)
    } else {
      favoriteArticles.value.push(articleId)
    }
  }

  /** 检查是否已收藏 */
  function isFavorite(articleId: string): boolean {
    return favoriteArticles.value.includes(articleId)
  }

  return {
    isLoggedIn,
    currentUser,
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
    initialize,
    refreshUserData,
    loginByEmail,
    logout,
    toggleFavorite,
    isFavorite,
  }
})
