<template>
  <div class="product-list-wrap">
    <div class="product-list-content">
      <header class="category-header wrap">
        <i class="nbicon nbfanhui" @click="goBack"></i>
        <div class="header-search">
          <i class="nbicon nbSearch"></i>
          <input
            type="text"
            class="search-title"
            v-model="state.keyword"/>
        </div>
        <span class="search-btn" @click="getSearch">搜索</span>
      </header>
      <div class="tabs-container">
        <div class="tab" v-for="(tab, index) in state.tabs" :key="index" @click="changeTab(index)">
          {{ tab.title }}
        </div>
      </div>
    </div>
    <div class="content">
      <van-pull-refresh v-model="state.refreshing" @refresh="onRefresh" class="product-list-refresh">
        <van-list
          v-model:loading="state.loading"
          :finished="state.finished"
          :finished-text="state.productList.length ? '没有更多了' : '搜索想要的商品'"
          @load="onLoad"
          @offset="10"
        >
          <!-- <p v-for="item in list" :key="item">{{ item }}</p> -->
          <template v-if="state.productList.length">
            <div class="product-item" v-for="(item, index) in state.productList" :key="index" @click="productDetail(item)">
              <img :src="$filters.prefix(item.goodsCoverImg)" />
              <div class="product-info">
                <p class="name">{{item.goodsName}}</p>
                <p class="subtitle">{{item.goodsIntro}}</p>
                <span class="price">￥ {{item.sellingPrice}}</span>
              </div>
            </div>
          </template>
          <div v-if="state.loading" class="loader"></div>
<!--          <img class="empty" v-else src="https://s.yezgea02.com/1604041313083/kesrtd.png" alt="搜索">-->
        </van-list>


      </van-pull-refresh>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { search } from '@/service/good'
const route = useRoute()
const router = useRouter()
const state = reactive({
  tabs: [
    { title: "推荐", name: "" },
    { title: "新品", name: "new" },
    { title: "价格", name: "price" },
    {title: "好评", name: ""}
  ],
  activeTabIndex: 0,
  keyword: route.query.keyword || '',
  searchBtn: false,
  seclectActive: false,
  refreshing: false,
  list: [],
  loading: false,
  finished: false,
  productList: [],
  totalPage: 0,
  page: 1,
  orderBy: ''
})
const init = async () => {
  const { categoryId } = route.query
  if (!categoryId && !state.keyword) {
    state.finished = true
    state.loading = false;
    return
  }
  const { data, data: { list } } = await search({ pageNumber: state.page, goodsCategoryId: categoryId, keyword: state.keyword, orderBy: state.orderBy })

  state.productList = state.productList.concat(list)
  state.totalPage = data.totalPage
  state.loading = false;
  if (state.page >= data.totalPage) state.finished = true
}

const goBack = () => {
  router.go(-1)
}
const goToPreviousPage = () => {
  if (state.page > 1) {
    state.page--;
    onRefresh();
  }
}

const goToNextPage = () => {
  // alert(state.totalPage)
  if (state.page < state.totalPage) {
    state.page++;
    onRefresh();
  }
}
const productDetail = (item) => {
  router.push({ path: `/product/${item.goodsId}` })
}

const getSearch = () => {
  onRefresh()
}

const onLoad = () => {
  if (!state.refreshing && state.page < state.totalPage) {
    state.page = state.page + 1
  }
  if (state.refreshing) {
    state.productList = [];
    state.refreshing = false;
  }
  init()
}

const onRefresh = () => {
  state.refreshing = true
  state.finished = false
  state.loading = true
  state.page = 1
  onLoad()
}
const changeTab = (index) => {
  state.activeTabIndex = index
  console.log('name', state.tabs[index])
  state.orderBy = state.tabs[index].name
  onRefresh()
}



</script>

