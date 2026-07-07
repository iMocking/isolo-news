<template>
  <section class="relative overflow-hidden">
    <!-- 背景渐变层 -->
    <div class="absolute inset-0" :style="heroBgStyle"></div>
    <div class="absolute inset-0" :style="heroOverlayStyle"></div>
    <div class="relative z-10 max-w-5xl mx-auto px-6 py-12 md:py-16">
      <div class="flex flex-col md:flex-row items-center md:items-end gap-6 md:gap-8">
        <!-- 头像区块 -->
        <div class="relative shrink-0">
          <div class="w-28 h-28 md:w-36 md:h-36 overflow-hidden" :style="avatarStyle">
            <img :src="userAvatar" :alt="user.username" class="w-full h-full object-cover" />
            <!-- HUD 角标（nexus / ironcore 主题） -->
            <template v-if="theme !== 'comiket'">
              <span class="absolute top-1.5 left-1.5 text-lg font-bold leading-none" :style="avatarCornerStyle">[</span>
              <span class="absolute top-1.5 right-1.5 text-lg font-bold leading-none" :style="avatarCornerStyle">]</span>
              <span class="absolute bottom-1.5 left-1.5 text-lg font-bold leading-none" :style="{ ...avatarCornerStyle, color: 'var(--color-secondary)' }">[</span>
              <span class="absolute bottom-1.5 right-1.5 text-lg font-bold leading-none" :style="{ ...avatarCornerStyle, color: 'var(--color-secondary)' }">]</span>
            </template>
            <!-- 底部渐变遮罩 -->
            <div class="absolute bottom-0 left-0 right-0 h-1/3 pointer-events-none"
                 :style="{ background: 'linear-gradient(to top, rgba(0,0,0,0.3), transparent)' }"></div>
          </div>
          <!-- 等级徽章 -->
          <div class="absolute -bottom-2 -right-2 px-2.5 py-1 text-xs font-bold" :style="levelBadgeStyle">
            Lv.{{ user.level }}
          </div>
          <!-- 外发光 -->
          <div class="absolute -inset-1 -z-10 opacity-40 blur-lg"
               :style="{ background: 'var(--color-primary)', borderRadius: avatarStyle.borderRadius }"></div>
        </div>

        <!-- 用户信息 -->
        <div class="text-center md:text-left flex-1 min-w-0">
          <h1 class="text-2xl md:text-4xl font-bold tracking-wider mb-2"
              :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }">
            {{ user.username }}
          </h1>
          <p class="text-sm mb-4 tracking-wide"
             :style="{ fontFamily: 'var(--font-mono)', color: 'var(--color-textSecondary)' }">
            {{ user.title }}
          </p>
          <div class="flex flex-wrap justify-center md:justify-start gap-2">
            <span v-for="tag in user.tags" :key="tag"
                  class="text-xs px-3 py-1 rounded-full" :style="tagStyle">
              {{ tag }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useThemeStore } from '@/stores/themeStore'
import { useUserStore } from '@/stores/userStore'

const themeStore = useThemeStore()
const userStore = useUserStore()

const theme = computed(() => themeStore.currentTheme)

const user = computed(() => ({
  username: userStore.playerName,
  title: userStore.title,
  level: userStore.level,
  tags: userStore.tags
}))

const defaultAvatar = computed(() => {
  const avatars: Record<string, string> = {
    nexus: '/images/image_1_yi19x4.jpg',
    comiket: '/images/image_3_yi19x4.jpg',
    ironcore: '/images/image_5_yi19x4.jpg'
  }
  return avatars[theme.value] || avatars.nexus
})

const userAvatar = computed(() => {
  return userStore.currentUser?.avatarUrl || defaultAvatar.value
})

const heroBgStyle = computed(() => {
  const gradients: Record<string, string> = {
    nexus: 'linear-gradient(180deg, rgba(0,240,255,0.08) 0%, rgba(0,240,255,0.02) 50%, transparent 100%)',
    comiket: 'linear-gradient(180deg, rgba(255,107,43,0.08) 0%, rgba(255,107,43,0.02) 50%, transparent 100%)',
    ironcore: 'linear-gradient(180deg, rgba(240,160,48,0.08) 0%, rgba(240,160,48,0.02) 50%, transparent 100%)'
  }
  return { background: gradients[theme.value] || gradients.nexus }
})

const heroOverlayStyle = computed(() => {
  const colors: Record<string, string> = {
    nexus: 'radial-gradient(ellipse at 20% 0%, rgba(0,240,255,0.12) 0%, transparent 50%), radial-gradient(ellipse at 80% 100%, rgba(255,45,149,0.08) 0%, transparent 50%)',
    comiket: 'radial-gradient(ellipse at 20% 0%, rgba(255,107,43,0.12) 0%, transparent 50%), radial-gradient(ellipse at 80% 100%, rgba(59,158,255,0.08) 0%, transparent 50%)',
    ironcore: 'radial-gradient(ellipse at 20% 0%, rgba(240,160,48,0.12) 0%, transparent 50%), radial-gradient(ellipse at 80% 100%, rgba(78,205,196,0.08) 0%, transparent 50%)'
  }
  return { background: colors[theme.value] || colors.nexus }
})

const avatarStyle = computed(() => {
  const styles: Record<string, Record<string, string>> = {
    nexus: { border: '2px solid var(--color-primary)', boxShadow: '0 0 24px var(--color-primary200), 0 0 48px var(--color-primary50)', borderRadius: '12px' },
    comiket: { border: '2px solid var(--color-primary)', boxShadow: '0 4px 20px rgba(255,107,43,0.25), 0 0 0 4px rgba(255,107,43,0.1)', borderRadius: '20px' },
    ironcore: { border: '2px solid var(--color-primary)', boxShadow: '0 0 20px var(--color-primary200)', borderRadius: '6px' }
  }
  return styles[theme.value] || styles.nexus
})

const avatarCornerStyle = computed(() => ({
  color: 'var(--color-primary)',
  opacity: theme.value === 'nexus' ? '0.6' : theme.value === 'ironcore' ? '0.5' : '0'
}))

const levelBadgeStyle = computed(() => {
  const styles: Record<string, Record<string, string>> = {
    nexus: { background: 'linear-gradient(135deg, var(--color-secondary), var(--color-secondaryDark))', color: 'var(--color-textInverse)', fontFamily: 'var(--font-display)', border: '1px solid var(--color-secondaryLight)', boxShadow: '0 0 12px var(--color-secondary200)', borderRadius: '4px' },
    comiket: { background: 'var(--color-secondary)', color: 'white', fontFamily: 'var(--font-display)', borderRadius: '999px', padding: '2px 8px' },
    ironcore: { background: 'var(--color-secondary)', color: 'var(--color-textInverse)', fontFamily: 'var(--font-display)', border: '1px solid var(--color-secondaryDark)', borderRadius: '2px' }
  }
  return styles[theme.value] || styles.nexus
})

const tagStyle = computed(() => ({
  background: 'var(--color-primary50)', color: 'var(--color-primary)',
  border: '1px solid var(--color-border)', fontFamily: 'var(--font-mono)'
}))
</script>
