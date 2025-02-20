import { useUserStore } from '@/stores'
import axios from 'axios'

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
    return res.data
  },
  (err) => {
    return Promise.reject(err)
  },
)

export default instance
