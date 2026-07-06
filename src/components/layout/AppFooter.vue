<template>
  <footer :style="footerStyle">
    <div class="max-w-7xl mx-auto px-6 py-8">
      <div class="flex flex-col md:flex-row items-center justify-between gap-4">
        <!-- Logo -->
        <a href="#" class="flex items-center gap-2 shrink-0" :style="logoStyle">
          <span
            class="text-sm font-bold tracking-wider"
            :style="brandTextStyle"
          >
            {{ currentBrandName }}
          </span>
          <span
            class="text-xs"
            :style="subTextStyle"
          >
            DAILY
          </span>
        </a>

        <!-- Social Icons -->
        <div class="flex items-center gap-3">
          <a
            v-for="social in socialLinks"
            :key="social.name"
            :href="social.url"
            class="w-8 h-8 rounded-md flex items-center justify-center transition-colors duration-150"
            :style="socialIconStyle"
          >
            <component :is="social.icon" class="w-3.5 h-3.5" :style="iconStyle" />
          </a>
        </div>

        <!-- Copyright -->
        <p class="text-xs" :style="copyrightStyle">
          {{ copyrightText }}
        </p>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useThemeStore } from '@/stores/themeStore'
import { Twitter, Github, Youtube, Rss } from 'lucide-vue-next'

const themeStore = useThemeStore()

const socialLinks = [
  { name: 'twitter', url: '#', icon: Twitter },
  { name: 'github', url: '#', icon: Github },
  { name: 'youtube', url: '#', icon: Youtube },
  { name: 'rss', url: '#', icon: Rss }
]

const currentBrandName = computed(() => {
  const themeNames = {
    nexus: 'NEXUS',
    comiket: 'COMIKET',
    ironcore: 'IRONCORE'
  }
  return themeNames[themeStore.currentTheme]
})

const copyrightText = computed(() => {
  const texts = {
    nexus: '2026 NEXUS Daily. All rights reserved.',
    comiket: '2026 COMIKET. All rights reserved.',
    ironcore: '© 2026 IRONCORE NETWORK. ALL INTEL CLASSIFIED.'
  }
  return texts[themeStore.currentTheme]
})

const footerStyle = computed(() => ({
  background: `var(--color-bgSecondary)`,
  borderTop: `1px solid var(--color-borderSubtle)`
}))

const logoStyle = computed(() => ({
  fontFamily: `var(--font-display)`
}))

const brandTextStyle = computed(() => ({
  color: `var(--color-primary)`
}))

const subTextStyle = computed(() => ({
  color: `var(--color-textTertiary)`
}))

const socialIconStyle = computed(() => ({
  background: `var(--color-bgCard)`,
  border: `1px solid var(--color-borderSubtle)`
}))

const iconStyle = computed(() => ({
  color: `var(--color-textTertiary)`
}))

const copyrightStyle = computed(() => ({
  color: `var(--color-textTertiary)`
}))
</script>