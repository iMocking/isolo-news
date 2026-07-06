# NEXUS Daily — 赛博朋克科技资讯聚合平台

> **赛博资讯 · 游戏化阅读 · 多主题沉浸体验**

NEXUS Daily（代号 `isolo-news`）是一个面向科技爱好者、游戏玩家和动漫迷的科技资讯聚合平台。后端通过 RSS 采集器自动聚合多源科技资讯，前端以**游戏化（Game-Fi）** 方式呈现，包含 XP 经验值、等级、成就、每日任务等机制，增强用户参与感。支持**三种视觉主题**自由切换，满足不同审美偏好。

---

## 核心特性

### 📡 智能资讯聚合
- 内置 **10 个 RSS 数据源**，覆盖 AI 前沿、机器人、编程、硬件、二次元、Godot 游戏等领域
- **定时调度采集**，基于 cron 表达式自动抓取最新资讯
- 文章去重、自动分类、智能精选

### 🎮 游戏化阅读体验
- **XP 经验值系统** — 阅读文章获得经验，按阅读时长分级奖励
- **等级 & 称号** — 随经验增长自动升级，解锁独特称号
- **成就系统** — 5 大成就等你解锁（初次阅读、连续登录、评论之星等）
- **每日任务** — 3 个日常任务驱动内容消费
- **周排行榜** — 与社区玩家竞技，争夺榜首

### 🎨 三套沉浸式主题
| 主题 | 风格 | 特点 |
|------|------|------|
| **NEXUS** | 赛博游戏 | 霓虹青蓝 + 深黑背景，像素感装饰 |
| **COMIKET** | 清爽漫画 | 暖白底 + 活力橙红，圆润柔和 |
| **IRONCORE** | 暗黑机甲 | 暗灰金属 + 琥珀色点缀，硬朗工业风 |

### 🔐 完善的权限体系
- **RBAC 角色模型**（guest → user → editor → admin）
- **JWT 双 Token 认证**（Access + Refresh）
- 邮箱验证码注册，保障账号安全

### 🌐 国际化
- 完整的中英文双语支持
- 语言偏好持久化存储

---

## 技术栈

### 后端

| 分类 | 技术 |
|------|------|
| **语言** | Go 1.25 |
| **Web 框架** | Gin v1.12 |
| **ORM** | Ent v0.14 (entgo.io) |
| **数据库** | PostgreSQL |
| **缓存** | Redis（预留） |
| **权限** | Casbin v2（RBAC） |
| **认证** | JWT (golang-jwt/v5) |
| **RSS** | gofeed v1.3 |
| **定时任务** | robfig/cron v3 |
| **配置** | Viper v1.21 |
| **邮件** | gomail v2 |

### 前端

| 分类 | 技术 |
|------|------|
| **框架** | Vue 3 + TypeScript (Composition API) |
| **构建** | Vite 5 |
| **状态管理** | Pinia v2 |
| **路由** | Vue Router v4 |
| **样式** | Tailwind CSS 3 + CSS 变量动态主题 |
| **国际化** | vue-i18n v9 |
| **HTTP** | Axios v1.7 |
| **图标** | Lucide Vue Next |

---

## 快速开始

### 环境要求

- **Go** >= 1.25
- **Node.js** >= 20
- **PostgreSQL** >= 15
- **Redis**（可选，预留）

### 1. 克隆项目

```bash
git clone https://github.com/your-org/isolo-news.git
cd isolo-news
```

### 2. 后端配置与启动

```bash
cd server

# 复制配置文件并修改数据库连接等参数
cp config.yaml config.local.yaml
# 编辑 config.local.yaml，填入 PostgreSQL 连接信息

# 启动 HTTP 服务（自动执行数据库迁移、种子数据初始化）
go run cmd/server/main.go

# 可选：单独执行种子数据采集
go run cmd/seed/main.go
```

### 3. 前端配置与启动

```bash
cd web

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

访问 `http://localhost:5173` 即可看到应用。

---

## 项目结构

