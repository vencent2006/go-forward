<script setup lang="ts">
const fn = () => {
  console.log('click right')
}
</script>

<template>
  <div class="login-page">
    <cp-nav-bar @click-right="fn" title="登录" right-text="注册"></cp-nav-bar>
    <van-button>van-button</van-button>
    <van-checkbox>van-checkbox</van-checkbox>
  </div>
</template>

<style lang="scss" scoped></style>
