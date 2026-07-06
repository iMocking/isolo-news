<script setup lang="ts">
import { computed } from 'vue'
import { cn } from '@/lib/utils'

interface Props {
  variant?: 'primary' | 'secondary' | 'outline' | 'ghost'
  size?: 'sm' | 'md' | 'lg'
  disabled?: boolean
  loading?: boolean
  fullWidth?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'primary',
  size: 'md',
  disabled: false,
  loading: false,
  fullWidth: false
})

const emit = defineEmits<{
  click: [event: MouseEvent]
}>()

const buttonClasses = computed(() => {
  const baseClasses = 'font-medium transition-all duration-150 inline-flex items-center justify-center gap-2 focus:outline-none focus:ring-2 focus:ring-offset-2'

  const variantClasses = {
    primary: 'bg-[var(--color-primary)] text-[var(--color-text-inverse)] hover:opacity-80 shadow-[0_0_20px_rgba(0,240,255,0.2)] hover:shadow-[0_0_30px_rgba(0,240,255,0.35)] hover:-translate-y-px',
    secondary: 'bg-[var(--color-secondary)] text-[var(--color-text-inverse)] hover:opacity-80 shadow-[0_0_20px_rgba(255,45,149,0.2)] hover:shadow-[0_0_30px_rgba(255,45,149,0.35)] hover:-translate-y-px',
    outline: 'bg-transparent border border-[var(--color-border-subtle)] text-[var(--color-text-secondary)] hover:border-[var(--color-primary)] hover:text-[var(--color-primary)]',
    ghost: 'bg-transparent text-[var(--color-text-secondary)] hover:text-[var(--color-primary)] hover:bg-[var(--color-primary-50)]'
  }

  const sizeClasses = {
    sm: 'px-3 py-1.5 text-xs font-mono rounded-[var(--radius-sm)]',
    md: 'px-4 py-2.5 text-sm font-mono rounded-[var(--radius-md)]',
    lg: 'px-6 py-3 text-base font-display tracking-widest rounded-[var(--radius-md)]'
  }

  const disabledClasses = props.disabled ? 'opacity-50 cursor-not-allowed pointer-events-none' : 'cursor-pointer'

  const widthClasses = props.fullWidth ? 'w-full' : ''

  return cn(
    baseClasses,
    variantClasses[props.variant],
    sizeClasses[props.size],
    disabledClasses,
    widthClasses
  )
})

const handleClick = (event: MouseEvent) => {
  if (!props.disabled && !props.loading) {
    emit('click', event)
  }
}
</script>

<template>
  <button
    :class="buttonClasses"
    :disabled="disabled || loading"
    @click="handleClick"
  >
    <svg
      v-if="loading"
      class="animate-spin h-4 w-4"
      fill="none"
      viewBox="0 0 24 24"
    >
      <circle
        class="opacity-25"
        cx="12"
        cy="12"
        r="10"
        stroke="currentColor"
        stroke-width="4"
      />
      <path
        class="opacity-75"
        fill="currentColor"
        d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
      />
    </svg>
    <slot />
  </button>
</template>