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
    const productList = cartList[shopId].productList || []
    return productList
  })

  // 商品名称
  const shopName = computed(() => {
    // const 是限定作用域的 所以下面的productList 不会和 上面的productList 发生语法错误
    const shopName = cartList[shopId]?.shopName || []
    return shopName
  })

  return { cartList, shopName, productList, changeCartItemInfo }
}
