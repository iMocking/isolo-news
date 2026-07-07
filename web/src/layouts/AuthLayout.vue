<template>
  <div class="min-h-screen flex items-center justify-center px-4 py-8 overflow-hidden" :class="loginConfig.bgClass">
    <!-- Scanline overlay -->
    <div v-if="loginConfig.showScanline" class="absolute inset-0 pointer-events-none" :style="{ backgroundImage: 'repeating-linear-gradient(0deg, transparent, transparent 2px, rgba(255,255,255,0.05) 2px, rgba(255,255,255,0.05) 4px), linear-gradient(rgba(255,255,255,0.03) 1px, transparent 1px), linear-gradient(90deg, rgba(255,255,255,0.03) 1px, transparent 1px)', backgroundSize: '100% 4px, 50px 50px, 50px 50px' }" />
    <!-- Glow corners -->
    <div class="absolute top-0 left-0 w-96 h-96 pointer-events-none" :style="{ background: `radial-gradient(circle, ${loginConfig.glowColor1} 0%, transparent 70%)` }" />
    <div class="absolute bottom-0 right-0 w-96 h-96 pointer-events-none" :style="{ background: `radial-gradient(circle, ${loginConfig.glowColor2} 0%, transparent 70%)` }" />

    <!-- Auth Card -->
    <div class="relative w-full max-w-md mx-auto" :style="cardStyle">
      <!-- Corner brackets -->
      <span class="absolute top-4 left-4 text-lg font-mono" :style="{ color: loginConfig.cornerBracketColor }">[</span>
      <span class="absolute top-4 right-4 text-lg font-mono" :style="{ color: loginConfig.cornerBracketColor }">]</span>
      <span class="absolute bottom-4 left-4 text-lg font-mono" :style="{ color: loginConfig.cornerBracketColor }">[</span>
      <span class="absolute bottom-4 right-4 text-lg font-mono" :style="{ color: loginConfig.cornerBracketColor }">]</span>

      <!-- Status bar -->
      <div class="flex items-center justify-between mb-6">
        <span class="text-xs font-mono tracking-wider" :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478' }">SYS.STATUS: ONLINE</span>
        <span class="text-xs font-mono tracking-wider" :style="{ color: theme === 'comiket' ? '#22c55e' : '#00ff88' }">CONN: SECURE</span>
      </div>

      <!-- Logo -->
      <div class="text-center mb-8">
        <h1 class="text-3xl md:text-4xl font-bold tracking-widest mb-2" :style="{ color: 'var(--color-primary)', textShadow: loginConfig.logoTextShadow }">{{ logoTitle }}</h1>
        <p class="text-xs tracking-[0.2em] font-mono" :style="{ color: theme === 'comiket' ? '#6b6158' : '#8892a8' }">{{ systemText }}</p>
      </div>

      <!-- 子路由内容：Login / Register -->
      <router-view />

      <!-- Divider -->
      <div class="flex items-center gap-3 my-8">
        <div class="flex-1 h-px" :style="{ background: loginConfig.inputBorder }" />
        <span class="text-xs font-mono tracking-wider" :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478' }">{{ $t('login.divider') }}</span>
        <div class="flex-1 h-px" :style="{ background: loginConfig.inputBorder }" />
      </div>

      <!-- Social login buttons -->
      <LoginSocialButtons :loginConfig="loginConfig" :theme="theme" />

      <!-- Back home -->
      <div class="text-center mt-8">
        <router-link to="/" class="inline-flex items-center gap-2 text-xs font-mono transition-colors duration-150 hover:text-[var(--color-primary)]" :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478' }">
          <ArrowLeft class="w-4 h-4" /> {{ $t('login.backHome') }}
        </router-link>
      </div>

      <!-- Bottom HUD -->
      <div class="mt-6 pt-4 flex justify-between" :style="{ borderTop: `1px solid ${loginConfig.inputBorder}` }">
        <span class="text-xs font-mono" :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478', opacity: 0.5 }">v2.4.1</span>
        <span class="text-xs font-mono" :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478', opacity: 0.5 }">{{ logoTitle }} CORP.</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, provide } from 'vue'
import { useThemeStore } from '@/stores/themeStore'
import { useI18n } from 'vue-i18n'
import { useLoginThemeConfig } from '@/components/business/loginThemeConfig'
import LoginSocialButtons from '@/components/business/LoginSocialButtons.vue'
import { ArrowLeft } from 'lucide-vue-next'

const { t } = useI18n()
const themeStore = useThemeStore()
const theme = computed(() => themeStore.currentTheme)
const loginConfig = useLoginThemeConfig(theme)

/** 向下级子路由提供 loginConfig 和 theme，LoginForm/RegisterForm 需要使用 */
provide('loginConfig', loginConfig)
provide('theme', theme)

const logoTitle = computed(() => {
  const titles: Record<string, string> = { nexus: 'NEXUS', comiket: 'COMIKET', ironcore: 'IRONCORE' }
  return titles[theme.value]
})

const systemText = computed(() => t(`login.systemText.${theme.value}`))

const cardStyle = computed(() => ({
  background: loginConfig.value.cardBg,
  border: loginConfig.value.cardBorder,
  borderRadius: theme.value === 'comiket' ? '20px' : theme.value === 'ironcore' ? '8px' : '12px',
  boxShadow: loginConfig.value.cardShadow,
  padding: theme.value === 'comiket' ? '36px' : theme.value === 'ironcore' ? '30px' : '40px'
}))
</script>
