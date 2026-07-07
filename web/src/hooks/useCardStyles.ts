import { computed } from 'vue'
import { useThemeStore } from '@/stores/themeStore'

export type CardStyle = Record<string, string | undefined>

interface InternalCardStyle extends CardStyle {
  hoverBorder?: string
  hoverShadow?: string
  hoverBorderTop?: string
  hoverBorderLeft?: string
  hoverBorderRight?: string
  hoverBorderBottom?: string
}

interface InternalThemeCardStyles {
  default: InternalCardStyle
  elevated: InternalCardStyle
  outlined: InternalCardStyle
  panel: InternalCardStyle
  article: InternalCardStyle
  comment: InternalCardStyle
  stat: InternalCardStyle
}

export interface ThemeCardStyles {
  default: CardStyle
  elevated: CardStyle
  outlined: CardStyle
  panel: CardStyle
  article: CardStyle
  comment: CardStyle
  stat: CardStyle
}

export function useCardStyles() {
  const themeStore = useThemeStore()
  const theme = computed(() => themeStore.currentTheme)

  const nexusCardStyles: InternalThemeCardStyles = {
    default: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-border)',
      borderRadius: 'var(--radius-lg)',
      boxShadow: '0 0 0 rgba(0, 240, 255, 0), inset 0 1px 0 rgba(255, 255, 255, 0.02)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 0 20px rgba(0, 240, 255, 0.15), 0 0 40px rgba(0, 240, 255, 0.05), inset 0 0 20px rgba(0, 240, 255, 0.02)'
    },
    elevated: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-border)',
      borderRadius: 'var(--radius-lg)',
      boxShadow: '0 4px 20px rgba(0, 0, 0, 0.3), 0 0 20px rgba(0, 240, 255, 0.05), inset 0 1px 0 rgba(255, 255, 255, 0.03)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 8px 30px rgba(0, 0, 0, 0.4), 0 0 30px rgba(0, 240, 255, 0.15), inset 0 1px 0 rgba(255, 255, 255, 0.05)'
    },
    outlined: {
      background: 'rgba(21, 28, 46, 0.8)',
      border: '1px solid var(--color-border)',
      borderRadius: 'var(--radius-lg)',
      boxShadow: 'none',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 0 15px rgba(0, 240, 255, 0.1)'
    },
    panel: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-border)',
      borderTop: '1px solid var(--color-primary)',
      borderRadius: 'var(--radius-lg)',
      boxShadow: '0 4px 24px rgba(0, 0, 0, 0.3), 0 0 0 1px rgba(0, 240, 255, 0.03), inset 0 1px 0 rgba(255, 255, 255, 0.02)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverBorderTop: '2px solid var(--color-primary)',
      hoverShadow: '0 8px 32px rgba(0, 0, 0, 0.4), 0 0 20px rgba(0, 240, 255, 0.1)'
    },
    article: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-border)',
      borderRadius: 'var(--radius-lg)',
      boxShadow: '0 2px 12px rgba(0, 0, 0, 0.2), inset 0 1px 0 rgba(255, 255, 255, 0.02)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 0 20px rgba(0, 240, 255, 0.15), 0 4px 16px rgba(0, 0, 0, 0.3), inset 0 0 15px rgba(0, 240, 255, 0.03)'
    },
    comment: {
      background: 'rgba(21, 28, 46, 0.6)',
      border: '1px solid var(--color-borderSubtle)',
      borderRadius: 'var(--radius-md)',
      boxShadow: 'none',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 0 15px rgba(0, 240, 255, 0.08)'
    },
    stat: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-borderSubtle)',
      borderRadius: 'var(--radius-md)',
      boxShadow: '0 1px 8px rgba(0, 0, 0, 0.2), inset 0 1px 0 rgba(255, 255, 255, 0.02)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 2px 12px rgba(0, 0, 0, 0.3), 0 0 10px rgba(0, 240, 255, 0.08)'
    }
  }

  const comiketCardStyles: InternalThemeCardStyles = {
    default: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-borderSubtle)',
      borderRadius: 'var(--radius-lg)',
      boxShadow: '0 2px 12px rgba(255, 107, 43, 0.06), inset 0 1px 0 rgba(255, 255, 255, 0.8)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 4px 20px rgba(255, 107, 43, 0.12), inset 0 1px 0 rgba(255, 255, 255, 0.9)'
    },
    elevated: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-border)',
      borderRadius: 'var(--radius-xl)',
      boxShadow: '0 6px 24px rgba(255, 107, 43, 0.08), 0 0 0 1px rgba(255, 107, 43, 0.05)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 10px 32px rgba(255, 107, 43, 0.15), 0 0 0 1px rgba(255, 107, 43, 0.1)'
    },
    outlined: {
      background: 'rgba(255, 255, 255, 0.6)',
      border: '1px dashed var(--color-border)',
      borderRadius: 'var(--radius-lg)',
      boxShadow: 'none',
      hoverBorder: '1px dashed var(--color-primary)',
      hoverShadow: '0 0 15px rgba(255, 107, 43, 0.1)'
    },
    panel: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-borderSubtle)',
      borderRadius: 'var(--radius-xl)',
      boxShadow: '0 4px 20px rgba(255, 107, 43, 0.06), inset 0 1px 0 rgba(255, 255, 255, 0.9)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 8px 28px rgba(255, 107, 43, 0.12)'
    },
    article: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-borderSubtle)',
      borderRadius: 'var(--radius-lg)',
      boxShadow: '0 2px 12px rgba(255, 107, 43, 0.05), inset 0 1px 0 rgba(255, 255, 255, 0.9)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 4px 20px rgba(255, 107, 43, 0.1), inset 0 1px 0 rgba(255, 255, 255, 1)'
    },
    comment: {
      background: 'rgba(255, 255, 255, 0.8)',
      border: '1px solid var(--color-borderSubtle)',
      borderRadius: 'var(--radius-lg)',
      boxShadow: '0 1px 8px rgba(255, 107, 43, 0.04)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 2px 12px rgba(255, 107, 43, 0.08)'
    },
    stat: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-borderSubtle)',
      borderRadius: 'var(--radius-lg)',
      boxShadow: '0 2px 8px rgba(255, 107, 43, 0.04), inset 0 1px 0 rgba(255, 255, 255, 0.9)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 4px 16px rgba(255, 107, 43, 0.08)'
    }
  }

  const ironcoreCardStyles: InternalThemeCardStyles = {
    default: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-border)',
      borderRadius: 'var(--radius-md)',
      boxShadow: '0 2px 8px rgba(0, 0, 0, 0.4), inset 0 1px 0 rgba(255, 255, 255, 0.02)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 4px 16px rgba(0, 0, 0, 0.5), 0 0 15px rgba(240, 160, 48, 0.1), inset 0 1px 0 rgba(255, 255, 255, 0.03)'
    },
    elevated: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-border)',
      borderRadius: 'var(--radius-md)',
      boxShadow: '0 4px 16px rgba(0, 0, 0, 0.5), 0 0 0 1px rgba(240, 160, 48, 0.05)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 6px 24px rgba(0, 0, 0, 0.6), 0 0 0 1px rgba(240, 160, 48, 0.1), 0 0 20px rgba(240, 160, 48, 0.1)'
    },
    outlined: {
      background: 'rgba(30, 33, 40, 0.9)',
      border: '1px solid var(--color-border)',
      borderRadius: 'var(--radius-md)',
      boxShadow: 'none',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 0 15px rgba(240, 160, 48, 0.08)'
    },
    panel: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-border)',
      borderRadius: 'var(--radius-sm)',
      boxShadow: '0 4px 16px rgba(0, 0, 0, 0.5), 0 0 0 1px rgba(240, 160, 48, 0.05), inset 0 1px 0 rgba(255, 255, 255, 0.02)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 6px 24px rgba(0, 0, 0, 0.6), 0 0 15px rgba(240, 160, 48, 0.1)'
    },
    article: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-border)',
      borderRadius: 'var(--radius-sm)',
      boxShadow: '0 2px 8px rgba(0, 0, 0, 0.4), inset 0 1px 0 rgba(255, 255, 255, 0.02)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 4px 16px rgba(0, 0, 0, 0.5), 0 0 10px rgba(240, 160, 48, 0.08)'
    },
    comment: {
      background: 'rgba(30, 33, 40, 0.8)',
      border: '1px solid var(--color-border)',
      borderRadius: 'var(--radius-sm)',
      boxShadow: '0 1px 6px rgba(0, 0, 0, 0.3)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 2px 12px rgba(0, 0, 0, 0.4), 0 0 8px rgba(240, 160, 48, 0.06)'
    },
    stat: {
      background: 'var(--color-bgCard)',
      border: '1px solid var(--color-border)',
      borderRadius: 'var(--radius-sm)',
      boxShadow: '0 2px 8px rgba(0, 0, 0, 0.4), inset 0 1px 0 rgba(255, 255, 255, 0.02)',
      hoverBorder: '1px solid var(--color-primary)',
      hoverShadow: '0 3px 12px rgba(0, 0, 0, 0.5), 0 0 8px rgba(240, 160, 48, 0.06)'
    }
  }

  const cardStyles = computed<InternalThemeCardStyles>(() => {
    const styles: Record<string, InternalThemeCardStyles> = {
      nexus: nexusCardStyles,
      comiket: comiketCardStyles,
      ironcore: ironcoreCardStyles
    }
    return styles[theme.value] || nexusCardStyles
  })

  const getCardStyle = (variant: keyof ThemeCardStyles = 'default', isHovered = false): CardStyle => {
    const baseStyle = cardStyles.value[variant]
    if (isHovered) {
      const hoverStyle: CardStyle = {
        ...baseStyle,
        border: baseStyle.hoverBorder || baseStyle.border,
        boxShadow: baseStyle.hoverShadow || baseStyle.boxShadow
      }
      if (baseStyle.hoverBorderTop) hoverStyle.borderTop = baseStyle.hoverBorderTop
      if (baseStyle.hoverBorderLeft) hoverStyle.borderLeft = baseStyle.hoverBorderLeft
      if (baseStyle.hoverBorderRight) hoverStyle.borderRight = baseStyle.hoverBorderRight
      if (baseStyle.hoverBorderBottom) hoverStyle.borderBottom = baseStyle.hoverBorderBottom
      return hoverStyle
    }
    return baseStyle
  }

  return {
    cardStyles,
    getCardStyle,
    nexusCardStyles,
    comiketCardStyles,
    ironcoreCardStyles
  }
}
