<template>
  <div class="top">
    <div class="top__bgcolor"></div>
    <div class="top__header">
      <!-- 后退 -->
      <div class="top__header__back iconfont">&#xe6f2;</div>
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
      <div class="top__receiver__enter iconfont">&#xe6f2;</div>
    </div>
  </div>
</template>

<script>
export default {}
</script>

<style lang='scss' scoped>
.top {
  height: 1.96rem;
  background-size: 100% 1.59rem;
  background-image: linear-gradient(0deg, rgba(0, 145, 255, 0.00) 4%, #0091FF 50%); // 渐变
  background-repeat: no-repeat; // 不要重复

  &__header {
    position: relative;
    padding-top: .4rem;
    line-height: .24rem;
    color: #FFF;
    text-align: center;
    font-size: .16rem;

    &__back {
      position: absolute;
      left: .18rem;
      font-size: .22rem;
    }
  }

}
</style>
