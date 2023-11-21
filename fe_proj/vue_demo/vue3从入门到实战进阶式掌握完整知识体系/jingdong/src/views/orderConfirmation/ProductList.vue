<template>
  <!-- 购物车商品列表 -->
  <div class="products">
    <!-- 商店名称 -->
    <div class="products__title">{{ shopName }}</div>
    <!-- 用wrapper包一下list，让list自己撑开高度 -->
    <div class="products__wrapper">
      <!-- product 列表 -->
      <div class="products__list">
        <template v-for="item in productList" :key="item._id">
          <div class="products__item" v-if="item.count > 0">
            <img class="products__item__img" :src="item.imgUrl">
            <div class="products__item__detail">
              <h4 class="products__item__title">{{ item.name }}</h4>
              <p class="products__item__price">
                <!-- 单价 -->
                <span>
                  <span class="products__item__yen">&yen;</span>
                  {{ item.price }} x {{ item.count }}
                </span>
                <!-- 总价 -->
                <span class="products__item__total">
                  <span class="products__item__yen">&yen;</span>
                  {{ (item.price * item.count).toFixed(2) }}
                </span>
              </p>
            </div>
          </div>
        </template>
      </div>
    </div>
  </div>
</template>

<script>
import { useRoute } from 'vue-router'
import { useCommonCartEffect } from '@/effects/cartEffects'

export default {
  name: 'ProductList',
  setup() {
    const route = useRoute()
    const shopId = route.params.id
    const { shopName, productList } = useCommonCartEffect(shopId)
    return { shopName, productList }
  }
}
</script>

<style lang='scss' scoped>
@import '@/style/variables.scss';
@import '@/style/mixins.scss';

.products {
  margin: .16rem .18rem .1rem .18rem;
  background: $bgColor;

  &__title {
    padding: .16rem;
    font-size: .16rem;
    color: $content-fontcolor;
  }

  &__wrapper {
    overflow-y: scroll;
    margin: 0 .18rem;
    position: absolute;
    left: 0;
    right: 0;
    bottom: .6rem;
    top: 2.6rem;
  }

  // list 自己撑开高度
  &__list {
    background: $bgColor;
  }

  &__item {
    position: relative;
    display: flex;
    padding: 0 .16rem .16rem .16rem;

    &__img {
      width: .46rem;
      height: .46rem;
      margin-right: .16rem;
    }

    &__detail {
      flex: 1;
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
      display: flex;
      margin: .06rem 0 0 0; // price 是p标签，都设置为0，清除默认样式
      line-height: .2rem;
      font-size: .14rem;
      color: $highlight-fontColor;
    }

    &__total {
      flex: 1;
      text-align: right;
      color: $dark-fontColor;
    }

    &__yen {
      font-size: .12rem;
    }
  }
}
</style>
