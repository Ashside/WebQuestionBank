import {createRouter, createWebHistory} from 'vue-router'
// import HelloWorld from "@/components/HelloWorld.vue";

// Vue.use(VueRouter); // 确保使用 Vue.use() 注册 Vue Router 插件
import store from "@/store";

const routes = [
    {
        path: '/',
        component: () => import('@/views/LoginView.vue')
    },
    {
        path: '/Unauthenticated',
        component: () => import('@/views/UnauthenticatedPage.vue')
    },
    {
        path: '/about',
        component: () => import('@/views/AboutVue.vue')
    },
    {
        path: '/loadQuestion',
        children: [
            {
                path: '/shortAnswer',
                component: () => import('@/views/LoadShortAnswer.vue'),
                meta: {
                    requiresAuth: true  // 需要认证
                }
            },
            {
                path: '/multipleChoice',
                component: () => import('@/views/LoadMultipleChoice.vue'),
                meta: {
                    requiresAuth: true  // 需要认证
                }
            }
            // 可以根据需要添加更多子路由
        ]
    },
    {
        path: '/home',
        component: () => import('@/views/HomeVue.vue'),
        meta: {
            requiresAuth: true  // 需要认证
        }
    },
    {
        path: '/viewQuestion',
        component: () => import('@/views/ViewQuestions.vue'),
        meta: {
            requiresAuth: true  // 需要认证
        }
    },
    {
        path: '/:pathMatch(.*)*',
        component: () => import('@/views/NotFound.vue'),
    }
];

const router = createRouter({
    // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
    history: createWebHistory(process.env.BASE_URL),
    // mode: 'history',
    routes // `routes: routes` 的缩写
})

// 添加全局前置守卫
router.beforeEach((to, from, next) => {
    console.log(store.state.username)
    // 检查路由元信息
    if (to.matched.some(record => record.meta.requiresAuth)) {
        // 检查Vuex存储中的用户名
        if (store.state.username !== null) {
            next()  // 如果用户名存在，正常导航
        } else {
            console.log("Here")
            next({ path: 'Unauthenticated' })  // 如果用户名不存在，重定向到登录页
            // next({ name: '' })  // 如果用户名不存在，重定向到登录页
        }
    } else {
        next()  // 确保一定要调用 next()
    }
})

export default router;
