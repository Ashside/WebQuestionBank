<template>
  <div class="login-container">
    <!-- 使用v-if和v-else来切换登录和注册表单 -->
    <div v-if="isLogin" class="login-card shadow-lg">
      <!-- 登录表单 -->
      <h1 class="text-center mb-4">欢迎使用网络题库</h1>
      <form @submit.prevent="handleLogin">
        <div class="form-floating mb-3 text-center">
          <div><label for="email">电子邮箱地址</label></div>
          <p></p>
          <input type="email" class="form-control text-center" id="email" v-model="email" required />
        </div>
        <p></p>
        <div class="form-floating mb-3 text-center">
          <div><label for="password">密码</label></div>
          <p></p>
          <input type="password" class="form-control text-center" id="password" v-model="password" required />
        </div>
        <p></p>
        <button class="btn btn-primary w-100" type="submit" :disabled="isLoading">
          <span v-if="isLoading" class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>
          <span v-if="!isLoading">登录</span>
          <span v-if="isLoading">登录中...</span>
        </button>
      </form>
      <!-- 切换到注册按钮 -->
      <p></p>
      <button @click="toggleForm" class="btn btn-primary w-100">还没有账号？注册</button>
    </div>

    <div v-else class="login-card shadow-lg">
      <!-- 注册表单 -->
      <h1 class="text-center mb-4">注册</h1>
      <form @submit.prevent="handleRegister">
        <div class="form-floating mb-3 text-center">
          <div><label for="new-email">电子邮箱地址</label></div>
          <p></p>
          <input type="email" class="form-control text-center" id="new-email" v-model="newEmail" required />
        </div>
        <p></p>
        <div class="form-floating mb-3 text-center">
          <div><label for="new-password">密码</label></div>
          <p></p>
          <input type="password" class="form-control text-center" id="new-password" v-model="newPassword" required />
        </div>
        <p></p>
        <!-- 新增用户角色选择 -->
        <div class="form-floating mb-3 text-center">
          <div><label for="role">身份</label></div>
          <select id="role" class="form-control" v-model="role">
            <option value="user">用户</option>
            <option value="admin">管理员</option>
            <option value="student">学生</option>
          </select>
        </div>
        <p></p>
        <button class="btn btn-primary w-100" type="submit" :disabled="isLoading">
          <!-- 注册按钮内容 -->
            <span v-if="isLoading" class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>
            <span v-if="!isLoading">注册</span>
            <span v-if="isLoading">注册中</span>
        </button>
      </form>
      <p></p>
      <button @click="toggleForm" class="btn btn-primary w-100">已经有账号了？登录</button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import axios from 'axios';
import router from "@/router";
import store from "@/store";

const email = ref('');
const password = ref('');
const newEmail = ref('');
const newPassword = ref('');
const isLoading = ref(false);
const isLogin = ref(true);  // 初始为登录视图
const role = ref('user');  // 默认为用户角色

const handleLogin = async () => {
  isLoading.value = true;
  try {
    const response = await axios.post(process.env['VUE_APP_API_URL'] + '/api/usr/loginCheck', {
      username: email.value,
      password: password.value},
    {
      headers: {
        'Content-Type': 'application/json',
      }
    });
    if (response.status === 200 && response.data.success) {
      store.dispatch('login', { username: email.value, role: response.data.type });
      alert('Logged in successfully!');
      await router.push('/home');
    }
  } catch (error) {
    alert('Failed to login.');
  } finally {
    isLoading.value = false;
  }
};

const handleRegister = async () => {
  isLoading.value = true;
  try {
    const response = await axios.post(process.env['VUE_APP_API_URL'] + '/api/usr/registerCheck', {
          username: newEmail.value,
          password: newPassword.value,
          type: role.value,
        },
        {
          headers: {
            'Content-Type': 'application/json',
          }
        });
    if (response.status === 200 && response.data.success) {
      store.dispatch('login', { username: newEmail.value, role: role.value });
      alert('Register in successfully!');
      await router.push('/home');
    }
  } catch (error) {
    alert('Failed to register.');
  } finally {
    isLoading.value = false;
  }
};

const toggleForm = () => {
  isLogin.value = !isLogin.value;  // 切换表单
};
</script>

<style scoped>
.login-container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  position: relative;
}

.login-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 50%;
  width: 200%;
  height: 100%;
  background-size: contain;
  z-index: -1;
  transform: translateX(-50%) skewX(-20deg);
}

.login-card {
  background: #fff;
  padding: 3rem;
  border-radius: 10px;
  width: 100%;
  max-width: 500px;
  animation: fadeIn 1s ease-in-out;
  text-align: center;
}

@keyframes fadeIn {
  0% {
    opacity: 0;
    transform: scale(0.95);
  }
  100% {
    opacity: 1;
    transform: scale(1);
  }
}

.btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  font-size: 1.5rem;
  color: white;
  padding: 10px 20px;
  border-radius: 5px;
  cursor: pointer;
  transition: background 0.3s ease;
}

.btn-primary:focus, .btn-primary:hover {
  background: linear-gradient(135deg, #5a67d8 0%, #6b46c1 100%);
}

.form-floating label {
  width: 100%;
  text-align: center;
  font-size: 1.5rem;
}

.form-control, select {
  text-align: center;
  font-size: 1.5rem;
  border: 2px solid #ccc;
  padding: 10px;
  border-radius: 5px;
  cursor: pointer;
  appearance: none; /* Remove default styling */
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) no-repeat right .75rem center/8px 10px;
  color: black;
}

select option {
  background: white;
  color: black;
}

/* Add custom arrow using CSS */
select::after {
  content: "";
  position: absolute;
  right: 15px;
  top: 50%;
  border-left: 6px solid transparent;
  border-right: 6px solid transparent;
  border-top: 6px solid white; /* Arrow color */
  transform: translateY(-50%);
}
</style>
