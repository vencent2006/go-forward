<template>
  <!-- wrapper -->
  <div class="wrapper">
    <div class="position">
      <span class="iconfont position__icon">&#xe619;</span>
      北京理工大学国防科技园2号楼10层
      <span class="iconfont position__notice">&#xe60b;</span>
    </div>
  </div>
  <!-- docker -->
  <div class="docker">
    <!-- BEM命名 block__element--modifier -->
    <div class="docker__item docker__item--active">
      <div class="iconfont">&#xe6f3;</div>
      <div class="docker__item__title">首页</div>
    </div>
    <div class="docker__item">
      <div class="iconfont">&#xe7e5;</div>
      <div class="docker__item__title">购物车</div>
    </div>
    <div class="docker__item">
      <div class="iconfont">&#xe61e;</div>
      <div class="docker__item__title">订单</div>
    </div>
    <div class="docker__item">
      <div class="iconfont">&#xe660;</div>
      <div class="docker__item__title">我的</div>
    </div>
  </div>
</template>

<style lang="scss">
@import './style/variables.scss';
@import './style/mixins.scss';

// wrapper
.wrapper {
  position: absolute;
  left: 0;
  top: 0;
  bottom: .5rem;
  right: 0;
  padding: 0 .18rem;
}

.position {
  position: relative;
  padding: .16rem .24rem .16rem 0;
  background: pink;
  line-height: .22rem;
  font-size: .16rem;
  @include ellipsis; // 从mixin中引入

  // 提升下优先级，直接用.，而不是&
  .position__icon {
    position: relative;
    top: .01rem;
    font-size: .2rem;
  }

  .position__notice {
    position: absolute;
    right: 0;
    top: .17rem;
    font-size: .2rem;
  }

  color: $content-fontcolor;
}

// docker
.docker {
  display: flex;
  box-sizing: border-box;
  padding: 0 .18rem;
  position: absolute;
  left: 0;
  bottom: 0;
  width: 100%;
  height: .49rem;
  border-top: .01rem solid #F1F1F1;
  color: $content-fontcolor;

  &__item {
    flex: 1;
    text-align: center;

    .iconfont {
      font-size: .18rem;
      margin: .07rem 0 .02rem 0;
    }

    &--active {
      color: #1FA4FC;
    }

    &__title {
      font-size: .2rem;
      // 老师说，移动端最小显示像素是12px，如果想显示10px，就需要用transform进行缩放
      transform: scale(.5, .5);
      transform-origin: center top;
    }
  }
}
</style>
