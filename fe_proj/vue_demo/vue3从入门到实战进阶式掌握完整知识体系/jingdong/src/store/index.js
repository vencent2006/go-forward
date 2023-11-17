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
    addItemToCart(state, payload) {
      const { shopId, productId, productInfo } = payload
      console.log(shopId, productId, productInfo)
    }
  },
  actions: {
  },
  modules: {
  }
})
