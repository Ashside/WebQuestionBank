<template>
  <div class="markdown-editor" :style="{ height: containerHeight }">
    <div class="container">
      <h2>{{ title }}</h2>
      <div class="editor-preview-container">
        <div class="display-pane">
          <MarkdownRenderer :content="question" />
        </div>
        <div class="editor-pane">
          <div v-if="questionType === 'multiple-choice'" class="choices">
            <div v-for="(option, index) in internalOptions" :key="index" class="choice">
              <input type="checkbox" :id="'option-' + index" v-model="option.selected">
              <label :for="'option-' + index">
                <MarkdownRenderer :content="option.content" />
              </label>
            </div>
          </div>
          <div v-if="!isMarkdown && questionType === 'short-answer'">
            <textarea v-model="answer" v-auto-resize placeholder="请输入答案..."></textarea>
          </div>
          <div v-if="isMarkdown">
            <MarkdownRenderer :content="answer" v-if="questionType === 'short-answer'" />
            <div v-if="questionType === 'multiple-choice'">
              <div v-for="(option, index) in internalOptions" :key="index" class="choice">
                <MarkdownRenderer :content="option.content" />
              </div>
            </div>
          </div>
          <button @click="toggleMarkdown" v-if="questionType === 'short-answer'">
            {{ isMarkdown ? '编辑' : '渲染' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, watch, computed } from 'vue';
import MarkdownRenderer from '@/components/MarkdownRenderer.vue';
import autoResize from '@/directives/autoResize';

export default {
  components: {
    MarkdownRenderer,
  },
  directives: {
    autoResize,
  },
  props: {
    title: {
      type: String,
      required: true,
    },
    question: {
      type: String,
      required: true,
    },
    modelValue: {
      type: [String, Array],
      required: true,
    },
    questionType: {
      type: String,
      default: 'short-answer',
    },
    options: {
      type: Array,
      default: () => [],
    },
  },
  setup(props, { emit }) {
    const answer = ref(props.modelValue);
    const isMarkdown = ref(false);

    watch(() => props.modelValue, (newValue) => {
      answer.value = newValue;
    });

    watch(answer, (newValue) => {
      emit('update:modelValue', newValue);
    });

    watch(() => props.options, (newValue) => {
      internalOptions.value = newValue;
    }, { deep: true });

    const internalOptions = ref(props.options.map(option => ({
      content: option.content || '',
      selected: option.selected || false,
    })));

    const toggleMarkdown = () => {
      isMarkdown.value = !isMarkdown.value;
    };

    const containerHeight = computed(() => {
      const baseHeight = 15;
      const lineHeight = 1.4;
      let extraHeight = 0;

      if (props.questionType === 'short-answer') {
        const content = typeof answer.value === 'string' ? answer.value : '';
        extraHeight = (content.split('\n').length * lineHeight) + 15;
      } else if (props.questionType === 'multiple-choice') {
        extraHeight = internalOptions.value.reduce((total, option) => {
          return total + (option.content.split('\n').length * lineHeight) + 2;
        }, 15);
      }

      return `${Math.min(baseHeight + extraHeight, 70)}vh`;
    });

    return {
      answer,
      isMarkdown,
      toggleMarkdown,
      containerHeight,
      internalOptions,
    };
  },
};
</script>


<style scoped>
.markdown-editor {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
  max-height: 70vh;
  overflow: hidden;
}

.container {
  display: flex;
  flex-direction: column;
  width: 70%;
  height: 100%;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.12);
  overflow: hidden;
  padding: 15px;
}

.container h2 {
  margin-bottom: 15px;
  font-size: 20px;
  color: #333;
  text-align: center;
}

.editor-preview-container {
  display: flex;
  flex: 1;
  overflow: hidden;
}

.display-pane,
.editor-pane {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: auto;
  padding: 15px;
}

.editor-pane textarea {
  width: 100%;
  height: auto;
  border: 1px solid #ccc;
  resize: none;
  padding: 10px;
  font-size: 14px;
  line-height: 1.4;
  border-radius: 4px;
  box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.1);
}

.editor-pane button {
  margin-top: 10px;
  padding: 10px;
  font-size: 14px;
  cursor: pointer;
}

.display-pane {
  border-right: 1px solid #ddd;
}

.display-pane > * {
  width: 100%;
}

.editor-pane textarea:focus {
  outline: none;
  box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.2);
}

button {
  background-color: #1e88e5;
  color: white;
  border: none;
  padding: 10px 20px;
  font-size: 16px;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
  display: block;
  text-align: center;
}

button:hover {
  background-color: #2a2a72;
}

.choices {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.choice {
  display: flex;
  align-items: center;
  gap: 10px;
}

.choice input {
  cursor: pointer;
}

.choice label {
  flex: 1;
}

/* 自定义复选框样式 */
input[type="checkbox"] {
  -webkit-appearance: none; /* 移除默认外观 */
  appearance: none;
  background-color: #fff;
  margin: 0 10px 0 0;
  font-size: 1.5em;
  color: #42a5f5;
  width: 20px;
  height: 20px;
  border: 2px solid #42a5f5;
  border-radius: 4px;
  cursor: pointer;
  position: relative;
}

input[type="checkbox"]:checked {
  background-color: #42a5f5;
}

input[type="checkbox"]:checked::after {
  content: "✔";
  position: absolute;
  top: -2px;
  left: 2px;
  color: #fff;
  font-size: 16px;
}
</style>
