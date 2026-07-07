/**
 * English UI text locale — Common/Nav/Home
 * Organized by domain, mirrors zh-CN structure
 */
export default {
  // ==================== Category names (global, keyed by slug) ====================
  categories: {
    all: 'All',
    tech: 'Tech',
    ai_robot: 'AI Robot',
    ai_coding: 'AI Coding',
    hardware: 'Hardware',
    anime: 'Anime',
    godot: 'Godot',
  },

  // ==================== Common ====================
  common: {
    all: 'All',
    confirm: 'Confirm',
    cancel: 'Cancel',
    save: 'Save',
    close: 'Close',
    loading: 'Loading...',
    noData: 'No Data',
    back: 'Back',
    search: 'Search',
    minutes: '{count} minutes ago',
    hours: '{count} hours ago',
    days: '{count} days ago',
    justNow: 'Just now',
    errors: {
      networkError: 'Network error, please try again later.'
    },
    empty: {
      title: 'No Data',
      description: 'Nothing to display right now'
    },
    offline: {
      title: 'System Offline',
      description: 'Unable to connect to the intel server. Some features may be unavailable.'
    },
    searchEmpty: {
      title: 'No Results Found',
      description: 'Try searching with different keywords'
    }
  },

  // ==================== Navigation ====================
  nav: {
    home: 'Home',
    tech: 'Tech',
    game: 'Gaming',
    hardware: 'Hardware',
    anime: 'Anime',
    intelligenceCenter: 'Intel Center',
    intelligenceWarehouse: 'Intel Warehouse',
    aboutMe: 'About Me',
    profile: 'Profile',
    settings: 'Settings',
    logout: 'Log Out',
    switchTheme: 'Switch Theme'
  },

  // ==================== Hero Section ====================
  home: {
    hero: {
      title: {
        nexus: 'NEXUS DAILY',
        comiket: 'COMIKET STATION',
        ironcore: 'IRONCORE NETWORK'
      },
      subtitle: {
        nexus: 'Intel Updated',
        comiket: 'Latest Tech × Hot Games',
        ironcore: 'Intel Synced'
      },
      highlight: {
        nexus: 'DAILY INTEL UPDATED',
        comiket: 'Your daily dose of tech news, anime & gaming — all in one place!',
        ironcore: 'INTELLIGENCE SYNCHRONIZED'
      },
      hud: {
        nexus: '// SYSTEM ONLINE',
        comiket: 'Today\'s Special',
        ironcore: 'SYSTEM ONLINE // INTEL FEED ACTIVE'
      },
      btnPrimary: {
        nexus: 'Explore Now',
        comiket: 'Start Reading',
        ironcore: 'View Intel'
      },
      btnSecondary: {
        nexus: 'View Missions',
        comiket: 'Browse Manga',
        ironcore: 'Briefing'
      }
    },
    trending: {
      nexus: ['PS6 Launch', 'AI Art Revolution', 'Genshin 5.0', 'RTX5090 Review', 'Quantum Chip Breakthrough'],
      comiket: ['Genshin 5.0', 'RTX 6090', 'July Anime', 'Foldable Phones', 'AI Art', 'Gaming Laptops'],
      ironcore: ['Quantum Processor', 'RTX 5090 Review', 'Elden Ring DLC2', 'Gundam Movie', 'Starship Test 7']
    },
    player: {
      title: {
        nexus: 'Player Stats',
        comiket: 'My Data',
        ironcore: 'Dossier'
      },
      statsName: {
        nexus: ['Read', 'Achievements', 'Login Streak', 'Rank'],
        comiket: ['Read', 'Achievements', 'Level', 'Points'],
        ironcore: ['Intel Read', 'Medals', 'Streak', 'Rank']
      },
      playerName: {
        nexus: 'NEXUS Commander',
        comiket: 'COMIKET User',
        ironcore: 'Iron Agent'
      },
      playerTitle: {
        nexus: 'Cyber先锋 // Intel Hunter',
        comiket: 'Manga Master // Tech Lover',
        ironcore: 'Intel Expert // Field Veteran'
      }
    },
    quests: {
      title: {
        nexus: 'Daily Quests',
        comiket: 'Daily Challenges',
        ironcore: 'Active Missions'
      },
      tag: {
        nexus: 'DAILY QUEST',
        comiket: 'TASKS',
        ironcore: 'MISSION'
      },
      statusCompleted: 'Completed',
      statusNotStarted: 'Not Started',
      progress: 'Progress {progress}/{target}',
      questTitles: {
        read_articles: 'Read 3 Articles',
        login: 'Daily Check-in',
        comment: 'Comment on an Article'
      }
    },
    grid: {
      title: {
        nexus: 'Featured Intel',
        comiket: 'Manga Grid',
        ironcore: 'Encrypted Feed'
      },
      tag: {
        nexus: 'DAILY PICK',
        comiket: 'FEATURED',
        ironcore: 'CLASSIFIED & VERIFIED'
      }
    }
  }
}
