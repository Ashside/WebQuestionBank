<template>
  <div class="unauth-container">
    <h1>您尚未登录</h1>
    <button @click="redirectToLogin">登录</button>
  </div>
</template>

<script setup>
import router from "@/router";

function redirectToLogin() {
  router.push('/')  // 确保指定正确的路由名
}
</script>

<style scoped>
.unauth-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background-color: #f0f2f5;
}

h1 {
  color: #333;
  margin-bottom: 20px;
}

button {
  padding: 10px 20px;
  font-size: 16px;
  color: #fff;
  background-color: #007BFF;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
}

button:hover {
  background-color: #0056b3;
}
</style>
