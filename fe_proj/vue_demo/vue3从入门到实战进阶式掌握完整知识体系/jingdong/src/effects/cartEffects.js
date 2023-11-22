// import { toRefs } from 'vue'
import { useStore } from 'vuex'
import { computed } from 'vue'

// 购物车相关逻辑
export const useCommonCartEffect = (shopId) => {
  const store = useStore()
  // const { cartList } = toRefs(store.state)
  // 不加toRefs也能用
  const cartList = store.state.cartList
  // 点击+号的操作
  const changeCartItemInfo = (shopId, productId, productInfo, num) => {
    // console.log(shopId, productId, productInfo)
    // store.js mutations里要有 addItemToCart 这个方法
    store.commit('changeCartItemInfo', {
      shopId, productId, productInfo, num
    })
  }

  // 商品列表
  const productList = computed(() => {
    // const 是限定作用域的 所以下面的productList 不会和 上面的productList 发生语法错误
    // productList 是个对象，而不是个数组
    const productList = cartList[shopId].productList || {}
    // 去掉count=0的商品
    const notEmptyProductList = {} // 记住是对象
    for (const i in productList) {
      const product = productList[i]
      if (product.count > 0) { // 非空的才往里添加
        notEmptyProductList[i] = product
      }
    }
    return notEmptyProductList
  })

  // 商品名称
  // 这个地方得用计算属性（computed，即响应式的）
  // 单纯使用 const shopName = cartList[shopId]?.shopName || ''，
  // 如果 shopName=''了，就不会再变化，即没有computed严谨
  const shopName = computed(() => {
    // const 是限定作用域的 所以下面的productList 不会和 上面的productList 发生语法错误
    const shopName = cartList[shopId]?.shopName || ''
    return shopName
  })

  // 计算属性 购物车的
  // total 商品个数
  // price 总价
  // allChecked 是否都被选中
  const calculations = computed(() => {
    const proudctList = cartList[shopId]?.productList
    // total 商品个数
    // price 总价
    // allChecked 是否都被选中
    const result = { total: 0, price: 0, allChecked: true }
    if (proudctList) {
      for (const i in proudctList) {
        const product = proudctList[i]
        // 总数
        result.total += product.count
        // 总价
        if (product.check) {
          // 选中的，才累加金额
          result.price += product.count * product.price
        }
        // 是否全选
        if (product.count > 0 && !product.check) {
          result.allChecked = false
        }
      }
    }
    result.price = result.price.toFixed(2) // price 保留2位小数
    return result
  })

  return { cartList, shopName, productList, calculations, changeCartItemInfo }
}
