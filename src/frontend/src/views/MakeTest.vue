<template>
  <NavigateBar />
  <div v-if="!aiModalOpen">
    <div>
      <h1>问题列表</h1>
      <div class="button-container">
        <button @click="toggleAIGeneration" class="ai-generate-btn">找题太麻烦？不如试试 AI 组卷</button>
      </div>
      <div class="subject-selector-container">
        <div class="subject-selector">
          <label for="subject">选择科目:</label>
          <select id="subject" v-model="selectedSubject">
            <option value="all">所有科目</option>
            <option v-for="subject in subjects" :key="subject.value" :value="subject.value">{{ subject.label }}</option>
          </select>
        </div>
        <div class="difficulty-selector">
          <label for="difficulty">选择难度:</label>
          <select id="difficulty" v-model="selectedDifficulty">
            <option value="all">所有难度</option>
            <option value="1">简单</option>
            <option value="2">中等</option>
            <option value="3">困难</option>
          </select>
        </div>
      </div>
      <ul>
        <!-- 渲染接收到的问题的描述 -->
        <li v-for="(item, index) in filteredQuestions" :key="index">
          <div class="question-header">
            <input type="checkbox" v-model="item.selected">
            <h3>题目{{ index + 1 }}</h3>
            本题分数：<input type="number" v-model="item.score" class="score-input">
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
            <n-tag v-else-if="item.subject === 'chinese'" style="background-color: #42a5f5">语文</n-tag>

            <n-tag v-if="item.difficulty === 1" type="success">简单</n-tag>
            <n-tag v-else-if="item.difficulty === 2" type="warning">中等</n-tag>
            <n-tag v-else-if="item.difficulty === 3" type="error">困难</n-tag>
            <n-tag v-for="(keywordObj, i_keyword) in item.keywords.slice(0, 3)" :key="i_keyword">
              {{ keywordObj.keyword }}
            </n-tag>
          </div>
        </li>
      </ul>
    </div>
    <div class="button-container">
      <button @click="openModal">提交选中的题目</button>
      <transition name="fade">
        <button v-if="submissionSuccess" @click="goToViewTestDocument">查看试卷</button>
      </transition>
    </div>
    <div v-if="isModalOpen" class="modal">
      <div class="modal-content">
        <span class="close" @click="closeModal">&times;</span>
        <center>
          <h2>请输入试卷信息</h2>
          <input type="text" v-model="testName" placeholder="试卷名称">
          <br><br>
          <button @click="submitSelectedQuestions">确认提交</button>
        </center>
      </div>
    </div>
  </div>
  <!-- AI组卷模态窗口 -->
  <div v-if="aiModalOpen">
    <div>
      <div class="subject-selector-container">
        <div class="subject-selector">
          <label for="subject">选择科目:</label>
          <select id="subject" v-model="selectedSubject">
            <option value="all">所有科目</option>
            <option v-for="subject in subjects" :key="subject.value" :value="subject.value">{{ subject.label }}</option>
          </select>
        </div>
        <div class="difficulty-selector">
          <label for="difficulty">选择难度:</label>
          <select id="difficulty" v-model="selectedDifficulty">
            <option value="all">所有难度</option>
            <option value="1">简单</option>
            <option value="2">中等</option>
            <option value="3">困难</option>
          </select>
        </div>
      </div>
      <center>
        <div class="keywords-input">
          <label for="keywords">输入关键词:</label>
          <input id="keywords" v-model="keyword" placeholder="请输入关键词">
        </div>
        <button @click="generateAIQuestions">AI组卷</button>
        <h2>AI组卷结果</h2>
      </center>
      <ul>
        <!-- 渲染接收到的问题的描述 -->
        <li v-for="(item, index) in aiGeneratedQuestions" :key="index">
          <div class="question-header">
            <input type="checkbox" v-model="item.selected">
            <h3>题目{{ index + 1 }}</h3>
            本题分数：<input type="number" v-model="item.score" class="score-input">
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
            <n-tag v-else-if="item.subject === 'chinese'" style="background-color: #42a5f5">语文</n-tag>

            <n-tag v-if="item.difficulty === 1" type="success">简单</n-tag>
            <n-tag v-else-if="item.difficulty === 2" type="warning">中等</n-tag>
            <n-tag v-else-if="item.difficulty === 3" type="error">困难</n-tag>
            <n-tag v-for="(keywordObj, i_keyword) in item.keywords.slice(0, 3)" :key="i_keyword">
              {{ keywordObj.keyword }}
            </n-tag>
          </div>
        </li>
      </ul>
    </div>
    <div class="button-container">
      <button v-if="aiSubmitButton" @click="openModal">提交选中的题目</button>
      <transition name="fade">
        <button v-if="submissionSuccess" @click="goToViewTestDocument">查看试卷</button>
      </transition>
    </div>
    <div v-if="isModalOpen" class="modal">
      <div class="modal-content">
        <span class="close" @click="closeModal">&times;</span>
        <center>
          <h2>请输入试卷信息</h2>
          <input type="text" v-model="testName" placeholder="试卷名称">
          <br><br>
          <button @click="submitSelectedAIQuestions">确认提交</button>
        </center>
      </div>
    </div>
    <div class="button-container">
      <button @click="closeAIModal">返回</button>
    </div>
    <div class="powered-by">Powered by Machine Learning Model</div>
  </div>
</template>

