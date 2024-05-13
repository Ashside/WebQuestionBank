<template>
  <div class="login-form">
    <transition name="fade">
      <div v-if="!loggedIn" class="login-box">
        <h2>Welcome Back!</h2>
        <form @submit.prevent="login">
          <input type="text" placeholder="Username" v-model="username">
          <input type="password" placeholder="Password" v-model="password">
          <button type="submit">Login</button>
        </form>
      </div>
    </transition>
    <transition name="fade">
      <div v-if="loggedIn" class="success-message">
        <h2>Welcome {{ username }}!</h2>
        <p>Login successful!</p>
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  name: 'LoginForm', // 添加组件名称
  data() {
    return {
      username: '',
      password: '',
      loggedIn: false
    };
  },
  methods: {
    login() {
      // 在这里实现你的登录逻辑
      // 这里假设登录成功后将 loggedIn 设置为 true
      // 你需要替换成你的实际登录逻辑
      this.loggedIn = true;
    }
  }
};
</script>

<style scoped>
.login-form {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.login-box {
  background: #f0f0f0;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

input {
  width: 100%;
  padding: 10px;
  margin: 10px 0;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.success-message {
  background: #d4edda;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}

.fade-enter, .fade-leave-to /* .fade-leave-active in <2.1.8 */ {
  opacity: 0;
}
</style>
