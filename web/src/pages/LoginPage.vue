<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { ArrowLeft } from 'lucide-vue-next'
import { useUserStore } from '@/stores/userStore'
import LoginForm from '@/components/business/LoginForm.vue'
import RegisterForm from '@/components/business/RegisterForm.vue'
import LoginSocialButtons from '@/components/business/LoginSocialButtons.vue'
import { useThemeStore } from '@/stores/themeStore'
import { useLoginThemeConfig } from '@/components/business/loginThemeConfig'

const { t } = useI18n()
const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const themeStore = useThemeStore()

const activeTab = ref('login')
const email = ref('')
const password = ref('')
const rememberMe = ref(false)

// 注册表单
const registerEmail = ref('')
const registerPassword = ref('')
const confirmPassword = ref('')
const playerName = ref('')
const registerCode = ref('')
const codeSent = ref(false)
const sendCodeError = ref('')

const loginError = ref('')
const registerError = ref('')
const isSubmitting = ref(false)

/** 处理登录 */
const handleLogin = async () => {
  loginError.value = ''
  if (!email.value || !password.value) {
    loginError.value = '请输入邮箱和密码'
    return
  }
  isSubmitting.value = true
  try {
    const res = await userStore.loginByEmail(email.value, password.value, rememberMe.value)
    if (res.code !== 0) {
      loginError.value = res.message || '登录失败，请检查邮箱和密码'
      return
    }
    const redirect = route.query.redirect as string || '/'
    router.push(redirect)
  } catch {
    loginError.value = '网络错误，请稍后再试'
  } finally {
    isSubmitting.value = false
  }
}

/** 发送邮箱验证码 */
const handleSendCode = async () => {
  sendCodeError.value = ''
  if (!registerEmail.value) {
    sendCodeError.value = '请先输入邮箱'
    return
  }
  try {
    const { sendEmailCode } = await import('@/api/auth')
    const res = await sendEmailCode(registerEmail.value)
    if (res.code !== 0) {
      sendCodeError.value = res.message || '验证码发送失败'
      return
    }
    codeSent.value = true
  } catch {
    sendCodeError.value = '网络错误，请稍后再试'
  }
}

/** 处理注册 */
const handleRegister = async () => {
  registerError.value = ''
  sendCodeError.value = ''

  // 基础校验
  if (!registerEmail.value || !registerPassword.value || !confirmPassword.value || !playerName.value) {
    registerError.value = '请填写所有注册字段'
    return
  }
  if (registerPassword.value !== confirmPassword.value) {
    registerError.value = '两次输入的密码不一致'
    return
  }
  if (registerPassword.value.length < 6) {
    registerError.value = '密码长度不能少于6位'
    return
  }

  // 验证码校验
  if (!registerCode.value) {
    registerError.value = '请先获取邮箱验证码并填写'
    return
  }

  isSubmitting.value = true
  try {
    const { register: registerApi, verifyEmailCode } = await import('@/api/auth')

    // 先验证邮箱验证码
    const verifyRes = await verifyEmailCode(registerEmail.value, registerCode.value)
    if (verifyRes.code !== 0) {
      registerError.value = verifyRes.message || '验证码错误或已过期'
      isSubmitting.value = false
      return
    }

    // 验证通过后提交注册
    const res = await registerApi({
      email: registerEmail.value,
      password: registerPassword.value,
      confirmPassword: confirmPassword.value,
      playerName: playerName.value,
    })
    if (res.code !== 0) {
      registerError.value = res.message || '注册失败'
      return
    }
    // 注册成功后切换到登录页，并填充邮箱
    activeTab.value = 'login'
    email.value = registerEmail.value
    registerError.value = ''
    loginError.value = '注册成功，请使用注册的邮箱登录'
  } catch {
    registerError.value = '网络错误，请稍后再试'
  } finally {
    isSubmitting.value = false
  }
}

const theme = computed(() => themeStore.currentTheme)

const loginConfig = useLoginThemeConfig(theme)

const logoTitle = computed(() => {
  const titles = {
    nexus: 'NEXUS',
    comiket: 'COMIKET',
    ironcore: 'IRONCORE'
  }
  return titles[theme.value]
})

const systemText = computed(() => t(`login.systemText.${theme.value}`))
</script>


