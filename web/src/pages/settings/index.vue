<script setup lang="ts">
import { Bell, Palette, EyeOff, Shield } from 'lucide-vue-next'
import NotificationToggles from '@/components/business/NotificationToggles.vue'
import ProfileEditCard from '@/pages/profile/EditCard.vue'
import SystemInfoCard from '@/pages/profile/SystemInfoCard.vue'
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import BaseCard from '@/components/base/BaseCard.vue'
import { useThemeStore } from '@/stores/themeStore'
import { useUserStore } from '@/stores/userStore'
import { useCardStyles } from '@/hooks/useCardStyles'
import { updateProfile } from '@/api/user'

const { t } = useI18n()
const themeStore = useThemeStore()
const userStore = useUserStore()
const { getCardStyle } = useCardStyles()

const notifications = ref({
  push: true,
  email: true,
  inApp: false
})

const systemPrefs = ref({
  compactMode: false,
  reducedMotion: false,
  highContrast: false
})

const themes = ['nexus', 'comiket', 'ironcore']

const selectTheme = (theme: string) => {
  themeStore.setTheme(theme as any)
  if (userStore.isLoggedIn) {
    updateProfile({ themePreference: theme }).catch((err) => {
      console.warn('[Settings] 主题保存失败:', err)
    })
  }
}

const theme = computed(() => themeStore.currentTheme)

const pageTitle = computed(() => t(`settings.pageTitle.${theme.value}`))
const pageTag = computed(() => t(`settings.pageTag.${theme.value}`))
const pageSubtitle = computed(() => t(`settings.pageSubtitle.${theme.value}`))

const themeLabels: Record<string, string> = {
  nexus: t('settings.theme.nexus'),
  comiket: t('settings.theme.comiket'),
  ironcore: t('settings.theme.ironcore')
}

const systemInfo = [
  { icon: 'monitor', label: t('settings.system.clientVersion'), value: 'v2.4.0' },
  { icon: 'globe', label: t('settings.system.locale'), value: 'zh-CN' },
  { icon: 'database', label: t('settings.system.dataSync'), value: t('settings.system.connected') },
  { icon: 'hard-drive', label: t('settings.system.localCache'), value: '128 MB' }
]
</script>

