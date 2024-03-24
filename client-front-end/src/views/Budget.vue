<template>
  <div>
    <s-header :name="'借贷管理'"></s-header>
    <div class="wallet-info">
      <!-- 添加类名 wallet-balance 和 debt 分别用于钱包余额和贷款金额 -->
      <van-cell title="钱包余额" class="wallet-balance">{{ walletBalance }}</van-cell>
      <van-cell title="贷款金额" class="debt">{{ debt }}</van-cell>
      <van-cell title="还款期限" v-if="repaymentTerm > 0">{{ repaymentTerm }}天</van-cell>
      <van-field v-model="repaymentAmount" label="充值/还款金额" input-align="right" placeholder="请输入金额" type="number"></van-field>
      <van-button type="primary" @click="handleRepayment" class="center-button">充值/还款</van-button>
    </div>
    <div v-if="repaymentTerm === 0" class="loan-application">
      <van-field v-model="loanAmount" label="借款金额" input-align="right" placeholder="请输入借款金额" type="number"></van-field>
      <van-field v-model="selectedTerm" label="借款期限" input-align="right" placeholder="请选择借款期限" type="number"></van-field>

      <van-button type="info" @click="tryGetLoan" class="center-button">申请借款</van-button>
    </div>
    <!-- 贷款成功后显示的按钮 -->
    <van-button v-if="loanResult" type="success" @click="doLoanFunction" class="center-button">确认申请</van-button>
    <van-loading v-if="isLoading" type="spinner" size="24px">数据获取中...</van-loading>
  </div>
</template>


<script setup>
import { ref, onMounted } from 'vue';
// import { getWalletInfo, repayLoan, applyLoan } from '@/service/loanService'; // 假设这些是API调用
import sHeader from '@/components/SimpleHeader.vue';
import {tryLoan,doLoan,payLoan,getUserFinance} from "@/service/user";
const walletBalance = ref(0);
const debt=ref(0);
const repaymentTerm = ref(60); // 默认60天
const loanAmount = ref('');
const selectedTerm = ref(30);
const isLoading = ref(false);
const loanResult = ref(null);
const repaymentAmount = ref('');
const tryGetLoan = async () => {
  isLoading.value = true; // 开始显示加载指示器
  try {
    const response = await tryLoan({  amount: parseInt(loanAmount.value, 10),
      term: parseInt(selectedTerm.value, 10), }); // 传递参数
    loanResult.value = response.data; // 假设后端响应的数据在response.data中
    console.log(loanResult.value); // 输出结果以供调试
    updateAomunt()
  } catch (error) {
    console.error('贷款尝试失败', error);
  } finally {
    isLoading.value = false; // 完成请求，隐藏加载指示器
  }
};


const updateAomunt = async () =>{
  try{
    const res=await getUserFinance();
    walletBalance.value=res.data.amount;
    repaymentTerm.value=res.data.term;
    debt.value=res.data.debt;
    console.log({res})
  }catch (error){
    console.error('获取钱包余额失败', error);
  }
}

const doLoanFunction = async () => {
  try {
    await doLoan({ amount: parseInt(loanAmount.value, 10), term: parseInt(selectedTerm.value, 10) });
    console.log("贷款操作完成"); // 根据实际情况展示操作结果
    await updateAomunt();
  } catch (error) {
    console.error('执行贷款操作失败', error);
    // 根据实际情况处理错误，比如显示错误消息
  }
};

onMounted(async () => {
  updateAomunt()
});




const handleRepayment = async () => {
  isLoading.value = true; // 显示加载指示器
  try {
    const response = await payLoan( parseInt(repaymentAmount.value, 10) );
    console.log('还款成功', response); // 根据实际情况输出日志或进行UI更新
    await updateAomunt();
  } catch (error) {
    console.error('还款失败', error); // 处理错误情况
  } finally {
    isLoading.value = false; // 请求结束，隐藏加载指示器
  }
};

</script>


<style scoped>
.wallet-info, .loan-application {
  margin-bottom: 20px;
  padding: 10px;
  border: 2px solid #eee;
  border-radius: 5px;
}

.loan-application {
  display: flex;
  flex-direction: column;
}

.van-button.center-button {
  display: block; /* 新增 */
  margin-top: 10px;
  margin-left: auto; /* 新增 */
  margin-right: auto; /* 新增 */
}

/* 新增钱包余额和贷款金额的颜色样式 */
.wallet-balance {
  color: green;
}

.debt {
  color: red;
}
</style>

