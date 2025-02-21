import type CpNavBar from '@/components/CpNavBar.vue'

declare module 'vue' {
  interface GlobalComponents {
    // 添加全局组件类型
    CpNavBar: typeof CpNavBar // typeof 作用：从js变量中取出对应的类型
  }
}
