<template>
  <div class="markdown-editor">
    <div class="container">
      <h2>{{ title }}</h2>
      <div class="editor-preview-container">
        <div class="editor-pane">
          <textarea v-model="localInputText" v-auto-resize placeholder="请输入Markdown内容..."></textarea>
        </div>
        <div class="preview-pane">
          <MarkdownRenderer :content="localInputText" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, watch } from 'vue';
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
    modelValue: {
      type: String,
      required: true,
    },
  },
  setup(props, { emit }) {
    const localInputText = ref(props.modelValue);

    watch(() => props.modelValue, (newValue) => {
      localInputText.value = newValue;
    });

    watch(localInputText, (newValue) => {
      emit('update:modelValue', newValue);
    });

    return {
      localInputText,
    };
  },
};
</script>

<style scoped>
.markdown-editor {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60vh;
  padding: 20px;
}

.container {
  display: flex;
  flex-direction: column;
  width: 80%;
  height: 80%;
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: auto;
  padding: 20px;
}

.container h2 {
  margin: 0 0 20px;
  font-size: 24px;
  text-align: center;
}

.editor-preview-container {
  display: flex;
  flex: 1;
  overflow: auto;
}

.editor-pane,
.preview-pane {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: auto;
  padding: 20px;
}

.editor-pane textarea {
  width: 100%;
  height: 100%;
  border: none;
  resize: none;
  padding: 10px;
  font-size: 16px;
  line-height: 1.5;
  border-radius: 4px;
  box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.1);
}

.preview-pane {
  border-left: 1px solid #ddd;
  padding: 20px;
}

.preview-pane > * {
  width: 90%;
}

.editor-pane textarea:focus {
  outline: none;
  box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.2);
}
</style>
