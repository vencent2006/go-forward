<script setup lang="ts">
import RoomStatus from './components/RoomStatus.vue'
import RoomAction from './components/RoomAction.vue'
import RoomMessage from './components/RoomMessage.vue'
import { io, Socket } from 'socket.io-client'
import { onMounted, onUnmounted, ref } from 'vue'
import { baseURL } from '@/utils/request'
import { useUserStore } from '@/stores'
import { useRoute } from 'vue-router'
import type { Message, TimeMessages } from '@/types/room'
import { MsgType } from '@/enums'

const store = useUserStore()
const route = useRoute()
const list = ref<Message[]>([])
let socket: Socket
onMounted(() => {
  // 连接服务器
  socket = io(baseURL, {
    auth: {
      token: `Bearer ${store.user?.token}`,
    },
    query: {
      // 订单id
      // orderId: route.query.orderId,
      // TODO 当前先固定了，不然chatMsgList取不到有内容的数据，后续要改掉
      orderId: '7133337225187328',
    },
  })
  socket.on('connect', () => {
    console.log('连接成功')
  })
  socket.on('disconnect', () => {
    console.log('连接关闭')
  })
  socket.on('error', () => {
    console.log('发生错误')
  })
  // 获取聊天记录，如果是第一次（默认消息）
  socket.on('chatMsgList', ({ data }: { data: TimeMessages[] }) => {
    // data 数据 ===> [{createTime}, ...items]
    const arr: Message[] = []
    data.forEach((item) => {
      arr.push({
        msgType: MsgType.Notify, // 通用的消息
        msg: {
          content: item.createTime,
        },
        createTime: item.createTime,
        id: item.createTime,
      })
      arr.push(...item.items)
    })
    console.log(arr)
    list.value.unshift(...arr) // 往前添加 unshift
  })
})

onUnmounted(() => {
  // 关闭连接
  socket.close()
})
</script>
<template>
  <div class="room-page">
    <cp-nav-bar title="问诊室"></cp-nav-bar>
    <!-- 状态栏组件 -->
    <room-status></room-status>
    <!-- 消息 -->
    <room-message></room-message>
    <!-- 操作栏 -->
    <room-action :disabled="true"></room-action>
  </div>
</template>
<style lang="scss" scoped>
.room-page {
  padding-top: 90px;
  padding-bottom: 60px;
  min-height: 100vh;
  box-sizing: border-box;
  background-color: var(--cp-bg);
  .van-pull-refresh {
    width: 100%;
    min-height: calc(100vh - 150px);
  }
}
</style>
