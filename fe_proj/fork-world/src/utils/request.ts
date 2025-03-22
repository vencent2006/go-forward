import router from '@/router'
import { useUserStore } from '@/stores'
import axios, { AxiosError, type Method } from 'axios'
import { showToast } from 'vant'

export const baseURL = 'https://consult-api.itheima.net/'
const instance = axios.create({
  // 1. 基础地址 超时时间
  baseURL, // 基础路径
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
  (err: AxiosError) => {
    // 5. 处理401错误
    if (err.response?.status === 401) {
      // 登录失效
      // 清理用户信息
      const store = useUserStore()
      store.delUser()
      // 跳转登录 携带当前访问页面的地址（包含参数的）
      router.push({
        path: '/login',
        query: { returnUrl: router.currentRoute.value.fullPath }, // 登录成功后跳转回之前的页面
      })
    }
    return Promise.reject(err)
  },
)

export default instance

type Data<T> = {
  code: number
  message: string
  data: T
}

export const request = <T>(
  url: string,
  method: Method = 'GET', // 方法
  submitData?: object,
) => {
  // 参数：请求地址，请求方式，提交的数据
  // 返回值：Promise
  return instance.request<any, Data<T>>({
    url,
    method,
    [method.toUpperCase() === 'GET' ? 'params' : 'data']: submitData,
  })
}
