/**
 * English article data locale
 * Mirrors zh-CN/articles.ts structure
 */

// ==================== Mock Article List (articleStore) ====================
export const articleList = [
  {
    id: '1',
    title: 'GPT-5 Arrives: How AI Will Change Game Development',
    summary: 'Deep dive into GPT-5\'s revolutionary impact on game narrative, NPC interaction, and procedural generation for indie and AAA studios.',
    author: 'Cyber Observer',
    category: 'tech',
    publishDate: '2h ago',
    readTime: 15,
    xpReward: 200,
    tags: ['AI', 'Game Dev', 'GPT-5']
  },
  {
    id: '2',
    title: 'Black Myth 2 Revealed: Unreal Engine 6 Demo',
    summary: 'Game Science unveils first gameplay footage showcasing Unreal Engine 6 breakthroughs in lighting, physics, and open worlds.',
    author: 'Game Frontline',
    category: 'game',
    publishDate: '1d ago',
    readTime: 8,
    xpReward: 180,
    tags: ['Black Myth', 'UE6', 'Gaming']
  },
  {
    id: '3',
    title: 'AMD RX 9070 XT In-Depth Review',
    summary: 'Benchmarking the new RDNA 5 architecture: 4K ray tracing vs NVIDIA flagships, power efficiency, and value analysis.',
    author: 'Hardware Channel',
    category: 'hardware',
    publishDate: '3h ago',
    readTime: 12,
    xpReward: 150,
    tags: ['AMD', 'GPU', 'Hardware']
  },
  {
    id: '4',
    title: 'Frieren Season 2 Confirmed',
    summary: 'The acclaimed fantasy anime returns this fall with the original production team, diving into the Demon King castle arc.',
    author: 'Anime Intel',
    category: 'anime',
    publishDate: '5h ago',
    readTime: 6,
    xpReward: 120,
    tags: ['Frieren', 'Anime', 'Fantasy']
  },
  {
    id: '5',
    title: 'Apple Vision Pro 2 Launched',
    summary: 'Second-gen spatial computer weighs half as much, starts at $2499, with all-new gesture control system.',
    author: 'Tech Flash',
    category: 'tech',
    publishDate: '8h ago',
    readTime: 10,
    xpReward: 250,
    tags: ['Apple', 'VR', 'Tech']
  },
  {
    id: '6',
    title: 'Steam Summer Sale: Best Deals',
    summary: '30 must-buy discounted games from indie darlings to AAA blockbusters — your ultimate money-saving guide.',
    author: 'Game Picks',
    category: 'game',
    publishDate: '12h ago',
    readTime: 20,
    xpReward: 100,
    tags: ['Steam', 'Sale', 'Gaming']
  }
]

// ==================== Trending Topics (articleStore) ====================
export const trendingTopics: string[] = [
  'PS6 Launch', 'AI Art Revolution', 'Genshin 5.0', 'RTX5090 Review', 'Quantum Chip'
]

// ==================== Article Detail Body ====================
export const articleBody = `<p style="color: var(--color-text-secondary);">
  <span style="font-family: var(--font-mono); color: var(--color-primary);">$ </span>
  In July 2026, OpenAI officially released GPT-5, a milestone that sent ripples through the game development industry. From procedural content generation to intelligent NPC dialogue systems, from automated bug detection to art asset optimization, GPT-5's multimodal reasoning capabilities far exceeded industry expectations. Internal benchmarks show GPT-5 improves code generation accuracy by 47% over its predecessor, approaching senior engineer level on complex logical reasoning tasks.
</p>

<h2 class="pt-4 pb-2" style="font-family: var(--font-display); font-size: 22px; font-weight: 600; color: var(--color-primary); border-left: 3px solid var(--color-primary); padding-left: 12px; letter-spacing: -0.01em;">
  A New Era of Procedural Generation
</h2>

<p style="color: var(--color-text-secondary);">
  Game world building has historically been one of the most costly aspects of development. Traditional AAA open worlds require hundreds of artists and level designers working for years. GPT-5's 3D spatial understanding elevates procedural generation from "random stitching" to "semantic-driven" creation. Engines can now automatically generate terrain, architecture, and even weather systems that follow narrative logic while maintaining visual consistency.
</p>

<h2 class="pt-4 pb-2" style="font-family: var(--font-display); font-size: 22px; font-weight: 600; color: var(--color-primary); border-left: 3px solid var(--color-primary); padding-left: 12px; letter-spacing: -0.01em;">
  Smart NPCs & Dynamic Storytelling
</h2>

<p style="color: var(--color-text-secondary);">
  NPC behavior logic has long been one of the toughest bottlenecks in game development. No matter how complex pre-written dialogue trees are, they eventually run out. GPT-5's context window expands to millions of tokens, enabling NPCs to maintain cross-chapter long-term memory and dynamically adjust interaction strategies based on player behavior patterns.
</p>`

