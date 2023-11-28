/**
 * 判断是否为外部资源
 * @param {*} path
 */
export function isExternal(path) {
  // https? 表示的是http或者https，?是修饰s的
  return /^(https?:|mailto:|tel:)/.test(path)
}
