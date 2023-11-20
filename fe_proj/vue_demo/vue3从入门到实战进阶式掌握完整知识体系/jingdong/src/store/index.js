import { createStore } from 'vuex'

export default createStore({
  state: {
    // 购物车
    cartList: {
      // 第一层级：商铺id
      // 第二层级：商品id
    }
  },
  getters: {
  },
  mutations: {
    // 修改product到购物车
    changeCardItemInfo(state, payload) {
      const { shopId, productId, productInfo, num } = payload
      // console.log(shopId, productId, productInfo)
      // const cartList = state.cartList
      let shopInfo = state.cartList[shopId]
      if (!shopInfo) { shopInfo = {} } // shopInfo不存在就给一个初值
      let product = shopInfo[productId]
      if (!product) {
        console.log('shopInfo ', productId, ' not exist')
        product = productInfo
        product.count = 0
      }
      product.count += num
      if (num > 0) { product.check = true } // 选中状态为true，表示选中了
      if (product.count < 0) {
        product.count = 0
      }

      // 写回到state
      // 必须要写回吗？？
      shopInfo[productId] = product
      state.cartList[shopId] = shopInfo

      console.log('shopInfo  ', shopInfo[productId])
    },
    // 改变购物车选中状态
    changeCartItemChecked(state, payload) {
      const { shopId, productId } = payload
      const product = state.cartList[shopId][productId]
      product.check = !product.check
    },
    // 清空购物车
    cleanCartProducts(state, payload) {
      const { shopId } = payload
      state.cartList[shopId] = {}
    },
    // 全选购物车
    setCartItemsChecked(state, payload) {
      const { shopId } = payload
      const products = state.cartList[shopId]
      if (products) {
        for (const key in products) {
          const product = products[key]
          product.check = true
        }
      }
    }
  },
  actions: {
  },
  modules: {
  }
})
