<template>
    <a-modal v-model:open="open" title="Basic Modal" @ok="handleOk">
        <a-button type="primary" @click="selectFile" size="large">
            <span><UploadOutlined/>选择音频</span>
        </a-button>
        <input type="file"
               accept=".mp3,.wav,.m4a"
               style="display: none"
               ref="fileUploadCom"
               @change="uploadFile"
        />
        <p>Some contents...</p>
        <p>Some contents...</p>
        <p>Some contents...</p>
    </a-modal>
</template>
<script setup>
import {ref} from 'vue';
import {notification} from "ant-design-vue";

const open = ref(false);
const showModal = () => {
    open.value = true;
};
const handleOk = e => {
    console.log(e);
    open.value = false;
};

// ----------- 选择文件 -----------
const fileUploadCom = ref();
const selectFile = () => {
    fileUploadCom.value.value = '';// 清空input
    fileUploadCom.value.click();
}

// ----------- 上传文件 -----------
const uploadFile = () => {
    // message.info(fileUploadCom.value.files[0].name)
    const file = fileUploadCom.value.files[0];
    console.log(file)

    // 音频文件最大为500M
    const max = 500 * 1024 * 1024;
    const size = file.size;
    if (size > max) {
        // message.error('文件大小不能超过500M')
        notification['warning']({
            message: '系统提示',
            description: '文件大小超过最大限制，最大为500M',
        })
    }
}

// 使用 defineExpose 向外暴露指定的数据和方法
defineExpose({
    showModal
})
</script>