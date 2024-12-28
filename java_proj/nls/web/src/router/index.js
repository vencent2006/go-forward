import {createRouter, createWebHistory} from "vue-router";
import Home from "../view/home.vue";

const routes = [
    {
        path: "/home",
        component: Home
    }
]
const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router