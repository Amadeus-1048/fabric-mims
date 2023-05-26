<template>
  <div class="app-container">
    <el-form ref="ruleForm" v-loading="loading" :model="ruleForm" :rules="rules" label-width="100px">

      <el-form-item label="病历" prop="prescription">
        <el-select v-model="ruleForm.prescription" placeholder="请选择病历" @change="selectGetPrescriptionList">
          <el-option
              v-for="item in prescriptionList"
              :key="item.id"
              :label="item.diagnosis"
              :value="item.id"
          >
            <span style="float: left">{{ item.id }}</span>
            <span style="float: right; color: #8492a6; font-size: 13px">{{ item.diagnosis }}</span>
          </el-option>
        </el-select>
      </el-form-item>


      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/accountV2'
import {createInsuranceCover} from '@/api/insuranceCover'
import { queryPrescriptionList} from '@/api/prescription'


export default {
  name: 'AddInsuranceCover',
  data() {
    return {
      isPatient: /病人/.test(this.account_name) ,
      ruleForm: {
        patient: this.account_id,
        prescription: '',
        status: 'processing',
      },
      accountList: [],
      prescriptionList: [],
      rules: {
        patient: [
          { required: true, message: '请选择病人', trigger: 'change' }
        ],

      },
      loading: false,
    }
  },
  computed: {
    ...mapGetters([
      'account_id',
      'account_name'
    ])
  },
  created() {
    queryAccountList().then(response => {
      if (response !== null) {
        // 过滤掉管理员
        this.accountList = response.filter(item =>
            //item.account_name !== '医生'
            /病人$/.test(item.account_name)
        )
      }
    })
    queryPrescriptionList({'patient':this.account_id}).then(response => {
        this.ruleForm.patient = this.account_id
          this.prescriptionList = response
    })
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即创建?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loading = true
            createInsuranceCover({
              patient: this.ruleForm.patient,
              prescription: this.ruleForm.prescription,
              status:'processing',
            }).then(response => {
              this.loading = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '创建成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '创建失败!'
                })
              }
            }).catch(_ => {
              this.loading = false
            })
          }).catch(() => {
            this.loading = false
            this.$message({
              type: 'info',
              message: '已取消创建'
            })
          })
        } else {
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },

    selectGetPrescriptionList(prescription) {
      this.ruleForm.prescription = prescription
    },
  }
}
</script>

<style scoped>
</style>
