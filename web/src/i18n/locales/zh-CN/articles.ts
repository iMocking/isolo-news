/**
 * 中文文章数据语言包
 * 包含：文章 mock 数据、文章正文、评论数据、相关文章、热搜标签
 */

// ==================== 文章 Mock 数据（articleStore 用） ====================
export const articleList = [
  {
    id: '1',
    title: 'GPT-5来了:AI将如何改变游戏开发',
    summary: '深度解析GPT-5在游戏叙事、NPC交互和程序化生成领域的革命性突破，以及它对独立开发者和大型工作室的深远影响。',
    author: '赛博观察者',
    category: 'tech',
    publishDate: '2小时前',
    readTime: 15,
    xpReward: 200,
    tags: ['AI', '游戏开发', 'GPT-5']
  },
  {
    id: '2',
    title: '《黑神话2》首曝:虚幻引擎6实机演示',
    summary: '游戏科学最新力作首次公开实机画面，展示虚幻引擎6在光照、物理和开放世界领域的突破性进展。',
    author: '游戏前线',
    category: 'game',
    publishDate: '1天前',
    readTime: 8,
    xpReward: 180,
    tags: ['黑神话', '虚幻引擎6', '游戏']
  },
  {
    id: '3',
    title: 'AMD RX 9070 XT深度评测',
    summary: '新一代RDNA 5架构显卡实测，4K光追性能对比NVIDIA旗舰，功耗与性价比全面分析。',
    author: '硬件频道',
    category: 'hardware',
    publishDate: '3小时前',
    readTime: 12,
    xpReward: 150,
    tags: ['AMD', '显卡', '硬件评测']
  },
  {
    id: '4',
    title: '《葬送的芙莉莲》第二季定档',
    summary: '人气奇幻动画续作确认秋季开播，原班制作团队回归，新篇章将深入魔王城篇。',
    author: '动漫情报',
    category: 'anime',
    publishDate: '5小时前',
    readTime: 6,
    xpReward: 120,
    tags: ['芙莉莲', '动画', '二次元']
  },
  {
    id: '5',
    title: '苹果Vision Pro 2正式发布',
    summary: '第二代空间计算设备重量减半，售价下探至2499美元，支持全新手势交互系统。',
    author: '科技快报',
    category: 'tech',
    publishDate: '8小时前',
    readTime: 10,
    xpReward: 250,
    tags: ['苹果', 'VR', '科技']
  },
  {
    id: '6',
    title: 'Steam夏季特卖最佳折扣推荐',
    summary: '精选30款不可错过的折扣游戏，从独立佳作到3A大作，省钱攻略一网打尽。',
    author: '游戏推荐',
    category: 'game',
    publishDate: '12小时前',
    readTime: 20,
    xpReward: 100,
    tags: ['Steam', '折扣', '游戏']
  }
]

// ==================== 热搜标签（articleStore 用） ====================
export const trendingTopics: string[] = [
  'PS6首发', 'AI绘画革命', '原神5.0', 'RTX5090评测', '量子芯片突破'
]

// ==================== 文章详情页正文 ====================
export const articleBody = `<p style="color: var(--color-text-secondary);">
  <span style="font-family: var(--font-mono); color: var(--color-primary);">$ </span>
  2026年7月，OpenAI正式发布GPT-5，这一里程碑事件迅速在游戏开发行业引发了连锁反应。从程序化内容生成到智能NPC对话系统，从自动Bug检测到美术资源优化，GPT-5展现出的多模态推理能力远超业内预期。据内部测试数据显示，GPT-5在代码生成准确率上较前代提升了47%，在复杂逻辑推理任务上的表现已接近高级工程师水平。
</p>

<h2 class="pt-4 pb-2" style="font-family: var(--font-display); font-size: 22px; font-weight: 600; color: var(--color-primary); border-left: 3px solid var(--color-primary); padding-left: 12px; letter-spacing: -0.01em;">
  程序化内容生成的新纪元
</h2>

<p style="color: var(--color-text-secondary);">
  游戏世界构建历来是开发成本最高的环节之一。传统AAA级游戏的开放世界地图往往需要上百名美术师和关卡设计师耗费数年才能完成。GPT-5的3D空间理解能力使得程序化生成从"随机拼凑"进化到了"语义驱动"。引擎能够根据剧情需要自动生成符合叙事逻辑的地形、建筑群落甚至天气系统，同时保持视觉风格的高度一致性。
</p>

<h2 class="pt-4 pb-2" style="font-family: var(--font-display); font-size: 22px; font-weight: 600; color: var(--color-primary); border-left: 3px solid var(--color-primary); padding-left: 12px; letter-spacing: -0.01em;">
  智能NPC与动态叙事
</h2>

<p style="color: var(--color-text-secondary);">
  NPC的行为逻辑一直是游戏开发中最难突破的瓶颈。预编写的对话树无论多么复杂，终究有穷尽之时。GPT-5的上下文窗口扩展至百万级Token，使得NPC能够维持跨章节的长期记忆，并根据玩家的行为模式动态调整交互策略。
</p>`

