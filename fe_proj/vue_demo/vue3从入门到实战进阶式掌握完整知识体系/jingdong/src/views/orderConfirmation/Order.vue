<template>
  <!-- 底部 订单 -->
  <div class="order">
    <div class="order__price">实付金额 <b>￥{{ calculations.price }}</b></div>
    <div class="order__btn" @click="() => handleSubmitOrder(true)">提交订单</div>
  </div>
  <!-- toast -->
  <div class="mask" v-show="showConfirm" @click="() => handleSubmitOrder(false)">
    <div class="mask__content" @click.stop>
      <h3 class="mask__content__title">确认离开收银台吗?</h3>
      <p class="mask__content__desc">请尽快完成支付，否则将被取消</p>
      <div class="mask__content__btns">
        <div class="mask__content__btn mask__content__btn--first" @click="() => handleConfirmlOrder(true)">取消订单</div>
        <div class="mask__content__btn mask__content__btn--last" @click="() => handleConfirmlOrder(false)">确认支付</div>
      </div>
    </div>
  </div>
</template>

<script>
import { post } from '@/utils/request'
import { useStore } from 'vuex'
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useCommonCartEffect } from '@/effects/cartEffects'

export default {
  name: 'OrderConfirmation',
  setup() {
    const route = useRoute()
    const router = useRouter()
    const store = useStore()

    const showConfirm = ref(false) // 是否展示mask（蒙层）

    const shopId = parseInt(route.params.id, 10)
    const { calculations, shopName, productList } = useCommonCartEffect(shopId)

    const handleConfirmlOrder = async (isCanceled) => {
      const products = []
      for (const i in productList.value) {
        const product = productList.value[i]
        products.push({
          id: parseInt(product._id, 10),
          num: product.count
        })
      }
      console.log(products, '--- products ---')
      try {
        const reqBody = {
          addressId: 1,
          shopId,
          shopName: shopName.value,
          isCanceled: isCanceled, // true, 取消订单; false, 下单
          products
        }
        console.log('-- reqBody --', reqBody)
        const result = await post('/api/order', reqBody)
        console.log('result', result)
        if (result?.errno === 0) {
          store.commit('cleanCartProducts', { shopId }) // 清空该shopId 的购物车
          router.push({ name: 'Home' })
        } else {
          alert('登录失败')
        }
      } catch (error) {
        console.log('请求失败', error.message)
        alert('请求失败')
      }
      // 关闭蒙层
      showConfirm.value = false
    }

    const handleSubmitOrder = (status) => {
      showConfirm.value = status
    }
    return { calculations, showConfirm, handleConfirmlOrder, handleSubmitOrder }
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

// 弹窗
.mask {
  z-index: 1;
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  top: 0;
  background: rgba(0, 0, 0, .5);

  &__content {
    // 以下4行用来绝对定位的居中
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 3rem;
    height: 1.56rem;
    background: #FFF;
    text-align: center;
    border-radius: .04rem;

    &__title {
      margin: .24rem 0 0 0;
      line-height: .26rem;
      font-size: .18rem;
      color: #333;
    }

    &__desc {
      margin: .08rem 0 0 0;
      font-size: .14rem;
      color: #666;
    }

    &__btns {
      margin: .24rem .58rem;
      display: flex;
    }

    &__btn {
      flex: 1; // 等分
      width: .8rem;
      line-height: .32rem;
      border-radius: .16rem;
      font-size: .14rem;

      &--first {
        margin-right: .12rem;
        border: .01rem solid #4FB0F9;
      }

      &--last {
        margin-left: .12rem;
        background: #4FB0F9;
        color: #FFF;
      }
    }
  }

}
</style>
