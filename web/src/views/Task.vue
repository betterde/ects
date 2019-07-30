<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header" :class="classes">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="16">
              <el-button type="primary" plain @click="handleCreate">Create</el-button>
            </el-col>
            <el-col :span="8">
              <el-input placeholder="Search in here" v-model="params.search">
                <i slot="prefix" class="el-input__icon el-icon-search"></i>
              </el-input>
            </el-col>
          </el-row>
        </div>
      </div>
      <el-dialog title="Create task" :visible.sync="create.dialog" @close="handleClose('create')" width="600px">
        <el-form :model="create.params" :rules="create.rules" ref="create" label-position="top">
          <el-row :gutter="10">
           <el-col :span="12">
             <el-form-item label="Name" prop="name">
               <el-input v-model="create.params.name" autocomplete="off"></el-input>
             </el-form-item>
           </el-col>
           <el-col :span="12">
             <el-form-item label="Mode" prop="mode">
               <el-select v-model="create.params.mode" placeholder="Please select a mode">
                 <el-option label="shell" value="shell"></el-option>
                 <el-option label="http" value="http"></el-option>
               </el-select>
             </el-form-item>
           </el-col>
          </el-row>
          <el-form-item label="Description" prop="description">
            <el-input v-model="create.params.description" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Content" prop="content">
            <el-input v-model="create.params.content" autocomplete="off" @keyup.enter.native="submitCreateForm"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="create.dialog = false">Cancel</el-button>
          <el-button type="primary" @click="submitCreateForm">Confirm</el-button>
        </div>
      </el-dialog>
      <div class="panel-body" :class="classes">
        <el-table :data="tasks" style="width: 100%">
          <el-table-column type="expand">
            <template slot-scope="props">
              <el-form label-position="top" inline class="table-expand">
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
          <el-table-column prop="id" label="ID" width="300"></el-table-column>
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
        create: {
          dialog: false,
          params: {
            name: '',
            mode: '',
            description: '',
            content: ''
          },
          rules: {
            name: [
              {type: 'string', required: true, message: 'Please enter a task name', trigger: 'blur'}
            ],
            description: [
              {type: 'string', required: false, message: 'Please enter a task description', trigger: 'blur'}
            ],
            mode: [
              {type: 'string', required: true, message: 'Please select a mode', trigger: 'change'}
            ],
            content: [
              {type: 'string', required: true, message: 'Please enter task command', trigger: 'blur'}
            ]
          }
        },
        edit: {
          dialog: false,
          params: {
            name: '',
            mode: '',
            description: '',
            content: ''
          },
          rules: {
            name: [
              {type: 'string', required: true, message: 'Please enter a task name', trigger: 'blur'}
            ],
            description: [
              {type: 'string', required: false, message: 'Please enter a task description', trigger: 'blur'}
            ],
            mode: [
              {type: 'string', required: true, message: 'Please select a mode', trigger: 'change'}
            ],
            content: [
              {type: 'string', required: true, message: 'Please enter task command', trigger: 'blur'}
            ]
          }
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
      handleCreate() {
        this.create.dialog = true;
      },
      submitCreateForm() {
        this.$refs.create.validate((valid) => {
          if (valid) {
            api.task.create(this.create.params).then(res => {
              this.meta.total += 1;
              // 判断是否需要跳转到最后一页
              if (this.meta.total > (this.meta.limit * this.meta.page)) {
                this.changePage(Math.ceil(this.meta.total / this.meta.limit));
              } else {
                // 如果不需要跳转则直接将数据追加到当前列表，减少API请求
                this.tasks.push(res.data);
              }
              this.handleClose('create');
              this.$message.success(res.message);
            }).catch(err => {
              this.$message.warning(err.message);
            });
          } else {
            return false;
          }
        });
      },
      handleEdit(index, row) {
        window.console.log(index, row);
      },
      /**
       * Close create or edit dialog handler
       * @param form
       */
      handleClose(form) {
        switch (form) {
          case 'create':
            this.$refs.create.resetFields();
            this.create.dialog = false;
            break;
          case 'edit':
            this.$refs.edit.resetFields();
            this.edit.dialog = false;
            break;
        }
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
  .el-select {
    width: 100%;
  }
</style>
