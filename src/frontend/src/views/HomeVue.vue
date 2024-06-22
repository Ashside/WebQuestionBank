<template>
  <Navigate></Navigate>
  <div class="home">
    <div class="card-container">
      <CustomCard v-if="!isStudent" title="录入简答题" cover="/simple_answer.png" @card-click="navigateTo('LoadQuestion/shortAnswer')"></CustomCard>
      <CustomCard v-if="!isStudent" title="录入选择题" cover="/multiple_choice.png" @card-click="navigateTo('LoadQuestion/multipleChoice')"></CustomCard>
      <CustomCard v-if="!isStudent" title="查看题库" cover="/question_bank.png" @card-click="navigateTo('/ViewQuestion')"></CustomCard>
      <CustomCard v-if="!isStudent" title="组卷" cover="/make_test.png" @card-click="navigateTo('/MakeTest')"></CustomCard>
      <CustomCard v-if="isStudent" title="完成试卷" cover="/exam.png" @card-click="navigateTo('/MakeTest')"></CustomCard>
      <CustomCard title="查看试卷" cover="/view_test.png" @card-click="navigateTo('/ViewAllTests')"></CustomCard>
      <CustomCard v-if="!isStudent" title="批阅试卷" cover="/test_score.png" @card-click="navigateTo('/CheckStudentAnswer')"></CustomCard>
      <CustomCard v-if="!isStudent" title="分发试卷" cover="/distribute_exam.png" @card-click="navigateTo('/DistributeTest')"></CustomCard>
    </div>
  </div>
</template>

<script>
import NavigateBar from "@/components/NavigateBar.vue";
import CustomCard from "@/components/CustomCard.vue";
import {useRouter} from "vue-router";
import {computed} from "vue";
import store from "@/store";

const storeRole = computed(() => store.state.role);
const isStudent = computed(() => storeRole.value === 'student');

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
      navigateTo,
      isStudent
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
  padding: 100px; /* 卡片与边框之间的间距 */
}

.card-container > * {
  width: 250px; /* 固定宽度 */
  flex-grow: 0;
  flex-shrink: 0;
}
</style>
