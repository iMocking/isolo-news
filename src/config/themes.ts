import type { ThemeConfig } from '@/types'

export const nexusTheme: ThemeConfig = {
  name: 'nexus',
  displayName: 'NEXUS',
  colors: {
    primary: '#00f0ff',
    primaryLight: '#66f7ff',
    primaryDark: '#00b8c5',
    primary50: 'rgba(0, 240, 255, 0.08)',
    primary100: 'rgba(0, 240, 255, 0.15)',
    primary200: 'rgba(0, 240, 255, 0.25)',
    secondary: '#ff2d95',
    secondaryLight: '#ff6db5',
    secondaryDark: '#cc1f78',
    secondary50: 'rgba(255, 45, 149, 0.08)',
    secondary100: 'rgba(255, 45, 149, 0.15)',
    bgPrimary: '#0a0e1a',
    bgSecondary: '#111827',
    bgTertiary: '#1a2035',
    bgCard: '#151c2e',
    bgElevated: '#1e2640',
    border: 'rgba(0, 240, 255, 0.15)',
    borderSubtle: 'rgba(255, 255, 255, 0.06)',
    textPrimary: '#e8edf5',
    textSecondary: '#8892a8',
    textTertiary: '#5a6478',
    textInverse: '#0a0e1a',
    stateSuccess: '#00ff88',
    stateWarning: '#ffaa00',
    stateError: '#ff3366',
    stateInfo: '#00aaff'
  },
  fonts: {
    display: "'Orbitron', 'Noto Sans SC', sans-serif",
    body: "'Inter', 'Noto Sans SC', sans-serif",
    mono: "'JetBrains Mono', 'Fira Code', monospace"
  },
  radius: {
    sm: '4px',
    md: '8px',
    lg: '12px',
    full: '9999px'
  }
}

export const comiketTheme: ThemeConfig = {
  name: 'comiket',
  displayName: 'COMIKET',
  colors: {
    primary: '#ff6b2b',
    primaryLight: '#ff9a66',
    primaryDark: '#cc5522',
    primary50: 'rgba(255, 107, 43, 0.08)',
    primary100: 'rgba(255, 107, 43, 0.15)',
    primary200: 'rgba(255, 107, 43, 0.25)',
    secondary: '#3b9eff',
    secondaryLight: '#70b8ff',
    secondaryDark: '#2e7ecc',
    secondary50: 'rgba(59, 158, 255, 0.08)',
    secondary100: 'rgba(59, 158, 255, 0.15)',
    bgPrimary: '#faf8f5',
    bgSecondary: '#f2eee8',
    bgTertiary: '#ebe6de',
    bgCard: '#ffffff',
    bgElevated: '#ffffff',
    border: 'rgba(0, 0, 0, 0.1)',
    borderSubtle: 'rgba(0, 0, 0, 0.05)',
    textPrimary: '#2a2520',
    textSecondary: '#6b6158',
    textTertiary: '#9e958b',
    textInverse: '#faf8f5',
    stateSuccess: '#22c55e',
    stateWarning: '#f59e0b',
    stateError: '#ef4444',
    stateInfo: '#3b82f6'
  },
  fonts: {
    display: "'Zen Maru Gothic', 'Noto Sans SC', sans-serif",
    body: "'Inter', 'Noto Sans SC', sans-serif",
    mono: "'JetBrains Mono', 'Fira Code', monospace"
  },
  radius: {
    sm: '4px',
    md: '10px',
    lg: '16px',
    full: '9999px'
  }
}

export const ironcoreTheme: ThemeConfig = {
  name: 'ironcore',
  displayName: 'IRONCORE',
  colors: {
    primary: '#f0a030',
    primaryLight: '#f4be6a',
    primaryDark: '#c08020',
    primary50: 'rgba(240, 160, 48, 0.08)',
    primary100: 'rgba(240, 160, 48, 0.15)',
    primary200: 'rgba(240, 160, 48, 0.25)',
    secondary: '#4ecdc4',
    secondaryLight: '#7eddd6',
    secondaryDark: '#3aa39c',
    secondary50: 'rgba(78, 205, 196, 0.08)',
    secondary100: 'rgba(78, 205, 196, 0.15)',
    bgPrimary: '#121418',
    bgSecondary: '#1a1d24',
    bgTertiary: '#22262e',
    bgCard: '#1e2128',
    bgElevated: '#282c35',
    border: 'rgba(240, 160, 48, 0.12)',
    borderSubtle: 'rgba(255, 255, 255, 0.06)',
    textPrimary: '#d8dce5',
    textSecondary: '#808898',
    textTertiary: '#555c6c',
    textInverse: '#121418',
    stateSuccess: '#4ecdc4',
    stateWarning: '#f0a030',
    stateError: '#e74c5e',
    stateInfo: '#5b8def'
  },
  fonts: {
    display: "'Rajdhani', 'Noto Sans SC', sans-serif",
    body: "'Inter', 'Noto Sans SC', sans-serif",
    mono: "'JetBrains Mono', 'Fira Code', monospace"
  },
  radius: {
    sm: '2px',
    md: '6px',
    lg: '10px',
    full: '9999px'
  }
}

export const themes = {
  nexus: nexusTheme,
  comiket: comiketTheme,
  ironcore: ironcoreTheme
}