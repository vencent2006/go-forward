<script setup lang="ts">
import { passwordInputProps, Button as VantButton } from 'vant'
import { useUserStore } from './stores'
import axios from './utils/request'
const store = useUserStore()
const getUser = () => {
  axios.request({
    url: 'patient/myUser',
    method: 'GET',
  })
}
const login = () => {
  axios
    .request({
      url: 'login/password',
      method: 'POST',
      data: {
        mobile: '13211112222',
        password: 'abc12345',
      },
    })
    .then(() => {
      console.log('登录成功')
    })
    .catch((err) => {
      console.log('登录失败', err)
    })
}
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
    <vant-button type="primary" @click="getUser">获取用户信息</vant-button>
    <vant-button type="primary" @click="login">接口登录</vant-button>
  </div>
</template>
