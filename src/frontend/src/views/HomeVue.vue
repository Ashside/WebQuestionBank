<template>
  <Navigate></Navigate>
  <div class="home">
    <h1>欢迎来到首页！</h1>
    <p>请选择你要进行的操作</p>
    <div class="card-container">
      <CustomCard title="录入简答题" cover="/simple_answer.png" @card-click="navigateTo('LoadQuestion/shortAnswer')"></CustomCard>
      <CustomCard title="录入选择题" cover="/multiple_choice.png" @card-click="navigateTo('LoadQuestion/multipleChoice')"></CustomCard>
      <CustomCard title="查看题库" cover="/question_bank.png" @card-click="navigateTo('/ViewQuestion')"></CustomCard>
      <!-- 可以根据需要添加更多卡片 -->
    </div>
  </div>
</template>

<script>
import NavigateBar from "@/components/NavigateBar.vue";
import CustomCard from "@/components/CustomCard.vue";
import {useRouter} from "vue-router";

export default {
  name: 'HomeVue',
  components: {
    CustomCard,
    Navigate: NavigateBar
  },
  setup() {
    const router = useRouter();

    const navigateTo = (path) => {
      router.push({ path });
    };

    return {
      navigateTo
    };
  }
};
</script>

<style scoped>
.home {
  text-align: center;
}

.card-container {
  display: flex;
  justify-content: center;
  gap: 16px; /* 卡片之间的间距 */
  flex-wrap: wrap; /* 当空间不足时，换行排列 */
}

.card-container > * {
  flex: 1;
  min-width: 200px; /* 每个卡片的最小宽度 */
}
</style>
