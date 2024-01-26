<template>
  <div class="product-detail">
    <s-header :name="'商品详情'"></s-header>
    <div class="detail-content">
      <div class="detail-swipe-wrap">
        <van-swipe class="my-swipe" indicator-color="#1baeae">
          <van-swipe-item v-for="(item, index) in state.detail.goodsCarouselList" :key="index">
            <img :src="item" alt="">
          </van-swipe-item>
        </van-swipe>
      </div>
      <div class="product-info">
        <div class="product-title">
          {{ state.detail.goodsName || '' }}
        </div>
        <div class="product-desc">免邮费 顺丰快递</div>
        <div class="product-price">
          <span>¥{{ state.detail.sellingPrice || '' }}</span>
          <!-- <span>库存203</span> -->
        </div>
      </div>
      <div class="product-intro">
        <ul>
          <li>详细信息</li>
        </ul>
        <div class="product-content" v-html="state.detail.goodsDetailContent || ''"></div>
        <ul>
          <li>评论</li>
        </ul>
      </div>


      <div class="product-reviews">
        <ul>
          <li v-for="(review, index) in state.reviews" :key="index">
            <div class="review-user">
              <img :src="$filters.prefix(review.avatar)" class="review-avatar" alt="User Avatar">
              <div class="review-nickname">{{ review.nickName }}</div>
            </div>


            <div class="review-title">评论标题: {{ review.reviewTitle }}</div>
            <div class="review-content">评论内容: {{ review.reviewContent }}</div>
            <div class="review-star">
              <span v-for="star in 5" :key="star" class="star" v-bind:class="{ filled: star <= review.reviewStar }"></span>
            </div>
            <div class="review-time">评论时间: {{ formatDate(review.reviewTime) }}</div>
          </li>
        </ul>
      </div>
    </div>
    <van-action-bar>
      <van-action-bar-icon icon="chat-o" text="客服" />
      <van-action-bar-icon icon="cart-o" :badge="!cart.count ? '' : cart.count" @click="goTo()" text="购物车" />
      <van-action-bar-button type="warning" @click="handleAddCart" text="加入购物车" />
      <van-action-bar-button type="danger" @click="goToCart" text="立即购买" />
    </van-action-bar>
  </div>
</template>

<script setup>
import { reactive, onMounted, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useCartStore } from '@/stores/cart'
import {getDetail, getReview} from '@/service/good'
import { addCart } from '@/service/cart'
import sHeader from '@/components/SimpleHeader.vue'
import { showSuccessToast } from 'vant'
import { prefix } from '@/common/js/utils'
const route = useRoute()
const router = useRouter()
const cart = useCartStore()

const state = reactive({
  detail: {
    goodsCarouselList: []
  },
  reviews: []
})

onMounted(async () => {
  const { id } = route.params
  const { data } = await getDetail(id)
  data.goodsCarouselList = data.goodsCarouselList.map(i => prefix(i))
  state.detail = data
  cart.updateCart()
  await fetchReviews()
})

const fetchReviews = async (pageNumber = 1) => {
  const { id } = route.params
  const { data } = await getReview(id, { pageNumber })
  state.reviews = data || []
}

const formatDate = (timestamp) => {
  const date = new Date(timestamp * 1000)
  return `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}`
}

nextTick(() => {
  // 一些和DOM有关的东西
  const content = document.querySelector('.detail-content')
  content.scrollTop = 0
})

const goBack = () => {
  router.go(-1)
}

const goTo = () => {
  router.push({ path: '/cart' })
}

const handleAddCart = async () => {
  const { resultCode } = await addCart({ goodsCount: 1, goodsId: state.detail.goodsId })
  if (resultCode == 200 ) showSuccessToast('添加成功')
  cart.updateCart()
}

const goToCart = async () => {
  await addCart({ goodsCount: 1, goodsId: state.detail.goodsId })
  cart.updateCart()
  router.push({ path: '/cart' })
}

</script>

<style lang="less">
  @import '../common/style/mixin';
  .product-detail {
    .detail-header {
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
      .product-name {
        font-size: 14px;
      }
    }
    .detail-content {
      height: calc(100vh - 50px);
      overflow: hidden;
      overflow-y: auto;
      .detail-swipe-wrap {
        .my-swipe .van-swipe-item {
          img {
            width: 100%;
            // height: 300px;
          }
        }
      }
      .product-info {
        padding: 0 10px;
        .product-title {
          font-size: 18px;
          text-align: left;
          color: #333;
        }
        .product-desc {
          font-size: 14px;
          text-align: left;
          color: #999;
          padding: 5px 0;
        }
        .product-price {
          .fj();
          span:nth-child(1) {
            color: #F63515;
            font-size: 22px;
          }
          span:nth-child(2) {
            color: #999;
            font-size: 16px;
          }
        }
      }
      .product-intro {
        width: 100%;
        padding-bottom: 50px;
        ul {
          .fj();
          width: 100%;
          margin: 10px 0;
          li {
            flex: 1;
            padding: 5px 0;
            text-align: center;
            font-size: 15px;
            border-right: 1px solid #999;
            box-sizing: border-box;
            &:last-child {
              border-right: none;
            }
          }
        }
        .product-content {
          padding: 0 20px;
          img {
            width: 100%;
          }
        }
      }
    }
    .van-action-bar-button--warning {
      background: linear-gradient(to right,#6bd8d8, @primary)
    }
    .van-action-bar-button--danger {
      background: linear-gradient(to right, #0dc3c3, #098888)
    }
  }


  .product-reviews ul {
    list-style-type: none;
    padding: 0;
    margin: 0;
  }

  .product-reviews li {
    padding: 10px 0;
    border-bottom: 1px solid #ddd; /* 每个评论之间的横线 */
  }

  .review-title, .review-content {
    margin-top: 5px;
  }

  .review-star {
    display: inline-block;
    margin: 5px 0px;
    left: 33% ;
  }

  .star {
    display: inline-block;
    width: 26px;
    height: 26px;
    background-image: url('../assets/empty_star.png');
    background-size: cover;
  }

  .star.filled {
    background-image: url('../assets/full_star.png');
  }

  .review-time {
    text-align: right;
    font-size: 0.8em;
    color: #777;
    margin-top: 5px;
  }
  .review-avatar {
    width: 50px;
    height: 50px;
    border-radius: 25px;
  }
  .review-user {
    display: flex;
    align-items: center;
    font-size: 15px;
  }

  .review-avatar {
    width: 50px;
    height: 50px;
    border-radius: 25px;
    margin-right: 10px; /* 在头像和昵称之间添加一些间距 */
  }

  .review-nickname {
    font-size: 1.2em; /* 增大字体大小 */
    color: black; /* 将字体颜色设置为纯黑色 */
  }

</style>
