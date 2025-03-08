import { useUserStore } from '@/stores'
import { createRouter, createWebHistory } from 'vue-router'

import NProgress from 'nprogress' // 引入进度条
import 'nprogress/nprogress.css' // 引入样式
NProgress.configure({
  showSpinner: false, // 不显示环形进度图标
})

// createWebHashHistory 哈希路由 会有 # 号
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL), // 路由模式
  routes: [
    {
      path: '/login',
      component: () => import('@/views/Login/index.vue'),
      meta: { title: '家庭档案' },
    },
    {
      path: '/user/patient',
      component: () => import('@/views/User/PatientPage.vue'),
      meta: { title: '登录' },
    },
    {
      path: '/consult/fast',
      component: () => import('@/views/Consult/ConsultFast.vue'),
      meta: { title: '极速问诊' },
    },
    {
      path: '/consult/dep',
      component: () => import('@/views/Consult/ConsultDep.vue'),
      meta: { title: '选择科室' },
    },
    {
      path: '/consult/pay',
      component: () => import('@/views/Consult/ConsultPay.vue'),
      meta: { title: '问诊支付' },
    },
    {
      path: '/consult/illness',
      component: () => import('@/views/Consult/ConsultIllness.vue'),
      meta: { title: '病情描述' },
    },
    {
      path: '/',
      redirect: '/home',
      component: () => import('@/views/Layout/index.vue'),
      children: [
        // 二级路由
        {
          path: '/home',
          component: () => import('@/views/Home/index.vue'),
          meta: { title: '首页' },
        },
        {
          path: '/article',
          component: () => import('@/views/Article/index.vue'),
          meta: { title: '健康百科' },
        },
        {
          path: '/notify',
          component: () => import('@/views/Notify/index.vue'),
          meta: { title: '消息通知' },
        },
        {
          path: '/user',
          component: () => import('@/views/User/index.vue'),
          meta: { title: '个人中心' },
        },
      ],
    },
  ],
})

// 全局的前置导航
router.beforeEach((to) => {
  // 开启进度条
  NProgress.start()
  // 1.获取用户信息对象
  const store = useUserStore()
  // 2.白名单
  const whiteList = ['/login']
  // 3.如果没有token，并且不在白名单中，就跳转到登录页
  if (!store.user?.token && !whiteList.includes(to.path)) return '/login'
  // 不返回，和return true 和 return undefined 一样
})

// 全局的后置导航
router.afterEach((to) => {
  document.title = `${to.meta.title || ''}-优医问诊` // 动态设置标题
  // 关闭进度条
  NProgress.done()
})

export default router
