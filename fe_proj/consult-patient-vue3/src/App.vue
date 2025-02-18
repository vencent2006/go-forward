<script setup lang="ts">
import { reactive, ref } from 'vue'

type Item = {
  id: string
  name: string
  price: number
}
const list = ref<Item[]>([])
list.value = [
  { id: '1', name: 'item 1', price: 100 },
  { id: '2', name: 'item 2', price: 200 },
  { id: '3', name: 'item 3', price: 300 },
]
const form = reactive({
  username: '',
  password: '',
})
</script>
<template>
  <div>
    <li v-for="item in list" :key="item.id">{{ item.name }}</li>
  </div>
</template>
