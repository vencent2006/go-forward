<template>
  <div class="wrapper">
    <div class="search">
      <div class="search__back iconfont" @click="handleBackClick">&#xe6f2;</div>
      <div class="search__content">
        <span class="search__content__icon iconfont">&#xe62d;</span>
        <input class="search__content__input" placeholder="请输入商品名称" />
      </div>
    </div>
    <!-- v-show item.imgUrl有数据了，才展示组件 -->
    <ShopInfo :item="item" :hideBorder="true" v-show="item.imgUrl" />
    <Content :shopName="item.name" />
    <Cart />
  </div>
</template>

<script>
import { reactive, toRefs } from 'vue'
import { useRoute } from 'vue-router'
import { get } from '@/utils/request'
import ShopInfo from '@/components/ShopInfo'
import Content from './Content.vue'
import Cart from './Cart.vue'
import { useBackRouterEffect } from '@/effects/routeEffects'

const useShopInfoEffect = () => {
  const data = reactive({ item: {} }) // 记住：这种reactive的方式

  const getItemData = async () => {
    const route = useRoute()
    // console.log(route.params.id)
    const result = await get(`/api/shop/${route.params.id}`)
    // console.log(result)
    if (result?.errno === 0 && result?.data) {
      data.item = result.data
    }
  }

  const { item } = toRefs(data)
  return { item, getItemData }
}

export default {
  name: 'Shop',
  components: { ShopInfo, Content, Cart },
  setup() {
    const { item, getItemData } = useShopInfoEffect()
    const { handleBackClick } = useBackRouterEffect()
    getItemData() // 执行一下，取下数据
    return { item, handleBackClick }
  }
}
</script>

<style lang="scss" scoped>
@import '@/style/variables.scss';

.wrapper {
  padding: 0 .18rem;
}

.search {
  display: flex;
  margin: .14rem 0 .16rem 0;
  line-height: .32rem;

  &__back {
    width: .3rem;
    font-size: .24rem;
    color: #B6B6B6;
  }

  &__content {
    flex: 1; // flex:1 直接填满剩余空间
    background: $search-bgColor;
    border-radius: .16rem;
    display: flex;

    &__icon {
      width: .44rem;
      text-align: center;
      color: $search-fontColor;
    }

    &__input {
      display: block;
      width: 100%;
      padding-right: .2rem;
      border: none;
      outline: none;
      background: none;
      height: .32rem;
      font-size: .14rem;
      color: $content-fontcolor;

      &::placeholder {
        color: $content-fontcolor;
      }
    }
  }
}
</style>
