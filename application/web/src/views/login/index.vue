<template>
  <div class="login-container">
    <el-form ref="loginForm" class="login-form" auto-complete="on" label-position="left">

      <div class="title-container">
        <h3 class="title">基于区块链的医疗信息管理系统</h3>
      </div>
      <el-select v-model="value" placeholder="请选择用户角色" class="login-select" @change="selectGet">
        <el-option
          v-for="item in accountList"
          :key="item.account_id"
          :label="item.account_name"
          :value="item.account_id"
        >
          <span style="float: left">{{ item.account_name }}</span>
          <span style="float: right; color: #8492a6; font-size: 13px">{{ item.account_id }}</span>
        </el-option>
      </el-select>

      <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;" @click.native.prevent="handleLogin">立即进入</el-button>

      <div class="tips">
        <span style="margin-right:20px;">请选择不同的用户角色体验系统</span>
      </div>

    </el-form>
  </div>
</template>

<script>
import { queryAccountList } from '@/api/accountV2'

export default {
  name: 'Login',
  data() {
    return {
      loading: false,
      redirect: undefined,
      accountList: [],
      value: ''
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  created() {
    queryAccountList().then(response => {
      if (response !== null) {
        this.accountList = response
      }
    })
  },
  methods: {
    handleLogin() {
      if (this.value) {
        this.loading = true
        this.$store.dispatch('account/login', this.value).then(() => {
          this.$router.push({ path: '/' })
          this.loading = false
        }).catch(() => {
          this.loading = false
        })
      } else {
        this.$message('请选择用户角色')
      }
    },
    selectGet(account_id) {
      this.value = account_id
    }
  }
}
</script>

<style lang="scss" scoped>
$bg:#2d3a4b;
$dark_gray:#889aa4;
$light_gray:#eee;

.login-container {
  min-height: 100%;
  width: 100%;
  background-color: $bg;
  overflow: hidden;

  .login-form {
    position: relative;
    width: 520px;
    max-width: 100%;
    padding: 160px 35px 0;
    margin: 0 auto;
    overflow: hidden;
  }
  .login-select{
   padding: 20px 0px 30px 0px;
   min-height: 100%;
   width: 100%;
   background-color: $bg;
   overflow: hidden;
   text-align: center;
  }
  .tips {
    font-size: 14px;
    color: #fff;
    margin-bottom: 10px;

    span {
      &:first-of-type {
        margin-right: 16px;
      }
    }
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $dark_gray;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }

  .title-container {
    position: relative;

    .title {
      font-size: 26px;
      color: $light_gray;
      margin: 0px auto 40px auto;
      text-align: center;
      font-weight: bold;
    }
  }

  .show-pwd {
    position: absolute;
    right: 10px;
    top: 7px;
    font-size: 16px;
    color: $dark_gray;
    cursor: pointer;
    user-select: none;
  }
}
</style>
