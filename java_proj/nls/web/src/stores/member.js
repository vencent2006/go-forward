// 导入一个方法 defineStore

import {defineStore} from "pinia";
import {ref} from "vue";

const MEMBER = 'member'
export const useMemberStore = defineStore('member', () => {
    // 定义数据(state)
    const member = ref({})
    member.value = window.SessionStorage.get(MEMBER) || {}

    const setMember = (data) => {
        member.value = data
        window.SessionStorage.set(MEMBER, data)
    }


    // 以对象的形式return供组件使用
    return {
        member,
        setMember,
    }
})