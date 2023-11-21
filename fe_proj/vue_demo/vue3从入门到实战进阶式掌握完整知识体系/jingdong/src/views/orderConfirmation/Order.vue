<template>
  <!-- 底部 订单 -->
  <div class="order">
    <div class="order__price">实付金额 <b>￥{{ calculations.price }}</b></div>
    <div class="order__btn">提交订单</div>
  </div>
</template>

<script>
import { useRoute } from 'vue-router'
import { useCommonCartEffect } from '@/effects/cartEffects'

export default {
  name: 'OrderConfirmation',
  setup() {
    const route = useRoute()
    const shopId = route.params.id
    const { calculations } = useCommonCartEffect(shopId)
    return { calculations }
  }
}
</script>

<style lang='scss' scoped>
@import '@/style/variables.scss';

// 底部 订单
.order {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  background: $bgColor;
  height: .49rem;
  line-height: .49rem;
  display: flex;

  &__price {
    flex: 1;
    text-indent: .24rem;
    font-size: .14rem;
    color: $content-fontcolor;
  }

  &__btn {
    width: .98rem;
    background: #4FB0F9;
    color: $bgColor;
    text-align: center;
    font-size: .14rem;
  }
}
</style>
