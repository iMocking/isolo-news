<script setup lang="ts">
import { onMounted } from 'vue'
import { RouterView } from 'vue-router'
import { useUserStore } from '@/stores/userStore'
import { useArticleStore } from '@/stores/articleStore'

const userStore = useUserStore()
const articleStore = useArticleStore()

onMounted(async () => {
  // 恢复用户登录会话
  await userStore.initialize()
  // 初始化文章数据
  articleStore.initialize()
})
</script>

<template>
  <RouterView />
</template>

<style>
/* Global styles are imported via main.css */
</style>