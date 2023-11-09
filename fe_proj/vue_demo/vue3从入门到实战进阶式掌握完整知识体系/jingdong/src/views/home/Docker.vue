<template>
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

<script>
export default {
  name: 'Docker'
}
</script>

<style lang="scss">
@import '../../style/variables.scss';

.docker {
  display: flex;
  box-sizing: border-box;
  padding: 0 .18rem;
  position: absolute;
  left: 0;
  bottom: 0;
  width: 100%;
  height: .49rem;
  border-top: .01rem solid $content-bgColor;
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
