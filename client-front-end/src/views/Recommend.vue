<template>

  <div class="search-container">
    <s-header :name="'单品推荐'"></s-header>
    <div class="search-box">
      <input v-model="queryString" type="text" placeholder="请输入搜索内容" />
      <button @click="startVoiceRecognition">语音</button>
      <button @click="performSearch">完成</button>
    </div>
    <div v-if="searchResult" class="result-container">
      <div class="result-item" v-for="item in searchResult" :key="item.goodsId">
        <img :src="item.goodsCoverImg" alt="商品图片">
        <div class="item-info">
          <h5>{{ item.goodsName }}</h5>
          <p>价格: ¥{{ item.sellingPrice }}</p>
          <button @click="addToCart(item)">添加到购物车</button>
        </div>
      </div>
    </div>
    <nav-bar></nav-bar>
  </div>
</template>
<script>
import axios from 'axios';
import sHeader from '@/components/SimpleHeader.vue';
export default {
  data() {
    return {
      queryString: '',
      searchResult: null
    };
  },
  methods: {

    async performSearch() {
      try {
        const response = await axios.post('http://localhost:8000/api/v1/query', {
          query_string: this.queryString
        });
        this.searchResult = response.data;
      } catch (error) {
        console.error('搜索失败:', error);
        // 这里可以添加错误处理逻辑
      }
    },
    addToCart(item) {
      // 添加到购物车的逻辑
      console.log('添加到购物车:', item);
    },
    startVoiceRecognition() {
      // 语音识别逻辑（暂未实现）
    }
  }
};
</script>
<style>
.search-container {
  padding: 20px;
}

.search-box {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

.search-box input[type="text"] {
  flex-grow: 1;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
}
.back-button {
  padding: 10px 15px;
  border: none;
  background-color: #ff5722;
  color: white;
  border-radius: 4px;
  cursor: pointer;
  margin-bottom: 20px;
}
.search-box button {
  padding: 10px 15px;
  border: none;
  background-color: #007bff;
  color: white;
  border-radius: 4px;
  cursor: pointer;
}

.result-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.result-item {
  display: flex;
  gap: 20px;
  align-items: center;
}

.result-item img {
  width: 100px;
  height: 100px;
  object-fit: cover;
}

.item-info h5 {
  margin: 0;
  font-size: 18px;
}

.item-info p {
  margin: 5px 0;
}

.item-info button {
  padding: 5px 10px;
  border: none;
  background-color: #28a745;
  color: white;
  border-radius: 4px;
  cursor: pointer;
}
</style>