<template>
  <div class="min-h-screen" :style="{ background: 'var(--color-bgPrimary)' }">
    <main class="max-w-3xl mx-auto px-6 pt-24 pb-16">
      <!-- Header -->
      <section class="relative overflow-hidden rounded-2xl mb-10 p-8" :style="getCardStyle('panel', false)">
        <div class="relative z-10">
          <div class="flex items-center gap-3 mb-4">
            <span class="w-3 h-3 rounded-sm" :style="{ background: 'var(--color-primary)', boxShadow: '0 0 12px var(--color-primary200)' }" />
            <h1 class="text-2xl md:text-3xl font-bold tracking-wider" :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }">
              {{ pageTitle }}
            </h1>
          </div>
          <p class="text-sm mb-3" :style="{ color: 'var(--color-textSecondary)', lineHeight: 1.7, maxWidth: '520px' }">
            {{ pageSubtitle }}
          </p>
          <span class="text-xs font-mono" :style="{ color: 'var(--color-primary)' }">{{ pageTag }}</span>
        </div>
        <div
          class="absolute top-0 right-0 w-1/2 h-full pointer-events-none opacity-20"
          :style="{ background: 'radial-gradient(circle at 80% 20%, var(--color-primary), transparent 60%)' }"
        />
      </section>

      <!-- Settings Cards -->
      <div class="space-y-4">
        <!-- Theme Select -->
        <BaseCard padding="md">
          <div class="flex items-center gap-3 mb-5">
            <div
              class="w-8 h-8 rounded-md flex items-center justify-center"
              :style="{ background: 'var(--color-primary50)', border: '1px solid var(--color-border)' }"
            >
              <Palette class="w-4 h-4" :style="{ color: 'var(--color-primary)' }" />
            </div>
            <h3 class="text-sm font-semibold" :style="{ color: 'var(--color-textPrimary)' }">
              {{ t('settings.theme.title') }}
            </h3>
            <span class="text-[10px] ml-auto font-mono" :style="{ color: 'var(--color-textTertiary)' }">
              THEME
            </span>
          </div>
          <div class="grid grid-cols-1 sm:grid-cols-3 gap-3 ml-11">
            <button
              v-for="themeItem in themes"
              :key="themeItem"
              class="px-3 py-3 text-xs font-semibold rounded-md transition-all duration-150 text-left"
              :style="{
                background: themeStore.currentTheme === themeItem ? 'var(--color-primary50)' : 'transparent',
                border: themeStore.currentTheme === themeItem ? '1px solid var(--color-primary)' : '1px solid var(--color-borderSubtle)',
                color: themeStore.currentTheme === themeItem ? 'var(--color-primary)' : 'var(--color-textTertiary)',
                fontFamily: 'var(--font-mono)'
              }"
              @click="selectTheme(themeItem)"
            >
              <span class="block text-[10px] opacity-70 mb-0.5">{{ themeItem.toUpperCase() }}</span>
              {{ themeLabels[themeItem] }}
            </button>
          </div>
        </BaseCard>

        <!-- Notification Settings -->
        <BaseCard padding="md">
          <div class="flex items-center gap-3 mb-5">
            <div
              class="w-8 h-8 rounded-md flex items-center justify-center"
              :style="{ background: 'var(--color-primary50)', border: '1px solid var(--color-border)' }"
            >
              <Bell class="w-4 h-4" :style="{ color: 'var(--color-primary)' }" />
            </div>
            <h3 class="text-sm font-semibold" :style="{ color: 'var(--color-textPrimary)' }">
              {{ t('settings.notifications.title') }}
            </h3>
            <span class="text-[10px] ml-auto font-mono" :style="{ color: 'var(--color-textTertiary)' }">
              NOTIFICATION
            </span>
          </div>
          <div class="ml-11">
            <NotificationToggles
              v-model="notifications"
              :items="[
                { key: 'push', label: t('settings.notifications.push') },
                { key: 'email', label: t('settings.notifications.email') },
                { key: 'inApp', label: t('settings.notifications.inApp') },
              ]" />
          </div>
        </BaseCard>

        <!-- 个人信息编辑（折叠卡片） -->
        <ProfileEditCard />

        <!-- System Preferences -->
        <BaseCard padding="md">
          <div class="flex items-center gap-3 mb-5">
            <div
              class="w-8 h-8 rounded-md flex items-center justify-center"
              :style="{ background: 'var(--color-secondary50)', border: '1px solid var(--color-border)' }"
            >
              <Monitor class="w-4 h-4" :style="{ color: 'var(--color-secondary)' }" />
            </div>
            <h3 class="text-sm font-semibold" :style="{ color: 'var(--color-textPrimary)' }">
               {{ t('settings.system.title') }}
            </h3>
            <span class="text-[10px] ml-auto font-mono" :style="{ color: 'var(--color-textTertiary)' }">
              PREFERENCE
            </span>
          </div>
          <div class="ml-11 space-y-3">
            <label
              v-for="(value, key) in systemPrefs"
              :key="key"
              class="flex items-center justify-between cursor-pointer"
            >
              <span class="text-sm" :style="{ color: 'var(--color-textSecondary)' }">
                {{ key === 'compactMode' ? t('settings.system.compactMode') : key === 'reducedMotion' ? t('settings.system.reduceAnimation') : t('settings.system.highContrast') }}
              </span>
              <input
                v-model="systemPrefs[key as keyof typeof systemPrefs]"
                type="checkbox"
                class="w-4 h-4"
                :style="{ accentColor: 'var(--color-primary)' }"
              />
            </label>
          </div>
        </BaseCard>

        <!-- Privacy Settings -->
        <BaseCard padding="md" clickable>
          <div class="flex items-center gap-3">
            <div
              class="w-8 h-8 rounded-md flex items-center justify-center"
              :style="{ background: 'var(--color-secondary50)', border: '1px solid var(--color-border)' }"
            >
              <EyeOff class="w-4 h-4" :style="{ color: 'var(--color-secondary)' }" />
            </div>
            <div class="flex-1 min-w-0">
              <h3 class="text-sm font-semibold" :style="{ color: 'var(--color-textPrimary)' }">
                {{ t('settings.privacy.title') }}
              </h3>
              <p class="text-xs mt-0.5" :style="{ color: 'var(--color-textTertiary)' }">
                {{ t('settings.privacy.desc') }}
              </p>
            </div>
            <ChevronRight class="w-4 h-4" :style="{ color: 'var(--color-textTertiary)' }" />
          </div>
        </BaseCard>

        <!-- Account Security -->
        <BaseCard padding="md" clickable>
          <div class="flex items-center gap-3">
            <div
              class="w-8 h-8 rounded-md flex items-center justify-center"
              :style="{ background: 'var(--color-state-success-50, rgba(0,255,136,0.08))', border: '1px solid var(--color-border)' }"
            >
              <Shield class="w-4 h-4" :style="{ color: 'var(--state-success)' }" />
            </div>
            <div class="flex-1 min-w-0">
              <h3 class="text-sm font-semibold" :style="{ color: 'var(--color-textPrimary)' }">
                {{ t('settings.security.title') }}
              </h3>
              <p class="text-xs mt-0.5" :style="{ color: 'var(--color-textTertiary)' }">
                {{ t('settings.security.desc') }}
              </p>
            </div>
            <ChevronRight class="w-4 h-4" :style="{ color: 'var(--color-textTertiary)' }" />
          </div>
        </BaseCard>

        <!-- 系统信息 -->
        <SystemInfoCard :items="systemInfo" />
      </div>
    </main>
  </div>
</template>
