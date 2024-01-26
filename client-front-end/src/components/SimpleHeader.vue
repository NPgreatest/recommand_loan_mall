<template>
  <header class="simple-header van-hairline--bottom">
    <i v-if="!isback" class="nbicon nbfanhui" @click="goBack"></i>
    <i v-else>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</i>
    <div class="simple-header-name">{{ name }}</div>
    <i class="nbicon nbmore"></i>
  </header>
  <div class="block" />
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
export default {
  props: {
    name: {
      type: String,
      default: ''
    },
    back: {
      type: String,
      default: ''
    },
    noback: {
      type: Boolean,
      default: false
    }
  },
  emits: ['callback'],
  setup(props, ctx) {
    const isback = ref(props.noback)
    const router = useRouter()
    const goBack = () => {
      if (!props.back) {
        router.go(-1)
      } else {
        router.push({ path: props.back })
      }
      ctx.emit('callback')
    }
    return {
      goBack,
      isback
    }
  }
}
</script>

<style lang="less" scoped>
@import '../common/style/mixin';
.simple-header {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 10000;
  .fj();
  .wh(100%, 50px); /* 增加高度 */
  line-height: 50px;
  padding: 0 15px;
  .boxSizing();
  color: #606060; /* 深灰色文字 */
  background: #f0f0f0; /* 浅灰色背景 */
  box-shadow: 0 2px 5px rgba(0,0,0,0.1); /* 添加阴影效果 */
  transition: background-color 0.3s, color 0.3s; /* 平滑过渡动画 */

  .simple-header-name {
    font-size: 18px; /* 增加字体大小 */
    font-weight: bold; /* 加粗字体 */
    transition: font-size 0.3s; /* 字体大小变化动画 */
  }

  &:hover {
    background: #e0e0e0; /* 悬停时背景颜色变化 */
    color: #333; /* 悬停时文字颜色变化 */
    .simple-header-name {
      font-size: 20px; /* 悬停时字体大小变化 */
    }
  }
}
.block {
  height: 50px;
}
</style>
