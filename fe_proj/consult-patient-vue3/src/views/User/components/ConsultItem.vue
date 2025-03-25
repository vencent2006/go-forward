<script setup lang="ts">
import { OrderType } from '@/enums'
import { ConsultOrderItem } from '@/types/consult'
import { log } from 'console'
import { computed, ref } from 'vue'

const props = defineProps<{
  item: ConsultOrderItem
}>()

// 更多操作
const showPopover = ref(false)
const actions = computed(() => [
  {
    text: '查看处方',
    disabled: !props.item.prescriptionId,
  },
  {
    text: '删除处方',
  },
])
const onSelect = () => {
  console.log('点击了更多操作')
}
</script>
<template>
  <div class="consult-item">
    <div class="head van-hairline--bottom">
      <img class="img" src="@/assets/avatar-doctor.svg" />
      <p>极速问诊(自动分配医生)</p>
      <span> 待支付 </span>
    </div>
    <div class="body">
      <div class="body-row">
        <div class="body-label">病情描述</div>
        <div class="body-value">腹痛腹泻 胃部有些痉挛</div>
      </div>
      <div class="body-row">
        <div class="body-label">价格</div>
        <div class="body-value">¥ 39.00</div>
      </div>
      <div class="body-row">
        <div class="body-label">创建时间</div>
        <div class="body-value tip">2019-07-08 09:55:54</div>
      </div>
    </div>
    <div class="foot">
      <van-button class="gray" plain size="small" round>取消问诊</van-button>
      <van-button type="primary" plain size="small" round> 继续沟通 </van-button>
    </div>
    <div class="foot" v-if="item.status === OrderType.ConsultComplete">
      <!-- 更多组件 -->
      <div class="more">
        <van-popover
          v-model:show="showPopover"
          placement="top-start"
          :actions="actions"
          @select="onSelect"
        >
          <template #reference>更多</template>
        </van-popover>
      </div>
      <van-button class="gray" plain size="small" round :to="`/room?orderId=${item.id}`">
        问诊记录
      </van-button>
      <van-button v-if="item.evaluateId" class="gray" plain size="small" round>
        查看评价
      </van-button>
      <van-button v-else type="primary" plain size="small" round> 写评价 </van-button>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.consult-item {
  border-radius: 4px;
  box-shadow: 0px 0px 22px 0px rgba(245, 245, 245, 0.1);
  background-color: #fff;
  margin-bottom: 10px;
  .head {
    display: flex;
    align-items: center;
    height: 50px;
    padding: 0 15px;
    .img {
      width: 20px;
      height: 20px;
    }
    > p {
      flex: 1;
      font-size: 15px;
      padding-left: 10px;
    }
    > span {
      color: var(--cp-tag);
      &.orange {
        color: #f2994a;
      }
      &.green {
        color: var(--cp-primary);
      }
    }
  }
  .body {
    padding: 15px 15px 8px 15px;
    .body-row {
      display: flex;
      margin-bottom: 7px;
    }
    .body-label {
      width: 62px;
      font-size: 13px;
      color: var(--cp-tip);
    }
    .body-value {
      width: 250px;
      &.tip {
        color: var(--cp-tip);
      }
    }
  }
  .foot {
    padding: 0 15px 15px 15px;
    display: flex;
    justify-content: flex-end;
    align-items: center;
    .van-button {
      margin-left: 10px;
      padding: 0 16px;
      &.gray {
        color: var(--cp-text3);
        background-color: var(--cp-bg);
      }
    }
    .more {
      color: var(--cp-tag);
      flex: 1;
      font-size: 13px;
    }
  }
}
</style>
