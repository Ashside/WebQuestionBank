<template>
  <div v-html="compiledMarkdown"></div>
</template>

<script>
import { marked } from 'marked';

export default {
  name: 'MarkdownRenderer',
  props: {
    content: {
      type: String,
      required: true
    }
  },
  computed: {
    compiledMarkdown() {
      return marked(this.content);
    }
  }
};
</script>
