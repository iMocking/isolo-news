export interface Achievement {
  id: string
  name: string
  description: string
  icon: string
  unlockedAt?: string
}

export interface UserState {
  playerName: string
  level: number
  xp: number
  maxXp: number
  achievements: Achievement[]
  readArticles: number
  loginDays: number
  rank: number
  title: string
}

export interface Quest {
  id: string
  title: string
  description: string
  xpReward: number
  progress: number
  target: number
  status: 'not_started' | 'in_progress' | 'completed'
  icon?: any
  statusText?: string
}