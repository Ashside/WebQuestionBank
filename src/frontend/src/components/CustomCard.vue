<template>
  <div class="custom-card" @click="handleClick">
    <n-card :title="title" hoverable=true>
      <template v-if="cover" #cover>
        <img :src="cover">
      </template>
      <slot></slot>
    </n-card>
  </div>
</template>

<script>
export default {
  name: 'CustomCard',
  props: {
    title: String,
    cover: String
  },
  methods: {
    handleClick() {
      this.$emit('card-click');
    }
  }
};
</script>

<style scoped>
.custom-card {
  max-width: 300px;
  max-height: 500px;
}
.custom-card img {
  width: 100%;  /* 使图片宽度充满容器 */
  height: 300px; /* 设置固定的高度 */
  object-fit: cover; /* 保持图片的比例，多余的部分会被裁剪 */
}
</style>
