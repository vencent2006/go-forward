<template>
    <a-space direction="vertical" :style="{ width: '100%' }" :size="[0, 48]">
        <a-layout>
            <a-layout-header :style="headerStyle">数据异动</a-layout-header>
            <a-layout-content>
                <EchartsDemo/>
                <HeatMap/>
            </a-layout-content>
            <!--            <a-layout-footer :style="footerStyle">Footer</a-layout-footer>-->
        </a-layout>
    </a-space>
</template>

<script setup>
import EchartsDemo from './components/EchartsDemo.vue'
import HeatMap from "./components/HeatMap.vue";

const headerStyle = {
    textAlign: 'center',
    color: '#fff',
    height: 64,
    paddingInline: 50,
    lineHeight: '64px',
    backgroundColor: '#7dbcea',
};
const contentStyle = {
    textAlign: 'center',
    minHeight: 120,
    lineHeight: '120px',
    color: '#fff',
    backgroundColor: '#108ee9',
};
const siderStyle = {
    textAlign: 'center',
    lineHeight: '120px',
    color: '#fff',
    backgroundColor: '#3ba0e9',
};
const footerStyle = {
    textAlign: 'center',
    color: '#fff',
    backgroundColor: '#7dbcea',
};
</script>