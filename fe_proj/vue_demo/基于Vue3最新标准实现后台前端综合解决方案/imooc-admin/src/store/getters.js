import variables from '@/styles/variables.module.scss'
// 快捷访问
const getters = {
  token: (state) => state.user.token,
  userInfo: (state) => state.user.userInfo,
  /**
   * 是否有用户信息
   * @param state
   * @return true 表示已存在用户信息
   */
  hasUserInfo: (state) => {
    return JSON.stringify(state.user.userInfo) !== '{}'
  },
  /**
   * css 全局变量
   * @param state
   * @returns {*}
   */
  cssVar: (state) => variables,
  /**
   * 侧边栏(sidebar) 打开状态
   * @param state
   * @returns {boolean} true: 打开状态, false: 关闭状态
   */
  sidebarOpened: (state) => state.app.sidebarOpened
}

export default getters
