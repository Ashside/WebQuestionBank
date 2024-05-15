<template>
  <div class="markdown-editor">
    <div class="editor-pane">
      <textarea v-model="inputText" @input="updateDisplay"></textarea>
    </div>
    <div class="preview-pane">
      <!-- 直接将inputText传递给显示器组件 -->
      <MarkdownRenderer :content="inputText"></MarkdownRenderer>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'; // 使用Vue 3的引入方式
import MarkdownRenderer from "@/components/MarkdownRenderer.vue";

export default {
  components: {
    MarkdownRenderer,
  },
  setup() {
    const inputText = ref('');

    const updateDisplay = (event) => {
      inputText.value = event.target.value;
    };

    return {
      inputText,
      updateDisplay
    };
  }
};
</script>

<style scoped>
.markdown-editor {
  display: flex;
  height: 100%;
}

.editor-pane,
.preview-pane {
  flex: 1;
  padding: 10px;
}

.editor-pane textarea {
  width: 100%;
  height: 100%;
}

.preview-pane {
  border-left: 1px solid #ddd;
  overflow-y: auto;
}
</style>