// ==================== Article Comments ====================
export const comments = [
  { id: 1, author: 'Quantum Traveler', avatar: 'Q', content: 'GPT-5 will definitely reshape game development. Looking forward to more AI-assisted masterpieces!', time: '2h ago', likes: 23 },
  { id: 2, author: 'Code Ninja', avatar: 'D', content: 'As an indie dev, I can finally focus more on creativity. Lowering the technical barrier is great.', time: '5h ago', likes: 15 },
  { id: 3, author: 'Pixel Artist', avatar: 'X', content: 'Worried about AI replacing artists, but then again, AI is just a tool — it depends on who uses it.', time: '8h ago', likes: 42 }
]

// ==================== Related Articles ====================
export const relatedArticles = [
  { title: 'Unity 2026: AI-Native Workflow Officially Launches', time: '2h ago' },
  { title: 'Sony Reveals PS6 Dev Kit: Ray Tracing Up 300%', time: '5h ago' },
  { title: 'Deep Learning Physics Engine: Next-Gen Game Simulation', time: '1d ago' },
  { title: 'AI-Assisted Art: From Concept to 3D Asset Pipeline', time: '2d ago' }
]

// ==================== Article List Page Static Data ====================
export const listPageArticles = [
  { id: '1', title: 'Apple Vision Pro 2: Spatial Computing Enters a New Era', summary: 'Second-gen Vision Pro achieves major breakthroughs in weight, display resolution, and hand tracking, with a 30% price reduction.', category: 'tech', author: 'Kira', publishDate: '2025-06-28', readTime: 13, xpReward: 200 },
  { id: '2', title: 'Elden Ring DLC Nightblade Global Launch', summary: 'FromSoftware\'s new expansion brings 40+ hours of content, three massive underground maps, and a new weapon skill tree system.', category: 'game', author: 'Rein', publishDate: '2025-06-27', readTime: 26, xpReward: 350 },
  { id: '3', title: 'Tesla Optimus Robot Mass Production Plan Announced', summary: 'Musk announces Optimus Gen-3 enters mass production in 2026 at under $20K per unit, first deployed at Gigafactories.', category: 'tech', author: 'Zero', publishDate: '2025-06-26', readTime: 9, xpReward: 180 },
  { id: '4', title: 'NVIDIA RTX 5090 Performance Review: Ray Tracing Records Broken', summary: 'RTX 5090 delivers 62% higher average FPS in 4K ray tracing with 15% lower power draw — Blackwell architecture impresses.', category: 'hardware', author: 'Vex', publishDate: '2025-06-25', readTime: 34, xpReward: 250 },
  { id: '5', title: 'Jujutsu Kaisen Final Arc Movie Announced', summary: 'MAPPA confirms original ending anime film for summer 2026, with creator Gege Akutami personally involved in scripting.', category: 'anime', author: 'Mika', publishDate: '2025-06-24', readTime: 18, xpReward: 150 },
  { id: '6', title: 'OpenAI Launches GPT-5 Multimodal Model', summary: 'GPT-5 achieves native video understanding and real-time code generation, with 4x inference speed and 60% lower API pricing.', category: 'tech', author: 'Ash', publishDate: '2025-06-23', readTime: 51, xpReward: 300 },
  { id: '7', title: 'Nintendo Switch 2 Launch Lineup Leaked', summary: 'Includes Mario Kart 9, Metroid Prime 4, and new IPs, with 4K output and magnetic Joy-Con controllers.', category: 'game', author: 'Yuki', publishDate: '2025-06-22', readTime: 20, xpReward: 280 },
  { id: '8', title: 'Spy x Family Season 3 Premieres October', summary: 'WIT STUDIO and CloverWorks co-produce the Yor arc, revealing more of the twilight organization\'s secrets.', category: 'anime', author: 'Luna', publishDate: '2025-06-21', readTime: 15, xpReward: 120 }
]

// ==================== Sidebar Leaderboard ====================
export const leaderboard = [
  { title: 'Apple Vision Pro 2 Deep Dive', views: '12.3K' },
  { title: 'RTX 5090 vs Competition', views: '9.8K' },
  { title: 'GPT-5 Dev Experience Report', views: '8.1K' },
  { title: 'Elden Ring DLC Boss Guide', views: '6.5K' },
  { title: 'Switch 2 Launch Games Roundup', views: '5.2K' }
]

// ==================== Sidebar Hot Tags ====================
export const hotTags = [
  { name: 'AI', type: 'primary' },
  { name: 'VR', type: 'secondary' },
  { name: 'GPU', type: 'info' },
  { name: 'Anime', type: 'warning' },
  { name: 'Robotics', type: 'primary' },
  { name: 'RPG', type: 'secondary' },
  { name: 'Chips', type: 'primary' },
  { name: 'Manga', type: 'warning' }
]
