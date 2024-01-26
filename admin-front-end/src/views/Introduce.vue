<template>
  <div class="about">
    <div class="about-body">
      <img class="intro-img" src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRFFYQkmI8yyOoPSvob9j-pZa0W6l-mz1ZblRUSzhOleoj-FCrdzZoQoWk6BNcTz83vHhM&usqp=CAU" alt="华东理工大学">

      <van-divider class="custom-divider"> </van-divider>
      <div class="content">基于大语言模型的电商金融分析系统</div>

      <van-divider class="custom-divider">作者信息</van-divider>
      <div class="content">华东理工大学</div>
      <div class="content">NP_123</div>
      <div class="content">学号：20000000</div>

      <van-divider class="custom-divider">技术栈</van-divider>
      <div class="content">Golang, Mysql, Postgresql, embedding, fine-tuning, Vue.js, Python, GPT-3.5</div>

      <van-divider class="custom-divider">系统亮点</van-divider>
      <div class="content">该系统的亮点在于通过大语言模型接受用户的输入，结合fine-tuning模型将输入拆分成一个个的所需商品信息，再通过词向量嵌入搜索出符合该商品信息的对应商品。最后，结合用户预算和商品价格等信息，通过机器学习算法推荐出商品列表。</div>
    </div>
  </div>
</template>

<script setup>

</script>

<style scoped>
.about-body {
  max-width: 800px;
  margin: auto;
  padding: 20px;
  text-align: center;
}

.intro-img {
  width: 30%;
  height: auto;
  border-radius: 10px;
  margin-bottom: 20px;
}

.custom-divider {
  color: #1baeae;
  border-color: #1baeae;
  font-size: 20px;
  font-weight: 600;
}

.content {
  margin-bottom: 15px;
  font-size: 20px;
  color: #333;
}
</style>
