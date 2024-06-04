<template>
  <div class="markdown-body" v-html="compiledMarkdown"></div>
</template>

<script>
import { marked } from 'marked';
import 'github-markdown-css/github-markdown.css';

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

<style scoped>
.markdown-body {
  color: black;
  padding: 20px;
  border: 1px solid #e1e4e8;
  border-radius: 6px;
  background-color: #ffffff;
  box-shadow: 0 1px 3px rgba(27, 31, 35, 0.12);
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji";
}

/* You can add more custom styles here */
</style>
