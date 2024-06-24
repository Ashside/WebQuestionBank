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
</style>
