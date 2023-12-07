import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'
import zhCn from 'element-plus/lib/locale/lang/zh-cn'
import en from 'element-plus/lib/locale/lang/en'
import store from '@/store'
export default (app) => {
  // 注册 `element` 时，根据当前语言选择使用哪种语言包
  // todo 会随着 store.getters.language 变化来导入不同的包吗？
  app.use(ElementPlus, { locale: store.getters.language === 'en' ? en : zhCn })
}
