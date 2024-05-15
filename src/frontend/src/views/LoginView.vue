<template>
  <div class="login-container">
    <div class="login-card shadow-lg">
      <h1 class="text-center mb-4">Welcome</h1>
      <form @submit.prevent="handleLogin">
        <div class="form-floating mb-3 text-center">
          <div><label for="email">Email address</label></div>
          <p></p>
          <input
              type="email"
              class="form-control text-center"
              id="email"
              v-model="email"
              required
          />
        </div>
        <p></p>
        <div class="form-floating mb-3 text-center">
          <div><label for="password">Password</label></div>
          <p></p>
          <input
              type="password"
              class="form-control text-center"
              id="password"
              v-model="password"
              required
          />
        </div>
        <p></p>
        <button
            class="btn btn-primary w-100"
            type="submit"
            :disabled="isLoading"
        >
          <span v-if="isLoading" class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>
          <span v-if="!isLoading">Login</span>
          <span v-if="isLoading">Logging in...</span>
        </button>
      </form>
    </div>
    <div class="circle-button-container">
      <van-button round type="success" color="#7232dd" to="about">about US</van-button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import axios from 'axios';
import router from "@/router";

const email = ref('');
const password = ref('');
const isLoading = ref(false);

const handleLogin = async () => {
  isLoading.value = true;
  try {
    const response = await axios.post('http://localhost:3000/api/login', {
      email: email.value,
      password: password.value,
    });
    alert('Logged in successfully!');
    if (response.status === 200) {
      router.push('/about');
    }
  } catch (error) {
    alert('Failed to login.');
  } finally {
    isLoading.value = false;
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}

.login-card {
  background: #fff;
  padding: 3rem; /* 增加padding以增大卡片 */
  border-radius: 10px;
  width: 100%;
  max-width: 500px; /* 增加max-width以增大卡片 */
  animation: fadeIn 1s ease-in-out;
  text-align: center; /* 设置卡片内容居中 */
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
  font-size: 1.5rem; /* 调整这个值以设置所需的字号 */
  color: white; /* 设置文字颜色为白色 */
  padding: 10px 20px; /* 添加一些内边距，使按钮更大一些 */
  border-radius: 5px; /* 添加圆角 */
  cursor: pointer; /* 鼠标悬停时显示指针 */
  transition: background 0.3s ease; /* 平滑过渡效果 */
}

.btn-primary:focus, .btn-primary:hover {
  background: linear-gradient(135deg, #5a67d8 0%, #6b46c1 100%);
}

/* 让表单文本居中 */
.form-floating label {
  width: 100%;
  text-align: center;
  font-size: 1.5rem; /* 调整这个值以设置所需的字号 */
}

.form-control {
  text-align: center;
  font-size: 1.5rem; /* 调整这个值以设置所需的字号 */
}

.circle-button-container {
  position: absolute;
  bottom: 20px;
  right: 20px;
}
</style>
