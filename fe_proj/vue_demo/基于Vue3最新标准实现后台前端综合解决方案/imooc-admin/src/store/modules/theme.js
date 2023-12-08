import { getItem, setItem } from '@/utils/storage'
import { DEFAULT_COLOR, MAIN_COLOR } from '@/constant'

export default {
  namespaced: true,
  state: () => ({
    mainColor: getItem(MAIN_COLOR) || DEFAULT_COLOR
  }),
  mutations: {
    /**
     * 设置主题色
     */
    setMainColor(state, newColor) {
      state.mainColor = newColor
      setItem(MAIN_COLOR, newColor) // 写入localstorage
    }
  }
}
