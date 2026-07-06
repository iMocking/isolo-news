<template>
  <!-- pt-16 由父级 <main> 统一提供，此处不再重复 -->
  <section class="relative w-full" :style="sectionStyle">
    <!-- Background Image -->
    <img
      :src="backgroundImage"
      :alt="backgroundAlt"
      class="absolute inset-0 w-full h-full object-cover"
    />
    
    <!-- Dark overlay -->
    <div class="absolute inset-0" :style="overlayStyle"></div>
    
    <!-- Scanline texture overlay -->
    <div v-if="showScanline" class="absolute inset-0 pointer-events-none opacity-[0.03]" :style="scanlineStyle"></div>

    <!-- Hero Content -->
    <div class="relative z-10 max-w-7xl mx-auto px-6 flex flex-col justify-end" :style="contentStyle">
      <!-- HUD Frame Decorations -->
      <div class="flex items-center gap-3 mb-4">
        <span
          class="w-3 h-3 rounded-sm"
          :style="hudDotStyle"
        ></span>
        <span
          class="text-xs tracking-widest uppercase"
          :style="hudTextStyle"
        >
          {{ hudText }}
        </span>
        <span
          class="flex-1 h-px"
          :style="hudLineStyle"
        ></span>
      </div>

      <h1
        class="mb-4"
        :style="titleStyle"
      >
        {{ heroTitle }}
      </h1>
      <p
        class="text-lg mb-8 max-w-xl"
        :style="subtitleStyle"
      >
        {{ heroSubtitle }}
      </p>

      <!-- CTA Buttons -->
      <div class="flex flex-wrap items-center gap-3">
        <button
          class="inline-flex items-center justify-center gap-2 px-6 py-3 text-sm font-semibold rounded-md whitespace-nowrap transition-all duration-150"
          :style="primaryButtonStyle"
          @mouseenter="handlePrimaryHover"
          @mouseleave="handlePrimaryLeave"
        >
          <component :is="primaryIcon" class="w-4 h-4" />
          {{ primaryButtonText }}
        </button>
        <button
          class="inline-flex items-center justify-center gap-2 px-6 py-3 text-sm font-semibold rounded-md whitespace-nowrap transition-all duration-150"
          :style="secondaryButtonStyle"
          @mouseenter="handleSecondaryHover"
          @mouseleave="handleSecondaryLeave"
        >
          <component :is="secondaryIcon" class="w-4 h-4" />
          {{ secondaryButtonText }}
        </button>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed, ref, type StyleValue } from 'vue'
import { useI18n } from 'vue-i18n'
import { useThemeStore } from '@/stores/themeStore'
import { Compass, ScrollText, Radar, FileText } from 'lucide-vue-next'

const { t } = useI18n()
const themeStore = useThemeStore()

const theme = computed(() => themeStore.currentTheme)

const heroImages: Record<string, { bg: string; alt: string }> = {
  nexus: {
    bg: '/images/image_0_yi19x4.jpg',
    alt: 'Cyberpunk cityscape with neon lights and futuristic technology'
  },
  comiket: {
    bg: '/images/image_2_yi19x4.jpg',
    alt: 'Anime and tech special edition'
  },
  ironcore: {
    bg: '/images/image_4_yi19x4.jpg',
    alt: 'Mecha cockpit interior'
  }
}

const heroIcons: Record<string, { primary: any; secondary: any }> = {
  nexus: { primary: Compass, secondary: ScrollText },
  comiket: { primary: Compass, secondary: ScrollText },
  ironcore: { primary: Radar, secondary: FileText }
}

const showScanline = computed(() => theme.value === 'nexus')

const backgroundImage = computed(() => heroImages[theme.value]?.bg || heroImages.nexus.bg)
const backgroundAlt = computed(() => heroImages[theme.value]?.alt || heroImages.nexus.alt)
const heroTitle = computed(() => t(`home.hero.title.${theme.value}`))
const heroSubtitle = computed(() => `${t(`home.hero.subtitle.${theme.value}`)} // ${t(`home.hero.highlight.${theme.value}`)}`)
const hudText = computed(() => t(`home.hero.hud.${theme.value}`))
const primaryButtonText = computed(() => t(`home.hero.btnPrimary.${theme.value}`))
const secondaryButtonText = computed(() => t(`home.hero.btnSecondary.${theme.value}`))
const primaryIcon = computed(() => heroIcons[theme.value]?.primary || Compass)
const secondaryIcon = computed(() => heroIcons[theme.value]?.secondary || ScrollText)

