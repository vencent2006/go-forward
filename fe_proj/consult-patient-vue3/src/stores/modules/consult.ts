import type { ConsultType } from '@/enums'
import type { PartialConsult } from '@/types/consult'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useConsultStore = defineStore(
  'cp-consult',
  () => {
    // 问诊信息
    const consult = ref(<PartialConsult>{})
    // 设置问诊信息
    const setType = (type: ConsultType) => {
      consult.value.type = type
    }
    // 记录问诊级别 0 普通 1 三甲
    const setIllnessType = (type: 0 | 1) => {
      consult.value.illnessType = type
    }
    // 记录科室
    const setDep = (depId: string) => {
      consult.value.depId = depId
    }

    return { consult, setType, setIllnessType, setDep }
  },
  {
    persist: true,
  },
)
