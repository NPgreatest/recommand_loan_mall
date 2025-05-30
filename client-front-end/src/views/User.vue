<template>
  <div class="user-box">
    <s-header :name="'我的'"></s-header>
    <van-skeleton title :avatar="true" :row="3" :loading="state.loading">
      <div class="user-info">
        <div class="info">
          <img :src="$filters.prefix(state.user.avatar)" alt="User Avatar"/>
          <div class="user-desc">
            <span>昵称：{{ state.user.nickName }}</span>
            <span>登录名：{{ state.user.loginName }}</span>
            <span class="name">个性签名：{{ state.user.introduceSign }}</span>
          </div>
        </div>
      </div>
    </van-skeleton>
    <ul class="user-list">
      <li class="van-hairline--bottom" @click="goTo('/order')">
        <van-icon name="orders-o" class="list-icon" />
        <span>我的订单</span>
        <van-icon name="arrow" />
      </li>
      <li class="van-hairline--bottom" @click="goTo('/setting')">
        <van-icon name="setting-o" class="list-icon" />
        <span>账号管理</span>
        <van-icon name="arrow" />
      </li>
      <li class="van-hairline--bottom" @click="goTo('/address', { from: 'mine' })">
        <van-icon name="location-o" class="list-icon" />
        <span>地址管理</span>
        <van-icon name="arrow" />
      </li>
      <li class="van-hairline--bottom" @click="goTo('/budget')">
        <van-icon name="chart-trending-o" class="list-icon" />
        <span>预算管理</span>
        <van-icon name="arrow" />
      </li>
      <li class="van-hairline--bottom" @click="goTo('/finance')">
        <van-icon name="gold-coin-o" class="list-icon" />
        <span>电商金融行为分析</span>
        <van-icon name="arrow" />
      </li>
      <li @click="goTo('/about')">
        <van-icon name="info-o" class="list-icon" />
        <span>关于我们</span>
        <van-icon name="arrow" />
      </li>
    </ul>
    <nav-bar></nav-bar>
  </div>
</template>


<script setup>
import { reactive, onMounted, toRefs } from 'vue'
import navBar from '@/components/NavBar.vue'
import sHeader from '@/components/SimpleHeader.vue'
import { getUserInfo } from '@/service/user'
import { useRouter } from 'vue-router'
import axios from "axios";
import * as filters from "@/common/js/utils";
const router = useRouter()
const state = reactive({
  user: {},
  loading: true
})

onMounted(async () => {
  const { data } = await getUserInfo()
  state.user = data
  state.loading = false
})

const goTo = (r, query) => {
  router.push({ path: r, query: query || {} })
}
</script>

<style lang="less" scoped>
  @import '../common/style/mixin';
  .user-box {
    .user-header {
      position: fixed;
      top: 0;
      left: 0;
      z-index: 10000;
      .fj();
      .wh(100%, 44px);
      line-height: 44px;
      padding: 0 10px;
      .boxSizing();
      color: #252525;
      background: #fff;
      border-bottom: 1px solid #dcdcdc;
      .user-name {
        font-size: 14px;
      }
    }
    .user-info {
      width: 94%;
      margin: 10px;
      height: 115px;
      background: linear-gradient(90deg, @primary, #1a9b49);
      box-shadow: 0 2px 5px #6cd392;
      border-radius: 6px;
      .info {
        position: relative;
        display: flex;
        width: 100%;
        height: 100%;
        padding: 25px 20px;
        .boxSizing();
        img {
          .wh(60px, 60px);
          border-radius: 50%;
          margin-top: 4px;
        }
        .user-desc {
          display: flex;
          flex-direction: column;
          margin-left: 10px;
          line-height: 20px;
          font-size: 14px;
          color: #fff;
          span {
            color: #fff;
            font-size: 14px;
            padding: 2px 0;
          }
        }
        .account-setting {
          position: absolute;
          top: 10px;
          right: 20px;
          font-size: 13px;
          color: #fff;
          .van-icon-setting-o {
            font-size: 16px;
            vertical-align: -3px;
            margin-right: 4px;
          }
        }
      }
    }
    .user-list {
      padding: 0 30px;
      margin-top: 40px;
      li {
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: 10px 0; /* 增加垂直内边距 */
        margin-bottom: 10px; /* 增加列表项之间的间距 */
        background-color: #f5f5f5; /* 浅灰色背景 */
        border-radius: 8px; /* 圆角边框 */
        transition: background-color 0.3s ease; /* 平滑背景颜色过渡 */
        cursor: pointer; /* 鼠标悬停时的手型指针 */

        &:hover {
          background-color: #e0e0e0; /* 深灰色背景 */
        }

        .list-icon {
          margin-right: 10px; /* 调整图标与文本之间的间距 */
          font-size: 20px; /* 增加图标大小 */
        }
        span {
          font-size: 16px; /* 增加字体大小 */
          color: #333; /* 深色字体 */
        }
        .van-icon-arrow {
          font-size: 20px; /* 增加箭头图标大小 */
          color: #999; /* 箭头图标的颜色 */
        }
      }
    }
  }
</style>
