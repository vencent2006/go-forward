<template>
    <a-layout-header class="header">
        <div class="logo"/>
        <div style="float: right; color: white;">
            您好: {{ member.name }} &nbsp;&nbsp;
            <router-link to="/login" style="color: white;">
                退出登录
            </router-link>
        </div>
        <a-menu
                v-model:selectedKeys="selectedKeys"
                theme="dark"
                mode="horizontal"
                :style="{ lineHeight: '64px' }"
        >
            <a-menu-item key="/home/welcome">
                <router-link to="/home/welcome">
                    <CoffeeOutlined/>
                    <span>欢迎使用</span>
                </router-link>
            </a-menu-item>
            <a-menu-item key="/home/filetrans">
                <router-link to="/home/filetrans">
                    <VideoCameraOutlined/>
                    <span>语音识别</span>
                </router-link>
            </a-menu-item>
            <a-menu-item key="/home/help">
                <router-link to="/home/help">
                    <QuestionCircleOutlined/>
                    <span>帮助文档</span>
                </router-link>
            </a-menu-item>
        </a-menu>
    </a-layout-header>
</template>
<script setup>
import {ref, watch} from 'vue';
import {useMemberStore} from "../stores/member.js";
import {useRouter} from "vue-router";

const selectedKeys = ref(['/home/welcome']);
const router = useRouter();
watch(() => router.currentRoute.value.path, (newValue, oldValue) => {
    // console.log('watch', newValue, oldValue)
    selectedKeys.value = [newValue]
}, {immediate: true})
const memberStore = useMemberStore()
const member = memberStore.member
</script>

<style scoped>
.logo {
    float: left;
    width: 120px;
    height: 31px;
    margin: 16px 24px 16px 0;
    background: rgba(255, 255, 255, 0.3);
}

.ant-row-rtl .logo {
    float: right;
    margin: 16px 0 16px 24px;
}
</style>