import {createRouter, createWebHistory} from "vue-router";
import Home from "../view/home.vue";
import Login from "../view/login.vue";
import Register from "../view/register.vue";
import Reset from "../view/reset.vue";
import Count from "../view/count.vue";
import Welcome from "../view/home/welcome.vue";
import Help from "../view/home/help.vue";
import Filetrans from "../view/home/filetrans.vue";

const routes = [
    {
        path: "/",
        redirect: "/login", // 重定向到login
    }, {
        path: "/home",
        component: Home,
        children: [
            {
                path: "welcome",
                component: Welcome,
            }, {
                path: "help",
                component: Help,
            }, {
                path: "filetrans",
                component: Filetrans,
            },
        ]
    }, {
        path: "/login",
        component: Login
    }, {
        path: "/register",
        component: Register
    }, {
        path: "/reset",
        component: Reset
    }, {
        path: "/count",
        component: Count
    }, {
        path: "/yidong",
        component: () => import("../view/yidong/index.vue")
    }
]
const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router