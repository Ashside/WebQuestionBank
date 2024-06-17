<template>
  <NavigateBar></NavigateBar>
  <h1>查看所有测试</h1>
  <ul>
    <li v-for="(test, index) in tests" :key="index">
      <div class="question-header">
        <h3>{{ test.name }}</h3>
        <button @click="viewTestDetails(test.id)">查看试卷</button>
      </div>
    </li>
  </ul>
</template>

<script>
import axios from 'axios';
import NavigateBar from "@/components/NavigateBar.vue";

export default {
  components: {NavigateBar},

  created() {
    this.fetchTests();
  },

  data() {
    return {
      tests: []
    }
  },

  methods: {
    async fetchTests() {
      try {
        const response = await axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/queryTest');
        if (response.data.success) {
          this.tests = response.data.test;
        } else {
          console.error('Failed to fetch tests:', response.data.reason);
        }
      } catch (error) {
        console.error('Error fetching tests:', error);
      }
    },

    viewTestDetails(testId) {
      // 这里可以根据你的应用需求进行调整，例如跳转到试卷详细页面或者显示一个模态框等
      console.log('Viewing details for test ID:', testId);
      // 例如使用 this.$router.push(`/test/${testId}`); 来进行页面跳转
    }
  }
}
</script>

<style scoped>

li {
  border: 1px solid #ccc;
  margin-top: 10px;
  border-radius: 8px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
  padding: 10px;
  transition: box-shadow 0.3s ease-in-out, transform 0.3s ease;
}

li:hover {
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
  transform: translateY(-2px);
}

.question-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.question-header h3 {
  margin: 0;
  color: #333;
  font-size: 18px;
}

button {
  padding: 6px 12px;
  background-color: #42a5f5;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  outline: none;
}

button:hover {
  background-color: #2a2a72;
}
</style>
