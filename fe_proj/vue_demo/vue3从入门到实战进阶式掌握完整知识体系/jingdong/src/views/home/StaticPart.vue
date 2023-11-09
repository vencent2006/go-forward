<template>
  <div class="position">
    <span class="iconfont position__icon">&#xe619;</span>
    北京理工大学国防科技园2号楼10层
    <span class="iconfont position__notice">&#xe60b;</span>
  </div>
  <div class="search">
    <span class="iconfont">&#xe62d;</span>
    <span class="search__text">山姆会员商店优惠商品</span>
  </div>
  <div class="banner">
    <img class="banner__img" src="http://www.dell-lee.com/imgs/vue3/banner.jpg" alt="">
  </div>
  <div class="icons">
    <div v-for="item in iconList" :key="item.desc" class="icons__item">
      <img :src="item.imgUrl" alt="" class="icons__item__img">
      <p class="icons__item__desc">{{ item.desc }}</p>
    </div>
  </div>
  <div class="gap"></div>
</template>

<script>
export default {
  name: 'StaticPart',
  setup() {
    const iconList = [
      { imgUrl: 'http://www.dell-lee.com/imgs/vue3/超市.png', desc: '超市便利' },
      { imgUrl: 'http://www.dell-lee.com/imgs/vue3/菜市场.png', desc: '菜市场' },
      { imgUrl: 'http://www.dell-lee.com/imgs/vue3/水果店.png', desc: '水果店' },
      { imgUrl: 'http://www.dell-lee.com/imgs/vue3/鲜花.png', desc: '鲜花绿植' },
      { imgUrl: 'http://www.dell-lee.com/imgs/vue3/医药健康.png', desc: '医药健康' },
      { imgUrl: 'http://www.dell-lee.com/imgs/vue3/家居.png', desc: '家居时尚' },
      { imgUrl: 'http://www.dell-lee.com/imgs/vue3/蛋糕.png', desc: '烘焙蛋糕' },
      { imgUrl: 'http://www.dell-lee.com/imgs/vue3/签到.png', desc: '签到' },
      { imgUrl: 'http://www.dell-lee.com/imgs/vue3/大牌免运.png', desc: '大牌免运' },
      { imgUrl: 'http://www.dell-lee.com/imgs/vue3/红包.png', desc: '红包套餐' }
    ]
    return { iconList }
  }
}
</script>

<style lang="scss">
@import '../../style/variables.scss';
@import '../../style/mixins.scss';

// 1. position 位置信息
.position {
  position: relative;
  padding: .16rem .24rem .16rem 0;
  // background: pink;
  line-height: .22rem;
  font-size: .16rem;
  @include ellipsis; // 从mixin中引入

  // 提升下优先级，直接用.，而不是&
  .position__icon {
    position: relative;
    top: .01rem;
    font-size: .2rem;
  }

  .position__notice {
    position: absolute;
    right: 0;
    top: .17rem;
    font-size: .2rem;
  }

  color: $content-fontcolor;
}

// 2. search 搜索
.search {
  margin-bottom: .12rem;
  line-height: .32rem; // 行高自动撑开
  background: #F5F5F5;
  color: #b7b7b7;
  border-radius: .16rem;

  .iconfont {
    display: inline-block;
    font-size: .16rem;
    padding: 0 .08rem 0 .16rem;
  }

  &__text {
    display: inline-block;
    font-size: .14rem;
  }
}

// 1.4 banner
.banner {
  // skill: 下面三行是为了避免抖动，提前占据好图片的空间
  height: 0;
  overflow: hidden;
  padding-bottom: 25.4%;

  &__img {
    width: 100%;
  }
}

// 1.5 icons
.icons {
  display: flex;
  flex-wrap: wrap;
  margin-top: .16rem;

  &__item {
    width: 20%;

    &__img {
      display: block;
      width: .4rem;
      height: .4rem;
      margin: 0 auto;
    }

    &__desc {
      margin: .06rem 0 .16rem 0;
      text-align: center;
      color: $content-fontcolor;
    }
  }
}

// 1.6 gap
.gap {
  margin: 0 -.18rem;
  height: .1rem;
  background: $content-bgColor;
}
</style>
