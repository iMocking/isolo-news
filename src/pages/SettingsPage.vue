<script setup lang="ts">
import { Bell, User, Palette, EyeOff, Shield, ChevronRight } from 'lucide-vue-next'
import { ref } from 'vue'
import BaseCard from '@/components/base/BaseCard.vue'
import BaseButton from '@/components/base/BaseButton.vue'
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
import { useThemeStore } from '@/stores/themeStore'

const themeStore = useThemeStore()

const notifications = ref({
  push: true,
  email: true,
  inApp: false
})

const themes = ['nexus', 'comiket', 'ironcore']

const toggleNotification = (type: 'push' | 'email' | 'inApp') => {
  notifications.value[type] = !notifications.value[type]
}

const selectTheme = (theme: string) => {
  themeStore.setTheme(theme as any)
}
</script>

<template>
  <div class="min-h-screen">
    <AppNavigation />
    
    <!-- Settings Content -->
    <div class="max-w-2xl mx-auto px-6 pt-24 pb-16">
      <!-- Page Title -->
      <div class="flex items-center gap-3 mb-8">
        <span
          class="w-3 h-3 rounded-sm"
          style="background: var(--color-primary); box-shadow: 0 0 8px rgba(0, 240, 255, 0.5);"
        />
        <h1
          class="text-2xl font-bold"
          style="font-family: var(--font-display); color: var(--color-text-primary);"
        >
          SYSTEM CONFIG
        </h1>
        <span
          class="text-sm"
          style="color: var(--color-text-tertiary);"
        >
          // 系统设置
        </span>
      </div>

      <!-- Settings Cards -->
      <div class="space-y-4">
        <!-- Notification Settings -->
        <BaseCard padding="md">
          <div class="flex items-center gap-3 mb-4">
            <div
              class="w-8 h-8 rounded-md flex items-center justify-center"
              style="background: var(--color-primary-50); border: 1px solid var(--color-border);"
            >
              <Bell class="w-4 h-4" style="color: var(--color-primary);" />
            </div>
            <h3
              class="text-sm font-semibold"
              style="color: var(--color-text-primary);"
            >
              通知设置
            </h3>
            <span
              class="text-[10px] ml-auto"
              style="font-family: var(--font-mono); color: var(--color-text-tertiary);"
            >
              NOTIFICATION
            </span>
          </div>
          <div class="space-y-3 ml-11">
            <label class="flex items-center justify-between cursor-pointer">
              <span
                class="text-sm"
                style="color: var(--color-text-secondary);"
              >
                推送通知
              </span>
              <div
                class="w-10 h-5 rounded-full relative cursor-pointer"
                :style="{ background: notifications.push ? 'var(--color-primary)' : 'var(--color-bg-primary)' }"
                @click="toggleNotification('push')"
              >
                <div
                  class="absolute w-4 h-4 rounded-full transition-all duration-200"
                  :style="{
                    background: notifications.push ? 'white' : 'var(--color-text-tertiary)',
                    right: notifications.push ? '2px' : 'auto',
                    left: notifications.push ? 'auto' : '2px',
                    top: '2px'
                  }"
                />
              </div>
            </label>
            <label class="flex items-center justify-between cursor-pointer">
              <span
                class="text-sm"
                style="color: var(--color-text-secondary);"
              >
                邮件通知
              </span>
              <div
                class="w-10 h-5 rounded-full relative cursor-pointer"
                :style="{ background: notifications.email ? 'var(--color-primary)' : 'var(--color-bg-primary)' }"
                @click="toggleNotification('email')"
              >
                <div
                  class="absolute w-4 h-4 rounded-full transition-all duration-200"
                  :style="{
                    background: notifications.email ? 'white' : 'var(--color-text-tertiary)',
                    right: notifications.email ? '2px' : 'auto',
                    left: notifications.email ? 'auto' : '2px',
                    top: '2px'
                  }"
                />
              </div>
            </label>
            <label class="flex items-center justify-between cursor-pointer">
              <span
                class="text-sm"
                style="color: var(--color-text-secondary);"
              >
                应用内通知
              </span>
              <div
                class="w-10 h-5 rounded-full relative cursor-pointer"
                :style="{ background: notifications.inApp ? 'var(--color-primary)' : 'var(--color-bg-primary)' }"
                @click="toggleNotification('inApp')"
              >
                <div
                  class="absolute w-4 h-4 rounded-full transition-all duration-200"
                  :style="{
                    background: notifications.inApp ? 'white' : 'var(--color-text-tertiary)',
                    right: notifications.inApp ? '2px' : 'auto',
                    left: notifications.inApp ? 'auto' : '2px',
                    top: '2px'
                  }"
                />
              </div>
            </label>
          </div>
        </BaseCard>

        <!-- Profile Edit -->
        <BaseCard
          padding="md"
          clickable
        >
          <div class="flex items-center gap-3">
            <div
              class="w-8 h-8 rounded-md flex items-center justify-center"
              style="background: var(--color-primary-50); border: 1px solid var(--color-border);"
            >
              <User class="w-4 h-4" style="color: var(--color-primary);" />
            </div>
            <div class="flex-1 min-w-0">
              <h3
                class="text-sm font-semibold"
                style="color: var(--color-text-primary);"
              >
                个人信息编辑
              </h3>
              <p
                class="text-xs mt-0.5"
                style="color: var(--color-text-tertiary);"
              >
                修改头像、昵称、签名
              </p>
            </div>
            <ChevronRight class="w-4 h-4" style="color: var(--color-text-tertiary);" />
          </div>
        </BaseCard>

        <!-- Theme Select -->
        <BaseCard padding="md">
          <div class="flex items-center gap-3 mb-4">
            <div
              class="w-8 h-8 rounded-md flex items-center justify-center"
              style="background: var(--color-primary-50); border: 1px solid var(--color-border);"
            >
              <Palette class="w-4 h-4" style="color: var(--color-primary);" />
            </div>
            <h3
              class="text-sm font-semibold"
              style="color: var(--color-text-primary);"
            >
              主题切换
            </h3>
            <span
              class="text-[10px] ml-auto"
              style="font-family: var(--font-mono); color: var(--color-text-tertiary);"
            >
              THEME
            </span>
          </div>
          <div class="flex gap-2 ml-11">
            <button
              v-for="theme in themes"
              :key="theme"
              class="px-4 py-2 text-xs font-semibold rounded-md transition-all duration-150"
              :style="{
                background: themeStore.currentTheme === theme ? 'var(--color-primary-50)' : 'transparent',
                border: themeStore.currentTheme === theme ? '1px solid var(--color-primary)' : '1px solid var(--color-border-subtle)',
                color: themeStore.currentTheme === theme ? 'var(--color-primary)' : 'var(--color-text-tertiary)',
                fontFamily: 'var(--font-mono)'
              }"
              @click="selectTheme(theme)"
            >
              {{ theme.toUpperCase() }}
            </button>
          </div>
        </BaseCard>

        <!-- Privacy Settings -->
        <BaseCard
          padding="md"
          clickable
        >
          <div class="flex items-center gap-3">
            <div
              class="w-8 h-8 rounded-md flex items-center justify-center"
              style="background: var(--color-secondary-50); border: 1px solid var(--color-border);"
            >
              <EyeOff class="w-4 h-4" style="color: var(--color-secondary);" />
            </div>
            <div class="flex-1 min-w-0">
              <h3
                class="text-sm font-semibold"
                style="color: var(--color-text-primary);"
              >
                隐私设置
              </h3>
              <p
                class="text-xs mt-0.5"
                style="color: var(--color-text-tertiary);"
              >
                数据可见性与隐私保护
              </p>
            </div>
            <ChevronRight class="w-4 h-4" style="color: var(--color-text-tertiary);" />
          </div>
        </BaseCard>

        <!-- Account Security -->
        <BaseCard
          padding="md"
          clickable
        >
          <div class="flex items-center gap-3">
            <div
              class="w-8 h-8 rounded-md flex items-center justify-center"
              style="background: var(--color-state-success-50, rgba(0,255,136,0.08)); border: 1px solid var(--color-border);"
            >
              <Shield class="w-4 h-4" style="color: var(--color-state-success);" />
            </div>
            <div class="flex-1 min-w-0">
              <h3
                class="text-sm font-semibold"
                style="color: var(--color-text-primary);"
              >
                账户安全
              </h3>
              <p
                class="text-xs mt-0.5"
                style="color: var(--color-text-tertiary);"
              >
                密码修改、登录验证
              </p>
            </div>
            <ChevronRight class="w-4 h-4" style="color: var(--color-text-tertiary);" />
          </div>
        </BaseCard>
      </div>
    </div>

    <AppFooter />
  </div>
</template>