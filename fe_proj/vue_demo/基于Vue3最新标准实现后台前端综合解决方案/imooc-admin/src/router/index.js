import { createRouter, createWebHashHistory } from 'vue-router'
/**
 * 公开路由表
 */
const publicRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/IndexView')
  },
  {
    path: '/',
    component: () => import('@/layout/index')
  },
  {
    path: '/chart',
    component: () => import('@/views/ChartView')
  },
  {
    path: '/chart2',
    component: () => import('@/views/Chart2View')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes: publicRoutes
})

export default router
