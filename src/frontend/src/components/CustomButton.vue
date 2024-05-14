<template>
  <button :class="buttonClass" @click="handleClick">
    <slot></slot>
  </button>
</template>

<script>
export default {
  name: 'CustomButton',
  props: {
    type: {
      type: String,
      default: 'button',
    },
    primary: {
      type: Boolean,
      default: false,
    },
    secondary: {
      type: Boolean,
      default: false,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    to: {
      type: String,
      default: '',
    },
  },
  computed: {
    buttonClass() {
      return {
        'btn-primary': this.primary,
        'btn-secondary': this.secondary,
        'btn-disabled': this.disabled,
      };
    },
  },
  methods: {
    handleClick(event) {
      if (!this.disabled) {
        if (this.to) {
          this.$router.push(this.to);
        } else {
          this.$emit('click', event);
        }
      }
    },
  },
};
</script>

<style scoped>
button {
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease, transform 0.2s ease;
}

button:active {
  transform: scale(0.98);
}

.btn-primary {
  background-color: #007bff;
  color: white;
}

.btn-primary:hover {
  background-color: #0056b3;
}

.btn-secondary {
  background-color: #6c757d;
  color: white;
}

.btn-secondary:hover {
  background-color: #5a6268;
}

.btn-disabled {
  background-color: #d6d6d6;
  color: #888;
  cursor: not-allowed;
}
</style>
