import { followOrUnfollow } from '@/services/consult'
import type { FollowType } from '@/types/consult'
import { ref } from 'vue'

// Vue3概念：
// 1. 组合式函数（useXxx）：useXxx，组合式函数可以组合多个状态和逻辑（数据和逻辑），返回一个对象，对象中包含状态和逻辑
// 2. 普通函数（xxx）：普通函数可以单独封装一个状态或逻辑，返回一个值
export const useFollow = (type: FollowType = 'doc') => {
  const loading = ref(false)
  // item只需要id和likeFlag两个属性
  const follow = async (item: { id: string; likeFlag: 0 | 1 }) => {
    loading.value = true

    try {
      await followOrUnfollow(item.id, type)
      item.likeFlag = item.likeFlag === 1 ? 0 : 1
    } finally {
      loading.value = false
    }
  }
  return { loading, follow }
}
