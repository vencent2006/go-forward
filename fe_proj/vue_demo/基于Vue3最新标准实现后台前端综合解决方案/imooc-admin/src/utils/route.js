import path from 'path'

/**
 * 返回所有子路由
 */
const getChildrenRoutes = (routes) => {
  const result = []
  routes.forEach((route) => {
    if (route.children && route.children.length > 0) {
      result.push(...route.children)
    }
  })
  return result
}
/**
 * 处理脱离层级的路由：某个一级路由为其他子路由，则剔除该一级路由，保留路由层级
 * @param {*} routes router.getRoutes()
 */
export const filterRouters = (routes) => {
  const childrenRoutes = getChildrenRoutes(routes)
  return routes.filter((route) => {
    return !childrenRoutes.find((childrenRoute) => {
      return childrenRoute.path === route.path
    })
  })
}

/**
 * 判断数据是否为空值
 */
function isNull(data) {
  if (!data) return true
  if (JSON.stringify(data) === '{}') return true
  if (JSON.stringify(data) === '[]') return true
  return false
}
/**
 * 根据 routes 数据，返回对应 menu 规则数组
 */
export function generateMenus(routes, basePath = '') {
  const result = []
  // 遍历路由表
  routes.forEach((item) => {
    // 1. 不存在 children && 不存在 meta 直接 return
    if (isNull(item.meta) && isNull(item.children)) return
    // 2. 存在 children 不存在 meta，进入迭代
    if (isNull(item.meta) && !isNull(item.children)) {
      result.push(...generateMenus(item.children))
      return
    }

    // 存在 meta

    // 合并 path 作为跳转路径
    const routePath = path.resolve(basePath, item.path)
    console.log(routePath)
    // 路由分离之后，存在同名父路由的情况，需要单独处理
    // todo 没看懂啥时候会同名父路由
    // find() 方法返回通过测试（函数内判断）的数组的第一个元素的值。
    let route = result.find((item) => item.path === routePath)
    if (!route) {
      route = {
        // todo 这个"..."的意思，应该是把item，变成了数组[item]; 然后是key和value同名？ route.item = [item]
        // todo 还是找个机会测试下吧
        ...item,
        path: routePath,
        children: []
      }

      // icon 与 title 必须全部存在
      if (route.meta.icon && route.meta.title) {
        // meta 存在生成 route 对象，放入 arr
        result.push(route)
      }
    }

    // 存在 children 进入迭代到children
    if (item.children) {
      // todo route应该是可以为undefined, 这还能route.path吗
      route.children.push(...generateMenus(item.children, route.path))
    }
  })
  return result
}
