// 快捷访问
const getters = {
  token: (state) => state.user.token,
  userInfo: (state) => state.user.userInfo,
  /**
   * 是否有用户信息
   * @return true 表示已存在用户信息
   */
  hasUserInfo: (state) => {
    return JSON.stringify(state.user.userInfo) !== '{}'
  },
  /**
   * css 全局变量
   * @returns {*}
   */
  cssVar: (state) => state.theme.variables,
  /**
   * 侧边栏(sidebar) 打开状态
   * @returns {boolean} true: 打开状态, false: 关闭状态
   */
  sidebarOpened: (state) => state.app.sidebarOpened,
  /**
   * 国际化的语言
   * @returns {*}
   */
  language: (state) => state.app.language,
  /**
   * 主题色
   */
  mainColor: (state) => state.theme.mainColor,
  /**
   * tag view
   */
  tagsViewList: (state) => state.app.tagsViewList
}

export default getters
