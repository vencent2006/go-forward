<template>
  <div class="nearby">
    <h3 class="nearby_title">附近店铺</h3>
    <div class="nearby__item">
      <img class="nearby__item__img" src="http://www.dell-lee.com/imgs/vue3/near.png">
      <div class="nearby__item__content">
        <div class="nearby__item__content__title">沃尔玛</div>
        <div class="nearby__item__content__tags">
          <span class="nearby__item__content__tags__tag">月售1万+</span>
          <span class="nearby__item__content__tags__tag">月售1万+</span>
          <span class="nearby__item__content__tags__tag">月售1万+</span>
        </div>
        <p class="nearby__item__content__highlight">VIP尊享89元减4元运费券(每月3张)</p>
      </div>
    </div>
    <div class="nearby__item">
      <img class="nearby__item__img" src="http://www.dell-lee.com/imgs/vue3/near.png">
      <div class="nearby__item__content">
        <div class="nearby__item__content__title">沃尔玛</div>
        <div class="nearby__item__content__tags">
          <span class="nearby__item__content__tags__tag">月售1万+</span>
          <span class="nearby__item__content__tags__tag">月售1万+</span>
          <span class="nearby__item__content__tags__tag">月售1万+</span>
        </div>
        <p class="nearby__item__content__highlight">VIP尊享89元减4元运费券(每月3张)</p>
      </div>
    </div>
    <div class="nearby__item">
      <img class="nearby__item__img" src="http://www.dell-lee.com/imgs/vue3/near.png">
      <div class="nearby__item__content">
        <div class="nearby__item__content__title">沃尔玛</div>
        <div class="nearby__item__content__tags">
          <span class="nearby__item__content__tags__tag">月售1万+</span>
          <span class="nearby__item__content__tags__tag">月售1万+</span>
          <span class="nearby__item__content__tags__tag">月售1万+</span>
        </div>
        <p class="nearby__item__content__highlight">VIP尊享89元减4元运费券(每月3张)</p>
      </div>
    </div>
    <div class="nearby__item">
      <img class="nearby__item__img" src="http://www.dell-lee.com/imgs/vue3/near.png">
      <div class="nearby__item__content">
        <div class="nearby__item__content__title">沃尔玛</div>
        <div class="nearby__item__content__tags">
          <span class="nearby__item__content__tags__tag">月售1万+</span>
          <span class="nearby__item__content__tags__tag">月售1万+</span>
          <span class="nearby__item__content__tags__tag">月售1万+</span>
        </div>
        <p class="nearby__item__content__highlight">VIP尊享89元减4元运费券(每月3张)</p>
      </div>
    </div>
    <div class="nearby__item">
      <img class="nearby__item__img" src="http://www.dell-lee.com/imgs/vue3/near.png">
      <div class="nearby__item__content">
        <div class="nearby__item__content__title">沃尔玛</div>
        <div class="nearby__item__content__tags">
          <span class="nearby__item__content__tags__tag">月售1万+</span>
          <span class="nearby__item__content__tags__tag">月售1万+</span>
          <span class="nearby__item__content__tags__tag">月售1万+</span>
        </div>
        <p class="nearby__item__content__highlight">VIP尊享89元减4元运费券(每月3张)</p>
      </div>
    </div>
    <div class="nearby__item">
      <img class="nearby__item__img" src="http://www.dell-lee.com/imgs/vue3/near.png">
      <div class="nearby__item__content">
        <div class="nearby__item__content__title">沃尔玛</div>
        <div class="nearby__item__content__tags">
          <span class="nearby__item__content__tags__tag">月售1万+</span>
          <span class="nearby__item__content__tags__tag">月售1万+</span>
          <span class="nearby__item__content__tags__tag">月售1万+</span>
        </div>
        <p class="nearby__item__content__highlight">VIP尊享89元减4元运费券(每月3张)</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Nearby'
}
</script>

<style lang="scss">
@import '../../style/variables.scss';

.nearby {
  &__title {
    margin: .16rem 0 .02rem 0;
    font-size: .18rem;
    font-weight: normal; // 不加粗展示
    color: $content-fontcolor;
  }

  &__item {
    display: flex;
    padding-top: .12rem 0;

    &__img {
      margin-right: .16rem;
      width: .56rem;
      height: .56rem;
    }

    &__content {
      flex: 1; // flex-grow:1 flex-shrink:1 flex-basis:auto
      padding-bottom: .12rem;
      border-bottom: 1px solid $content-bgColor;

      &__title {
        line-height: .22rem;
        font-size: .16rem;
        color: $content-fontcolor;
      }

      &__tags {
        margin-top: .08rem;
        line-height: .18rem;
        font-size: .13rem;
        color: $content-fontcolor;

        &__tag {
          margin-right: .16rem;
        }
      }

      &__highlight {
        margin: .08rem 0 0 0; // 去除p标签自带的margin
        line-height: .18rem;
        font-size: .13rem;
        color: #E93B3B;
      }
    }
  }
}
</style>
