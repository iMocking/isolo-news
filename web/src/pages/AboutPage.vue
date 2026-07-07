<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { Twitter, Github, Mail, Cpu, Users, Layers, MessageCircle } from 'lucide-vue-next'
import BaseCard from '@/components/base/BaseCard.vue'
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
import { useThemeStore } from '@/stores/themeStore'
import { useCardStyles } from '@/hooks/useCardStyles'

const { t, tm } = useI18n()
const themeStore = useThemeStore()
const { getCardStyle } = useCardStyles()

const teamMembers = ref<Array<{ initials: string; name: string; role: string; level: number }>>([])

onMounted(() => {
  const members = tm('about.teamMembers') as unknown as Array<{ name: string; role: string }>
  teamMembers.value = members.map((m: any, i: number) => ({
    initials: ['NX', 'GX', 'AX'][i] || 'NX',
    name: m.name,
    role: m.role,
    level: [58, 45, 51][i] || 1
  }))
})

const theme = computed(() => themeStore.currentTheme)

const pageTitle = computed(() => t(`about.pageTitle.${theme.value}`))
const pageTag = computed(() => t(`about.pageTag.${theme.value}`))

const techStack = ['VUE 3', 'TYPESCRIPT', 'TAILWIND CSS', 'VITE']

const featureIcons = [Cpu, Users, Layers, MessageCircle]

const platformFeatures = computed(() => {
  const features = tm('about.features') as Array<{ title: string; desc: string }>
  return features.map((f, i) => ({
    icon: featureIcons[i] || Cpu,
    title: f.title,
    desc: f.desc
  }))
})
</script>

<template>
  <div class="min-h-screen" :style="{ background: 'var(--color-bgPrimary)' }">
    <AppNavigation />

    <!-- About Content -->
    <main class="max-w-5xl mx-auto px-6 pt-24 pb-16">
      <!-- Header -->
      <section class="relative overflow-hidden rounded-2xl mb-10 p-8 md:p-12" :style="getCardStyle('panel', false)">
        <div class="relative z-10">
          <div class="flex items-center gap-3 mb-4">
            <span class="w-3 h-3 rounded-sm" :style="{ background: 'var(--color-primary)', boxShadow: '0 0 12px var(--color-primary200)' }" />
            <h1 class="text-2xl md:text-4xl font-bold tracking-wider" :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }">
              {{ pageTitle }}
            </h1>
          </div>
          <p class="text-sm md:text-base mb-4" :style="{ color: 'var(--color-textSecondary)', lineHeight: 1.7, maxWidth: '640px' }">
            {{ $t('about.intro') }}
          </p>
          <span class="text-xs font-mono" :style="{ color: 'var(--color-primary)' }">{{ pageTag }}</span>
        </div>
        <!-- Decorative gradient -->
        <div
          class="absolute top-0 right-0 w-1/2 h-full pointer-events-none opacity-20"
          :style="{ background: 'radial-gradient(circle at 80% 20%, var(--color-primary), transparent 60%)' }"
        />
      </section>

      <!-- Platform Intro -->
      <BaseCard padding="md" class="mb-8">
        <div class="flex items-center gap-3 mb-6">
          <span class="w-1 h-5 rounded-sm" :style="{ background: 'var(--color-primary)' }" />
          <h2 class="text-lg font-semibold" :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }">
            {{ $t('about.sections.platform') }}
          </h2>
        </div>
        <p class="text-sm leading-relaxed mb-6" :style="{ color: 'var(--color-textSecondary)', lineHeight: 1.8 }">
          {{ $t('about.platformDesc') }}
        </p>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div
            v-for="feature in platformFeatures"
            :key="feature.title"
            class="rounded-lg p-4 transition-all duration-150"
            :style="getCardStyle('default', false)"
          >
            <div class="flex items-center gap-3 mb-2">
              <component :is="feature.icon" class="w-4 h-4" :style="{ color: 'var(--color-primary)' }" />
              <h4 class="text-sm font-semibold" :style="{ color: 'var(--color-textPrimary)' }">{{ feature.title }}</h4>
            </div>
            <p class="text-xs" :style="{ color: 'var(--color-textSecondary)', lineHeight: 1.6 }">{{ feature.desc }}</p>
          </div>
        </div>
      </BaseCard>

      <!-- Team Section -->
      <section class="mb-8">
        <div class="flex items-center gap-3 mb-6">
          <span class="w-1 h-5 rounded-sm" :style="{ background: 'var(--color-secondary)' }" />
          <h2 class="text-lg font-semibold" :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }">
            {{ $t('about.sections.team') }}
          </h2>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
          <div
            v-for="member in teamMembers"
            :key="member.initials"
            class="rounded-lg p-5 text-center transition-all duration-160"
            :style="getCardStyle('elevated', false)"
          >
            <div
              class="w-14 h-14 rounded-full mx-auto mb-3 flex items-center justify-center text-sm font-bold"
              :style="{ background: 'linear-gradient(135deg, var(--color-primary), var(--color-secondary))', color: 'var(--color-textInverse)' }"
            >
              {{ member.initials }}
            </div>
            <h4 class="text-sm font-semibold mb-1" :style="{ color: 'var(--color-textPrimary)' }">
              {{ member.name }}
            </h4>
            <p class="text-xs mb-2" :style="{ color: 'var(--color-textTertiary)' }">
              {{ member.role }}
            </p>
            <span class="text-[10px] px-2 py-0.5 rounded font-mono" :style="{ background: 'var(--color-primary50)', color: 'var(--color-primary)' }">
              Lv.{{ member.level }}
            </span>
          </div>
        </div>
      </section>

      <!-- Tech Stack -->
      <BaseCard padding="md" class="mb-8">
        <div class="flex items-center gap-3 mb-6">
          <span class="w-1 h-5 rounded-sm" :style="{ background: 'var(--state-success)' }" />
          <h2 class="text-lg font-semibold" :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }">
            {{ $t('about.sections.tech') }}
          </h2>
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-3">
          <div
            v-for="tech in techStack"
            :key="tech"
            class="rounded-md p-3 text-center transition-all duration-150"
            :style="getCardStyle('stat', false)"
          >
            <div class="text-xs font-semibold" :style="{ color: 'var(--color-primary)', fontFamily: 'var(--font-mono)' }">
              {{ tech }}
            </div>
          </div>
        </div>
      </BaseCard>

      <!-- Contact -->
      <BaseCard padding="md">
        <div class="flex items-center gap-3 mb-6">
          <span class="w-1 h-5 rounded-sm" :style="{ background: 'var(--state-warning)' }" />
          <h2 class="text-lg font-semibold" :style="{ fontFamily: 'var(--font-display)', color: 'var(--color-textPrimary)' }">
            {{ $t('about.sections.contact') }}
          </h2>
        </div>
        <div class="flex items-center gap-4">
          <button
            v-for="(icon, idx) in [Twitter, Github, Mail]"
            :key="idx"
            class="w-10 h-10 rounded-md flex items-center justify-center transition-all duration-150"
            :style="getCardStyle('default', false)"
          >
            <component :is="icon" class="w-4 h-4" :style="{ color: 'var(--color-textSecondary)' }" />
          </button>
        </div>
      </BaseCard>
    </main>

    <AppFooter />
  </div>
</template>
