<template>
  <!-- overflow-x-hidden 防止任何元素超出视口导致水平滚动条 -->
  <div class="min-h-screen overflow-x-hidden">
    <!-- Navigation（fixed 定位，脱离文档流） -->
    <AppNavigation />

    <!-- Main Content（pt-16 为固定导航栏让出 64px 空间） -->
    <main class="pt-16">
      <!-- Hero Section -->
      <HeroSection />

      <!-- Trending Topics -->
      <TrendingTopics />

      <!-- Article Grid with Sidebar -->
      <ArticleGrid :articles="articles" :empty-type="articleStore.apiOffline ? 'offline' : 'empty'">
        <template #sidebar>
          <PlayerStats />
        </template>
      </ArticleGrid>

      <!-- Daily Quests -->
      <DailyQuests />
    </main>

    <!-- Footer -->
    <AppFooter />
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useArticleStore } from '@/stores/articleStore'
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
import HeroSection from '@/components/business/HeroSection.vue'
import TrendingTopics from '@/components/business/TrendingTopics.vue'
import ArticleGrid from '@/components/business/ArticleGrid.vue'
import PlayerStats from '@/components/business/PlayerStats.vue'
import DailyQuests from '@/components/business/DailyQuests.vue'

const articleStore = useArticleStore()

const articles = computed(() => articleStore.articles)

/** 首页挂载时加载数据 */
onMounted(async () => {
  // 加载文章列表（如果尚未加载）
  if (articleStore.articles.length === 0) {
    await articleStore.loadArticles()
  }
  // 加载热门话题
  await articleStore.loadTrendingTopics()
})
</script>