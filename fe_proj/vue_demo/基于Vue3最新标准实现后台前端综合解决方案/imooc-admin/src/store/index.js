import { createStore } from 'vuex'
import user from './modules/user'
import app from './modules/app'
import getters from './getters'

export default createStore({
  getters,
  modules: {
    user, // 用户数据
    app // 应用，eg：侧边栏是否展开
  }
})
