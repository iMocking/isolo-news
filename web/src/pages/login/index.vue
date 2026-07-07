<script setup lang="ts">
/**
 * 登录页面（仅登录表单+逻辑，外壳由 AuthLayout 提供）
 */
import { ref, inject } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useUserStore } from '@/stores/userStore'
import LoginForm from './LoginForm.vue'
import type { LoginThemeConfig } from '@/components/business/loginThemeConfig'

const { t } = useI18n()
const router = useRouter()
const route = useRoute()
const userStore = useUserStore()

/** 从 AuthLayout 注入主题配置 */
const loginConfig = inject('loginConfig') as LoginThemeConfig
const theme = inject('theme') as string

const email = ref('')
const password = ref('')
const rememberMe = ref(false)
const loginError = ref('')
const isSubmitting = ref(false)

/** 处理登录 */
const handleLogin = async () => {
  loginError.value = ''
  if (!email.value || !password.value) {
    loginError.value = t('login.errors.emptyEmailPassword')
    return
  }
  isSubmitting.value = true
  try {
    const res = await userStore.loginByEmail(email.value, password.value, rememberMe.value)
    if (res.code !== 0) {
      loginError.value = res.message || t('login.errors.loginFail')
      return
    }
    const redirect = (route.query.redirect as string) || '/'
    router.push(redirect)
  } catch {
    loginError.value = t('common.errors.networkError')
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <LoginForm
    v-model:email="email"
    v-model:password="password"
    v-model:rememberMe="rememberMe"
    :loginError="loginError"
    :isSubmitting="isSubmitting"
    :loginConfig="loginConfig"
    :theme="theme"
    @submit="handleLogin"
  />
</template>
