<template>
  <Navigate></Navigate>
  <div class="home">
    <h1>Welcome to Home!</h1>
    <p>This is the Home view.</p>
    <div>
    </div>
  </div>
</template>

<script>
import NavigateBar from "@/components/NavigateBar.vue";

export default {
  name: 'HomeVue',
  components: {Navigate: NavigateBar}
};
</script>

<style scoped>
.home {
  text-align: center;
}
</style>
