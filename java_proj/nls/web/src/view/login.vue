<template>
    <a-layout style="height: 100vh; ">
        <a-layout-content
                style="margin: auto; width: 400px; height: 100%;display: flex; flex-direction: column;align-items: center;justify-content: center">
            <div class="reset_title">登录</div>
            <a-form
                    :model="loginMember"
                    name="basic"
                    autocomplete="off"
                    @finish="login"
                    @finishFailed="onFinishFailed"
            >
                <a-form-item
                        name="mobile"
                        :rules="[{ required: true, message: '请输入手机号' }]"
                >
                    <a-input v-model:value="loginMember.mobile" placeholder="手机号">
                        <template #prefix>
                            <MobileOutlined style="margin-left: 5px"/>
                        </template>
                    </a-input>
                </a-form-item>

                <a-form-item
                        name="password"
                        :rules="[{ required: true, message: '请输入密码' }]"
                >
                    <a-input-password v-model:value="loginMember.password" placeholder="密码">
                        <template #prefix>
                            <LockOutlined style="margin-left: 5px"/>
                        </template>
                    </a-input-password>
                </a-form-item>

                <a-form-item>
                    <a-button type="primary" block html-type="submit">登&nbsp;录</a-button>
                </a-form-item>
                <div style="display: flex;align-items: center;justify-content: space-between;">
                    <router-link to="/register">
                        <ArrowLeftOutlined/>
                        注册
                    </router-link>
                    <router-link to="/reset">
                        忘记密码
                        <ArrowRightOutlined/>
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
import {hexMd5Key} from "../utils/md5.js";
import {useRouter} from "vue-router";

let router = useRouter();
const loginMember = ref({
    mobile: '',
    password: '',
});
const login = values => {
    console.log('开始登录:', values);

    axios.post('/nls/web/member/login', {
        mobile: loginMember.value.mobile,
        password: hexMd5Key(loginMember.value.password),
    }).then(response => {
        console.log(response);
        let data = response.data;
        if (data.success) {
            message.success("登录成功！");
            router.push('/home');
        } else {
            message.error(data.message);
        }
    });
};
const onFinishFailed = errorInfo => {
    console.log('Failed:', errorInfo);
};
</script>
