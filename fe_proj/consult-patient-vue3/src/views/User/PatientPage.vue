<script setup lang="ts">
import { addPatient, delPatient, editPatient, getPatientList } from '@/services/user'
import type { PatientList, Patient } from '@/types/user'
import { ref, onMounted, computed } from 'vue'
// import { nameRules, idCardRules } from '@/utils/rules'
import { showConfirmDialog, showSuccessToast, showToast, type FormInstance } from 'vant'
import { useRoute } from 'vue-router'
// import { useConsultStore } from '@/stores'
import router from '@/router'
import { nameRules, idCardRules } from '@/utils/rules'
import { useConsultStore } from '@/stores'

// 组件挂载完毕，获取数据
const list = ref<PatientList>([])
const loadList = async () => {
  const res = await getPatientList()
  list.value = res.data
  // 默认选中患者
  if (isChange.value && list.value.length) {
    const defaultPatient = list.value.find((item) => item.defaultFlag === 1)
    if (defaultPatient) {
      patientId.value = defaultPatient.id
    } else {
      // 没有默认患者，默认选中第一个
      patientId.value = list.value[0].id
    }
  }
}

onMounted(() => {
  loadList()
})

const options = [
  { label: '男', value: 1 },
  { label: '女', value: 0 },
]

// 控制popup
const show = ref(false)
const showPopup = (item?: Patient) => {
  if (item) {
    // 有数据，说明是编辑
    const { id, name, idCard, gender, defaultFlag } = item
    patient.value = { id, name, idCard, gender, defaultFlag }
  } else {
    // 没有数据，说明是添加

    // 清空表单校验
    form.value?.resetValidation()
    // 清空数据
    patient.value = { ...initPatient }
  }
  show.value = true
}

const initPatient: Patient = {
  name: '',
  idCard: '',
  gender: 1,
  defaultFlag: 0,
}
const patient = ref<Patient>({ ...initPatient }) // 把initPatient的解构赋给patient

// 支持复选框：因为patient的defaultFlag是number（0|1），随意计算属性支持复选框（true转1，false转0）
const defaultFlag = computed({
  get: () => {
    return patient.value.defaultFlag === 1 ? true : false
  },
  set: (value) => {
    patient.value.defaultFlag = value ? 1 : 0
  },
})

// 进行提交
const form = ref<FormInstance>() // 通过ref绑定实例
const onSubmit = async () => {
  // 表单整体校验 validate 进行校验
  await form.value?.validate()
  // 性别校验
  // 取出身份证倒数第二位，%2之后， 1男，0女
  const gender = +patient.value.idCard.slice(-2, -1) % 2 // +号是把字符串转数字
  if (gender !== patient.value.gender) {
    await showConfirmDialog({
      title: '温馨提示',
      message: '填写的性别与身份证上的不一致\n您确认提交吗？',
    })
    // 确认了，会继续往下走
  }

  // 提交即可 添加 或者 编辑
  patient.value.id ? await editPatient(patient.value) : await addPatient(patient.value)

  // 成功: 关闭添加患者界面，加载患者列表，成功提示
  show.value = false
  await loadList() // 重新加载数据
  showSuccessToast(patient.value.id ? '编辑成功' : '添加成功')
}

// 删除患者
const remove = async () => {
  if (!patient.value.id) {
    return showToast('请选择患者')
  }
  // 确认框，删除请求，关闭，加载，提示
  // 1. 确认框
  await showConfirmDialog({
    title: '温馨提示',
    message: `您确认删除 ${patient.value.name} 患者信息？`,
  })
  // 2. 删除请求
  await delPatient(patient.value.id!)
  show.value = false
  await loadList()
  showSuccessToast('删除成功')
}

// 路由传参 isChange 是不是选择患者
const route = useRoute()
const isChange = computed(() => route.query.isChange === '1')

// 选中效果
const patientId = ref<string>()
const selectedPatient = (item: Patient) => {
  if (isChange.value) {
    // 选择患者
    patientId.value = item.id
  }
}

// 下一步
const store = useConsultStore()
const next = () => {
  if (!patientId.value) {
    return showToast('请选择患者')
  }
  // 保存患者id
  store.setPatientId(patientId.value)
  // 跳转
  router.push('/consult/pay')
}
</script>

