<template>
  <div class="min-h-screen">
    <!-- Scanline overlay for NEXUS theme -->
    <div v-if="themeStore.currentTheme === 'nexus'" class="fixed inset-0 pointer-events-none z-50 opacity-[0.03]" style="background: repeating-linear-gradient(0deg, transparent, transparent 2px, rgba(0,240,255,0.08) 2px, rgba(0,240,255,0.08) 4px);"></div>
    
    <AppNavigation />
    
    <!-- Page Header -->
    <section class="pt-24 pb-6 px-6">
      <div class="max-w-7xl mx-auto">
        <h1 class="text-2xl md:text-4xl tracking-wider" style="font-family: var(--font-display); color: var(--color-primary); word-break: keep-all;">任务列表 // MISSION SELECT</h1>
        <p class="mt-3 text-sm" style="font-family: var(--font-mono); color: var(--color-text-secondary);">共 2,847 条情报</p>
        
        <!-- Search bar -->
        <div class="mt-5 max-w-xl relative">
          <div class="flex items-center px-4 h-12" style="background: var(--color-bg-card); border: 1px solid var(--color-border); border-radius: var(--radius-md);">
            <span class="text-sm mr-3" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">&gt;_</span>
            <input 
              type="text" 
              v-model="searchQuery"
              placeholder="搜索情报关键词..." 
              class="flex-1 bg-transparent outline-none text-sm" 
              style="color: var(--color-text-primary); font-family: var(--font-mono); caret-color: var(--color-primary);"
            />
            <div class="w-0.5 h-5 animate-pulse" style="background: var(--color-primary);"></div>
          </div>
        </div>
      </div>
    </section>

    <!-- Category Filter Bar -->
    <section class="px-6 pb-4">
      <div class="max-w-7xl mx-auto">
        <div class="flex gap-2 overflow-x-auto no-scrollbar py-1">
          <button 
            v-for="cat in categories" 
            :key="cat.id"
            @click="currentCategory = cat.id"
            class="shrink-0 px-4 h-9 text-sm font-medium transition-all duration-150 whitespace-nowrap"
            :style="{
              background: currentCategory === cat.id ? 'var(--color-primary-200)' : 'transparent',
              border: currentCategory === cat.id ? '1px solid var(--color-primary)' : '1px solid var(--color-border-subtle)',
              color: currentCategory === cat.id ? 'var(--color-primary)' : 'var(--color-text-secondary)',
              borderRadius: 'var(--radius-md)',
              fontFamily: 'var(--font-mono)',
              boxShadow: currentCategory === cat.id ? '0 0 12px rgba(0,240,255,0.2)' : 'none'
            }"
          >
            {{ cat.name }} <span class="ml-1 text-xs" :style="{
              background: currentCategory === cat.id ? 'var(--color-primary)' : 'var(--color-bg-elevated)',
              color: currentCategory === cat.id ? 'var(--color-text-inverse)' : 'var(--color-text-tertiary)',
              padding: '1px 6px',
              borderRadius: 'var(--radius-full)'
            }">{{ cat.count }}</span>
          </button>
        </div>
      </div>
    </section>

    <!-- Main Content Area: List + Sidebar -->
    <section class="px-6 pb-12">
      <div class="max-w-7xl mx-auto flex gap-6">
        <!-- Article List -->
        <div class="flex-1 min-w-0 space-y-3">
          <article 
            v-for="(article, index) in filteredArticles" 
            :key="article.id"
            class="flex items-stretch gap-0 p-5 transition-all duration-150 hover:translate-y-[-1px] cursor-pointer"
            :style="{
              background: 'var(--color-bg-card)',
              border: '1px solid var(--color-border-subtle)',
              borderLeft: `3px solid ${getCategoryColor(article.category)}`,
              borderRadius: 'var(--radius-md)'
            }"
            @click="router.push(`/article/${article.id}`)"
          >
            <div class="shrink-0 w-12 flex items-center justify-center mr-5">
              <span class="text-2xl font-bold" :style="{ fontFamily: 'var(--font-display)', color: getCategoryColor(article.category) }">
                {{ String(index + 1).padStart(2, '0') }}
              </span>
            </div>
            <div class="flex-1 min-w-0">
              <div class="flex items-center gap-2 mb-2">
                <span class="px-2 py-0.5 text-xs font-medium whitespace-nowrap" :style="{
                  background: getCategoryBg(article.category),
                  color: getCategoryColor(article.category),
                  borderRadius: 'var(--radius-sm)',
                  fontFamily: 'var(--font-mono)'
                }">{{ getCategoryName(article.category) }}</span>
                <h2 class="text-base font-semibold truncate" style="color: var(--color-text-primary);">{{ article.title }}</h2>
              </div>
              <p class="text-sm line-clamp-2 mb-3" style="color: var(--color-text-secondary); line-height: 1.6;">{{ article.summary }}</p>
              <div class="flex items-center gap-3">
                <div class="w-5 h-5 flex items-center justify-center" style="background: var(--color-bg-elevated); border-radius: var(--radius-full);">
                  <span class="text-[10px] font-bold" :style="{ color: getCategoryColor(article.category) }">{{ article.author.charAt(0).toUpperCase() }}</span>
                </div>
                <span class="text-xs" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">{{ article.author }} · {{ article.publishDate }} · {{ article.readTime * 10 }} 评论</span>
              </div>
            </div>
            <div class="shrink-0 ml-5 flex flex-col items-center gap-3 justify-center">
              <div class="px-3 py-1.5 text-xs font-bold whitespace-nowrap" :style="{
                background: getCategoryBg(article.category),
                border: `1px solid ${getCategoryBgLight(article.category)}`,
                color: getCategoryColor(article.category),
                borderRadius: 'var(--radius-sm)',
                fontFamily: 'var(--font-display)'
              }">+{{ article.xpReward }} XP</div>
              <div class="flex gap-1">
                <span v-for="i in 5" :key="i" class="w-2 h-2" :style="{
                  background: i <= Math.min(5, Math.floor(article.xpReward / 100)) ? getCategoryColor(article.category) : 'var(--color-bg-elevated)',
                  borderRadius: 'var(--radius-full)'
                }"></span>
              </div>
              <button class="px-3 py-1.5 text-xs font-medium whitespace-nowrap transition-all duration-150 hover:opacity-80" :style="{
                background: getCategoryColor(article.category),
                color: 'var(--color-text-inverse)',
                borderRadius: 'var(--radius-sm)',
                fontFamily: 'var(--font-mono)'
              }">开始任务</button>
            </div>
          </article>

          <!-- Pagination -->
          <div class="flex items-center justify-center gap-2 pt-6">
            <button class="px-3 py-2 text-xs font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); color: var(--color-text-tertiary); border-radius: var(--radius-sm); font-family: var(--font-mono);">PREV</button>
            <button class="w-9 h-9 flex items-center justify-center text-sm font-medium whitespace-nowrap" style="background: var(--color-primary); color: var(--color-text-inverse); border-radius: var(--radius-sm); font-family: var(--font-display);">1</button>
            <button class="w-9 h-9 flex items-center justify-center text-sm font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); color: var(--color-text-secondary); border-radius: var(--radius-sm); font-family: var(--font-mono);">2</button>
            <button class="w-9 h-9 flex items-center justify-center text-sm font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); color: var(--color-text-secondary); border-radius: var(--radius-sm); font-family: var(--font-mono);">3</button>
            <span class="text-sm px-1" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">...</span>
            <button class="w-9 h-9 flex items-center justify-center text-sm font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); color: var(--color-text-secondary); border-radius: var(--radius-sm); font-family: var(--font-mono);">42</button>
            <button class="px-4 py-2 text-xs font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-primary-50); border: 1px solid var(--color-primary); color: var(--color-primary); border-radius: var(--radius-sm); font-family: var(--font-display);">NEXT STAGE →</button>
          </div>
        </div>

        <!-- Sidebar -->
        <aside class="hidden lg:block w-72 shrink-0 space-y-6">
          <!-- Filter Panel -->
          <div class="p-5" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);">
            <h3 class="text-sm font-semibold mb-4" style="font-family: var(--font-display); color: var(--color-primary);">筛选条件</h3>
            <div class="space-y-3">
              <div>
                <label class="block text-xs mb-1.5" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">时间范围</label>
                <select class="w-full px-3 h-9 text-sm outline-none" style="background: var(--color-bg-elevated); border: 1px solid var(--color-border-subtle); color: var(--color-text-secondary); border-radius: var(--radius-sm); font-family: var(--font-mono);">
                  <option>全部时间</option>
                  <option>24小时内</option>
                  <option>一周内</option>
                  <option>一个月内</option>
                </select>
              </div>
              <div>
                <label class="block text-xs mb-1.5" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">热度排序</label>
                <select class="w-full px-3 h-9 text-sm outline-none" style="background: var(--color-bg-elevated); border: 1px solid var(--color-border-subtle); color: var(--color-text-secondary); border-radius: var(--radius-sm); font-family: var(--font-mono);">
                  <option>最新发布</option>
                  <option>最多评论</option>
                  <option>最高热度</option>
                </select>
              </div>
              <div>
                <label class="block text-xs mb-1.5" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">来源</label>
                <select class="w-full px-3 h-9 text-sm outline-none" style="background: var(--color-bg-elevated); border: 1px solid var(--color-border-subtle); color: var(--color-text-secondary); border-radius: var(--radius-sm); font-family: var(--font-mono);">
                  <option>全部来源</option>
                  <option>官方渠道</option>
                  <option>社区投稿</option>
                  <option>行业媒体</option>
                </select>
              </div>
            </div>
          </div>

          <!-- Hot Tags -->
          <div class="p-5" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);">
            <h3 class="text-sm font-semibold mb-4" style="font-family: var(--font-display); color: var(--color-primary);">热门标签</h3>
            <div class="flex flex-wrap gap-2">
              <span v-for="tag in hotTags" :key="tag.name" class="px-2.5 py-1 text-xs whitespace-nowrap transition-all duration-150" :style="{
                background: tag.bg,
                border: tag.border,
                color: tag.color,
                borderRadius: 'var(--radius-sm)',
                fontFamily: 'var(--font-mono)'
              }">{{ tag.name }}</span>
            </div>
          </div>

          <!-- Weekly Leaderboard -->
          <div class="p-5" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);">
            <h3 class="text-sm font-semibold mb-4" style="font-family: var(--font-display); color: var(--color-primary);">本周排行</h3>
            <div class="space-y-3">
              <div v-for="(item, index) in leaderboard" :key="index" class="flex items-center gap-3 group cursor-pointer">
                <span class="w-5 h-5 flex items-center justify-center text-[10px] font-bold shrink-0" :style="{
                  background: getLeaderboardColor(index),
                  color: index < 3 ? 'var(--color-text-inverse)' : 'var(--color-text-tertiary)',
                  borderRadius: 'var(--radius-full)',
                  fontFamily: 'var(--font-display)'
                }">{{ index + 1 }}</span>
                <p class="text-sm truncate flex-1 min-w-0 transition-colors duration-150" style="color: var(--color-text-primary);">{{ item.title }}</p>
                <span class="text-[10px] shrink-0" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">{{ item.views }}</span>
              </div>
            </div>
          </div>
        </aside>
      </div>
    </section>

    <AppFooter />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import type { ArticleCategory } from '@/types'
