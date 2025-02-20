import { useUserStore } from '@/stores'
import axios from 'axios'
import { showToast } from 'vant'

const instance = axios.create({
  // 1. 基础地址 超时时间
  baseURL: 'https://consult-api.itheima.net/', // 基础路径
  timeout: 5000, // 超时时间
})

instance.interceptors.request.use(
  (config) => {
    // 携带token
    const store = useUserStore()
    if (store.user?.token && config.headers) {
      config.headers.Authorization = `Bearer ${store.user.token}`
    }
    return config
  },
  (err) => {
    return Promise.reject(err)
  },
)

instance.interceptors.response.use(
  (res) => {
    // 3. 处理业务失败
    if (res.data.code !== 10000) {
      // 错误提示
      showToast(res.data.message || '业务失败')
      // 错误的promise
      // 传入 code 将来catch的时候可以使用
      return Promise.reject(res.data)
    }
    // 4. 摘取核心响应数据
    return res.data
  },
  (err) => {
    // 5. 处理401错误
    return Promise.reject(err)
  },
)

export default instance
