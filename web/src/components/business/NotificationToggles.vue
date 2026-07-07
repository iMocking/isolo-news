<template>
  <div class="space-y-3">
    <label
      v-for="item in items"
      :key="item.key"
      class="flex items-center justify-between cursor-pointer"
    >
      <span class="text-sm" :style="{ color: 'var(--color-textSecondary)' }">
        {{ item.label }}
      </span>
      <div
        class="w-10 h-5 rounded-full relative cursor-pointer"
        :style="{ background: modelValue[item.key] ? 'var(--color-primary)' : 'var(--color-bgPrimary)' }"
        @click="$emit('update:modelValue', { ...modelValue, [item.key]: !modelValue[item.key] })"
      >
        <div
          class="absolute w-4 h-4 rounded-full transition-all duration-200"
          :style="{
            background: modelValue[item.key] ? 'white' : 'var(--color-textTertiary)',
            right: modelValue[item.key] ? '2px' : 'auto',
            left: modelValue[item.key] ? 'auto' : '2px',
            top: '2px'
          }"
        />
      </div>
    </label>
  </div>
</template>

<script setup lang="ts">
interface ToggleItem {
  key: string
  label: string
}

defineProps<{
  modelValue: Record<string, boolean>
  items: ToggleItem[]
}>()

defineEmits<{
  'update:modelValue': [value: Record<string, boolean>]
}>()
</script>
