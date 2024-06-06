<template>
  <div class="answer-selector">
    <div v-for="option in options" :key="option" class="option">
      <input
          id="checkbox-{{option}}"
          type="checkbox"
          :value="option"
          @change="handleSelectionChange($event, option)"
          :checked="selectedOptions.includes(option)"
      />
      <label :for="`checkbox-${option}`">
        {{ option.toUpperCase() }}
      </label>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MultipleChoiceSelector',
  data() {
    return {
      options: ['a', 'b', 'c', 'd'],
      selectedOptions: [],
    };
  },
  methods: {
    handleSelectionChange(event, option) {
      if (event.target.checked) {
        this.selectedOptions.push(option);
      } else {
        this.selectedOptions = this.selectedOptions.filter(opt => opt !== option);
      }
      this.$emit('update:selected', this.selectedOptions);
    }
  }
};
</script>

<style scoped>
.answer-selector {
  display: flex;
  flex-direction: column;
  padding: 20px;
  border: 2px solid #ccc;
  border-radius: 8px;
  max-width: 300px;
  background-color: #f9f9f9;
}

.option {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.option input[type="checkbox"] {
  accent-color: #0066cc;  /* Change the color of the checkbox */
  margin-right: 10px;     /* Space between the checkbox and the label */
}

.option label {
  cursor: pointer;
  user-select: none;      /* Prevent text selection */
}
</style>
