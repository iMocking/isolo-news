import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useUserStore } from '@/stores/userStore'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/pages/HomePage.vue')
  },
  {
    path: '/list',
    name: 'ArticleList',
    component: () => import('@/pages/ArticleListPage.vue')
  },
  {
    path: '/article/:id',
    name: 'ArticleDetail',
    component: () => import('@/pages/ArticleDetailPage.vue')
  },
  {
    path: '/search',
    name: 'SearchResults',
    component: () => import('@/pages/SearchResultsPage.vue')
  },
  {
    path: '/profile',
    name: 'UserProfile',
    meta: { requiresAuth: true },
    component: () => import('@/pages/UserProfilePage.vue')
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('@/pages/AboutPage.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/pages/LoginPage.vue')
  },
  {
    path: '/settings',
    name: 'Settings',
    meta: { requiresAuth: true },
    component: () => import('@/pages/SettingsPage.vue')
  },
  {
    path: '/404',
    name: 'NotFound',
    component: () => import('@/pages/NotFoundPage.vue')
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/404'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  /** 每次页面跳转都滚动到顶部 */
  scrollBehavior() {
    return { top: 0 }
  }
})

/** 路由守卫：受保护页面需登录 */
router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    // 由于 pinia 在路由守卫中可能尚未完全初始化，使用同步检查 localStorage
    const token = localStorage.getItem('access_token')
    if (!token) {
      next({ name: 'Login', query: { redirect: to.fullPath } })
      return
    }
  }
  next()
})

export default router