// 导入一个方法 defineStore

import {defineStore} from "pinia";
import {computed, ref} from "vue";
import axios from "axios";

const API_URL = 'http://geek.itheima.net/v1_0/channels'
export const useCounterStore = defineStore('counter', () => {
    // 定义数据(state)
    const count = ref(0)

    // 定义修改数据的方法(action 同步+异步)
    const increment = () => {
        count.value++
    }

    // 定义获取数据的方法(getter)
    const doubleCount = computed(() => {
        return count.value * 2
    })

    // 定义异步action
    const list = ref([])
    const getList = async () => {
        const res = await axios.get(API_URL);
        list.value = res.data.data.channels
    }

    // 以对象的形式return供组件使用
    return {
        count,
        increment,
        doubleCount,
        list,
        getList
    }
})