```
isolo-news/
├── server/                          # 后端服务（Go + Gin）
│   ├── cmd/                         # 入口命令
│   │   ├── server/main.go           # HTTP 服务主入口
│   │   └── seed/main.go             # 种子数据 + 立即采集
│   ├── internal/                    # 内部业务逻辑
│   │   ├── config/                  # 配置加载
│   │   ├── handler/                 # HTTP 处理器（控制器层）
│   │   ├── middleware/              # Gin 中间件
│   │   ├── model/                   # 统一数据模型
│   │   ├── repository/ent/          # Ent ORM Schema + 生成代码
│   │   ├── router/                  # 路由注册
│   │   ├── seeder/                  # 种子数据初始化
│   │   └── service/                 # 业务服务层
│   │       └── collector/           # RSS 采集器 + 调度器
│   └── pkg/                         # 公共工具包
│       ├── casbin/                  # 权限模型 + 策略
│       └── utils/                   # 工具函数（JWT/密码/响应）
│
├── web/                             # 前端应用（Vue 3 + TS）
│   └── src/
│       ├── api/                     # API 请求层（Axios）
│       ├── components/              # 组件
│       │   ├── base/                # 基础组件
│       │   ├── layout/              # 布局组件
│       │   └── business/            # 业务组件
│       ├── config/                  # 主题配置
│       ├── i18n/                    # 国际化（中/英）
│       ├── lib/                     # 工具函数
│       ├── pages/                   # 页面组件（9 个路由页面）
│       ├── router/                  # Vue Router 路由定义
│       ├── stores/                  # Pinia 状态管理
│       ├── styles/                  # 全局样式
│       └── types/                   # TypeScript 类型定义
│
├── README.md
└── .trae/documents/                 # 产品需求 & 技术架构文档
```

---

## 功能模块详解

### 资讯系统
- **多源聚合**：自动从 10 个 RSS 源采集资讯
- **分类浏览**：AI 前沿、AI 机器人、AI 编程、AI 硬件、二次元、Godot 游戏
- **智能精选**：首页展示精选推荐内容
- **全文检索**：支持标题/内容关键词搜索
- **互动功能**：评论、点赞、收藏

### 用户系统
- **注册/登录**：邮箱注册 + 密码登录
- **Token 认证**：Access Token + Refresh Token 双机制
- **个人主页**：展示等级、经验、称号、游戏化数据
- **成就/任务**：追踪进度，解锁奖励

### 采集调度
- **定时采集**：按分类配置不同的采集频率（30 分钟 ~ 4 小时）
- **自动去重**：基于标题和来源 URL 去重
- **失败重试**：采集失败自动记录，支持降级

### 权限控制
- **4 级角色体系**：guest（游客）< user（用户）< editor（编辑）< admin（管理员）
- **接口级权限**：基于 Casbin 的细粒度权限校验

---

## API 文档索引

### 公开接口（无需认证）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/health` | 健康检查 |
| GET | `/api/v1/articles` | 资讯列表（分页/分类/搜索/排序） |
| GET | `/api/v1/articles/featured` | 首页精选资讯 |
| GET | `/api/v1/articles/:id` | 文章详情 |
| GET | `/api/v1/articles/:id/comments` | 文章评论列表 |
| GET | `/api/v1/categories` | 分类列表 |
| GET | `/api/v1/tags/hot` | 热门标签 |
| GET | `/api/v1/topics/trending` | 热门话题 |
| GET | `/api/v1/leaderboard/weekly` | 周排行榜 |
| POST | `/api/v1/auth/login` | 登录 |
| POST | `/api/v1/auth/register` | 注册 |
| POST | `/api/v1/auth/email/code` | 发送邮箱验证码 |
| POST | `/api/v1/auth/email/verify` | 验证邮箱 |
| POST | `/api/v1/auth/refresh` | 刷新 Token |

