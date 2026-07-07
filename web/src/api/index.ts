/**
 * API 客户端封装
 * 基于 axios，自动处理 Token 刷新、统一错误处理
 */

import axios, { type AxiosInstance, type AxiosResponse, type InternalAxiosRequestConfig } from 'axios'

/** API 基础路径，优先从环境变量读取 */
const API_BASE = import.meta.env.VITE_API_BASE_URL || '/api/v1'

/** API 统一响应格式 */
export interface ApiResponse<T = unknown> {
  code: number
  message: string
  data: T
}

/** 分页响应 */
export interface PaginatedData<T> {
  items: T[]
  total: number
  page: number
  page_size: number
  total_pages: number
}

/** 创建 axios 实例 */
const apiClient: AxiosInstance = axios.create({
  baseURL: API_BASE,
  timeout: 8000,
  headers: {
    'Content-Type': 'application/json',
  },
})

/** 是否正在刷新 Token 的锁 */
let isRefreshing = false
/** 等待 Token 刷新的请求队列 */
let pendingRequests: Array<{
  resolve: (token: string) => void
  reject: (err: unknown) => void
}> = []

/** 刷新 Token */
async function refreshAccessToken(): Promise<string> {
  const refreshToken = localStorage.getItem('refresh_token')
  if (!refreshToken) {
    throw new Error('无 refresh_token')
  }
  const res = await axios.post(`${API_BASE}/auth/refresh`, {
    refreshToken,
  })
  const { accessToken, refreshToken: newRefreshToken } = res.data.data
  localStorage.setItem('access_token', accessToken)
  if (newRefreshToken) {
    localStorage.setItem('refresh_token', newRefreshToken)
  }
  return accessToken
}

/** 清除登录态 */
function clearAuth() {
  localStorage.removeItem('access_token')
  localStorage.removeItem('refresh_token')
  // 使用 window.location 跳转确保完全重置
  if (window.location.pathname !== '/login') {
    window.location.href = '/login'
  }
}

/** 请求拦截器：自动注入 Token */
apiClient.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const token = localStorage.getItem('access_token')
    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => Promise.reject(error)
)

/**
 * 构建降级错误响应（网络/超时错误时返回，不抛异常）
 */
function buildErrorResponse(message: string): { data: ApiResponse<null> } {
  return {
    data: { code: -1, message, data: null },
  }
}

/** 响应拦截器：统一处理错误码 + 自动刷新 Token */
apiClient.interceptors.response.use(
  (response: AxiosResponse<ApiResponse>) => {
    const { data } = response
    if (data.code !== 0) {
      console.warn(`[API Error] ${data.message} (code: ${data.code})`)
    }
    return response
  },
  async (error) => {
    const originalRequest = error.config

    // 401 且尚未重试 → 尝试刷新 Token
    if (error.response?.status === 401 && !originalRequest._retry) {
      // 如果是刷新 Token 请求本身失败，直接清除登录态
      if (originalRequest.url?.includes('/auth/refresh')) {
        clearAuth()
        return Promise.reject(error)
      }

      // 正在刷新中 → 排队等待
      if (isRefreshing) {
        return new Promise<string>((resolve, reject) => {
          pendingRequests.push({ resolve, reject })
        }).then((newToken) => {
          originalRequest.headers.Authorization = `Bearer ${newToken}`
          return apiClient(originalRequest)
        })
      }

      originalRequest._retry = true
      isRefreshing = true

      try {
        const newToken = await refreshAccessToken()
        // 重放排队的请求
        pendingRequests.forEach((p) => p.resolve(newToken))
        pendingRequests = []
        originalRequest.headers.Authorization = `Bearer ${newToken}`
        return apiClient(originalRequest)
      } catch (refreshError) {
        clearAuth()
        pendingRequests.forEach((p) => p.reject(refreshError))
        pendingRequests = []
        return Promise.reject(refreshError)
      } finally {
        isRefreshing = false
      }
    }

    // === 以下所有错误均返回降级响应，绝不抛异常 ===
    // 因此所有调用方无需 try/catch 也能安全消费

    // 超时
    if (error.code === 'ECONNABORTED') {
      console.warn(`[API] 请求超时: ${originalRequest.url}`)
      return buildErrorResponse('请求超时，请检查网络连接')
    }

    // 无响应（网络断开等）
    if (!error.response) {
      console.warn(`[API] 网络不可用: ${originalRequest.url}`)
      return buildErrorResponse('网络连接失败，请检查网络')
    }

    // 其他 HTTP 错误（403/500 等）
    const { status } = error.response
    console.warn(`[API] HTTP ${status}: ${originalRequest.url}`)
    return buildErrorResponse(`服务器错误 (${status})`)
  }
)

export default apiClient
