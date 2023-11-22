<template>
  <div class="wrapper">
    <div class="title">我的订单</div>
    <div class="orders">
      <div class="order" v-for="(item, index) in list" :key="index">
        <div class="order__title">
          {{ item.shopName }}
          <span class="order__status">
            {{ item.isCanceled ? '已取消' : '已下单' }}</span>
        </div>
        <div class="order__content">
          <div class="order__content__imgs">
            <template v-for="(innerItem, innerIndex) in item.products" :key="innerIndex">
              <img class="order__content__img" :src="innerItem.product.img" v-if="innerIndex <= 3">
            </template>
          </div>
          <div class="order__content_info">
            <div class="order__content__price">￥{{ item.totalPrice }}</div>
            <div class="order__content__count">共{{ item.totalNumber }}件</div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <Docker :currentIndex="2" />
</template>

<script>
import { reactive, toRefs } from 'vue'
import { get } from '@/utils/request'
import Docker from '@/components/Docker.vue'

// 处理订单列表逻辑
const useOrderListEffect = (shopId) => {
  const data = reactive({ list: [] })
  const getOrders = async () => {
    const result = await get('/api/order')
    if (result?.errno === 0 && result?.data?.length) {
      const orderList = result.data
      orderList.forEach((order) => {
        const products = order.products || []
        let totalPrice = 0
        let totalNumber = 0
        products.forEach((productItem) => {
          totalPrice += (productItem?.product?.price * productItem?.orderSales) || 0
          totalNumber += productItem?.orderSales || 0
        })
        order.totalPrice = totalPrice.toFixed(2)
        order.totalNumber = totalNumber
      })
      console.log(orderList)
      data.list = result.data
    }
  }

  getOrders()

  const { list } = toRefs(data)
  // console.log(list, '++ list ++')
  // for (const i in list) {
  //   const shop = list[i]
  //   console.log(shop, '--shop--')
  //   shop.count = shop.products.length
  //   shop.totalPrice = 0
  //   for (const j in shop) {
  //     shop.totalPrice = shop.totalPrice + shop.products[j].price
  //   }
  //   shop.totalPrice.toFixed(2)
  // }
  return { list }
}

export default {
  name: 'OrderList',
  components: {
    Docker
  },
  setup() {
    const { list } = useOrderListEffect()
    console.log(list, '-- list --')
    return { list }
  }
}
</script>

<style lang='scss' scoped>
@import '@/style/variables.scss';

// 老师说，wrapper可以抽象到mixin里
.wrapper {
  overflow-y: auto; // 纵向滚动
  position: absolute;
  left: 0;
  top: 0;
  bottom: .5rem;
  right: 0;
  background: rgb(248, 248, 248);
}

.title {
  line-height: .44rem;
  background: #FFF;
  font-size: .16rem;
  color: #333;
  text-align: center;
}

.orders {}

.order {
  margin: .16rem .18rem;
  padding: .16rem;
  background: $bgColor;

  &__title {
    margin-bottom: .16rem;
    line-height: .22rem;
    font-size: .16rem;
    color: $content-fontcolor;
  }

  &__status {
    float: right;
    font-size: .14rem;
    color: $light-fontColor;
  }

  &__content {
    display: flex;

    &__imgs {
      flex: 1;
    }

    &__img {
      width: .4rem;
      height: .4rem;
      margin-right: .12rem;
    }

    &__info {
      width: .7rem
    }

    &__price {
      margin-bottom: .04rem;
      line-height: .2rem;
      font-size: .14rem;
      color: $highlight-fontColor;
      text-align: right;
    }

    &__count {
      line-height: .14rem;
      font-size: .12rem;
      color: $content-fontcolor;
      text-align: right;
    }

    &__price {
      margin-bottom: .04rem;
      line-height: .2rem;
      font-size: .14rem;
      color: $highlight-fontColor;
      text-align: right;
    }
  }
}
</style>
