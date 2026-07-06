/**
 * 用户相关 API
 */

import apiClient, { type ApiResponse } from './index'
import type { UserVO } from './auth'
import type { ArticleVO } from './articles'

/** 成就视图对象 */
export interface AchievementVO {
  id: string
  name: string
  description: string
  icon: string
  unlockedAt?: number
}

/** 任务视图对象 */
export interface QuestVO {
  id: string
  title: string
  description: string
  xpReward: number
  progress: number
  target: number
  status: 'not_started' | 'in_progress' | 'completed'
  icon?: string
}

/** 获取用户信息 */
export async function fetchProfile(): Promise<ApiResponse<UserVO>> {
  const res = await apiClient.get('/user/profile')
  return res.data
}

/** 更新用户信息 */
export async function updateProfile(updates: Record<string, unknown>): Promise<ApiResponse<UserVO>> {
  const res = await apiClient.put('/user/profile', updates)
  return res.data
}

/** 获取收藏列表 */
export async function fetchFavorites(): Promise<ApiResponse<ArticleVO[]>> {
  const res = await apiClient.get('/user/favorites')
  return res.data
}

/** 获取成就列表 */
export async function fetchAchievements(): Promise<ApiResponse<AchievementVO[]>> {
  const res = await apiClient.get('/user/achievements')
  return res.data
}

/** 获取每日任务 */
export async function fetchQuests(): Promise<ApiResponse<QuestVO[]>> {
  const res = await apiClient.get('/user/quests')
  return res.data
}

/** 获取用户统计数据 */
export async function fetchStats(): Promise<ApiResponse<{
  readArticles: number
  achievementCount: number
  loginDays: number
  level: number
  xp: number
  maxXp: number
  title: string
  playerName: string
}>> {
  const res = await apiClient.get('/user/stats')
  return res.data
}

/** 获取周排行 */
export async function fetchLeaderboard(): Promise<ApiResponse<Array<{ title: string; views: string }>>> {
  const res = await apiClient.get('/leaderboard/weekly')
  return res.data
}
