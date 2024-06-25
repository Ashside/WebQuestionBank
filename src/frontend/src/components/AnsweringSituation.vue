<template>

<!--  题目模块-->
  <div class="markdown-editor" :style="{ height: containerHeight }">
    <div class="container">
      <h2>{{ title }}</h2>
      <div class="editor-preview-container">
        <div class="display-pane">
          <h2>题目</h2>
          <MarkdownRenderer :content="question" />
        </div>

<!--        标准答案模块-->
        <div class="display-pane">
          <h2>标准答案</h2>
          <div v-if="questionType === 'multipleChoice'" class="choices">
            <div v-for="(option, index) in internalOptions" :key="index" class="choice">
              <input type="checkbox" :id="'option-' + index" :checked="option.selected">
              <label :for="'option-' + index">
                <MarkdownRenderer :content="option.content" />
              </label>
            </div>
          </div>
            <MarkdownRenderer :content="answer" v-if="questionType === 'simpleAnswer'" />
        </div>

<!--        学生答案模块-->
        <div class="editor-pane">
          <h2>你的答案</h2>
          <div v-if="questionType === 'multipleChoice'" class="choices">
            <div v-for="(studentOption, index) in internalStudentOptions" :key="index" class="choice">
              <input type="checkbox" :id="'option-' + index" :checked="studentOption.selected">
              <label :for="'option-' + index">
                <MarkdownRenderer :content="studentOption.content" />
              </label>
            </div>
          </div>
          <MarkdownRenderer :content="studentAnswer" v-if="questionType === 'simpleAnswer'" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue';
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
      default: 'simpleAnswer',
    },
    answer: {
      type: String,
      default: '',
    },
    studentAnswer: {
      type: String,
      default: '',
    },
    options: {
      type: Array,
      default: () => [],
    },
    studentOptions: {
      type: Array,
      default: () => [],
    },
  },
  setup(props) {
    const answer = ref(props.modelValue);
    const isMarkdown = ref(false);
    const internalOptions = ref(
        props.options.map((option, index) => ({
          content: option.content || '',
          selected: option.selected || false,
          index: index,
        }))
    );
    const internalStudentOptions = ref(
        props.studentOptions.map((studentOption, index) => ({
          content: studentOption.content || '',
          selected: studentOption.selected || false,
          index: index,
        }))
    );


    const toggleMarkdown = () => {
      isMarkdown.value = !isMarkdown.value;
    };

    const containerHeight = computed(() => {
      const baseHeight = 30;
      const lineHeight = 1.4;
      let extraHeight = 0;

      if (props.questionType === 'simpleAnswer') {
        const content = typeof answer.value === 'string' ? answer.value : '';
        extraHeight = (content.split('\n').length * lineHeight) + 15;
      } else if (props.questionType === 'multipleChoice') {
        extraHeight = internalOptions.value.reduce((total, option) => {
          return total + (option.content.split('\n').length * lineHeight) + 2;
        }, 15);
      }

      return `${Math.min(baseHeight + extraHeight, 70)}vh`;
    });

    return {
      isMarkdown,
      toggleMarkdown,
      containerHeight,
      internalOptions,
      internalStudentOptions,
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
