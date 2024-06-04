<template>
  <div class="subject-selector">
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
.subject-selector {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 20px;
  margin: 20px 0;
}

.subject-selector input[type="radio"] {
  display: none;
}

.subject-selector label {
  padding: 10px 20px;
  border-radius: 25px;
  cursor: pointer;
  transition: background-color 0.3s ease, color 0.3s ease;
  font-size: 16px;
  font-weight: bold;
  text-transform: uppercase;
}

.subject-selector label.math {
  background-color: black;
  color: white;
}

.subject-selector label.history {
  background-color: black;
  color: white;
}

.subject-selector label.english {
  background-color: black;
  color: white;
}

.subject-selector input[type="radio"]:checked + label.math {
  background-color: #1e88e5;
  color: #fff;
  box-shadow: 0 0 15px rgba(30, 136, 229, 0.5);
}

.subject-selector input[type="radio"]:checked + label.history {
  background-color: #fb8c00;
  color: #fff;
  box-shadow: 0 0 15px rgba(251, 140, 0, 0.5);
}

.subject-selector input[type="radio"]:checked + label.english {
  background-color: #8e24aa;
  color: #fff;
  box-shadow: 0 0 15px rgba(142, 36, 170, 0.5);
}
</style>
