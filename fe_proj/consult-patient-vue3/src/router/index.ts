import { createRouter, createWebHistory } from 'vue-router'

// createWebHashHistory 哈希路由 会有 # 号
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL), // 路由模式
  routes: [
    {
      path: '/login',
      component: () => import('@/views/Login/index.vue'),
    },
    {
      path: '/',
      redirect: '/home',
      component: () => import('@/views/Layout/index.vue'),
      children: [
        // 二级路由
        { path: '/home', component: () => import('@/views/Home/index.vue') },
        { path: '/article', component: () => import('@/views/Article/index.vue') },
        { path: '/notify', component: () => import('@/views/Notify/index.vue') },
        { path: '/user', component: () => import('@/views/User/index.vue') },
      ],
    },
  ],
})

export default router
