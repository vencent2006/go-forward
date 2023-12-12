<template>
  <router-view />
</template>

<script>
// 解决 element-plus 的 resize 报错 | start
// 参考 https://blog.csdn.net/qq_45112637/article/details/131740110
const debounce = (fn, delay) => {
  let timer = null
  return function () {
    const context = this
    const args = arguments
    clearTimeout(timer)
    timer = setTimeout(function () {
      fn.apply(context, args)
    }, delay)
  }
}

const _ResizeObserver = window.ResizeObserver
window.ResizeObserver = class ResizeObserver extends _ResizeObserver {
  constructor(callback) {
    callback = debounce(callback, 16)
    super(callback)
  }
}
// 解决 element-plus 的 resize 报错 | end
</script>

<style></style>