<template>
  <div class="min-h-screen flex items-center justify-center px-4 py-8 overflow-hidden" :class="loginConfig.bgClass">
    <!-- Scanline -->
    <div v-if="loginConfig.showScanline" class="absolute inset-0 pointer-events-none" style="background-image: repeating-linear-gradient(0deg, transparent, transparent 2px, rgba(255,255,255,0.05) 2px, rgba(255,255,255,0.05) 4px), linear-gradient(rgba(255,255,255,0.03) 1px, transparent 1px), linear-gradient(90deg, rgba(255,255,255,0.03) 1px, transparent 1px); background-size: 100% 4px, 50px 50px, 50px 50px;" />
    <!-- Glow corners -->
    <div class="absolute top-0 left-0 w-96 h-96 pointer-events-none" :style="{ background: `radial-gradient(circle, ${loginConfig.glowColor1} 0%, transparent 70%)` }" />
    <div class="absolute bottom-0 right-0 w-96 h-96 pointer-events-none" :style="{ background: `radial-gradient(circle, ${loginConfig.glowColor2} 0%, transparent 70%)` }" />

    <!-- Login Card -->
    <div class="relative w-full max-w-md mx-auto" :style="{ background: loginConfig.cardBg, border: loginConfig.cardBorder, borderRadius: theme === 'comiket' ? '20px' : theme === 'ironcore' ? '8px' : '12px', boxShadow: loginConfig.cardShadow, padding: theme === 'comiket' ? '36px' : theme === 'ironcore' ? '30px' : '40px' }">
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

      <!-- Tab switch -->
      <div class="flex mb-8" :style="{ borderBottom: `1px solid ${loginConfig.inputBorder}` }">
        <button v-for="tab in ['login', 'register']" :key="tab"
          class="flex-1 pb-3 text-sm text-center tracking-wide transition-all duration-200 relative"
          :class="tab === 'login' ? 'font-semibold' : ''"
          :style="{ color: activeTab === tab ? 'var(--color-primary)' : (theme === 'comiket' ? '#9e958b' : '#5a6478') }"
          @click="activeTab = tab">
          {{ tab === 'login' ? $t('login.tabs.login') : $t('login.tabs.register') }}
          <span v-if="activeTab === tab" class="absolute bottom-0 left-2 right-2 h-0.5" :style="{ background: 'var(--color-primary)', boxShadow: loginConfig.tabIndicatorShadow, borderRadius: theme === 'ironcore' ? '0' : '4px' }"></span>
        </button>
      </div>

      <!-- Form -->
      <LoginForm v-if="activeTab === 'login'" v-model:email="email" v-model:password="password" v-model:rememberMe="rememberMe" :loginError="loginError" :isSubmitting="isSubmitting" :loginConfig="loginConfig" :theme="theme" @submit="handleLogin" />
      <RegisterForm v-if="activeTab === 'register'" v-model:registerEmail="registerEmail" v-model:registerPassword="registerPassword" v-model:confirmPassword="confirmPassword" v-model:playerName="playerName" v-model:emailCode="registerCode" :registerError="registerError" :isSubmitting="isSubmitting" :loginConfig="loginConfig" :theme="theme" :codeSent="codeSent" :sendCodeError="sendCodeError" @submit="handleRegister" @sendCode="handleSendCode" />

      <!-- Divider -->
      <div class="flex items-center gap-3 my-8">
        <div class="flex-1 h-px" :style="{ background: loginConfig.inputBorder }" />
        <span class="text-xs font-mono tracking-wider" :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478' }">{{ $t('login.divider') }}</span>
        <div class="flex-1 h-px" :style="{ background: loginConfig.inputBorder }" />
      </div>

      <LoginSocialButtons :loginConfig="loginConfig" :theme="theme" />

      <!-- Back home -->
      <div class="text-center mt-8">
        <a href="#" class="inline-flex items-center gap-2 text-xs font-mono transition-colors duration-150 hover:text-[var(--color-primary)]" :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478' }"><ArrowLeft class="w-4 h-4" /> {{ $t('login.backHome') }}</a>
      </div>

      <!-- Bottom HUD -->
      <div class="mt-6 pt-4 flex justify-between" :style="{ borderTop: `1px solid ${loginConfig.inputBorder}` }">
        <span class="text-xs font-mono" :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478', opacity: 0.5 }">v2.4.1</span>
        <span class="text-xs font-mono" :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478', opacity: 0.5 }">{{ logoTitle }} CORP.</span>
      </div>
    </div>
  </div>
</template>
