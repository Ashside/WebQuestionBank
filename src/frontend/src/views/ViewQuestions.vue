<template>
  <NavigateBar />
  <div>
    <h1>问题列表</h1>
    <ul>
      <!-- 渲染接收到的问题的描述 -->
      <li v-for="(item, index) in questions" :key="index">
        <h3> 题目{{ index + 1 }} </h3>
        <div v-if="item.type === 'simpleAnswer'">
          <MarkdownRenderer :content="item.question" />
        </div>
        <div v-else-if="item.type === 'multipleChoice'">
          <MarkdownRenderer :content="item.question + '<br>' +
          'option1: ' + item.option.option1 + '<br>' +
          'option2: ' + item.option.option2 + '<br>' +
          'option3: ' + item.option.option3 + '<br>' +
          'option4: ' + item.option.option4" />
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';
import MarkdownRenderer from "@/components/MarkdownRenderer.vue";
import NavigateBar from "@/components/NavigateBar.vue";

export default {
  name: 'ViewQuestions',
  components: {NavigateBar, MarkdownRenderer},

  data() {
    return {
      questions: []  // 存储从API获取的问题数据
    }
  },

  created() {
    this.fetchQuestions();
  },

  methods: {
    async fetchQuestions() {
      try {
        // 向API发起请求并获取数据
        const response = await axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/queryQuestion');
        if (response.data.success) {
          this.questions = response.data.questions;  // 从返回数据中获取问题列表
        } else {
          console.error('Failed to fetch questions:', response.data.reason);
          // 处理API返回的错误
        }
      } catch (error) {
        console.error('Error fetching questions:', error);
        // 处理请求错误
      }
    }
  }
}
</script>

<style scoped>
/* 整体背景、字体设置和内边距 */
div {
  background-color: #f4f4f9;
  font-family: 'Arial', sans-serif;
  padding: 13px;  /* 新增内边距 */
}

/* 标题样式 */
h1 {
  color: #333;
  text-align: center;
  margin-bottom: 20px; /* 增加标题与列表间的间隔 */
}

/* 列表样式 */
ul {
  list-style-type: none;
  margin: 0 auto; /* 居中显示 */
  padding: 0;
  max-width: 800px; /* 限制最大宽度 */
}

li {
  background-color: #ffffff;
  border: 1px solid #ddd;
  margin-top: 10px;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
  transition: transform 0.3s ease-in-out;
}

li:hover {
  transform: scale(1.02);
}

/* 问题标题样式 */
h3 {
  color: #2a2a72;
  padding: 10px 15px;
  margin: 0;
  border-bottom: 1px solid #eee;
}
</style>
