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
            <component :is="topic.icon" class="w-3 h-3" :style="iconStyle(topic)" />
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
import { Flame, Bot, Gamepad2, Cpu, Atom, Tv, Smartphone, Sparkles } from 'lucide-vue-next'

const router = useRouter()
const themeStore = useThemeStore()

const handleTopicClick = (topicName: string) => {
  router.push({ path: '/search', query: { q: topicName } })
}

const trendingTopicsData = {
  nexus: [
    { id: 1, name: 'PS6首发', icon: Flame, colorType: 'error' },
    { id: 2, name: 'AI绘画革命', icon: Bot, colorType: 'primary' },
    { id: 3, name: '原神5.0', icon: Gamepad2, colorType: 'warning' },
    { id: 4, name: 'RTX5090评测', icon: Cpu, colorType: 'success' },
    { id: 5, name: '量子芯片突破', icon: Atom, colorType: 'secondary' }
  ],
  comiket: [
    { id: 1, name: '原神5.0', icon: Gamepad2, colorType: 'primary' },
    { id: 2, name: 'RTX 6090', icon: Cpu, colorType: 'secondary' },
    { id: 3, name: '七月新番', icon: Tv, colorType: 'primary' },
    { id: 4, name: '折叠屏', icon: Smartphone, colorType: 'secondary' },
    { id: 5, name: 'AI绘图', icon: Sparkles, colorType: 'primary' },
    { id: 6, name: '游戏本横评', icon: Gamepad2, colorType: 'secondary' }
  ],
  ironcore: [
    { id: 1, name: '量子处理器突破', icon: Atom, colorType: 'primary' },
    { id: 2, name: 'RTX 5090评测', icon: Cpu, colorType: 'secondary' },
    { id: 3, name: '艾尔登法环DLC2', icon: Gamepad2, colorType: 'primary' },
    { id: 4, name: '高达新作剧场版', icon: Tv, colorType: 'secondary' },
    { id: 5, name: '星舰第七次试飞', icon: Atom, colorType: 'primary' }
  ]
}

const trendingTopics = computed(() => trendingTopicsData[themeStore.currentTheme])

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
  const colorMap = {
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