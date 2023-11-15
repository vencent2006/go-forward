<template>
  <div class="shop">
    <img class="shop__img" :src="item.imgUrl">
    <div class="shop__content">
      <div class="shop__content__title">{{ item.name }}</div>
      <div class="shop__content__tags">
        <span class="shop__content__tags__tag">月销:{{ item.sales }}</span>
        <span class="shop__content__tags__tag">起送:{{ item.expressLimit }}</span>
        <span class="shop__content__tags__tag">基础运费:{{ item.expressPrice }}</span>
      </div>
      <p class="shop__content__highlight">{{ item.slogan }}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ShopInfo',
  props: ['item']
}
</script>

<style lang="scss" scoped>
@import '../style/variables.scss';

.shop {
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
</style>
