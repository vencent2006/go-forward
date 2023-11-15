import axios from 'axios'

// axios的实例
const instance = axios.create({
  baseURL: 'https://fastmock.site/mock/ae8e9031947a302fed5f92425995aa19/jd',
  timeout: 10000, // 超时时间为10s
  headers: {
    'Content-Type': 'application/json' // 指定content-type 为json
  }
})

export const get = (url, params = {}) => {
  return new Promise((resolve, reject) => {
    instance.get(url, { params }).then((response) => {
      resolve(response.data)
    }, err => {
      reject(err)
    })
  })
}
export const post = (url, data = {}) => {
  return new Promise((resolve, reject) => {
    instance.post(url, data, {
      // 我把这个header挪到instance里，如果post还有其他要求，可以继续在这里写
      // headers: {
      //   'Content-Type': 'application/json' // 指定content-type 为json
      // }
    }).then((response) => {
      resolve(response.data)
    }, err => {
      reject(err)
    })
  })
}
