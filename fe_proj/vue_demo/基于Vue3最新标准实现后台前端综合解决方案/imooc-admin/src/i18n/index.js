import { createI18n } from 'vue-i18n'

const messages = {
  en: {
    msg: {
      test: 'hello world'
    }
  },
  zh: {
    msg: {
      test: '你好世界'
    }
  }
}

const locale = 'en'

// 初始化 i18n 实例
const i18n = createI18n({
  //   使用 Composition API 模式，则需要将其设置为false
  legacy: false,
  //   全局注入 $t 函数
  globalInjection: true,
  locale,
  messages
})

// 把 `i18n` 注册到 `vue` 实例
export default i18n
