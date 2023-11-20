<template>
  <!-- 蒙层 -->
  <div class="mask" v-if="showCart" @click="handleCartShowChange" />
  <!-- 购物车信息 -->
  <div class="cart">
    <!-- 展示商品信息 -->
    <div class="product" v-if="showCart">
      <!-- 全选 和 清空购物车 -->
      <div class="product__header">
        <div class="product__header__all">
          <span class="product__header__icon iconfont" v-html="allChecked ? '&#xe652;' : '&#xe667;'"
            @click="() => setCartItemsChecked(shopId)"></span>
          全选
        </div>
        <div class="product__header__clear">
          <span class="product__header__clear__btn" @click="() => cleanCartProducts(shopId)">清空购物车</span>
        </div>
      </div>
      <!-- 这个template主要是占位，没有具体样式内容 -->
      <template v-for="item in productList" :key="item._id">
        <div class="product__item" v-if="item.count > 0">
          <div class="product__item__checked iconfont" v-html="item.check ? '&#xe652;' : '&#xe667;'"
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
        <img src="http://www.dell-lee.com/imgs/vue3/basket.png" class="check__icon__img"
          @click="() => handleCartShowChange(showCart)" />
        <div class="check__icon__tag">{{ total }}</div>
      </div>
      <div class="check__info">
        总计：<span class="check__info__price">&yen; {{ price }}</span>
      </div>
      <div class="check__btn">
        <router-link :to="{ name: 'Home' }">
          去结算
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { computed, ref } from 'vue'
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

  // 都被选中
  const allChecked = computed(() => {
    const productList = cartList[shopId]
    let result = true
    let cnt = 0 // 购物车中选中的商品的数量
    if (productList) {
      for (const i in productList) {
        const product = productList[i]
        if (product.count > 0) {
          if (!product.check) {
            // 商品数量>0 且 没有选中
            // 那么，全选
            result = false
            break
          } else {
            cnt++
          }
        }
      }
    }
    if (cnt === 0) {
      // 一个都没选，就谈不上"全选"了
      result = false
    }
    // console.log('allChecked', result)
    return result
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

  // 清空购物车
  const cleanCartProducts = (shopId) => {
    store.commit('cleanCartProducts', { shopId })
  }

  // 全选 购物车, 老师只实现了全选，没有实现全不选
  const setCartItemsChecked = (shopId) => {
    console.log('setCartItemsChecked', shopId)
    store.commit('setCartItemsChecked', { shopId })
  }

  return {
    // 变量
    total,
    price,
    productList,
    allChecked,
    // 方法
    changeCardItemInfo,
    changeCartItemChecked,
    cleanCartProducts,
    setCartItemsChecked
  }
}

// 展示/隐藏购物车逻辑
const toggleCartEffect = () => {
  // 购物车是否展示
  const showCart = ref(false) // 是否展示购物车
  const handleCartShowChange = () => {
    showCart.value = !showCart.value
  }
  return { showCart, handleCartShowChange }
}

export default {
  name: 'Cart',
  setup() {
    const route = useRoute()
    const shopId = route.params.id

    const { total, price, productList, allChecked, changeCardItemInfo, changeCartItemChecked, cleanCartProducts, setCartItemsChecked } = useCartEffect(shopId)
    const { showCart, handleCartShowChange } = toggleCartEffect()

    return {
      // 变量
      shopId,
      total,
      price,
      productList,
      allChecked,
      showCart,
      // 方法
      changeCardItemInfo,
      changeCartItemChecked,
      cleanCartProducts,
      setCartItemsChecked,
      handleCartShowChange
    }
  }
}
</script>

<style lang='scss' scoped>
@import '@/style/variables.scss';
@import '@/style/mixins.scss';

.mask {
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
  top: 0;
  background: rgba(0, 0, 0, .5);
  z-index: 1;
}

.cart {
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 2;
  background: $bgColor;
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

    a {
      color: $bgColor;
      text-decoration: none; // 去除a标签的下划线
    }
  }
}

.product {
  overflow-y: scroll; // 超出区域可以上下滚
  flex: 1; // 右侧填满
  background: $bgColor;

  &__header {
    display: flex;
    line-height: .52rem;
    border-bottom: 1px solid #F1F1F1;
    font-size: .14rem;
    color: $content-fontcolor;

    &__all {
      width: .64rem;
      margin-left: .16rem;
    }

    &__icon {
      display: inline-block;
      margin-right: .1rem;
      vertical-align: top;
      color: $btn-bgColor;
      font-size: .2rem;
    }

    &__clear {
      flex: 1;
      margin-right: .16rem;
      text-align: right;

      &__btn {
        display: inline-block; // 让高度撑满
      }
    }
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
      color: $btn-bgColor;
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
      bottom: .3rem;

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
