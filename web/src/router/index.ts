import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('@/layouts/MainLayout.vue'),
    redirect: { name: 'Home' },
    children: [
      // ==================== 首页 ====================
       {
    path: '/home',
    name: 'Home',
    component: () => import('@/pages/home/index.vue')
  },
  // ==================== 文章路由（嵌套布局） ====================
  {
    path: 'article',
    name: 'Article',
    component: () => import('@/layouts/ArticleLayout.vue'),
    children: [
      {
        path: '',
        redirect: { name: 'ArticleList' }
      },
      {
        path: 'list',
        name: 'ArticleList',
        component: () => import('@/pages/list/index.vue')
      },
      {
        path: 'detail/:id',
        name: 'ArticleDetail',
        component: () => import('@/pages/detail/index.vue')
      }
    ]
  },
  // ==================== 认证路由（嵌套布局） ====================
  {
    path: 'auth',
    component: () => import('@/layouts/AuthLayout.vue'),
    children: [
      {
        path: '',
        redirect: { name: 'Login' }
      },
      {
        path: 'login',
        name: 'Login',
        component: () => import('@/pages/login/index.vue')
      },
      {
        path: 'register',
        name: 'Register',
        component: () => import('@/pages/register/index.vue')
      }
    ]
  },
  // ==================== 其他页面 ====================
  {
    path: 'search',
    name: 'SearchResults',
    component: () => import('@/pages/search/index.vue')
  },
  {
    path: 'profile',
    name: 'UserProfile',
    meta: { requiresAuth: true },
    component: () => import('@/pages/profile/index.vue')
  },
  {
    path: 'about',
    name: 'About',
    component: () => import('@/pages/about/index.vue')
  },
  {
    path: 'settings',
    name: 'Settings',
    meta: { requiresAuth: true },
    component: () => import('@/pages/settings/index.vue')
  },
  {
    path: '404',
    name: 'NotFound',
    component: () => import('@/pages/404.vue')
  },
    ]
  },
 
  // 通配路由放在根层级，确保在所有子路由之后匹配
  {
    path: '/:pathMatch(.*)*',
    redirect: { name: 'NotFound' }
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
