<template>
  <!-- docker -->
  <div class="docker">
    <!-- BEM命名 block__element--modifier -->
    <!-- 第一项设置了active属性 -->
    <div v-for="(item, index) in dockerList" :key="index"
      :class="{ 'docker__item': true, 'docker__item--active': index === currentIndex }">
      <router-link :to="item.to">
        <div class="iconfont" v-html="item.icon" />
        <div class="docker__item__title">{{ item.text }}</div>
      </router-link>
    </div>

  </div>
</template>

<script>
export default {
  name: 'Docker',
  props: ['currentIndex'],
  setup() {
    const dockerList = [
      { icon: '&#xe6f3;', text: '首页', to: { name: 'Home' } },
      { icon: '&#xe7e5;', text: '购物车', to: { name: 'CartList' } },
      { icon: '&#xe61e;', text: '订单', to: { name: 'OrderList' } },
      { icon: '&#xe660;', text: '我的', to: { name: 'Home' } }
    ]
    return { dockerList }
  }
}
</script>

<style lang="scss" scoped>
@import '@/style/variables.scss';

.docker {
  display: flex;
  box-sizing: border-box;
  padding: 0 .18rem;
  position: absolute;
  left: 0;
  bottom: 0;
  width: 100%;
  height: .49rem;
  border-top: .01rem solid $content-bgColor;

  &__item {
    flex: 1;
    text-align: center;

    a {
      color: $content-fontcolor;
      text-decoration: none;
    }

    .iconfont {
      font-size: .18rem;
      margin: .07rem 0 .02rem 0;
    }

    &--active {
      a {
        color: #1FA4FC;
      }
    }

    &__title {
      font-size: .2rem;
      // 老师说，移动端最小显示像素是12px，如果想显示10px，就需要用transform进行缩放
      transform: scale(.5, .5);
      transform-origin: center top;
    }
  }
}
</style>
