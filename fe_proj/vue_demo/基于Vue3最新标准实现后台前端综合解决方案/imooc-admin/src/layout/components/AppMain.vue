<template>
  <div class="app-main">
    <router-view></router-view>
  </div>
</template>

<script setup>
import { useRoute } from 'vue-router'
import { generateTitle } from '@/utils/i18n'
import { useStore } from 'vuex'
import { watch } from 'vue'
import { isTags } from '@/utils/tags'

const route = useRoute()

/**
 * 生成 title
 */
const getTitle = (route) => {
  let title = ''
  if (!route.meta) {
    // 处理无 meta 的路由
    const pathArr = route.path.split('/')
    title = pathArr[pathArr.length - 1] // 就是最后一项
  } else {
    title = generateTitle(route.meta.title)
  }
  return title
}

/**
 * 监听 路由变化
 */
const store = useStore()
watch(
  route,
  (to, from) => {
    if (!isTags(to.path)) return
    const { fullPath, meta, name, params, path, query } = to
    store.commit('app/addTagsViewList', {
      fullPath,
      meta,
      name,
      params,
      path,
      query,
      title: getTitle(to)
    })
  },
  {
    immediate: true // 不是惰性，而是立即生效
  }
)
</script>

<style lang="scss" scoped>
.app-main {
  // 浏览器可视区域的高度 100vh
  min-height: calc(100vh - 50px);
  width: 100%;
  position: relative;
  overflow: hidden;
  padding: 61px 20px 20px 20px;
  box-sizing: border-box;
}
</style>
