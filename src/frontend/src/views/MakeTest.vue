<template>
  <NavigateBar />
  <div>
    <h1>问题列表</h1>
    <ul>
      <!-- 渲染接收到的问题的描述 -->
      <li v-for="(item, index) in questions" :key="index">
        <div class="question-header">
          <input type="checkbox" v-model="item.selected">
          <h3>题目{{ index + 1 }}</h3>
        </div>
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
        <div class="tag-container">
          <n-tag v-if="item.subject === 'history'" style="background-color: #ffa726">历史</n-tag>
          <n-tag v-else-if="item.subject === 'math'" style="background-color: #66bb6a">数学</n-tag>
          <n-tag v-else-if="item.subject === 'english'" style="background-color: #42a5f5">英语</n-tag>

          <n-tag v-if="item.difficulty === 1" type="success">简单</n-tag>
          <n-tag v-else-if="item.difficulty === 2" type="warning">中等</n-tag>
          <n-tag v-else-if="item.difficulty === 3" type="error">困难</n-tag>
          <n-tag v-for="(keywordObj, i_keyword) in item.keywords" :key="i_keyword">
            {{ keywordObj.keyword }}
          </n-tag>
        </div>
      </li>
    </ul>
  </div>
  <div class="button-container">
    <button @click="submitSelectedQuestions">提交选中的题目</button>
    <transition name="fade">
      <button v-if="submissionSuccess" @click="handleExtraAction">执行额外的操作</button>
    </transition>
  </div>
</template>


<script>
import axios from 'axios';
import MarkdownRenderer from "@/components/MarkdownRenderer.vue";
import NavigateBar from "@/components/NavigateBar.vue";
// import router from "@/router";

export default {
  name: 'ViewQuestions',
  components: {NavigateBar, MarkdownRenderer},

  data() {
    return {
      questions: [],  // 存储从API获取的问题数据
      submissionSuccess: false,
      pdfURL: ''
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
          this.pdfURL = ''
        } else {
          console.error('Failed to fetch questions:', response.data.reason);
          // 处理API返回的错误
        }
      } catch (error) {
        console.error('Error fetching questions:', error);
        // 处理请求错误
      }
    },
    submitSelectedQuestions() {
      const selectedQuestions = this.questions.filter(q => q.selected).map(q => ({ id: q.id }));
      axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/makeTest', { questions: selectedQuestions })
          .then(response => {
            if(response.data.success) {
              this.submissionSuccess = true;
              this.pdfURL = 'https://' + response.data.pdfURL;
            } else {
              console.error("提交失败:", response.data.reason);
            }
          })
          .catch(error => {
            console.error("提交时出错:", error);
          });
    },
    handleExtraAction() {
      window.location.href = this.pdfURL;
    }
  }
}
</script>

<style scoped>
/* 添加标签容器样式 */
.tag-container {
  display: flex;
  gap: 8px;  /* 设置较小的间隔 */
  padding: 8px;  /* 设置内边距 */
}

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

/* 自定义复选框样式 */
input[type="checkbox"] {
  -webkit-appearance: none; /* 移除默认外观 */
  appearance: none;
  background-color: #fff;
  margin: 0 10px 0 0;
  font-size: 1.5em;
  color: #42a5f5;
  width: 20px;
  height: 20px;
  border: 2px solid #42a5f5;
  border-radius: 4px;
  cursor: pointer;
  position: relative;
}

input[type="checkbox"]:checked {
  background-color: #42a5f5;
}

input[type="checkbox"]:checked::after {
  content: "✔";
  position: absolute;
  top: -2px;
  left: 2px;
  color: #fff;
  font-size: 16px;
}

/* 提交按钮样式 */
button {
  background-color: #1e88e5; /* 绿色背景 */
  color: white;
  border: none;
  padding: 10px 20px;
  font-size: 16px;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
  display: block;
  text-align: center;
}

button:hover {
  background-color: #2a2a72; /* 深绿色背景 */
}

.question-header {
  display: flex;
  align-items: center; /* 垂直居中对齐复选框和标题 */
  padding: 10px 15px; /* 提供一些内部空间 */
}

h3 {
  margin: 0 0 0 10px; /* 为标题添加左侧间距 */
  color: #2a2a72;
  flex-grow: 1; /* 允许标题占用剩余空间 */
}

input[type="checkbox"] {
  flex-shrink: 0; /* 防止复选框大小调整 */
}

.button-container {
  display: flex;  /* 启用flex布局 */
  justify-content: center;  /* 水平居中 */
  gap: 10px;  /* 按钮之间的间隔 */
}

</style>
