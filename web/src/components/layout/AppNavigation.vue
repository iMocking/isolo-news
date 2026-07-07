<template>
  <nav
    class="fixed top-0 left-0 right-0 z-50"
    :style="navStyle"
  >
    <div class="max-w-7xl mx-auto px-6 h-16 flex items-center justify-between">
      <!-- Logo -->
      <router-link to="/" class="flex items-center gap-2 shrink-0" :style="logoStyle">
        <span
          class="text-xl font-bold tracking-wider"
          :style="brandTextStyle"
        >
          {{ currentBrandName }}
        </span>
      </router-link>

      <!-- Nav Links -->
      <div class="hidden md:flex items-center gap-1">
        <router-link
          v-for="link in navLinks"
          :key="link.path"
          :to="link.path"
          class="px-4 py-2 text-sm font-medium rounded-md whitespace-nowrap transition-all duration-150"
          :style="getLinkStyle(link.path)"
        >
          {{ link.label }}
        </router-link>
      </div>

      <!-- Right Actions -->
      <div class="flex items-center gap-4">
        <!-- Language Switcher -->
        <LocaleSwitcher />

        <!-- Theme Switcher -->
        <ThemeSwitcher />

        <!-- User Avatar + Level -->
        <div
          class="relative flex items-center gap-2 cursor-pointer"
          @click="handleAvatarClick"
          @mouseenter="handleMouseEnter"
          @mouseleave="handleMouseLeave"
        >
          <div class="relative">
            <div
              class="w-8 h-8 rounded-full overflow-hidden transition-transform duration-150 hover:scale-110"
              :style="avatarStyle"
            >
              <img
                v-if="isLoggedIn"
                :src="userAvatar"
                :alt="userStore.playerName"
                class="w-full h-full object-cover"
              />
              <span v-else class="text-xs font-bold w-full h-full flex items-center justify-center" :style="{ color: 'var(--color-textPrimary)' }">
                {{ avatarText }}
              </span>
            </div>
            <span
              v-if="isLoggedIn"
              class="absolute -top-1 -right-1 text-[9px] font-bold px-1 rounded"
              :style="levelBadgeStyle"
            >
              {{ userLevel }}
            </span>
          </div>

          <!-- Hover Menu -->
          <div
            v-if="isLoggedIn && showMenu"
            class="absolute top-14 right-0 w-40 z-50"
            @mouseenter="handleMenuMouseEnter"
            @mouseleave="handleMenuMouseLeave"
          >
            <div
              class="rounded-lg overflow-hidden shadow-xl transition-all duration-150"
              :style="{
                background: 'var(--color-bgCard)',
                border: '1px solid var(--color-border)',
                backdropFilter: 'blur(12px)'
              }"
            >
              <div class="py-2">
                <button
                  class="w-full flex items-center gap-3 px-4 py-2.5 text-sm transition-colors duration-100"
                  :style="{ color: 'var(--color-textSecondary)' }"
                  @mouseenter="(e: any) => e.currentTarget.style.color = 'var(--color-primary)'"
                  @mouseleave="(e: any) => e.currentTarget.style.color = 'var(--color-textSecondary)'"
                  @click.stop="navigateToProfile"
                >
                  <User class="w-4 h-4" />
                  <span>{{ t('nav.profile') }}</span>
                </button>
                <button
                  class="w-full flex items-center gap-3 px-4 py-2.5 text-sm transition-colors duration-100"
                  :style="{ color: 'var(--color-textSecondary)' }"
                  @mouseenter="(e: any) => e.currentTarget.style.color = 'var(--color-primary)'"
                  @mouseleave="(e: any) => e.currentTarget.style.color = 'var(--color-textSecondary)'"
                  @click.stop="navigateToSettings"
                >
                  <Settings class="w-4 h-4" />
                  <span>{{ t('nav.settings') }}</span>
                </button>
                <div class="border-t my-1" :style="{ borderColor: 'var(--color-borderSubtle)' }" />
                <button
                  class="w-full flex items-center gap-3 px-4 py-2.5 text-sm transition-colors duration-100"
                  :style="{ color: 'var(--stateError)' }"
                  @mouseenter="logoutMouseEnter"
                  @mouseleave="logoutMouseLeave"
                  @click.stop="handleLogout"
                >
                  <LogOut class="w-4 h-4" />
                  <span>{{ t('nav.logout') }}</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useThemeStore } from '@/stores/themeStore'