import { useThemeStore } from '@/stores/themeStore'
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'

const router = useRouter()

const themeStore = useThemeStore()
const searchQuery = ref('')
const currentCategory = ref<string>('all')

const categories = [
  { id: 'all', name: '全部', count: 2847 },
  { id: 'tech', name: '科技', count: 1203 },
  { id: 'game', name: '游戏', count: 856 },
  { id: 'hardware', name: '硬件', count: 412 },
  { id: 'anime', name: '二次元', count: 376 }
]

const hotTags = [
  { name: 'AI', bg: 'var(--color-primary-50)', border: '1px solid var(--color-primary-100)', color: 'var(--color-primary)' },
  { name: 'VR', bg: 'var(--color-secondary-50)', border: '1px solid var(--color-secondary-100)', color: 'var(--color-secondary)' },
  { name: '显卡', bg: 'rgba(0,170,255,0.08)', border: '1px solid rgba(0,170,255,0.15)', color: 'var(--color-state-info)' },
  { name: '新番', bg: 'rgba(255,170,0,0.08)', border: '1px solid rgba(255,170,0,0.15)', color: 'var(--color-state-warning)' },
  { name: '机器人', bg: 'var(--color-primary-50)', border: '1px solid var(--color-primary-100)', color: 'var(--color-primary)' },
  { name: 'RPG', bg: 'var(--color-secondary-50)', border: '1px solid var(--color-secondary-100)', color: 'var(--color-secondary)' },
  { name: '芯片', bg: 'var(--color-primary-50)', border: '1px solid var(--color-primary-100)', color: 'var(--color-primary)' },
  { name: '漫画', bg: 'rgba(255,170,0,0.08)', border: '1px solid rgba(255,170,0,0.15)', color: 'var(--color-state-warning)' }
]

