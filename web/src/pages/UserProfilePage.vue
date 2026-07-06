<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { BookOpen, Trophy, Flame, BarChart3, MessageSquare, Bookmark, LogOut, ArrowRight, Award, Star, Zap, Target } from 'lucide-vue-next'
import { useRouter } from 'vue-router'
import { useThemeStore } from '@/stores/themeStore'
import { useUserStore } from '@/stores/userStore'
import { useArticleStore } from '@/stores/articleStore'
import BaseProgressBar from '@/components/base/BaseProgressBar.vue'
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'

import { fetchFavorites } from '@/api/user'

const router = useRouter()
const themeStore = useThemeStore()
const userStore = useUserStore()
const articleStore = useArticleStore()

const isLogoutHovered = ref(false)
const statLabels = ref<string[]>([])
const achievementList = ref<Array<{name: string; description: string; icon: any; unlocked: boolean}>>([])

const achievementIcons = [Award, Star, Zap, Target]

onMounted(async () => {
  statLabels.value = ['已读文章', '获得成就', '连续登录', '积分排名', '发表评论', '收藏情报']
  
  try {
    if (userStore.achievements.length > 0) {
      achievementList.value = userStore.achievements.map((a: any, i: number) => ({
        name: a.name,
        description: a.description,
        icon: achievementIcons[i % achievementIcons.length],
        unlocked: !!a.unlockedAt,
      }))
    }
  } catch {
    // 使用默认数据
  }
  
  if (achievementList.value.length === 0) {
    achievementList.value = [
      { name: '初入情报网', description: '完成首次情报阅读', icon: Award, unlocked: true },
      { name: '连续签到达人', description: '连续签到30天', icon: Star, unlocked: true },
      { name: '评论之星', description: '发表100条评论', icon: Zap, unlocked: false },
      { name: '收藏家', description: '收藏50篇情报', icon: Target, unlocked: true },
    ]
  }

  try {
    const res = await fetchFavorites()
    if (res.code === 0 && res.data) {
      const favIds = res.data.map((a) => a.id)
      userStore.favoriteArticles = favIds
      res.data.forEach((a) => {
        if (!articleStore.articles.find((ea) => ea.id === a.id)) {
          articleStore.articles.push(a)
        }
      })
    }
  } catch {
    console.warn('收藏加载失败')
  }
})

const theme = computed(() => themeStore.currentTheme)

const user = computed(() => ({
  username: userStore.playerName,
  title: userStore.title,
  level: userStore.level,
  exp: userStore.xp,
  maxExp: userStore.maxXp,
  tags: userStore.tags
}))

const avatarText = computed(() => {
  const texts: Record<string, string> = {
    nexus: 'NX',
    comiket: 'CO',
    ironcore: 'IR'
  }
  return texts[theme.value] || 'NX'
})

const favoriteArticleList = computed(() => {
  return articleStore.articles.filter(a => userStore.favoriteArticles.includes(a.id))
})

const handleLogout = () => {
  userStore.logout()
  router.push('/')
}

const stats = computed(() => [
  { icon: BookOpen, label: statLabels.value[0] || '已读文章', value: userStore.readArticles, color: 'primary' },
  { icon: Trophy, label: statLabels.value[1] || '获得成就', value: userStore.achievementCount, color: 'secondary' },
  { icon: Flame, label: statLabels.value[2] || '连续登录', value: `${userStore.loginDays}天`, color: 'primary' },
  { icon: BarChart3, label: statLabels.value[3] || '积分排名', value: `#${userStore.rank}`, color: 'primary' },
  { icon: MessageSquare, label: statLabels.value[4] || '发表评论', value: userStore.achievementCount, color: 'secondary' },
  { icon: Bookmark, label: statLabels.value[5] || '收藏情报', value: userStore.favoriteArticles.length, color: 'primary' }
])

const achievements = computed(() => achievementList.value)

const heroBgStyle = computed(() => {
  const gradients: Record<string, string> = {
    nexus: 'linear-gradient(180deg, rgba(0,240,255,0.08) 0%, rgba(0,240,255,0.02) 50%, transparent 100%)',
    comiket: 'linear-gradient(180deg, rgba(255,107,43,0.08) 0%, rgba(255,107,43,0.02) 50%, transparent 100%)',
    ironcore: 'linear-gradient(180deg, rgba(240,160,48,0.08) 0%, rgba(240,160,48,0.02) 50%, transparent 100%)'
  }
  return {
    background: gradients[theme.value] || gradients.nexus
  }
})

