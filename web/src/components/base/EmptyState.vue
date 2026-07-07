<template>
  <div class="flex flex-col items-center justify-center py-16 px-6 text-center">
    <!-- 图标 -->
    <div
      class="w-16 h-16 rounded-full flex items-center justify-center mb-5"
      :style="{ background: iconBgColor, border: `1px solid ${iconBorderColor}` }"
    >
      <component :is="currentIcon" class="w-7 h-7" :style="{ color: iconColor }" />
    </div>

    <!-- 标题 -->
    <h3
      class="text-lg font-semibold mb-2"
      :style="{ color: 'var(--color-textPrimary)', fontFamily: 'var(--font-display)' }"
    >
      {{ resolvedTitle }}
    </h3>

    <!-- 描述 -->
    <p
      class="text-sm max-w-xs"
      :style="{ color: 'var(--color-textTertiary)', fontFamily: 'var(--font-mono)', lineHeight: 1.6 }"
    >
      {{ resolvedDescription }}
    </p>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { WifiOff, Inbox, SearchX } from 'lucide-vue-next'
import type { Component } from 'vue'

type EmptyStateType = 'empty' | 'offline' | 'search-empty'

const props = withDefaults(defineProps<{
  /** 空状态类型，决定默认图标和文案 */
  type?: EmptyStateType
  /** 自定义标题（覆盖 type 的默认值） */
  title?: string
  /** 自定义描述（覆盖 type 的默认值） */
  description?: string
}>(), {
  type: 'empty',
})

const { t } = useI18n()

/** 根据 type 返回对应图标组件 */
const currentIcon = computed<Component>(() => {
  const iconMap: Record<EmptyStateType, Component> = {
    'empty': Inbox,
    'offline': WifiOff,
    'search-empty': SearchX,
  }
  return iconMap[props.type]
})

const iconBgColor = computed(() => {
  if (props.type === 'offline') return 'rgba(255, 77, 77, 0.08)'
  if (props.type === 'search-empty') return 'rgba(255, 170, 0, 0.08)'
  return 'var(--color-bgElevated)'
})

const iconBorderColor = computed(() => {
  if (props.type === 'offline') return 'rgba(255, 77, 77, 0.2)'
  if (props.type === 'search-empty') return 'rgba(255, 170, 0, 0.2)'
  return 'var(--color-borderSubtle)'
})

const iconColor = computed(() => {
  if (props.type === 'offline') return 'rgb(255, 77, 77)'
  if (props.type === 'search-empty') return 'rgb(255, 170, 0)'
  return 'var(--color-textTertiary)'
})

/** 分辨率标题：优先自定义 title，否则从 i18n 取 type 对应文案 */
const resolvedTitle = computed(() => {
  if (props.title) return props.title
  const keyMap: Record<EmptyStateType, string> = {
    'empty': 'common.empty.title',
    'offline': 'common.offline.title',
    'search-empty': 'common.searchEmpty.title',
  }
  return t(keyMap[props.type])
})

const resolvedDescription = computed(() => {
  if (props.description) return props.description
  const keyMap: Record<EmptyStateType, string> = {
    'empty': 'common.empty.description',
    'offline': 'common.offline.description',
    'search-empty': 'common.searchEmpty.description',
  }
  return t(keyMap[props.type])
})
</script>
