<script setup lang="ts">
import { computed } from 'vue'
import { cn } from '@/lib/utils'

interface Props {
  variant?: 'primary' | 'secondary' | 'success' | 'warning' | 'error' | 'info' | 'default'
  size?: 'sm' | 'md' | 'lg'
  outlined?: boolean
  closable?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'default',
  size: 'md',
  outlined: false,
  closable: false
})

const emit = defineEmits<{
  close: []
}>()

const tagClasses = computed(() => {
  const baseClasses = 'inline-flex items-center gap-1 font-mono whitespace-nowrap transition-all duration-150'

  const variantClasses = {
    primary: props.outlined
      ? 'bg-[var(--color-primary-50)] text-[var(--color-primary)] border border-[var(--color-primary-100)]'
      : 'bg-[var(--color-primary-50)] text-[var(--color-primary)] border border-[var(--color-primary-100)]',
    secondary: props.outlined
      ? 'bg-[var(--color-secondary-50)] text-[var(--color-secondary)] border border-[var(--color-secondary-100)]'
      : 'bg-[var(--color-secondary-50)] text-[var(--color-secondary)] border border-[var(--color-secondary-100)]',
    success: props.outlined
      ? 'bg-transparent text-[var(--state-success)] border border-[var(--state-success)]'
      : 'bg-[rgba(0,255,136,0.08)] text-[var(--state-success)] border border-[rgba(0,255,136,0.15)]',
    warning: props.outlined
      ? 'bg-transparent text-[var(--state-warning)] border border-[var(--state-warning)]'
      : 'bg-[rgba(255,170,0,0.08)] text-[var(--state-warning)] border border-[rgba(255,170,0,0.15)]',
    error: props.outlined
      ? 'bg-transparent text-[var(--state-error)] border border-[var(--state-error)]'
      : 'bg-[rgba(255,51,102,0.08)] text-[var(--state-error)] border border-[rgba(255,51,102,0.15)]',
    info: props.outlined
      ? 'bg-transparent text-[var(--state-info)] border border-[var(--state-info)]'
      : 'bg-[rgba(0,170,255,0.08)] text-[var(--state-info)] border border-[rgba(0,170,255,0.15)]',
    default: 'bg-[var(--color-bg-elevated)] text-[var(--color-text-secondary)] border border-[var(--color-border-subtle)]'
  }

  const sizeClasses = {
    sm: 'px-2 py-0.5 text-xs rounded-[var(--radius-sm)]',
    md: 'px-2.5 py-1 text-xs rounded-[var(--radius-sm)]',
    lg: 'px-3 py-1.5 text-sm rounded-[var(--radius-md)]'
  }

  return cn(
    baseClasses,
    variantClasses[props.variant],
    sizeClasses[props.size]
  )
})

const handleClose = () => {
  emit('close')
}
</script>

<template>
  <span :class="tagClasses">
    <slot />
    <button
      v-if="closable"
      class="hover:opacity-70 transition-opacity cursor-pointer"
      @click.stop="handleClose"
    >
      <svg
        class="w-3 h-3"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M6 18L18 6M6 6l12 12"
        />
      </svg>
    </button>
  </span>
</template>