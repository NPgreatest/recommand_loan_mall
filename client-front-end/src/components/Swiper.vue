<template>
  <van-swipe class="my-swipe" :autoplay="3000" indicator-color="#1baeae">
    <van-swipe-item v-for="(item, index) in list" :key="index" class="swipe-item">
      <img :src="item.carouselUrl" alt="" @click="goTo(item.redirectUrl)">
    </van-swipe-item>
  </van-swipe>
</template>

<script>
export default {
  props: {
    list: Array
  },
  methods: {
    goTo(url) {
      window.open(url)
    }
  }
}
</script>

<style lang='less' scoped>
.my-swipe {
  border-radius: 8px; /* 圆角边框 */
  overflow: hidden; /* 隐藏溢出的部分 */
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.15); /* 轻微的阴影效果 */

  .swipe-item {
    display: flex;
    justify-content: center;
    align-items: center;

    img {
      width: 100%;
      height: auto; /* 保持图片比例 */
      object-fit: cover; /* 图片填充整个容器 */
    }
  }
}
</style>
