<script setup lang="ts">
import { OrderType } from '@/enums'
defineProps<{
  status?: OrderType
  countdown?: number
}>()
</script>

<template>
  <div class="room-status">
    <div v-if="status === OrderType.ConsultWait" class="wait">
      已通知医生尽快接诊，24小时内医生未回复将自动退款
    </div>
    <div v-if="status === OrderType.ConsultChat" class="chat">
      <span>咨询中</span>
      <!-- van-count-down 是以毫秒为单位 -->
      <span> 剩余时间：<van-count-down v-if="countdown" :time="countdown * 1000" /> </span>
    </div>
    <div
      v-if="status === OrderType.ConsultComplete || status === OrderType.ConsultCancel"
      class="end"
    >
      <van-icon name="passed" /> 已结束
    </div>
  </div>
</template>

<style lang="scss" scoped>
.room-status {
  position: fixed;
  left: 0;
  top: 46px;
  height: 44px;
  width: 100%;
  background-color: #fff;
  font-size: 13px;
  z-index: 1;
  .wait {
    padding: 0 15px;
    background-color: var(--cp-plain);
    height: 100%;
    line-height: 44px;
    text-align: center;
    color: var(--cp-primary);
  }
  .chat {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 15px;
    span {
      &:first-child {
        color: var(--cp-primary);
      }
      &:last-child {
        color: var(--cp-text2);
        width: 130px;
        .van-count-down {
          display: inline;
        }
      }
    }
  }
  .end {
    display: flex;
    align-items: center;
    height: 44px;
    padding: 0 15px;
    font-weight: 500;
    .van-icon {
      font-size: 14px;
      margin-right: 4px;
      position: relative;
      top: 1px;
    }
  }
}
</style>
