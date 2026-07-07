<template>
  <BaseCard padding="md">
    <div class="flex items-center gap-3 mb-5">
      <div class="w-8 h-8 rounded-md flex items-center justify-center"
        :style="{ background: 'var(--color-bgTertiary)', border: '1px solid var(--color-border)' }">
        <Database class="w-4 h-4" :style="{ color: 'var(--color-textTertiary)' }" />
      </div>
      <h3 class="text-sm font-semibold" :style="{ color: 'var(--color-textPrimary)' }">
        {{ t('settings.system.info') }}
      </h3>
      <span class="text-[10px] ml-auto font-mono" :style="{ color: 'var(--color-textTertiary)' }">SYS.INFO</span>
    </div>
    <div class="grid grid-cols-2 gap-3 ml-11">
      <div v-for="info in resolvedItems" :key="info.label" class="rounded-md p-3"
        :style="{ background: 'var(--color-bgPrimary)', border: '1px solid var(--color-borderSubtle)' }">
        <div class="flex items-center gap-2 mb-1">
          <component :is="info.icon" class="w-3 h-3" :style="{ color: 'var(--color-textTertiary)' }" />
          <span class="text-[10px]" :style="{ color: 'var(--color-textTertiary)', fontFamily: 'var(--font-mono)' }">{{ info.label }}</span>
        </div>
        <span class="text-xs font-semibold" :style="{ color: 'var(--color-textPrimary)' }">{{ info.value }}</span>
      </div>
    </div>
  </BaseCard>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { Monitor, Globe, Database, HardDrive } from 'lucide-vue-next'
import BaseCard from '@/components/base/BaseCard.vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const iconMap: Record<string, any> = {
  monitor: Monitor,
  globe: Globe,
  database: Database,
  'hard-drive': HardDrive,
}

const props = defineProps<{
  items: Array<{ icon: string; label: string; value: string }>
}>()

const resolvedItems = computed(() =>
  props.items.map((item) => ({ ...item, icon: iconMap[item.icon] || Monitor }))
)
</script>
