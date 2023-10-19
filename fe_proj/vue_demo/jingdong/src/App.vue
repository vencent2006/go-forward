<template>
  <div class="wrapper">
    <div class="position">
      <span class="iconfont position__icon">&#xe64c;</span>
      北京理工大学国防科技园2号楼10层
      <span class="iconfont position__notice">&#xe66a;</span>
    </div>
  </div>
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
.wrapper {
  position: absolute;
  left: 0;
  top: 0;
  bottom: .5rem;
  right: 0;
  padding: 0 .18rem;
  // background: blue;
}
.position {
  position: relative;
  padding: .16rem 0;
  line-height: .22rem;
  font-size: .16rem;
  .position__icon { // 优先级更高 会编译成 .position .position__icon
    position: relative;
    top: .01rem;
    font-size: .2rem;
  }
  .position__notice {
    position: absolute;
    right: 0;
    top: .01rem;
    font-size: .2rem;
  }
}
.docker {
  display: flex;
  box-sizing: border-box; // css中的width所包含的是盒子模型中的content的宽度+padding+border的宽度。此时和怪异盒子模型的表现一致
  position: absolute;
  padding: 0 .18rem; // 0 是上下，0.18rem是左右
  left: 0;
  bottom: 0;
  width: 100%;
  height: .49rem;
  border-top: .01rem solid #F1F1F1;
  // background: green;

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
