import { computed, type ComputedRef } from 'vue'

export interface LoginThemeConfig {
  bgClass: string
  cardBg: string
  cardBorder: string
  cardShadow: string
  inputBg: string
  inputBorder: string
  inputFocusBorder: string
  inputFocusShadow: string
  scanlineOpacity: number
  glowColor1: string
  glowColor2: string
  cornerBracketColor: string
  logoTextShadow: string
  tabIndicatorShadow: string
  socialHoverBorder: string
  socialHoverColor: string
  showScanline: boolean
}

export function useLoginThemeConfig(theme: ComputedRef<string>): ComputedRef<LoginThemeConfig> {
  const configs: Record<string, LoginThemeConfig> = {
    nexus: {
      bgClass: 'bg-[#0a0e1a]',
      cardBg: '#151c2e',
      cardBorder: '1px solid rgba(0, 240, 255, 0.15)',
      cardShadow: '0 0 40px rgba(0, 240, 255, 0.15), 0 0 80px rgba(0, 240, 255, 0.05), inset 0 1px 0 rgba(255, 255, 255, 0.05)',
      inputBg: '#111827',
      inputBorder: 'rgba(255, 255, 255, 0.06)',
      inputFocusBorder: '#00f0ff',
      inputFocusShadow: '0 0 0 2px rgba(0, 240, 255, 0.1)',
      scanlineOpacity: 0.03,
      glowColor1: 'rgba(0, 240, 255, 0.08)',
      glowColor2: 'rgba(255, 45, 149, 0.06)',
      cornerBracketColor: 'rgba(0, 240, 255, 0.3)',
      logoTextShadow: '0 0 30px rgba(0, 240, 255, 0.4)',
      tabIndicatorShadow: '0 0 15px rgba(0, 240, 255, 0.5)',
      socialHoverBorder: '#00f0ff',
      socialHoverColor: '#00f0ff',
      showScanline: true
    },
    comiket: {
      bgClass: 'bg-[#faf8f5]',
      cardBg: '#ffffff',
      cardBorder: '1px solid rgba(0, 0, 0, 0.08)',
      cardShadow: '0 20px 60px rgba(255, 107, 43, 0.15), 0 10px 30px rgba(59, 158, 255, 0.1), inset 0 1px 0 rgba(255, 255, 255, 0.8)',
      inputBg: '#f8f6f2',
      inputBorder: 'rgba(0, 0, 0, 0.06)',
      inputFocusBorder: '#ff6b2b',
      inputFocusShadow: '0 0 0 2px rgba(255, 107, 43, 0.1)',
      scanlineOpacity: 0,
      glowColor1: 'rgba(255, 107, 43, 0.06)',
      glowColor2: 'rgba(59, 158, 255, 0.05)',
      cornerBracketColor: 'rgba(255, 107, 43, 0.3)',
      logoTextShadow: '0 0 20px rgba(255, 107, 43, 0.3)',
      tabIndicatorShadow: '0 0 10px rgba(255, 107, 43, 0.4)',
      socialHoverBorder: '#ff6b2b',
      socialHoverColor: '#ff6b2b',
      showScanline: false
    },
    ironcore: {
      bgClass: 'bg-[#121418]',
      cardBg: '#1e2128',
      cardBorder: '1px solid rgba(240, 160, 48, 0.1)',
      cardShadow: '0 0 40px rgba(240, 160, 48, 0.1), 0 0 80px rgba(240, 160, 48, 0.03), inset 0 1px 0 rgba(255, 255, 255, 0.03)',
      inputBg: '#1a1d24',
      inputBorder: 'rgba(255, 255, 255, 0.05)',
      inputFocusBorder: '#f0a030',
      inputFocusShadow: '0 0 0 2px rgba(240, 160, 48, 0.1)',
      scanlineOpacity: 0,
      glowColor1: 'rgba(240, 160, 48, 0.06)',
      glowColor2: 'rgba(78, 205, 196, 0.05)',
      cornerBracketColor: 'rgba(240, 160, 48, 0.3)',
      logoTextShadow: '0 0 20px rgba(240, 160, 48, 0.3)',
      tabIndicatorShadow: '0 0 10px rgba(240, 160, 48, 0.4)',
      socialHoverBorder: '#f0a030',
      socialHoverColor: '#f0a030',
      showScanline: false
    }
  }
  return computed(() => configs[theme.value] || configs.nexus)
}
