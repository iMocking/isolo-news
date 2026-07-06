import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

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
  routes
})

export default router