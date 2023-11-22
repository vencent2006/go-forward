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
            <img class="order__content__img" src="http://www.dell-lee.com/imgs/vue3/tomato.png">
            <img class="order__content__img" src="http://www.dell-lee.com/imgs/vue3/tomato.png">
            <img class="order__content__img" src="http://www.dell-lee.com/imgs/vue3/tomato.png">
            <img class="order__content__img" src="http://www.dell-lee.com/imgs/vue3/tomato.png">
          </div>
          <div class="order__content_info">
            <div class="order__content__price">￥36.88</div>
            <div class="order__content__count">共2件</div>
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
    /*
    {
    "errno": 0,
    "data": [
        {
            "address": {
                "city": "北京",
                "department": "xx小区",
                "houseNumber": "门牌号",
                "name": "张三",
                "phone": "18611112222"
            },
            "shopId": "1",
            "shopName": "沃尔玛",
            "isCanceled": false,
            "products": [
                {
                    "orderSales": 5,
                    "product": {
                        "name": "番茄 250g / 份",
                        "img": "http://www.dell-lee.com/imgs/vue3/tomato.png",
                        "price": 33.6,
                        "sales": 6
                    }
                },
                {
                    "orderSales": 10,
                    "product": {
                        "name": "车厘子 500g / 份",
                        "img": "http://www.dell-lee.com/imgs/vue3/cherry.png",
                        "price": 33.6,
                        "sales": 6
                    }
                }
            ]
        }
    ],
    "message": "errno !== 0 时的错误信息"
}
    */
    if (result?.errno === 0 && result?.data?.length) {
      data.list = result.data
    }
  }

  getOrders()

  const { list } = toRefs(data)
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