import { useUserStore } from '@/stores/userStore'
import ThemeSwitcher from './ThemeSwitcher.vue'
import LocaleSwitcher from './LocaleSwitcher.vue'
import { User, Settings, LogOut } from 'lucide-vue-next'

const router = useRouter()
const { t } = useI18n()
const themeStore = useThemeStore()
const userStore = useUserStore()

const showMenu = ref(false)
let menuCloseTimer: ReturnType<typeof setTimeout> | null = null

const handleMouseEnter = () => {
  if (menuCloseTimer) {
    clearTimeout(menuCloseTimer)
    menuCloseTimer = null
  }
  showMenu.value = true
}

const handleMouseLeave = () => {
  menuCloseTimer = setTimeout(() => {
    showMenu.value = false
    menuCloseTimer = null
  }, 200)
}

const handleMenuMouseEnter = () => {
  if (menuCloseTimer) {
    clearTimeout(menuCloseTimer)
    menuCloseTimer = null
  }
}

const handleMenuMouseLeave = () => {
  showMenu.value = false
}

const navigateToProfile = () => {
  showMenu.value = false
  router.push('/profile')
}

const navigateToSettings = () => {
  showMenu.value = false
  router.push('/settings')
}

const handleLogout = () => {
  showMenu.value = false
  userStore.logout()
  router.push('/')
}

const logoutMouseEnter = (e: any) => {
  e.currentTarget.style.color = 'var(--stateError)'
  e.currentTarget.style.background = 'var(--color-primary50)'
}

const logoutMouseLeave = (e: any) => {
  e.currentTarget.style.color = 'var(--stateError)'
  e.currentTarget.style.background = 'transparent'
}

const handleAvatarClick = () => {
  if (!userStore.isLoggedIn) {
    router.push({ path: '/login', query: { redirect: router.currentRoute.value.fullPath } })
  }
}

const navLinks = computed(() => [
  { path: '/', label: t('nav.intelligenceCenter') },
  { path: '/article/list', label: t('nav.intelligenceWarehouse') },
  { path: '/about', label: t('nav.aboutMe') }
])

const currentBrandName = computed(() => {
  const themeNames = {
    nexus: 'NEXUS',
    comiket: 'COMIKET',
    ironcore: 'IRONCORE'
  }
  return themeNames[themeStore.currentTheme]
})

const currentRoutePath = computed(() => router.currentRoute.value.path)
const isLoggedIn = computed(() => userStore.isLoggedIn)
const userLevel = computed(() => userStore.level)
const avatarText = computed(() => {
  const texts = {
    nexus: 'NX',
    comiket: 'CO',
    ironcore: 'IR'
  }
  return texts[themeStore.currentTheme]
})

const defaultAvatar = computed(() => {
  const avatars: Record<string, string> = {
    nexus: '/images/image_1_yi19x4.jpg',
    comiket: '/images/image_3_yi19x4.jpg',
    ironcore: '/images/image_5_yi19x4.jpg'
  }
  return avatars[themeStore.currentTheme] || avatars.nexus
})

const userAvatar = computed(() => {
  return userStore.currentUser?.avatarUrl || defaultAvatar.value
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
  border: `1.5px solid var(--color-primary)`,
  color: `var(--color-textPrimary)`
}))

const levelBadgeStyle = computed(() => ({
  fontFamily: `var(--font-mono)`,
  background: `var(--color-secondary)`,
  color: 'white',
  lineHeight: '1.3'
}))

const getLinkStyle = (path: string) => {
  const routePath = currentRoutePath.value
  const isActive =
    routePath === path ||
    (path === '/article/list' && (routePath.startsWith('/article/list') || routePath.startsWith('/article/detail')))

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