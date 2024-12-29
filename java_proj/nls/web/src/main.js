import {createApp} from 'vue'
import './style.css'
import App from './App.vue'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';
// 将antDesign的图标都引进
import * as Icons from '@ant-design/icons-vue';
import router from "./router";
import axios from "axios";

const app = createApp(App);
app
    .use(Antd) // ant design vue
    .use(router) // vue-router
    .mount('#app');

// 全局使用图标
const icons = Icons;
for (const i in icons) {
    app.component(i, icons[i])
}

/**
 * axios拦截器
 */
axios.interceptors.request.use(function (config) {
    console.log('请求参数: ', config);
    return config;
}, error => {
    return Promise.reject(error);
})

axios.interceptors.response.use(function (response) {
    console.log('返回结果: ', response);
    return response;
}, error => {
    console.log('返回错误: ', response);
    return Promise.reject(error);
})