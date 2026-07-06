<template>
  <nav
    class="fixed top-0 left-0 right-0 z-50"
    :style="navStyle"
  >
    <div class="max-w-7xl mx-auto px-6 h-16 flex items-center justify-between">
      <!-- Logo -->
      <a href="/" class="flex items-center gap-2 shrink-0" :style="logoStyle">
        <span
          class="text-xl font-bold tracking-wider"
          :style="brandTextStyle"
        >
          {{ currentBrandName }}
        </span>
      </a>

      <!-- Nav Links -->
      <div class="hidden md:flex items-center gap-1">
        <a
          v-for="link in navLinks"
          :key="link.path"
          :href="link.path"
          class="px-4 py-2 text-sm font-medium rounded-md whitespace-nowrap transition-all duration-150"
          :style="getLinkStyle(link.path)"
        >
          {{ link.label }}
        </a>
      </div>

      <!-- Right Actions -->
      <div class="flex items-center gap-4">
        <!-- Theme Switcher -->
        <ThemeSwitcher />

        <!-- User Avatar + Level -->
        <div class="flex items-center gap-2 cursor-pointer" @click="handleAvatarClick">
          <div class="relative">
            <div
              class="w-8 h-8 rounded-full flex items-center justify-center text-xs font-bold transition-transform duration-150 hover:scale-110"
              :style="avatarStyle"
            >
              {{ avatarText }}
            </div>
            <span
              class="absolute -top-1 -right-1 text-[9px] font-bold px-1 rounded"
              :style="levelBadgeStyle"
            >
              {{ userLevel }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useThemeStore } from '@/stores/themeStore'
import { useUserStore } from '@/stores/userStore'
import ThemeSwitcher from './ThemeSwitcher.vue'

const router = useRouter()
const themeStore = useThemeStore()
const userStore = useUserStore()

const handleAvatarClick = () => {
  if (userStore.isLoggedIn) {
    router.push('/profile')
  } else {
    router.push({ path: '/login', query: { redirect: router.currentRoute.value.fullPath } })
  }
}

const navLinks = [
  { path: '/', label: '首页' },
  { path: '/list', label: '科技资讯' },
  { path: '/list', label: '游戏情报' },
  { path: '/list', label: '硬件评测' },
  { path: '/list', label: '二次元' }
]

const currentBrandName = computed(() => {
  const themeNames = {
    nexus: 'NEXUS',
    comiket: 'COMIKET',
    ironcore: 'IRONCORE'
  }
  return themeNames[themeStore.currentTheme]
})

const currentRoutePath = computed(() => router.currentRoute.value.path)

const userLevel = computed(() => userStore.level)
const avatarText = computed(() => {
  const texts = {
    nexus: 'NX',
    comiket: 'CO',
    ironcore: 'IR'
  }
  return texts[themeStore.currentTheme]
})

const navStyle = computed(() => {
  const bgColors = {
    nexus: 'rgba(10, 14, 26, 0.92)',
    comiket: 'rgba(250, 248, 245, 0.95)',
    ironcore: 'rgba(18, 20, 24, 0.95)'
  }
  return {
    background: bgColors[themeStore.currentTheme],
    backdropFilter: 'blur(12px)',
    borderBottom: `1px solid var(--color-borderSubtle)`
  }
})

const logoStyle = computed(() => ({
  fontFamily: `var(--font-display)`
}))

const brandTextStyle = computed(() => {
  const shadows = {
    nexus: '0 0 20px rgba(0, 240, 255, 0.4)',
    comiket: '0 0 20px rgba(255, 107, 43, 0.3)',
    ironcore: '0 0 20px rgba(240, 160, 48, 0.3)'
  }
  return {
    color: `var(--color-primary)`,
    textShadow: shadows[themeStore.currentTheme]
  }
})

const avatarStyle = computed(() => ({
  background: `linear-gradient(135deg, var(--color-primaryDark), var(--color-secondaryDark))`,
  color: `var(--color-textPrimary)`
}))

const levelBadgeStyle = computed(() => ({
  fontFamily: `var(--font-mono)`,
  background: `var(--color-secondary)`,
  color: 'white',
  lineHeight: '1.3'
}))

const getLinkStyle = (path: string) => {
  const isActive = currentRoutePath.value === path || (path === '/list' && currentRoutePath.value.startsWith('/list'))
  
  if (isActive) {
    return {
      color: `var(--color-primary)`,
      background: `var(--color-primary50)`
    }
  } else {
    return {
      color: `var(--color-textSecondary)`
    }
  }
}
</script>