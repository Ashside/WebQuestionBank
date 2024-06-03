import { createApp } from 'vue';
import App from './App.vue';
import router from './router'; // 引入路由文件
// 1. 引入你需要的组件
import { Button } from 'vant';
// 2. 引入组件样式
import 'vant/lib/index.css';
import { BackTop } from "vant";
import { ActionBar, ActionBarIcon, ActionBarButton } from 'vant';
import naive from "naive-ui";



createApp(App)
    .use(router)
    .use(Button)
    .use(BackTop)
    .use(ActionBar)
    .use(ActionBarIcon)
    .use(ActionBarButton)
    .use(naive)
    .mount('#app');
