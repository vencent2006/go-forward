import { createApp } from 'vue'
// i18n （PS：导入放到 APP.vue 导入之前，因为后面我们会在 app.vue 中使用国际化内容）
import i18n from '@/i18n'
import App from './App.vue'
import router from './router'
import store from './store'
import installElementPlus from './plugins/element'
// 初始化样式表
import '@/styles/index.scss'
// 导入 svgIcon
import installIcons from '@/icons'
// 导入路由鉴权
import '@/permission'

const app = createApp(App)
installElementPlus(app)
installIcons(app)
app
  .use(store) // 全局状态
  .use(router) // 路由
  .use(i18n) // 国际化
  .mount('#app') // 挂载点
