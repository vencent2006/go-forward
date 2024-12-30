import {createRouter, createWebHistory} from "vue-router";
import Home from "../view/home.vue";
import Login from "../view/login.vue";
import Register from "../view/register.vue";

const routes = [
    {
        path: "/",
        redirect: "/login", // 重定向到login
    }, {
        path: "/home",
        component: Home
    }, {
        path: "/login",
        component: Login
    }, {
        path: "/register",
        component: Register
    },
]
const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router