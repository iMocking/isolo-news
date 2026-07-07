import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useUserStore } from '@/stores/userStore'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/pages/HomePage.vue')
  },
  // ==================== 文章路由（嵌套布局） ====================
  {
    path: '/article',
    component: () => import('@/layouts/ArticleLayout.vue'),
    children: [
      {
        path: '',
        redirect: { name: 'ArticleList' }
      },
      {
        path: 'list',
        name: 'ArticleList',
        component: () => import('@/pages/ArticleListPage.vue')
      },
      {
        path: 'detail/:id',
        name: 'ArticleDetail',
        component: () => import('@/pages/ArticleDetailPage.vue')
      }
    ]
  },

  // ==================== 认证路由（嵌套布局） ====================
  {
    path: '/auth',
    component: () => import('@/layouts/AuthLayout.vue'),
    children: [
      {
        path: '',
        redirect: { name: 'Login' }
      },
      {
        path: 'login',
        name: 'Login',
        component: () => import('@/pages/LoginPage.vue')
      },
      {
        path: 'register',
        name: 'Register',
        component: () => import('@/pages/RegisterPage.vue')
      }
    ]
  },

  // ==================== 其他页面 ====================
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
  scrollBehavior() {
    return { top: 0 }
  }
})

/** 路由守卫：受保护页面需登录 */
router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    const token = localStorage.getItem('access_token')
    if (!token) {
      next({ name: 'Login', query: { redirect: to.fullPath } })
      return
    }
  }
  next()
})

export default router
