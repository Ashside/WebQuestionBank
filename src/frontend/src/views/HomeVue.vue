<template>
  <div class="home">
    <h1>Welcome to Home!</h1>
    <p>This is the Home view.</p>
    <div>
    <MarkdownRenderer :content="markdownContent"></MarkdownRenderer>
    </div>
  </div>
</template>

<script>
import MarkdownRenderer from "@/components/MarkdownRenderer.vue";

export default {
  name: 'HomeVue',
  components: {MarkdownRenderer},
  data() {
    return {
      markdownContent: `
  # 这是一个标题

  这是一些文本。

  - 列表项1
  - 列表项2
  - 列表项3

  \`\`\`javascript
  console.log('Hello, Markdown!');
  \`\`\`
      `
    };
  }
};
</script>

<style scoped>
.home {
  text-align: center;
}
</style>
