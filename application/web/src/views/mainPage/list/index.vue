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

      <div class="mainPageNote" style="float: left;">
        <el-tag>说明示例: </el-tag>
        <span style="margin-left: 5px;"> {{ 'amadeus' }}</span>
      </div>
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
    height: 440px;
    margin: 18px auto; /* auto可以使其水平居中 */
  }

  .mainPage-card .mainPageHeader {
    text-align: center;
    font-size: 20px;
  }

  .mainPage-card .mainPageNote {
    float: left;
    font-size: 16px;
  }
</style>
