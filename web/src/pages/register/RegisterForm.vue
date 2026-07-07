<script setup lang="ts">
import { ref, watch, onUnmounted } from 'vue'
import BaseButton from '@/components/base/BaseButton.vue'

const props = defineProps({
  registerEmail: { type: String, required: true },
  registerPassword: { type: String, required: true },
  confirmPassword: { type: String, required: true },
  playerName: { type: String, required: true },
  registerError: { type: String, required: true },
  isSubmitting: { type: Boolean, required: true },
  loginConfig: { type: Object, required: true },
  theme: { type: String, required: true },
  emailCode: { type: String, default: '' },
  codeSent: { type: Boolean, default: false },
  sendCodeError: { type: String, default: '' },
})

const emit = defineEmits([
  'update:registerEmail',
  'update:registerPassword',
  'update:confirmPassword',
  'update:playerName',
  'update:emailCode',
  'submit',
  'sendCode',
])

const countdown = ref(0)
let countdownTimer: number | undefined

/** 检测到 codeSent 变为 true 时启动 60s 倒计时 */
watch(() => props.codeSent, (val) => {
  if (val) {
    countdown.value = 60
    countdownTimer = window.setInterval(() => {
      countdown.value--
      if (countdown.value <= 0) {
        if (countdownTimer) clearInterval(countdownTimer)
      }
    }, 1000)
  }
})

/** 邮箱改变时重置倒计时（旧验证码失效） */
watch(() => props.registerEmail, () => {
  if (countdownTimer) clearInterval(countdownTimer)
  countdown.value = 0
})

onUnmounted(() => {
  if (countdownTimer) clearInterval(countdownTimer)
})
</script>

