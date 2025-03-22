import 'vue-router'

// 扩展路由的 meta 属性
declare module 'vue-router' {
  interface RouteMeta {
    title?: string
  }
}
