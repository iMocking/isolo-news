<template>
  <div class="min-h-screen">
    <AppNavigation />

    <!-- Article Hero -->
    <section class="w-full relative overflow-hidden" style="background: linear-gradient(180deg, var(--color-bg-secondary) 0%, var(--color-bg-card) 100%);">
      <!-- Scanline overlay for NEXUS theme -->
      <div v-if="themeStore.currentTheme === 'nexus'" class="absolute inset-0 pointer-events-none opacity-5" style="background: repeating-linear-gradient(0deg, transparent, transparent 2px, var(--color-text-primary) 2px, var(--color-text-primary) 3px);"></div>

      <div class="max-w-7xl mx-auto px-6 pt-12 pb-10 relative z-10">
        <!-- Breadcrumb -->
        <div class="flex items-center gap-2 mb-8">
          <a href="/" class="text-xs whitespace-nowrap transition-colors duration-150" style="color: var(--color-text-tertiary);" @mouseenter="($event.target as HTMLElement).style.color = 'var(--color-primary)'" @mouseleave="($event.target as HTMLElement).style.color = 'var(--color-text-tertiary)'">首页</a>
          <span class="text-xs" style="color: var(--color-text-tertiary);">/</span>
          <a href="/list" class="text-xs whitespace-nowrap transition-colors duration-150" style="color: var(--color-text-tertiary);" @mouseenter="($event.target as HTMLElement).style.color = 'var(--color-primary)'" @mouseleave="($event.target as HTMLElement).style.color = 'var(--color-text-tertiary)'">科技资讯</a>
          <span class="text-xs" style="color: var(--color-text-tertiary);">/</span>
          <span class="text-xs whitespace-nowrap" style="color: var(--color-text-secondary);">任务详情</span>
        </div>

        <!-- Category badge -->
        <div class="mb-6">
          <span class="inline-flex items-center px-3 py-1 text-xs font-semibold whitespace-nowrap" :style="{ background: getCategoryBg(currentArticle.category), color: getCategoryColor(currentArticle.category), border: `1px solid ${getCategoryBorder(currentArticle.category)}`, fontFamily: 'var(--font-display)', letterSpacing: '0.05em' }">{{ getCategoryName(currentArticle.category) }}</span>
        </div>

        <!-- Article title -->
        <h1 class="mb-6" style="font-size: clamp(28px, 3.5vw, 48px); line-height: 1.2; font-weight: 600; color: var(--color-text-primary); word-break: keep-all;">
          {{ currentArticle.title }}
        </h1>

        <!-- Author info row -->
        <div class="flex flex-wrap items-center gap-4 mb-6">
          <div class="flex items-center gap-3">
            <div class="w-9 h-9 rounded-full flex items-center justify-center shrink-0" style="background: linear-gradient(135deg, var(--color-primary-200), var(--color-secondary-100)); border: 1px solid var(--color-border);">
              <span class="text-sm font-bold" style="color: var(--color-text-inverse);">{{ currentArticle.author.charAt(0).toUpperCase() }}</span>
            </div>
            <div class="flex items-center gap-2">
              <span class="text-sm font-medium whitespace-nowrap" style="color: var(--color-text-primary);">{{ currentArticle.author }}</span>
              <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">{{ currentArticle.publishDate }}</span>
            </div>
          </div>
          <div class="flex items-center gap-1">
            <Eye class="w-4 h-4" style="color: var(--color-text-tertiary);" />
            <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">阅读 {{ currentArticle.readTime * 1000 }}</span>
          </div>
        </div>

        <!-- XP Progress Bar -->
        <div class="mb-8 p-4 max-w-2xl" style="background: var(--color-bg-elevated); border: 1px solid var(--color-border); border-radius: var(--radius-md);">
          <div class="flex items-center justify-between mb-2">
            <span class="text-xs font-medium whitespace-nowrap" style="color: var(--color-primary); font-family: var(--font-display);">XP 奖励进度</span>
            <span class="text-xs whitespace-nowrap" style="color: var(--color-text-secondary);">+{{ currentArticle.xpReward }} XP</span>
          </div>
          <div class="w-full h-2 rounded-full overflow-hidden" style="background: var(--color-bg-tertiary);">
            <div class="h-full rounded-full transition-all duration-300" :style="{ width: `${readProgress}%`, background: 'linear-gradient(90deg, var(--color-primary), var(--color-secondary))' }"></div>
          </div>
          <div class="flex items-center justify-between mt-2">
            <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">阅读进度 {{ readProgress }}%</span>
            <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">读完可获得 +500 XP</span>
          </div>
        </div>

        <!-- Tag pills -->
        <div class="flex flex-wrap gap-2">
          <span v-for="tag in articleTags" :key="tag" class="inline-flex items-center px-3 py-1 text-xs whitespace-nowrap" style="background: var(--color-bg-elevated); color: var(--color-text-secondary); border: 1px solid var(--color-border-subtle);">{{ tag }}</span>
        </div>
      </div>
    </section>

    <!-- Article Body + Sidebar -->
    <section class="max-w-7xl mx-auto px-6 py-12">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
        <!-- Main article content -->
        <article class="lg:col-span-2">
          <!-- Terminal frame container -->
          <div class="relative" style="background: var(--color-bg-card); border: 1px solid var(--color-border); border-radius: var(--radius-lg);">
            <!-- Terminal header bar -->
            <div class="flex items-center gap-2 px-4 py-3" style="background: var(--color-bg-elevated); border-bottom: 1px solid var(--color-border); border-radius: var(--radius-lg) var(--radius-lg) 0 0;">
              <div class="w-3 h-3 rounded-full shrink-0" style="background: var(--state-error);"></div>
              <div class="w-3 h-3 rounded-full shrink-0" style="background: var(--state-warning);"></div>
              <div class="w-3 h-3 rounded-full shrink-0" style="background: var(--state-success);"></div>
              <span class="text-xs ml-2 whitespace-nowrap" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">nexus://article/gpt5-game-dev.log</span>
            </div>

            <!-- Scanline overlay on terminal body -->
            <div v-if="themeStore.currentTheme === 'nexus'" class="absolute inset-0 pointer-events-none opacity-[0.03] z-10" style="background: repeating-linear-gradient(0deg, transparent, transparent 1px, var(--color-text-primary) 1px, var(--color-text-primary) 2px); border-radius: var(--radius-lg);"></div>

            <!-- Article text body -->
            <div class="relative z-10 p-6 md:p-10 space-y-6" style="font-size: 15px; line-height: 1.7; color: var(--color-text-primary);">
              <p style="color: var(--color-text-secondary);">
                <span style="font-family: var(--font-mono); color: var(--color-primary);">$ </span>
                2026年7月，OpenAI正式发布GPT-5，这一里程碑事件迅速在游戏开发行业引发了连锁反应。从程序化内容生成到智能NPC对话系统，从自动Bug检测到美术资源优化，GPT-5展现出的多模态推理能力远超业内预期。据内部测试数据显示，GPT-5在代码生成准确率上较前代提升了47%，在复杂逻辑推理任务上的表现已接近高级工程师水平。
              </p>

              <h2 class="pt-4 pb-2" style="font-family: var(--font-display); font-size: 22px; font-weight: 600; color: var(--color-primary); border-left: 3px solid var(--color-primary); padding-left: 12px; letter-spacing: -0.01em;">
                程序化内容生成的新纪元
              </h2>

              <p style="color: var(--color-text-secondary);">
                游戏世界构建历来是开发成本最高的环节之一。传统AAA级游戏的开放世界地图往往需要上百名美术师和关卡设计师耗费数年才能完成。GPT-5的3D空间理解能力使得程序化生成从"随机拼凑"进化到了"语义驱动"。引擎能够根据剧情需要自动生成符合叙事逻辑的地形、建筑群落甚至天气系统，同时保持视觉风格的高度一致性。
              </p>

              <!-- Image placeholder -->
              <div class="w-full h-48 md:h-64 flex items-center justify-center relative overflow-hidden" style="background: linear-gradient(135deg, var(--color-bg-tertiary) 0%, var(--color-bg-elevated) 100%); border: 1px solid var(--color-border-subtle);">
                <div class="absolute inset-0 flex items-center justify-center">
                  <span class="text-xs tracking-widest" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">[ RENDERING ASSET... ]</span>
                </div>
                <div class="absolute bottom-2 right-2">
                  <span class="text-xs px-2 py-0.5 whitespace-nowrap" style="color: var(--color-text-tertiary); background: var(--color-bg-primary); font-family: var(--font-mono);">IMG_001.jpg</span>
                </div>
              </div>

              <!-- Pull quote -->
              <div class="my-8 p-6 relative" style="background: var(--color-secondary-50); border-left: 3px solid var(--color-secondary);">
                <span class="absolute -top-2 -left-1 text-4xl" style="color: var(--color-secondary); opacity: 0.3; font-family: var(--font-display);">"</span>
                <p class="text-base italic pl-2" style="color: var(--color-text-primary); line-height: 1.6;">
                  "GPT-5不是在替代开发者，而是在重新定义创作的边界。当机器能够处理80%的重复性工作，剩余20%的创意空间反而变得更加广阔。"
                </p>
                <p class="mt-3 pl-2 text-xs" style="color: var(--color-text-tertiary);">
                  -- 某大型工作室技术总监
                </p>
              </div>

              <h2 class="pt-4 pb-2" style="font-family: var(--font-display); font-size: 22px; font-weight: 600; color: var(--color-primary); border-left: 3px solid var(--color-primary); padding-left: 12px; letter-spacing: -0.01em;">
                智能NPC与动态叙事
              </h2>

              <p style="color: var(--color-text-secondary);">
                NPC的行为逻辑一直是游戏开发中最难突破的瓶颈。预编写的对话树无论多么复杂，终究有穷尽之时。GPT-5的上下文窗口扩展至百万级Token，使得NPC能够维持跨章节的长期记忆，并根据玩家的行为模式动态调整交互策略。早期测试中，部分玩家甚至无法分辨对手是AI还是真人操控——这在多人竞技游戏中引发了关于"AI伪装"的伦理讨论。
              </p>

              <p style="color: var(--color-text-secondary);">
                更令人瞩目的是，GPT-5展现出了跨模态的叙事理解能力。它不仅能生成自然语言对话，还能同步控制角色的面部表情、肢体动作和语音语调的参数化输出。这意味着游戏中的过场动画可以根据玩家此前的选择实时生成，而非依赖预渲染的视频片段。
              </p>

              <!-- Code snippet -->
              <div class="my-6 p-4 overflow-x-auto" style="background: var(--color-bg-primary); border: 1px solid var(--color-border-subtle);">
                <div class="flex items-center justify-between mb-3">
                  <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">npc_ai_config.py</span>
                  <span class="text-xs whitespace-nowrap" style="color: var(--color-primary); font-family: var(--font-mono);">Python</span>
                </div>
                <pre class="text-xs leading-relaxed whitespace-pre" style="color: var(--color-text-secondary); font-family: var(--font-mono);"><span style="color: var(--color-secondary);">class</span> <span style="color: var(--color-primary);">NPCEngine</span>:
    <span style="color: var(--color-secondary);">def</span> <span style="color: var(--color-text-primary);">__init__</span>(self, model=<span style="color: var(--color-state-warning);">"gpt-5-turbo"</span>):
        self.model = model
        self.memory = ConversationMemory(
            max_tokens=<span style="color: var(--color-secondary-light);">1_000_000</span>
        )
        self.persona = <span style="color: var(--color-state-warning);">None</span>

    <span style="color: var(--color-secondary);">async def</span> <span style="color: var(--color-text-primary);">generate_response</span>(self, player_input):
        context = self.memory.get_context()
        response = <span style="color: var(--color-secondary);">await</span> self.model.chat(
            messages=context + [player_input],
            persona=self.persona
        )
        <span style="color: var(--color-secondary);">return</span> response</pre>
              </div>

              <h2 class="pt-4 pb-2" style="font-family: var(--font-display); font-size: 22px; font-weight: 600; color: var(--color-primary); border-left: 3px solid var(--color-primary); padding-left: 12px; letter-spacing: -0.01em;">
                行业冲击与开发者反馈
              </h2>

              <p style="color: var(--color-text-secondary);">
                并非所有声音都是乐观的。独立开发者群体对GPT-5的态度呈现出明显的分化。一部分人将其视为"创作民主化"的工具，认为它大幅降低了中小团队的技术门槛，使得小规模团队也能产出接近AAA品质的内容。另一部分人则担心，大型发行商会借此进一步压缩开发周期，将更多人力成本转嫁给AI工具，导致行业岗位进一步缩减。
              </p>

              <p style="color: var(--color-text-secondary);">
                从技术生态的角度看，Unity和Unreal Engine均已宣布将在下一版本中深度集成GPT-5的API接口。这意味着AI辅助开发将从"可选的外部工具"转变为"引擎原生能力"。对于开发者而言，掌握Prompt工程和AI协作流程正在成为与传统编程技能同等重要的核心竞争力。
              </p>

              <!-- Second image placeholder -->
              <div class="w-full h-48 md:h-64 flex items-center justify-center relative overflow-hidden" style="background: linear-gradient(135deg, var(--color-bg-elevated) 0%, var(--color-primary-50) 100%); border: 1px solid var(--color-border-subtle);">
                <div class="absolute inset-0 flex items-center justify-center">
                  <span class="text-xs tracking-widest" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">[ LOADING VISUAL DATA... ]</span>
                </div>
                <div class="absolute bottom-2 right-2">
                  <span class="text-xs px-2 py-0.5 whitespace-nowrap" style="color: var(--color-text-tertiary); background: var(--color-bg-primary); font-family: var(--font-mono);">IMG_002.jpg</span>
                </div>
              </div>

              <p style="color: var(--color-text-secondary);">
                GPT-5的到来并非AI替代人类的终点，而是人机协作新范式的起点。游戏开发行业的每一次技术跃迁——从2D到3D、从单机到在线——最终都带来了市场总量的扩张而非萎缩。这一次，或许也不例外。
              </p>

              <div class="flex items-center gap-1 mt-6 pt-4" style="border-top: 1px solid var(--color-border-subtle);">
                <span class="text-xs" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">EOF</span>
              </div>
            </div>
          </div>

          <!-- Article Footer -->
          <div class="mt-8 space-y-6">
            <!-- Tags row -->
            <div class="flex flex-wrap gap-2">
              <span v-for="tag in currentArticle.tags" :key="tag" class="inline-flex items-center px-3 py-1 text-xs whitespace-nowrap" style="background: var(--color-primary-50); color: var(--color-primary); border: 1px solid var(--color-border);">#{{ tag }}</span>
            </div>

            <!-- Author bio card -->
            <div class="flex items-center gap-4 p-5" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);">
              <div class="w-12 h-12 rounded-full flex items-center justify-center shrink-0" style="background: linear-gradient(135deg, var(--color-primary-200), var(--color-secondary-100)); border: 1px solid var(--color-border);">
                <span class="text-lg font-bold" style="color: var(--color-text-inverse);">A</span>
              </div>
              <div class="min-w-0 flex-1">
                <p class="text-sm font-semibold truncate" style="color: var(--color-text-primary);">量子记者 · Alice</p>
                <p class="text-xs mt-1 line-clamp-2" style="color: var(--color-text-tertiary);">资深科技记者，专注AI与游戏产业交叉领域。曾报道多届E3与TGS，对技术趋势有独到见解。</p>
              </div>
            </div>

            <!-- Back button -->
            <router-link to="/list" class="inline-flex items-center justify-center gap-2 px-5 py-2.5 text-sm font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-primary-50); color: var(--color-primary); border: 1px solid var(--color-border);" @mouseenter="($event.target as HTMLElement).style.background = 'var(--color-primary-100)'; ($event.target as HTMLElement).style.transform = 'translateY(-1px)'" @mouseleave="($event.target as HTMLElement).style.background = 'var(--color-primary-50)'; ($event.target as HTMLElement).style.transform = 'translateY(0)'">
              <ArrowLeft class="w-4 h-4" />
              <span>返回任务列表</span>
            </router-link>
          </div>

          <!-- Comments Section -->
          <div class="mt-12">
            <div class="flex items-center gap-3 mb-6">
              <span class="w-1 h-5 rounded-sm" style="background: var(--color-primary);"></span>
              <h3 class="text-lg font-semibold" style="color: var(--color-text-primary); font-family: var(--font-display);">
                评论区 <span class="text-sm font-normal" style="color: var(--color-text-tertiary);">({{ comments.length }})</span>
              </h3>
            </div>

            <!-- Comment Input -->
            <div class="p-4 mb-6" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);">
              <div class="flex items-start gap-3">
                <div class="w-9 h-9 rounded-full flex items-center justify-center shrink-0" style="background: linear-gradient(135deg, var(--color-primary-200), var(--color-secondary-100)); border: 1px solid var(--color-border);">
                  <span class="text-sm font-bold" style="color: var(--color-text-inverse);">{{ userStore.playerName.charAt(0).toUpperCase() }}</span>
                </div>
                <div class="flex-1">
                  <textarea
                    v-model="newComment"
                    class="w-full p-3 text-sm outline-none resize-none transition-all duration-150"
                    style="background: var(--color-bg-primary); border: 1px solid var(--color-border); border-radius: var(--radius-sm); color: var(--color-text-primary); font-family: var(--font-body); min-height: 80px;"
                    placeholder="写下你的评论..."
                  ></textarea>
                  <div class="flex justify-end mt-3">
                    <button
                      class="flex items-center gap-2 px-4 py-2 text-sm font-medium transition-all duration-150"
                      :style="{
                        background: newComment.trim() ? 'var(--color-primary)' : 'var(--color-bg-tertiary)',
                        color: newComment.trim() ? 'white' : 'var(--color-text-tertiary)',
                        borderRadius: 'var(--radius-sm)',
                        cursor: newComment.trim() ? 'pointer' : 'not-allowed'
                      }"
                      :disabled="!newComment.trim()"
                      @click="handleAddComment"
                    >
                      <Send class="w-4 h-4" />
                      <span>发表评论</span>
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <!-- Comment List -->
            <div class="space-y-4">
              <div
                v-for="comment in comments"
                :key="comment.id"
                class="p-4 transition-all duration-150 hover:translate-y-[-1px]"
                style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);"
              >
                <div class="flex items-start gap-3">
                  <div class="w-9 h-9 rounded-full flex items-center justify-center shrink-0" style="background: linear-gradient(135deg, var(--color-primary-200), var(--color-secondary-100)); border: 1px solid var(--color-border);">
                    <span class="text-sm font-bold" style="color: var(--color-text-inverse);">{{ comment.avatar }}</span>
                  </div>
                  <div class="flex-1 min-w-0">
                    <div class="flex items-center justify-between mb-2">
                      <span class="text-sm font-medium" style="color: var(--color-text-primary);">{{ comment.author }}</span>
                      <span class="text-xs" style="color: var(--color-text-tertiary);">{{ comment.time }}</span>
                    </div>
                    <p class="text-sm mb-3" style="color: var(--color-text-secondary); line-height: 1.6;">{{ comment.content }}</p>
                    <div class="flex items-center gap-4">
                      <button class="flex items-center gap-1 text-xs transition-colors duration-150" style="color: var(--color-text-tertiary);" @mouseenter="($event.target as HTMLElement).style.color = 'var(--color-primary)'" @mouseleave="($event.target as HTMLElement).style.color = 'var(--color-text-tertiary)'">
                        <Heart class="w-3.5 h-3.5" />
                        <span>{{ comment.likes }}</span>
                      </button>
                      <button class="text-xs transition-colors duration-150" style="color: var(--color-text-tertiary);" @mouseenter="($event.target as HTMLElement).style.color = 'var(--color-primary)'" @mouseleave="($event.target as HTMLElement).style.color = 'var(--color-text-tertiary)'">
                        回复
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </article>

        <!-- Sidebar -->
        <aside class="lg:col-span-1 space-y-8">
          <!-- Related Missions -->
          <div class="p-5" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);">
            <h3 class="text-sm font-semibold mb-4 pb-3 whitespace-nowrap" style="color: var(--color-text-primary); font-family: var(--font-display); letter-spacing: 0.02em; border-bottom: 1px solid var(--color-border-subtle);">相关任务</h3>
            <div class="space-y-4">
              <a v-for="(item, index) in relatedArticles" :key="index" href="#" class="block group">
                <div class="flex items-start gap-3">
                  <div class="w-16 h-10 shrink-0 flex items-center justify-center overflow-hidden" :style="{ background: `linear-gradient(135deg, ${index % 2 === 0 ? 'var(--color-bg-tertiary)' : 'var(--color-bg-elevated)'}, ${index % 2 === 0 ? 'var(--color-bg-elevated)' : 'var(--color-bg-tertiary)'})` }">
                    <span class="text-xs" style="color: var(--color-text-tertiary); font-family: var(--font-mono);">{{ String(index + 1).padStart(2, '0') }}</span>
                  </div>
                  <div class="min-w-0 flex-1">
                    <p class="text-sm font-medium line-clamp-2 group transition-colors duration-150" style="color: var(--color-text-primary);" @mouseenter="($event.target as HTMLElement).style.color = 'var(--color-primary)'" @mouseleave="($event.target as HTMLElement).style.color = 'var(--color-text-primary)'">{{ item.title }}</p>
                    <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">{{ item.time }}</span>
                  </div>
                </div>
              </a>
            </div>
          </div>

          <!-- Interaction Area -->
          <div class="p-5" style="background: var(--color-bg-card); border: 1px solid var(--color-border-subtle); border-radius: var(--radius-md);">
            <h3 class="text-sm font-semibold mb-4 pb-3 whitespace-nowrap" style="color: var(--color-text-primary); font-family: var(--font-display); letter-spacing: 0.02em; border-bottom: 1px solid var(--color-border-subtle);">互动区</h3>
            <div class="space-y-3">
              <!-- Comment count -->
              <div class="flex items-center justify-between">
                <div class="flex items-center gap-2">
                  <MessageCircle class="w-4 h-4" style="color: var(--color-text-tertiary);" />
                  <span class="text-sm whitespace-nowrap" style="color: var(--color-text-secondary);">{{ comments.length }} 条评论</span>
                </div>
              </div>
              <!-- Like button -->
              <button 
                class="w-full flex items-center justify-center gap-2 px-4 py-2.5 text-sm font-medium whitespace-nowrap transition-all duration-150"
                :style="{
                  background: isLiked ? 'var(--color-primary)' : 'var(--color-primary-50)',
                  color: isLiked ? 'white' : 'var(--color-primary)',
                  border: '1px solid var(--color-border)'
                }"
                @mouseenter="($event.target as HTMLElement).style.background = isLiked ? 'var(--color-primary-dark)' : 'var(--color-primary-100)'"
                @mouseleave="($event.target as HTMLElement).style.background = isLiked ? 'var(--color-primary)' : 'var(--color-primary-50)'"
                @click="handleLike"
              >
                <Heart class="w-4 h-4" :fill="isLiked ? 'currentColor' : 'none'" />
                <span>{{ isLiked ? '已点赞' : `+${likeCount} 点赞` }}</span>
              </button>
              <!-- Favorite button -->
              <button 
                class="w-full flex items-center justify-center gap-2 px-4 py-2.5 text-sm font-medium whitespace-nowrap transition-all duration-150"
                :style="{
                  background: userStore.isFavorite(currentArticle.id) ? 'var(--state-warning)' : 'var(--color-bg-elevated)',
                  color: userStore.isFavorite(currentArticle.id) ? 'white' : 'var(--color-text-secondary)',
                  border: '1px solid var(--color-border-subtle)'
                }"
                @mouseenter="($event.target as HTMLElement).style.background = userStore.isFavorite(currentArticle.id) ? '#e6a700' : 'var(--color-bg-tertiary)'"
                @mouseleave="($event.target as HTMLElement).style.background = userStore.isFavorite(currentArticle.id) ? 'var(--state-warning)' : 'var(--color-bg-elevated)'"
                @click="userStore.toggleFavorite(currentArticle.id)"
              >
                <Bookmark class="w-4 h-4" :fill="userStore.isFavorite(currentArticle.id) ? 'currentColor' : 'none'" />
                <span>{{ userStore.isFavorite(currentArticle.id) ? '已收藏' : '收藏' }}</span>
              </button>
              <!-- Share button -->
              <button class="w-full flex items-center justify-center gap-2 px-4 py-2.5 text-sm font-medium whitespace-nowrap transition-all duration-150" style="background: var(--color-bg-elevated); color: var(--color-text-secondary); border: 1px solid var(--color-border-subtle);" @mouseenter="($event.target as HTMLElement).style.background = 'var(--color-bg-tertiary)'" @mouseleave="($event.target as HTMLElement).style.background = 'var(--color-bg-elevated)'">
                <Share2 class="w-4 h-4" />
                <span>分享</span>
              </button>
            </div>
          </div>

          <!-- XP Summary -->
          <div class="p-5" style="background: var(--color-bg-card); border: 1px solid var(--color-border); border-radius: var(--radius-md);">
            <h3 class="text-sm font-semibold mb-4 pb-3 whitespace-nowrap" style="color: var(--color-primary); font-family: var(--font-display); letter-spacing: 0.02em; border-bottom: 1px solid var(--color-border-subtle);">XP 总览</h3>
            <div class="space-y-3">
              <div class="flex items-center justify-between">
                <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">本篇奖励</span>
                <span class="text-sm font-semibold whitespace-nowrap" style="color: var(--color-primary); font-family: var(--font-display);">+500 XP</span>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">今日已获</span>
                <span class="text-sm font-semibold whitespace-nowrap" style="color: var(--color-state-success); font-family: var(--font-display);">+1,200 XP</span>
              </div>
              <div class="pt-2" style="border-top: 1px solid var(--color-border-subtle);">
                <div class="flex items-center justify-between mb-1">
                  <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">距离 LV.13</span>
                  <span class="text-xs whitespace-nowrap" style="color: var(--color-text-tertiary);">800 / 2,000 XP</span>
                </div>
                <div class="w-full h-2 rounded-full overflow-hidden" style="background: var(--color-bg-tertiary);">
                  <div class="h-full rounded-full" style="width: 40%; background: linear-gradient(90deg, var(--color-primary), var(--color-primaryLight));"></div>
                </div>
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
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import { useThemeStore } from '@/stores/themeStore'
import { useArticleStore } from '@/stores/articleStore'
import { useUserStore } from '@/stores/userStore'
import { Eye, ArrowLeft, Heart, Share2, MessageCircle, Send, Bookmark } from 'lucide-vue-next'
import AppNavigation from '@/components/layout/AppNavigation.vue'
import AppFooter from '@/components/layout/AppFooter.vue'

