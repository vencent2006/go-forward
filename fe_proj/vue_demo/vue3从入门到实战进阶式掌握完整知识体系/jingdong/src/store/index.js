import { createStore } from 'vuex'

export default createStore({
  state: {
    // 购物车
    cartList: {
      // 第一层级：商铺id
      // 第二层级：商品id
      /*
        {
          "shopId": {
            "shopName": "沃尔玛",
            "productList": {
                "productId_1": {
                    "_id": "1",
                    "name": "番茄 250g / 份",
                    "imgUrl": "http://www.dell-lee.com/imgs/vue3/tomato.png",
                    "sales": 10,
                    "price": 33.6,
                    "oldPrice": 39.6,
                    "count": 1,
                    "check": true
                },
                "productId_2": {
                    "_id": "2",
                    "name": "土豆 250g / 份",
                    "imgUrl": "http://www.dell-lee.com/imgs/vue3/tomato.png",
                    "sales": 12,
                    "price": 23.6,
                    "oldPrice": 49.6,
                    "count": 3,
                    "check": false
                }
            }
          }
        }
      */
    }
  },
  getters: {
  },
  mutations: {
    // 修改product到购物车
    changeCartItemInfo(state, payload) {
      const { shopId, productId, productInfo, num } = payload
      const shopInfo = state.cartList[shopId] || {
        shopName: '',
        productList: {}
      }// shopInfo不存在就给一个初值
      let product = shopInfo.productList[productId]
      if (!product) {
        console.log('shopInfo ', productId, ' not exist')
        product = productInfo
        product.count = 0
      }
      product.count = product.count + num
      // 选中状态为true，表示选中了
      // num > 0 && (product.check = true) // num > 0 没法像老师一样，加上括号
      // 跟下面的等价
      if (num > 0) { product.check = true }

      // 跟下面的等价 product.count < 0 && (product.count = 0)
      if (product.count < 0) {
        product.count = 0
      }

      // 写回到state
      // 必须要写回吗？？js有引用的概念吗？
      shopInfo.productList[productId] = product
      state.cartList[shopId] = shopInfo

      console.log('cartList  ', state.cartList)
    },
    // 改变购物车选中状态
    changeCartItemChecked(state, payload) {
      const { shopId, productId } = payload
      const product = state.cartList[shopId].productList[productId]
      product.check = !product.check
    },
    // 清空购物车
    cleanCartProducts(state, payload) {
      const { shopId } = payload
      state.cartList[shopId].productList = {}
    },
    // 全选购物车
    setCartItemsChecked(state, payload) {
      const { shopId } = payload
      const products = state.cartList[shopId].productList
      if (products) {
        for (const key in products) {
          const product = products[key]
          product.check = true
        }
      }
    },
    // 修改店铺名称
    changeShopName(state, payload) {
      const { shopId, shopName } = payload
      const shopInfo = state.cartList[shopId] || {
        shopName: '',
        productList: {}
      }
      shopInfo.shopName = shopName
      state.cartList[shopId] = shopInfo
    }
  },
  actions: {
  },
  modules: {
  }
})
