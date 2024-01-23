<template>
  <div class="about">
    <s-header :name="'关于作者'"></s-header>
    <div class="about-body">
      <img class="intro-img" src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRFFYQkmI8yyOoPSvob9j-pZa0W6l-mz1ZblRUSzhOleoj-FCrdzZoQoWk6BNcTz83vHhM&usqp=CAU" alt="华东理工大学">
      <van-divider :style="{ color: '#1baeae', borderColor: '#1baeae', fontSize: '20px', fontWeight: 500 }">项目名称</van-divider>
      <div>基于大语言模型的电商金融分析系统</div>

      <van-divider :style="{ color: '#1baeae', borderColor: '#1baeae', fontSize: '20px', fontWeight: 500 }">作者信息</van-divider>
      <div>华东理工大学</div>
      <div>NP_123</div>
      <div>学号：20000000</div>

      <van-divider :style="{ color: '#1baeae', borderColor: '#1baeae', fontSize: '20px', fontWeight: 500 }">系统亮点</van-divider>
      <div>该系统的亮点在于通过大语言模型接受用户的输入，结合fine-tuning模型将输入拆分成一个个的所需商品信息，再通过词向量嵌入搜索出符合该商品信息的对应商品。最后，结合用户预算和商品价格等信息，通过机器学习算法推荐出商品列表。</div>
    </div>
  </div>
</template>

<script setup>
import sHeader from '@/components/SimpleHeader.vue';
</script>

<style lang="less" scoped>
.about {
  box-sizing: border-box;
  padding: 20px;
  .intro-img {
    width: 100%;
    max-width: 400px; /* 调整图片大小 */
    display: block;
    margin: auto; /* 图片居中显示 */
  }
  .about-body {
    font-size: 16px;
    a {
      color: #007fff;
    }
  }
}
</style>
