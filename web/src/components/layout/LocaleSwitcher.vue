<template>
  <button
    class="flex items-center gap-1 px-2.5 py-1.5 text-xs font-semibold rounded-md whitespace-nowrap transition-all duration-150"
    :style="buttonStyle"
    @click="handleToggle"
    :title="toggleTitle"
  >
    <Globe class="w-3.5 h-3.5" />
    <span>{{ currentLabel }}</span>
  </button>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { setLocale, type SupportedLocale } from '@/i18n'
import { Globe } from 'lucide-vue-next'

const { locale } = useI18n()

const currentLabel = computed(() => {
  return locale.value === 'zh-CN' ? '中' : 'EN'
})

const toggleTitle = computed(() => {
  return locale.value === 'zh-CN' ? '切换至英文' : 'Switch to Chinese'
})

const buttonStyle = computed(() => {
  const isCN = locale.value === 'zh-CN'
  return {
    background: isCN ? 'var(--color-primary50)' : 'transparent',
    border: `1px solid ${isCN ? 'var(--color-primary)' : 'var(--color-borderSubtle)'}`,
    color: isCN ? 'var(--color-primary)' : 'var(--color-textTertiary)',
    fontFamily: 'var(--font-mono)',
    cursor: 'pointer'
  }
})

const handleToggle = () => {
  const newLocale: SupportedLocale = locale.value === 'zh-CN' ? 'en-US' : 'zh-CN'
  setLocale(newLocale)
}
</script>