const leaderboard = [
  { title: '苹果Vision Pro 2深度评测', views: '12.3K' },
  { title: 'RTX 5090对比测试', views: '9.8K' },
  { title: 'GPT-5开发者体验报告', views: '8.1K' },
  { title: '夜行者之剑Boss攻略', views: '6.5K' },
  { title: 'Switch 2首发游戏清单', views: '5.2K' }
]

const articles: { id: string; title: string; summary: string; category: ArticleCategory; author: string; publishDate: string; readTime: number; xpReward: number }[] = [
  { id: '1', title: '苹果发布Vision Pro 2：空间计算进入新时代', summary: '第二代Vision Pro在重量、显示分辨率和手势追踪方面实现重大突破，售价降低30%，有望推动空间计算设备进入消费市场。', category: 'tech', author: 'Kira', publishDate: '2025-06-28', readTime: 13, xpReward: 200 },
  { id: '2', title: '《艾尔登法环》DLC夜行者之剑全球首发', summary: 'FromSoftware全新资料片带来超过40小时的新剧情内容、三张大型地下地图和全新武器技能树系统。', category: 'game', author: 'Rein', publishDate: '2025-06-27', readTime: 26, xpReward: 350 },
  { id: '3', title: '特斯拉Optimus机器人批量生产计划公布', summary: '马斯克宣布Optimus Gen-3将于2026年进入量产，单台售价降至2万美元以下，首批部署在特斯拉超级工厂。', category: 'tech', author: 'Zero', publishDate: '2025-06-26', readTime: 9, xpReward: 180 },
  { id: '4', title: 'NVIDIA RTX 5090性能实测：光追再破纪录', summary: 'RTX 5090在4K光追环境下平均帧率提升62%，功耗却降低15%，Blackwell架构效率令人瞩目。', category: 'hardware', author: 'Vex', publishDate: '2025-06-25', readTime: 34, xpReward: 250 },
  { id: '5', title: '《咒术回战》完结篇动画制作决定', summary: 'MAPPA确认将制作原创新结局动画电影，预计2026年夏季上映，芥见下监督亲自参与脚本撰写。', category: 'anime', author: 'Mika', publishDate: '2025-06-24', readTime: 18, xpReward: 150 },
  { id: '6', title: 'OpenAI发布GPT-5多模态模型', summary: 'GPT-5实现原生视频理解和实时代码生成，推理速度提升4倍，API定价降低60%。', category: 'tech', author: 'Ash', publishDate: '2025-06-23', readTime: 51, xpReward: 300 },
  { id: '7', title: '任天堂Switch 2首批游戏阵容曝光', summary: '包括马力欧卡丁车9、密特罗德4和全新IP，支持4K输出和磁吸Joy-Con手柄。', category: 'game', author: 'Yuki', publishDate: '2025-06-22', readTime: 20, xpReward: 280 },
  { id: '8', title: '《间谍过家家》第三季十月开播', summary: 'WIT STUDIO和CloverWorks联合制作，约尔篇为主线，揭露更多twilight组织内幕。', category: 'anime', author: 'Luna', publishDate: '2025-06-21', readTime: 15, xpReward: 120 }
]

