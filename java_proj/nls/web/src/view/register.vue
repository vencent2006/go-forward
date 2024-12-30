<template>
    <a-layout style="height: 100vh; ">
        <a-layout-content style="margin: auto; width: 400px; height: 100%;display: flex; align-items: center;">
            <a-form
                    :model="registerMember"
                    name="basic"
                    :label-col="{ span: 8 }"
                    :wrapper-col="{ span: 16 }"
                    autocomplete="off"
                    @finish="register"
                    @finishFailed="onFinishFailed"
            >
                <a-form-item
                        label="Username"
                        name="username"
                        :rules="[{ required: true, message: 'Please input your username!' }]"
                >
                    <a-input v-model:value="registerMember.username"/>
                </a-form-item>

                <a-form-item
                        label="Email"
                        name="email"
                        :rules="[{ required: true, message: 'Please input your email!' }]"
                >
                    <a-input v-model:value="registerMember.email"/>
                </a-form-item>

                <a-form-item
                        label="Password"
                        name="password"
                        :rules="[{ required: true, message: 'Please input your password!' }]"
                >
                    <a-input-password v-model:value="registerMember.password"/>
                </a-form-item>

                <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
                    <a-button type="primary" block html-type="submit">Register</a-button>
                </a-form-item>
            </a-form>
        </a-layout-content>
    </a-layout>
</template>

<script setup>
import {ref} from 'vue';

const registerMember = ref({
    username: '',
    email: '',
    password: '',
});

const register = values => {
    console.log('开始注册:', values);
};

const onFinishFailed = errorInfo => {
    console.log('注册失败:', errorInfo);
};
</script>
