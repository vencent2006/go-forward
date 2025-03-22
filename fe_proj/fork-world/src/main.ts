import { createApp } from 'vue'
import pinia from './stores'

import App from './App.vue'
import router from './router'

// 在main.scss前面引入vant的样式，后面需要去覆盖vant
import 'vant/lib/index.css'
import './styles/main.scss'

// 导入svg插件需要的代码
import 'virtual:svg-icons-register'

const app = createApp(App)

app.use(pinia)
app.use(router)

app.mount('#app')
