<template>
  <div class="setting-box">
    <s-header :name="'财务管理'"></s-header>
    <div class="input-item">
      <div class="radio-group-item">
        <div class="label">性别</div>
        <van-radio-group v-model="financeData.gender" direction="horizontal">
          <van-radio :name="false">男</van-radio>
          <van-radio :name="true">女</van-radio>
        </van-radio-group>
      </div>
      <div class="radio-group-item">
        <div class="label">家庭成员数</div>
        <van-radio-group v-model="financeData.dependents" direction="horizontal">
          <van-radio :name="0">0</van-radio>
          <van-radio :name="1">1</van-radio>
          <van-radio :name="2">2</van-radio>
          <van-radio :name="3">3+</van-radio>
        </van-radio-group>
      </div>
      <div class="radio-group-item">
        <div class="label">婚姻状况</div>
        <van-radio-group v-model="financeData.married" direction="horizontal">
          <van-radio :name="false">未婚</van-radio>
          <van-radio :name="true">已婚</van-radio>
        </van-radio-group>
      </div>
      <div class="radio-group-item">
        <div class="label">教育水平</div>
        <van-radio-group v-model="financeData.education" direction="horizontal">
          <van-radio :name="false">其它</van-radio>
          <van-radio :name="true">大学及以上</van-radio>
        </van-radio-group>
      </div>
      <div class="radio-group-item">
        <div class="label">是否自雇</div>
        <van-radio-group v-model="financeData.selfEmployed" direction="horizontal">
          <van-radio :name="false">否</van-radio>
          <van-radio :name="true">是</van-radio>
        </van-radio-group>
      </div>
      <div class="radio-group-item">
        <div class="label">居住地</div>
        <van-radio-group v-model="financeData.city" direction="horizontal">
          <van-radio :name="false">农村</van-radio>
          <van-radio :name="true">城市</van-radio>
        </van-radio-group>
      </div>
      <div class="label">* 申请人收入</div>
      <van-slider v-model="financeData.applicantIncome" max="40000" label="申请人收入" />
      <div class="label">共同申请人收入(配偶)</div>
      <van-slider v-model="financeData.coapplicantIncome" max="40000" label="共同申请人收入" />
    </div>
    <van-button round class="save-btn" color="#1baeae" type="primary" @click="saveFinanceData" block>保存</van-button>
  </div>
</template>

<script setup>
import { reactive, onMounted } from 'vue';
import sHeader from '@/components/SimpleHeader.vue';
import { RadioGroup, Radio, Slider, Button } from 'vant';
import { showSuccessToast } from 'vant';
import { setUserFinance, getUserFinance } from "@/service/user";

const financeData = reactive({
  // gender: false, // 假定性别字段现在是布尔值
  // dependents: 0, // 整数
  // married: false, // 布尔值
  // education: false, // 布尔值
  // selfEmployed: false, // 布尔值
  // applicantIncome: 0.0, // 浮点数
  // coapplicantIncome: 0.0, // 浮点数
  // city: false // 布尔值，之前是residency
})

onMounted(async () => {
  const response = await getUserFinance();
  if (response.resultCode === 200 && response.message === "SUCCESS") {
    // 直接使用返回的 data 对象进行赋值
    const { data } = response;
    Object.assign(financeData, {
      gender: data.gender,
      dependents: data.dependents,
      married: data.married,
      education: data.education,
      selfEmployed: data.self_employed, // 后端返回的是 self_employed，前端使用 selfEmployed
      applicantIncome: data.applicant_income, // 使用正确的字段映射
      coapplicantIncome: data.coapplicant_income, // 使用正确的字段映射
      city: data.city,
    });
    showSuccessToast('数据加载成功'); // 显示数据加载成功的提示
  } else {
    showSuccessToast('数据加载失败'); // 显示数据加载失败的提示，根据实际情况调整
  }
});

const saveFinanceData = async () => {
  const params = {
    gender: financeData.gender,
    dependents: parseInt(financeData.dependents), // 确保是整数
    married: financeData.married,
    education: financeData.education,
    self_employed: financeData.selfEmployed, // 注意字段名的改变
    applicant_income: parseFloat(financeData.applicantIncome), // 确保是浮点数
    coapplicant_income: parseFloat(financeData.coapplicantIncome), // 确保是浮点数
    city: financeData.city === 'city' // 假设之前的residency字段需要转换成布尔值
  };
  await setUserFinance(params);
  showSuccessToast('保存成功');
};
</script>

<style scoped>
.radio-group-item {
  margin-bottom: 16px;
}

.radio-group-item .label {
  margin-bottom: 8px;
}
</style>


<style scoped>
.setting-box {
  max-width: 600px;
  margin: 20px auto;
  padding: 20px;
  box-shadow: 0 0 10px rgba(0,0,0,0.1);
}

.input-item {
  margin-bottom: 20px;
}

.input-item > div {
  margin-bottom: 15px;
}

.save-btn {
  margin-top: 20px;
}

.van-radio-group {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
}

.van-radio {
  margin-right: 10px;
  margin-bottom: 10px;
}

.van-slider {
  margin-bottom: 20px;
}
</style>

