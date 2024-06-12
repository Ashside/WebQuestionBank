<template>
  <NavigateBar/>
  <div>
    <div>
      <center>
        <h1>Question Difficulty</h1>
        <DifficultySelector v-model:difficulty="difficulty" />
        <SubjectSelector v-model:subject="subject" />
      </center>
    </div>
    <div>
      <markdown-editor title="Question" v-model="question"></markdown-editor>
    </div>
    <div>
      <markdown-editor title="Answer" v-model="answer"></markdown-editor>
    </div>
    <div class="submit-button-container">
      <button class="submit-button" @click="handleSubmit">Submit</button>
    </div>
  </div>
  <van-back-top />
</template>

<script setup>
import { ref, computed } from 'vue';
import axios from 'axios';
import router from "@/router";
import { useStore } from "vuex";
import MarkdownEditor from '@/components/MarkdownEditor.vue';
import DifficultySelector from "@/components/DifficultySelector.vue";
import NavigateBar from "@/components/NavigateBar.vue";
import SubjectSelector from "@/components/SubjectSelector.vue";

const question = ref('');
const answer = ref('');
const difficulty = ref(2); // 默认难度
const subject = ref('history');
const store = useStore();

const storeUsername = computed(() => store.state.username);

const handleSubmit = async () => {
  const payload = {
    question: question.value,
    answer: answer.value,
    difficulty: difficulty.value,
    subject: subject.value,
    username: storeUsername.value
  };
  try {
    const response = await axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/addQuestion/simpleAnswer', payload);
    console.log('Data submitted successfully:', response.data);
    if (response.status === 200 && response.data.success) {
      router.push('/home');
    }
  } catch (error) {
    console.error('Error submitting data:', error);
  }
};
</script>

<style scoped>
.submit-button-container {
  display: flex;
  justify-content: center;
  margin: 20px 0;
}

.submit-button {
  background: linear-gradient(45deg, #f39c12, #e74c3c);
  border: none;
  border-radius: 25px;
  color: white;
  padding: 10px 20px;
  font-size: 18px;
  font-weight: bold;
  text-transform: uppercase;
  cursor: pointer;
  transition: background 0.3s ease, transform 0.3s ease;
}

.submit-button:hover {
  background: linear-gradient(45deg, #e67e22, #d35400);
  transform: scale(1.05);
}

.submit-button:active {
  transform: scale(0.95);
}

.submit-button:focus {
  outline: none;
}
</style>
