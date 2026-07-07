<template>
  <section class="relative" :style="sectionStyle">
    <div class="max-w-7xl mx-auto px-6 py-4">
      <div class="flex items-center gap-3 overflow-x-auto no-scrollbar">
        <span
          class="text-xs tracking-widest uppercase shrink-0"
          :style="labelStyle"
        >
          TRENDING
        </span>
        <div class="flex items-center gap-2 overflow-x-auto no-scrollbar">
          <span
            v-for="topic in trendingTopics"
            :key="topic.id"
            class="shrink-0 inline-flex items-center gap-1.5 px-3 py-1.5 rounded-md text-xs font-medium whitespace-nowrap cursor-pointer transition-all duration-150"
            :style="topicStyle(topic)"
            @mouseenter="handleTopicHover(topic)"
            @mouseleave="handleTopicLeave(topic)"
            @click="handleTopicClick(topic.name)"
          >
            <component :is="getIcon(topic.icon)" class="w-3 h-3" :style="iconStyle(topic)" />
            {{ topic.name }}
          </span>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useThemeStore } from '@/stores/themeStore'
import { useArticleStore } from '@/stores/articleStore'
import { Flame, Bot, Gamepad2, Cpu, Atom, Tv, Smartphone, Sparkles, Zap, Star, TrendingUp } from 'lucide-vue-next'

const router = useRouter()
const themeStore = useThemeStore()
const articleStore = useArticleStore()

const handleTopicClick = (topicName: string) => {
  router.push({ path: '/search', query: { q: topicName } })
}

const theme = computed(() => themeStore.currentTheme)

const iconMap: Record<string, any> = {
  Flame, Bot, Gamepad2, Cpu, Atom, Tv, Smartphone, Sparkles, Zap, Star, TrendingUp
}

function getIcon(iconName: string) {
  return iconMap[iconName] || Flame
}

const trendingTopics = computed(() => articleStore.trendingTopics)

const hoveredTopics = ref<Set<number>>(new Set())

const sectionStyle = computed(() => ({
  background: `var(--color-bgPrimary)`,
  borderBottom: `1px solid var(--color-borderSubtle)`
}))

const labelStyle = computed(() => ({
  fontFamily: `var(--font-mono)`,
  color: `var(--color-textTertiary)`
}))

const topicStyle = computed(() => (topic: any) => {
  const isHovered = hoveredTopics.value.has(topic.id)
  
  return {
    background: `var(--color-bgCard)`,
    border: isHovered ? '1px solid var(--color-primary)' : '1px solid var(--color-border)',
    color: `var(--color-textPrimary)`,
    boxShadow: isHovered ? '0 0 12px rgba(0, 240, 255, 0.15)' : 'none',
    borderRadius: themeStore.currentTheme === 'comiket' ? 'var(--radius-md)' : 'var(--radius-md)'
  }
})

const iconStyle = computed(() => (topic: any) => {
  const colorMap: Record<string, string> = {
    primary: 'var(--color-primary)',
    secondary: 'var(--color-secondary)',
    success: 'var(--state-success)',
    warning: 'var(--state-warning)',
    error: 'var(--state-error)',
    info: 'var(--state-info)'
  }
  
  return {
    color: colorMap[topic.colorType as keyof typeof colorMap] || 'var(--color-primary)'
  }
})

const handleTopicHover = (topic: any) => {
  hoveredTopics.value.add(topic.id)
}

const handleTopicLeave = (topic: any) => {
  hoveredTopics.value.delete(topic.id)
}
</script>

<style scoped>
.no-scrollbar::-webkit-scrollbar {
  display: none;
}
.no-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
</style>