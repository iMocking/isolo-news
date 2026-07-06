<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { Twitter, Github, Mail } from 'lucide-vue-next'
import BaseCard from '@/components/base/BaseCard.vue'
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'

const { t, tm } = useI18n()

const teamMembers = ref<Array<{initials: string; name: string; role: string; level: number}>>([])

onMounted(() => {
  const members = tm('about.teamMembers') as unknown as Array<{name: string; role: string}>
  teamMembers.value = members.map((m: any, i: number) => ({
    initials: ['NX', 'GX', 'AX'][i] || 'NX',
    name: m.name,
    role: m.role,
    level: [58, 45, 51][i] || 1
  }))
})

const techStack = ['VUE', 'TAILWIND', 'NODE.JS', 'POSTGRES']
</script>

<template>
  <div class="min-h-screen">
    <AppNavigation />
    
    <!-- About Content -->
    <div class="max-w-4xl mx-auto px-6 pt-24 pb-16">
      <!-- Title -->
      <div class="flex items-center gap-3 mb-4">
        <span
          class="w-3 h-3 rounded-sm"
          style="background: var(--color-primary); box-shadow: 0 0 8px rgba(0, 240, 255, 0.5);"
        />
        <h1
          class="text-2xl font-bold"
          style="font-family: var(--font-display); color: var(--color-text-primary);"
        >
          ABOUT NEXUS
        </h1>
        <span
          class="text-sm"
          style="color: var(--color-text-tertiary);"
        >
          {{ $t('about.subtitle') }}
        </span>
      </div>
      <p
        class="text-sm mb-12"
        style="color: var(--color-text-secondary); line-height: 1.6;"
      >
        {{ $t('about.intro') }}
      </p>

      <!-- Intro Section -->
      <BaseCard
        padding="md"
        class="mb-8"
      >
        <div class="flex items-center gap-3 mb-4">
          <span
            class="w-1 h-5 rounded-sm"
            style="background: var(--color-primary);"
          />
          <h2
            class="text-lg font-semibold"
            style="font-family: var(--font-display); color: var(--color-text-primary);"
          >
             {{ $t('about.sections.platform') }}
          </h2>
        </div>
        <p
          class="text-sm leading-relaxed"
          style="color: var(--color-text-secondary); line-height: 1.7;"
        >
          {{ $t('about.platformDesc') }}
        </p>
      </BaseCard>

      <!-- Team Section -->
      <section class="mb-8">
        <div class="flex items-center gap-3 mb-6">
          <span
            class="w-1 h-5 rounded-sm"
            style="background: var(--color-secondary);"
          />
          <h2
            class="text-lg font-semibold"
            style="font-family: var(--font-display); color: var(--color-text-primary);"
          >
             {{ $t('about.sections.team') }}
          </h2>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
          <div
            v-for="member in teamMembers"
            :key="member.initials"
            class="rounded-lg p-5 text-center transition-all duration-160"
            style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle);"
          >
            <div
              class="w-14 h-14 rounded-full mx-auto mb-3 flex items-center justify-center text-sm font-bold"
              style="background: linear-gradient(135deg, var(--color-primary-dark), var(--color-secondary-dark)); color: var(--color-text-primary);"
            >
              {{ member.initials }}
            </div>
            <h4
              class="text-sm font-semibold mb-1"
              style="color: var(--color-text-primary);"
            >
              {{ member.name }}
            </h4>
            <p
              class="text-xs"
              style="color: var(--color-text-tertiary);"
            >
              {{ member.role }} // Lv.{{ member.level }}
            </p>
          </div>
        </div>
      </section>

      <!-- Tech Stack -->
      <BaseCard
        padding="md"
        class="mb-8"
      >
        <div class="flex items-center gap-3 mb-4">
          <span
            class="w-1 h-5 rounded-sm"
            style="background: var(--color-state-success);"
          />
          <h2
            class="text-lg font-semibold"
            style="font-family: var(--font-display); color: var(--color-text-primary);"
          >
             {{ $t('about.sections.tech') }}
          </h2>
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-3">
          <div
            v-for="tech in techStack"
            :key="tech"
            class="rounded-md p-3 text-center"
            style="background: var(--color-bg-primary); border: 1px solid var(--color-border-subtle);"
          >
            <div
              class="text-xs font-semibold"
              style="color: var(--color-primary); font-family: var(--font-mono);"
            >
              {{ tech }}
            </div>
          </div>
        </div>
      </BaseCard>

      <!-- Contact -->
      <BaseCard padding="md">
        <div class="flex items-center gap-3 mb-4">
          <span
            class="w-1 h-5 rounded-sm"
            style="background: var(--color-state-warning);"
          />
          <h2
            class="text-lg font-semibold"
            style="font-family: var(--font-display); color: var(--color-text-primary);"
          >
             {{ $t('about.sections.contact') }}
          </h2>
        </div>
        <div class="flex items-center gap-4">
          <button
            class="w-10 h-10 rounded-md flex items-center justify-center transition-colors duration-150"
            style="background: var(--color-bg-primary); border: 1px solid var(--color-border-subtle);"
          >
            <Twitter class="w-4 h-4" style="color: var(--color-text-secondary);" />
          </button>
          <button
            class="w-10 h-10 rounded-md flex items-center justify-center transition-colors duration-150"
            style="background: var(--color-bg-primary); border: 1px solid var(--color-border-subtle);"
          >
            <Github class="w-4 h-4" style="color: var(--color-text-secondary);" />
          </button>
          <button
            class="w-10 h-10 rounded-md flex items-center justify-center transition-colors duration-150"
            style="background: var(--color-bg-primary); border: 1px solid var(--color-border-subtle);"
          >
            <Mail class="w-4 h-4" style="color: var(--color-text-secondary);" />
          </button>
        </div>
      </BaseCard>
    </div>

    <AppFooter />
  </div>
</template>