<template>
  <NavigateBar></NavigateBar>
  <br>
  <div class="about-us">
    <h1>关于我们</h1>
    <p>我们的团队由三位成员组成，每位成员都拥有独特的技能和职责，致力于为用户提供卓越的软件解决方案。</p>

    <h1>Slogan</h1>
    <p>有用的事情赚取金钱，没用的事情改变世界。</p>

    <h2>我们的目标</h2>
    <p>我们的目标是为用户提供一个方便的在线题库，用户可以在这里查看、录入、组卷、测试题目。
      我们希望借助我们的努力，让需要题库管理的用户零成本、低成本搭建好用高效的系统。</p>

    <h2>我们的项目</h2>
    <a href="https://github.com/Ashside/WebQuestionBank" target="_blank">
      <img src="../assets/github.png" alt="GitHub Project">
    </a>

    <h2>团队成员</h2>
      <div class="member">
        <h3>Ashside</h3>
        <p> 搭建并测试后端框架 </p>
        <p> 编写部分路由分组代码 </p>
        <p> 封装部分数据库接口 </p>
      </div>
    <br>
      <div class="member">
        <h3>Meng Yueqi</h3>
        <p> API 接口设计 </p>
        <p> 前端页面开发及部署 </p>
        <p> 部分服务器部署测试 </p>
      </div>
    <br>
      <div class="member">
        <h3>MaxSuperMax33</h3>
        <p> 编写部分路由分组代码 </p>
        <p> 封装部分数据库接口 </p>
        <p> 识别并优化性能瓶颈，进行必要的代码重构 </p>
      </div>
  </div>
</template>

<script>
import NavigateBar from "@/components/NavigateBar.vue";

export default {
  name: 'GitHubProject',
  components: {NavigateBar}
};
</script>

<style scoped>
.about-us {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  text-align: center;
  background-color: #fff; /* Optional: can change to other colors */
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.15);
}

h1, h2 {
  color: #333;
}

p, li, h3 {
  font-size: 16px;
  line-height: 1.5;
}

a {
  display: inline-block;
  margin: 20px;
}

img {
  width: 100px; /* Adjust as needed */
  height: auto;
}


.member {
  flex: 1;
  padding: 10px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
  border-radius: 5px;
  background-color: #f9f9f9;
}
</style>
