<script setup lang="ts">
import { computed, ref } from 'vue'
import { cn } from '@/lib/utils'
import { useCardStyles } from '@/hooks/useCardStyles'

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

const { getCardStyle } = useCardStyles()
const isHovered = ref(false)

const cardStyle = computed(() => {
  return getCardStyle(props.variant, isHovered.value)
})

const borderClasses = props.borderColor ? {
  primary: 'border-l-[3px] border-l-[var(--color-primary)]',
  secondary: 'border-l-[3px] border-l-[var(--color-secondary)]',
  success: 'border-l-[3px] border-l-[var(--state-success)]',
  warning: 'border-l-[3px] border-l-[var(--state-warning)]',
  error: 'border-l-[3px] border-l-[var(--state-error)]',
  info: 'border-l-[3px] border-l-[var(--state-info)]'
}[props.borderColor] : ''

const hoverClasses = props.hoverable ? 'hover:-translate-y-px transition-all duration-150' : ''

const clickableClasses = props.clickable ? 'cursor-pointer' : ''

const paddingClasses = {
  none: '',
  sm: 'p-3',
  md: 'p-5',
  lg: 'p-6'
}

const cardClasses = cn(
  'overflow-hidden transition-all duration-200',
  borderClasses,
  hoverClasses,
  clickableClasses,
  paddingClasses[props.padding]
)

const handleClick = (event: MouseEvent) => {
  if (props.clickable) {
    emit('click', event)
  }
}
</script>

<template>
  <div
    :class="cardClasses"
    :style="cardStyle"
    @mouseenter="isHovered = true"
    @mouseleave="isHovered = false"
    @click="handleClick"
  >
    <slot />
  </div>
</template>
