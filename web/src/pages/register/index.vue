<script setup lang="ts">
/**
 * 注册页面（仅注册表单+逻辑，外壳由 AuthLayout 提供）
 */
import { ref, inject } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useUserStore } from '@/stores/userStore'
import RegisterForm from './RegisterForm.vue'
import type { LoginThemeConfig } from '@/components/business/loginThemeConfig'

const { t } = useI18n()
const router = useRouter()
const userStore = useUserStore()

/** 从 AuthLayout 注入主题配置 */
const loginConfig = inject('loginConfig') as LoginThemeConfig
const theme = inject('theme') as string

const registerEmail = ref('')
const registerPassword = ref('')
const confirmPassword = ref('')
const playerName = ref('')
const registerCode = ref('')
const codeSent = ref(false)
const sendCodeError = ref('')
const registerError = ref('')
const isSubmitting = ref(false)

/** 发送邮箱验证码 */
const handleSendCode = async () => {
  sendCodeError.value = ''
  if (!registerEmail.value) {
    sendCodeError.value = t('login.errors.emptyEmail')
    return
  }
  try {
    const { sendEmailCode } = await import('@/api/auth')
    const res = await sendEmailCode(registerEmail.value)
    if (res.code !== 0) {
      sendCodeError.value = res.message || t('login.errors.sendCodeFail')
      return
    }
    codeSent.value = true
  } catch {
    sendCodeError.value = t('common.errors.networkError')
  }
}

/** 处理注册 */
const handleRegister = async () => {
  registerError.value = ''
  sendCodeError.value = ''

  if (!registerEmail.value || !registerPassword.value || !confirmPassword.value || !playerName.value) {
    registerError.value = t('login.errors.fillAllFields')
    return
  }
  if (registerPassword.value !== confirmPassword.value) {
    registerError.value = t('login.errors.passwordMismatch')
    return
  }
  if (registerPassword.value.length < 6) {
    registerError.value = t('login.errors.passwordTooShort')
    return
  }
  if (!registerCode.value) {
    registerError.value = t('login.errors.enterCode')
    return
  }

  isSubmitting.value = true
  try {
    const { register: registerApi, verifyEmailCode } = await import('@/api/auth')

    const verifyRes = await verifyEmailCode(registerEmail.value, registerCode.value)
    if (verifyRes.code !== 0) {
      registerError.value = verifyRes.message || t('login.errors.codeError')
      isSubmitting.value = false
      return
    }

    const res = await registerApi({
      email: registerEmail.value,
      password: registerPassword.value,
      confirmPassword: confirmPassword.value,
      playerName: playerName.value,
    })
    if (res.code !== 0) {
      registerError.value = res.message || t('login.errors.registerFail')
      return
    }
    // 注册成功后跳转到登录页
    router.push('/auth/login')
  } catch {
    registerError.value = t('common.errors.networkError')
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <RegisterForm
    v-model:registerEmail="registerEmail"
    v-model:registerPassword="registerPassword"
    v-model:confirmPassword="confirmPassword"
    v-model:playerName="playerName"
    v-model:emailCode="registerCode"
    :registerError="registerError"
    :isSubmitting="isSubmitting"
    :loginConfig="loginConfig"
    :theme="theme"
    :codeSent="codeSent"
    :sendCodeError="sendCodeError"
    @submit="handleRegister"
    @sendCode="handleSendCode"
  />
</template>
