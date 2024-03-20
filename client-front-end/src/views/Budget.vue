<template>
  <div>
    <s-header :name="'借贷管理'"></s-header>
    <van-cell title="钱包余额">{{ walletBalance }}</van-cell>
    <van-cell title="还款期限">{{ repaymentTerm }}天</van-cell>
    <van-field v-model="loanAmount" label="借款金额" input-align="right" placeholder="请输入借款金额" type="number"></van-field>
    <van-field v-model="selectedTerm" label="借款期限" input-align="right" placeholder="请选择借款期限" type="number"></van-field>
    <van-button type="primary" @click="handleRepayment">还款</van-button>
    <van-button type="info" @click="tryGetLoan">申请借款</van-button>
    <van-loading v-if="isLoading" type="spinner" size="24px">数据获取中...</van-loading>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue';
// import { getWalletInfo, repayLoan, applyLoan } from '@/service/loanService'; // 假设这些是API调用
import sHeader from '@/components/SimpleHeader.vue';
import {tryLoan} from "@/service/user";
const walletBalance = ref(0);
const repaymentTerm = ref(60); // 默认60天
const loanAmount = ref('');
const selectedTerm = ref(30);
const isLoading = ref(false); // 控制加载条的显示
const loanResult = ref(null); // 存储贷款尝试的结果

const tryGetLoan = async () => {
  isLoading.value = true; // 开始显示加载指示器
  try {
    const response = await tryLoan({  amount: parseInt(loanAmount.value, 10),
      term: parseInt(selectedTerm.value, 10), }); // 传递参数
    loanResult.value = response.data; // 假设后端响应的数据在response.data中
    console.log(loanResult.value); // 输出结果以供调试
  } catch (error) {
    console.error('贷款尝试失败', error);
  } finally {
    isLoading.value = false; // 完成请求，隐藏加载指示器
  }
};

onMounted(async () => {
  // const walletInfo = await getWalletInfo();
  // walletBalance.value = walletInfo.balance;
  // repaymentTerm.value = walletInfo.term;
});


const handleRepayment = async () => {
  // const result = await repayLoan();
  // 根据实际情况处理还款后的逻辑，例如显示提示信息
};

const onConfirm = (value) => {
  selectedTerm.value = value;
  repaymentTerm.value = value.getHours() * 60; // 转换为天数
};

const onCancel = () => {
  // 取消选择还款期限时的处理逻辑
};
</script>

<style scoped>
.time-picker {
  /* 样式调整 */
}
</style>
