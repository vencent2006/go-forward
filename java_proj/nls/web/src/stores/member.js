// 导入一个方法 defineStore

import {defineStore} from "pinia";
import {ref} from "vue";
import axios from "axios";

const API_URL = 'http://geek.itheima.net/v1_0/channels'
export const useMemberStore = defineStore('member', () => {
    // 定义数据(state)
    const user = ref({})

    const setMember = (data) => {
        user.value = data
        console.log(user)
    }

    // 定义异步action
    const list = ref([])
    const getList = async () => {
        const res = await axios.get(API_URL);
        list.value = res.data.data.channels
    }

    // 以对象的形式return供组件使用
    return {
        user,
        setMember,
        list,
        getList
    }
})