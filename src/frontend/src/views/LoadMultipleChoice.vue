<template>
  <NavigateBar/>
  <div>
    <div>
      <center>
        <h1>题目难度和科目选择</h1>
        <DifficultySelector v-model:difficulty="difficulty" />
        <SubjectSelector v-model:subject="subject" />
      </center>
    </div>
    <div>
      <markdown-editor title="题目" v-model="question"></markdown-editor>
    </div>
    <div>
      <markdown-editor title="选项 1" v-model="option1"></markdown-editor>
    </div>
    <div>
      <markdown-editor title="选项 2" v-model="option2"></markdown-editor>
    </div>
    <div>
      <markdown-editor title="选项 3" v-model="option3"></markdown-editor>
    </div>
    <div>
      <markdown-editor title="选项 4" v-model="option4"></markdown-editor>
    </div>
    <div>
      <center>
        <div class="multiple-choice-selector">
          <input type="checkbox" id="option1" name="choice" value="option1">
          <label for="option1">选项 1</label><p></p>
          <input type="checkbox" id="option2" name="choice" value="option2">
          <label for="option2">选项 2</label><p></p>
          <input type="checkbox" id="option3" name="choice" value="option3">
          <label for="option3">选项 3</label><p></p>
          <input type="checkbox" id="option4" name="choice" value="option4">
          <label for="option4">选项 4</label><p></p>
        </div>
      </center>
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
const selectedAnswers = ref('');
const difficulty = ref(2); // 默认难度
const subject = ref('history');
const store = useStore();
const option1 = ref('');
const option2 = ref('');
const option3 = ref('');
const option4 = ref('');

const storeUsername = computed(() => store.state.username);

const handleSubmit = async () => {
  const choices = document.getElementsByName('choice');
  let selectedValues = [];
  for (let choice of choices) {
    if (choice.checked) {
      selectedValues.push(choice.value);
    }
  }
  const payload = {
    question: question.value,
    answer: selectedAnswers.value = selectedValues.join(', '), // 将选中的答案拼接成字符串
    option: { option1: option1.value, option2: option2.value, option3: option3.value, option4: option4.value },
    difficulty: difficulty.value,
    subject: subject.value,
    username: storeUsername.value
  };
  try {
    const response = await axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/addQuestion/multipleChoice', payload);
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

.multiple-choice-selector {
  padding: 20px;
  margin: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
}

input[type="radio"] {
  margin: 10px;
}

label {
  margin-left: 8px;
  font-size: 16px;
}

button {
  padding: 10px 20px;
  margin-top: 20px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

</style>