const filteredArticles = computed(() => {
  return articles.filter(article => {
    const matchesCategory = currentCategory.value === 'all' || article.category === currentCategory.value
    const matchesSearch = article.title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
                          article.summary.toLowerCase().includes(searchQuery.value.toLowerCase())
    return matchesCategory && matchesSearch
  })
})

const getCategoryColor = (category: ArticleCategory) => {
  const colors: Record<ArticleCategory, string> = {
    tech: 'var(--color-primary)',
    game: 'var(--color-secondary)',
    hardware: 'var(--color-state-info)',
    anime: 'var(--color-state-warning)'
  }
  return colors[category] || 'var(--color-primary)'
}

const getCategoryBg = (category: ArticleCategory) => {
  const bgs: Record<ArticleCategory, string> = {
    tech: 'var(--color-primary-50)',
    game: 'var(--color-secondary-50)',
    hardware: 'rgba(0,170,255,0.08)',
    anime: 'rgba(255,170,0,0.08)'
  }
  return bgs[category] || 'var(--color-primary-50)'
}

const getCategoryBgLight = (category: ArticleCategory) => {
  const bgs: Record<ArticleCategory, string> = {
    tech: 'var(--color-primary-100)',
    game: 'var(--color-secondary-100)',
    hardware: 'rgba(0,170,255,0.15)',
    anime: 'rgba(255,170,0,0.15)'
  }
  return bgs[category] || 'var(--color-primary-100)'
}

const getCategoryName = (category: ArticleCategory) => {
  const names: Record<ArticleCategory, string> = {
    tech: '科技',
    game: '游戏',
    hardware: '硬件',
    anime: '二次元'
  }
  return names[category] || '科技'
}

const getLeaderboardColor = (index: number) => {
  const colors = ['var(--color-primary)', 'var(--color-secondary)', 'var(--color-state-info)', 'var(--color-state-warning)', 'var(--color-bg-elevated)']
  return colors[index] || 'var(--color-bg-elevated)'
}
</script>