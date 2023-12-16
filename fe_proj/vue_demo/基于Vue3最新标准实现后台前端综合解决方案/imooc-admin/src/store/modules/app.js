import { LANG, TAGS_VIEW } from '@/constant'
import { getItem, setItem } from '@/utils/storage'

export default {
  namespaced: true,
  // () => ({sidebarOpened: true}), 这么写才能返回一个对象
  // () => {}, 这么写是单纯的执行一个函数，并没有返回值
  state: () => ({
    sidebarOpened: true, // 侧边栏 是否 展开
    language: getItem(LANG) || 'zh', // 国际化 的 当前语言
    tagsViewList: getItem(TAGS_VIEW) || [] // tag view
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
    },
    /**
     * 添加 tags
     */
    addTagsViewList(state, tag) {
      const isFind = state.tagsViewList.find((item) => {
        return item.path === tag.path
      })
      // 处理重复
      if (!isFind) {
        state.tagsViewList.push(tag) // 如果没找到，就放进 tagsViewList
        setItem(TAGS_VIEW, state.tagsViewList) // 并存入 localStorage
      }
    },
    /**
     * 为指定的 tag 修改 title
     */
    changeTagsView(state, { index, tag }) {
      state.tagsViewList[index] = tag
      setItem(TAGS_VIEW, state.tagsViewList)
    }
  },
  actions: {}
}
