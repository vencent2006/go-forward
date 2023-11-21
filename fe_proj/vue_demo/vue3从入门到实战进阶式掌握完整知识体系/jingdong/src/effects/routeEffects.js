import { useRouter } from 'vue-router'

// 点击回退逻辑
export const useBackRouterEffect = () => {
  const router = useRouter()
  const handleBackClick = () => {
    // todo: 如果const router = useRouter()
    // 放在这里会报错，不知道为啥
    router.back()
  }
  return { handleBackClick }
}
