<template >
  <div class="login" id="HomePage"  :style="{ backgroundImage: 'url(' + state.backgroundImage + ')' }">
    <s-header :name="state.type == 'login' ? '登录' : '注册'" :back="'/home'"></s-header>
    <img class="logo" :src="`icon/ECUST.webp`" alt="">
    <div v-if="state.type == 'login'" class="login-body login">
      <van-cell-group inset >
        <van-form @submit="onSubmit">
          <van-field
              center
              left-icon="icon/username.png"
              right-icon="warning-o"
              v-model="state.username"
              name="username"
              label="用户名"
              label-width='50'
              placeholder="用户名"
              :rules="[{ required: true, message: '请填写用户名' }]"
          />
          <van-field
              label-width='50'
              left-icon ="icon/password.png"
              center
              id='password'
              v-model="state.password"
              type="password"
              name="password"
              label="密码"
              placeholder="密码"
              @click="changeBackground_S"
              :rules="[{ required: true, message: '请填写密码' }]">
            <template #button >
              <span class="solts" @click="switchPasswordType">
                <van-icon name="eye" v-if="state.see" />
                <van-icon name="closed-eye" v-else />
              </span>
            </template>
          </van-field>
          <van-field
              label-width='50'
              left-icon="icon/verify.jpg"
              center
              label="验证码"
              placeholder="验证码"
              v-model="state.verify"
          >
            <template #button>
              <vue-img-verify ref="verifyRef"/>
            </template>

          </van-field>
          <div style="margin: 8px;">
            <div class="link-register" @click="toggle('register')">还没有账号？立即注册</div>
            <van-button round block
                        icon="https://fastly.jsdelivr.net/npm/@vant/assets/user-active.png"
                        color="linear-gradient(to right, #1baeae, #ff6f00)"
                        native-type="submit">登录</van-button>
          </div>
        </van-form>
      </van-cell-group>


    </div>
    <div v-else class="login-body register">
      <van-cell-group inset size="large">
        <van-form @submit="onSubmit">
          <van-field
              left-icon="icon/username.png"
              right-icon="warning-o"
              label-width='50'
              center
              v-model="state.username1"
              name="username1"
              label="用户名"
              placeholder="用户名"
              :rules="[{ required: true, message: '请填写用户名' }]"
          />
          <van-field
              left-icon="icon/password.png"
              label-width='50'
              center
              v-model="state.password1"
              type="password"
              name="password1"
              id="password1"
              label="密码"
              placeholder="密码"

              :rules="[{ required: true, message: '请填写密码' }]">
            <template #button >
              <span class="solts" @click="switchPasswordType1">
                <van-icon name="eye" v-if="state.see" />
                <van-icon name="closed-eye" v-else />
              </span>
            </template>
          </van-field>
          <van-field
              left-icon="icon/verify.jpg"
              label-width='50'
              center
              clearable
              label="验证码"
              placeholder="验证码"
              v-model="state.verify"
          >
            <template #button>
              <vue-img-verify ref="verifyRef" />
            </template>
          </van-field>
          <div style="margin: 8px;">
            <div class="link-login"  @click="toggle('login')"  >已有账号，直接登录</div>
            <van-button round block
                        icon="https://fastly.jsdelivr.net/npm/@vant/assets/user-active.png"
                        color="linear-gradient(to right, #1baeae, #ff6f00)" native-type="submit">注册</van-button>
          </div>
        </van-form>
      </van-cell-group>

    </div>
    <div> </div>

    <div class="placeholder-container">
    <div class="placeholder"></div>
  </div>


  </div>
</template>



<script setup>
import { reactive, ref } from 'vue'
import sHeader from '@/components/SimpleHeader.vue'
import vueImgVerify from '@/components/VueImageVerify.vue'
import { login, register } from '@/service/user'
import { setLocal } from '@/common/js/utils'
import md5 from 'js-md5'
import { showSuccessToast, showFailToast } from 'vant'
/*通过Vue的Composition API创建一个响应式引用，初始值为null。这个引用可以用来访问vue-img-verify组件。
* ref函数接收一个参数，即初始值，并返回一个包含两个属性的对象：.value和.valueOf()。*/
const verifyRef = ref(null)
const state = reactive({
  username: '',
  password: '',
  username1: '',
  password1: '',
  type: 'login',
  imgCode: '',
  verify: '',
  see:0,
  backgroundImagePaths: [
    'icon/background_login.jpg',
    'icon/background_login_pass.jpg',
    // 添加更多背景图片路径
  ],
  currentBackgroundIndex: 0, // 当前显示的背景图片索引
  backgroundImage:'icon/background_login.jpg'
})

const switchPasswordType =() => {
  state.see=!state.see
  if (state.see){
    document.getElementById('password').type='text'
  }
  else {
    document.getElementById('password').type='password'
  }
}
const switchPasswordType1 =() => {
  state.see=!state.see
  if (state.see){
    document.getElementById('password1').type='text'
  }
  else {
    document.getElementById('password1').type='password'
  }
}

const toggle = (v) => {
  state.type = v
  state.verify = ''
  state.currentBackgroundIndex = (state.currentBackgroundIndex + 1) % state.backgroundImagePaths.length;
  state.backgroundImage = state.backgroundImagePaths[state.currentBackgroundIndex];
}

const onSubmit = async (values) => {
  state.imgCode = verifyRef.value.state.imgCode || ''
  if (state.verify.toLowerCase() != state.imgCode.toLowerCase()) {
    showFailToast('验证码有误')
    return
  }

  if (state.type == 'login') {
    const { data } = await login({
      "loginName": values.username,
      "passwordMd5": md5(values.password)
    })
    setLocal('token', data)
    window.location.href = '/'
  } else {
    await register({
      "loginName": values.username1,
      "password": values.password1
    })
    showSuccessToast('注册成功')
    state.type = 'login'
    state.verify = ''
  }
}
</script>






<style lang="less">
.placeholder-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.placeholder {
  width: 100px;
  height: 100px;
  /* 其他样式 */
}

.login {
  .logo {
    width: 120px;
    height: 120px;
    display: block;
    margin: 80px auto 20px;
  }

  .login-body {
    padding: 0 20px;
    margin-bottom: 200px;
  }
  .login {
    .link-register {
      font-size: 14px;
      margin-bottom: 20px;
      color: #1989fa;
      display: inline-block;
    }
  }
  .register {
    .link-login {
      font-size: 14px;
      margin-bottom: 20px;
      color: #1989fa;
      display: inline-block;
    }
  }
  .verify-bar-area {
    margin-top: 24px;
    .verify-left-bar {
      border-color: #1baeae;
    }
    .verify-move-block {
      background-color: #1baeae;
      color: #fff;
    }
  }
  .verify {
    >div {
      width: 100%;
    }
    display: flex;
    justify-content: center;
    .cerify-code-panel {
      margin-top: 16px;
    }
    .verify-code {
      width: 40%!important;
      float: left!important;
    }
    .verify-code-area {
      float: left!important;
      width: 54%!important;
      margin-left: 14px!important;
      .varify-input-code {
        width: 90px;
        height: 38px!important;
        border: 1px solid #e9e9e9;
        padding-left: 10px;
        font-size: 16px;
      }
      .verify-change-area {
        line-height: 44px;
      }
    }
  }
}
</style>
