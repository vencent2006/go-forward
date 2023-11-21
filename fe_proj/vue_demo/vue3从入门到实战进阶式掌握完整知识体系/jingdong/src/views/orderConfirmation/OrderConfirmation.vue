<template>
  <div class="wrapper">
    <!-- 顶部信息 -->
    <div class="top">
      <div class="top__bgcolor"></div>
      <div class="top__header">
        <!-- 后退 -->
        <div class="top__header__back iconfont" @click="handleBackClick">&#xe6f2;</div>
        确认订单
      </div>
      <div class="top__receiver">
        <div class="top__receiver__title">收货地址</div>
        <div class="top__receiver__address">北京理工大学国防科技园2号楼10层</div>
        <div class="top__receiver__info">
          <span class="top__receiver__info__name">瑶妹（先生）</span>
          <span class="top__receiver__info__name">18911024266</span>
        </div>
        <!-- 进入 -->
        <div class="top__receiver__icon iconfont">&#xe6f2;</div>
      </div>
    </div>
    <!-- 商品 -->
    <div class="products">
      <!-- 商店名称 -->
      <div class="products__title">{{ shopName }}</div>
      <!-- 商品列表 -->
      <div class="products__list">
        <div class="products__item" v-for="item in productList" :key="item._id">
          <img class="products__item__img" :src="item.imgUrl">
          <div class="products__item__detail">
            <h4 class="products__item__title">{{ item.name }}</h4>
            <p class="products__item__price">
              <span class="products__item__price__yen">&yen;{{ item.price }} * {{ item.count }}</span>
              <span class="products__item__price__yen">&yen;{{ item.price * item.count }}</span>
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { toRefs } from 'vue'
import { useRoute } from 'vue-router'
import { useCommonCartEffect } from '@/effects/cartEffects'
import { useBackRouterEffect } from '@/effects/routeEffects'

export default {
  setup() {
    const route = useRoute()
    const shopId = route.params.id
    const { cartList, productList } = useCommonCartEffect(shopId)
    const { shopName } = toRefs(cartList[shopId])
    const { handleBackClick } = useBackRouterEffect()
    return { shopName, productList, handleBackClick }
  }
}
</script>

<style lang='scss' scoped>
@import '@/style/variables.scss';
@import '@/style/mixins.scss';

.wrapper {
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  background: #EEE;
}

.top {
  position: relative;
  height: 1.96rem;
  background-size: 100% 1.59rem;
  background-image: linear-gradient(0deg, rgba(0, 145, 255, 0.00) 4%, #0091FF 50%); // 渐变
  background-repeat: no-repeat; // 不要重复

  &__header {
    position: relative;
    padding-top: .26rem;
    line-height: .24rem;
    color: #EEE;
    text-align: center;
    font-size: .16rem;

    &__back {
      position: absolute;
      left: .18rem;
      font-size: .22rem;
    }
  }

  &__receiver {
    position: absolute;
    left: .18rem;
    right: .18rem;
    bottom: 0;
    height: 1.11rem;
    background: #FFF;
    border-radius: .04rem;

    &__title {
      line-height: .22rem;
      padding: .16rem 0 .14rem .16rem;
      font-size: .16rem;
      color: #333;
    }

    &__address {
      line-height: .2rem;
      padding: 0 .4rem 0 .16rem;
      font-size: .14rem;
      color: #333;
    }

    &__info {
      padding: .06rem 0 0 .16rem;

      &__name {
        margin-right: .06rem;
        line-height: .18rem;
        font-size: .12rem;
        color: #666;
      }
    }

    &__icon {
      transform: rotate(180deg);
      position: absolute;
      right: .16rem;
      top: .5rem;
      color: #666;
      font-size: .2rem;
    }
  }
}

.products {
  margin: .16rem .18rem .55rem .18rem;
  background: #FFF;

  &__title {
    font-size: .2rem;
  }

  &__list {
    height: 4rem;
    overflow-y: scroll; // 超出区域可以上下滚
    flex: 1; // 右侧填满
  }

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
    }
  }
}
</style>
