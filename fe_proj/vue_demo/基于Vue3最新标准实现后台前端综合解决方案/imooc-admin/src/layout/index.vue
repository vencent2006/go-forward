<template>
  <div class="app-wrapper">
    <!-- 左侧 menu -->
    <!-- <sidebar class="sidebar-container"></sidebar> -->
    <div class="main-container">
      <div class="fixed-header">
        <!-- 顶部 navbar -->
        <!-- <navbar /> -->
      </div>
      <!-- 内容区 -->
      <!-- <app-main></app-main> -->
    </div>
    <!-- test -->
    <!-- <div><h1 :style="{ color: color }">JS 获取 SCSS 变量值</h1></div> -->
  </div>
</template>

<script setup>
// import Navbar from './components/Navbar.vue'
// import Sidebar from './components/Sidebar' // 因为是index.vue，所以可以省略
// import AppMain from './components/AppMain.vue'
import variables from '@/styles/variables.module.scss'
import { ref } from 'vue'

console.log(variables)
const color = ref(variables.menuText)
console.log(color)
</script>

<style lang="scss" scoped>
@import '~@/styles/mixin.scss';
@import '~@/styles/variables.module.scss';

.app-wrapper {
  @include clearfix;
  position: relative;
  height: 100%;
  width: 100%;
}

.fixed-header {
  position: fixed;
  top: 0;
  right: 0;
  z-index: 9;
  width: calc(100% - #{$sideBarWidth});
}
</style>
