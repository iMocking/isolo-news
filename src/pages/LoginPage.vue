<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { Github, Chrome, Twitter, ArrowLeft } from 'lucide-vue-next'
import BaseButton from '@/components/base/BaseButton.vue'
import { useUserStore } from '@/stores/userStore'
import { useThemeStore } from '@/stores/themeStore'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const themeStore = useThemeStore()

const activeTab = ref('login')
const username = ref('')
const password = ref('')
const rememberMe = ref(false)

const handleLogin = () => {
  userStore.login()
  const redirect = route.query.redirect as string || '/'
  router.push(redirect)
}

const theme = computed(() => themeStore.currentTheme)

const loginConfig = computed(() => {
  const configs = {
    nexus: {
      bgClass: 'bg-[#0a0e1a]',
      cardBg: '#151c2e',
      cardBorder: '1px solid rgba(0, 240, 255, 0.15)',
      cardShadow: '0 0 40px rgba(0, 240, 255, 0.15), 0 0 80px rgba(0, 240, 255, 0.05), inset 0 1px 0 rgba(255, 255, 255, 0.05)',
      inputBg: '#111827',
      inputBorder: 'rgba(255, 255, 255, 0.06)',
      inputFocusBorder: '#00f0ff',
      inputFocusShadow: '0 0 0 2px rgba(0, 240, 255, 0.1)',
      scanlineOpacity: 0.03,
      glowColor1: 'rgba(0, 240, 255, 0.08)',
      glowColor2: 'rgba(255, 45, 149, 0.06)',
      cornerBracketColor: 'rgba(0, 240, 255, 0.3)',
      logoTextShadow: '0 0 30px rgba(0, 240, 255, 0.4)',
      tabIndicatorShadow: '0 0 15px rgba(0, 240, 255, 0.5)',
      socialHoverBorder: '#00f0ff',
      socialHoverColor: '#00f0ff',
      showScanline: true
    },
    comiket: {
      bgClass: 'bg-[#faf8f5]',
      cardBg: '#ffffff',
      cardBorder: '1px solid rgba(0, 0, 0, 0.08)',
      cardShadow: '0 20px 60px rgba(255, 107, 43, 0.15), 0 10px 30px rgba(59, 158, 255, 0.1), inset 0 1px 0 rgba(255, 255, 255, 0.8)',
      inputBg: '#f8f6f2',
      inputBorder: 'rgba(0, 0, 0, 0.06)',
      inputFocusBorder: '#ff6b2b',
      inputFocusShadow: '0 0 0 2px rgba(255, 107, 43, 0.1)',
      scanlineOpacity: 0,
      glowColor1: 'rgba(255, 107, 43, 0.06)',
      glowColor2: 'rgba(59, 158, 255, 0.05)',
      cornerBracketColor: 'rgba(255, 107, 43, 0.3)',
      logoTextShadow: '0 0 20px rgba(255, 107, 43, 0.3)',
      tabIndicatorShadow: '0 0 10px rgba(255, 107, 43, 0.4)',
      socialHoverBorder: '#ff6b2b',
      socialHoverColor: '#ff6b2b',
      showScanline: false
    },
    ironcore: {
      bgClass: 'bg-[#121418]',
      cardBg: '#1e2128',
      cardBorder: '1px solid rgba(240, 160, 48, 0.1)',
      cardShadow: '0 0 40px rgba(240, 160, 48, 0.1), 0 0 80px rgba(240, 160, 48, 0.03), inset 0 1px 0 rgba(255, 255, 255, 0.03)',
      inputBg: '#1a1d24',
      inputBorder: 'rgba(255, 255, 255, 0.05)',
      inputFocusBorder: '#f0a030',
      inputFocusShadow: '0 0 0 2px rgba(240, 160, 48, 0.1)',
      scanlineOpacity: 0,
      glowColor1: 'rgba(240, 160, 48, 0.06)',
      glowColor2: 'rgba(78, 205, 196, 0.05)',
      cornerBracketColor: 'rgba(240, 160, 48, 0.3)',
      logoTextShadow: '0 0 20px rgba(240, 160, 48, 0.3)',
      tabIndicatorShadow: '0 0 10px rgba(240, 160, 48, 0.4)',
      socialHoverBorder: '#f0a030',
      socialHoverColor: '#f0a030',
      showScanline: false
    }
  }
  return configs[theme.value]
})

const logoTitle = computed(() => {
  const titles = {
    nexus: 'NEXUS',
    comiket: 'COMIKET',
    ironcore: 'IRONCORE'
  }
  return titles[theme.value]
})

