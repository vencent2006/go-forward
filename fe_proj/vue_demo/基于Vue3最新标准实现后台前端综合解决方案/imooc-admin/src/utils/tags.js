const whitelist = ['/login', '/import', '/404', '/401']

/**
 * path 是否需要被缓存
 * @param path
 * @returns {boolean}
 */
export function isTags(path) {
  // 不在白名单里，就是需要被缓存的
  return !whitelist.includes(path)
}
