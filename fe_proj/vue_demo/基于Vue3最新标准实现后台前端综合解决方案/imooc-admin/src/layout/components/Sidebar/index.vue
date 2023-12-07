<template>
  <div class="">
    <div class="logo-container">
      <el-avatar
        :size="logoHeight"
        shape="square"
        src="https://img2.baidu.com/it/u=648933335,1401463189&fm=253&fmt=auto&app=138&f=JPEG?w=400&h=400"
      />
      <h1 class="logo-title" v-if="$store.getters.sidebarOpened">admin</h1>
    </div>
    <el-scrollbar>
      <sidebar-menu></sidebar-menu>
    </el-scrollbar>
  </div>
</template>

<script setup>
import SidebarMenu from './SidebarMenu'
import {} from 'vue'
const logoHeight = 44
</script>
<style lang="scss" scoped>
.logo-container {
  height: v-bind(logoHeight) + 'px'; // 使用vue3.2新特性: 组件状态驱动的动态 `CSS` 值
  padding: 10px 0 22px 0;
  display: flex;
  align-items: center;
  justify-content: center;
  .logo-title {
    margin-left: 10px;
    color: #fff;
    font-weight: 600;
    line-height: 50px;
    font-size: 16px;
    white-space: nowrap;
  }
}
</style>
