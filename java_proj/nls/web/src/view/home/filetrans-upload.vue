<template>
    <a-modal v-model:open="open" title="Basic Modal" @ok="handleOk">
        <a-button type="primary" @click="selectFile" size="large">
            <span><UploadOutlined/>选择音频</span>
        </a-button>
        <input type="file" accept=".mp3,.wav,.m4a" style="display: none" ref="fileUploadCom"/>
        <p>Some contents...</p>
        <p>Some contents...</p>
        <p>Some contents...</p>
    </a-modal>
</template>
<script setup>
import {ref} from 'vue';

const open = ref(false);
const showModal = () => {
    open.value = true;
};
const handleOk = e => {
    console.log(e);
    open.value = false;
};

// 选择文件
const fileUploadCom = ref();
const selectFile = () => {
    fileUploadCom.value.click();
}

// 使用 defineExpose 向外暴露指定的数据和方法
defineExpose({
    showModal
})
</script>