<template>
  <div class="main-content">
    <el-table :data="tableData" style="width: 100%" height="250">
      <el-table-column fixed prop="date" label="名称" width="150"></el-table-column>
      <el-table-column prop="name" label="运行方式" width="120"></el-table-column>
      <el-table-column prop="province" label="运行设备"></el-table-column>
      <el-table-column prop="city" label="运行次数" width="120"></el-table-column>
      <el-table-column prop="address" label="异常次数" width="300"></el-table-column>
      <el-table-column prop="zip" label="日志" width="120"></el-table-column>
      <el-table-column prop="zip" label="操作" width="120"></el-table-column>
    </el-table>
    <el-pagination background layout="prev, pager, next" :total="1000"></el-pagination>
  </div>
</template>

<script>
  export default {
    name: "Task",
    data() {
      return {
        tableData: [{
          date: '删除日志',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '同步用户信息',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '计算通勤时间',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: 'HTTPS证书认证',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '备份数据库',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '给用户发送通知',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '生成统计报表',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }]
      }
    }
  }
</script>

<style lang="scss">

</style>
