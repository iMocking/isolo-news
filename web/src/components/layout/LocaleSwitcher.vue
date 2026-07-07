<template>
  <!-- 语言切换：点击 Globe 图标呼出下拉菜单 -->
  <div class="relative" ref="dropdownRef">
    <!-- 触发按钮：Globe 图标 -->
    <button
      class="flex items-center justify-center w-8 h-8 rounded-lg transition-all duration-150"
      :style="triggerBtnStyle"
      @click.stop="toggleDropdown"
      :title="toggleTitle"
    >
      <Globe class="w-4 h-4" />
    </button>

    <!-- 下拉菜单面板 -->
    <Transition name="dropdown">
      <div
        v-if="isOpen"
        class="absolute top-10 right-0 min-w-[130px] z-50 rounded-lg overflow-hidden shadow-xl"
        :style="menuPanelStyle"
      >
        <div class="py-1">
          <button
            v-for="option in localeOptions"
            :key="option.value"
            class="w-full flex items-center gap-3 px-4 py-2.5 text-sm transition-colors duration-100"
            :style="getItemStyle(option.value)"
            @click.stop="handleSelect(option.value)"
            @mouseenter="onItemEnter"
            @mouseleave="onItemLeave"
          >
            <span>{{ option.label }}</span>
            <Check
              v-if="option.value === currentLocale"
              class="w-3.5 h-3.5 ml-auto"
              :style="{ color: 'var(--color-primary)' }"
            />
          </button>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
/**
 * 语言切换下拉菜单组件
 * 点击 Globe 图标展开下拉菜单，选择中/英文切换
 * 适配 NEXUS / COMIKET / IRONCORE 三主题色调
 */
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { setLocale } from '@/i18n'
import { Globe, Check } from 'lucide-vue-next'

const { locale } = useI18n()

/** 下拉菜单开关状态 */
const isOpen = ref(false)
/** 容器 DOM 引用，用于判断点击外部区域 */
const dropdownRef = ref<HTMLElement | null>(null)

/** 语言选项列表 */
const localeOptions = [
  { value: 'zh-CN', label: '中文' },
  { value: 'en-US', label: 'English' }
]

/** 当前语言代码 */
const currentLocale = computed(() => locale.value)

/** 触发按钮的 tooltip 文字 */
const toggleTitle = computed(() =>
  locale.value === 'zh-CN' ? '切换语言' : 'Switch Language'
)

/** 触发按钮样式：打开时高亮，关闭时半透明 */
const triggerBtnStyle = computed(() => ({
  background: isOpen.value ? 'var(--color-primary50)' : 'transparent',
  border: `1px solid ${isOpen.value ? 'var(--color-primary)' : 'var(--color-borderSubtle)'}`,
  color: isOpen.value ? 'var(--color-primary)' : 'var(--color-textTertiary)',
  cursor: 'pointer'
}))

/** 下拉菜单面板样式：使用 CSS 变量适配三主题 */
const menuPanelStyle = computed(() => ({
  background: 'var(--color-bgCard)',
  border: '1px solid var(--color-border)',
  backdropFilter: 'blur(12px)' as const
}))

/**
 * 获取菜单项的样式
 * @param val 选项值
 * @returns 样式对象：选中态使用主色，否则使用次要文字色
 */
const getItemStyle = (val: string) => {
  const isActive = val === currentLocale.value
  return {
    color: isActive ? 'var(--color-primary)' : 'var(--color-textSecondary)',
    fontFamily: 'var(--font-body)'
  }
}

/** 菜单项鼠标悬停恢复默认色 */
const onItemEnter = (e: Event) => {
  const target = e.currentTarget as HTMLElement
  target.style.background = 'var(--color-primary50)'
}

/** 菜单项鼠标离开恢复透明 */
const onItemLeave = (e: Event) => {
  const target = e.currentTarget as HTMLElement
  target.style.background = 'transparent'
}

/** 切换下拉菜单的打开/关闭 */
const toggleDropdown = () => {
  isOpen.value = !isOpen.value
}

/**
 * 选中语言选项
 * @param val 目标语言代码
 */
const handleSelect = (val: string) => {
  setLocale(val as 'zh-CN' | 'en-US')
  isOpen.value = false
}

/**
 * 点击文档其他区域时关闭菜单
 */
const handleDocumentClick = (e: MouseEvent) => {
  if (dropdownRef.value && !dropdownRef.value.contains(e.target as Node)) {
    isOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleDocumentClick)
})

onUnmounted(() => {
  document.removeEventListener('click', handleDocumentClick)
})
</script>

<style scoped>
/* 下拉菜单展开/收起过渡动画 */
.dropdown-enter-active,
.dropdown-leave-active {
  transition: opacity 0.15s ease, transform 0.15s ease;
}
.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-6px);
}
</style>
