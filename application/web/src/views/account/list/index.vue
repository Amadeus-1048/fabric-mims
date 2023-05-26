<template>
  <div class="container">
    <el-alert
      type="success"
    >
      <p>账户ID: {{ account_id }}</p>
      <p>用户名: {{ account_name }}</p>

    </el-alert>
    <div v-if="accountList.length==0" style="text-align: center;">
      <el-alert
        title="查询不到数据"
        type="warning"
      />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in accountList" :key="index" :span="6" :offset="1">
        <el-card class="prescription-card">
          <div slot="header" class="clearfix">
            账户ID:
            <span style="color: rgb(255, 0, 0);">{{ val.account_id }}</span>
          </div>

          <div class="item">
            <el-tag>账户名: </el-tag>
            <span style="margin-left: 5px;">{{ val.account_name }}</span>
          </div>


        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/accountV2'

export default {
  name: 'Account',
  data() {
    return {
      loading: true,
      loadingDialog: false,
      accountList: [],
      valItem: {}
    }
  },
  computed: {
    ...mapGetters([
      'account_id',
      'roles',
      'account_name',
    ])
  },
  created() {
    queryAccountList().then(response => {
        if (response !== null) {
          this.accountList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    }
}

</script>

<style>
  .container{
    width: 100%;
    /*text-align: center;*/
    min-height: 100%;
    overflow: hidden;
    font-size: 15px;
  }
  .tag {
    float: left;
  }

  .item {
    font-size: 14px;
    margin-bottom: 18px;
    color: #999;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
  }
  .clearfix:after {
    clear: both
  }

  .prescription-card {
    width: 280px;
    height: 150px;
    margin: 18px;
  }
</style>
