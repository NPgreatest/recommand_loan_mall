<template>
  <div class="setting-box">
    <s-header :name="'财务管理'"></s-header>
    <div class="input-item">
      <van-field v-model.number="financeData.monthlyIncome" label="每月收入" type="number" />
      <van-field v-model.number="financeData.monthlyExpenses" label="每月支出" type="number" />
      <van-field v-model.number="financeData.creditScore" label="信用评分" type="number" />
      <van-field v-model.number="financeData.debtStatus" label="债务状况" type="number" />
    </div>
    <van-button round class="save-btn" color="#1baeae" type="primary" @click="saveFinanceData" block>保存</van-button>
  </div>
</template>

<script setup>
import { reactive, onMounted } from 'vue'
import sHeader from '@/components/SimpleHeader.vue'
import { showSuccessToast } from 'vant'
import axios from 'axios'
import {EditUserInfo, getUserInfo} from "@/service/user";
import md5 from "js-md5";

const financeData = reactive({
  monthlyIncome: 0,
  monthlyExpenses: 0,
  creditScore: 0,
  debtStatus: 0
})


onMounted(async () => {
  const { data } = await getUserFinance()
  financeData.monthlyIncome = data.monthlyIncome
  financeData.monthlyExpenses = data.monthlyExpenses
  financeData.creditScore = data.creditScore
  financeData.debtStatus = data.debtStatus
})


const saveFinanceData = async () => {
  const params = {
    monthlyIncome:  financeData.monthlyIncome,
    monthlyExpenses:financeData.monthlyExpenses,
    creditScore:financeData.creditScore,
    debtStatus:financeData.debtStatus
  }
  await EditUserInfo(params)
  showSuccessToast('保存成功')
}


</script>
