<template>
  <div class="container">
    <el-alert
      type="success"
    >
      <p>账户ID: {{ account_id }}</p>
      <p>用户名: {{ account_name }}</p>

    </el-alert>
    <div v-if="prescriptionList.length==0" style="text-align: center;">
      <el-alert
        title="查询不到数据"
        type="warning"
      />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in prescriptionList" :key="index" :span="6" :offset="1">
        <el-card class="prescription-card">
          <div slot="header" class="clearfix">
            病历ID:
            <span style="color: rgb(255, 0, 0);">{{ val.id }}</span>
          </div>

          <div class="item">
            <el-tag>病人ID: </el-tag>
            <span style="margin-left: 5px;">{{ val.patient }}</span>
          </div>
          <div class="item">
            <el-tag type="success">医生ID: </el-tag>
            <span style="margin-left: 5px;">{{ val.doctor }}</span>
          </div>
          <div class="item">
            <el-tag type="warning">诊断: </el-tag>
            <span style="margin-left: 5px;">{{ val.diagnosis }}</span>
          </div>
          <div class="item">
            <el-tag type="danger">药品: </el-tag>
            <span style="margin-left: 5px;">
    <span v-for="(drug, index) in val.drug" :key="index">{{ drug.Name }} {{drug.amount}}份 <br v-if="index !== val.drug.length - 1"></span>
  </span>
          </div>
          <div class="item">
            <el-tag type="danger">备注: </el-tag>
            <span style="margin-left: 5px;">{{ val.comment }}</span>
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
import { queryPrescriptionList } from '@/api/prescription'

export default {
  name: 'Prescription',
  data() {
    return {
      loading: true,
      loadingDialog: false,
      prescriptionList: [],
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
    if (this.roles[0] === 'admin' || this.roles[0] === 'doctor') {
      queryPrescriptionList().then(response => {
        if (response !== null) {
          this.prescriptionList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    } else {
      queryPrescriptionList({ patient: this.account_id }).then(response => {
        if (response !== null) {
          this.prescriptionList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    }
  },
  methods: {
    openDialog(item) {
      this.dialogCreateSelling = true
      this.valItem = item
    },
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
    height: 380px;
    margin: 18px;
  }
</style>
