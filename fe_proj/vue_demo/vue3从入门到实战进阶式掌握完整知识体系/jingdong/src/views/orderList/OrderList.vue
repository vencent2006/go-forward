<template>
  <div class="wrapper">
    <div class="title">我的订单</div>
    <div class="orders">
      <div class="order">
        <div class="order__title">
          沃尔玛
          <span class="order__status">已取消</span>
        </div>
        <div class="order__content">
          <div class="order__content__imgs">
            <img class="order__content__img" src="">
            <img class="order__content__img" src="">
            <img class="order__content__img" src="">
            <img class="order__content__img" src="">
          </div>
          <div class="order__info">
            <div class="order__info__price">￥36.88</div>
            <div class="order__info__count">共2件</div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <Docker :currentIndex="2" />
</template>

<script>
import Docker from '@/components/Docker.vue'
export default {
  name: 'OrderList',
  components: {
    Docker
  }
}
</script>

<style lang='scss' scoped>
// 老师说，wrapper可以抽象到mixin里
.wrapper {
  overflow-y: auto; // 纵向滚动
  position: absolute;
  left: 0;
  top: 0;
  bottom: .5rem;
  right: 0;
  background: #EEE;
}

.title {
  line-height: .44rem;
  background: #FFF;
  font-size: .16rem;
  color: #333;
  text-align: center;
}
</style>