const heroOverlayStyle = computed(() => {
  const colors: Record<string, string> = {
    nexus: 'radial-gradient(ellipse at 20% 0%, rgba(0,240,255,0.12) 0%, transparent 50%), radial-gradient(ellipse at 80% 100%, rgba(255,45,149,0.08) 0%, transparent 50%)',
    comiket: 'radial-gradient(ellipse at 20% 0%, rgba(255,107,43,0.12) 0%, transparent 50%), radial-gradient(ellipse at 80% 100%, rgba(59,158,255,0.08) 0%, transparent 50%)',
    ironcore: 'radial-gradient(ellipse at 20% 0%, rgba(240,160,48,0.12) 0%, transparent 50%), radial-gradient(ellipse at 80% 100%, rgba(78,205,196,0.08) 0%, transparent 50%)'
  }
  return {
    background: colors[theme.value] || colors.nexus
  }
})

const avatarStyle = computed(() => ({
  border: '2px solid var(--color-primary)',
  boxShadow: '0 0 20px var(--color-primary200), 0 0 40px var(--color-primary50)',
  background: 'var(--color-bgCard)'
}))

const avatarTextStyle = computed(() => ({
  color: 'var(--color-primary)',
  fontFamily: 'var(--font-display)'
}))

const levelBadgeStyle = computed(() => ({
  background: 'var(--color-secondary)',
  color: 'var(--color-textInverse)',
  fontFamily: 'var(--font-display)',
  border: '1px solid var(--color-secondaryDark)'
}))

const tagStyle = computed(() => ({
  background: 'var(--color-primary50)',
  color: 'var(--color-primary)',
  border: '1px solid var(--color-border)',
  fontFamily: 'var(--font-mono)'
}))

const statCardStyle = (stat: { color: string }) => ({
  background: 'var(--color-bgCard)',
  border: `1px solid ${stat.color === 'primary' ? 'var(--color-borderSubtle)' : 'var(--color-borderSubtle)'}`,
})

const achievementCardStyle = (achievement: { unlocked: boolean }) => ({
  background: 'var(--color-bgCard)',
  border: `1px solid ${achievement.unlocked ? 'var(--color-borderSubtle)' : 'var(--color-borderSubtle)'}`,
  opacity: achievement.unlocked ? 1 : 0.6
})

const achievementIconBoxStyle = (achievement: { unlocked: boolean }) => ({
  background: achievement.unlocked ? 'var(--color-primary50)' : 'var(--color-bgTertiary)',
  border: `1px solid ${achievement.unlocked ? 'var(--color-border)' : 'var(--color-borderSubtle)'}`
})

const logoutButtonStyle = computed(() => ({
  background: isLogoutHovered.value ? 'var(--color-stateError)' : 'var(--color-bgCard)',
  color: isLogoutHovered.value ? 'var(--color-textInverse)' : 'var(--color-stateError)',
  border: `1px solid var(--color-stateError)`,
}))
</script>

