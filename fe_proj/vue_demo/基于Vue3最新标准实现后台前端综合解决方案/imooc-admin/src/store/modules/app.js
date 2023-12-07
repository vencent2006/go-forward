import { LANG } from '@/constant'
import { getItem, setItem } from '@/utils/storage'

export default {
  namespaced: true,
  // () => ({sidebarOpened: true}), 这么写才能返回一个对象
  // () => {}, 这么写是单纯的执行一个函数，并没有返回值
  state: () => ({
    sidebarOpened: true,
    language: getItem(LANG) || 'zh'
  }),
  mutations: {
    /**
     * 触发 侧边栏 开关切换
     */
    triggerSidebarOpened(state) {
      state.sidebarOpened = !state.sidebarOpened
    },
    /**
     * 设置国际化
     */
    setLanguage(state, lang) {
      setItem(LANG, lang) // 设置到localstorage里
      state.language = lang
    }
  },
  actions: {}
}
