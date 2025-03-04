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
    return { consult, setType }
  },
  {
    persist: true,
  },
)
