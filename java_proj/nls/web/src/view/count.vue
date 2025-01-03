<script setup>
// 1. 导入 use 打头的方法
import {useCounterStore} from "../stores/counter.js";
// 2. 执行方法得到store实例对象
const counterStore = useCounterStore();
console.log(counterStore);
</script>

<template>
    <div>
        <button @click="counterStore.increment">{{ counterStore.count }}</button>
        {{ counterStore.doubleCount }}
    </div>
</template>

<style scoped></style>