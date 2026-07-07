/**
 * English UI text locale — Page-specific text
 * Mirrors zh-CN/pages.ts structure
 */
export default {
  // ==================== Article List ====================
  articleList: {
    title: 'Mission List',
    subtitle: 'MISSION SELECT',
    total: '{count} intel reports',
    searchPlaceholder: 'Search intel keywords...',
    startMission: 'Start Mission',
    comments: '{count} comments',
    prev: 'PREV',
    next: 'NEXT STAGE →',
    pagination: {
      perPage: 'Show',
      unit: 'per page',
    },
    filter: {
      title: 'Filters',
      timeRange: 'Time Range',
      allTime: 'All Time',
      within24h: 'Within 24h',
      withinWeek: 'Within Week',
      withinMonth: 'Within Month',
      sortBy: 'Sort By',
      latest: 'Latest',
      mostComments: 'Most Comments',
      hottest: 'Hottest',
      source: 'Source',
      allSources: 'All Sources',
      official: 'Official',
      community: 'Community',
      industry: 'Industry'
    },
    sidebar: {
      hotTags: 'Hot Tags',
      weeklyRank: 'Weekly Rank'
    },
    categories: {
      all: 'All',
      tech: 'Tech',
      game: 'Gaming',
      hardware: 'Hardware',
      anime: 'Anime'
    }
  },

  // ==================== Article Detail ====================
  articleDetail: {
    breadcrumb: {
      home: 'Home',
      list: 'Tech',
      current: 'Mission Detail'
    },
    reading: 'Read {count}',
    readTime: '{count} min read',
    xpProgress: 'XP Progress',
    readProgress: 'Reading {progress}%',
    readReward: 'Finish for +{xp} XP',
    authorBio: 'Quantum Journalist · Alice',
    authorDesc: 'Senior tech journalist covering AI & gaming. Veteran of E3 and TGS with unique insights into tech trends.',
    backToList: 'Back to List',
    comment: {
      title: 'Comments ({count})',
      placeholder: 'Write a comment...',
      submit: 'Post Comment',
      reply: 'Reply',
      count: '{count} comments',
      liked: 'Liked',
      like: '+{count} likes',
      favorited: 'Saved',
      favorite: 'Save',
      share: 'Share',
      loginRequired: 'Please log in to comment.',
      postSuccess: 'Comment posted!',
      postFail: 'Failed to post comment. Please try again.'
    },
    sidebar: {
      related: 'Related Missions',
      interactions: 'Interactions',
      comment: 'Comment',
      like: 'Like',
      favorite: 'Favorite',
      share: 'Share',
      xpOverview: 'XP Overview',
      xpThisArticle: 'This Article',
      xpToday: 'Today Earned',
      xpToNextLevel: 'To LV.{level}'
    }
  },

  // ==================== Search ====================
  search: {
    title: 'Search Terminal',
    subtitle: 'SEARCH TERMINAL',
    placeholder: 'Enter search keywords...',
    results: '{count} results',
    scanResults: 'Scan Results',
    source: 'Source: ',
    sidebar: {
      suggestions: 'Suggestions',
      hotSearches: 'Trending'
    },
    filters: {
      time: 'Time (All)',
      category: 'Category (All)',
      sort: 'Sort (Relevance)'
    }
  },

  // ==================== Profile ====================
  profile: {
    stats: ['Intel Read', 'Achievements', 'Login Streak', 'Global Rank', 'Comments', 'Saved'],
    achievements: [
      { name: 'Rookie Agent', desc: 'First intel read' },
      { name: 'Streak Master', desc: '30-day login streak' },
      { name: 'Comment Star', desc: '100 comments posted' },
      { name: 'Collector', desc: '50 intel saved' }
    ],
    panels: {
      ability: 'Ability Panel',
      achievements: 'Achievements ({count})',
      favorites: 'My Favorites',
      favoriteCount: '{count} articles'
    },
    empty: {
      title: 'No saved intel yet',
      desc: 'Click the save button on article detail pages to add favorites'
    },
    logout: 'Log Out',
    playerName: {
      nexus: 'NEXUS Commander',
      comiket: 'COMIKET User',
      ironcore: 'Iron Agent'
    },
    playerTitle: {
      nexus: 'Cyber Pioneer // Intel Hunter',
      comiket: 'Manga Master // Tech Lover',
      ironcore: 'Intel Expert // Field Veteran'
    }
  },

  // ==================== Settings ====================
  settings: {
    subtitle: '// System Config',
    notifications: {
      title: 'Notifications',
      push: 'Push Notifications',
      email: 'Email Notifications',
      inApp: 'In-App Notifications'
    },
    profileEdit: {
      title: 'Edit Profile',
      desc: 'Change avatar, nickname, bio',
      nickname: 'Nickname',
      fieldTitle: 'Title',
      tags: 'Tags (comma separated)',
      saveSuccess: 'Saved successfully',
      saveFail: 'Save failed'
    },
    theme: {
      title: 'Theme',
      label: 'THEME',
      nexus: 'NEXUS Cyberpunk',
      comiket: 'COMIKET Anime',
      ironcore: 'IRONCORE Mecha'
    },
    privacy: {
      title: 'Privacy',
      desc: 'Data visibility & privacy'
    },
    security: {
      title: 'Security',
      desc: 'Password & authentication'
    },
    system: {
      title: 'Preferences',
      info: 'System Info',
      clientVersion: 'Client Version',
      locale: 'Language',
      dataSync: 'Data Sync',
      connected: 'Connected',
      localCache: 'Local Cache',
      compactMode: 'Compact Mode',
      reduceAnimation: 'Reduce Animation',
      highContrast: 'High Contrast'
    }
  },

  // ==================== Login ====================
  login: {
    systemText: {
      nexus: 'SYSTEM LOGIN',
      comiket: 'COMIC INFO',
      ironcore: 'COMMAND CENTER'
    },
    tabs: {
      login: 'Login',
      register: 'Register'
    },
    form: {
      username: 'Player ID',
      password: 'Password',
      remember: 'Remember Me',
      forgot: 'Forgot Password?'
    },
    buttons: {
      nexus: 'CONNECT // LOGIN',
      comiket: 'START READING',
      ironcore: 'AUTHORIZE'
    },
    divider: 'Quick Access',
    backHome: 'Back to Home',
    statusOnline: 'SYS.STATUS: ONLINE',
    connSecure: 'CONN: SECURE',
    errors: {
      emptyEmailPassword: 'Please enter email and password',
      loginFail: 'Login failed. Please check your email and password.',
      emptyEmail: 'Please enter your email first',
      sendCodeFail: 'Failed to send verification code',
      fillAllFields: 'Please fill in all registration fields',
      passwordMismatch: 'Passwords do not match',
      passwordTooShort: 'Password must be at least 6 characters',
      enterCode: 'Please get a verification code first',
      codeError: 'Code is invalid or expired',
      registerFail: 'Registration failed',
      registerSuccess: 'Registration successful! Please log in with your email.'
    }
  },

  // ==================== About ====================
  about: {
    subtitle: '// About System',
    intro: 'NEXUS DAILY is a cyberpunk-themed tech & anime intel platform, blending tech news, game reviews, hardware analysis, and anime culture for an immersive reading experience.',
    sections: {
      platform: 'Platform',
      team: 'Team',
      tech: 'Tech Stack',
      contact: 'Contact'
    },
    platformDesc: 'Founded in 2024, NEXUS DAILY is dedicated to building the most immersive cyberpunk-style news platform. We combine gamification mechanisms with cutting-edge tech reporting for a unique reading experience.',
    teamMembers: [
      { name: 'Cyber Observer', role: 'Chief Tech Analyst' },
      { name: 'Game Frontline', role: 'Head of Gaming Intel' },
      { name: 'Anime Probe', role: 'Cultural Analyst' }
    ]
  },

  // ==================== 404 ====================
  notFound: {
    subtitle: '// Intel Not Found',
    message: 'This intel line seems to be disconnected...'
  }
}