<template>
  <div class="min-h-screen" :style="{ background: 'var(--color-bgPrimary)' }">
    <AppNavigation />

    <main class="pt-16">
      <!-- Profile Hero -->
      <section class="relative overflow-hidden">
        <div class="absolute inset-0" :style="heroBgStyle"></div>
        <div class="absolute inset-0" :style="heroOverlayStyle"></div>
        
        <div class="relative z-10 max-w-5xl mx-auto px-6 py-12 md:py-16">
          <div class="flex flex-col md:flex-row items-center md:items-end gap-6 md:gap-8">
            <!-- Avatar -->
            <div class="relative shrink-0">
              <div
                class="w-28 h-28 md:w-36 md:h-36 rounded-xl flex items-center justify-center transition-all duration-200"
                :style="avatarStyle"
              >
                <span
                  class="text-4xl md:text-5xl font-bold"
                  :style="avatarTextStyle"
                >
                  {{ avatarText }}
                </span>
              </div>
              <!-- Level Badge -->
              <div
                class="absolute -bottom-2 -right-2 px-2.5 py-1 text-xs font-bold rounded-md"
                :style="levelBadgeStyle"
              >
                Lv.{{ user.level }}
              </div>
              <!-- Glow effect -->
              <div
                class="absolute -inset-1 rounded-xl -z-10 opacity-30 blur-md"
                :style="{ background: 'var(--color-primary)' }"
              ></div>
            </div>

            <!-- User Info -->
            <div class="text-center md:text-left flex-1 min-w-0">
              <h1
                class="text-2xl md:text-4xl font-bold tracking-wider mb-2"
                :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }"
              >
                {{ user.username }}
              </h1>
              <p
                class="text-sm mb-4 tracking-wide"
                :style="{ fontFamily: 'var(--font-mono)', color: 'var(--color-textSecondary)' }"
              >
                {{ user.title }}
              </p>
              <div class="flex flex-wrap justify-center md:justify-start gap-2">
                <span
                  v-for="tag in user.tags"
                  :key="tag"
                  class="text-xs px-3 py-1 rounded-full transition-all duration-150"
                  :style="tagStyle"
                >
                  {{ tag }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Level Bar -->
      <section class="max-w-5xl mx-auto px-6 -mt-4 relative z-20">
        <div
          class="rounded-lg p-5"
          :style="{ background: 'var(--color-bgCard)', border: '1px solid var(--color-borderSubtle)' }"
        >
          <BaseProgressBar
            :value="user.exp"
            :max="user.maxExp"
            variant="gradient"
            size="md"
            :show-value="false"
          />
          <div class="flex items-center justify-between text-xs mt-3" :style="{ fontFamily: 'var(--font-mono)', color: 'var(--color-textTertiary)' }">
            <span :style="{ color: 'var(--color-primary)' }">Lv.{{ user.level }}</span>
            <span>EXP {{ user.exp }} / {{ user.maxExp }}</span>
            <span :style="{ color: 'var(--color-secondary)' }">Lv.{{ user.level + 1 }}</span>
          </div>
        </div>
      </section>

      <!-- Stats Dashboard -->
      <section class="max-w-5xl mx-auto px-6 py-12">
        <div class="flex items-center gap-3 mb-6">
          <span class="w-1 h-5 rounded-sm" :style="{ background: 'var(--color-primary)' }"></span>
          <h2
            class="text-xl font-semibold tracking-wider"
            :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }"
          >
            能力面板
          </h2>
        </div>
        <div class="grid grid-cols-2 md:grid-cols-3 gap-4">
          <div
            v-for="stat in stats"
            :key="stat.label"
            class="rounded-lg p-5 transition-all duration-150 cursor-pointer hover:-translate-y-0.5"
            :style="statCardStyle(stat)"
          >
            <div class="flex items-center gap-2.5 mb-3">
              <div
                class="w-8 h-8 rounded-md flex items-center justify-center"
                :style="{ background: stat.color === 'primary' ? 'var(--color-primary50)' : 'var(--color-secondary50)' }"
              >
                <component
                  :is="stat.icon"
                  class="w-4 h-4"
                  :style="{ color: stat.color === 'primary' ? 'var(--color-primary)' : 'var(--color-secondary)' }"
                />
              </div>
              <span
                class="text-xs"
                :style="{ fontFamily: 'var(--font-mono)', color: 'var(--color-textTertiary)' }"
              >
                {{ stat.label }}
              </span>
            </div>
            <span
              class="text-2xl md:text-3xl font-bold tabular-nums"
              :style="{
                fontFamily: 'var(--font-display)',
                color: stat.color === 'secondary' ? 'var(--color-secondary)' : 'var(--color-textPrimary)'
              }"
            >
              {{ stat.value }}
            </span>
          </div>
        </div>
      </section>

      <!-- Achievements -->
      <section class="max-w-5xl mx-auto px-6 pb-12">
        <div class="flex items-center gap-3 mb-6">
          <span class="w-1 h-5 rounded-sm" :style="{ background: 'var(--color-secondary)' }"></span>
          <h2
            class="text-xl font-semibold tracking-wider"
            :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }"
          >
            成就徽章 ({{ achievements.length }})
          </h2>
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div
            v-for="(achievement, index) in achievements"
            :key="achievement.name"
            class="rounded-lg p-5 transition-all duration-150"
            :style="achievementCardStyle(achievement)"
          >
            <div
              class="w-12 h-12 mx-auto mb-3 rounded-lg flex items-center justify-center"
              :style="achievementIconBoxStyle(achievement)"
            >
              <component
                :is="achievement.icon"
                class="w-6 h-6"
                :style="{ color: achievement.unlocked ? 'var(--color-primary)' : 'var(--color-textTertiary)' }"
              />
            </div>
            <h3
              class="text-sm font-semibold mb-1 text-center"
              :style="{ color: achievement.unlocked ? 'var(--color-textPrimary)' : 'var(--color-textTertiary)' }"
            >
              {{ achievement.name }}
            </h3>
            <p
              class="text-xs text-center"
              :style="{ color: 'var(--color-textSecondary)', lineHeight: 1.5 }"
            >
              {{ achievement.description }}
            </p>
          </div>
        </div>
      </section>

      <!-- Favorite Articles -->
      <section class="max-w-5xl mx-auto px-6 pb-12">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center gap-3">
            <span class="w-1 h-5 rounded-sm" :style="{ background: 'var(--state-success)' }"></span>
            <h2
              class="text-xl font-semibold tracking-wider"
              :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }"
            >
              我的收藏
            </h2>
          </div>
          <span class="text-xs" :style="{ fontFamily: 'var(--font-mono)', color: 'var(--color-textTertiary)' }">
            共 {{ favoriteArticleList.length }} 篇
          </span>
        </div>
        
        <div v-if="favoriteArticleList.length > 0" class="space-y-3">
          <div
            v-for="article in favoriteArticleList"
            :key="article.id"
            class="rounded-lg p-4 cursor-pointer transition-all duration-150 hover:-translate-y-0.5"
            :style="{ background: 'var(--color-bgCard)', border: '1px solid var(--color-borderSubtle)' }"
            @click="router.push(`/article/${article.id}`)"
          >
            <div class="flex items-start gap-3">
              <div class="min-w-0 flex-1">
                <div class="flex items-center gap-2 mb-2">
                  <span
                    class="px-2 py-0.5 rounded text-[10px] font-bold"
                    :style="{ background: 'var(--color-primary50)', color: 'var(--color-primary)', fontFamily: 'var(--font-mono)' }"
                  >
                    {{ typeof article.category === 'string' ? article.category : article.category?.name }}
                  </span>
                  <span v-if="article.xpReward" class="text-[10px]" :style="{ color: 'var(--state-success)', fontFamily: 'var(--font-mono)' }">
                    +{{ article.xpReward }} XP
                  </span>
                </div>
                <h3 class="text-sm font-semibold mb-1 truncate" :style="{ color: 'var(--color-textPrimary)' }">
                  {{ article.title }}
                </h3>
                <p class="text-xs line-clamp-1" :style="{ color: 'var(--color-textSecondary)' }">
                  {{ article.summary }}
                </p>
              </div>
              <ArrowRight class="w-4 h-4 shrink-0 mt-2" :style="{ color: 'var(--color-textTertiary)' }" />
            </div>
          </div>
        </div>
        
        <div v-else class="rounded-lg p-10 text-center" :style="{ background: 'var(--color-bgCard)', border: '1px solid var(--color-borderSubtle)' }">
          <Bookmark class="w-12 h-12 mx-auto mb-3" :style="{ color: 'var(--color-textTertiary)', opacity: 0.5 }" />
          <p class="text-sm mb-1" :style="{ color: 'var(--color-textSecondary)' }">还没有收藏任何情报</p>
          <p class="text-xs" :style="{ color: 'var(--color-textTertiary)' }">点击文章详情页的收藏按钮添加收藏</p>
        </div>
      </section>

      <!-- Logout Button -->
      <section class="max-w-5xl mx-auto px-6 pb-16">
        <button
          class="w-full flex items-center justify-center gap-2 py-3.5 text-sm font-medium rounded-lg transition-all duration-150"
          :style="logoutButtonStyle"
          @mouseenter="isLogoutHovered = true"
          @mouseleave="isLogoutHovered = false"
          @click="handleLogout"
        >
          <LogOut class="w-4 h-4" />
          <span>退出登录</span>
        </button>
      </section>
    </main>

    <AppFooter />
  </div>
</template>

<style scoped>
</style>
