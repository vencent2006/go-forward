<script setup lang="ts">
import RoomStatus from './components/RoomStatus.vue'
import RoomAction from './components/RoomAction.vue'
import RoomMessage from './components/RoomMessage.vue'
import { io, Socket } from 'socket.io-client'
import { nextTick, onMounted, onUnmounted, ref } from 'vue'
import { baseURL } from '@/utils/request'
import { useUserStore } from '@/stores'
import { useRoute } from 'vue-router'
import type { Message, TimeMessages } from '@/types/room'
import { MsgType, OrderType } from '@/enums'
import type { ConsultOrderItem, Image } from '@/types/consult'
import { getConsultOrderDetail } from '@/services/consult'
import dayjs from 'dayjs'
import { showToast } from 'vant'
const consult = ref<ConsultOrderItem>()
const orderIdMock = '7134595992539136'
const loadConsult = async () => {
  // TODO mock 数据
  // const res = await getConsultOrderDetail(route.query.id as string) // as string 类型断言
  const res = await getConsultOrderDetail(orderIdMock as string) // as string 类型断言
  consult.value = res.data
}
const store = useUserStore()
const route = useRoute()
const list = ref<Message[]>([])
const initialMsg = ref(true) // 是不是第一次消息
let socket: Socket
onMounted(() => {
  // 获取订单数据
  loadConsult()

  // 连接服务器
  socket = io(baseURL, {
    auth: {
      token: `Bearer ${store.user?.token}`,
    },
    query: {
      // 订单id
      // orderId: route.query.orderId,
      // TODO 当前先固定了，不然chatMsgList取不到有内容的数据，后续要改掉
      orderId: orderIdMock,
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
    data.forEach((item, i) => {
      // 记录每一段消息中最早的消息时间，后去聊天记录需要使用
      if (i === 0) {
        // 最早的消息
        time.value = item.createTime
      }
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

    loading.value = false // 关闭刷新
    if (!arr.length) return showToast('没有更多聊天记录了')

    // 第一次消息的逻辑
    if (initialMsg.value) {
      // 已读最后一条消息
      socket.emit('updateMsgStatus', arr[arr.length - 1].id)
      // 第一次需要滚到到最新的消息
      nextTick(() => {
        // 等待dom更新
        window.scrollTo(0, document.body.scrollHeight)
        initialMsg.value = false // 不是第一次了
      })
    }
  })

  // 监听订单状态变化
  socket.on('statusChange', () => loadConsult())
  // 接收聊天消息
  socket.on('receiveChatMsg', async (event) => {
    // 已读该条消息
    socket.emit('updateMsgStatus', event.id)
    // 追加消息
    list.value.push(event)
    // 滚动到底部
    await nextTick() // 等待dom更新
    window.scrollTo(0, document.body.scrollHeight) // 滚动到底部
  })
})

onUnmounted(() => {
  // 关闭连接
  socket.close()
})

// 发送文字信息
const onSendText = (text: string) => {
  socket.emit('sendChatMsg', {
    from: store.user?.id,
    to: consult.value?.docInfo?.id,
    msgType: MsgType.MsgText,
    msg: {
      content: text,
    },
  })
}

// 发送图片信息
const onSendImage = (image: Image) => {
  socket.emit('sendChatMsg', {
    from: store.user?.id,
    to: consult.value?.docInfo?.id,
    msgType: MsgType.MsgImage, // 图片类型
    msg: {
      picture: image,
    },
  })
}

// 下拉刷新
const loading = ref(false)
const time = ref(dayjs().format('YYYY-MM-DD HH:mm:ss'))
const onRefresh = async () => {
  socket.emit('getChatMsgList', 20, time.value, consult.value?.id)
}
</script>
<template>
  <div class="room-page">
    <cp-nav-bar title="问诊室"></cp-nav-bar>
    <!-- 状态栏组件 -->
    <room-status :status="consult?.status" :countdown="consult?.countdown"></room-status>
    <van-pull-refresh v-model="loading" @refresh="onRefresh">
      <!-- 消息 -->
      <room-message v-for="item in list" :key="item.id" :item="item"></room-message>
    </van-pull-refresh>
    <!-- 操作栏 -->
    <room-action
      @send-text="onSendText"
      @send-image="onSendImage"
      :disabled="consult?.status !== OrderType.ConsultChat"
    ></room-action>
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
