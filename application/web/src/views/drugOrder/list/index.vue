<template>
  <div class="container">
    <el-alert
      type="success"
    >
      <p>账户ID: {{ account_id }}</p>
      <p>用户名: {{ account_name }}</p>

    </el-alert>
    <div v-if="drugOrderList.length==0" style="text-align: center;">
      <el-alert
        title="查询不到数据"
        type="warning"
      />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in drugOrderList" :key="index" :span="6" :offset="1">
        <el-card class="drugOrder-card">
          <div slot="header" class="clearfix">
            药品订单ID:
            <span style="color: rgb(255, 0, 0);">{{ val.id }}</span>
          </div>

          <div class="item">
            <el-tag>病历ID: </el-tag>
            <span style="margin-left: 5px;">{{ val.prescription }}</span>
          </div>
          <div class="item">
            <el-tag type="success">病人ID: </el-tag>
            <span style="margin-left: 5px;">{{ val.patient }}</span>
          </div>
          <div class="item">
            <el-tag type="warning">药品名: </el-tag>
            <span style="margin-left: 5px;">{{ val.Name }}</span>
          </div>
          <div class="item">
            <el-tag type="warning">药品数量: </el-tag>
            <span style="margin-left: 5px;">{{ val.amount }} 份</span>
          </div>
          <div class="item">
            <el-tag type="info">创建时间: </el-tag>
            <span style="margin-left: 5px;">{{ val.created }}</span>
          </div>
        </el-card>
      </el-col>
    </el-row>

  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/accountV2'
import { queryDrugOrderList } from '@/api/drugOrder'

export default {
  name: 'DrugOrder',
  data() {
    return {
      loading: true,
      drugOrderList: [],
      accountList: [],
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
    if (this.roles[0] === 'admin') {
      queryDrugOrderList().then(response => {
        if (response !== null) {
          this.drugOrderList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    } else {
      queryDrugOrderList(/*{ patient: this.account_id }*/).then(response => {
        if (response !== null) {
          this.drugOrderList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    }
  },
  methods: {
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

  .drugOrder-card {
    width: 280px;
    height: 330px;
    margin: 18px;
  }
</style>
