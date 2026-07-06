export type ThemeName = 'nexus' | 'comiket' | 'ironcore'

export interface ThemeColors {
  primary: string
  primaryLight: string
  primaryDark: string
  primary50: string
  primary100: string
  primary200: string
  secondary: string
  secondaryLight: string
  secondaryDark: string
  secondary50: string
  secondary100: string
  bgPrimary: string
  bgSecondary: string
  bgTertiary: string
  bgCard: string
  bgElevated: string
  border: string
  borderSubtle: string
  textPrimary: string
  textSecondary: string
  textTertiary: string
  textInverse: string
  stateSuccess: string
  stateWarning: string
  stateError: string
  stateInfo: string
}

export interface ThemeFonts {
  display: string
  body: string
  mono: string
}

export interface ThemeRadius {
  sm: string
  md: string
  lg: string
  full: string
}

export interface ThemeConfig {
  name: ThemeName
  displayName: string
  colors: ThemeColors
  fonts: ThemeFonts
  radius: ThemeRadius
}

export interface ThemeState {
  currentTheme: ThemeName
  themeConfig: ThemeConfig
}