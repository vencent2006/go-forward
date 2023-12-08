import { createStore } from 'vuex'
import getters from './getters'
import user from './modules/user'
import app from './modules/app'
import theme from './modules/theme'

export default createStore({
  getters,
  modules: {
    user, // 用户数据
    app, // 应用，eg：侧边栏是否展开
    theme // 主题色
  }
})
