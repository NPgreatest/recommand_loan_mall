<template>
  <div>
    <van-cell title="钱包余额">{{ walletBalance }}</van-cell>
    <van-cell title="还款期限">{{ repaymentTerm }}天</van-cell>
    <van-button type="primary" @click="handleRepayment">还款</van-button>

    <van-field v-model="loanAmount" label="借款金额" input-align="right" placeholder="请输入借款金额">
      <template #input>
        <van-number-keyboard v-model="loanAmount" />
      </template>
    </van-field>


    <van-field v-model="selectedTerm" label="借款期限" input-align="right" placeholder="请输入借款金额">
      <template #input>
        <van-number-keyboard v-model="selectedTerm" />
      </template>
    </van-field>


    <van-button type="info" @click="applyForLoan">申请借款</van-button>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue';
// import { getWalletInfo, repayLoan, applyLoan } from '@/service/loanService'; // 假设这些是API调用

const walletBalance = ref(0);
const repaymentTerm = ref(60); // 默认60天
const loanAmount = ref('');
const selectedTerm = ref(new Date());

// 在页面加载时获取钱包信息和还款期限
onMounted(async () => {
  // const walletInfo = await getWalletInfo();
  // walletBalance.value = walletInfo.balance;
  // repaymentTerm.value = walletInfo.term;
});

const handleRepayment = async () => {
  // const result = await repayLoan();
  // 根据实际情况处理还款后的逻辑，例如显示提示信息
};

const applyForLoan = async () => {
  // const result = await applyLoan({ amount: loanAmount.value, term: repaymentTerm.value });
  // 根据实际情况处理借款申请后的逻辑，例如显示提示信息
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
