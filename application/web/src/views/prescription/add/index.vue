<template>
  <div class="app-container">
    <el-form ref="ruleForm" v-loading="loading" :model="ruleForm" :rules="rules" label-width="100px">

      <el-form-item label="病人" prop="patient">
        <el-select v-model="ruleForm.patient" placeholder="请选择病人" @change="selectGet">
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
      </el-form-item>
      <el-form-item label="诊断" prop="diagnosis">
        <el-input v-model="ruleForm.diagnosis" style="width: 197px"/>
      </el-form-item>
      <el-form-item label="药品名" prop="drug_name">
        <el-input v-model="ruleForm.drug_name" style="width: 197px" />
      </el-form-item>
      <el-form-item label="药品数量" prop="drug_amount">
        <el-input-number v-model="ruleForm.drug_amount" :precision="0" :step="1" :min="1" style="width: 197px" />
      </el-form-item>
      <el-form-item label="备注" prop="comment">
        <el-input v-model="ruleForm.comment" style="width: 197px" />
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
import { createPrescription } from '@/api/prescription'

export default {
  name: 'AddPrescription',
  data() {
    var checkArea = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error('必须大于0'))
      } else {
        callback()
      }
    }
    return {
      ruleForm: {
        patient: '',
        diagnosis: '',
        drug_name: '',
        drug_amount: '',
        comment:'',
      },
      accountList: [],
      rules: {
        patient: [
          { message: '请选择病人', trigger: 'change' }
        ],
        // totalArea: [
        //   { validator: checkArea, trigger: 'blur' }
        // ],
        // livingSpace: [
        //   { validator: checkArea, trigger: 'blur' }
        // ]
      },
      loading: false,
    }
  },
  computed: {
    ...mapGetters([
      'account_id'
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
            createPrescription({
              doctor: this.account_id,
              patient: this.ruleForm.patient,
              diagnosis: this.ruleForm.diagnosis,
              drug_name: this.ruleForm.drug_name,
              drug_amount: this.ruleForm.drug_amount.toString(),
              hospital:'0feceb66ffc1',
              comment: this.ruleForm.comment,
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
    selectGet(account_id) {
      this.ruleForm.patient = account_id
    },
  }
}
</script>

<style scoped>
</style>
