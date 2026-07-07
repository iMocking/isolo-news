/**
 * 已读文章追踪 composable
 * 使用 localStorage 持久化已读文章 ID 列表
 * 模块级共享状态，所有组件共用同一个响应式数组
 */

import { ref } from 'vue'

const STORAGE_KEY = 'read_articles'

/** 从 localStorage 加载已读 ID 列表 */
function loadFromStorage(): string[] {
  try {
    const raw = localStorage.getItem(STORAGE_KEY)
    if (!raw) return []
    const parsed = JSON.parse(raw)
    return Array.isArray(parsed) ? parsed : []
  } catch {
    return []
  }
}

/** 保存到 localStorage */
function saveToStorage(ids: string[]) {
  try {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(ids))
  } catch {
    console.warn('[useReadTracker] 写入 localStorage 失败')
  }
}

/** 模块级共享响应式数组 */
const readIds = ref<string[]>(loadFromStorage())

export function useReadTracker() {
  /**
   * 标记指定文章为已读
   * @param id 文章 ID
   */
  function markAsRead(id: string) {
    if (!readIds.value.includes(id)) {
      readIds.value.push(id)
      saveToStorage(readIds.value)
    }
  }

  /**
   * 检查文章是否已读
   * @param id 文章 ID
   * @returns 是否已读
   */
  function isRead(id: string): boolean {
    return readIds.value.includes(id)
  }

  return {
    /** 已读文章 ID 列表（响应式） */
    readIds,
    /** 标记已读 */
    markAsRead,
    /** 查询是否已读 */
    isRead,
  }
}
