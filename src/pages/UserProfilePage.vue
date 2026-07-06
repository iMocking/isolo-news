<script setup lang="ts">
import { computed } from 'vue'
import { BookOpen, Trophy, Flame, BarChart3, MessageSquare, Bookmark, LogOut, ArrowRight } from 'lucide-vue-next'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/userStore'
import { useArticleStore } from '@/stores/articleStore'
import BaseProgressBar from '@/components/base/BaseProgressBar.vue'
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'

const router = useRouter()
const userStore = useUserStore()
const articleStore = useArticleStore()

const user = computed(() => ({
  username: userStore.playerName,
  title: userStore.title,
  level: userStore.level,
  exp: userStore.xp,
  maxExp: userStore.maxXp,
  tags: userStore.tags
}))

const favoriteArticleList = computed(() => {
  return articleStore.articles.filter(a => userStore.favoriteArticles.includes(a.id))
})

const handleLogout = () => {
  userStore.logout()
  router.push('/')
}

const stats = [
  { icon: BookOpen, label: '已读情报', value: 284, color: 'primary' },
  { icon: Trophy, label: '获得成就', value: 42, color: 'secondary' },
  { icon: Flame, label: '连续签到', value: '7天', color: 'primary' },
  { icon: BarChart3, label: '全站排名', value: '#1,247', color: 'primary' },
  { icon: MessageSquare, label: '发表评论', value: 89, color: 'secondary' },
  { icon: Bookmark, label: '收藏情报', value: 156, color: 'primary' }
]

const achievements = [
  { name: '初入情报网', description: '完成首次情报阅读', icon: '📖', unlocked: true },
  { name: '连续签到达人', description: '连续签到30天', icon: '🔥', unlocked: true },
  { name: '评论之星', description: '发表100条评论', icon: '💬', unlocked: false },
  { name: '收藏家', description: '收藏50篇情报', icon: '📚', unlocked: true }
]
</script>

