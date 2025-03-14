<script setup lang="ts">
import { IllnessTime, MsgType } from '@/enums'
import type { Message } from '@/types/room'
import { timeOptions, flagOptions } from '@/services/constants'
import type { Image } from '@/types/consult'
import { showImagePreview, showToast } from 'vant'
import { useUserStore } from '@/stores'
import dayjs from 'dayjs'
defineProps<{ item: Message }>()

// 获取患病时间
const getIllnessTime = (time: IllnessTime) => {
  return timeOptions.find((item) => item.value === time)?.label
}
// 获取是否就诊
const getConsultFlagText = (flag: 0 | 1) => {
  return flagOptions.find((item) => item.value === flag)?.label
}

// 预览图片
const onPreviewImage = (images?: Image[]) => {
  if (images && images.length)
    showImagePreview(images.map((item) => item.url)) // 图片预览
  else showToast('暂无图片')
}

const store = useUserStore()

const formatTime = (time: string) => dayjs(time).format('HH:mm')
</script>

<template>
  <!-- 患者卡片 -->
  <div class="msg msg-illness" v-if="item.msgType === MsgType.CardPat">
    <div class="patient van-hairline--bottom">
      <p>
        {{ item.msg.consultRecord?.patientInfo.name }} |
        {{ item.msg.consultRecord?.patientInfo.genderValue }} |
        {{ item.msg.consultRecord?.patientInfo.age }}岁
      </p>
      <p v-if="item.msg.consultRecord">
        {{ getIllnessTime(item.msg.consultRecord?.illnessTime) }} |
        {{ getConsultFlagText(item.msg.consultRecord?.consultFlag) }}
      </p>
    </div>
    <van-row>
      <van-col span="6">病情描述</van-col>
      <van-col span="18">{{ item.msg.consultRecord?.illnessDesc }}</van-col>
      <van-col span="6">图片</van-col>
      <van-col span="18" @click="onPreviewImage(item.msg.consultRecord?.pictures)"
        >点击查看</van-col
      >
    </van-row>
  </div>
  <!-- 通知-通用 -->
  <div class="msg msg-tip" v-if="item.msgType === MsgType.Notify">
    <div class="content">
      <span>{{ item.msg.content }}</span>
    </div>
  </div>
  <!-- 通知-温馨提示 -->
  <div class="msg msg-tip" v-if="item.msgType === MsgType.NotifyTip">
    <div class="content">
      <span class="green">温馨提示：</span>
      <span>{{ item.msg.content }}</span>
    </div>
  </div>
  <!-- 发送文字 -->
  <div class="msg msg-to" v-if="item.msgType === MsgType.MsgText && item.from === store.user?.id">
    <div class="content">
      <div class="time">{{ formatTime(item.createTime) }}</div>
      <div class="pao">{{ item.msg.content }}</div>
    </div>
    <van-image :src="item.fromAvatar" />
  </div>
  <!-- 发送图片 -->
  <div class="msg msg-to" v-if="item.msgType === MsgType.MsgImage && item.from === store.user?.id">
    <div class="content">
      <div class="time">{{ formatTime(item.createTime) }}</div>
      <van-image fit="contain" :src="item.msg.picture?.url" />
    </div>
    <van-image :src="item.fromAvatar" />
  </div>
  <!-- 接收图片 -->
  <div
    class="msg msg-from"
    v-if="item.msgType === MsgType.MsgImage && item.from !== store.user?.id"
  >
    <van-image :src="item.fromAvatar" />
    <div class="content">
      <div class="time">{{ formatTime(item.createTime) }}</div>
      <van-image fit="contain" :src="item.msg.picture?.url" />
    </div>
  </div>
  <!-- 接收文字 -->
  <div class="msg msg-from" v-if="item.msgType === MsgType.MsgText && item.from !== store.user?.id">
    <van-image :src="item.fromAvatar" />
    <div class="content">
      <div class="time">{{ formatTime(item.createTime) }}</div>
      <div class="pao">{{ item.msg.content }}</div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@import '@/styles/room.scss';
</style>
