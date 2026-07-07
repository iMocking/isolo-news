/**
 * 用户状态管理
 * 对接后端 API
 */

import { defineStore } from 'pinia'
import { ref } from 'vue'
import { login, type LoginResponse, type UserVO } from '@/api/auth'
import type { AchievementVO, QuestVO } from '@/api/user'
import { fetchProfile, fetchAchievements, fetchQuests, fetchStats } from '@/api/user'

export const useUserStore = defineStore('user', () => {
  // 状态
  const isLoggedIn = ref(false)
  const currentUser = ref<UserVO | null>(null)
  const playerName = ref('')
  const level = ref(0)
  const xp = ref(0)
  const maxXp = ref(0)
  const readArticles = ref(0)
  const loginDays = ref(0)
  const rank = ref(0)
  const title = ref('')
  const tags = ref<string[]>([])
  const achievements = ref<AchievementVO[]>([])
  const dailyQuests = ref<QuestVO[]>([])
  const achievementCount = ref(0)
  const favoriteArticles = ref<string[]>([])
  const apiOffline = ref(false)

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
      apiOffline.value = false
      const [profileRes, achRes, questRes, statsRes] = await Promise.all([
        fetchProfile(),
        fetchAchievements().catch(() => null),
        fetchQuests().catch(() => null),
        fetchStats().catch(() => null),
      ])

      if (profileRes.code === 0 && profileRes.data) {
        applyUserData(profileRes.data)
      } else if (profileRes.code !== 0) {
        apiOffline.value = true
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
      console.warn('[userStore] 用户数据API不可用')
      apiOffline.value = true
    }
  }

  /** 重置 API 离线状态 */
  function resetApiState() {
    apiOffline.value = false
  }

  /** 应用用户数据到状态 */
  function applyUserData(data: UserVO) {
    currentUser.value = data
    playerName.value = data.playerName
    level.value = data.level
    xp.value = data.xp
    maxXp.value = data.maxXp
    title.value = data.title
    tags.value = data.tags ?? []
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
    playerName.value = ''
    level.value = 0
    xp.value = 0
    maxXp.value = 0
    title.value = ''
    tags.value = []
    readArticles.value = 0
    loginDays.value = 0
    rank.value = 0
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
    apiOffline,
    initialize,
    refreshUserData,
    loginByEmail,
    logout,
    toggleFavorite,
    isFavorite,
    resetApiState,
  }
})
