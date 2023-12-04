// 快捷访问
const getters = {
  token: (state) => state.user.token,
  /**
   * @return true 表示已存在用户信息
   */
  hasUserInfo: (state) => {
    return JSON.stringify(state.user.userInfo) !== '{}'
  }
}

export default getters
