<script setup lang="ts">
import { computed } from 'vue'
import { cn } from '@/lib/utils'

interface Props {
  value?: number
  max?: number
  variant?: 'primary' | 'secondary' | 'success' | 'warning' | 'error' | 'gradient'
  size?: 'sm' | 'md' | 'lg'
  showValue?: boolean
  animated?: boolean
  striped?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  value: 0,
  max: 100,
  variant: 'primary',
  size: 'md',
  showValue: false,
  animated: false,
  striped: false
})

const percentage = computed(() => {
  const percent = (props.value / props.max) * 100
  return Math.min(Math.max(percent, 0), 100)
})

const progressClasses = computed(() => {
  const baseClasses = 'w-full overflow-hidden transition-all duration-300'

  const sizeClasses = {
    sm: 'h-1',
    md: 'h-2',
    lg: 'h-3'
  }

  return cn(
    baseClasses,
    sizeClasses[props.size]
  )
})

const barClasses = computed(() => {
  const baseClasses = 'h-full transition-all duration-300'

  const variantClasses = {
    primary: 'bg-[var(--color-primary)]',
    secondary: 'bg-[var(--color-secondary)]',
    success: 'bg-[var(--state-success)]',
    warning: 'bg-[var(--state-warning)]',
    error: 'bg-[var(--state-error)]',
    gradient: 'bg-gradient-to-r from-[var(--color-primary)] to-[var(--color-secondary)]'
  }

  const animatedClasses = props.animated ? 'animate-pulse' : ''
  const stripedClasses = props.striped ? 'bg-[length:1rem_1rem]' : ''

  return cn(
    baseClasses,
    variantClasses[props.variant],
    animatedClasses,
    stripedClasses
  )
})
</script>

<template>
  <div class="flex flex-col gap-2">
    <div
      v-if="showValue"
      class="flex items-center justify-between text-xs font-mono"
      style="color: var(--color-text-secondary);"
    >
      <span>{{ percentage.toFixed(0) }}%</span>
      <span>{{ value }} / {{ max }}</span>
    </div>

    <div
      :class="progressClasses"
      class="rounded-[var(--radius-full)]"
      style="background: var(--color-bg-tertiary); border: 1px solid var(--color-border);"
    >
      <div
        :class="barClasses"
        :style="{ width: `${percentage}%` }"
        class="rounded-[var(--radius-full)]"
      >
        <slot name="bar" />
      </div>
    </div>
  </div>
</template>