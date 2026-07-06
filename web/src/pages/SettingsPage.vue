<script setup lang="ts">
import { Bell, User, Palette, EyeOff, Shield, ChevronRight, Check, X } from 'lucide-vue-next'
import NotificationToggles from '@/components/business/NotificationToggles.vue'
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import BaseCard from '@/components/base/BaseCard.vue'
import BaseButton from '@/components/base/BaseButton.vue'
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
import { useThemeStore } from '@/stores/themeStore'
import { useUserStore } from '@/stores/userStore'
import { updateProfile } from '@/api/user'

const { t } = useI18n()
const themeStore = useThemeStore()
const userStore = useUserStore()

const notifications = ref({
  push: true,
  email: true,
  inApp: false
})

const themes = ['nexus', 'comiket', 'ironcore']

const selectTheme = (theme: string) => {
  themeStore.setTheme(theme as any)
  // 同步主题偏好到后端
  if (userStore.isLoggedIn) {
    updateProfile({ themePreference: theme }).catch(() => {})
  }
}

// 个人资料编辑
const showEditForm = ref(false)
const editName = ref(userStore.playerName)
const editTitle = ref(userStore.title)
const editTagsText = ref(userStore.tags.join(', '))
const saveMessage = ref('')
const isSaving = ref(false)
const toggleEditForm = () => {
  showEditForm.value = !showEditForm.value
  if (showEditForm.value) {
    editName.value = userStore.playerName
    editTitle.value = userStore.title
    editTagsText.value = userStore.tags.join(', ')
    saveMessage.value = ''
  }
}

const handleSaveProfile = async () => {
  if (!editName.value.trim()) { saveMessage.value = '昵称不能为空'; return }
  isSaving.value = true
  saveMessage.value = ''
  try {
    const tags = editTagsText.value.split(',').map((t) => t.trim()).filter(Boolean)
    const res = await updateProfile({
      playerName: editName.value.trim(),
      title: editTitle.value.trim(),
      tags,
    })
    if (res.code === 0) {
      saveMessage.value = '保存成功 ✓'
      showEditForm.value = false
      // 刷新用户数据
      await userStore.refreshUserData()
    } else {
      saveMessage.value = res.message || '保存失败'
    }
  } catch {
    saveMessage.value = '网络错误，请稍后重试'
  } finally {
    isSaving.value = false
  }
}
const pageSubtitle = computed(() => t('settings.subtitle'))
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
          {{ pageSubtitle }}
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
                {{ t('settings.notifications.title') }}
              </h3>
            <span
              class="text-[10px] ml-auto"
              style="font-family: var(--font-mono); color: var(--color-text-tertiary);"
            >
              NOTIFICATION
            </span>
          </div>
          <NotificationToggles
            v-model="notifications"
            :items="[
              { key: 'push', label: t('settings.notifications.push') },
              { key: 'email', label: t('settings.notifications.email') },
              { key: 'inApp', label: t('settings.notifications.inApp') },
            ]" />
        </BaseCard>

        <!-- Profile Edit -->
        <BaseCard padding="md">
          <div
            class="flex items-center gap-3 cursor-pointer"
            @click="toggleEditForm"
          >
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
                {{ t('settings.profileEdit.title') }}
              </h3>
              <p
                class="text-xs mt-0.5"
                style="color: var(--color-text-tertiary);"
              >
                {{ userStore.playerName }} · {{ userStore.title }}
              </p>
            </div>
            <ChevronRight class="w-4 h-4" style="color: var(--color-text-tertiary);" />
          </div>

          <!-- 编辑表单 -->
          <div v-if="showEditForm" class="mt-4 ml-11 space-y-3">
            <div>
              <label class="block text-xs mb-1" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">昵称</label>
              <input
                v-model="editName"
                type="text"
                class="w-full px-3 py-2 text-sm outline-none"
                style="background: var(--color-bg-primary); border: 1px solid var(--color-border); color: var(--color-text-primary); border-radius: var(--radius-sm);"
              />
            </div>
            <div>
              <label class="block text-xs mb-1" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">头衔</label>
              <input
                v-model="editTitle"
                type="text"
                class="w-full px-3 py-2 text-sm outline-none"
                style="background: var(--color-bg-primary); border: 1px solid var(--color-border); color: var(--color-text-primary); border-radius: var(--radius-sm);"
              />
            </div>
            <div>
              <label class="block text-xs mb-1" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">标签（逗号分隔）</label>
              <input
                v-model="editTagsText"
                type="text"
                class="w-full px-3 py-2 text-sm outline-none"
                style="background: var(--color-bg-primary); border: 1px solid var(--color-border); color: var(--color-text-primary); border-radius: var(--radius-sm);"
              />
            </div>
            <div v-if="saveMessage" class="text-xs" :style="{ color: saveMessage.includes('成功') ? 'var(--state-success)' : 'var(--state-error)' }">
              {{ saveMessage }}
            </div>
            <div class="flex gap-2">
              <BaseButton size="sm" variant="primary" :disabled="isSaving" @click="handleSaveProfile">
                <template v-if="isSaving">保存中...</template>
                <template v-else><Check class="w-3 h-3" /> 保存</template>
              </BaseButton>
              <BaseButton size="sm" variant="ghost" @click="toggleEditForm">
                <X class="w-3 h-3" /> 取消
              </BaseButton>
            </div>
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
                {{ t('settings.theme.title') }}
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
                {{ t('settings.privacy.title') }}
              </h3>
              <p
                class="text-xs mt-0.5"
                style="color: var(--color-text-tertiary);"
              >
                {{ t('settings.privacy.desc') }}
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
                {{ t('settings.security.title') }}
              </h3>
              <p
                class="text-xs mt-0.5"
                style="color: var(--color-text-tertiary);"
              >
                {{ t('settings.security.desc') }}
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