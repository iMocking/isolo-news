<template>
  <div class="min-h-screen" :style="{ background: 'var(--color-bgPrimary)' }">
    <AppNavigation />
    <router-view />
    <AppFooter />
  </div>
</template>

<script setup lang="ts">
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
</script>
