<template>
  <div class="layout">
    <el-container v-if="state.showMenu" class="container">
      <el-aside class="aside">
        <div class="head">
          <div>
            <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRFFYQkmI8yyOoPSvob9j-pZa0W6l-mz1ZblRUSzhOleoj-FCrdzZoQoWk6BNcTz83vHhM&usqp=CAU" alt="logo">
            <span>管理员后端</span>
          </div>
        </div>
        <div class="line" />
        <el-menu
          background-color="#222832"
          text-color="#fff"
          :router="true"
           :default-openeds="state.defaultOpen"
           :default-active='state.currentPath'
        >
          <el-sub-menu index="1">
            <template #title>
              <span>主要功能</span>
            </template>
            <el-menu-item-group>
              <el-menu-item index="/"><el-icon><Odometer /></el-icon>首页</el-menu-item>
              <el-menu-item index="/add"><el-icon><Plus /></el-icon>添加商品</el-menu-item>
            </el-menu-item-group>
          </el-sub-menu>
           <el-sub-menu index="2">
            <template #title>
              <span>首页配置</span>
            </template>
            <el-menu-item-group>
              <el-menu-item index="/swiper"><el-icon><Picture /></el-icon>轮播图配置</el-menu-item>
              <el-menu-item index="/hot"><el-icon><StarFilled /></el-icon>热销商品配置</el-menu-item>
              <el-menu-item index="/new"><el-icon><Sell /></el-icon>新品上线配置</el-menu-item>
              <el-menu-item index="/recommend"><el-icon><ShoppingCart /></el-icon>为你推荐配置</el-menu-item>
            </el-menu-item-group>
          </el-sub-menu>
          <el-sub-menu index="3">
            <template #title>
              <span>模块管理</span>
            </template>
            <el-menu-item-group>
              <el-menu-item index="/category"><el-icon><Menu /></el-icon>分类管理</el-menu-item>
              <el-menu-item index="/good"><el-icon><Goods /></el-icon>商品管理</el-menu-item>
              <el-menu-item index="/guest"><el-icon><User /></el-icon>会员管理</el-menu-item>
              <el-menu-item index="/order"><el-icon><List /></el-icon>订单管理</el-menu-item>
            </el-menu-item-group>
          </el-sub-menu>
          <el-sub-menu index="4">
            <template #title>
              <span>系统管理</span>
            </template>
            <el-menu-item-group>
              <el-menu-item index="/account"><el-icon><Lock /></el-icon>修改密码</el-menu-item>
            </el-menu-item-group>
          </el-sub-menu>
        </el-menu>
      </el-aside>
      <el-container class="content">
        <Header />
        <div class="main">
          <router-view />
        </div>
        <Footer />
      </el-container>
    </el-container>
    <el-container v-else class="container">
      <router-view />
    </el-container>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import { localGet, pathMap } from '@/utils'

const noMenu = ['/login']
const router = useRouter()
const state = reactive({
  showMenu: true,
  defaultOpen: ['1', '2', '3', '4'],
  currentPath: '/',
})

router.afterEach((to, from) => {
  state.showMenu = !noMenu.includes(to.path)
})

router.beforeEach((to, from, next) => {
  if (to.path == '/login') {
    // 如果路径是 /login 则正常执行
    next()
  } else {
    // 如果不是 /login，判断是否有 token
    if (!localGet('token')) {
      // 如果没有，则跳至登录页面
      next({ path: '/login' })
    } else {
      // 否则继续执行
      next()
    }
  }
  state.currentPath = to.path
  document.title = pathMap[to.name]
})
</script>

<style scoped>
.layout {
  min-height: 100vh;
  background-color: #f7f7f7; /* 浅灰背景色 */
}
.container {
  height: 100vh;
  display: flex; /* 使用 flex 布局 */
}
.aside {
  width: 250px!important; /* 增加侧边栏宽度 */
  background-color: #37474f; /* 深蓝灰色 */
  padding: 20px; /* 增加内边距 */
}
.head {
  display: flex;
  align-items: center;
  justify-content: space-between; /* 调整为两端对齐 */
  height: 60px; /* 增加头部高度 */
  background-color: #263238; /* 深灰色 */
}
.head > div {
  display: flex;
  align-items: center;
  padding: 0 15px; /* 增加内边距 */
}

.head img {
  width: 40px; /* 缩小图片大小 */
  height: 40px;
  border-radius: 20px; /* 圆形图片 */
}
.head span {
  font-size: 24px; /* 增加字体大小 */
  color: #FFC107; /* 金色字体 */
  font-weight: bold; /* 加粗字体 */
}
.line {
  border-top: 1px solid #eceff1; /* 调整颜色 */
  border-bottom: 1px solid #cfd8dc;
}
.content {
  display: flex;
  flex-direction: column;
  max-height: 100vh;
  overflow: hidden;
  flex-grow: 1; /* 填充剩余空间 */
}
.main {
  overflow: auto;
  padding: 20px; /* 增加内边距 */
  background-color: #eceff1; /* 调整背景色 */
}
</style>

<style>
body {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
  font-family: 'Roboto', sans-serif; /* 更改字体 */
}
.el-menu {
  border-right: none!important;
  background-color: #263238; /* 统一侧边栏颜色 */
}
.el-submenu {
  border-top: 1px solid #cfd8dc;
  border-bottom: 1px solid #b0bec5;
}
.el-submenu:first-child {
  border-top: none;
}
.el-submenu [class^="el-icon-"] {
  vertical-align: middle!important; /* 调整图标垂直对齐 */
}
a {
  color: #FFC107; /* 调整链接颜色 */
  text-decoration: none;
}
.el-pagination {
  text-align: center;
  margin-top: 20px;
}
.el-popper__arrow {
  display: none;
}
</style>

