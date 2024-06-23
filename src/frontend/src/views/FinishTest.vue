<template>
  <NavigateBar></NavigateBar>
  <h1>查看我所有的测试</h1>
  <ul>
    <li v-for="(test, index) in tests" :key="index">
      <div class="question-header">
        <div class="title-and-tag">
          <h3>{{ test.name }}</h3>
          <n-tag v-if="test.state === 'to_be_finish'" type="error">待完成</n-tag>
          <n-tag v-if="test.state === 'finish'" type="success">已完成</n-tag>
        </div>
        <button v-if="test.state === 'to_be_finish'">去完成试卷</button>
        <button v-if="test.state === 'finish'">查看试卷评阅进度</button>
      </div>
    </li>
  </ul>
</template>

<script>
import axios from 'axios';
import NavigateBar from "@/components/NavigateBar.vue";
import store from "@/store";

export default {
  components: {NavigateBar},

  created() {
    this.fetchTests();
  },

  data() {
    return {
      tests: [],
      showModal: false,
      testDetails: '',
      sameTestDetails: '',
      showSameTest: false,
      testID: -1,
    }
  },

  methods: {

    async fetchTests() {
      try {
        const response = await axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/queryAllTestsByStudentID', {
          username: store.state.username
        });
        if (response.data.success) {
          this.tests = response.data.test;
        } else {
          console.error('Failed to fetch tests:', response.data.reason);
        }
      } catch (error) {
        console.error('Error fetching tests:', error);
      }
    },
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

.title-and-tag {
  display: flex;
  align-items: center;
}

.title-and-tag h3 {
  margin: 0;
  margin-right: 10px; /* 调整标题与标签的间距 */
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
