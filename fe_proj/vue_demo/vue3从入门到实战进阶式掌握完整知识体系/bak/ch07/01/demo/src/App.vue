<template>
  <div>
    <input v-model="inputValue"/>
    <button class="button" @click="handleAddItem">提交</button>
  </div>
  <ul>
    <list-item
      v-for="(item, index) in list"
      :key="index"
      :msg="item"
    />
  </ul>
</template>

<script>
// 单文件组件
import { reactive, ref } from 'vue';
import ListItem from './components/ListItem.vue';

export default {
  name: 'App',
  components: { ListItem },
  setup() {
    const inputValue = ref('')
    const list = reactive([])
    const handleAddItem = () => {
      list.push(inputValue.value)
      inputValue.value = ''
    }
    return { inputValue, list, handleAddItem }
  }
}
</script>

<style>
  .button {
    margin-left: 20px;
    color: red;
  }
</style>
