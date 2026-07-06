export type ArticleCategory = 'tech' | 'game' | 'hardware' | 'anime'

export interface Article {
  id: string
  title: string
  summary: string
  content: string
  category: ArticleCategory
  author: string
  publishDate: string
  readTime: number
  xpReward: number
  imageUrl: string
  tags: string[]
  isFeatured: boolean
  timeAgo?: string
}

export interface ArticleState {
  articles: Article[]
  currentCategory: ArticleCategory | 'all'
  searchQuery: string
}