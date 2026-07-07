<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { LogOut, Award, Star, Zap, Target } from 'lucide-vue-next'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/userStore'
import { useArticleStore } from '@/stores/articleStore'
import BaseProgressBar from '@/components/base/BaseProgressBar.vue'
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
import ProfileHeroSection from '@/components/profile/ProfileHeroSection.vue'
import StatsPanel from '@/components/profile/StatsPanel.vue'
import AchievementGrid from '@/components/profile/AchievementGrid.vue'
import FavoriteArticleList from '@/components/profile/FavoriteArticleList.vue'

import { fetchFavorites } from '@/api/user'
import { useCardStyles } from '@/hooks/useCardStyles'

const router = useRouter()
const userStore = useUserStore()
const articleStore = useArticleStore()
const { getCardStyle } = useCardStyles()

const isLogoutHovered = ref(false)
const achievementList = ref<Array<{name: string; description: string; icon: any; unlocked: boolean}>>([])

const achievementIcons = [Award, Star, Zap, Target]

onMounted(async () => {
  await userStore.refreshUserData()

  // 从后端获取成就列表
  try {
    achievementList.value = userStore.achievements.map((a: any, i: number) => ({
      name: a.name,
      description: a.description,
      icon: achievementIcons[i % achievementIcons.length],
      unlocked: !!a.unlockedAt,
    }))
  } catch {
    achievementList.value = []
  }

  // 加载收藏
  try {
    const res = await fetchFavorites()
    if (res.code === 0 && res.data) {
      const favIds = res.data.map((a: any) => a.id)
      userStore.favoriteArticles = favIds
      res.data.forEach((a: any) => {
        if (!articleStore.articles.find((ea) => ea.id === a.id)) {
          articleStore.articles.push(a)
        }
      })
    }
  } catch {
    console.warn('收藏加载失败')
  }
})

/** 等级 / XP 数据（供等级条使用） */
const userLevelInfo = computed(() => ({
  level: userStore.level,
  exp: userStore.xp,
  maxExp: userStore.maxXp,
}))

/** 收藏文章列表 */
const favoriteArticleList = computed(() => {
  return articleStore.articles.filter(a => userStore.favoriteArticles.includes(a.id))
})

const levelBarStyle = computed(() => getCardStyle('panel', false))

const logoutButtonStyle = computed(() => ({
  background: isLogoutHovered.value ? 'var(--color-stateError)' : 'var(--color-bgCard)',
  color: isLogoutHovered.value ? 'var(--color-textInverse)' : 'var(--color-stateError)',
  border: '1px solid var(--color-stateError)',
}))

const handleLogout = () => {
  userStore.logout()
  router.push('/')
}
</script>

<template>
  <div class="min-h-screen" :style="{ background: 'var(--color-bgPrimary)' }">
    <AppNavigation />

    <main class="pt-16">
      <!-- 用户头像与信息 -->
      <ProfileHeroSection />

      <!-- 等级 / XP 进度条 -->
      <section class="max-w-5xl mx-auto px-6 -mt-4 relative z-20">
        <div class="rounded-lg p-5" :style="levelBarStyle">
          <BaseProgressBar
            :value="userLevelInfo.exp"
            :max="userLevelInfo.maxExp"
            variant="gradient"
            size="md"
            :show-value="false"
          />
          <div class="flex items-center justify-between text-xs mt-3" :style="{ fontFamily: 'var(--font-mono)', color: 'var(--color-textTertiary)' }">
            <span :style="{ color: 'var(--color-primary)' }">Lv.{{ userLevelInfo.level }}</span>
            <span>EXP {{ userLevelInfo.exp }} / {{ userLevelInfo.maxExp }}</span>
            <span :style="{ color: 'var(--color-secondary)' }">Lv.{{ userLevelInfo.level + 1 }}</span>
          </div>
        </div>
      </section>

      <!-- 能力面板 -->
      <StatsPanel />

      <!-- 成就徽章 -->
      <AchievementGrid :achievements="achievementList" />

      <!-- 我的收藏 -->
      <FavoriteArticleList :articles="favoriteArticleList" />

      <!-- 退出登录 -->
      <section class="max-w-5xl mx-auto px-6 pb-16">
        <button
          class="w-full flex items-center justify-center gap-2 py-3.5 text-sm font-medium rounded-lg transition-all duration-150"
          :style="logoutButtonStyle"
          @mouseenter="isLogoutHovered = true"
          @mouseleave="isLogoutHovered = false"
          @click="handleLogout"
        >
          <LogOut class="w-4 h-4" />
          <span>{{ $t('profile.logout') }}</span>
        </button>
      </section>
    </main>

    <AppFooter />
  </div>
</template>

<style scoped>
</style>
