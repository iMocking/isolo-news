<script setup lang="ts">
import BaseButton from '@/components/base/BaseButton.vue'

defineProps({
  email: { type: String, required: true },
  password: { type: String, required: true },
  rememberMe: { type: Boolean, required: true },
  loginError: { type: String, required: true },
  isSubmitting: { type: Boolean, required: true },
  loginConfig: { type: Object, required: true },
  theme: { type: String, required: true },
})

defineEmits(['update:email', 'update:password', 'update:rememberMe', 'submit'])
</script>

<template>
  <form
    class="space-y-5"
    @submit.prevent="$emit('submit')"
  >
    <!-- 错误提示 -->
    <div
      v-if="loginError"
      class="text-xs font-mono px-3 py-2 rounded"
      :style="{
        color: '#ff4444',
        background: 'rgba(255,68,68,0.08)',
        border: '1px solid rgba(255,68,68,0.2)'
      }"
    >
      {{ loginError }}
    </div>

    <!-- Email -->
    <div>
      <label
        class="block text-xs font-mono mb-2 tracking-wider"
        :style="{ color: theme === 'comiket' ? '#6b6158' : '#8892a8' }"
      >
        {{ $t('login.form.username').toUpperCase() }}
      </label>
      <input
        :value="email"
        type="email"
        :placeholder="$t('login.form.username')"
        class="w-full px-4 py-3 text-sm outline-none transition-all duration-150"
        :style="{
          background: loginConfig.inputBg,
          border: `1px solid ${loginConfig.inputBorder}`,
          color: theme === 'comiket' ? '#2a2520' : '#e8edf5',
          borderRadius: theme === 'ironcore' ? '4px' : '8px',
          fontFamily: 'var(--font-mono)',
          caretColor: 'var(--color-primary)'
        }"
        @input="$emit('update:email', ($event.target as HTMLInputElement).value)"
        @focus="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputFocusBorder; ($event.target as HTMLInputElement).style.boxShadow = loginConfig.inputFocusShadow;"
        @blur="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputBorder; ($event.target as HTMLInputElement).style.boxShadow = 'none';"
      >
    </div>

    <!-- Password -->
    <div>
      <label
        class="block text-xs font-mono mb-2 tracking-wider"
        :style="{ color: theme === 'comiket' ? '#6b6158' : '#8892a8' }"
      >
        {{ $t('login.form.password').toUpperCase() }}
      </label>
      <input
        :value="password"
        type="password"
        :placeholder="$t('login.form.password')"
        class="w-full px-4 py-3 text-sm outline-none transition-all duration-150"
        :style="{
          background: loginConfig.inputBg,
          border: `1px solid ${loginConfig.inputBorder}`,
          color: theme === 'comiket' ? '#2a2520' : '#e8edf5',
          borderRadius: theme === 'ironcore' ? '4px' : '8px',
          fontFamily: 'var(--font-mono)',
          caretColor: 'var(--color-primary)'
        }"
        @input="$emit('update:password', ($event.target as HTMLInputElement).value)"
        @focus="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputFocusBorder; ($event.target as HTMLInputElement).style.boxShadow = loginConfig.inputFocusShadow;"
        @blur="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputBorder; ($event.target as HTMLInputElement).style.boxShadow = 'none';"
      >
    </div>

    <!-- Remember + Forgot -->
    <div class="flex items-center justify-between">
      <label class="flex items-center gap-2 cursor-pointer">
        <input
          :checked="rememberMe"
          type="checkbox"
          class="w-4 h-4"
          :style="{ accentColor: 'var(--color-primary)' }"
          @change="$emit('update:rememberMe', ($event.target as HTMLInputElement).checked)"
        >
        <span
          class="text-xs"
          :style="{ color: theme === 'comiket' ? '#6b6158' : '#8892a8' }"
        >
          {{ $t('login.form.remember') }}
        </span>
      </label>
      <a
        href="#"
        class="text-xs transition-colors duration-150 hover:text-[var(--color-primary)]"
        :style="{ color: theme === 'comiket' ? '#9e958b' : '#5a6478' }"
      >
        {{ $t('login.form.forgot') }}
      </a>
    </div>

    <!-- 登录按钮 -->
    <BaseButton
      variant="primary"
      size="lg"
      full-width
      :disabled="isSubmitting"
    >
      <template v-if="isSubmitting">连接中...</template>
      <template v-else>{{ $t('login.buttons.' + theme) }}</template>
    </BaseButton>
  </form>
</template>
