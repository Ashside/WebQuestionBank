import { createStore } from 'vuex';

const store = createStore({
    state: {
        username: null,
    },
    mutations: {
        setUsername(state, username) {
            state.username = username;
        },
    },
    actions: {
        login({ commit }, username) {
            // 保存用户名到状态管理
            commit('setUsername', username);
            // 保存用户名到本地存储
            localStorage.setItem('username', username);
        },
        logout({ commit }) {
            // 清除用户名
            commit('setUsername', null);
            localStorage.removeItem('username');
        },
        initializeStore({ commit }) {
            const username = localStorage.getItem('username');
            if (username) {
                commit('setUsername', username);
            }
        },
    },
});

export default store;
