import { login } from '@/api/sys'
import md5 from 'md5'
import { setItem, getItem } from '@/utils/storage'
import { TOKEN } from '@/constant'

/**
{
  "code": 200,
  "data": {
    "token": "dc7f7964-60e5-4f61-b65e-0ac43dce6b1f"
  },
  "message": "登录成功",
  "sccess": true
}
 */

export default {
  namespaced: true, // 表示是一个单独的模块，不会被合并
  state: () => ({
    token: getItem(TOKEN) || ''
  }),
  mutations: {
    setToken(state, token) {
      state.token = token
      setItem(TOKEN, token)
    }
  },
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
            console.log(data.data.data.token)
            const token = data.data.data.token
            // 这也是user
            this.commit('user/setToken', token)
            resolve()
          })
          .catch((err) => {
            reject(err)
          })
      })
    }
  }
}
