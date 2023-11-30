import { login } from '@/api/sys'
import md5 from 'md5'

export default {
  namespaced: true, // 表示是一个单独的模块，不会被合并
  state: () => ({}),
  mutations: {},
  actions: {
    /**
     * 登录请求动作
     */
    login(context, userInfo) {
      const { username, password } = userInfo
      return new Promise((resolve, reject) => {
        login({
          username,
          password: md5(password)
        })
          .then((data) => {
            resolve()
          })
          .catch((err) => {
            reject(err)
          })
      })
    }
  }
}
