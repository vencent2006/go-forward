import {createApp} from 'vue'
import './style.css'
import App from './App.vue'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';
// 将antDesign的图标都引进
import * as Icons from '@ant-design/icons-vue';

const app = createApp(App);
app.use(Antd).mount('#app');

// 全局使用图标
const icons = Icons;
for (const i in icons) {
    app.component(i, icons[i])
}