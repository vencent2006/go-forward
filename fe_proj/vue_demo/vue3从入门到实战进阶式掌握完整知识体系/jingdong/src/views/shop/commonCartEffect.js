import { toRefs } from 'vue'
import { useStore } from 'vuex'

// 购物车相关逻辑
export const useCommonCartEffect = () => {
  const store = useStore()
  const { cartList } = toRefs(store.state)

  // 点击+号的操作
  const changeCardItemInfo = (shopId, productId, productInfo, delta) => {
    // console.log(shopId, productId, productInfo)
    // store.js mutations里要有 addItemToCart 这个方法
    store.commit('changeCardItemInfo', {
      shopId, productId, productInfo, delta
    })
  }
  return { cartList, changeCardItemInfo }
}
