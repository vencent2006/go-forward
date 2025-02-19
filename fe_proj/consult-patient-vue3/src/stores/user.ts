import type { User } from '@/types/user'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore(
  'cp-user',
  () => {
    // 1. 用户信息状态
    const user = ref<User>()
    // 2. 设置用户信息的函数
    const setUser = (u: User) => {
      user.value = u
    }
    // 3. 删除用户信息的函数
    const delUser = () => {
      user.value = undefined
    }
    return { user, setUser, delUser }
  },
  {
    persist: true, // 开启数据缓存
  },
)
