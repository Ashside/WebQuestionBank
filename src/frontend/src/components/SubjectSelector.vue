<template>
  <div class="load-subject-selector">
    <input type="radio" id="math" value="math" v-model="localSubject" @change="updateValue" />
    <label for="math" class="math">Math</label>
    <input type="radio" id="history" value="history" v-model="localSubject" @change="updateValue" />
    <label for="history" class="history">History</label>
    <input type="radio" id="english" value="english" v-model="localSubject" @change="updateValue" />
    <label for="english" class="english">English</label>
  </div>
</template>

<script>
export default {
  name: 'SubjectSelector',
  props: {
    subject: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      localSubject: this.subject
    };
  },
  watch: {
    subject(newVal) {
      this.localSubject = newVal;
    }
  },
  methods: {
    updateValue() {
      this.$emit('update:subject', this.localSubject);
    }
  }
};
</script>

<style scoped>
.load-subject-selector {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 20px 0;
}

.load-subject-selector input[type="radio"] {
  display: none;
}

.load-subject-selector label {
  padding: 10px 20px;
  border: none; /* 移除边框 */
  cursor: pointer;
  transition: background-color 0.3s ease, color 0.3s ease;
  font-size: 16px;
  font-weight: bold;
  text-transform: uppercase;
  /* 添加如下属性来使标签变成方形 */
  min-width: 100px;
  margin: 0; /* 移除边距让它们紧贴 */
}

.load-subject-selector label.math {
  background-color: #333333;
  color: white;
}

.load-subject-selector label.history {
  background-color: #333333;
  color: white;
}

.load-subject-selector label.english {
  background-color: #333333;
  color: white;
}

.load-subject-selector input[type="radio"]:checked + label.math {
  background-color: #1e88e5;
  color: #fff;
  box-shadow: 0 0 15px rgba(30, 136, 229, 0.5);
}

.load-subject-selector input[type="radio"]:checked + label.history {
  background-color: #fb8c00;
  color: #fff;
  box-shadow: 0 0 15px rgba(251, 140, 0, 0.5);
}

.load-subject-selector input[type="radio"]:checked + label.english {
  background-color: #8e24aa;
  color: #fff;
  box-shadow: 0 0 15px rgba(142, 36, 170, 0.5);
}
</style>