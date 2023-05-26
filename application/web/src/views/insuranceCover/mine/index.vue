<template>
  <div class="container">
    <el-alert
      type="success"
    >
      <p>账户ID: {{ account_id }}</p>
      <p>用户名: {{ account_name }}</p>

    </el-alert>
    <div v-if="insuranceCoverList.length==0" style="text-align: center;">
      <el-alert
        title="查询不到数据"
        type="warning"
      />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in insuranceCoverList" :key="index" :span="6" :offset="1">
        <el-card class="insuranceCover-card">
          <div slot="header" class="clearfix">
            报销状态:
            <span :style="{ color: getStatusColor(val.status) }">{{ val.status }}</span>
          </div>

          <div class="item">
            <el-tag>报销记录ID: </el-tag>
            <span style="margin-left: 5px;">{{ val.id }}</span>
          </div>
          <div class="item">
            <el-tag type="success">病历ID: </el-tag>
            <span style="margin-left: 5px;">{{ val.prescription }}</span>
          </div>
          <div class="item">
            <el-tag type="warning">病人ID: </el-tag>
            <span style="margin-left: 5px;">{{ val.patient }}</span>
          </div>
          <div class="item">
            <el-tag type="info">创建时间: </el-tag>
            <span style="margin-left: 5px;">{{ val.created }}</span>
          </div>

        </el-card>
      </el-col>
    </el-row>

    <el-dialog v-loading="loadingDialog" :visible.sync="dialogCreateProcessing" :close-on-click-modal="false" @close="resetForm('ProcessForm')">
      <el-form ref="ProcessForm" :model="ProcessForm" :rules="rulesProcess" label-width="100px">
        <el-form-item label="操作" prop="operation">
          <el-select v-model="ProcessForm.operation" placeholder="请选择操作" @change="selectGetOperation">
            <el-option
              v-for="item in operationList"
              :key="item"
              :label="item"
              :value="item"
            >
              <span style="float: left">{{ item }}</span>
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="createProcessing('ProcessingForm')">确 定</el-button>
        <el-button @click="dialogCreateProcessing = false">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryInsuranceCoverList } from '@/api/insuranceCover'

export default {
  name: 'insuranceCover',
  data() {
    return {
      loading: true,
      loadingDialog: false,
      insuranceCoverList: [],
      dialogCreateProcessing: false,
      realForm: {
        price: 0,
        salePeriod: 0
      },
      ProcessingForm: {
        proprietor: ''
      },
      rulesProcess: {
        operation: [
          { required: true, message: '请选择操作', trigger: 'change' }
        ]
      },
      accountList: [],
      valItem: {},
      ProcessForm: {
        operation: ''
      },
      operationList: ['通过', '拒绝'],
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
      queryInsuranceCoverList().then(response => {
        if (response !== null) {
          this.insuranceCoverList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    } else {
      queryInsuranceCoverList({ patient: this.account_id }).then(response => {
        if (response !== null) {
          this.insuranceCoverList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    }
  },
  methods: {
    openProcessingDialog(item) {
      this.dialogCreateProcessing = true
      this.valItem = item
    },

    createProcessing(formName) {
      this.loadingDialog = true
      createProcessing({
        objectOfProcessing: this.valItem.realEstateId,
        donor: this.valItem.proprietor,
        grantee: this.ProcessingForm.proprietor
      }).then(response => {
        this.loadingDialog = false
        this.dialogCreateProcessing = false
        if (response !== null) {
          this.$message({
            type: 'success',
            message: '操作成功!'
          })
        } else {
          this.$message({
            type: 'error',
            message: '操作失败!'
          })
        }
        setTimeout(() => {
          window.location.reload()
        }, 1000)
      }).catch(_ => {
        this.loadingDialog = false
        this.dialogCreateProcessing = false
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    selectGet(account_id) {
      this.ProcessingForm.proprietor = account_id
    },
    selectGetOperation(operation) {
      this.ProcessForm.operation = operation
    },
    // 根据状态返回不同的颜色值
    getStatusColor(status) {
      if (status === '处理中') {
        return 'blue'; // 如果状态为 '处理中'，返回红色
      } else if (status === '已通过') {
        return 'green'; // 如果状态为 'approved'，返回绿色
      } else {
        return 'red'; // 默认返回黑色
      }
    }
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

  .insuranceCover-card {
    width: 280px;
    height: 280px;
    margin: 18px;
  }
</style>