<template>
  <div class="patient-page">
    <!-- 导航栏 -->
    <cp-nav-bar :title="isChange ? '选择患者' : '家庭档案'"></cp-nav-bar>

    <!-- 头部提示 -->
    <div class="patient-change" v-if="isChange">
      <h3>请选择患者信息</h3>
      <p>以便医生给出更准确的治疗，信息仅医生可见</p>
    </div>

    <div class="patient-list">
      <div
        class="patient-item"
        v-for="item in list"
        :key="item.id"
        @click="selectedPatient(item)"
        :class="{ selected: patientId == item.id }"
      >
        <div class="info">
          <span class="name">{{ item.name }}</span>
          <span class="id">
            {{ item.idCard.replace(/^(.{6}).+(.{4})$/, '$1********$2') }}
          </span>
          <span>{{ item.genderValue }}</span>
          <span>{{ item.age }}岁</span>
        </div>
        <div class="icon" @click.stop="showPopup(item)">
          <cp-icon name="user-edit" />
        </div>
        <div class="tag" v-if="item.defaultFlag === 1">默认</div>
      </div>
      <div class="patient-add" v-if="list.length < 6" @click="showPopup()">
        <cp-icon name="user-add" />
        <p>添加患者</p>
      </div>
      <div class="patient-tip">最多可添加 6 人</div>
    </div>
    <!-- 底部按钮 -->
    <div class="patient-next" @click="next" v-if="isChange">
      <van-button type="primary" round block>下一步</van-button>
    </div>
    <!-- 使用 popup组件 -->
    <van-popup position="right" v-model:show="show">
      <cp-nav-bar
        :title="patient.id ? '编辑患者' : '添加患者'"
        right-text="保存"
        :back="() => (show = false)"
        @click-right="onSubmit"
      ></cp-nav-bar>
      <van-form autocomplete="off" ref="form">
        <van-field
          v-model="patient.name"
          label="真实姓名"
          placeholder="请输入真实姓名"
          :rules="nameRules"
        />
        <van-field
          v-model="patient.idCard"
          label="身份证号"
          placeholder="请输入身份证号"
          :rules="idCardRules"
        />
        <van-field label="性别" class="pb4">
          <!-- 单选按钮组件 -->
          <template #input>
            <cp-radio-btn v-model="patient.gender" :options="options"></cp-radio-btn>
          </template>
        </van-field>
        <van-field label="默认就诊人">
          <template #input>
            <van-checkbox v-model="defaultFlag" :icon-size="18" round />
          </template>
        </van-field>
      </van-form>
      <!-- 删除按钮 -->
      <van-action-bar v-if="patient.id">
        <van-action-bar-button text="删除" @click="remove" />
      </van-action-bar>
    </van-popup>
  </div>
</template>

<style lang="scss" scoped>
.patient-page {
  padding: 46px 0 80px;
  :deep() {
    .van-popup {
      width: 100%;
      height: 100%;
      padding-top: 46px;
      box-sizing: border-box;
    }
  }
}
.patient-change {
  padding: 15px;
  > h3 {
    font-weight: normal;
    margin-bottom: 5px;
  }
  > p {
    color: var(--cp-text3);
  }
}
.patient-next {
  padding: 15px;
  background-color: #fff;
  position: fixed;
  left: 0;
  bottom: 0;
  width: 100%;
  height: 80px;
  box-sizing: border-box;
}

// 底部操作栏
.van-action-bar {
  padding: 0 10px;
  margin-bottom: 10px;
  .van-button {
    color: var(--cp-price);
    background-color: var(--cp-bg);
  }
}

.patient-list {
  padding: 15px;
}
.patient-item {
  display: flex;
  align-items: center;
  padding: 15px;
  background-color: var(--cp-bg);
  border-radius: 8px;
  margin-bottom: 15px;
  position: relative;
  border: 1px solid var(--cp-bg);
  transition: all 0.3s;
  overflow: hidden;
  .info {
    display: flex;
    flex-wrap: wrap;
    flex: 1;
    span {
      color: var(--cp-tip);
      margin-right: 20px;
      line-height: 30px;
      &.name {
        font-size: 16px;
        color: var(--cp-text1);
        width: 80px;
        margin-right: 0;
      }
      &.id {
        color: var(--cp-text2);
        width: 180px;
      }
    }
  }
  .icon {
    color: var(--cp-tag);
    width: 20px;
    text-align: center;
  }
  .tag {
    position: absolute;
    right: 60px;
    top: 21px;
    width: 30px;
    height: 16px;
    font-size: 10px;
    color: #fff;
    background-color: var(--cp-primary);
    border-radius: 2px;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  &.selected {
    border-color: var(--cp-primary);
    background-color: var(--cp-plain);
    .icon {
      color: var(--cp-primary);
    }
  }
}
.patient-add {
  background-color: var(--cp-bg);
  color: var(--cp-primary);
  text-align: center;
  padding: 15px 0;
  border-radius: 8px;
  .cp-icon {
    font-size: 24px;
  }
}
.patient-tip {
  color: var(--cp-tag);
  padding: 12px 0;
}
.pb4 {
  padding-bottom: 4px;
}
</style>
