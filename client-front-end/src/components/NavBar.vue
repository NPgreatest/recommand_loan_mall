<template>
  <div class="nav-bar van-hairline--top">
    <ul class="nav-list">
      <router-link class="nav-list-item" to="home">
        <i class="nbicon nblvsefenkaicankaoxianban-1"></i>
        <span>首页</span>
      </router-link>
      <router-link class="nav-list-item" to="category">
        <i class="nbicon nbfenlei"></i>
        <span>分类</span>
      </router-link>
      <router-link class="nav-list-item" to="cart">
        <i><van-icon name="shopping-cart-o" :badge="!cart.count ? '' : cart.count" /></i>
        <span>购物车</span>
      </router-link>
      <router-link class="nav-list-item" to="user">
        <i class="nbicon nblvsefenkaicankaoxianban-"></i>
        <span>我的</span>
      </router-link>
    </ul>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useCartStore } from '@/stores/cart'
import { getLocal } from '@/common/js/utils'

const route = useRoute()
const cart = useCartStore()

onMounted(() => {
  const token = getLocal('token')
  const path = route.path
  if (token && !['/home', '/category'].includes(path)) {
    cart.updateCart()
  }
})
</script>

<style lang="less" scoped >
@import '../common/style/mixin';

.nav-bar {
  position: fixed;
  left: 0;
  bottom: 0;
  width: 100%;
  padding: 0px 0;
  z-index: 1000;
  background: #ffffff; /* 深色背景 */
  color: #000000; /* 白色文字 */
  box-shadow: 0 -2px 5px rgba(0,0,0,0.2); /* 添加阴影 */
}

.nav-list {
  width: 100%;
  .fj();
  flex-direction: row;
  padding: 0;
  margin: 0;
  list-style: none; /* 移除列表样式 */

  .nav-list-item {
    display: flex;
    flex: 1;
    flex-direction: column;
    text-align: center;
    align-items: center; /* 水平居中 */
    justify-content: center; /* 垂直居中 */
    color: #b0bec5; /* 更淡的字体颜色 */
    transition: color 0.3s; /* 平滑颜色变换 */

    &.router-link-active {
      color: #0029ff; /* 激活状态下的颜色 */
    }

    i {
      text-align: center;
      font-size: 26px; /* 增加图标大小 */
      margin-bottom: 5px; /* 增加图标与文字间距 */
    }

    span {
      font-size: 14px; /* 增加字体大小 */
    }

    .van-icon-shopping-cart-o {
      margin: 0 auto;
      color: #fa6181; /* 购物车图标颜色 */
    }
  }
}
</style>
