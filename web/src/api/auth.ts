/**
 * 认证相关 API
 */

import apiClient, { type ApiResponse } from './index'

/** 用户视图对象 */
export interface UserVO {
  id: string
  email: string
  playerName: string
  level: number
  xp: number
  maxXp: number
  title: string
  tags: string[]
  readArticles: number
  loginDays: number
  avatarUrl?: string
  themePreference: string
  achievementCount: number
}

/** 登录响应 */
export interface LoginResponse {
  accessToken: string
  refreshToken: string
  expiresIn: number
  user: UserVO
}

/** 用户注册 */
export async function register(params: {
  email: string
  password: string
  confirmPassword: string
  playerName: string
}): Promise<ApiResponse<UserVO>> {
  const res = await apiClient.post('/auth/register', params)
  return res.data
}

/** 用户登录 */
export async function login(params: {
  email: string
  password: string
  rememberMe?: boolean
}): Promise<ApiResponse<LoginResponse>> {
  const res = await apiClient.post('/auth/login', params)
  return res.data
}

/** 发送邮箱验证码 */
export async function sendEmailCode(email: string): Promise<ApiResponse<{ message: string }>> {
  const res = await apiClient.post('/auth/email/code', { email })
  return res.data
}

/** 验证邮箱验证码 */
export async function verifyEmailCode(email: string, code: string): Promise<ApiResponse<{ verified: boolean }>> {
  const res = await apiClient.post('/auth/email/verify', { email, code })
  return res.data
}
