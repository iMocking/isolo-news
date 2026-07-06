<template>
  <div class="min-h-screen">
    <!-- Navigation -->
    <AppNavigation />

    <!-- Main Content -->
    <main>
      <!-- Hero Section -->
      <HeroSection />

      <!-- Trending Topics -->
      <TrendingTopics />

      <!-- Article Grid with Sidebar -->
      <ArticleGrid :articles="articles">
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
import { computed } from 'vue'
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
</script>