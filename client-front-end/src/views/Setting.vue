
<template>
  <div class="seting-box">
    <s-header :name="'账号管理'"></s-header>
    <div class="input-item">
      <van-field v-model="state.nickName" label="昵称" />
      <van-field v-model="state.introduceSign" label="个性签名" />
      <van-field v-model="state.password" type='password' label="修改密码" />
      <van-field v-model="state.avatar" type='avatar' label="头像(必须):" />
    </div>
    <el-upload
        :on-change="handleelchange"
        :limit="1"
        :auto-upload="false"
        list-type="picture-card"
    >
      <i class="el-icon-plus"></i>
    </el-upload>

    <van-button round class="save-btn" color="#1baeae" type="primary" @click="save" block>保存</van-button>
    <van-button round class="save-btn" color="#1baeae" type="primary" @click="handleLogout" block>退出登录</van-button>
  </div>
</template>

<script setup>
import { reactive, onMounted } from 'vue'
import md5 from 'js-md5'
import sHeader from '@/components/SimpleHeader.vue'
import { getUserInfo, EditUserInfo, logout } from '@/service/user'
import { setLocal } from '@/common/js/utils'
import { showSuccessToast } from 'vant'


const state = reactive({
  nickName: '',
  introduceSign: '',
  password: '',
  uploadava: ''
})

const handleelchange = (file) => {
  const reader = new FileReader();
  reader.readAsDataURL(file.raw);
  reader.onload = () => {
    state.uploadava = reader.result;
  };
}

onMounted(async () => {
  const { data } = await getUserInfo()
  state.nickName = data.nickName
  state.introduceSign = data.introduceSign
})

const save = async () => {
  const params = {
    introduceSign: state.introduceSign,
    nickName: state.nickName
  }
  if (state.uploadava!==''){
    params.avatar=state.uploadava;
  }
  if (state.password) {
    params.passwordMd5 = md5(state.password)
  }
  await EditUserInfo(params)
  showSuccessToast('保存成功')
}

const handleLogout = async () => {
    setLocal('token', '')
    window.location.href = '/home'
}
</script>

<style lang="less" scoped>
  .seting-box {
    .save-btn {
      width: 80%;
      margin: 20px auto ;
    }
  }
</style>