<template>
  <form
    class="space-y-5"
    @submit.prevent="$emit('submit')"
  >
    <!-- 错误提示 -->
    <div
      v-if="registerError"
      class="text-xs font-mono px-3 py-2 rounded"
      :style="{
        color: '#ff4444',
        background: 'rgba(255,68,68,0.08)',
        border: '1px solid rgba(255,68,68,0.2)'
      }"
    >
      {{ registerError }}
    </div>

    <!-- 昵称 -->
    <div>
      <label
        class="block text-xs font-mono mb-2 tracking-wider"
        :style="{ color: theme === 'comiket' ? '#6b6158' : '#8892a8' }"
      >
        昵称
      </label>
      <input
        :value="playerName"
        type="text"
        placeholder="你的游戏代号"
        class="w-full px-4 py-3 text-sm outline-none transition-all duration-150"
        :style="{
          background: loginConfig.inputBg,
          border: `1px solid ${loginConfig.inputBorder}`,
          color: theme === 'comiket' ? '#2a2520' : '#e8edf5',
          borderRadius: theme === 'ironcore' ? '4px' : '8px',
          fontFamily: 'var(--font-mono)',
          caretColor: 'var(--color-primary)'
        }"
        @input="$emit('update:playerName', ($event.target as HTMLInputElement).value)"
        @focus="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputFocusBorder; ($event.target as HTMLInputElement).style.boxShadow = loginConfig.inputFocusShadow;"
        @blur="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputBorder; ($event.target as HTMLInputElement).style.boxShadow = 'none';"
      >
    </div>

    <!-- 邮箱 + 发送验证码按钮 -->
    <div>
      <label
        class="block text-xs font-mono mb-2 tracking-wider"
        :style="{ color: theme === 'comiket' ? '#6b6158' : '#8892a8' }"
      >
        邮箱
      </label>
      <div class="flex gap-2">
        <input
          :value="registerEmail"
          type="email"
          placeholder="your@email.com"
          class="flex-1 px-4 py-3 text-sm outline-none transition-all duration-150"
          :style="{
            background: loginConfig.inputBg,
            border: `1px solid ${loginConfig.inputBorder}`,
            color: theme === 'comiket' ? '#2a2520' : '#e8edf5',
            borderRadius: theme === 'ironcore' ? '4px' : '8px',
            fontFamily: 'var(--font-mono)',
            caretColor: 'var(--color-primary)'
          }"
          @input="$emit('update:registerEmail', ($event.target as HTMLInputElement).value)"
          @focus="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputFocusBorder; ($event.target as HTMLInputElement).style.boxShadow = loginConfig.inputFocusShadow;"
          @blur="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputBorder; ($event.target as HTMLInputElement).style.boxShadow = 'none';"
        >
        <button
          type="button"
          :disabled="countdown > 0 || !registerEmail || isSubmitting"
          class="shrink-0 px-4 py-3 text-xs font-mono font-semibold tracking-wider transition-all duration-150 whitespace-nowrap"
          :style="{
            background: countdown > 0 || !registerEmail ? 'var(--color-bg-tertiary)' : 'var(--color-primary)',
            color: countdown > 0 || !registerEmail ? 'var(--color-text-tertiary)' : '#ffffff',
            border: 'none',
            borderRadius: theme === 'ironcore' ? '4px' : '8px',
            opacity: countdown > 0 || !registerEmail ? 0.6 : 1,
            cursor: countdown > 0 || !registerEmail ? 'not-allowed' : 'pointer',
          }"
          @click="$emit('sendCode')"
        >
          {{ countdown > 0 ? `${countdown}s` : '发送验证码' }}
        </button>
      </div>
      <!-- 发送验证码错误提示 -->
      <p
        v-if="sendCodeError"
        class="text-xs font-mono mt-1"
        :style="{ color: '#ff4444' }"
      >
        {{ sendCodeError }}
      </p>
    </div>

    <!-- 验证码（始终可见） -->
    <div>
      <label
        class="block text-xs font-mono mb-2 tracking-wider"
        :style="{ color: theme === 'comiket' ? '#6b6158' : '#8892a8' }"
      >
        验证码
      </label>
      <input
        :value="emailCode"
        type="text"
        placeholder="请输入邮箱收到的验证码"
        maxlength="6"
        class="w-full px-4 py-3 text-sm outline-none transition-all duration-150"
        :style="{
          background: loginConfig.inputBg,
          border: `1px solid ${loginConfig.inputBorder}`,
          color: theme === 'comiket' ? '#2a2520' : '#e8edf5',
          borderRadius: theme === 'ironcore' ? '4px' : '8px',
          fontFamily: 'var(--font-mono)',
          caretColor: 'var(--color-primary)'
        }"
        @input="$emit('update:emailCode', ($event.target as HTMLInputElement).value)"
        @focus="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputFocusBorder; ($event.target as HTMLInputElement).style.boxShadow = loginConfig.inputFocusShadow;"
        @blur="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputBorder; ($event.target as HTMLInputElement).style.boxShadow = 'none';"
      >
    </div>

    <!-- 注册密码 -->
    <div>
      <label
        class="block text-xs font-mono mb-2 tracking-wider"
        :style="{ color: theme === 'comiket' ? '#6b6158' : '#8892a8' }"
      >
        密码
      </label>
      <input
        :value="registerPassword"
        type="password"
        placeholder="至少6位字符"
        class="w-full px-4 py-3 text-sm outline-none transition-all duration-150"
        :style="{
          background: loginConfig.inputBg,
          border: `1px solid ${loginConfig.inputBorder}`,
          color: theme === 'comiket' ? '#2a2520' : '#e8edf5',
          borderRadius: theme === 'ironcore' ? '4px' : '8px',
          fontFamily: 'var(--font-mono)',
          caretColor: 'var(--color-primary)'
        }"
        @input="$emit('update:registerPassword', ($event.target as HTMLInputElement).value)"
        @focus="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputFocusBorder; ($event.target as HTMLInputElement).style.boxShadow = loginConfig.inputFocusShadow;"
        @blur="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputBorder; ($event.target as HTMLInputElement).style.boxShadow = 'none';"
      >
    </div>

    <!-- 确认密码 -->
    <div>
      <label
        class="block text-xs font-mono mb-2 tracking-wider"
        :style="{ color: theme === 'comiket' ? '#6b6158' : '#8892a8' }"
      >
        确认密码
      </label>
      <input
        :value="confirmPassword"
        type="password"
        placeholder="再次输入密码"
        class="w-full px-4 py-3 text-sm outline-none transition-all duration-150"
        :style="{
          background: loginConfig.inputBg,
          border: `1px solid ${loginConfig.inputBorder}`,
          color: theme === 'comiket' ? '#2a2520' : '#e8edf5',
          borderRadius: theme === 'ironcore' ? '4px' : '8px',
          fontFamily: 'var(--font-mono)',
          caretColor: 'var(--color-primary)'
        }"
        @input="$emit('update:confirmPassword', ($event.target as HTMLInputElement).value)"
        @focus="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputFocusBorder; ($event.target as HTMLInputElement).style.boxShadow = loginConfig.inputFocusShadow;"
        @blur="($event.target as HTMLInputElement).style.borderColor = loginConfig.inputBorder; ($event.target as HTMLInputElement).style.boxShadow = 'none';"
      >
    </div>

    <!-- 注册按钮 -->
    <BaseButton
      variant="primary"
      size="lg"
      full-width
      :disabled="isSubmitting"
    >
      <template v-if="isSubmitting">注册中...</template>
      <template v-else>创建账号</template>
    </BaseButton>
  </form>
</template>
