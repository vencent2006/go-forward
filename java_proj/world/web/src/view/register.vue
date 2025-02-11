<template>
    <a-layout class="login-layout">
        <a-layout-content class="login-content">
            <div class="login-container">
                <div class="login-left">
                    <h1 class="app-title">my-world</h1>
                    <p class="app-desc">one entry to worlds</p>
                </div>
                <div class="login-right">
                    <a-form
                            :model="registerMember"
                            name="basic"
                            autocomplete="off"
                            @finish="register"
                            @finishFailed="onFinishFailed"
                            class="login-form"
                    >
                        <a-form-item
                                name="mobile"
                                :rules="[{ required: true, message: 'Please input your phone number' }]"
                        >
                            <a-input v-model:value="registerMember.mobile" placeholder="phone" size="large">
                                <template #prefix>
                                    <MobileOutlined/>
                                </template>
                            </a-input>
                        </a-form-item>

                        <a-form-item
                                name="password"
                                :rules="[{ required: true, message: 'Please input your password' }]"
                        >
                            <a-input-password v-model:value="registerMember.password" placeholder="password"
                                              size="large">
                                <template #prefix>
                                    <LockOutlined/>
                                </template>
                            </a-input-password>
                        </a-form-item>

                        <a-form-item
                                name="confirmPassword"
                                :rules="[
                                { required: true, message: 'Please confirm your password' },
                                { validator: validatePassword }
                            ]"
                        >
                            <a-input-password v-model:value="registerMember.confirmPassword"
                                              placeholder="confirm password" size="large">
                                <template #prefix>
                                    <LockOutlined/>
                                </template>
                            </a-input-password>
                        </a-form-item>

                        <a-form-item>
                            <a-button type="primary" block html-type="submit" size="large" class="login-button">
                                Sign Up
                            </a-button>
                        </a-form-item>

                        <div class="divider">
                            <span class="divider-line"></span>
                        </div>

                        <div class="register-button-wrapper">
                            <router-link to="/login">
                                <a-button size="large" class="register-button">Back to Sign In</a-button>
                            </router-link>
                        </div>
                    </a-form>
                </div>
            </div>
        </a-layout-content>
    </a-layout>
</template>

<script setup>
import {ref} from 'vue';
import {message} from "ant-design-vue";
import axios from "axios";
import {hexMd5Key} from "../utils/md5.js";
import {useRouter} from "vue-router";

const router = useRouter();
const registerMember = ref({
    mobile: '',
    password: '',
    confirmPassword: ''
});

const validatePassword = async (rule, value) => {
    if (value && value !== registerMember.value.password) {
        throw new Error('Two passwords do not match!');
    }
};

const register = values => {
    console.log('注册:', values);

    axios.post('/myworld/web/member/register', {
        mobile: registerMember.value.mobile,
        password: hexMd5Key(registerMember.value.password)
    }).then(response => {
        console.log(response);
        let data = response.data;
        if (data.success) {
            message.success("注册成功！");
            router.push('/login');
        } else {
            message.error(data.message);
        }
    });
};

const onFinishFailed = errorInfo => {
    console.log('Failed:', errorInfo);
};
</script>

<style scoped>
.login-layout {
    min-height: 100vh;
    background-color: #f0f2f5;
}

.login-content {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 20px;
}

.login-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    max-width: 980px;
    width: 100%;
}

.login-left {
    flex: 1;
    padding-right: 32px;
    margin-bottom: 100px;
}

.app-title {
    font-size: 56px;
    font-weight: bold;
    color: #1877f2;
    margin-bottom: 0;
}

.app-desc {
    font-size: 28px;
    line-height: 32px;
    width: 500px;
}

.login-right {
    flex: 0 0 396px;
}

.login-form {
    background: #fff;
    border-radius: 8px;
    padding: 20px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, .1), 0 8px 16px rgba(0, 0, 0, .1);
}

.login-button {
    background-color: #1877f2;
    font-size: 20px;
    font-weight: bold;
    height: 48px;
    margin-top: 16px;
}

.divider {
    margin: 20px 0;
    display: flex;
    align-items: center;
}

.divider-line {
    height: 1px;
    background-color: #dadde1;
    flex: 1;
}

.register-button-wrapper {
    text-align: center;
}

.register-button {
    background-color: #42b72a;
    color: white;
    font-size: 17px;
    font-weight: bold;
    height: 48px;
    border: none;
    width: 60%;
}

.register-button:hover {
    background-color: #36a420;
    color: white;
}

@media (max-width: 900px) {
    .login-container {
        flex-direction: column;
    }

    .login-left {
        text-align: center;
        padding-right: 0;
        margin-bottom: 40px;
    }

    .app-desc {
        width: 100%;
    }
}
</style>