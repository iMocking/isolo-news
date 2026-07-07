/**
 * 中文 UI 文本语言包 — 页面专用文本
 * 包含文章列表、文章详情、搜索、用户中心、设置、登录、关于、404
 */
export default {
  // ==================== 文章列表页 ====================
  articleList: {
    title: '任务列表',
    subtitle: 'MISSION SELECT',
    total: '共 {count} 条情报',
    searchPlaceholder: '搜索情报关键词...',
    startMission: '开始任务',
    comments: '{count} 条评论',
    prev: 'PREV',
    next: 'NEXT STAGE →',
    pagination: {
      perPage: '每页',
      unit: '条',
    },
    filter: {
      title: '筛选条件',
      timeRange: '时间范围',
      allTime: '全部时间',
      within24h: '24小时内',
      withinWeek: '一周内',
      withinMonth: '一个月内',
      sortBy: '热度排序',
      latest: '最新发布',
      mostComments: '最多评论',
      hottest: '最高热度',
      source: '来源',
      allSources: '全部来源',
      official: '官方渠道',
      community: '社区投稿',
      industry: '行业媒体'
    },
    sidebar: {
      hotTags: '热门标签',
      weeklyRank: '本周排行'
    },
    categories: {
      all: '全部',
      tech: '科技',
      ai_robot: 'AI机器人',
      ai_coding: 'AI编程',
      hardware: '硬件',
      anime: '二次元',
      godot: 'Godot'
    },
    read: '已读'
  },

  // ==================== 文章详情页 ====================
  articleDetail: {
    breadcrumb: {
      home: '首页',
      list: '科技资讯',
      current: '任务详情'
    },
    reading: '阅读 {count}',
    readTime: '{count} 分钟阅读',
    xpProgress: 'XP 奖励进度',
    readProgress: '阅读进度 {progress}%',
    readReward: '读完可获得 +{xp} XP',
    authorBio: '量子记者 · Alice',
    authorDesc: '资深科技记者，专注AI与游戏产业交叉领域。曾报道多届E3与TGS，对技术趋势有独到见解。',
    backToList: '返回任务列表',
    comment: {
      title: '评论区 ({count})',
      placeholder: '写下你的评论...',
      submit: '发表评论',
      reply: '回复',
      count: '{count} 条评论',
      liked: '已点赞',
      like: '+{count} 点赞',
      favorited: '已收藏',
      favorite: '收藏',
      share: '分享',
      loginRequired: '请先登录后再发表评论',
      postSuccess: '评论发表成功',
      postFail: '评论发表失败，请稍后重试'
    },
    sidebar: {
      related: '相关任务',
      interactions: '互动区',
      comment: '评论',
      like: '点赞',
      favorite: '收藏',
      share: '分享',
      xpOverview: 'XP 总览',
      xpThisArticle: '本篇奖励',
      xpToday: '今日已获',
      xpToNextLevel: '距离 LV.{level}',
      currentLevel: '当前等级'
    }
  },

  // ==================== 搜索页 ====================
  search: {
    title: '搜索终端',
    subtitle: 'SEARCH TERMINAL',
    placeholder: '输入搜索关键词...',
    results: '{count} 个结果',
    scanResults: '扫描结果',
    source: '来源: ',
    sidebar: {
      suggestions: '搜索建议',
      hotSearches: '热门搜索'
    },
    filters: {
      time: '时间 (全部)',
      category: '分类 (全部)',
      sort: '排序 (相关度)'
    }
  },

  // ==================== 用户中心 ====================
  profile: {
    stats: ['已读情报', '获得成就', '连续签到', '全站排名', '发表评论', '收藏情报'],
    achievements: [
      { name: '初入情报网', desc: '完成首次情报阅读' },
      { name: '连续签到达人', desc: '连续签到30天' },
      { name: '评论之星', desc: '发表100条评论' },
      { name: '收藏家', desc: '收藏50篇情报' }
    ],
    panels: {
      ability: '能力面板',
      achievements: '成就徽章 ({count})',
      favorites: '我的收藏',
      favoriteCount: '{count} 篇情报'
    },
    empty: {
      title: '还没有收藏任何情报',
      desc: '点击文章详情页的收藏按钮添加收藏'
    },
    logout: '退出登录',
    playerName: {
      nexus: 'NEXUS指挥官',
      comiket: 'COMIKET用户',
      ironcore: '铁核探员'
    },
    unitDays: '天',
    playerTitle: {
      nexus: '赛博先锋 // 资讯猎人',
      comiket: '漫画达人 // 科技爱好者',
      ironcore: '情报专家 // 战场老兵'
    }
  },

  // ==================== 设置页 ====================
  settings: {
    subtitle: '// 系统设置',
    pageTitle: {
      nexus: 'SYSTEM CONFIG',
      comiket: '系统配置',
      ironcore: 'COMMAND CONFIG'
    },
    pageTag: {
      nexus: '// 系统参数调整',
      comiket: '// 个性化设置',
      ironcore: '// 作战参数配置'
    },
    pageSubtitle: {
      nexus: '调整系统参数、通知偏好与个人资料',
      comiket: '个性化你的阅读体验与界面主题',
      ironcore: '配置作战参数与系统偏好'
    },
    notifications: {
      title: '通知设置',
      push: '推送通知',
      email: '邮件通知',
      inApp: '应用内通知'
    },
    profileEdit: {
      title: '个人信息编辑',
      desc: '修改头像、昵称、签名',
      nickname: '昵称',
      fieldTitle: '头衔',
      tags: '标签（逗号分隔）',
      saveSuccess: '保存成功',
      saveFail: '保存失败'
    },
    theme: {
      title: '主题切换',
      label: 'THEME',
      nexus: 'NEXUS 赛博朋克',
      comiket: 'COMIKET 清新动漫',
      ironcore: 'IRONCORE 硬核机甲'
    },
    privacy: {
      title: '隐私设置',
      desc: '数据可见性与隐私保护'
    },
    security: {
      title: '账户安全',
      desc: '密码修改、登录验证'
    },
    system: {
      title: '系统偏好',
      info: '系统信息',
      clientVersion: '客户端版本',
      locale: '语言环境',
      dataSync: '数据同步',
      connected: '已连接',
      localCache: '本地缓存',
      compactMode: '紧凑模式',
      reduceAnimation: '减少动画',
      highContrast: '高对比度'
    }
  },

  // ==================== 登录页 ====================
  login: {
    systemText: {
      nexus: 'SYSTEM LOGIN // 系统登录',
      comiket: '动漫资讯 // COMIC INFO',
      ironcore: 'COMMAND CENTER // 指挥中心'
    },
    tabs: {
      login: '登录',
      register: '注册'
    },
    form: {
      username: '玩家ID',
      password: '密码',
      remember: '记住我',
      forgot: '忘记密码?'
    },
    buttons: {
      nexus: '接入系统 // LOGIN',
      comiket: '开始阅读 // START',
      ironcore: 'AUTHORIZE // 授权'
    },
    divider: '快速接入',
    backHome: '返回首页',
    statusOnline: 'SYS.STATUS: ONLINE',
    connSecure: 'CONN: SECURE',
    errors: {
      emptyEmailPassword: '请输入邮箱和密码',
      loginFail: '登录失败，请检查邮箱和密码',
      emptyEmail: '请先输入邮箱',
      sendCodeFail: '验证码发送失败',
      fillAllFields: '请填写所有注册字段',
      passwordMismatch: '两次输入的密码不一致',
      passwordTooShort: '密码长度不能少于6位',
      enterCode: '请先获取邮箱验证码并填写',
      codeError: '验证码错误或已过期',
      registerFail: '注册失败',
      registerSuccess: '注册成功，请使用注册的邮箱登录'
    }
  },

  // ==================== 关于页 ====================
  about: {
    subtitle: '// 关于系统',
    pageTitle: {
      nexus: 'ABOUT NEXUS',
      comiket: 'ABOUT COMIKET',
      ironcore: 'ABOUT IRONCORE'
    },
    pageTag: {
      nexus: '// SYSTEM INFO',
      comiket: '// 关于我们',
      ironcore: '// UNIT PROFILE'
    },
    features: [
      { title: '前沿科技', desc: '聚合全球最新科技动态与深度评测' },
      { title: '游戏情报', desc: '覆盖主机、PC、手游全平台资讯' },
      { title: '二次元文化', desc: '动漫、漫画、轻小说文化速递' },
      { title: '社区互动', desc: '评论、收藏、点赞与成就系统' }
    ],
    intro: 'NEXUS DAILY 是一个赛博朋克风格的科技与二次元情报平台，融合了科技资讯、游戏评测、硬件测评和二次元文化等内容，致力于为用户提供沉浸式的资讯阅读体验。',
    sections: {
      platform: '平台简介',
      team: '情报团队',
      tech: '技术栈',
      contact: '联系我们'
    },
    platformDesc: 'NEXUS DAILY 成立于2024年，致力于打造全球最具沉浸感的赛博朋克风格资讯平台。我们融合游戏化机制与前沿科技报道，为用户带来全新的阅读体验。',
    teamMembers: [
      { name: '赛博观察者', role: '首席科技情报员' },
      { name: '游戏前线', role: '游戏情报主管' },
      { name: '二次元探针', role: '动漫文化分析师' }
    ]
  },

  // ==================== 404 ====================
  notFound: {
    subtitle: '// 情报未找到',
    message: '这条情报线似乎已断开...'
  }
}
