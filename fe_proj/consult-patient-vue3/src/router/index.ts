import { createRouter, createWebHistory } from 'vue-router'

// createWebHashHistory 哈希路由 会有 # 号
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL), // 路由模式
  routes: [],
})

export default router
