<template>
  <!-- wrapper -->
  <div class="wrapper">
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
      <div class="icons__item">
        <img src="http://www.dell-lee.com/imgs/vue3/超市.png" alt="" class="icons__item__img">
        <p class="icons__item__desc">超市便利</p>
      </div>
      <div class="icons__item">
        <img src="http://www.dell-lee.com/imgs/vue3/菜市场.png" alt="" class="icons__item__img">
        <p class="icons__item__desc">菜市场</p>
      </div>
      <div class="icons__item">
        <img src="http://www.dell-lee.com/imgs/vue3/水果店.png" alt="" class="icons__item__img">
        <p class="icons__item__desc">水果店</p>
      </div>
      <div class="icons__item">
        <img src="http://www.dell-lee.com/imgs/vue3/鲜花.png" alt="" class="icons__item__img">
        <p class="icons__item__desc">鲜花绿植</p>
      </div>
      <div class="icons__item">
        <img src="http://www.dell-lee.com/imgs/vue3/医药健康.png" alt="" class="icons__item__img">
        <p class="icons__item__desc">医药健康</p>
      </div>
      <div class="icons__item">
        <img src="http://www.dell-lee.com/imgs/vue3/家居.png" alt="" class="icons__item__img">
        <p class="icons__item__desc">家居时尚</p>
      </div>
      <div class="icons__item">
        <img src="http://www.dell-lee.com/imgs/vue3/蛋糕.png" alt="" class="icons__item__img">
        <p class="icons__item__desc">烘焙蛋糕</p>
      </div>
      <div class="icons__item">
        <img src="http://www.dell-lee.com/imgs/vue3/签到.png" alt="" class="icons__item__img">
        <p class="icons__item__desc">签到</p>
      </div>
      <div class="icons__item">
        <img src="http://www.dell-lee.com/imgs/vue3/大牌免运.png" alt="" class="icons__item__img">
        <p class="icons__item__desc">大牌免运</p>
      </div>
      <div class="icons__item">
        <img src="http://www.dell-lee.com/imgs/vue3/红包.png" alt="" class="icons__item__img">
        <p class="icons__item__desc">红包套餐</p>
      </div>
    </div>
    <div class="gap">

    </div>
  </div>
  <!-- docker -->
  <div class="docker">
    <!-- BEM命名 block__element--modifier -->
    <div class="docker__item docker__item--active">
      <div class="iconfont">&#xe6f3;</div>
      <div class="docker__item__title">首页</div>
    </div>
    <div class="docker__item">
      <div class="iconfont">&#xe7e5;</div>
      <div class="docker__item__title">购物车</div>
    </div>
    <div class="docker__item">
      <div class="iconfont">&#xe61e;</div>
      <div class="docker__item__title">订单</div>
    </div>
    <div class="docker__item">
      <div class="iconfont">&#xe660;</div>
      <div class="docker__item__title">我的</div>
    </div>
  </div>
</template>

<style lang="scss">
@import './style/variables.scss';
@import './style/mixins.scss';

// 1. wrapper
.wrapper {
  position: absolute;
  left: 0;
  top: 0;
  bottom: .5rem;
  right: 0;
  padding: 0 .18rem;
}

// 1.2 position 位置信息
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

// 1.3 search 搜索
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

.banner {
  // skill: 下面三行是为了避免抖动，提前占据好图片的空间
  height: 0;
  overflow: hidden;
  padding-bottom: 25.4%;

  &__img {
    width: 100%;
  }
}

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

.gap {
  margin: 0 -.18rem;
  height: .1rem;
  background: $content-bgColor;
}

// 2. docker
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
  color: $content-fontcolor;

  &__item {
    flex: 1;
    text-align: center;

    .iconfont {
      font-size: .18rem;
      margin: .07rem 0 .02rem 0;
    }

    &--active {
      color: #1FA4FC;
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
