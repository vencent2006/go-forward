/**
 * 存储数据
 */
export const setItem = (key, value) => {
  // value 分为两种情况:
  // 1. 基本数据类型
  // 2. 复杂数据类型
  if (typeof value === 'object') {
    // 复杂数据类型
    value = JSON.stringify(value)
  }
  window.localStorage.setItem(key, value)
}

/**
 * 获取数据
 */
export const getItem = (key) => {
  const data = window.localStorage.getItem(key)
  // 直接判断是json比较麻烦，上来先json parse，如果不能parse就直接返回
  try {
    return JSON.parse(data)
  } catch (err) {
    return data
  }
}

/**
 * 删除指定数据
 */
export const removeItem = (key) => {
  window.localStorage.removeItem(key)
}

/**
 * 删除所有数据
 */
export const removeAllItem = () => {
  window.localStorage.clear()
}
