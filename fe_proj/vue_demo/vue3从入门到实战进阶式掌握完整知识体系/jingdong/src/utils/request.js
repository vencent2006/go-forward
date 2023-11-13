import axios from 'axios'
export const post = (url, data = {}) => {
  return new Promise((resolve, reject) => {
    axios.post(url, data, {
      baseURL: 'https://fastmock.site/mock/ae8e9031947a302fed5f92425995aa19/jd',
      headers: {
        'Content-Type': 'application/json' // 指定content-type 为json
      }
    }).then((response) => {
      resolve(response.data)
    }, err => {
      reject(err)
    })
  })
}
