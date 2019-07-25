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
              <el-input placeholder="Search in here" v-model="params.search"><i slot="prefix"
                                                                                class="el-input__icon el-icon-search"></i>
              </el-input>
            </el-col>
          </el-row>
        </div>
      </div>
      <div class="panel-body" :class="classes">
        <el-table :data="tasks" style="width: 100%">
          <el-table-column type="expand">
            <template slot-scope="props">
              <el-form label-position="top" inline class="table-expand">
                <el-row :gutter="10">
                 <el-col :span="12">
                   <el-form-item label="ID">
                     <span>{{ props.row.id }}</span>
                   </el-form-item>
                 </el-col>
                  <el-col :span="12">
                    <el-form-item label="TeamID">
                      <span>{{ props.row.team_id }}</span>
                    </el-form-item>
                  </el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="24">
                    <el-form-item label="Content">
                      <pre class="task-pre"><code class="task-content">{{ props.row.content }}</code></pre>
                    </el-form-item>
                  </el-col>
                </el-row>
              </el-form>
            </template>
          </el-table-column>
          <el-table-column prop="name" label="Name" width="200"></el-table-column>
          <el-table-column prop="mode" label="Mode" width="100"></el-table-column>
          <el-table-column prop="description" label="Description"></el-table-column>
          <el-table-column prop="created_at" label="CreatedAt" width="155"></el-table-column>
          <el-table-column prop="option" label="Action" width="130">
            <template slot-scope="scope">
              <el-button size="mini" icon="el-icon-edit" circle
                         @click="handleEdit(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-tickets" plain circle
                         @click="handleDelete(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-delete" type="danger" plain circle
                         @click="handleDelete(scope.$index, scope.row)"></el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="pagination">
          <el-pagination background layout="prev, pager, next" :current-page.sync="meta.page" :total="meta.total" @current-change="changePage"></el-pagination>
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
        classes: ['animated', 'fade-in', 'fast'],
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
        window.console.log(index, row);
      },
      handleDelete(index, row) {
        window.console.log(index, row);
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
      },
      changePage(page) {
        this.meta.page = page;
        this.params.page = page;
        this.fetchTasks();
      },
    },
    mounted() {
      this.fetchTasks();
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
  .task-pre {
    width: 100%;
    margin: 0 0;
    color: #5e6d82;
    padding: 10px 20px;
    background-color: #e6effb;
  }
  .task-content {
    width: 100%;
  }
</style>
