<template>
  <div class="wrapper">
    <!-- 顶部信息 -->
    <div class="top">
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
      <!-- 用wrapper包一下list，让list自己撑开高度 -->
      <div class="products__wrapper">
        <!-- 商品列表 -->
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
    <!-- 底部 订单 -->
    <div class="order">
      <div class="order__price">实付金额 <b>￥{{ calculations.price }}</b></div>
      <div class="order__btn">提交订单</div>
    </div>
  </div>
</template>

<script>
import { useRoute } from 'vue-router'
import { useCommonCartEffect } from '@/effects/cartEffects'
import { useBackRouterEffect } from '@/effects/routeEffects'

export default {
  name: 'OrderConfirmation',

  setup() {
    const route = useRoute()
    const shopId = route.params.id
    const { shopName, productList, calculations } = useCommonCartEffect(shopId)
    const { handleBackClick } = useBackRouterEffect()
    return { shopName, productList, calculations, handleBackClick }
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
  overflow-y: scroll; // 出现滚动条
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
  margin: .16rem .18rem .1rem .18rem;
  background: #FFF;

  &__title {
    padding: .16rem;
    font-size: .16rem;
    color: #333;
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
    background: #FFF;
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
      color: #000;
    }

    &__yen {
      font-size: .12rem;
    }
  }
}

// 底部 订单
.order {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  background: #FFF;
  height: .49rem;
  line-height: .49rem;
  display: flex;

  &__price {
    flex: 1;
    text-indent: .24rem;
    font-size: .14rem;
    color: #333;
  }

  &__btn {
    width: .98rem;
    background: #4FB0F9;
    color: #FFF;
    text-align: center;
    font-size: .14rem;
  }
}
</style>
