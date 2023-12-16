<template>
  <div
    class="app-wrapper"
    :class="[$store.getters.sidebarOpened ? 'openSidebar' : 'hideSidebar']"
  >
    <!-- 左侧 menu 使用了行内样式，方便变化主题色 -->
    <sidebar
      class="sidebar-container"
      :style="{ backgroundColor: variables.menuBg }"
    ></sidebar>
    <div class="main-container">
      <div class="fixed-header">
        <!-- 顶部 navbar -->
        <navbar />
        <!-- tags -->
        <tags-view></tags-view>
      </div>
      <!-- 内容区 -->
      <app-main></app-main>
    </div>
  </div>
</template>

<script setup>
import Navbar from './components/Navbar.vue'
import Sidebar from './components/Sidebar' // 因为是index.vue，所以可以省略
import AppMain from './components/AppMain.vue'
import TagsView from '@/components/TagsView'
import variables from '@/styles/variables.module.scss'
import { ref } from 'vue'

console.log(variables)
const color = ref(variables.menuBg)
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
  transition: width #{$sideBarDuration};
}

.hideSidebar .fixed-header {
  width: calc(100% - #{$hideSideBarWidth});
}
</style>
