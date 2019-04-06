<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header">
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
      <div class="panel-body">
        <el-table :data="tasks" style="width: 100%">
          <el-table-column prop="id" label="ID" width="300"></el-table-column>
          <el-table-column prop="name" label="名称" width="200"></el-table-column>
          <el-table-column prop="mode" label="方式" width="100"></el-table-column>
          <el-table-column prop="description" label="简介"></el-table-column>
          <el-table-column prop="status" label="状态" width="100"></el-table-column>
          <el-table-column prop="created_at" label="创建于" width="155"></el-table-column>
          <el-table-column prop="option" label="操作" width="130">
            <template slot-scope="scope">
              <el-button size="mini" icon="el-icon-edit" circle @click="handleEdit(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-tickets" plain circle @click="handleDelete(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-delete" type="danger" plain circle @click="handleDelete(scope.$index, scope.row)"></el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="pagination">
          <el-pagination background layout="prev, pager, next" :total="meta.total"></el-pagination>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import api from '../apis'

  export default {
    name: "Task",
    data() {
      return {
        loading: false,
        params: {
          search: ""
        },
        tasks: [],
        meta: {
          limit: 10,
          page: 1,
          total: 0
        }
      }
    },
    methods: {
      handleEdit(index, row) {
        window.console.log(index,row);
      },
      handleDelete(index, row) {
        window.console.log(index,row);
      },
      fetchTasks() {
        this.loading = true;
        api.task.fetch(this.params).then(res => {
          this.tasks = res.data;
          this.meta = res.meta;
        }).catch(err => {
          this.$message.warning(err.message)
        });
        this.loading = false;
      }
    },
    mounted() {
      this.fetchTasks();
    }
  }
</script>

<style lang="scss">

</style>