const systemText = computed(() => {
  const texts = {
    nexus: 'SYSTEM LOGIN // 系统登录',
    comiket: '动漫资讯 // COMIC INFO',
    ironcore: 'COMMAND CENTER // 指挥中心'
  }
  return texts[theme.value]
})
</script>

<template>
  <div
    class="min-h-screen flex items-center justify-center px-4 py-8 overflow-hidden"
    :class="loginConfig.bgClass"
  >
    <!-- Scanline grid overlay -->
    <div
      v-if="loginConfig.showScanline"
      class="absolute inset-0 pointer-events-none"
      style="background-image: repeating-linear-gradient(0deg, transparent, transparent 2px, rgba(255,255,255,0.05) 2px, rgba(255,255,255,0.05) 4px), linear-gradient(rgba(255,255,255,0.03) 1px, transparent 1px), linear-gradient(90deg, rgba(255,255,255,0.03) 1px, transparent 1px); background-size: 100% 4px, 50px 50px, 50px 50px;"
    />

    <!-- Ambient glow corners -->
    <div
      class="absolute top-0 left-0 w-96 h-96 pointer-events-none"
      :style="{ background: `radial-gradient(circle, ${loginConfig.glowColor1} 0%, transparent 70%)` }"
    />
    <div
      class="absolute bottom-0 right-0 w-96 h-96 pointer-events-none"
      :style="{ background: `radial-gradient(circle, ${loginConfig.glowColor2} 0%, transparent 70%)` }"
    />

    <!-- Login Card -->
    <div
      class="relative w-full max-w-md mx-auto"
      :style="{
        background: loginConfig.cardBg,
        border: loginConfig.cardBorder,
        borderRadius: theme === 'comiket' ? '20px' : theme === 'ironcore' ? '8px' : '12px',
        boxShadow: loginConfig.cardShadow,
        padding: theme === 'comiket' ? '36px' : theme === 'ironcore' ? '30px' : '40px'
      }"
    >
      <!-- Corner brackets HUD -->
      <span
        class="absolute top-4 left-4 text-lg font-mono"
        :style="{ color: loginConfig.cornerBracketColor }"
      >[</span>
      <span
        class="absolute top-4 right-4 text-lg font-mono"
        :style="{ color: loginConfig.cornerBracketColor }"
      >]</span>
      <span
        class="absolute bottom-4 left-4 text-lg font-mono"
        :style="{ color: loginConfig.cornerBracketColor }"
      >[</span>
      <span
        class="absolute bottom-4 right-4 text-lg font-mono"
        :style="{ color: loginConfig.cornerBracketColor }"
      >]</span>

      <!-- Status bar -->
      <div class="flex items-center justify-between mb-6">
        <span
          class="text-xs font-mono tracking-wider"
          :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478' }"
        >
          SYS.STATUS: ONLINE
        </span>
        <span
          class="text-xs font-mono tracking-wider"
          :style="{ color: theme === 'comiket' ? '#22c55e' : '#00ff88' }"
        >
          CONN: SECURE
        </span>
      </div>

      <!-- Logo -->
      <div class="text-center mb-8">
        <h1
          class="text-3xl md:text-4xl font-bold tracking-widest mb-2"
          :style="{
            fontFamily: theme === 'nexus' ? 'var(--font-display)' : theme === 'comiket' ? 'var(--font-display)' : 'var(--font-display)',
            color: 'var(--color-primary)',
            textShadow: loginConfig.logoTextShadow
          }"
        >
          {{ logoTitle }}
        </h1>
        <p
          class="text-xs tracking-[0.2em] font-mono"
          :style="{ color: theme === 'comiket' ? '#6b6158' : '#8892a8' }"
        >
          {{ systemText }}
        </p>
      </div>

      <!-- Tab switch -->
      <div
        class="flex mb-8"
        :style="{ borderBottom: `1px solid ${loginConfig.inputBorder}` }"
      >
        <button
          class="flex-1 pb-3 text-sm font-semibold text-center tracking-wide transition-all duration-200 relative"
          :style="{
            color: activeTab === 'login' ? 'var(--color-primary)' : (theme === 'comiket' ? '#9e958b' : '#5a6478'),
            fontFamily: theme === 'nexus' ? 'var(--font-display)' : theme === 'comiket' ? 'var(--font-display)' : 'var(--font-display)'
          }"
          @click="activeTab = 'login'"
        >
          {{ theme === 'ironcore' ? 'LOGIN' : '登录' }}
          <span
            v-if="activeTab === 'login'"
            class="absolute bottom-0 left-2 right-2 h-0.5"
            :style="{
              background: 'var(--color-primary)',
              boxShadow: loginConfig.tabIndicatorShadow,
              borderRadius: theme === 'ironcore' ? '0' : '4px'
            }"
          ></span>
        </button>
        <button
          class="flex-1 pb-3 text-sm text-center tracking-wide transition-all duration-200 relative"
          :style="{
            color: activeTab === 'register' ? 'var(--color-primary)' : (theme === 'comiket' ? '#9e958b' : '#5a6478'),
            fontFamily: theme === 'nexus' ? 'var(--font-display)' : theme === 'comiket' ? 'var(--font-display)' : 'var(--font-display)'
          }"
          @click="activeTab = 'register'"
        >
          {{ theme === 'ironcore' ? 'REGISTER' : '注册' }}
          <span
            v-if="activeTab === 'register'"
            class="absolute bottom-0 left-2 right-2 h-0.5"
            :style="{
              background: 'var(--color-primary)',
              boxShadow: loginConfig.tabIndicatorShadow,
              borderRadius: theme === 'ironcore' ? '0' : '4px'
            }"
          ></span>
        </button>
      </div>

      <!-- Form -->
      <form
        class="space-y-5"
        @submit.prevent="handleLogin"
      >
        <!-- Username -->
        <div>
          <label
            class="block text-xs font-mono mb-2 tracking-wider"
            :style="{ color: theme === 'comiket' ? '#6b6158' : '#8892a8' }"
          >
            {{ theme === 'ironcore' ? 'CALLSIGN' : 'USERNAME' }}
          </label>
          <input
            v-model="username"
            type="text"
            :placeholder="theme === 'ironcore' ? 'ENTER CALLSIGN' : '玩家ID'"
            class="w-full px-4 py-3 text-sm outline-none transition-all duration-150"
            :style="{
              background: loginConfig.inputBg,
              border: `1px solid ${loginConfig.inputBorder}`,
              color: theme === 'comiket' ? '#2a2520' : '#e8edf5',
              borderRadius: theme === 'ironcore' ? '4px' : '8px',
              fontFamily: 'var(--font-mono)',
              caretColor: 'var(--color-primary)'
            }"
            @focus="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputFocusBorder; ($event.target as HTMLInputElement).style.boxShadow = loginConfig.inputFocusShadow;"
            @blur="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputBorder; ($event.target as HTMLInputElement).style.boxShadow = 'none';"
          >
        </div>

        <!-- Password -->
        <div>
          <label
            class="block text-xs font-mono mb-2 tracking-wider"
            :style="{ color: theme === 'comiket' ? '#6b6158' : '#8892a8' }"
          >
            {{ theme === 'ironcore' ? 'ACCESS CODE' : 'PASSWORD' }}
          </label>
          <input
            v-model="password"
            type="password"
            :placeholder="theme === 'ironcore' ? 'ENTER ACCESS CODE' : '密码'"
            class="w-full px-4 py-3 text-sm outline-none transition-all duration-150"
            :style="{
              background: loginConfig.inputBg,
              border: `1px solid ${loginConfig.inputBorder}`,
              color: theme === 'comiket' ? '#2a2520' : '#e8edf5',
              borderRadius: theme === 'ironcore' ? '4px' : '8px',
              fontFamily: 'var(--font-mono)',
              caretColor: 'var(--color-primary)'
            }"
            @focus="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputFocusBorder; ($event.target as HTMLInputElement).style.boxShadow = loginConfig.inputFocusShadow;"
            @blur="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputBorder; ($event.target as HTMLInputElement).style.boxShadow = 'none';"
          >
        </div>

        <!-- Remember + Forgot -->
        <div class="flex items-center justify-between">
          <label class="flex items-center gap-2 cursor-pointer">
            <input
              v-model="rememberMe"
              type="checkbox"
              class="w-4 h-4"
              :style="{ accentColor: 'var(--color-primary)' }"
            >
            <span
              class="text-xs"
              :style="{ color: theme === 'comiket' ? '#6b6158' : '#8892a8' }"
            >
              {{ theme === 'ironcore' ? 'REMEMBER CREDENTIALS' : '记住我' }}
            </span>
          </label>
          <a
            href="#"
            class="text-xs transition-colors duration-150 hover:text-[var(--color-primary)]"
            :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478' }"
          >
            {{ theme === 'ironcore' ? 'RESET CODE?' : '忘记密码?' }}
          </a>
        </div>

        <!-- Submit -->
        <BaseButton
          variant="primary"
          size="lg"
          full-width
        >
          {{ theme === 'nexus' ? '接入系统 // LOGIN' : theme === 'comiket' ? '开始阅读 // START' : 'AUTHORIZE // 授权' }}
        </BaseButton>
      </form>

      <!-- Divider -->
      <div class="flex items-center gap-3 my-8">
        <div
          class="flex-1 h-px"
          :style="{ background: loginConfig.inputBorder }"
        />
        <span
          class="text-xs font-mono tracking-wider"
          :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478' }"
        >
          {{ theme === 'ironcore' ? 'QUICK ACCESS' : '快速接入' }}
        </span>
        <div
          class="flex-1 h-px"
          :style="{ background: loginConfig.inputBorder }"
        />
      </div>

      <!-- Social buttons -->
      <div class="flex justify-center gap-4">
        <button
          class="flex items-center justify-center w-12 h-12 transition-all duration-150"
          :style="{
            background: loginConfig.inputBg,
            border: `1px solid ${loginConfig.inputBorder}`,
            borderRadius: theme === 'ironcore' ? '4px' : '10px',
            color: theme === 'comiket' ? '#6b6158' : '#8892a8'
          }"
          @mouseenter="($event.target as HTMLElement).style.borderColor = loginConfig.socialHoverBorder; ($event.target as HTMLElement).style.color = loginConfig.socialHoverColor;"
          @mouseleave="($event.target as HTMLElement).style.borderColor = loginConfig.inputBorder; ($event.target as HTMLElement).style.color = (theme === 'comiket' ? '#6b6158' : '#8892a8');"
        >
          <Github class="w-5 h-5" />
        </button>
        <button
          class="flex items-center justify-center w-12 h-12 transition-all duration-150"
          :style="{
            background: loginConfig.inputBg,
            border: `1px solid ${loginConfig.inputBorder}`,
            borderRadius: theme === 'ironcore' ? '4px' : '10px',
            color: theme === 'comiket' ? '#6b6158' : '#8892a8'
          }"
          @mouseenter="($event.target as HTMLElement).style.borderColor = loginConfig.socialHoverBorder; ($event.target as HTMLElement).style.color = loginConfig.socialHoverColor;"
          @mouseleave="($event.target as HTMLElement).style.borderColor = loginConfig.inputBorder; ($event.target as HTMLElement).style.color = (theme === 'comiket' ? '#6b6158' : '#8892a8');"
        >
          <Chrome class="w-5 h-5" />
        </button>
        <button
          class="flex items-center justify-center w-12 h-12 transition-all duration-150"
          :style="{
            background: loginConfig.inputBg,
            border: `1px solid ${loginConfig.inputBorder}`,
            borderRadius: theme === 'ironcore' ? '4px' : '10px',
            color: theme === 'comiket' ? '#6b6158' : '#8892a8'
          }"
          @mouseenter="($event.target as HTMLElement).style.borderColor = loginConfig.socialHoverBorder; ($event.target as HTMLElement).style.color = loginConfig.socialHoverColor;"
          @mouseleave="($event.target as HTMLElement).style.borderColor = loginConfig.inputBorder; ($event.target as HTMLElement).style.color = (theme === 'comiket' ? '#6b6158' : '#8892a8');"
        >
          <Twitter class="w-5 h-5" />
        </button>
      </div>

      <!-- Back home -->
      <div class="text-center mt-8">
        <a
          href="#"
          class="inline-flex items-center gap-2 text-xs font-mono transition-colors duration-150 hover:text-[var(--color-primary)]"
          :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478' }"
        >
          <ArrowLeft class="w-4 h-4" />
          {{ theme === 'ironcore' ? 'RETURN TO MAIN' : '返回首页' }}
        </a>
      </div>

      <!-- Bottom HUD -->
      <div
        class="mt-6 pt-4 flex justify-between"
        :style="{ borderTop: `1px solid ${loginConfig.inputBorder}` }"
      >
        <span
          class="text-xs font-mono"
          :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478', opacity: 0.5 }"
        >
          v2.4.1
        </span>
        <span
          class="text-xs font-mono"
          :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478', opacity: 0.5 }"
        >
          {{ logoTitle }} CORP.
        </span>
      </div>
    </div>
  </div>
</template>