// ==================== 文章详情页评论 ====================
export const comments = [
  { id: 1, author: '量子旅人', avatar: 'Q', content: 'GPT-5的出现确实会改变游戏开发的格局，期待更多AI辅助创作的佳作！', time: '2小时前', likes: 23 },
  { id: 2, author: '代码忍者', avatar: 'D', content: '作为独立开发者，终于可以把更多精力放在创意上了，技术门槛降低是好事。', time: '5小时前', likes: 15 },
  { id: 3, author: '像素画师', avatar: 'X', content: '担心AI会让美术从业者失业，但转念一想，AI也是工具，关键看谁在用。', time: '8小时前', likes: 42 }
]

// ==================== 文章详情页相关文章 ====================
export const relatedArticles = [
  { title: 'Unity 2026发布：AI原生工作流正式上线', time: '2小时前' },
  { title: '索尼公布PS6开发套件：光线追踪性能提升300%', time: '5小时前' },
  { title: '深度学习驱动的物理引擎：下一代游戏仿真', time: '1天前' },
  { title: 'AI辅助美术：从概念草图到3D模型的自动化管线', time: '2天前' }
]

// ==================== 文章列表页静态数据 ====================
export const listPageArticles = [
  { id: '1', title: '苹果发布Vision Pro 2：空间计算进入新时代', summary: '第二代Vision Pro在重量、显示分辨率和手势追踪方面实现重大突破，售价降低30%，有望推动空间计算设备进入消费市场。', category: 'tech', author: 'Kira', publishDate: '2025-06-28', readTime: 13, xpReward: 200 },
  { id: '2', title: '《艾尔登法环》DLC夜行者之剑全球首发', summary: 'FromSoftware全新资料片带来超过40小时的新剧情内容、三张大型地下地图和全新武器技能树系统。', category: 'game', author: 'Rein', publishDate: '2025-06-27', readTime: 26, xpReward: 350 },
  { id: '3', title: '特斯拉Optimus机器人批量生产计划公布', summary: '马斯克宣布Optimus Gen-3将于2026年进入量产，单台售价降至2万美元以下，首批部署在特斯拉超级工厂。', category: 'tech', author: 'Zero', publishDate: '2025-06-26', readTime: 9, xpReward: 180 },
  { id: '4', title: 'NVIDIA RTX 5090性能实测：光追再破纪录', summary: 'RTX 5090在4K光追环境下平均帧率提升62%，功耗却降低15%，Blackwell架构效率令人瞩目。', category: 'hardware', author: 'Vex', publishDate: '2025-06-25', readTime: 34, xpReward: 250 },
  { id: '5', title: '《咒术回战》完结篇动画制作决定', summary: 'MAPPA确认将制作原创新结局动画电影，预计2026年夏季上映，芥见下监督亲自参与脚本撰写。', category: 'anime', author: 'Mika', publishDate: '2025-06-24', readTime: 18, xpReward: 150 },
  { id: '6', title: 'OpenAI发布GPT-5多模态模型', summary: 'GPT-5实现原生视频理解和实时代码生成，推理速度提升4倍，API定价降低60%。', category: 'tech', author: 'Ash', publishDate: '2025-06-23', readTime: 51, xpReward: 300 },
  { id: '7', title: '任天堂Switch 2首批游戏阵容曝光', summary: '包括马力欧卡丁车9、密特罗德4和全新IP，支持4K输出和磁吸Joy-Con手柄。', category: 'game', author: 'Yuki', publishDate: '2025-06-22', readTime: 20, xpReward: 280 },
  { id: '8', title: '《间谍过家家》第三季十月开播', summary: 'WIT STUDIO和CloverWorks联合制作，约尔篇为主线，揭露更多twilight组织内幕。', category: 'anime', author: 'Luna', publishDate: '2025-06-21', readTime: 15, xpReward: 120 }
]

// ==================== 文章列表侧栏排行榜 ====================
export const leaderboard = [
  { title: '苹果Vision Pro 2深度评测', views: '12.3K' },
  { title: 'RTX 5090对比测试', views: '9.8K' },
  { title: 'GPT-5开发者体验报告', views: '8.1K' },
  { title: '夜行者之剑Boss攻略', views: '6.5K' },
  { title: 'Switch 2首发游戏清单', views: '5.2K' }
]

// ==================== 文章列表侧栏热门标签 ====================
export const hotTags = [
  { name: 'AI', type: 'primary' },
  { name: 'VR', type: 'secondary' },
  { name: '显卡', type: 'info' },
  { name: '新番', type: 'warning' },
  { name: '机器人', type: 'primary' },
  { name: 'RPG', type: 'secondary' },
  { name: '芯片', type: 'primary' },
  { name: '漫画', type: 'warning' }
]
