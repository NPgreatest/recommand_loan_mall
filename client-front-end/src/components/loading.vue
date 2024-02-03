<template>
  <div class="page-loader" v-if="!isloaded">
    <div class="cube"></div>
    <div class="cube"></div>
    <div class="cube"></div>
    <div class="cube"></div>
  </div>
</template>

<script>
export default {
  data: () => {
    return {
      isloaded: false
    }
  },
  mounted() {
    document.onreadystatechange = () => {
      if (document.readyState == "complete") {
        this.isloaded = true;
      }
    }
  },
}
</script>

<style lang="scss" scoped>
$colors: #8CC271, #69BEEB, #F5AA39, #E9643B;

// -----------------------------------------------------

.page-loader {
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute; // 或者 "relative" 根据需要调整
  top: 50%; // 将加载器定位到父容器的中心
  left: 50%;
  width: auto; // 或指定一个宽度
  height: auto; // 或指定一个高度
  transform: translate(-50%, -50%); // 中心对齐
  background-color: rgba(255, 255, 255, 0.5); // 半透明背景
  z-index: 999;
}

// -----------------------------------------------------

.cube{
  width: 40px;
  height: 40px;
  margin-right: 10px;

  @for $i from 1 through length($colors) {
    &:nth-child(#{$i}) {
      background-color: nth($colors, $i);
    }
  }

  &:first-child {
    animation: left 1s infinite;
  }

  &:last-child {
    animation: right 1s infinite .5s;
  }
}

// -----------------------------------------------------

@keyframes left {
  40% {
    transform: translateX(-60px);
  }
  50% {
    transform: translateX(0);
  }
}

@keyframes right {
  40% {
    transform: translateX(60px);
  }
  50% {
    transform: translateX(0);
  }
}
</style>
