<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header" :class="classes">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="16">
              <el-button type="primary" plain>Create</el-button>
            </el-col>
            <el-col :span="8">
              <el-input placeholder="Search in here" v-model="params.search"><i slot="prefix" class="el-input__icon el-icon-search"></i></el-input>
            </el-col>
          </el-row>
        </div>
      </div>
      <div class="panel-body" :class="classes">
        <el-table :data="workers" style="width: 100%" height="100%" empty-text="No more data">
          <el-table-column prop="name" label="Pipeline" width="200"></el-table-column>
          <el-table-column prop="ip" label="Node" width="140"></el-table-column>
          <el-table-column prop="status" label="Level" width="120"></el-table-column>
          <el-table-column prop="remark" label="Created at"></el-table-column>
          <el-table-column prop="option" label="Action" width="130">
            <template slot-scope="scope">
              <el-button size="mini" icon="el-icon-edit" circle @click="handleEdit(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-tickets" plain circle @click="handleDelete(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-delete" type="danger" plain circle @click="handleDelete(scope.$index, scope.row)"></el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="pagination">
          <el-pagination background layout="prev, pager, next" :total="0"></el-pagination>
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
        workers: []
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
