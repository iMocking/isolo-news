<script setup lang="ts">
import { Home } from 'lucide-vue-next'
import BaseButton from '@/components/base/BaseButton.vue'
</script>

<template>
  <div class="min-h-screen">
    <!-- Scanline overlay -->
    <div
      class="fixed inset-0 z-0 pointer-events-none opacity-[0.03]"
      style="background: repeating-linear-gradient(0deg, transparent, transparent 2px, rgba(255, 255, 255, 0.08) 2px, rgba(255, 255, 255, 0.08) 4px);"
    />

    <div class="relative z-10 min-h-screen flex items-center justify-center px-6">
      <div class="text-center max-w-md">
        <!-- HUD frame decoration -->
        <div class="flex items-center justify-center gap-3 mb-6">
          <span
            class="block w-12 h-px"
            style="background: var(--color-border);"
          />
          <span
            class="w-3 h-3 rounded-sm"
            style="background: var(--state-error); box-shadow: 0 0 8px rgba(255, 51, 102, 0.5);"
          />
          <span
            class="block w-12 h-px"
            style="background: var(--color-border);"
          />
        </div>

        <!-- 404 Title -->
        <h1
          class="mb-3"
          style="font-family: var(--font-display); font-size: clamp(48px, 8vw, 96px); font-weight: 700; line-height: 1.05; color: var(--color-primary); text-shadow: 0 0 40px rgba(0, 240, 255, 0.3); letter-spacing: -0.02em;"
        >
          ERROR 404
        </h1>
        <p
          class="text-lg mb-2"
          style="font-family: var(--font-display); color: var(--color-text-secondary);"
        >
          {{ $t('notFound.subtitle') }}
        </p>
        <p
          class="text-sm mb-10"
          style="color: var(--color-text-tertiary); line-height: 1.6;"
        >
          {{ $t('notFound.message') }}
        </p>

        <!-- Glitch decoration -->
        <div
          class="mb-10 p-4"
          style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);"
        >
          <div
            class="text-xs space-y-1"
            style="font-family: var(--font-mono); color: var(--color-text-tertiary);"
          >
            <p>
              &gt; SIGNAL_STATUS: <span style="color: var(--color-state-error);">DISCONNECTED</span>
            </p>
            <p>
              &gt; ROUTE_TRACE: <span style="color: var(--color-state-warning);">UNKNOWN</span>
            </p>
            <p>
              &gt; TARGET_NODE: <span style="color: var(--color-text-tertiary);">NULL</span>
            </p>
            <p class="animate-pulse">
              &gt; RETRYING...
            </p>
          </div>
        </div>

        <!-- Return button -->
        <BaseButton
          variant="primary"
          size="lg"
        >
          <Home class="w-4 h-4" />
          RETURN TO BASE
        </BaseButton>
      </div>
    </div>
  </div>
</template>