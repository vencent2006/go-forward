<script setup lang="ts">
import { createConsultOrder, getConsultOrderPre } from '@/services/consult'
import { getPatientDetail } from '@/services/user'
import { useConsultStore } from '@/stores'
import type { ConsultOrderPreData, PartialConsult } from '@/types/consult'
import type { Patient } from '@/types/user'
import { onMounted, ref } from 'vue'
import { showConfirmDialog, showDialog, showToast } from 'vant'
import { onBeforeRouteLeave, useRouter } from 'vue-router'

// 获取预支付信息
const payInfo = ref<ConsultOrderPreData>()
const store = useConsultStore()
const loadData = async () => {
  const res = await getConsultOrderPre({
    type: store.consult.type, // 问诊类型
    illnessType: store.consult.illnessType, // 问诊级别
  })
  payInfo.value = res.data
  // 记录优惠券id
  store.setCoupon(res.data.couponId)
}
// 获取患者详情
const patient = ref<Patient>()
const loadPatient = async () => {
  if (!store.consult.patientId) return
  const res = await getPatientDetail(store.consult.patientId)
  patient.value = res.data
}

// 数据加载
type Key = keyof PartialConsult // keyof从一个对象当中获取所有的key，作为一个联合类型
onMounted(() => {
  // 生成订单需要的信息，不完整的时候需要提示
  const validKeys: Key[] = [
    'type',
    'illnessType',
    'depId',
    'illnessDesc',
    'illnessTime',
    'consultFlag',
    'patientId',
  ]
  const valid = validKeys.every((key) => store.consult[key] !== undefined)
  if (!valid) {
    // 校验不成功
    return showDialog({
      title: '温馨提示',
      message: '问诊信息不完整请重新填写，如有未支付的问诊订单可在问诊记录中继续支付',
      closeOnPopstate: false, // 主要作用是控制当用户通过浏览器的前进或后退按钮导航时，是否自动关闭当前的弹出窗口或模态框。
    }).then(() => {
      // 点击确定 跳到首页
      router.push('/')
    })
  }

  loadData()
  loadPatient()
})

// 同意
const agree = ref(false)

// 生成订单
const show = ref(false)
const paymentMethod = ref<0 | 1>() // 0 微信 1 支付宝
const loading = ref(false) // 加载中
const orderId = ref('')
const submit = async () => {
  // 先判断是否同意协议
  if (!agree.value) return showToast('请先同意支付协议')
  // 发送生成订单的请求
  loading.value = true
  const res = await createConsultOrder(store.consult)
  loading.value = false
  // 清理数据
  store.clear()
  // 记录订单id
  orderId.value = res.data.id
  // 显示支付面板
  show.value = true
}

// 提示1： 取消支付将无法获得医生回复，医生接诊名额有限，是否确认关闭？
// 提示2： 问诊信息不完整请重新填写，如有未支付的问诊订单可在问诊记录中继续支付

// 用户引导
onBeforeRouteLeave((to) => {
  // 有订单不允许回退
  if (orderId.value) return false
})
const router = useRouter()
const onClose = () => {
  // showConfirmDialog 是 vant 提供的一个弹窗组件，返回一个 Promise
  return showConfirmDialog({
    title: '提示',
    message: '取消支付将无法获得医生回复，医生接诊名额有限，是否确认关闭？',
    cancelButtonText: '狠心离开',
    confirmButtonText: '继续支付',
  })
    .then(() => {
      // 点击继续支付
      return false
    })
    .catch(() => {
      // 点击狠心离开
      orderId.value = '' // 清空订单id
      router.push('/user/consult') // 跳转页面
      return true // 关闭弹窗
    })
}
</script>

<template>
  <!-- 必须都加载到数据了，才能显示 -->
  <div class="consult-pay-page" v-if="payInfo && patient">
    <cp-nav-bar title="支付" />
    <div class="pay-info">
      <p class="tit">图文问诊 {{ payInfo.payment }} 元</p>
      <img class="img" src="@/assets/avatar-doctor.svg" />
      <p class="desc">
        <span>极速问诊</span>
        <span>自动分配医生</span>
      </p>
    </div>
    <van-cell-group>
      <van-cell title="优惠券" :value="`-¥${payInfo.couponDeduction}`" />
      <van-cell title="积分抵扣" :value="`-¥${payInfo.pointDeduction}`" />
      <van-cell title="实付款" :value="`¥${payInfo.actualPayment}`" class="pay-price" />
    </van-cell-group>
    <div class="pay-space"></div>
    <van-cell-group>
      <van-cell
        title="患者信息"
        :value="`${patient.name} | ${patient.genderValue} | ${patient.age}岁`"
      ></van-cell>
      <van-cell title="病情描述" :label="store.consult.illnessDesc"></van-cell>
    </van-cell-group>
    <div class="pay-schema">
      <van-checkbox v-model="agree"> 我已同意 <span class="text">支付协议</span> </van-checkbox>
    </div>
    <van-submit-bar
      button-type="primary"
      :price="payInfo.actualPayment * 100"
      button-text="立即支付"
      text-align="left"
      @click="submit"
      :loading="loading"
    />
    <!-- 支付抽屉，控制面板 close-on-popstate 表示回退不关闭弹窗; closable 表示不可关闭-->
    <van-action-sheet
      v-model:show="show"
      title="选择支付方式"
      :close-on-popstate="false"
      :closeable="false"
      :before-close="onClose"
    >
      <div class="pay-type">
        <p class="amount">￥ {{ payInfo.actualPayment.toFixed(2) }}</p>
        <van-cell-group>
          <van-cell title="微信支付" @click="paymentMethod = 0">
            <template #icon><cp-icon name="consult-wechat" /></template>
            <template #extra><van-checkbox :checked="paymentMethod == 0" /></template>
          </van-cell>
          <van-cell title="支付宝支付" @click="paymentMethod = 1">
            <template #icon><cp-icon name="consult-alipay" /></template>
            <template #extra><van-checkbox :checked="paymentMethod == 1" /></template>
          </van-cell>
        </van-cell-group>
        <div class="btn">
          <van-button type="primary" round block>立即支付</van-button>
        </div>
      </div>
    </van-action-sheet>
  </div>
  <!-- 这下面都是骨架屏 -->
  <div class="consult-pay-page" v-else>
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

// 支付方式
.pay-type {
  .amount {
    padding: 20px;
    text-align: center;
    font-size: 16px;
    font-weight: bold;
  }
  .btn {
    padding: 15px;
  }
  .van-cell {
    align-items: center;
    .cp-icon {
      margin-right: 10px;
      font-size: 18px;
    }
    .van-checkbox :deep(.van-checkbox__icon) {
      font-size: 16px;
    }
  }
}
</style>
