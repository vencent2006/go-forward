<template>
  <div class="docker">
    <!-- BEM -->
    <!-- 写法: block__element--modifier -->
    <div class="docker__item docker__item--active">
      <div class="iconfont">&#xe651;</div>
      <div class="docker__title">首页</div>
    </div>
    <div class="docker__item">
      <div class="iconfont">&#xe645;</div>
      <div class="docker__title">购物车</div>
    </div>
    <div class="docker__item">
      <div class="iconfont">&#xe605;</div>
      <div class="docker__title">订单</div>
    </div>
    <div class="docker__item">
      <div class="iconfont">&#xe660;</div>
      <div class="docker__title">我的</div>
    </div>
  </div>
</template>

<style lang="scss">
.docker {
  display: flex;
  box-sizing: border-box; // css中的width所包含的是盒子模型中的content的宽度+padding+border的宽度。此时和怪异盒子模型的表现一致
  position: absolute;
  padding: 0 .18rem; // 0 是上下，0.18rem是左右
  left: 0;
  bottom: 0;
  width: 100%;
  height: .49rem;
  border-top: 1px solid #F1F1F1;
  background: green;

  &__item {
    flex: 1; // 就是 flex-grow 表示占1份
    text-align: center;
    .iconfont {
      margin: .07rem 0 .02rem 0;
      font-size: .18rem;
    }
    &--active {
      color: #1FA4FC;
    }
  }

  &__title {
    font-size: .2rem;
    transform: scale(.5, .5); // 大小
    transform-origin: center top; // 位置
  }
}
</style>
