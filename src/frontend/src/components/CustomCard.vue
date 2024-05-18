<template>
  <div class="custom-card">
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
  }
};
</script>

<style scoped>
.custom-card {
  max-width: 300px;
}
</style>