const route = useRoute()
const themeStore = useThemeStore()
const articleStore = useArticleStore()
const userStore = useUserStore()
const readProgress = ref(0)
const isLiked = ref(false)
const likeCount = ref(100)
const newComment = ref('')

const comments = ref([
  { id: 1, author: '量子旅人', avatar: 'Q', content: 'GPT-5的出现确实会改变游戏开发的格局，期待更多AI辅助创作的佳作！', time: '2小时前', likes: 23 },
  { id: 2, author: '代码忍者', avatar: 'D', content: '作为独立开发者，终于可以把更多精力放在创意上了，技术门槛降低是好事。', time: '5小时前', likes: 15 },
  { id: 3, author: '像素画师', avatar: 'X', content: '担心AI会让美术从业者失业，但转念一想，AI也是工具，关键看谁在用。', time: '8小时前', likes: 42 }
])

const articleId = computed(() => route.params.id as string)

const currentArticle = computed(() => {
  return articleStore.articles.find(a => a.id === articleId.value) || articleStore.articles[0]
})

const articleTags = computed(() => currentArticle.value.tags)

const handleLike = () => {
  isLiked.value = !isLiked.value
  likeCount.value += isLiked.value ? 1 : -1
}

const handleAddComment = () => {
  if (newComment.value.trim()) {
    comments.value.unshift({
      id: Date.now(),
      author: userStore.playerName,
      avatar: userStore.playerName.charAt(0).toUpperCase(),
      content: newComment.value,
      time: '刚刚',
      likes: 0
    })
    newComment.value = ''
  }
}

