<template>
  <div class="nearby">
    <h3 class="nearby_title">附近店铺</h3>
    <ShopInfo v-for="item in nearbyList" :key="item._id" :item="item" />
  </div>
</template>

<script>
import { ref } from 'vue'
import { get } from '@/utils/request'
import ShopInfo from '@/components/ShopInfo'

const useNearbyListEffect = () => {
  // 使用ref，成为响应式的数据，ref是基本类型, reactive 是复合类型
  const nearbyList = ref([])
  const getNearbyList = async () => {
    const result = await get('/api/shop/hot-list')
    if (result?.errno === 0 && result?.data?.length) {
      // ref 封装的赋值要加value
      nearbyList.value = result.data
    }
  }
  return { nearbyList, getNearbyList }
}

export default {
  name: 'Nearby',
  components: { ShopInfo },
  setup() {
    const { nearbyList, getNearbyList } = useNearbyListEffect()
    getNearbyList()
    return { nearbyList }
  }
}
</script>

<style lang="scss" scoped>
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
