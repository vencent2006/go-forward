<template>
  <!-- 一级 menu 菜单 -->
  <el-menu
    :default-active="activeMenu"
    :background-color="$store.getters.cssVar.menuBg"
    :text-color="$store.getters.cssVar.menuText"
    :active-text-color="$store.getters.cssVar.menuActiveText"
    :uniqueOpened="true"
    router
  >
    <sidebar-item
      v-for="item in routes"
      :key="item.path"
      :route="item"
    ></sidebar-item>
  </el-menu>
</template>
<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { filterRouters, generateMenus } from '@/utils/route'
import SidebarItem from './SidebarItem'

const router = useRouter()
// console.log('router.getRoutes()', router.getRoutes())
const routes = computed(() => {
  const filterRoutes = filterRouters(router.getRoutes())
  // console.log('filterRoutes', filterRoutes)
  return generateMenus(filterRoutes)
})
// console.log(JSON.stringify(routes.value))
// console.log(routes.value)

// 计算高亮 menu 的方法
const route = useRoute()
const activeMenu = computed(() => {
  const { path } = route
  return path
})
</script>
