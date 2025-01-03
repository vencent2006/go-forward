<template>
    <h1>welcome to help</h1>
  {{ resp }}
    <a-input v-model:value="resp" @click="onChange"/>
</template>
<script setup>
import axios from "axios";
import {ref} from "vue";
import {message} from "ant-design-vue";

const resp = ref()

axios.get(import.meta.env.VITE_SERVER + "/nls/query", {
    params: {
        mobile: "1"
    }
}).then(response => {
    // console.log(response);
    let data = response.data
    if (data.success) {
        resp.value = data.content
    } else {
        message.error(data.message)
    }
})

const onChange = () => {
    console.log(resp.value)
}
</script>
<style scoped>

</style>