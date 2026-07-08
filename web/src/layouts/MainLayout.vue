<script setup lang="ts">
import { RouterView } from 'vue-router'
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
</script>
<template>
   <AppNavigation />
   <RouterView v-slot="{ Component }">
        <component :is="Component" />
    </RouterView>
    <AppFooter /> 
</template>