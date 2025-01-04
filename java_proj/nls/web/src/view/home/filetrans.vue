<template>
    <div>
        <a-button type="primary" @click="showModal">语音识别</a-button>
        <a-modal v-model:open="open" title="Basic Modal" @ok="handleOk">
            <p>Some contents...</p>
            <p>Some contents...</p>
            <p>Some contents...</p>
        </a-modal>
    </div>
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
</script>