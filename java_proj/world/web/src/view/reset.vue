<template>
    <a-layout style="height: 100vh; ">
        <a-layout-content
                style="margin: auto; width: 400px; height: 100%;display: flex; flex-direction: column;align-items: center;justify-content: center">
            <div class="reset_title">重置密码</div>
            <a-form
                    :model="resetMember"
                    name="basic"
                    autocomplete="off"
                    @finish="reset"
                    @finishFailed="onFinishFailed"
            >
                <a-form-item
                        name="mobile"
                        :rules="[
                            { required: true, message: '输入手机号', trigger: 'blur' },
                            {pattern: /^1[3456789]\d{9}$/, message: '手机号为11位', trigger: 'blur'}]"
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
                        :rules="[{ required: true, message: '输入验证码', trigger: 'blur' }]"
                >
                    <a-input-search
                            v-model:value="resetMember.code"
                            placeholder="短信验证码"
                            :enter-button="sendText"
                            @search="sendRegisterSmsCode"
                            :loading="sendBtnLoading"
                    >
                        <template #prefix>
                            <MessageOutlined style="margin-left: 5px"/>
                        </template>
                    </a-input-search>
                </a-form-item>

                <a-form-item
                        name="passwordOri"
                        :rules="[{ required: true, message: '输入密码', trigger: 'blur' },{pattern: /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,20}$/, message: '密码必须包含字母和数字，且长度6-20', trigger: 'blur'}]"
                >
                    <a-input-password v-model:value="resetMember.passwordOri" placeholder="密码">
                        <template #prefix>
                            <LockOutlined style="margin-left: 5px"/>
                        </template>
                    </a-input-password>
                </a-form-item>


                <a-form-item
                        name="passwordConfirm"
                        :rules="[{ required: true, message: '再次输入密码' }]"
                >
                    <a-input-password v-model:value="resetMember.passwordConfirm" placeholder="确认密码">
                        <template #prefix>
                            <CheckCircleOutlined style="margin-left: 5px"/>
                        </template>
                    </a-input-password>
                </a-form-item>
                <a-form-item>
                    <a-button type="primary" block html-type="submit">重&nbsp;置</a-button>
                </a-form-item>
                <div style="display: flex;align-items: center;justify-content: left;">
                    <router-link to="/login">
                        <ArrowLeftOutlined/>
                        返回登录
                    </router-link>
                </div>

            </a-form>


        </a-layout-content>
    </a-layout>
</template>

<script setup>
import {ref} from 'vue';
import {message} from "ant-design-vue";
import axios from "axios";
import {useRouter} from "vue-router";
import {hexMd5Key} from "../utils/md5";

let router = useRouter();

const resetMember = ref({
    mobile: '',
    code: '',
    password: '',
    passwordOri: '',
    passwordConfirm: '',
});

const reset = values => {
    console.log('开始重置:', values);
    if (resetMember.value.passwordOri !== resetMember.value.passwordConfirm) {
        message.error("两次密码不一致！");
        return;
    }

    resetMember.value.password = resetMember.value.passwordOri;
    axios.post('/myworld/web/member/reset', {
        mobile: resetMember.value.mobile,
        password: hexMd5Key(resetMember.value.password),
        code: resetMember.value.code,
    }).then(response => {
        console.log(response);
        let data = response.data;
        if (data.success) {
            message.success("重置密码成功！");
            router.push('/login');
        } else {
            message.error(data.message);
        }
    });
};

const onFinishFailed = errorInfo => {
    console.log('注册失败:', errorInfo);
};

// -------------- 短信验证码 --------------
const sendBtnLoading = ref(false);
const sendText = ref('获取验证码');
const COUNTDOWN = 5;// 倒计时秒数
let countdown = ref(COUNTDOWN);
const setTime = () => {
    if (countdown.value === 0) {
        sendBtnLoading.value = false;
        sendText.value = '获取验证码';
        countdown.value = COUNTDOWN;
        return;
    } else {
        sendBtnLoading.value = true;
        sendText.value = '重新发送(' + countdown.value + ')';
        countdown.value--;
    }
    setTimeout(function () {
        setTime();
    }, 1000);
}

const sendRegisterSmsCode = () => {
    console.log('发送短信验证码:');
    sendBtnLoading.value = true;
    axios.post('/myworld/web/sms-code/send-for-reset', {
        mobile: resetMember.value.mobile,
    }).then(response => {
        console.log(response);
        let data = response.data;
        if (data.success) {
            setTime();
            message.success("短信发送成功！");
        } else {
            sendBtnLoading.value = false;
            message.error(data.message);
        }
    });
};


</script>
