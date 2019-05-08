<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header" :class="classes">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="16">
              <el-button type="primary" plain>添加</el-button>
            </el-col>
            <el-col :span="8">
              <el-input placeholder="请输入内容" v-model="params.search"><i slot="prefix" class="el-input__icon el-icon-search"></i></el-input>
            </el-col>
          </el-row>
        </div>
      </div>
      <div class="panel-body" :class="classes">
        <el-table :data="workers" style="width: 100%" height="100%">
          <el-table-column prop="id" label="ID" width="300"></el-table-column>
          <el-table-column prop="name" label="名称" width="200"></el-table-column>
          <el-table-column prop="ip" label="执行方式" width="140"></el-table-column>
          <el-table-column prop="status" label="执行次数" width="120"></el-table-column>
          <el-table-column prop="remark" label="标记"></el-table-column>
          <el-table-column prop="option" label="操作" width="130">
            <template slot-scope="scope">
              <el-button size="mini" icon="el-icon-edit" circle @click="handleEdit(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-tickets" plain circle @click="handleDelete(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-delete" type="danger" plain circle @click="handleDelete(scope.$index, scope.row)"></el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="pagination">
          <el-pagination background layout="prev, pager, next" :total="11"></el-pagination>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Log",
    data() {
      return {
        classes: ['animated', 'fade-in', 'fast'],
        params: {
          search: ""
        },
        workers: [{
          id: '6dd14a84-759a-4676-a511-7c3802077db5',
          name: '删除一个月以前的日志',
          ip: '192.168.128.234',
          status: 'connected',
          remark: '线下',
        }, {
          id: '6dd14a84-759a-4676-a511-7c3802077db5',
          name: '检查Kafka消费者运行状态',
          ip: '192.168.128.234',
          status: 'connected',
          remark: '线下',
        }, {
          id: '6dd14a84-759a-4676-a511-7c3802077db5',
          name: '给用户发送通知',
          ip: '192.168.128.234',
          status: 'connected',
          remark: '线下',
        }, {
          id: '6dd14a84-759a-4676-a511-7c3802077db5',
          name: 'Betterde',
          ip: '192.168.128.234',
          status: 'connected',
          remark: '线下',
        }, {
          id: '6dd14a84-759a-4676-a511-7c3802077db5',
          name: 'Betterde',
          ip: '192.168.128.234',
          status: 'connected',
          remark: '线下',
        }, {
          id: '6dd14a84-759a-4676-a511-7c3802077db5',
          name: 'Betterde',
          ip: '192.168.128.234',
          status: 'connected',
          remark: '线下',
        }, {
          id: '6dd14a84-759a-4676-a511-7c3802077db5',
          name: 'Betterde',
          ip: '192.168.128.234',
          status: 'connected',
          remark: '线下',
        }, {
          id: '6dd14a84-759a-4676-a511-7c3802077db5',
          name: 'Betterde',
          ip: '192.168.128.234',
          status: 'connected',
          remark: '线下',
        }, {
          id: '6dd14a84-759a-4676-a511-7c3802077db5',
          name: 'Betterde',
          ip: '192.168.128.234',
          status: 'connected',
          remark: '线下',
        }, {
          id: '6dd14a84-759a-4676-a511-7c3802077db5',
          name: 'Betterde',
          ip: '192.168.128.234',
          status: 'connected',
          remark: '线下',
        }]
      }
    },
    methods: {
      handleEdit(index, row) {
        window.console.log(index, row)
      },
      handleDelete(index, row) {
        window.console.log(index, row)
      }
    },
    beforeRouteEnter(to, from, next) {
      next(vm => {
        vm.classes = ['animated', 'fade-in', 'fast'];
      });
    },
    beforeRouteLeave (to, from, next) {
      this.classes = ['animated', 'fade-out', 'faster'];
      setTimeout(next, 200);
    }
  }
</script>

<style lang="scss">

</style>
