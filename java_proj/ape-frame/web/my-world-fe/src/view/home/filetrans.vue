<template>
    <div>
        <a-button type="primary" @click="showModal">开始上传音频</a-button>
        <filetrans-upload ref="filetransUploadComponent"></filetrans-upload>
    </div>
</template>
<script setup>

import FiletransUpload from "./filetrans-upload.vue";
import {ref} from "vue";

const filetransUploadComponent = ref();
const showModal = () => {
    filetransUploadComponent.value.showModal();
};
</script>