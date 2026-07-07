<template>
  <div class="flex items-center justify-center gap-2 pt-6 flex-wrap">
    <!-- 每页条数选择 -->
    <div class="flex items-center gap-2 mr-6">
      <span class="text-xs whitespace-nowrap" :style="{ color: 'var(--color-textTertiary)', fontFamily: 'var(--font-mono)' }">{{ $t('articleList.pagination.perPage') }}</span>
      <div class="relative">
        <select v-model.number="pageSizeLocal" @change="onPageSizeChange"
          class="appearance-none px-3 h-8 text-xs outline-none cursor-pointer transition-all duration-150"
          :style="selectStyle"
          @mouseenter="selectHovered = true"
          @mouseleave="selectHovered = false"
          @focus="selectFocused = true"
          @blur="selectFocused = false">
          <option :value="10">10</option>
          <option :value="20">20</option>
          <option :value="50">50</option>
        </select>
        <ChevronDown
          class="absolute right-1.5 top-1/2 -translate-y-1/2 w-3 h-3 pointer-events-none transition-transform duration-150"
          :class="{ 'rotate-180': selectFocused }"
          :style="{ color: 'var(--color-textTertiary)' }"
        />
      </div>
      <span class="text-xs whitespace-nowrap" :style="{ color: 'var(--color-textTertiary)', fontFamily: 'var(--font-mono)' }">{{ $t('articleList.pagination.unit') }}</span>
    </div>

    <!-- PREV -->
    <button :disabled="store.currentPage <= 1" @click="goToPage(store.currentPage - 1)"
      class="px-3 py-2 text-xs font-medium whitespace-nowrap transition-all duration-150"
      :style="{ background: 'var(--color-bgCard)', border: '1px solid var(--color-borderSubtle)', color: store.currentPage <= 1 ? 'var(--color-textTertiary)' : 'var(--color-textSecondary)', borderRadius: 'var(--radius-sm)', fontFamily: 'var(--font-mono)', opacity: store.currentPage <= 1 ? 0.4 : 1, cursor: store.currentPage <= 1 ? 'not-allowed' : 'pointer' }">
      PREV
    </button>

    <!-- 页码按钮 -->
    <button v-for="p in visiblePages" :key="p" @click="goToPage(p)"
      class="w-9 h-9 flex items-center justify-center text-sm font-medium whitespace-nowrap transition-all duration-150"
      :style="p === store.currentPage
        ? { background: 'var(--color-primary)', color: '#ffffff', borderRadius: 'var(--radius-sm)', fontFamily: 'var(--font-display)' }
        : { background: 'var(--color-bgCard)', border: '1px solid var(--color-borderSubtle)', color: 'var(--color-textSecondary)', borderRadius: 'var(--radius-sm)', fontFamily: 'var(--font-mono)' }">
      {{ p }}
    </button>

    <!-- NEXT -->
    <button :disabled="store.currentPage >= store.totalPages" @click="goToPage(store.currentPage + 1)"
      class="px-4 py-2 text-xs font-medium whitespace-nowrap transition-all duration-150"
      :style="{ background: store.currentPage >= store.totalPages ? 'var(--color-bgCard)' : 'var(--color-primary50)', border: '1px solid var(--color-primary)', color: store.currentPage >= store.totalPages ? 'var(--color-textTertiary)' : 'var(--color-primary)', borderRadius: 'var(--radius-sm)', fontFamily: 'var(--font-display)', opacity: store.currentPage >= store.totalPages ? 0.4 : 1, cursor: store.currentPage >= store.totalPages ? 'not-allowed' : 'pointer' }">
      NEXT →
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useArticleStore } from '@/stores/articleStore'
import { ChevronDown } from 'lucide-vue-next'

const store = useArticleStore()
const pageSizeLocal = ref(store.pageSize)

/** select 悬浮状态 */
const selectHovered = ref(false)
/** select 聚焦状态 */
const selectFocused = ref(false)

/** select 动态样式：悬浮/聚焦时边框高亮，匹配主题按钮风格 */
const selectStyle = computed(() => ({
  background: 'var(--color-bgCard)',
  border: `1px solid ${selectFocused.value ? 'var(--color-primary)' : selectHovered.value ? 'var(--color-border)' : 'var(--color-borderSubtle)'}`,
  color: 'var(--color-textSecondary)',
  borderRadius: 'var(--radius-sm)',
  fontFamily: 'var(--font-mono)',
  paddingRight: '1.75rem'
}))

/** 可见页码（当前页前后各 2 个） */
const visiblePages = computed(() => {
  const total = store.totalPages
  const current = store.currentPage
  if (total <= 0) return []
  if (total <= 5) return Array.from({ length: total }, (_, i) => i + 1)

  let start = Math.max(1, current - 2)
  let end = Math.min(total, current + 2)
  if (end - start < 4) {
    if (start === 1) end = Math.min(total, start + 4)
    else start = Math.max(1, end - 4)
  }
  return Array.from({ length: end - start + 1 }, (_, i) => start + i)
})

/** 跳转到指定页码 */
function goToPage(page: number) {
  if (page < 1 || page > store.totalPages) return
  store.currentPage = page
  window.scrollTo({ top: 0, behavior: 'smooth' })
  store.loadArticles({ page })
}

/** 每页条数切换 */
function onPageSizeChange() {
  store.currentPage = 1
  store.pageSize = pageSizeLocal.value
  store.loadArticles({ page: 1, page_size: pageSizeLocal.value })
}
</script>

<style scoped>
/* 下拉框选项适配三主题色调 */
select option {
  background: var(--color-bgCard);
  color: var(--color-textSecondary);
}
</style>