### 用户接口（需登录）

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/v1/articles/:id/comments` | 创建评论 |
| POST | `/api/v1/articles/:id/like` | 切换点赞 |
| POST | `/api/v1/articles/:id/favorite` | 切换收藏 |
| POST | `/api/v1/articles/:id/read` | 记录阅读（获得 XP） |
| GET | `/api/v1/user/profile` | 获取个人信息 |
| PUT | `/api/v1/user/profile` | 更新个人信息 |
| GET | `/api/v1/user/favorites` | 收藏列表 |
| GET | `/api/v1/user/achievements` | 成就列表 |
| GET | `/api/v1/user/quests` | 每日任务 |
| GET | `/api/v1/user/stats` | 用户统计 |

---

## 游戏化系统

### XP 经验值
阅读文章根据文章阅读时长分级获得 XP：

| 阅读时长 | XP 奖励 |
|----------|---------|
| < 1 分钟 | 50 XP |
| 1~3 分钟 | 100 XP |
| 3~5 分钟 | 200 XP |
| > 5 分钟 | 300 XP |

### 成就系统
| 成就 | 条件 |
|------|------|
| 🏆 初次阅读 | 阅读第一篇资讯 |
| 🔥 连续登录 | 连续登录 7 天 |
| 🤖 科技爱好者 | 阅读 10 篇 AI 分类文章 |
| 💬 评论之星 | 发表 5 条评论 |
| 📌 收藏家 | 收藏 10 篇文章 |

### 每日任务
| 任务 | 目标 |
|------|------|
| 📰 资讯达人 | 阅读 3 篇资讯 |
| ✅ 每日签到 | 登录并阅读一篇文章 |
| ✍️ 发表评论 | 评论一篇评测 |

---

## 主题系统

平台内置三套完整视觉主题，通过 CSS 变量动态切换：

```typescript
// 主题配置位于 web/src/config/themes.ts
interface ThemeConfig {
  name: string          // 主题名称
  colors: {             // 色板（主色/辅色/背景/文字等）
    primary: string
    secondary: string
    // ...
  }
  fonts: {              // 字体族
    display: string
    body: string
  }
  borderRadius: string  // 全局圆角
}
```

主题切换通过 `themeStore` 管理，持久化存储到 `localStorage`。

---

## 开发指南

### 后端开发

```bash
# 代码生成（Ent Schema 变更后执行）
cd server
go generate ./...

# 运行测试
go test ./...

# 编译
go build -o server.exe cmd/server/main.go
go build -o seed.exe cmd/seed/main.go
```

### 前端开发

```bash
cd web

# 开发模式
npm run dev

# 构建生产版本
npm run build

# 代码检查
npm run lint
```

### 数据库迁移

Ent ORM 在服务启动时会自动执行迁移（`AutoMigrate`），无需手动执行 SQL。

---

## 部署指南

### 后端部署

1. 编译为可执行文件
   ```bash
   cd server
   go build -o server.exe cmd/server/main.go
   ```
2. 配置环境变量或 `config.yaml`
3. 运行 `./server.exe`

### 前端部署

1. 构建静态文件
   ```bash
   cd web
   npm run build   # 输出到 dist/
   ```
2. 将 `dist/` 目录部署到 Nginx / CDN
3. 配置反向代理到后端 API

---

## RSS 数据源一览

| 数据源 | 分类 | 采集频率 |
|--------|------|----------|
| Hacker News | AI前沿 | 每 30 分钟 |
| TechCrunch AI | AI前沿 | 每 30 分钟 |
| MIT Tech Review - AI | AI前沿 | 每 30 分钟 |
| IEEE Spectrum Robotics | AI机器人 | 每小时 |
| Dev.to AI | AI编程 | 每 30 分钟 |
| GitHub Trending | AI编程 | 每 30 分钟 |
| Tom's Hardware | AI硬件 | 每 2 小时 |
| Anime News Network | 二次元 | 每 3 小时 |
| Godot Blog | Godot游戏 | 每 4 小时 |
| Reddit r/godot | Godot游戏 | 每 4 小时 |

---

## 贡献指南

1. Fork 本仓库
2. 创建特性分支：`git checkout -b feature/amazing-feature`
3. 提交更改：`git commit -m 'feat: 添加某个精彩特性'`
4. 推送到分支：`git push origin feature/amazing-feature`
5. 提交 Pull Request

### 代码规范

- 后端：遵循 Go 标准项目布局，使用 `gofmt` 格式化
- 前端：遵循 Vue 3 Composition API + `<script setup>` 风格
- 提交信息：遵循 [Conventional Commits](https://www.conventionalcommits.org/) 规范

---

## 许可证

本项目基于 MIT 许可证开源。

---

> **NEXUS Daily** — 连接未来，每日资讯。
