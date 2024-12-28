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
                    Content
                </a-layout-content>
            </a-layout>
        </a-layout>
    </a-layout>
</template>
<script setup>
import TheHeader from "../components/the-header.vue";
import TheSider from "../components/the-sider.vue";
</script>
<style scoped>
.site-layout-background {
    background: #fff;
}
</style>