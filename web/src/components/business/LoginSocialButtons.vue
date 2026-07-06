<template>
  <div class="flex justify-center gap-4">
    <button
      v-for="btn in ['github', 'chrome', 'twitter']"
      :key="btn"
      class="flex items-center justify-center w-12 h-12 transition-all duration-150"
      :style="{
        background: loginConfig.inputBg,
        border: `1px solid ${loginConfig.inputBorder}`,
        borderRadius: theme === 'ironcore' ? '4px' : '10px',
        color: theme === 'comiket' ? '#6b6158' : '#8892a8'
      }"
      @mouseenter="onHover($event, loginConfig)"
      @mouseleave="onLeave($event, loginConfig, theme)"
    >
      <component :is="iconMap[btn]" class="w-5 h-5" />
    </button>
  </div>
</template>

<script setup lang="ts">
import { Github, Chrome, Twitter } from 'lucide-vue-next'
import type { LoginThemeConfig } from './loginThemeConfig'

defineProps<{
  loginConfig: LoginThemeConfig
  theme: string
}>()

const iconMap: Record<string, any> = {
  github: Github,
  chrome: Chrome,
  twitter: Twitter,
}

const onHover = (event: MouseEvent, config: LoginThemeConfig) => {
  const el = event.currentTarget as HTMLElement
  el.style.borderColor = config.socialHoverBorder
  el.style.color = config.socialHoverColor
}

const onLeave = (event: MouseEvent, config: LoginThemeConfig, theme: string) => {
  const el = event.currentTarget as HTMLElement
  el.style.borderColor = config.inputBorder
  el.style.color = theme === 'comiket' ? '#6b6158' : '#8892a8'
}
</script>
