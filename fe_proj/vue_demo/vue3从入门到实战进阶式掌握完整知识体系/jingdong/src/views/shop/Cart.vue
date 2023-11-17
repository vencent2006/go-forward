<template>
  <div class="cart">
    <div class="check">
      <div class="check__icon">
        <img src="http://www.dell-lee.com/imgs/vue3/basket.png" class="check__icon__img" />
        <div class="check__icon__tag">1</div>
      </div>
      <div class="check__info">
        总计：<span class="check__info__price">&yen;128</span>
      </div>
      <div class="check__btn">去结算</div>
    </div>
  </div>
</template>

<script>
export default {}
</script>

<style lang='scss' scoped>
@import '@/style/variables.scss';

.cart {
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
}

.check {
  display: flex;
  // box-sizing: border-box;
  height: .49rem;
  border-top: .01rem solid $content-bgColor;
  line-height: .49rem;

  &__icon {
    position: relative;
    width: .84rem;

    &__img {
      display: block;
      margin: .12rem auto;
      width: .28rem;
      height: .26rem;
    }

    &__tag {
      position: absolute;
      top: .04rem;
      right: .2rem;
      width: .2rem;
      height: .2rem;
      line-height: .2rem;
      background: $highlight-fontColor;
      border-radius: 50%;
      font-size: .12rem;
      text-align: center;
      color: $bgColor;
      transform: scale(.5);
    }
  }

  // 总计
  &__info {
    flex: 1;
    color: $content-fontcolor;
    font-size: .12rem;
    // font-weight: ;

    &__price {
      color: $highlight-fontColor;
      font-size: .18rem;
    }
  }

  // button
  &__btn {
    width: .98rem;
    font-size: .14rem;
    text-align: center;
    background: #4FB0F9;
    color: $bgColor;
  }
}
</style>
