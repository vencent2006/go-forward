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
    // 添加product到购物车
    addItemToCart(state, payload) {
      const { shopId, productId, productInfo } = payload
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
      product.count += 1

      // 写回到state
      shopInfo[productId] = product
      state.cartList[shopId] = shopInfo

      console.log('shopInfo  ', shopInfo[productId])
    },

    // 删除product，从购物车
    delItemToCart(state, payload) {
      const { shopId, productId } = payload
      // console.log(shopId, productId)
      const shopInfo = state.cartList[shopId]
      if (!shopInfo) { return } // shopInfo不存在, 直接返回
      const product = shopInfo[productId]
      if (!product) { return }
      product.count -= 1
      if (product.count < 0) {
        product.count = 0
      }

      // 写回到state
      shopInfo[productId] = product
      state.cartList[shopId] = shopInfo

      console.log('shopInfo  ', shopInfo[productId])
    }
  },
  actions: {
  },
  modules: {
  }
})
