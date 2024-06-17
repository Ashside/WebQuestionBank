<template>
  <NavigateBar />
  <div>
    <h1>问题列表</h1>
    <ul>
      <!-- 渲染接收到的问题的描述 -->
      <li v-for="(item, index) in questions" :key="index">
        <div class="question-header">
          <input v-if="isAdmin" type="checkbox" v-model="item.selected">
          <h3> 题目{{ index + 1 }} </h3>
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
          <n-tag v-for="(keywordObj, i_keyword) in item.keywords.slice(0, 3)" :key="i_keyword">
            {{ keywordObj.keyword }}
          </n-tag>
        </div>
      </li>
    </ul>
    <div class="button-container">
      <button v-if="isAdmin" @click="submitDeleteQuestions">删除选中的题目</button>
    </div>
  </div>
</template>


<script>
import axios from 'axios';
import MarkdownRenderer from "@/components/MarkdownRenderer.vue";
import NavigateBar from "@/components/NavigateBar.vue";
import {computed} from "vue";
import store from "@/store";

const storeRole = computed(() => store.state.role);
const isAdmin = computed(() => storeRole.value === 'admin');

export default {
  name: 'ViewQuestions',
  components: {NavigateBar, MarkdownRenderer},

  data() {
    return {
      questions: [],  // 存储从API获取的问题数据
      isAdmin
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
    },
    submitDeleteQuestions() {
      // 获取选中的问题
      const selectedQuestions = this.questions.filter(item => item.selected);
      // 获取选中问题的ID
      const selectedQuestionIds = selectedQuestions.map(item => ({ id: item.id }));
      // 向API发送请求删除选中的问题
      axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/deleteQuestion', {
        username: store.state.username,
        questions: selectedQuestionIds
      }).then(response => {
        if (response.data.success) {
          alert('Deleted successfully!');
          this.fetchQuestions();  // 重新获取问题列表
        } else {
          alert('Failed to delete questions.');
        }
      }).catch(error => {
        console.error('Error deleting questions:', error);
      });
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

/* 自定义复选框样式 */
input[type="checkbox"] {
  -webkit-appearance: none; /* 移除默认外观 */
  appearance: none;
  background-color: #fff;
  margin: 0 10px 0 0;
  font-size: 1.5em;
  color: #f44336; /* 红色 */
  width: 20px;
  height: 20px;
  border: 2px solid #f44336; /* 红色边框 */
  border-radius: 4px;
  cursor: pointer;
  position: relative;
}

input[type="checkbox"]:checked {
  background-color: #f44336; /* 红色背景 */
}

input[type="checkbox"]:checked::after {
  content: "✖"; /* 改为叉 */
  position: absolute;
  top: -4px; /* 根据叉的大小微调位置 */
  left: 4px;
  color: #fff;
  font-size: 16px;
}

/* 提交按钮样式 */
button {
  background-color: #f44336; /* 红色背景 */
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
  background-color: #b71c1c; /* 深红色背景 */
}


input[type="checkbox"] {
  flex-shrink: 0; /* 防止复选框大小调整 */
}

</style>
