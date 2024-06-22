<template>
  <NavigateBar></NavigateBar>
  <h1>分发测试</h1>
  <ul>
    <li v-for="(test, index) in tests" :key="index">
      <div class="question-header">
        <h3>{{ test.name }}</h3>
        <button @click="distributeTest(test.id)">分发试卷</button>
      </div>
    </li>
  </ul>
  <div v-if="showModal" class="modal">
    <div class="modal-content">
      <div class="test-details">
        <center>
          <h2>分发试卷</h2>
        </center>
        <MarkdownRenderer :content="testDetails"></MarkdownRenderer>
      </div>
      <div class="students-list">
        <h3>选择学生</h3>
        <ul>
          <li v-for="(student, index) in students" :key="index">
            <input type="checkbox" :id="student.studentID" v-model="selectedStudents" :value="student.studentID">
            <label :for="student.studentID">{{ student.student }}</label>
          </li>
        </ul>
        <center>
          <button @click="submitDistribute(testID)">提交分发</button>
        </center>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import NavigateBar from "@/components/NavigateBar.vue";
import store from "@/store";
import MarkdownRenderer from "@/components/MarkdownRenderer.vue";
import router from "@/router";

export default {
  components: {MarkdownRenderer, NavigateBar},

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
      students: [],
      selectedStudents: [],
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
      this.fetchStudents();
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

    async fetchStudents() {
      try {
        const response = await axios.post(process.env["VUE_APP_API_URL"] + '/api/usr/findAllStudents', {
          username: store.state.username
        });
        if (response.data.success) {
          this.students = response.data.students;
        } else {
          console.error('Failed to fetch students:', response.data.reason);
        }
      } catch (error) {
        console.error('Error fetching students:', error);
      }
    },

    distributeTest(testId) {
      axios.post(process.env["VUE_APP_API_URL"] + `/api/questionBank/queryTestByID`, {
        testId: testId,
        username: store.state.username
      })
          .then(response => {
            if (response.data.success) {
              this.testDetails = response.data.test;
              this.openModal();
              this.testID = testId;
              this.showSameTest = false;
              this.fetchStudents(); // Fetch students when opening the modal
            } else {
              console.error('Failed to fetch test details:', response.data.reason);
            }
          })
          .catch(error => {
            console.error('Error fetching test details:', error);
          });
    },

    submitDistribute(testId){
      this.showSameTest = true;
      axios.post(process.env["VUE_APP_API_URL"] + `/api/questionBank/distributeTest`, {
        testID: testId,
        username: store.state.username,
        students: this.selectedStudents, // Pass selected students
      })
          .then(response => {
            if (response.data.success) {
              router.push('/home');
            } else {
              console.error('Failed to fetch test details:', response.data.reason);
            }
          })
          .catch(error => {
            console.error('Error fetching test details:', error);
          });
    }
  }
}
</script>

<style scoped>
/* 保持现有样式不变 */
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
  display: flex;
  flex-direction: row;
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  max-width: 90%;
  box-shadow: 0 4px 8px rgba(0,0,0,0.3);
}

.test-details,
.students-list {
  flex: 1;
  margin: 0 10px;
}

div.markdown-container {
  gap: 10px;
  margin-bottom: 30px;
  display: flex;
  justify-content: center;
}

.students-list ul {
  list-style-type: none;
  padding: 0;
}

.students-list li {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.students-list input[type="checkbox"] {
  margin-right: 10px;
}

</style>
