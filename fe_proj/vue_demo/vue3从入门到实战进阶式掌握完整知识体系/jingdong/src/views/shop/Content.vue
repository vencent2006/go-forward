<template>
  <div class="content">
    <div class="category">
      <div class="category__item category__item--active">全部商品</div>
      <div class="category__item">秒杀</div>
      <div class="category__item">新鲜水果</div>
      <div class="category__item">休闲食品</div>
      <div class="category__item">时令蔬菜</div>
      <div class="category__item">肉蛋家禽</div>
    </div>
    <div class="product">
      <div class="product__item">
        <img class="product__item__img" src="http://www.dell-lee.com/imgs/vue3/near.png">
        <div class="product__item__detail">
          <h4 class="product__item__title">番茄250g/份</h4>
          <p class="product__item__sales">月售10件</p>
          <p class="product__item__price">
            <span class="product__item__price__yen">&yen;</span>33.6
            <span class="product__item__price__origin">&yen;66.6</span>
          </p>
        </div>
        <div class="product__number">
          <span class="product__number__minus">-</span>
          0
          <span class="product__number__plus">+</span>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
export default {}
</script>

<style lang='scss' scoped>
.content {
  display: flex;
  position: absolute;
  left: 0;
  right: 0;
  top: 1.5rem;
  bottom: .5rem;
}

.category {
  overflow-y: scroll;
  height: 100%;
  width: .76rem;
  background: #F5F5F5;

  &__item {
    line-height: .4rem;
    text-align: center;
    font-size: .14rem;
    color: #333;

    &--active {
      background: #FFF;
    }
  }
}

.product {
  overflow-y: scroll; // 超出区域可以上下滚
  flex: 1; // 右侧填满

  &__item {
    position: relative;
    display: flex;
    padding: .12rem 0;
    margin: 0 .16rem;
    border-bottom: .01rem solid #F1F1F1;

    &__img {
      width: .68rem;
      height: .68rem;
      margin-right: .16rem;
    }

    &__title {
      margin: 0;
      line-height: .2rem;
      font-size: .14rem;
      color: #333;
    }

    &__sales {
      margin: .06rem 0;
      line-height: .16rem;
      font-size: .12rem;
      color: #333;
    }

    &__price {
      margin: 0;
      line-height: .2rem;
      font-size: .14rem;
      color: #E93B3B;

      &__yen {
        font-size: .12rem;
      }

      &__origin {
        margin-left: .06rem;
        line-height: .2rem;
        font-size: .12rem;
        color: #999;
        text-decoration: line-through;
      }
    }

    .product__number {
      position: absolute;
      right: 0;
      bottom: .12rem;

      &__minus,
      &__plus {
        display: inline-block; // span
        width: .2rem;
        height: .2rem;
        line-height: .16rem;
        border-radius: 50%;
        font-size: .2rem;
        text-align: center;
      }

      &__minus {
        border: .01rem solid #666;
        color: #666;
        margin-right: .05rem;
      }

      &__plus {
        background: #0091FF;
        color: #FFF;
        margin-left: .05rem;
      }
    }
  }

}
</style>
