<template>
  <div class="footer">
    <div class="left">基于大语言模型的电商金融分析系统</div>
    <div class="right">
      <a target="_blank" href="https://github.com/NPgreatest/undergrad-thesis-proj">个人Github主页</a>
    </div>
  </div>
</template>

<script setup>
</script>

<style scoped>
.footer {
  height: 70px;
  background-color: #b3bed9; /* 更改背景颜色 */
  color: #e85d5d; /* 更改文本颜色 */
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 30px;
  box-shadow: 0 -4px 8px 0 rgba(110, 50, 50, 0.2); /* 添加阴影效果 */
  font-family: 'Arial', sans-serif; /* 更改字体 */
}

.footer .left {
  font-size: 18px; /* 左侧文本大小 */
  font-weight: bold; /* 字体加粗 */
}

.footer .right a {
  color: #000000; /* 更改链接颜色 */
  text-decoration: none; /* 去除下划线 */
  font-size: 16px; /* 链接字体大小 */
  font-weight: bold;
}

.footer .right a:hover {
  color: #02ec21; /* 鼠标悬停时的链接颜色 */
}
</style>
