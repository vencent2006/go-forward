<template>
  <div>
    <a href="javascript:" @click="onToggle">
      <svg-icon :icon="isFullscreen ? 'exit-fullscreen' : 'fullscreen'" />
    </a>
  </div>
</template>
<script setup>
// 是否全屏
import { onMounted, onUnmounted, ref } from 'vue'
import screenfull from 'screenfull'

const isFullscreen = ref(false)

// 监听变化
const change = () => {
  isFullscreen.value = screenfull.isFullscreen
}

// 切换事件
const onToggle = () => {
  screenfull.toggle()
}

// 设置侦听器
onMounted(() => {
  screenfull.on('change', change)
})

// 删除侦听器
onUnmounted(() => {
  screenfull.off('change', change)
})
</script>

<style scoped lang="scss"></style>
