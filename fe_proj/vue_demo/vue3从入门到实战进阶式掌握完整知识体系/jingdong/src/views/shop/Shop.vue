<template>
  <div class="wrapper">
    <div class="search">
      <div class="search__back iconfont" @click="handleBackClick">&#xe6f2;</div>
      <div class="search__content">
        <span class="search__content__icon iconfont">&#xe62d;</span>
        <input class="search__content__input" placeholder="请输入商品名称" />
      </div>
    </div>
    <ShopInfo :item="data.item" :hideBorder="true" />
  </div>
</template>

<script>
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import { get } from '@/utils/request'
import ShopInfo from '@/components/ShopInfo'
// https://fastmock.site/mock/ae8e9031947a302fed5f92425995aa19/jd/api/shop/hot-list
export default {
  name: 'Shop',
  components: { ShopInfo },
  setup() {
    const router = useRouter()
    const data = reactive({ item: {} })
    const getItemData = async () => {
      const result = await get('/api/shop/1')
      console.log(result)
      if (result?.errno === 0 && result?.data) {
        data.item = result.data
      }
    }
    getItemData()

    const handleBackClick = () => {
      router.back()
    }
    return { data, handleBackClick }
  }
}
</script>

<style lang="scss" scoped>
@import '../../style/variables.scss';

.wrapper {
  padding: 0 .18rem;
}

.search {
  display: flex;
  margin: .14rem 0 .16rem 0;
  line-height: .32rem;

  &__back {
    width: .3rem;
    font-size: .24rem;
    color: #B6B6B6;
  }

  &__content {
    flex: 1; // flex:1 直接填满剩余空间
    background: $search-bgColor;
    border-radius: .16rem;
    display: flex;

    &__icon {
      width: .44rem;
      text-align: center;
      color: $search-fontColor;
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
      color: $content-fontcolor;

      &::placeholder {
        color: $content-fontcolor;
      }
    }
  }
}
</style>
