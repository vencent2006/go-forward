<template>
    <a-form
            :model="loginMember"
            name="basic"
            :label-col="{ span: 8 }"
            :wrapper-col="{ span: 16 }"
            autocomplete="off"
            @finish="login"
            @finishFailed="onFinishFailed"
    >
        <a-form-item
                label="mobile"
                name="username"
                :rules="[{ required: true, message: '请输入手机号' }]"
        >
            <a-input v-model:value="loginMember.mobile">
                <template #prefix>
                    <MobileOutlined style="margin-left: 15px"/>
                </template>
            </a-input>
        </a-form-item>

        <a-form-item
                label="Password"
                name="password"
                :rules="[{ required: true, message: 'Please input your password!' }]"
        >
            <a-input-password v-model:value="loginMember.password">
                <template #prefix>
                    <LockOutlined style="margin-left: 15px"/>
                </template>
            </a-input-password>
        </a-form-item>

        <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
            <a-button type="primary" block html-type="submit">登&nbsp;录</a-button>
        </a-form-item>
    </a-form>
</template>
<script setup>
import {ref} from 'vue';

const loginMember = ref({
    mobile: '',
    password: '',
});
const login = values => {
    console.log('开始登录:', values);
};
const onFinishFailed = errorInfo => {
    console.log('Failed:', errorInfo);
};
</script>