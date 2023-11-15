<template>
  <div class="wrapper">
    <div class="search">
      <div class="search__back iconfont" @click="handleBackClick">&#xe6f2;</div>
      <div class="search__content">
        <span class="search__content__icon iconfont">&#xe62d;</span>
        <input class="search__content__input" placeholder="请输入商品名称" />
      </div>
    </div>
    <ShopInfo :item="item" :hideBorder="true" />
  </div>
</template>

<script>
import { useRouter } from 'vue-router'
import ShopInfo from '@/components/ShopInfo'
// https://fastmock.site/mock/ae8e9031947a302fed5f92425995aa19/jd/api/shop/hot-list
export default {
  name: 'Shop',
  components: { ShopInfo },
  setup() {
    const router = useRouter()
    const item = {
      _id: '1',
      name: '沃尔玛',
      imgUrl: 'http://www.dell-lee.com/imgs/vue3/near.png',
      sales: 10000,
      expressLimit: 0,
      expressPrice: 5,
      slogan: 'VIP尊享满89元减4元运费券'
    }

    const handleBackClick = () => {
      router.back()
    }
    return { item, handleBackClick }
  }
}
</script>

<style lang="scss" scoped>
.wrapper {
  padding: 0 .18rem;
}

.search {
  display: flex;
  margin: .2rem 0 .16rem 0;
  line-height: .32rem;

  &__back {
    width: .3rem;
    font-size: .24rem;
    color: #B6B6B6;
  }

  &__content {
    flex: 1; // flex:1 直接填满剩余空间
    background: #F5F5F5;
    border-radius: .16rem;
    display: flex;

    &__icon {
      width: .44rem;
      text-align: center;
      color: #B7B7B7;
    }

    &__input {
      display: block;
      width: 100%;
      padding-right: .2rem;
      border: none;
      outline: none;
      background: none;
      height: .32rem;
      font-size: .14rem;
      color: #333;

      &::placeholder {
        color: #333;
      }
    }
  }
}
</style>
