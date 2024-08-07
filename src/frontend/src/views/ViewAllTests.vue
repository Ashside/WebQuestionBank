<template>
  <NavigateBar></NavigateBar>
  <h1>查看所有测试</h1>
  <ul>
    <li v-for="(test, index) in tests" :key="index">
      <div class="question-header">
        <input v-if="isAdmin" type="checkbox" v-model="test.selected">
        <h3>{{ test.name }}</h3>
        <button @click="viewTestDetails(test.id)">查看试卷</button>
      </div>
    </li>
  </ul>
  <div v-if="showModal" class="modal">
    <div class="modal-content">
      <div class="markdown-container">
        <div class="test-details">
          <MarkdownRenderer :content="testDetails"></MarkdownRenderer>
          <br><br>
          <center>
            <button @click="findSameTest(testID)">显示相似试卷</button>
          </center>
        </div>
        <div class="same-test-details" v-if="showSameTest">
          <MarkdownRenderer :content="sameTestDetails"></MarkdownRenderer>
        </div>
      </div>
    </div>
  </div>
  <div class="button-container" v-if="isAdmin">
    <button @click="submitDeleteTests">删除选中的试卷</button>
  </div>
</template>

<script>
import axios from 'axios';
import NavigateBar from "@/components/NavigateBar.vue";
import store from "@/store";
import MarkdownRenderer from "@/components/MarkdownRenderer.vue";
import { computed } from "vue";

const storeRole = computed(() => store.state.role);
const isAdmin = computed(() => storeRole.value === 'admin');

export default {
  components: { MarkdownRenderer, NavigateBar },

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
      isAdmin
    }
  },

  methods: {
    onEscKey(event) {
      if (event.keyCode === 27) {
        this.closeModal();
      }
    },

    closeModal() {
      this.showModal = false;
      document.removeEventListener('keydown', this.onEscKey);
    },

    openModal() {
      this.showModal = true;
      document.addEventListener('keydown', this.onEscKey);
    },

    async fetchTests() {
      try {
        const response = await axios.post(process.env["VUE_APP_API_URL"] + '/api/questionBank/queryAllTests');
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
      axios.post(process.env["VUE_APP_API_URL"] + "/api/questionBank/queryTestByID", {
        testId: testId,
        username: store.state.username
      })
          .then(response => {
            if (response.data.success) {
              this.testDetails = response.data.test;
              this.openModal();
              this.testID = testId;
              this.showSameTest = false;
            } else {
              console.error('Failed to fetch test details:', response.data.reason);
            }
          })
          .catch(error => {
            console.error('Error fetching test details:', error);
          });
    },

    findSameTest(testId) {
      this.sameTestDetails = '';
      axios.post(process.env["VUE_APP_API_URL"] + "/api/questionBank/findSameTestByID", {
        testId: testId,
        username: store.state.username
      })
          .then(response => {
            if (response.data.success) {
              this.sameTestDetails = response.data.test;
              this.showSameTest = true;
            } else {
              alert("匹配失败。原因：" + response.data.reason);
              console.error('Failed to fetch test details:', response.data.reason);
            }
          })
          .catch(error => {
            alert("匹配失败", error);
            console.error('Error fetching test details:', error);
          });
    },

    submitDeleteTests() {
      const selectedTests = this.tests.filter(test => test.selected);
      const selectedTestIds = selectedTests.map(test => test.id);
      axios.post(process.env["VUE_APP_API_URL"] + "/api/questionBank/deleteTestByID", {
        username: store.state.username,
        id: selectedTestIds
      })
          .then(response => {
            if (response.data.success) {
              alert('Deleted successfully!');
              this.fetchTests();
            } else {
              alert('Failed to delete tests.');
            }
          })
          .catch(error => {
            console.error('Error deleting tests:', error);
          });
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
  align-items: center;
  justify-content: space-between;
}

.question-header h3 {
  margin: 0;
  color: #333;
  font-size: 18px;
  flex-grow: 1;
}

.question-header button {
  margin-left: 10px;
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

.modal {
  position: fixed;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0,0,0,0.7);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  max-width: 90%;
  box-shadow: 0 4px 8px rgba(0,0,0,0.3);
  display: flex;
  flex-direction: column;
  width: 80%; /* 调整模态框的宽度 */
}

.markdown-container {
  display: flex;
  justify-content: space-between; /* 水平分布子元素 */
  width: 100%;
}

.test-details,
.same-test-details {
  flex: 1; /* 两个子元素平分宽度 */
  margin: 0 10px; /* 调整间距 */
  max-height: 400px; /* 固定高度 */
  overflow-y: auto; /* 启用滚动条 */
}

.button-container {
  text-align: center;
  margin-top: 20px;
}

</style>