<script>
import axios from 'axios';
import MarkdownRenderer from "@/components/MarkdownRenderer.vue";
import NavigateBar from "@/components/NavigateBar.vue";
import store from "@/store";
import router from "@/router";

export default {
  name: 'ViewQuestions',
  components: {NavigateBar, MarkdownRenderer},

  data() {
    return {
      questions: [],  // 存储从API获取的问题数据
      submissionSuccess: false,
      pdfURL: '',
      isModalOpen: false,  // 控制模态框是否显示
      testName: '',  // 存储输入的试卷名称
      subjects: [{label: '历史', value: 'history'}, {label: '数学', value: 'math'}, {label: '语文', value: 'chinese'}],
      selectedSubject: 'all',
      selectedDifficulty: 'all',
      aiGeneratedQuestions: [],
      aiModalOpen: false,
      aiSubmitButton: false,
      keyword: '',  // 新增关键词数据绑定
    }
  },

  created() {
    this.fetchQuestions();
  },

  computed: {
    // 使用计算属性来动态过滤问题列表
    filteredQuestions() {
      if (this.selectedSubject === 'all' && this.selectedDifficulty === 'all') {
        return this.questions;
      } else if (this.selectedDifficulty === 'all') {
        return this.questions.filter(item => item.subject === this.selectedSubject);
      } else if (this.selectedSubject === 'all') {
        return this.questions.filter(item => item.difficulty === parseInt(this.selectedDifficulty));
      } else {
        return this.questions.filter(item => item.subject === this.selectedSubject && item.difficulty === parseInt(this.selectedDifficulty));
      }
    }
  },

  methods: {
    openModal() {
      this.isModalOpen = true;
    },
    closeModal() {
      this.isModalOpen = false;
    },
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
      const selectedQuestions = this.questions.filter(q => q.selected);
      if (selectedQuestions.every(q => q.score && q.score > 0)) {
        // 所有选中的题目都有有效分数
        this.closeModal();
        axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/makeTest', {
          username: store.state.username,
          testName: this.testName,
          questions: selectedQuestions.map(q => ({ id: q.id, score: q.score }))
        })
            .then(response => {
              if(response.data.success) {
                this.submissionSuccess = true;
                this.pdfURL = 'https://' + response.data.pdfURL;
              } else {
                console.error("提交失败:", response.data.reason);
                this.errorMessage = "提交失败: " + response.data.reason;  // 显示错误消息
              }
            })
            .catch(error => {
              console.error("提交时出错:", error);
              this.errorMessage = "提交时出错: " + error.message;  // 显示错误消息
            });
      } else {
        // 不是所有选中的题目都有分数
        alert('请为所有选中的题目输入有效分数');
      }
    },

    submitSelectedAIQuestions() {
      const selectedAIQuestions = this.aiGeneratedQuestions.filter(q => q.selected);
      if (selectedAIQuestions.every(q => q.score && q.score > 0)) {
        // 所有选中的题目都有有效分数
        this.closeModal();
        axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/makeTest', {
          username: store.state.username,
          testName: this.testName,
          questions: selectedAIQuestions.map(q => ({ id: q.id, score: q.score }))
        })
            .then(response => {
              if(response.data.success) {
                this.submissionSuccess = true;
                this.pdfURL = 'https://' + response.data.pdfURL;
              } else {
                console.error("提交失败:", response.data.reason);
                this.errorMessage = "提交失败: " + response.data.reason;  // 显示错误消息
              }
            })
            .catch(error => {
              console.error("提交时出错:", error);
              this.errorMessage = "提交时出错: " + error.message;  // 显示错误消息
            });
      } else {
        // 不是所有选中的题目都有分数
        alert('请为所有选中的题目输入有效分数');
      }
    },

    goToViewTestDocument() {
      router.push('/viewAllTests');
    },

    toggleAIGeneration() {
      this.aiModalOpen = !this.aiModalOpen;
    },

    closeAIModal() {
      this.aiModalOpen = false;
    },

    async generateAIQuestions() {
      try {
        const response = await axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/aiGenerate', {
          subject: this.selectedSubject,
          difficulty: this.selectedDifficulty,
          keyword: this.keyword === '' ? null : this.keyword
        });
        if (response.data.success) {
          this.aiGeneratedQuestions = response.data.questions;
          this.aiSubmitButton = true;
          this.aiGeneratedQuestions = response.data.questions.map(q => {
            q.selected = true; // 设置所有生成的问题为选中状态
            return q;
          });
        } else {
          console.error('AI组卷失败:', response.data.reason);
        }
      } catch (error) {
        console.error('调用AI组卷接口出错:', error);
      }
    },
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

.modal {
  position: fixed;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 5px;
  width: 300px;
}

.close {
  float: right;
  cursor: pointer;
}

.powered-by {
  position: fixed; /* 固定位置，不随滚动条滚动 */
  right: 0; /* 左下角 */
  bottom: 10px; /* 距底部10px */
  background-color: rgba(0, 0, 0, 0.5); /* 半透明黑色背景 */
  color: rgba(255, 255, 255, 0.5); /* 文字颜色为白色，半透明 */
  padding: 5px 10px; /* 内边距 */
  font-size: 14px; /* 字体大小增大 */
  border-radius: 10px; /* 移除边框圆角 */
  z-index: 1000; /* 确保在最前面 */
}
</style>
