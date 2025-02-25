import { useUserStore } from '@/stores'
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

// 全局的前置导航
router.beforeEach((to) => {
  // 1.获取用户信息对象
  const store = useUserStore()
  // 2.白名单
  const whiteList = ['/login']
  // 3.如果没有token，并且不在白名单中，就跳转到登录页
  if (!store.user?.token && !whiteList.includes(to.path)) return '/login'
  // 不返回，和return true 和 return undefined 一样
})

export default router
