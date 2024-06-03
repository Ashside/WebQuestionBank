<template>
  <div>
    <div>
      <markdown-editor title="Question"></markdown-editor>
    </div>
    <div>
      <markdown-editor title="Answer"></markdown-editor>
    </div>
  </div>
  <van-back-top />
</template>

<script>
import MarkdownEditor from '@/components/MarkdownEditor.vue';

export default {
  components: {
    MarkdownEditor
  },
  setup() {
    const list = [...Array(50).keys()];
    return { list };
  }
};
</script>
