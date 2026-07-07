<template>
  <div class="mt-12">
    <div class="flex items-center gap-3 mb-6">
      <span class="w-1 h-5 rounded-sm" style="background: var(--color-primary);"></span>
      <h3 class="text-lg font-semibold" style="color: var(--color-text-primary); font-family: var(--font-display);">
        评论区 <span class="text-sm font-normal" style="color: var(--color-text-tertiary);">({{ comments.length }})</span>
      </h3>
    </div>

    <!-- Comment Input -->
    <div class="p-4 mb-6" :style="commentInputStyle">
      <div class="flex items-start gap-3">
        <div class="w-9 h-9 rounded-full flex items-center justify-center shrink-0" style="background: linear-gradient(135deg, var(--color-primary-200), var(--color-secondary-100)); border: 1px solid var(--color-border);">
          <span class="text-sm font-bold" style="color: var(--color-text-inverse);">{{ userStore.playerName.charAt(0).toUpperCase() }}</span>
        </div>
        <div class="flex-1">
          <!-- 错误提示 -->
          <div v-if="commentError" class="mb-2 px-3 py-2 text-xs font-mono rounded" style="color: #ff4444; background: rgba(255,68,68,0.08); border: 1px solid rgba(255,68,68,0.2);">
            {{ commentError }}
          </div>
          <!-- 成功提示 -->
          <div v-if="commentSuccess" class="mb-2 px-3 py-2 text-xs font-mono rounded" style="color: #22c55e; background: rgba(34,197,94,0.08); border: 1px solid rgba(34,197,94,0.2);">
            {{ commentSuccess }}
          </div>
          <textarea
            v-model="newComment"
            class="w-full p-3 text-sm outline-none resize-none transition-all duration-150"
            style="background: var(--color-bg-primary); border: 1px solid var(--color-border); border-radius: var(--radius-sm); color: var(--color-text-primary); font-family: var(--font-body); min-height: 80px;"
            placeholder="写下你的评论..."
          ></textarea>
          <div class="flex justify-end mt-3">
            <button
              class="flex items-center gap-2 px-4 py-2 text-sm font-medium transition-all duration-150"
              :style="{
                background: newComment.trim() ? 'var(--color-primary)' : 'var(--color-bg-tertiary)',
                color: newComment.trim() ? 'white' : 'var(--color-text-tertiary)',
                borderRadius: 'var(--radius-sm)',
                cursor: newComment.trim() ? 'pointer' : 'not-allowed'
              }"
              :disabled="!newComment.trim()"
              @click="handleAddComment"
            >
              <Send class="w-4 h-4" />
              <span>发表评论</span>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Comment List -->
    <div class="space-y-4">
      <div
        v-for="comment in comments"
        :key="comment.id"
        class="p-4 transition-all duration-150 hover:translate-y-[-1px]"
        :style="commentCardStyle"
      >
        <div class="flex items-start gap-3">
          <div class="w-9 h-9 rounded-full flex items-center justify-center shrink-0" style="background: linear-gradient(135deg, var(--color-primary-200), var(--color-secondary-100)); border: 1px solid var(--color-border);">
            <span class="text-sm font-bold" style="color: var(--color-text-inverse);">{{ comment.avatar || comment.author.charAt(0).toUpperCase() }}</span>
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center justify-between mb-2">
              <span class="text-sm font-medium" style="color: var(--color-text-primary);">{{ comment.author }}</span>
              <span class="text-xs" style="color: var(--color-text-tertiary);">{{ formatCommentTime(comment.time) }}</span>
            </div>
            <p class="text-sm mb-3" style="color: var(--color-text-secondary); line-height: 1.6;">{{ comment.content }}</p>
            <div class="flex items-center gap-4">
              <button class="flex items-center gap-1 text-xs transition-colors duration-150" style="color: var(--color-text-tertiary);" @mouseenter="($event.target as HTMLElement).style.color = 'var(--color-primary)'" @mouseleave="($event.target as HTMLElement).style.color = 'var(--color-text-tertiary)'">
                <Heart class="w-3.5 h-3.5" />
                <span>{{ comment.likes }}</span>
              </button>
              <button class="text-xs transition-colors duration-150" style="color: var(--color-text-tertiary);" @mouseenter="($event.target as HTMLElement).style.color = 'var(--color-primary)'" @mouseleave="($event.target as HTMLElement).style.color = 'var(--color-text-tertiary)'">
                回复
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { useUserStore } from '@/stores/userStore'
import type { CommentVO } from '@/api/articles'
import { fetchComments, createComment } from '@/api/articles'
import { Heart, Send, MessageCircle } from 'lucide-vue-next'
import { useCardStyles } from '@/hooks/useCardStyles'

const props = defineProps<{
  articleId: string
}>()

const emit = defineEmits<{
  (e: 'update:commentCount', count: number): void
}>()

const { t } = useI18n()
const userStore = useUserStore()
const { getCardStyle } = useCardStyles()

const comments = ref<CommentVO[]>([])
const newComment = ref('')
const commentError = ref('')
const commentSuccess = ref('')

const commentInputStyle = computed(() => getCardStyle('default', false))
const commentCardStyle = computed(() => getCardStyle('comment', false))

const formatCommentTime = (timestamp: number): string => {
  if (!timestamp) return t('common.justNow')
  const now = Date.now() / 1000
  const diff = now - timestamp
  if (diff < 60) return t('common.justNow')
  if (diff < 3600) return t('common.minutes', { count: Math.floor(diff / 60) })
  if (diff < 86400) return t('common.hours', { count: Math.floor(diff / 3600) })
  if (diff < 2592000) return t('common.days', { count: Math.floor(diff / 86400) })
  return new Date(timestamp * 1000).toLocaleDateString()
}

/** 发表评论 */
const handleAddComment = async () => {
  commentError.value = ''
  commentSuccess.value = ''

  // 内容为空校验
  if (!newComment.value.trim()) return

  // 登录校验
  if (!userStore.isLoggedIn) {
    commentError.value = t('articleDetail.comment.loginRequired')
    return
  }

  try {
    const res = await createComment(props.articleId, newComment.value)
    if (res.code === 0 && res.data) {
      comments.value.unshift(res.data)
      emit('update:commentCount', comments.value.length)
      newComment.value = ''
      commentSuccess.value = t('articleDetail.comment.postSuccess')
      setTimeout(() => { commentSuccess.value = '' }, 3000)
    } else {
      commentError.value = res.message || t('articleDetail.comment.postFail')
    }
  } catch {
    commentError.value = t('common.errors.networkError')
  }
}

onMounted(async () => {
  try {
    const res = await fetchComments(props.articleId)
    if (res.code === 0) {
      comments.value = res.data || []
      emit('update:commentCount', comments.value.length)
    }
  } catch {
    console.warn('评论加载失败，使用空数据')
  }
})
</script>