<template>
  <div class="min-h-screen">
    <AppNavigation />

    <!-- Profile Hero -->
    <section class="px-4 md:px-8 pt-8 pb-6 md:pt-12 md:pb-8">
      <div class="max-w-5xl mx-auto flex flex-col md:flex-row items-center md:items-start gap-6">
        <div class="relative">
          <div
            class="w-28 h-28 md:w-36 md:h-36 rounded-xl flex items-center justify-center"
            style="border: 2px solid var(--color-primary); box-shadow: 0 0 20px var(--color-primary-200), 0 0 40px var(--color-primary-50); background: linear-gradient(135deg, var(--color-primary-200), var(--color-secondary-100));"
          >
            <span
              class="text-4xl md:text-5xl font-bold"
              style="color: var(--color-text-inverse); font-family: var(--font-display);"
            >
              N
            </span>
          </div>
          <div
            class="absolute -bottom-1 -right-1 w-8 h-8 flex items-center justify-center text-xs font-bold"
            style="background: var(--color-secondary); color: white; border-radius: var(--radius-sm); font-family: var(--font-display);"
          >
            {{ user.level }}
          </div>
        </div>
        <div class="text-center md:text-left">
          <h1
            class="text-2xl md:text-3xl tracking-wider"
            style="font-family: var(--font-display); color: var(--color-primary);"
          >
            {{ user.username }}
          </h1>
          <p
            class="mt-1 text-sm tracking-wide"
            style="color: var(--color-text-secondary); font-family: var(--font-mono);"
          >
            {{ user.title }}
          </p>
          <div class="mt-3 flex flex-wrap justify-center md:justify-start gap-2">
            <span
              v-for="tag in user.tags"
              :key="tag"
              class="text-xs px-2 py-1"
              style="background: var(--color-primary-50); color: var(--color-primary); border: 1px solid var(--color-border); border-radius: var(--radius-full);"
            >
              {{ tag }}
            </span>
          </div>
        </div>
      </div>
    </section>

    <!-- Level Bar -->
    <section class="px-4 md:px-8 pb-6">
      <div class="max-w-5xl mx-auto">
        <BaseProgressBar
          :value="user.exp"
          :max="user.maxExp"
          variant="gradient"
          size="md"
          :show-value="true"
        />
        <div class="flex items-center justify-between text-xs mt-2" style="font-family: var(--font-mono); color: var(--color-text-secondary);">
          <span>Lv.{{ user.level }}</span>
          <span>EXP {{ user.exp }} / {{ user.maxExp }}</span>
        </div>
      </div>
    </section>

    <!-- Stats Dashboard -->
    <section class="px-4 md:px-8 pb-8">
      <div class="max-w-5xl mx-auto">
        <h2
          class="text-lg mb-4 tracking-wider"
          style="font-family: var(--font-display); color: var(--color-primary);"
        >
          能力面板
        </h2>
        <div class="grid grid-cols-2 md:grid-cols-3 gap-3 md:gap-4">
          <div
            v-for="stat in stats"
            :key="stat.label"
            class="p-4 transition-colors duration-150 hover:border-[var(--color-primary)]"
            style="background: var(--color-bg-card); border: 1px solid var(--color-border); border-radius: var(--radius-md);"
          >
            <div class="flex items-center gap-2 mb-2">
              <component
                :is="stat.icon"
                class="w-4 h-4"
                :style="{ color: `var(--color-${stat.color})` }"
              />
              <span
                class="text-xs"
                style="color: var(--color-text-tertiary); font-family: var(--font-mono);"
              >
                {{ stat.label }}
              </span>
            </div>
            <span
              class="text-2xl font-bold tabular-nums"
              style="font-family: var(--font-display);"
              :style="{ color: stat.color === 'secondary' ? `var(--color-${stat.color})` : 'var(--color-text-primary)' }"
            >
              {{ stat.value }}
            </span>
          </div>
        </div>
      </div>
    </section>

    <!-- Achievements -->
    <section class="px-4 md:px-8 pb-8">
      <div class="max-w-5xl mx-auto">
        <h2
          class="text-lg mb-4 tracking-wider"
          style="font-family: var(--font-display); color: var(--color-primary);"
        >
          成就徽章
        </h2>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div
            v-for="achievement in achievements"
            :key="achievement.name"
            class="p-4 transition-all duration-150"
            :style="{
              background: achievement.unlocked ? 'var(--color-bg-card)' : 'var(--color-bg-tertiary)',
              border: achievement.unlocked ? '1px solid var(--color-primary)' : '1px solid var(--color-border-subtle)',
              'border-radius': 'var(--radius-md)',
              opacity: achievement.unlocked ? 1 : 0.5
            }"
          >
            <div
              class="text-3xl mb-2 text-center"
            >
              {{ achievement.icon }}
            </div>
            <h3
              class="text-sm font-semibold mb-1 text-center"
              :style="{ color: achievement.unlocked ? 'var(--color-primary)' : 'var(--color-text-tertiary)' }"
            >
              {{ achievement.name }}
            </h3>
            <p
              class="text-xs text-center"
              style="color: var(--color-text-secondary);"
            >
              {{ achievement.description }}
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Favorite Articles -->
    <section class="px-4 md:px-8 pb-8">
      <div class="max-w-5xl mx-auto">
        <div class="flex items-center justify-between mb-4">
          <h2
            class="text-lg tracking-wider"
            style="font-family: var(--font-display); color: var(--color-primary);"
          >
            我的收藏
          </h2>
          <span class="text-xs" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">
            {{ favoriteArticleList.length }} 篇情报
          </span>
        </div>
        <div v-if="favoriteArticleList.length > 0" class="space-y-3">
          <div
            v-for="article in favoriteArticleList"
            :key="article.id"
            class="p-4 cursor-pointer transition-all duration-150 hover:translate-y-[-1px]"
            style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);"
            @click="router.push(`/article/${article.id}`)"
          >
            <div class="flex items-start gap-3">
              <div class="min-w-0 flex-1">
                <div class="flex items-center gap-2 mb-1">
                  <span
                    class="px-2 py-0.5 rounded text-[10px] font-bold"
                    style="background: var(--color-primary-50); color: var(--color-primary); font-family: var(--font-mono);"
                  >
                    {{ article.category }}
                  </span>
                  <span class="text-[10px]" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">
                    +{{ article.xpReward }} XP
                  </span>
                </div>
                <h3 class="text-sm font-semibold mb-1 truncate" style="color: var(--color-text-primary);">
                  {{ article.title }}
                </h3>
                <p class="text-xs line-clamp-1" style="color: var(--color-text-secondary);">
                  {{ article.summary }}
                </p>
              </div>
              <ArrowRight class="w-4 h-4 shrink-0 mt-2" style="color: var(--color-text-tertiary);" />
            </div>
          </div>
        </div>
        <div v-else class="p-8 text-center" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);">
          <Bookmark class="w-12 h-12 mx-auto mb-3" style="color: var(--color-text-tertiary); opacity: 0.5;" />
          <p class="text-sm" style="color: var(--color-text-tertiary);">还没有收藏任何情报</p>
          <p class="text-xs mt-1" style="color: var(--color-text-tertiary);">点击文章详情页的收藏按钮添加收藏</p>
        </div>
      </div>
    </section>

    <!-- Logout Button -->
    <section class="px-4 md:px-8 pb-12">
      <div class="max-w-5xl mx-auto">
        <button
          class="w-full flex items-center justify-center gap-2 py-3 text-sm font-medium transition-all duration-150"
          style="background: var(--color-bg-card); color: var(--state-error); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);"
          @mouseenter="($event.target as HTMLElement).style.background = 'rgba(255, 68, 68, 0.05)'; ($event.target as HTMLElement).style.borderColor = 'var(--state-error)'"
          @mouseleave="($event.target as HTMLElement).style.background = 'var(--color-bg-card)'; ($event.target as HTMLElement).style.borderColor = 'var(--color-border-subtle)'"
          @click="handleLogout"
        >
          <LogOut class="w-4 h-4" />
          <span>退出登录</span>
        </button>
      </div>
    </section>

    <AppFooter />
  </div>
</template>