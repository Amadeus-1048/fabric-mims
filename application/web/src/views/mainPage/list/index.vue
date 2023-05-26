<template>
  <div class="container">
    <el-alert type="success">
      <p>账户ID: {{ account_id }}</p>
      <p>用户名: {{ account_name }}</p>
    </el-alert>

    <el-card class="mainPage-card" >
      <div slot="header" class="mainPageHeader" style="text-align: center;">
        项目说明
      </div>

      <div class="mainPageNote">
        <el-tag>系统角色: </el-tag>
        <span style="margin-left: 5px;"> {{ '管理员、医生、患者、药店、保险机构' }}</span>
      </div><br>

      <div class="mainPageNote">
        <el-tag type="success">医生功能: </el-tag>
        <span style="margin-left: 5px;"> {{ '查看患者信息、创建病历、查看病历' }}</span>
      </div><br>

      <div class="mainPageNote">
        <el-tag type="warning">患者功能: </el-tag>
        <span style="margin-left: 5px;"> {{ '查看病历、查看药品订单、查看保险报销申请、发起保险报销申请' }}</span>
      </div><br>

      <div class="mainPageNote">
        <el-tag type="success">药店功能: </el-tag>
        <span style="margin-left: 5px;"> {{ '查看病历、创建药品订单、查看药品订单' }}</span>
      </div><br>

      <div class="mainPageNote">
        <el-tag>保险机构功能: </el-tag>
        <span style="margin-left: 5px;"> {{ '查看病历、查看保险报销记录、创建保险报销记录、审核患者发起的报销申请' }}</span>
      </div><br>

      <div class="mainPageNote">
        <el-tag type="danger">管理员: </el-tag>
        <span style="margin-left: 5px;"> {{ '查看各类记录、查看账号、创建账号' }}</span>
      </div><br>

    </el-card>


  </div>

</template>


<script>
import { mapGetters } from 'vuex'

export default {
  computed: {
    ...mapGetters([
      'account_id',
      'roles',
      'account_name',
    ])
  },

}
</script>

<style>
  .container p{
    font-size: 15px;
    width: 100%;
    /*text-align: center;*/
    min-height: 100%;
    overflow: hidden;
  }
  .mainPage-card {
    width: 80%;
    height: 400px;
    margin: 18px auto; /* auto可以使其水平居中 */
  }

  .mainPage-card .mainPageHeader {
    text-align: center;
    font-size: 20px;
  }

  .mainPage-card .mainPageNote {
    font-size: 16px;
  }
</style>
