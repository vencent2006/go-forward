import type { ConsultType } from '@/enums'
import type { ConsultIllness, PartialConsult } from '@/types/consult'
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

    // 记录病情
    const setIllness = (illness: ConsultIllness) => {
      consult.value.illnessDesc = illness.illnessDesc
      consult.value.illnessTime = illness.illnessTime
      consult.value.consultFlag = illness.consultFlag
      consult.value.pictures = illness.pictures
    }

    // 记录患者id
    const setPatientId = (id: string) => {
      consult.value.patientId = id
    }

    return { consult, setType, setIllnessType, setDep, setIllness, setPatientId }
  },
  {
    persist: true,
  },
)
