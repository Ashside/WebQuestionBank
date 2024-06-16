import { createStore } from 'vuex';

const store = createStore({
    state: {
        username: null,
        role: null,  // 添加角色状态
    },
    mutations: {
        setUsername(state, username) {
            state.username = username;
        },
        setRole(state, role) {
            state.role = role;  // mutation来设置角色
        }
    },
    actions: {
        login({ commit }, { username, role }) {
            // 保存用户名和角色到状态管理
            commit('setUsername', username);
            commit('setRole', role);
            // 保存用户名和角色到本地存储
            localStorage.setItem('username', username);
            localStorage.setItem('role', role);
        },
        logout({ commit }) {
            // 清除用户名和角色
            commit('setUsername', null);
            commit('setRole', null);
            localStorage.removeItem('username');
            localStorage.removeItem('role');
        },
        initializeStore({ commit }) {
            const username = localStorage.getItem('username');
            const role = localStorage.getItem('role');
            if (username) {
                commit('setUsername', username);
            }
            if (role) {
                commit('setRole', role);
            }
        },
    },
});

export default store;
