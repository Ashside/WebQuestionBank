<template>
  <div>
<!--    <AnsweringSituation-->
<!--        title="editorTitle"-->
<!--        :question="editorQuestion"-->
<!--        :correctAnswer="correctAnswer"-->
<!--        questionType="simpleAnswer"-->
<!--        answer="studentAnswer"-->
<!--    />-->
    <AnsweringSituation
        :title="editorTitle"
        :question="editorQuestion"
        :correctAnswer="correctAnswer"
        :questionType="questionType"
        :options="options"
        :studentOptions="studentOptions"
        :answer="Answer"
        :studentAnswer="studentAnswer"
    />
  </div>
</template>

<script>
import AnsweringSituation from "@/components/AnsweringSituation.vue";

export default {
  components: {
    AnsweringSituation,
  },
  data() {
    return {
      editorTitle: '示例问题',
      editorQuestion: '这是一个示例问题的描述，你需要选择正确的答案。',
      correctAnswer: '正确答案是选项A。',
      questionType: 'multipleChoice',
      // questionType: 'simpleAnswer',
      studentAnswer: "option1, option2",
      Answer: "option2",
      options: [
        { content: '选项A', selected: true },
        { content: '选项B', selected: false },
        { content: '选项C', selected: false },
      ],
      studentOptions: [
        { content: '选项A', selected: true },
        { content: '选项B', selected: true },
        { content: '选项C', selected: false },
      ],
    };
  },
};
</script>

<style scoped>
/* 父组件样式 */
</style>