const sectionStyle = computed(() => ({
  minHeight: '520px'
}))

const overlayStyle = computed(() => {
  const theme = themeStore.currentTheme
  if (theme === 'nexus') {
    return {
      background: 'linear-gradient(180deg, rgba(10, 14, 26, 0.3) 0%, rgba(10, 14, 26, 0.6) 70%, var(--color-bgPrimary) 100%)'
    }
  } else if (theme === 'comiket') {
    return {
      background: 'linear-gradient(135deg, rgba(255,107,43,0.35) 0%, rgba(59,158,255,0.3) 100%)'
    }
  } else {
    return {
      background: 'linear-gradient(180deg, rgba(18,20,24,0.4) 0%, rgba(18,20,24,0.6) 50%, rgba(18,20,24,1) 100%)'
    }
  }
})

const scanlineStyle = computed(() => ({
  background: 'repeating-linear-gradient(0deg, transparent, transparent 2px, rgba(255, 255, 255, 0.08) 2px, rgba(255, 255, 255, 0.08) 4px)'
}))

const contentStyle = computed(() => ({
  minHeight: '520px',
  paddingTop: '120px',
  paddingBottom: '80px'
}))

const hudDotStyle = computed(() => ({
  background: `var(--color-primary)`,
  boxShadow: `0 0 8px rgba(0, 240, 255, 0.5)`
}))

const hudTextStyle = computed(() => ({
  fontFamily: `var(--font-mono)`,
  color: `var(--color-primary)`
}))

const hudLineStyle = computed(() => ({
  background: `linear-gradient(90deg, var(--color-border), transparent)`
}))

const titleStyle = computed<StyleValue>(() => ({
  fontFamily: `var(--font-display)`,
  fontSize: 'clamp(36px, 5vw, 72px)',
  fontWeight: 700,
  lineHeight: 1.05,
  letterSpacing: '-0.02em',
  color: themeStore.currentTheme === 'comiket' ? 'var(--color-textInverse)' : 'var(--color-textPrimary)',
  wordBreak: 'keep-all'
}))

const subtitleStyle = computed(() => ({
  color: themeStore.currentTheme === 'comiket' ? 'rgba(255,255,255,0.9)' : 'var(--color-textSecondary)',
  lineHeight: 1.6
}))

const isPrimaryHovered = ref(false)
const isSecondaryHovered = ref(false)

const primaryButtonStyle = computed(() => {
  const baseStyle = {
    background: themeStore.currentTheme === 'comiket' ? 'var(--color-textInverse)' : 'var(--color-primary)',
    color: themeStore.currentTheme === 'comiket' ? 'var(--color-primary)' : 'var(--color-textInverse)',
    fontFamily: `var(--font-display)`,
    borderRadius: `var(--radius-md)`
  }
  
  if (themeStore.currentTheme === 'nexus') {
    return {
      ...baseStyle,
      boxShadow: isPrimaryHovered.value 
        ? '0 0 30px rgba(0, 240, 255, 0.4)' 
        : '0 0 20px rgba(0, 240, 255, 0.25)',
      transform: isPrimaryHovered.value ? 'translateY(-1px)' : 'translateY(0)'
    }
  }
  
  return baseStyle
})

const secondaryButtonStyle = computed(() => ({
  background: isSecondaryHovered.value 
    ? (themeStore.currentTheme === 'comiket' ? 'rgba(255,255,255,0.2)' : 'var(--color-secondary50)')
    : 'transparent',
  color: themeStore.currentTheme === 'comiket' ? 'var(--color-textInverse)' : 'var(--color-secondary)',
  border: `1px solid ${themeStore.currentTheme === 'comiket' ? 'var(--color-textInverse)' : 'var(--color-secondary)'}`,
  fontFamily: `var(--font-display)`,
  borderRadius: `var(--radius-md)`,
  transform: isSecondaryHovered.value ? 'translateY(-1px)' : 'translateY(0)'
}))

const handlePrimaryHover = () => {
  isPrimaryHovered.value = true
}

const handlePrimaryLeave = () => {
  isPrimaryHovered.value = false
}

const handleSecondaryHover = () => {
  isSecondaryHovered.value = true
}

const handleSecondaryLeave = () => {
  isSecondaryHovered.value = false
}
</script>