<template>
  <div class="content">
    <div class="category">
      <!-- :class="{ 'docker__item': true, 'docker__item--active': index === 0 }"> -->
      <div v-for="cat in categories" :key="cat.name"
        :class="{ 'category__item': true, 'category__item--active': currentTab === cat.tab }"
        @click="() => handleTabClick(cat.tab)">{{ cat.name }}</div>
    </div>

    <div class="product">
      <div class="product__item" v-for="item in list" :key="item._id">
        <img class="product__item__img" :src="item.imgUrl">
        <div class="product__item__detail">
          <h4 class="product__item__title">{{ item.name }}</h4>
          <p class="product__item__sales">月售{{ item.sales }}件</p>
          <p class="product__item__price">
            <span class="product__item__price__yen">&yen;</span>{{ item.price }}
            <span class="product__item__price__origin">&yen;{{ item.oldPrice }}</span>
          </p>
        </div>
        <div class="product__number">
          <!-- 减号操作 -->
          <span class="product__number__minus"
            @click="() => { changeCartItem(shopId, item._id, item, -1, shopName) }">-</span>
          <!-- 从购物车的数据里拿值 -->
          {{ getProductCartCount(shopId, item._id) }}
          <!-- 加号操作 -->
          <span class="product__number__plus"
            @click="() => { changeCartItem(shopId, item._id, item, 1, shopName) }">+</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { reactive, ref, toRefs, watchEffect } from 'vue'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { get } from '@/utils/request'
import { useCommonCartEffect } from '@/effects/cartEffects'

const categories = [
  { name: '全部商品', tab: 'all' },
  { name: '秒杀', tab: 'seckill' },
  { name: '水果', tab: 'fruit' }
]

// 和 tab 切换相关的逻辑
const useTabEffect = () => {
  const currentTab = ref(categories[0].tab)
  const handleTabClick = (tab) => {
    // console.log('tab', tab)
    currentTab.value = tab
  }
  return { currentTab, handleTabClick }
}

// 列表内容相关的逻辑
const useCurrentListEffect = (currentTab, shopId) => {
  const content = reactive({ list: [] })
  const getContentData = async () => {
    // all 全部; seckill 秒杀; fruit 水果
    // 依赖了 currentTab 了
    const result = await get(`/api/shop/${shopId}/products`, { tab: currentTab.value })
    if (result?.errno === 0 && result?.data?.length) {
      content.list = result.data
    }
  }

  // 这个watchEffect 要学会
  watchEffect(() => {
    // 首次加载 or 依赖的数据发生变化时 要执行
    getContentData()
  })

  const { list } = toRefs(content)
  return { list }
}

// 购物车相关
const useCartEffect = (shopId) => {
  const store = useStore()
  const { cartList, changeCartItemInfo } = useCommonCartEffect(shopId)
  const changeShopName = (shopId, shopName) => {
    store.commit('changeShopName', { shopId, shopName })
  }
  const changeCartItem = (shopId, productId, productInfo, num, shopName) => {
    changeCartItemInfo(shopId, productId, productInfo, num)
    changeShopName(shopId, shopName)
  }
  // 获取指定shopId下的指定productId的数量
  const getProductCartCount = (shopId, productId) => {
    return cartList?.[shopId]?.productList?.[productId]?.count || 0
  }
  return { cartList, changeCartItem, getProductCartCount }
}

export default {
  name: 'Content',
  props: ['shopName'],
  setup() {
    const route = useRoute()
    const shopId = route.params.id

    // tab 和 列表
    const { currentTab, handleTabClick } = useTabEffect()
    const { list } = useCurrentListEffect(currentTab, shopId)
    // 购物车
    const { cartList, changeCartItem, getProductCartCount } = useCartEffect(shopId)

    return {
      cartList,
      categories,
      currentTab,
      handleTabClick,
      list,
      // shopId
      shopId,
      // 购物车
      changeCartItem,
      getProductCartCount
    }
  }
}
</script>

<style lang='scss' scoped>
@import '@/style/variables.scss';
@import '@/style/mixins.scss';

.content {
  display: flex;
  position: absolute;
  left: 0;
  right: 0;
  top: 1.5rem;
  bottom: .5rem;
}

.category {
  overflow-y: scroll;
  height: 100%;
  width: .76rem;
  background: $search-bgColor;

  &__item {
    line-height: .4rem;
    text-align: center;
    font-size: .14rem;
    color: $content-fontcolor;

    &--active {
      background: $bgColor;
    }
  }
}

.product {
  overflow-y: scroll; // 超出区域可以上下滚
  flex: 1; // 右侧填满

  &__item {
    position: relative;
    display: flex;
    padding: .12rem 0;
    margin: 0 .16rem;
    border-bottom: .01rem solid $content-bgColor;

    &__detail {
      overflow: hidden;
    }

    &__img {
      width: .68rem;
      height: .68rem;
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

    &__sales {
      margin: .06rem 0;
      line-height: .16rem;
      font-size: .12rem;
      color: $content-fontcolor;
    }

    &__price {
      margin: 0;
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
