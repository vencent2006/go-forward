<template>
    <a-layout style="height: 100vh; ">
        <a-layout-content
                style="margin: auto; width: 400px; height: 100%;display: flex; flex-direction: column;align-items: center;justify-content: center">
            <div class="register_title">注册</div>
            <a-form
                    :model="registerMember"
                    name="basic"
                    autocomplete="off"
                    @finish="register"
                    @finishFailed="onFinishFailed"
            >
                <a-form-item
                        name="mobile"
                        :rules="[{ required: true, message: '输入手机号' }]"
                >
                    <a-input
                            v-model:value="registerMember.mobile"
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
                            v-model:value="registerMember.code"
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
                    <a-input-password v-model:value="registerMember.passwordOri" placeholder="密码">
                        <template #prefix>
                            <LockOutlined style="margin-left: 5px"/>
                        </template>
                    </a-input-password>
                </a-form-item>


                <a-form-item
                        name="password"
                        :rules="[{ required: true, message: '再次输入密码' }]"
                >
                    <a-input-password v-model:value="registerMember.passwordConfirm" placeholder="确认密码">
                        <template #prefix>
                            <CheckCircleOutlined style="margin-left: 5px"/>
                        </template>
                    </a-input-password>
                </a-form-item>
                <a-form-item>
                    <a-button type="primary" block html-type="submit">注&nbsp;册</a-button>
                </a-form-item>
            </a-form>
        </a-layout-content>
    </a-layout>
</template>

<script setup>
import {ref} from 'vue';
import {message} from "ant-design-vue";
import axios from "axios";

const registerMember = ref({
    mobile: '',
    passwordOri: '',
    passwordConfirm: '',
});

const register = values => {
    console.log('开始注册:', values);
};

const onFinishFailed = errorInfo => {
    console.log('注册失败:', errorInfo);
};

// -------------- 短信验证码 --------------
const sendText = ref('获取验证码');
const COUNTDOWN = 5;// 倒计时秒数
let countdown = ref(COUNTDOWN);
const setTime = () => {
    if (countdown.value === 0) {
        sendText.value = '获取验证码';
        countdown.value = COUNTDOWN;
        return;
    } else {
        sendText.value = '重新发送(' + countdown.value + ')';
        countdown.value--;
    }
    setTimeout(function () {
        setTime();
    }, 1000);
}

const sendRegisterSmsCode = () => {
    console.log('发送短信验证码:');
    axios.post('/nls/web/sms-code/send-for-register', {
        mobile: registerMember.value.mobile,
    }).then(response => {
        console.log(response);
        let data = response.data;
        if (data.success) {
            setTime();
            message.error("短信发送成功！");
        } else {
            message.error(data.message);
        }
    });
};

</script>
