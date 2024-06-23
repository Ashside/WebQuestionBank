<template>
  <div class="markdown-editor" :style="{ height: containerHeight }">
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
import {computed, ref, watch} from 'vue';
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

    // 计算容器高度，最小为30vh，最大为70vh
    const containerHeight = computed(() => {
      const baseHeight = 15; // 最小高度
      const lineHeight = 1.4; // 行高
      const extraHeight = (localInputText.value.split('\n').length * lineHeight) + 15; // 额外高度计算
      return `${Math.min(baseHeight + extraHeight, 70)}vh`; // 限制最大高度为70vh
    });

    return {
      localInputText,
      containerHeight
    };
  },
};
</script>

<style scoped>
.markdown-editor {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 30vh; /* 将高度调整小一点 */
  padding: 20px;
}

.container {
  display: flex;
  flex-direction: column;
  width: 70%; /* 默认的容器宽度调小 */
  height: 100%; /* 调整为100%以适应markdown-editor的高度 */
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.12); /* 更明显的阴影效果 */
  overflow: hidden; /* 隐藏滚动条 */
  padding: 15px; /* 调整内部间距 */
}

.container h2 {
  margin-bottom: 15px; /* 减少标题与内容的间隔 */
  font-size: 20px; /* 字体大小调整 */
  color: #333; /* 字体颜色更深 */
  text-align: center;
}

.editor-preview-container {
  display: flex;
  flex: 1;
  overflow: hidden; /* 隐藏内部滚动条 */
}

.editor-pane,
.preview-pane {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: auto; /* 允许滚动 */
  padding: 15px; /* 内部间距调整 */
}

.editor-pane textarea {
  width: 100%;
  height: 100%;
  border: 1px solid #ccc; /* 添加边框 */
  resize: none;
  padding: 10px;
  font-size: 14px; /* 字体大小调整 */
  line-height: 1.4; /* 行高调整 */
  border-radius: 4px;
  box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.1);
}

.preview-pane {
  border-left: 1px solid #ddd;
  padding: 15px; /* 调整内部间距 */
}

.preview-pane > * {
  width: 100%; /* 宽度调整为100% */
}

.editor-pane textarea:focus {
  outline: none;
  box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.2);
}
</style>