const getCategoryColor = (category: string) => {
  const colors: Record<string, string> = {
    tech: 'var(--color-primary)',
    game: 'var(--color-secondary)',
    hardware: 'var(--state-info)',
    anime: 'var(--state-warning)'
  }
  return colors[category] || 'var(--color-primary)'
}

const getCategoryBg = (category: string) => {
  const colors: Record<string, string> = {
    tech: 'var(--color-primary-50)',
    game: 'var(--color-secondary-50)',
    hardware: 'rgba(0, 170, 255, 0.1)',
    anime: 'rgba(255, 170, 0, 0.1)'
  }
  return colors[category] || 'var(--color-primary-50)'
}

const getCategoryBorder = (category: string) => {
  const colors: Record<string, string> = {
    tech: 'var(--color-primary-200)',
    game: 'var(--color-secondary-200)',
    hardware: 'rgba(0, 170, 255, 0.3)',
    anime: 'rgba(255, 170, 0, 0.3)'
  }
  return colors[category] || 'var(--color-primary-200)'
}

const getCategoryName = (category: string) => {
  const names: Record<string, string> = {
    tech: '科技',
    game: '游戏',
    hardware: '硬件',
    anime: '二次元'
  }
  return names[category] || category
}

const relatedArticles = [
  { title: 'Unity 2026发布：AI原生工作流正式上线', time: '2小时前' },
  { title: '索尼公布PS6开发套件：光线追踪性能提升300%', time: '5小时前' },
  { title: '深度学习驱动的物理引擎：下一代游戏仿真', time: '1天前' },
  { title: 'AI辅助美术：从概念草图到3D模型的自动化管线', time: '2天前' }
]

let progressInterval: number | undefined

onMounted(() => {
  progressInterval = window.setInterval(() => {
    if (readProgress.value < 100) {
      readProgress.value += Math.random() * 5
      if (readProgress.value > 100) readProgress.value = 100
    }
  }, 1000)
})

onUnmounted(() => {
  if (progressInterval) {
    clearInterval(progressInterval)
  }
})
</script>