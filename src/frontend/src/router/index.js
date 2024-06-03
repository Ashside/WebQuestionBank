import {createRouter, createWebHistory} from 'vue-router'
// import HelloWorld from "@/components/HelloWorld.vue";

// Vue.use(VueRouter); // 确保使用 Vue.use() 注册 Vue Router 插件

const routes = [
    {
        path: '/',
        component: () => import('@/views/LoginView.vue')
    },
    {
        path: '/about',
        component: () => import('@/views/AboutVue.vue')
    },
    {
        path: '/loadQuestion',
        component: () => import('@/views/LoadShortAnswer.vue'),
        children: [
            {
                path: 'shortAnswer',
                component: () => import('@/views/LoadShortAnswer.vue')
            }
            // 可以根据需要添加更多子路由
        ]
    },
    {
        path: '/home',
        component: () => import('@/views/HomeVue.vue')
    },
    {
        path: '/:pathMatch(.*)*',
        component: () => import('@/views/NotFound.vue')
    }
];

const router = createRouter({
    // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
    history: createWebHistory(process.env.BASE_URL),
    // mode: 'history',
    routes // `routes: routes` 的缩写
})

export default router;
