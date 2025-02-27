import CpNavBar from '@/components/CpNavBar.vue'
import CpIcon from '@/components/CpIcon.vue'
import CpRadioBtn from '@/components/CpRadioBtn.vue'

declare module 'vue' {
  interface GlobalComponents {
    // 添加全局组件类型
    CpRadioBtn: typeof CpRadioBtn // 单选按钮
    CpNavBar: typeof CpNavBar // typeof 作用：从js变量中取出对应的类型
    CpIcon: typeof CpIcon // svg图标
  }
}
