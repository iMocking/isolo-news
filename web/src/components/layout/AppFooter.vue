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

        <!-- Social Icons: GitHub / Bilibili / Douyin -->
        <div class="flex items-center gap-3">
          <a
            v-for="s in socialLinks"
            :key="s.name"
            :href="s.url"
            target="_blank"
            rel="noopener noreferrer"
            class="w-8 h-8 rounded-md flex items-center justify-center transition-all duration-150"
            :style="socialIconStyle"
            @mouseenter="onIconEnter"
            @mouseleave="onIconLeave"
          >
            <component :is="s.icon" v-if="s.icon" class="w-3.5 h-3.5" :style="iconStyle" />
            <span v-else class="w-3.5 h-3.5 flex items-center justify-center" :style="iconStyle">
              <!-- Bilibili 小电视 SVG -->
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" class="w-full h-full">
                <rect x="3" y="7" width="18" height="13" rx="2.5" />
                <path d="M5 7L3 2L9 7" />
                <path d="M19 7L21 2L15 7" />
                <path d="M10 11v5l5-2.5z" fill="currentColor" stroke="none" />
              </svg>
            </span>
          </a>
        </div>

        <!-- Copyright -->
        <p class="text-xs" :style="copyrightStyle">
          {{ copyrightText }}
        </p>
      </div>

      <!-- 备案信息 -->
      <div class="flex flex-wrap items-center justify-center gap-x-5 gap-y-1 mt-4 pt-4" :style="{ borderTop: '1px solid var(--color-borderSubtle)' }">
        <a
          :href="icpRecord.url"
          target="_blank"
          rel="noopener noreferrer"
          class="flex items-center gap-1.5 text-[11px] transition-colors duration-150"
          :style="{ color: 'var(--color-textTertiary)' }"
          @mouseenter="(e) => (e.currentTarget as HTMLElement).style.color = 'var(--color-primary)'"
          @mouseleave="(e) => (e.currentTarget as HTMLElement).style.color = 'var(--color-textTertiary)'"
        >
          <Shield class="w-3 h-3" />
          {{ icpRecord.label }}
        </a>
        <a
          :href="psRecord.url"
          target="_blank"
          rel="noopener noreferrer"
          class="flex items-center gap-1.5 text-[11px] transition-colors duration-150"
          :style="{ color: 'var(--color-textTertiary)' }"
          @mouseenter="(e) => (e.currentTarget as HTMLElement).style.color = 'var(--color-primary)'"
          @mouseleave="(e) => (e.currentTarget as HTMLElement).style.color = 'var(--color-textTertiary)'"
        >
          <Shield class="w-3 h-3" />
          {{ psRecord.label }}
        </a>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useThemeStore } from '@/stores/themeStore'
import { Github, Music, Shield } from 'lucide-vue-next'

const themeStore = useThemeStore()

/** 社交链接：GitHub、Bilibili（自定义 SVG）、抖音（音符图标） */
const socialLinks = [
  { name: 'github', url: 'https://github.com/iMocking', icon: Github },
  { name: 'bilibili', url: 'https://space.bilibili.com/94897302?spm_id_from=333.1007.0.0', icon: null },
  { name: 'douyin', url: 'https://douyin.com', icon: Music }
]

/** 网站备案信息 */
const icpRecord = {
  label: '鄂ICP备2024070712号-2',
  url: 'https://beian.miit.gov.cn/'
}

/** 公安备案信息 */
const psRecord = {
  label: '鄂公网安备42000000000000号',
  url: 'http://www.beian.gov.cn/'
}

/** 图标悬浮高亮 */
const onIconEnter = (e: Event) => {
  const target = e.currentTarget as HTMLElement
  target.style.borderColor = 'var(--color-primary)'
  target.style.background = 'var(--color-primary50)'
  const iconEl = target.querySelector('.w-3\\.5\\.h-3\\.5, .w-3\\.5\\.h-3\\.5 span, svg')
  if (iconEl) (iconEl as HTMLElement).style.color = 'var(--color-primary)'
}

/** 图标悬浮恢复 */
const onIconLeave = (e: Event) => {
  const target = e.currentTarget as HTMLElement
  target.style.borderColor = 'var(--color-borderSubtle)'
  target.style.background = 'var(--color-bgCard)'
  const iconEl = target.querySelector('.w-3\\.5\\.h-3\\.5, .w-3\\.5\\.h-3\\.5 span, svg')
  if (iconEl) (iconEl as HTMLElement).style.color = 'var(--color-textTertiary)'
}

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