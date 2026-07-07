<template>
  <BaseCard padding="md">
    <div class="flex items-center gap-3 cursor-pointer" @click="editing = !editing">
      <div class="w-8 h-8 rounded-md flex items-center justify-center"
        :style="{ background: 'var(--color-primary50)', border: '1px solid var(--color-border)' }">
        <User class="w-4 h-4" :style="{ color: 'var(--color-primary)' }" />
      </div>
      <div class="flex-1 min-w-0">
        <h3 class="text-sm font-semibold" :style="{ color: 'var(--color-textPrimary)' }">{{ t('settings.profileEdit.title') }}</h3>
        <p class="text-xs mt-0.5" :style="{ color: 'var(--color-textTertiary)' }">{{ t('settings.profileEdit.desc') }}</p>
      </div>
      <ChevronRight class="w-4 h-4" :style="{ color: 'var(--color-textTertiary)', transform: editing ? 'rotate(90deg)' : 'rotate(0deg)', transition: 'transform 0.2s' }" />
    </div>

    <div v-if="editing" class="mt-5 ml-11 space-y-3">
      <div>
        <label class="block text-xs mb-1.5" :style="{ color: 'var(--color-textTertiary)', fontFamily: 'var(--font-mono)' }">{{ t('settings.profileEdit.nickname') }}</label>
        <input v-model="localName" type="text" class="w-full px-3 py-2 text-sm outline-none transition-all duration-150"
          :style="{ background: 'var(--color-bgPrimary)', border: '1px solid var(--color-border)', color: 'var(--color-textPrimary)', borderRadius: 'var(--radius-sm)' }" />
      </div>
      <div>
        <label class="block text-xs mb-1.5" :style="{ color: 'var(--color-textTertiary)', fontFamily: 'var(--font-mono)' }">{{ t('settings.profileEdit.fieldTitle') }}</label>
        <input v-model="localTitle" type="text" class="w-full px-3 py-2 text-sm outline-none transition-all duration-150"
          :style="{ background: 'var(--color-bgPrimary)', border: '1px solid var(--color-border)', color: 'var(--color-textPrimary)', borderRadius: 'var(--radius-sm)' }" />
      </div>
      <div>
        <label class="block text-xs mb-1.5" :style="{ color: 'var(--color-textTertiary)', fontFamily: 'var(--font-mono)' }">{{ t('settings.profileEdit.tags') }}</label>
        <input v-model="localTags" type="text" class="w-full px-3 py-2 text-sm outline-none transition-all duration-150"
          :style="{ background: 'var(--color-bgPrimary)', border: '1px solid var(--color-border)', color: 'var(--color-textPrimary)', borderRadius: 'var(--radius-sm)' }" />
      </div>
      <div v-if="message" class="text-xs" :style="{ color: msgColor }">{{ message }}</div>
      <div class="flex gap-2 pt-1">
        <BaseButton size="sm" variant="primary" :disabled="saving" @click="handleSave">
          <template v-if="saving">{{ t('common.loading') }}</template>
          <template v-else><Check class="w-3 h-3" /> {{ t('common.save') }}</template>
        </BaseButton>
        <BaseButton size="sm" variant="ghost" @click="resetForm">
          <X class="w-3 h-3" /> {{ t('common.cancel') }}
        </BaseButton>
      </div>
    </div>
  </BaseCard>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { User, ChevronRight, Check, X } from 'lucide-vue-next'
import BaseCard from '@/components/base/BaseCard.vue'
import BaseButton from '@/components/base/BaseButton.vue'
import { useUserStore } from '@/stores/userStore'
import { updateProfile } from '@/api/user'

const { t } = useI18n()
const userStore = useUserStore()

const editing = ref(false)
const saving = ref(false)
const message = ref('')
const localName = ref(userStore.playerName)
const localTitle = ref(userStore.title)
const localTags = ref((userStore.tags ?? []).join(', '))

const msgColor = computed(() => message.value.includes('成功') || message.value.includes('✓')
  ? 'var(--state-success)' : 'var(--state-error)')

async function handleSave() {
  if (!localName.value.trim()) { message.value = '不能为空'; return }
  saving.value = true
  message.value = ''
  try {
    const tags = localTags.value.split(',').map((t: string) => t.trim()).filter(Boolean)
    const res = await updateProfile({
      playerName: localName.value.trim(),
      title: localTitle.value.trim(),
      tags,
    })
    if (res.code === 0) {
      message.value = t('settings.profileEdit.saveSuccess') + ' ✓'
      editing.value = false
      await userStore.refreshUserData()
    } else {
      message.value = res.message || t('settings.profileEdit.saveFail')
    }
  } catch {
    message.value = t('common.errors.networkError')
  } finally {
    saving.value = false
  }
}

function resetForm() {
  editing.value = false
  localName.value = userStore.playerName
  localTitle.value = userStore.title
  localTags.value = (userStore.tags ?? []).join(', ')
  message.value = ''
}
</script>
