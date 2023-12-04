import { getUserInfo, login } from '@/api/sys'
import md5 from 'md5'
import { setItem, getItem } from '@/utils/storage'
import { TOKEN } from '@/constant'
import router from '@/router'
import { setTimeStamp } from '@/utils/auth'

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
    token: getItem(TOKEN) || '',
    userInfo: {}
  }),
  mutations: {
    // token
    setToken(state, token) {
      state.token = token
      setItem(TOKEN, token)
    },
    // userInfo
    setUserInfo(state, userInfo) {
      state.userInfo = userInfo
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
            console.log(data)
            // 这也是user module的，所以要加user
            this.commit('user/setToken', data.token)
            // 跳转
            router.push('/')
            // 保存登录时间
            setTimeStamp()

            resolve()
          })
          .catch((err) => {
            reject(err)
          })
      })
    },
    /**
     * 获取用户信息
     */
    async getUserInfo(context) {
      const res = await getUserInfo()
      this.commit('user/setUserInfo', res)
      return res
    }
  }
}
