<template>
  <div class="cart">
    <div class="check">
      <div class="check__icon">
        <img src="http://www.dell-lee.com/imgs/vue3/basket.png" class="check__icon__img" />
        <div class="check__icon__tag">123123{{ total }}</div>
      </div>
      <div class="check__info">
        总计：<span class="check__info__price">&yen; {{ price }}</span>
      </div>
      <div class="check__btn">去结算</div>
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'

// 获取购物车信息逻辑
const useCartEffect = () => {
  const store = useStore()
  const route = useRoute()
  const shopId = route.params.id
  const cartList = store.state.cartList
  // 商品个数
  const total = computed(() => {
    const proudctList = cartList[shopId]
    // console.log('lalalala ----- ', proudctList)
    let count = 0
    if (proudctList) {
      for (const i in proudctList) {
        const product = proudctList[i]
        count += product.count
      }
    }
    return count
  })
  // 总价
  const price = computed(() => {
    const proudctList = cartList[shopId]
    let count = 0
    if (proudctList) {
      for (const i in proudctList) {
        const product = proudctList[i]
        count += product.count * product.price
      }
    }
    // 保留2位小数
    return count.toFixed(2)
  })
  return { total, price }
}

export default {
  name: 'Cart',
  setup() {
    const { total, price } = useCartEffect()
    return { total, price }
  }
}
</script>

<style lang='scss' scoped>
@import '@/style/variables.scss';

.cart {
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
}

.check {
  display: flex;
  // box-sizing: border-box;
  height: .49rem;
  border-top: .01rem solid $content-bgColor;
  line-height: .49rem;

  &__icon {
    position: relative;
    width: .84rem;

    &__img {
      display: block;
      margin: .12rem auto;
      width: .28rem;
      height: .26rem;
    }

    &__tag {
      position: absolute;
      left: .46rem;
      top: .04rem;
      padding: 0 .04rem;
      min-width: .2rem;
      height: .2rem;
      line-height: .2rem;
      background: $highlight-fontColor;
      border-radius: .01rem;
      font-size: .12rem;
      text-align: center;
      color: $bgColor;
      transform: scale(.5);
      transform-origin: left center;
    }
  }

  // 总计
  &__info {
    flex: 1;
    color: $content-fontcolor;
    font-size: .12rem;
    // font-weight: ;

    &__price {
      color: $highlight-fontColor;
      font-size: .18rem;
    }
  }

  // button
  &__btn {
    width: .98rem;
    font-size: .14rem;
    text-align: center;
    background: #4FB0F9;
    color: $bgColor;
  }
}
</style>
