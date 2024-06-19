<template>
  <div>
    <NavigateBar></NavigateBar>
    <div class="grading-page">
      <h1>判分页面</h1>
      <div class="question-section">
        <h2>题目：</h2>
        <MarkdownRenderer :content="question" />
      </div>
      <div class="answer-section">
        <h2>标准答案：</h2>
        <MarkdownRenderer :content="standardAnswer" />
      </div>
      <div class="student-answer-section">
        <h2>学生答案：</h2>
        <MarkdownRenderer :content="studentAnswer" />
      </div>
      <div class="score-section">
        <div class="score-item">
          <h2>满分：{{ fullScore }}</h2>
        </div>
        <div class="score-item">
          <h2>老师给的分数：
          <input type="number" v-model="teacherScore" class="score-input">
          </h2>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import NavigateBar from "@/components/NavigateBar.vue";
import MarkdownRenderer from "@/components/MarkdownRenderer.vue";
import axios from "axios";
import store from "@/store";

export default {
  components: {NavigateBar, MarkdownRenderer},

  data() {
    return {
      question: "",
      standardAnswer: "",
      studentAnswer: "",
      fullScore: 10,
      teacherScore: 8
    }
  },

  created() {
    this.getStudentAnswer();
  },

  methods: {
    async getStudentAnswer() {
        try{
          const response = await axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/getStudentAnswers', {
            username: store.state.username
          });
          if (response.data.success) {
            this.standardAnswer = response.data.answer;
            this.studentAnswer = response.data.studentAnswer;
            this.fullScore = response.data.score;
            this.question = response.data.question;
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
.grading-page {
  padding: 20px;
  margin: 40px; /* 增加外部容器的边距 */
}

.question-section,
.answer-section,
.student-answer-section,
.score-section {
  margin-bottom: 20px;
}

h1, h2 {
  color: #333;
}

p {
  background-color: #f9f9f9;
  padding: 10px;
  border-radius: 5px;
}

input[type="number"] {
  width: 100px;
  padding: 5px;
  border-radius: 5px;
  border: 1px solid #ccc;
}

.score-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.score-item {
  flex: 1;
  margin-right: 20px;
}

.score-item:last-child {
  margin-right: 0;
}
</style>