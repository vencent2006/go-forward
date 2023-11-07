import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
// 抹平不同浏览器间的差异
import 'normalize.css'
import './style/base.scss'

createApp(App).use(store).use(router).mount('#app')
