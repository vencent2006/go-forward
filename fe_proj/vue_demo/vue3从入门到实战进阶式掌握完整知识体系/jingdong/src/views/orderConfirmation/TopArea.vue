<template>
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
</template>

<script>
import { useBackRouterEffect } from '@/effects/routeEffects'

export default {
  name: 'TopArea',

  setup() {
    const { handleBackClick } = useBackRouterEffect()
    return { handleBackClick }
  }
}
</script>

<style lang='scss' scoped>
@import '@/style/variables.scss';

.top {
  position: relative;
  height: 1.96rem;
  background-size: 100% 1.59rem;
  background-image: linear-gradient(0deg, rgba(0, 145, 255, 0.00) 4%, #0091FF 50%); // 渐变
  background-repeat: no-repeat; // 不要重复

  &__header {
    position: relative;
    padding-top: .2rem;
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
    background: $bgColor;
    border-radius: .04rem;

    &__title {
      line-height: .22rem;
      padding: .16rem 0 .14rem .16rem;
      font-size: .16rem;
      color: $content-fontcolor;
    }

    &__address {
      line-height: .2rem;
      padding: 0 .4rem 0 .16rem;
      font-size: .14rem;
      color: $content-fontcolor;
    }

    &__info {
      padding: .06rem 0 0 .16rem;

      &__name {
        margin-right: .06rem;
        line-height: .18rem;
        font-size: .12rem;
        color: $medium-fontColor;
      }
    }

    &__icon {
      transform: rotate(180deg);
      position: absolute;
      right: .16rem;
      top: .5rem;
      color: $medium-fontColor;
      font-size: .2rem;
    }
  }
}
</style>
