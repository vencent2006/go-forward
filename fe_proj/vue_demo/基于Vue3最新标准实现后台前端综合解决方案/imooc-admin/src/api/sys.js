import request from '@/utils/request'

/**
 * 登录
 * return promise
 */
export const login = (data) => {
  return request({
    url: '/sys/login',
    method: 'POST',
    data
  })
}
