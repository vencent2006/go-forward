<template>
    <a-layout>
        <the-header></the-header>
        <a-layout>
            <the-sider></the-sider>
            <a-layout style="padding: 0 24px 24px">
                <a-breadcrumb style="margin: 16px 0">
                    <a-breadcrumb-item>Home</a-breadcrumb-item>
                    <a-breadcrumb-item>List</a-breadcrumb-item>
                    <a-breadcrumb-item>App</a-breadcrumb-item>
                </a-breadcrumb>
                <a-layout-content
                        :style="{ background: '#fff', padding: '24px', margin: 0, minHeight: '280px' }"
                >
                    {{ resp }}
                    <a-input v-model:value="resp" @click="onChange"/>
                </a-layout-content>
            </a-layout>
        </a-layout>
    </a-layout>
</template>
<script setup>
import TheHeader from "../components/the-header.vue";
import TheSider from "../components/the-sider.vue";
import axios from "axios";
import {ref} from "vue";
import {message} from "ant-design-vue";

const resp = ref()

axios.get("http://localhost:8080/nls/query", {
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
.site-layout-background {
    background: #fff;
}
</style>