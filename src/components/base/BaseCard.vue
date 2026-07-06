<script setup lang="ts">
import { computed } from 'vue'
import { cn } from '@/lib/utils'

interface Props {
  variant?: 'default' | 'elevated' | 'outlined'
  borderColor?: 'primary' | 'secondary' | 'success' | 'warning' | 'error' | 'info'
  hoverable?: boolean
  clickable?: boolean
  padding?: 'none' | 'sm' | 'md' | 'lg'
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'default',
  borderColor: undefined,
  hoverable: false,
  clickable: false,
  padding: 'md'
})

const emit = defineEmits<{
  click: [event: MouseEvent]
}>()

const cardClasses = computed(() => {
  const baseClasses = 'rounded-[var(--radius-md)] overflow-hidden'

  const variantClasses = {
    default: 'bg-[var(--color-bg-card)]',
    elevated: 'bg-[var(--color-bg-elevated)] shadow-lg',
    outlined: 'bg-[var(--color-bg-card)] border border-[var(--color-border-subtle)]'
  }

  const borderClasses = props.borderColor ? {
    primary: 'border-l-[3px] border-l-[var(--color-primary)]',
    secondary: 'border-l-[3px] border-l-[var(--color-secondary)]',
    success: 'border-l-[3px] border-l-[var(--state-success)]',
    warning: 'border-l-[3px] border-l-[var(--state-warning)]',
    error: 'border-l-[3px] border-l-[var(--state-error)]',
    info: 'border-l-[3px] border-l-[var(--state-info)]'
  }[props.borderColor] : ''

  const hoverClasses = props.hoverable ? 'hover:-translate-y-px transition-all duration-150' : ''

  const clickableClasses = props.clickable ? 'cursor-pointer hover:shadow-[0_0_20px_rgba(0,240,255,0.08)]' : ''

  const paddingClasses = {
    none: '',
    sm: 'p-3',
    md: 'p-5',
    lg: 'p-6'
  }

  const borderVariantClasses = props.variant === 'outlined' ? '' : 'border border-[var(--color-border-subtle)]'

  return cn(
    baseClasses,
    variantClasses[props.variant],
    borderVariantClasses,
    borderClasses,
    hoverClasses,
    clickableClasses,
    paddingClasses[props.padding]
  )
})

const handleClick = (event: MouseEvent) => {
  if (props.clickable) {
    emit('click', event)
  }
}
</script>

<template>
  <div
    :class="cardClasses"
    @click="handleClick"
  >
    <slot />
  </div>
</template>