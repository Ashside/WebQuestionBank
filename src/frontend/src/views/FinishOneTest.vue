<template>
  <div>
    <ShortAnswerBlock
        title="Markdown 编辑器"
        :question="questionText"
        v-model="answerText"
        questionType="short-answer"
    />
    <ShortAnswerBlock
        title="选择题示例"
        question="以下哪项是Vue.js的特性？"
        v-model="selectedOptions"
        questionType="multiple-choice"
        :options="selectedOptions"
    />
  </div>
</template>

<script>
import ShortAnswerBlock from "@/components/AnswerBlock.vue";

export default {
  components: { ShortAnswerBlock },
  data() {
    return {
      questionText: "# 这是一个问题的Markdown内容\n你可以在这里展示问题的详情。",
      answerText: "",
      selectedOptions: [
        { content: '1. 反应式数据绑定', selected: false },
        { content: '2. 虚拟DOM', selected: false },
        { content: '3. 服务器端渲染', selected: false },
        { content: '4. 双向数据绑定', selected: false },
      ],
    };
  },
};
</script>

<style scoped>
#app {
  margin: 20px;
}
</style>
