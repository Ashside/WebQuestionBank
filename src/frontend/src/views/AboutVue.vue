<template style="height: 100%">
  <NavigateBar></NavigateBar>
  <div style="min-height: 100%" class="background">
    <div class="container">
      <h1>About Us</h1>
      <p>Welcome to <span class="highlight">Our Company</span>! We are a team of passionate individuals dedicated to excellence.</p>
      <div class="team-member">
        <h3>John Doe</h3>
        <p>Co-founder & CEO</p>
        <p>John is a visionary leader with a keen eye for innovation.</p>
      </div>
      <div class="team-member">
        <h3>Jane Smith</h3>
        <p>Co-founder & CTO</p>
        <p>Jane is a tech enthusiast who loves turning ideas into reality.</p>
      </div>
      <div class="team-member">
        <h3>David Lee</h3>
        <p>Lead Designer</p>
        <p>David's creativity knows no bounds, always pushing the envelope.</p>
      </div>
      <div class="team-member">
        <h3>Emily Johnson</h3>
        <p>Marketing Manager</p>
        <p>Emily is a storyteller at heart, crafting compelling narratives.</p>
      </div>
  </div>
  </div>
</template>

<script>
import NavigateBar from "@/components/NavigateBar.vue";

  export default {
    name: 'AboutUs',
    components: {NavigateBar}
  };
</script>

<style scoped>
  html, body {
    font-family: Arial, sans-serif;
    margin: 0;
    height: 100%;
    padding: 0;
  }
  .container {
    max-width: 800px;
    margin: 50px auto;
    padding: 20px;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
    /* 设置背景颜色及透明度 */
    background-color: rgba(255, 255, 255, 0.7); /* 0.5表示50%的透明度 */
  }
  h1 {
    color: #333;
    text-align: center;
    margin-bottom: 30px;
  }
  p {
    color: #666;
    line-height: 1.6;
    margin-bottom: 20px;
  }
  .highlight {
    color: #ff5722;
    font-weight: bold;
  }
  .team-member {
    margin-bottom: 30px;
    border-bottom: 1px solid #ddd;
    padding-bottom: 20px;
  }
  .team-member:last-child {
    border-bottom: none;
    margin-bottom: 0;
  }
  .team-member h3 {
    color: #333;
    margin: 0 0 10px;
  }
  .team-member p {
    margin: 0;
    color: #666;
  }
  .background {
    /* 设置背景图片的URL */
    background-image: url("../../public/aboutUs.png");
    /* 设置背景图片的大小和行为 */
    background-size: cover; /* 使背景图片充满整个容器 */
    background-repeat: no-repeat; /* 防止背景图片重复 */
    /* 可选：添加其他样式，如背景颜色、高度等 */
    min-height: 100vh;
    margin: 0;
  }
</style>
