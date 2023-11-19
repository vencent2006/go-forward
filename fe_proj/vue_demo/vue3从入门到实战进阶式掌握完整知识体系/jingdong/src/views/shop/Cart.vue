<template>
  <div class="cart">
    <!-- 展示商品信息 -->
    <div class="product">
      <div class="product__header">

      </div>
      <template v-for="item in productList" :key="item._id">
        <div class="product__item" v-if="item.count > 0">
          <div class="product__item__checked iconfont" v-html="item.check ? '&#xe652;' : '&#xe6f7;'"
            @click="() => changeCartItemChecked(shopId, item._id)"></div>
          <img class="product__item__img" :src="item.imgUrl">
          <div class="product__item__detail">
            <h4 class="product__item__title">{{ item.name }}</h4>
            <p class="product__item__price">
              <span class="product__item__price__yen">&yen;</span>{{ item.price }}
              <span class="product__item__price__origin">&yen;{{ item.oldPrice }}</span>
            </p>
          </div>
          <div class="product__number">
            <!-- 减号操作 -->
            <span class="product__number__minus"
              @click="() => { changeCardItemInfo(shopId, item._id, item, -1) }">-</span>
            {{ item.count || 0 }}
            <!-- 加号操作 -->
            <span class="product__number__plus" @click="() => { changeCardItemInfo(shopId, item._id, item, 1) }">+</span>
          </div>
        </div>
      </template>
    </div>
    <!-- 确认信息部分 -->
    <div class="check">
      <div class="check__icon">
        <img src="http://www.dell-lee.com/imgs/vue3/basket.png" class="check__icon__img" />
        <div class="check__icon__tag">{{ total }}</div>
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
import { useCommonCartEffect } from './commonCartEffect'

// 获取购物车信息逻辑
const useCartEffect = (shopId) => {
  const { changeCardItemInfo } = useCommonCartEffect()
  const store = useStore()
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
    const productList = cartList[shopId]
    let count = 0
    if (productList) {
      for (const i in productList) {
        const product = productList[i]
        if (product.check) {
          // 选中的，才累加金额
          count += product.count * product.price
        }
      }
    }
    // 保留2位小数
    return count.toFixed(2)
  })

  // 商品列表
  const productList = computed(() => {
    console.log('shopId', shopId)
    // const 是限定作用域的 所以下面的productList 不会和 上面的productList 发生语法错误
    const productList = cartList[shopId] || []
    return productList
  })

  const changeCartItemChecked = (shopId, productId) => {
    // 找到shopId和productId对应的商品，把选中状态设置为相反
    store.commit('changeCartItemChecked', { shopId, productId })
  }

  return { total, price, productList, changeCardItemInfo, changeCartItemChecked }
}

export default {
  name: 'Cart',
  setup() {
    const route = useRoute()
    const shopId = route.params.id
    const { total, price, productList, changeCardItemInfo, changeCartItemChecked } = useCartEffect(shopId)
    return { total, price, productList, changeCardItemInfo, changeCartItemChecked, shopId }
  }
}
</script>

<style lang='scss' scoped>
@import '@/style/variables.scss';
@import '@/style/mixins.scss';

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

.product {
  overflow-y: scroll; // 超出区域可以上下滚
  flex: 1; // 右侧填满
  background: #FFF;

  &__header {
    height: .52rem;
    border-bottom: 1px solid #F1F1F1;
  }

  &__item {
    position: relative;
    display: flex;
    padding: .12rem 0;
    margin: 0 .16rem;
    border-bottom: .01rem solid $content-bgColor;

    &__checked {
      line-height: .5rem;
      margin-right: .2rem;
      color: #0091FF;
      font-size: .2rem;
    }

    &__detail {
      overflow: hidden;
    }

    &__img {
      width: .46rem;
      height: .46rem;
      margin-right: .16rem;
    }

    &__title {
      margin: 0;
      line-height: .2rem;
      font-size: .14rem;
      color: $content-fontcolor;
      // 要求ellipsis的父组件，必须有overflow: hidden; 这个ellipsis才能生效
      @include ellipsis; // 从mixin中引入
    }

    &__price {
      margin: .06rem 0 0 0;
      line-height: .2rem;
      font-size: .14rem;
      color: $highlight-fontColor;

      &__yen {
        font-size: .12rem;
      }

      &__origin {
        margin-left: .06rem;
        line-height: .2rem;
        font-size: .12rem;
        color: $light-fontColor;
        text-decoration: line-through;
      }
    }

    .product__number {
      position: absolute;
      right: 0;
      bottom: .12rem;

      &__minus,
      &__plus {
        display: inline-block; // span
        width: .2rem;
        height: .2rem;
        line-height: .16rem;
        border-radius: 50%;
        font-size: .2rem;
        text-align: center;
      }

      &__minus {
        border: .01rem solid $medium-fontColor;
        color: $medium-fontColor;
        margin-right: .05rem;
      }

      &__plus {
        background: $btn-bgColor;
        color: $bgColor;
        margin-left: .05rem;
      }
    }
  }

}
</style>
