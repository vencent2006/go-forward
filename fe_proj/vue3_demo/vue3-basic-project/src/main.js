import { createApp } from "vue"

// 导入 ElementPlus 和 样式文件
import ElementPlus from "element-plus"
import "element-plus/dist/index.css"

import App from "./App.vue"
const app = createApp(App)

// 使用 ElementPlus
app.use(ElementPlus)

app.mount("#app")
