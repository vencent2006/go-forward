<script setup>
// 1. 导入 use 打头的方法
import {useCounterStore} from "../stores/counter.js";
import {onMounted} from "vue";
// 2. 执行方法得到store实例对象
const counterStore = useCounterStore();
console.log(counterStore);

onMounted(() => {
    counterStore.getList()
})
</script>

<template>
    <button @click="counterStore.increment">{{ counterStore.count }}</button>
  {{ counterStore.doubleCount }}
    <ul>
        <li v-for="item in counterStore.list" :key="item.id">
            {{ item.name }}
        </li>
    </ul>
</template>

<style scoped></style>