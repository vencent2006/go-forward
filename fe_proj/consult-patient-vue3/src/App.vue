<script setup lang="ts">
import { Button as VantButton } from 'vant'
import { useUserStore } from './stores/user'
const store = useUserStore()
</script>
<template>
  <div>
    App {{ store.user }}
    <vant-button
      type="primary"
      @click="
        store.setUser({
          id: '1',
          avatar: '1',
          token: '1',
          mobile: '1',
          account: '1',
        })
      "
      >登录</vant-button
    >
    <vant-button type="primary" @click="store.delUser()">退出</vant-button>
  </div>
</template>
