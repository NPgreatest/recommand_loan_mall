<template>
  <div class="setting-box">
    <s-header :name="'财务管理'"></s-header>
    <div class="input-item">
      <div class="slider-container">
        <div class="slider-title">每月收入</div>
        <van-slider v-model="financeData.monthlyIncome" :max="40000" :step="1000" />
        <div class="credit-score">备注：每月收入多少￥(元)</div>
        <van-field v-model.number="financeData.monthlyIncome" type="number" />
      </div>
      <div class="slider-container">
        <div class="slider-title">每月支出</div>
        <van-slider v-model="financeData.monthlyExpenses" :max="40000" :step="1000" />
        <div class="credit-score">备注：每月支出多少￥(元)</div>
        <van-field v-model.number="financeData.monthlyExpenses" type="number" />
      </div>
      <div class="slider-container">
        <div class="slider-title">债务状况</div>
        <van-slider v-model="financeData.debtStatus" :max="100000" :step="1000" />
        <div class="credit-score">备注：当前负债多少￥(元)</div>
        <van-field v-model.number="financeData.debtStatus" type="number" />
      </div>
      <div class="slider-container">
        <div class="slider-title">信用评分</div>
        <van-slider v-model="financeData.creditScore" :max="1000" :step="1" />
        <div class="credit-score">当前评分: {{ financeData.creditScore }}</div>
        <div class="credit-score-hint">（微信支付信用分）</div>
      </div>
    </div>
    <van-button round class="save-btn" color="#1baeae" type="primary" @click="saveFinanceData" block>保存</van-button>
  </div>
</template>

<script setup>
import { reactive, onMounted } from 'vue'
import sHeader from '@/components/SimpleHeader.vue'
import { showSuccessToast } from 'vant'
import { getUserFinance, setUserFinance } from "@/service/user"

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
  await setUserFinance(financeData)
  showSuccessToast('保存成功')
}
</script>

<style scoped>
.slider-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 20px;
}

.slider-title {
  margin-bottom: 10px;
  font-weight: bold;
}

.credit-score, .credit-score-hint {
  text-align: center;
  margin-top: 5px;
}
</style>
