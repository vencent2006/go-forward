<template>
    <a-layout style="height: 100vh; ">

        <a-layout-content
                style="margin: auto; width: 400px; height: 100%;display: flex; align-items: center;justify-content: center">
            <a-form
                    :model="resetMember"
                    name="basic"
                    autocomplete="off"
                    @finish="reset"
                    @finishFailed="onFinishFailed"
            >
                <a-form-item
                        name="mobile"
                        :rules="[{ required: true, message: '输入手机号' }]"
                >
                    <a-input
                            v-model:value="resetMember.mobile"
                            placeholder="手机号">
                        <template #prefix>
                            <mobileOutlined style="margin-left: 5px"/>
                        </template>
                    </a-input>
                </a-form-item>

                <a-form-item
                        name="code"
                        :rules="[{ required: true, message: '输入验证码' }]"
                >
                    <a-input-search
                            v-model:value="resetMember.code"
                            placeholder="短信验证码"
                            :enter-button="sendText"
                            @search="sendRegisterSmsCode"
                    >
                        <template #prefix>
                            <MessageOutlined style="margin-left: 5px"/>
                        </template>
                    </a-input-search>
                </a-form-item>

                <a-form-item
                        name="password"
                        :rules="[{ required: true, message: '输入密码' }]"
                >
                    <a-input-password v-model:value="resetMember.passwordOri" placeholder="密码">
                        <template #prefix>
                            <LockOutlined style="margin-left: 5px"/>
                        </template>
                    </a-input-password>
                </a-form-item>


                <a-form-item
                        name="password"
                        :rules="[{ required: true, message: '再次输入密码' }]"
                >
                    <a-input-password v-model:value="resetMember.passwordConfirm" placeholder="确认密码">
                        <template #prefix>
                            <CheckCircleOutlined style="margin-left: 5px"/>
                        </template>
                    </a-input-password>
                </a-form-item>
                <a-form-item>
                    <a-button type="primary" block html-type="submit">重置密码</a-button>
                </a-form-item>
            </a-form>
        </a-layout-content>
    </a-layout>
</template>

<script setup>
import {ref} from 'vue';

const resetMember = ref({
    mobile: '',
    passwordOri: '',
    passwordConfirm: '',
});

const reset = values => {
    console.log('开始重置:', values);
};

const onFinishFailed = errorInfo => {
    console.log('重置失败:', errorInfo);
};

const sendRegisterSmsCode = () => {
    console.log('发送短信验证码:');
};

const sendText = ref('获取验证码');
</script>
