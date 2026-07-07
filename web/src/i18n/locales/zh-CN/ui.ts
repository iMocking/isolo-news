/**
 * 中文 UI 文本语言包 — 通用/导航/首页
 * 按功能域组织，仅包含界面文字
 */
export default {
  // ==================== 通用 ====================
  common: {
    all: '全部',
    confirm: '确认',
    cancel: '取消',
    save: '保存',
    close: '关闭',
    loading: '加载中...',
    noData: '暂无数据',
    back: '返回',
    search: '搜索',
    minutes: '{count} 分钟前',
    hours: '{count} 小时前',
    days: '{count} 天前',
    justNow: '刚刚',
    errors: {
      networkError: '网络错误，请稍后再试'
    },
    empty: {
      title: '暂无数据',
      description: '当前没有可显示的内容'
    },
    offline: {
      title: '情报系统离线',
      description: '无法连接到情报服务器，部分功能可能不可用'
    },
    searchEmpty: {
      title: '未找到结果',
      description: '尝试使用其他关键词搜索'
    }
  },

  // ==================== 导航 ====================
  nav: {
    home: '首页',
    tech: '科技资讯',
    game: '游戏情报',
    hardware: '硬件评测',
    anime: '二次元',
    intelligenceCenter: '情报中心',
    intelligenceWarehouse: '情报仓库',
    aboutMe: '关于我',
    profile: '个人中心',
    settings: '偏好设置',
    logout: '退出账号'
  },

  // ==================== 首页 Hero ====================
  home: {
    hero: {
      title: {
        nexus: 'NEXUS DAILY',
        comiket: 'COMIKET 情报站',
        ironcore: 'IRONCORE NETWORK'
      },
      subtitle: {
        nexus: '今日情报已更新',
        comiket: '最新科技 x 热门游戏',
        ironcore: '情报已同步'
      },
      highlight: {
        nexus: 'DAILY INTEL UPDATED',
        comiket: '一网打尽！每天精选最前沿的科技资讯和最热门的动漫游戏情报',
        ironcore: 'INTELLIGENCE SYNCHRONIZED'
      },
      hud: {
        nexus: '// SYSTEM ONLINE',
        comiket: '今日特集',
        ironcore: 'SYSTEM ONLINE // INTEL FEED ACTIVE'
      },
      btnPrimary: {
        nexus: '开始探索',
        comiket: '开始阅读',
        ironcore: '查看情报'
      },
      btnSecondary: {
        nexus: '查看任务',
        comiket: '浏览漫画',
        ironcore: '任务简报'
      }
    },
    trending: {
      nexus: ['PS6首发', 'AI绘画革命', '原神5.0', 'RTX5090评测', '量子芯片突破'],
      comiket: ['原神5.0', 'RTX 6090', '七月新番', '折叠屏', 'AI绘图', '游戏本横评'],
      ironcore: ['量子处理器突破', 'RTX 5090评测', '艾尔登法环DLC2', '高达新作剧场版', '星舰第七次试飞']
    },
    player: {
      title: {
        nexus: '玩家数据',
        comiket: '我的数据',
        ironcore: '作战档案'
      },
      statsName: {
        nexus: ['已阅读', '获得成就', '连续登录', '排名'],
        comiket: ['阅读', '成就', '等级', '积分'],
        ironcore: ['已读情报', '勋章', '连续签到', '位阶']
      },
      playerName: {
        nexus: 'NEXUS指挥官',
        comiket: 'COMIKET用户',
        ironcore: '铁核探员'
      },
      playerTitle: {
        nexus: '赛博先锋 // 资讯猎人',
        comiket: '漫画达人 // 科技爱好者',
        ironcore: '情报专家 // 战场老兵'
      }
    },
    quests: {
      title: {
        nexus: '今日任务',
        comiket: '每日挑战',
        ironcore: '当前任务'
      },
      tag: {
        nexus: 'DAILY QUEST',
        comiket: 'TASKS',
        ironcore: 'MISSION'
      },
      statusCompleted: '已完成',
      statusNotStarted: '未开始',
      progress: '进度 {progress}/{target}'
    },
    grid: {
      title: {
        nexus: '精选情报',
        comiket: '漫画面板网格',
        ironcore: '加密情报流'
      },
      tag: {
        nexus: 'DAILY PICK',
        comiket: 'FEATURED',
        ironcore: 'CLASSIFIED & VERIFIED'
      }
    }
  }
}
