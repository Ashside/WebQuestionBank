<template>
  <h1>查看作答情况</h1>
  <div>
    <div v-if="loading">加载中...</div>
    <div v-else>
      <div v-for="(question, index) in questions" :key="index">
        <AnsweringSituation
            :title="`问题 ${index + 1}`"
            :question="question.question"
            :correctAnswer="question.answer"
            :questionType="question.type"
            :options="getOptions(question)"
            :studentOptions="getStudentOptions(question)"
            :answer="question.answer"
            :studentAnswer="question.studentAnswer"
            :isReviewComplete="question.isReviewComplete"
            :fullScore="question.fullScore"
            :student-score="question.studentScore"
        />
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import AnsweringSituation from "@/components/AnsweringSituation.vue";
import store from "@/store";

export default {
  components: {
    AnsweringSituation,
  },
  data() {
    return {
      questions: [],
      loading: true,
    };
  },
  methods: {
    getOptions(question) {
      if (question.type === 'multipleChoice' && question.option) {
        return Object.keys(question.option).map(key => ({
          content: question.option[key],
          selected: question.answer.includes(key)
        }));
      }
      return [];
    },
    getStudentOptions(question) {
      if (question.type === 'multipleChoice' && question.option) {
        return Object.keys(question.option).map(key => ({
          content: question.option[key],
          selected: question.studentAnswer.includes(key)
        }));
      }
      return [];
    },
    fetchQuestions() {
      axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/queryTestDetailByStudentID', {
        studentUsername: store.state.username,
        testID: this.$route.query.testID || -1,
      })
          .then(response => {
            if (response.data.success) {
              this.questions = response.data.questions;
            } else {
              console.error(response.data.reason);
            }
          })
          .catch(error => {
            console.error(error);
          })
          .finally(() => {
            this.loading = false;
          });
    }
  },
  created() {
    this.fetchQuestions();
  },
};
</script>

<style scoped>
/* 父组件样式 */
</style>
