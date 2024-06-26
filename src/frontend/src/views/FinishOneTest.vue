<template>
  <h1>答题界面</h1>
  <h1>片纸鉴心，诚信不败</h1>
  <div>
    <ShortAnswerBlock
        v-for="(question, index) in questions"
        :key="index"
        :title="'题目 ' + (index + 1)"
        :question="question.question"
        v-model="question.studentAnswer"
        :questionType="question.type === 'simpleAnswer' ? 'simpleAnswer' : 'multipleChoice'"
        :options="question.options"
    />
  </div>
  <div class="button-container">
    <button @click="saveAnswers">保存答案</button>
    <button @click="submitAnswers">提交答案</button>
  </div>
</template>


<script>
import ShortAnswerBlock from "@/components/AnswerBlock.vue";
import axios from 'axios';
import store from "@/store";

export default {
  components: { ShortAnswerBlock },
  data() {
    return {
      questions: [],
    };
  },
  methods: {
    async fetchQuestions() {
      try {
        const response = await axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/queryTestStateByStudentID', {
            studentUsername: store.state.username, // 将占位符替换为实际的学生ID
            testID: this.$route.query.testID || -1,
        });

        if (response.data.success) {
          this.questions = response.data.questions.map(q => ({
            ...q,
            options: q.type === 'multipleChoice' ? Object.entries(q.option).map(([key, value]) => ({
              content: value,
              selected: q.studentAnswer.includes(key),
            })) : [],
          }));
        }
      } catch (error) {
        console.error('Error fetching questions:', error);
      }
    },

    async saveAnswers() {
      const formattedAnswers = this.questions.map(q => {
        if (q.type === 'multipleChoice') {
          return {
            id: parseInt(q.id, 10),  // 确保ID为整数类型
            type: q.type,
            studentAnswer: q.studentAnswer,
          };
        } else {
          return {
            id: parseInt(q.id, 10),  // 确保ID为整数类型
            type: q.type,
            studentAnswer: q.studentAnswer,
          };
        }
      });

      try {
        const response = await axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/saveTestAnswerByStudentID', {testID: parseInt(this.$route.query.testID, 10), studentUsername: store.state.username, questions: formattedAnswers});
        if (response.data.success) {
          console.log('Answers submitted successfully');
        } else {
          console.error('Failed to submit answers:', response.data.reason);
        }
      } catch (error) {
        console.error('Error submitting answers:', error);
      }
    },

    submitAnswers() {
      // 提交答案的逻辑
      console.log('Submitting answers:', this.questions);
      // 可以通过API提交答案
      // axios.post('/api/submitAnswers', { answers: this.questions })
      //   .then(response => { ... })
      //   .catch(error => { ... });
    },
  },
  mounted() {
    this.fetchQuestions();
  },
};
</script>


<style scoped>
#app {
  margin: 20px;
}

button {
  background-color: #1e88e5;
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
  background-color: #2a2a72;
}
</style>
