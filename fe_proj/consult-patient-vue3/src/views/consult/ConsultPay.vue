<script setup lang="ts">
import { getConsultOrderPre } from '@/services/consult'
import { getPatientDetail } from '@/services/user'
import { useConsultStore } from '@/stores'
import type { ConsultOrderPreData } from '@/types/consult'
import type { Patient } from '@/types/user'
import { onMounted, ref } from 'vue'

// 获取预支付信息
const payInfo = ref<ConsultOrderPreData>()
const store = useConsultStore()
const loadData = async () => {
  const res = await getConsultOrderPre({
    type: store.consult.type, // 问诊类型
    illnessType: store.consult.illnessType, // 问诊级别
  })
  payInfo.value = res.data
}
// 获取患者详情
const patient = ref<Patient>()
const loadPatient = async () => {
  if (!store.consult.patientId) return
  const res = await getPatientDetail(store.consult.patientId)
  patient.value = res.data
}

// 数据加载
onMounted(() => {
  loadData()
  loadPatient()
})
</script>

<template>
  <div class="consult-pay-page">
    <cp-nav-bar title="支付" />
    <div class="pay-info">
      <p class="tit">图文问诊 49 元</p>
      <img class="img" src="@/assets/avatar-doctor.svg" />
      <p class="desc">
        <span>极速问诊</span>
        <span>自动分配医生</span>
      </p>
    </div>
    <van-cell-group>
      <van-cell title="优惠券" value="`-¥10`" />
      <van-cell title="积分抵扣" value="`-¥10`" />
      <van-cell title="实付款" value="`¥29`" class="pay-price" />
    </van-cell-group>
    <div class="pay-space"></div>
    <van-cell-group>
      <van-cell
        title="患者信息"
        value="`${patient.name} | ${patient.genderValue} | ${patient.age}岁`"
      ></van-cell>
      <van-cell title="病情描述" label="store.consult.illnessDesc"></van-cell>
    </van-cell-group>
    <div class="pay-schema">
      <van-checkbox> 我已同意 <span class="text">支付协议</span> </van-checkbox>
    </div>
    <van-submit-bar button-type="primary" :price="100" button-text="立即支付" text-align="left" />
    <!-- 支付抽屉，控制面板 -->
    <cp-pay-sheet
      order-id="orderId"
      actual-payment="payInfo.actualPayment"
      on-close="onClose"
      pay-callback="/room"
    ></cp-pay-sheet>
  </div>
  <div class="consult-pay-page">
    <cp-nav-bar title="支付" />
    <!-- 骨架组件 -->
    <van-skeleton title :row="10" style="margin-top: 18px" />
  </div>
</template>

<style lang="scss" scoped>
.consult-pay-page {
  padding: 46px 0 50px 0;
}
.pay-info {
  display: flex;
  padding: 15px;
  flex-wrap: wrap;
  align-items: center;
  .tit {
    width: 100%;
    font-size: 16px;
    margin-bottom: 10px;
  }
  .img {
    margin-right: 10px;
    width: 38px;
    height: 38px;
    border-radius: 4px;
    overflow: hidden;
  }
  .desc {
    flex: 1;
    > span {
      display: block;
      color: var(--cp-tag);
      &:first-child {
        font-size: 16px;
        color: var(--cp-text2);
      }
    }
  }
}
.pay-price {
  ::v-deep() {
    .vam-cell__title {
      font-size: 16px;
    }
    .van-cell__value {
      font-size: 16px;
      color: var(--cp-price);
    }
  }
}
.pay-space {
  height: 12px;
  background-color: var(--cp-bg);
}
.pay-schema {
  height: 56px;
  display: flex;
  align-items: center;
  justify-content: center;
  .text {
    color: var(--cp-primary);
  }
}
::v-deep() {
  .van-submit-bar__button {
    font-weight: normal;
    width: 160px;
  }
}
</style>
