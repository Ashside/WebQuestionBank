<template>
  <div class="difficulty-selector">
    <input type="radio" id="easy" value=1 v-model="localDifficulty" @change="updateValue" />
    <label for="easy" class="easy">Easy</label>
    <input type="radio" id="medium" value=2 v-model="localDifficulty" @change="updateValue" />
    <label for="medium" class="medium">Medium</label>
    <input type="radio" id="hard" value=3 v-model="localDifficulty" @change="updateValue" />
    <label for="hard" class="hard">Hard</label>
  </div>
</template>

<script>
export default {
  name: 'DifficultySelector',
  props: {
    difficulty: {
      type: Number,
      required: true
    }
  },
  data() {
    return {
      localDifficulty: this.difficulty
    };
  },
  watch: {
    difficulty(newVal) {
      this.localDifficulty = newVal;
    }
  },
  methods: {
    updateValue() {
      this.$emit('update:difficulty', parseInt(this.localDifficulty, 10));
    }
  }
};
</script>


<style scoped>
.difficulty-selector {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 20px 0;
}

.difficulty-selector input[type="radio"] {
  display: none;
}

.difficulty-selector label {
  padding: 10px 20px;
  border: none; /* 移除边框 */
  cursor: pointer;
  transition: background-color 0.3s ease, color 0.3s ease;
  font-size: 16px;
  font-weight: bold;
  text-transform: uppercase;
  /* 添加如下属性来使标签变成正方形 */
  min-width: 100px;
  margin: 0; /* 移除边距让它们紧贴 */
}

.difficulty-selector label.easy {
  background-color: #333333;
  color: white;
}

.difficulty-selector label.medium {
  background-color: #333333;
  color: white;
}

.difficulty-selector label.hard {
  background-color: #333333;
  color: white;
}

.difficulty-selector input[type="radio"]:checked + label.easy {
  background-color: #2e7d32;
  color: #fff;
  box-shadow: 0 0 15px rgba(46, 125, 50, 0.5);
}

.difficulty-selector input[type="radio"]:checked + label.medium {
  background-color: #ffb300;
  color: #fff;
  box-shadow: 0 0 15px rgba(255, 179, 0, 0.5);
}

.difficulty-selector input[type="radio"]:checked + label.hard {
  background-color: #d32f2f;
  color: #fff;
  box-shadow: 0 0 15px rgba(211, 47, 47, 0.5);
}
</style>