<style lang="less" scoped>
  @import '../common/style/mixin';
  .loader {
    border: 4px solid #f3f3f3;
    border-top: 4px solid #1f38cb;
    border-radius: 50%;
    width: 40px;
    height: 40px;
    animation: spin 2s linear infinite;
    margin: 50px auto;
  }

  @keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }
  .tabs-container {
    display: flex;
    justify-content: center; /* 居中对齐 */
    background-color: #ffffff; /* 淡绿色背景 */
    padding: 10px;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    margin-bottom: 5px; /* 增加下边距以避免重叠 */
  }

  .tab {
    padding: 12px 16px; /* 增加内边距 */
    color: #000000;
    font-size: 18px; /* 增加字体大小 */
    font-weight: bold;
    cursor: pointer;
    border-radius: 5px;
    transition: background-color 0.3s, transform 0.2s;
    margin: 0 5px; /* 为标签之间添加一些间距 */
  }

  .tab:hover {
    background-color: #a5e3a5; /* 鼠标悬停时更深的绿色 */
    transform: scale(1.05); /* 缩放效果 */
  }

  .tab:nth-child(2) {
    /* 为第二个标签添加不同的样式 */
  }

  .product-list-content {
    position: fixed;
    left: 0;
    top: 0;
    width: 100%;
    z-index: 1000;
    background: #fff;
    .category-header {
      .fj();
      width: 100%;
      height: 50px;
      line-height: 50px;
      padding: 0 15px;
      .boxSizing();
      font-size: 15px;
      color: #656771;
      z-index: 10000;
      &.active {
        background: @primary;
      }
      .icon-left {
        font-size: 25px;
        font-weight: bold;
      }
      .header-search {
        display: flex;
        width: 76%;
        line-height: 20px;
        margin: 10px 0;
        padding: 5px 0;
        color: #232326;
        background: #F7F7F7;
        .borderRadius(20px);
        .nbSearch {
          padding: 0 5px 0 20px;
          font-size: 17px;
        }
        .search-title {
          font-size: 12px;
          color: #666;
          background: #F7F7F7;
        }
    }
    .icon-More {
      font-size: 20px;
    }
      .search-btn {
        height: 32px; /* 增加按钮高度 */
        margin: 8px 0;
        line-height: 32px;
        padding: 0 15px; /* 增加按钮内边距 */
        color: #fff;
        background: #1c671c; /* 浅绿色背景 */
        .borderRadius(15px); /* 增加圆角 */
        transition: background-color 0.3s, transform 0.1s; /* 添加颜色和变形过渡效果 */
        cursor: pointer; /* 鼠标悬停时手型指针 */

        &:hover {
          background: #1b791b; /* 鼠标悬停时的背景色变化 */
        }

        &:active {
          transform: scale(0.95); /* 点击时的缩放效果 */
        }
      }

  }
}
  .content {
    height: calc(~"(100vh - 70px)");
    overflow: hidden;
    overflow-y: scroll;
    margin-top: 128px;
  }
  .product-list-refresh {
    .product-item {
      .fj();
      width: 100%;
      height: 120px;
      padding: 10px 0;
      border-bottom: 1px solid #dcdcdc;
      img {
        width: 140px;
        height: 120px;
        padding: 0 10px;
        .boxSizing();
      }
      .product-info {
          width: 56%;
          height: 120px;
          padding: 5px;
          text-align: left;
          .boxSizing();
          p {
            margin: 0
          }
          .name {
            width: 100%;
            max-height: 40px;
            line-height: 20px;
            font-size: 15px;
            color: #333;
            overflow: hidden;
            text-overflow:ellipsis;
            white-space: nowrap;
          }
          .subtitle {
            width: 100%;
            max-height: 20px;
            padding: 10px 0;
            line-height: 25px;
            font-size: 13px;
            color: #999;
            overflow: hidden;
          }
          .price {
            color: @primary;
            font-size: 16px;
          }
      }
  }
  .empty {
    display: block;
    width: 150px;
    margin: 50px auto 20px;
  }

}
</style>
