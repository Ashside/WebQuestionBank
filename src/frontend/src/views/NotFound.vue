<template>
  <div class="error-container">
    <div class="error-message">
      <h1>404</h1>
      <p>页面不存在</p>
      <CustomButton :to="'/'">回到主页</CustomButton>
    </div>
    <div class="error-image">
      <img src="../../public/404.png" alt="404 Error" />
    </div>
  </div>
</template>

<script>
import CustomButton from "@/components/CustomButton.vue";

export default {
  name: 'NotFound',
  components: {CustomButton},
};
</script>

<style scoped>
.error-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: #f5f5f5; /* 可以根据需要更改背景颜色 */
}

.error-message {
  text-align: center;
  animation: fadeInDown 1s ease-in-out;
}

.error-message h1 {
  font-size: 6em;
  margin: 0;
  color: #ff6f61; /* 可以根据自己的喜好修改颜色 */
}

.error-message p {
  font-size: 1.5em;
  margin-top: 0.5em;
  color: #333;
}

.error-image {
  margin-top: 2em; /* 根据需要调整距离 */
  animation: fadeInUp 1s ease-in-out;
}

.error-image img {
  max-width: 100%; /* 根据需要调整大小 */
  width: 300px;
}

@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translateY(-20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
