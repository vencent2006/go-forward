<template>
    <a-layout>
        <the-header></the-header>
        <a-layout>
            <the-sider></the-sider>
            <a-layout-content
                    :style="{ background: '#fff', padding: '24px', margin: 0, minHeight: '280px' }"
            >
                <router-view/>
            </a-layout-content>
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