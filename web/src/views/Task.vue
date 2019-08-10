<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header" :class="classes">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="8">
              <el-input placeholder="在这里搜索" v-model="params.search" @keyup.enter.native="fetchTasks" @clear="handleClear" clearable>
                <i slot="prefix" class="el-input__icon el-icon-search"></i>
              </el-input>
            </el-col>
            <el-col :span="16" style="text-align: right">
              <el-button type="primary" plain @click="handleCreate">创建</el-button>
            </el-col>
          </el-row>
        </div>
      </div>
      <el-dialog title="创建任务" :visible.sync="create.dialog" @close="handleClose('create')" width="600px" :close-on-click-modal="false">
        <el-form :model="create.params" :rules="create.rules" ref="create" label-position="top">
          <el-row :gutter="10">
            <el-col :span="18">
              <el-form-item label="名称" prop="name">
                <el-input v-model="create.params.name" autocomplete="off" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="类型" prop="mode">
                <el-select v-model="create.params.mode" placeholder="请选择">
                  <el-option key="shell" label="Shell命令" value="shell"></el-option>
                  <el-option key="http" label="HTTP请求" value="http"></el-option>
                  <el-option key="mail" label="邮件通知" value="mail"></el-option>
                  <el-option key="hook" label="Hook通知" value="hook"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="描述" prop="description">
                <el-input v-model="create.params.description" autocomplete="off" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="10" v-if="['http', 'hook'].includes(create.params.mode)">
            <el-col :span="4">
              <el-form-item label="方法" prop="method">
                <el-select v-model="create.params.method" placeholder="方法">
                  <el-option key="get" label="GET" value="get"></el-option>
                  <el-option key="post" label="POST" value="post"></el-option>
                  <el-option key="put" label="PUT" value="put"></el-option>
                  <el-option key="patch" label="PATCH" value="patch"></el-option>
                  <el-option key="delete" label="DELETE" value="delete"></el-option>
                  <el-option key="head" label="HEAD" value="head"></el-option>
                  <el-option key="options" label="OPTIONS" value="options"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="20">
              <el-form-item label="URL" prop="mode">
                <el-input v-model="create.params.url" placeholder="请输入请求的URL地址" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="24" v-if="['shell', 'mail', 'http'].includes(create.params.mode)">
              <el-form-item label="内容" prop="content">
                <el-input v-model="create.params.content" placeholder="此处填写Shell命令、邮件通知地址或HTTP请求体（JSON字符串）"
                          autocomplete="off" @keyup.enter.native="submit('create')" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="create.dialog = false">取消</el-button>
          <el-button type="primary" @click="submit('create')">确定</el-button>
        </div>
      </el-dialog>
      <el-dialog title="编辑任务" :visible.sync="update.dialog" @close="handleClose('update')" width="600px" :close-on-click-modal="false">
        <el-form :model="update.params" :rules="update.rules" ref="update" label-position="top">
          <el-row :gutter="10">
            <el-col :span="18">
              <el-form-item label="名称" prop="name">
                <el-input v-model="update.params.name" autocomplete="off" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="类型" prop="mode">
                <el-select v-model="update.params.mode" placeholder="请选择">
                  <el-option key="shell" label="Shell命令" value="shell"></el-option>
                  <el-option key="http" label="HTTP请求" value="http"></el-option>
                  <el-option key="mail" label="邮件通知" value="mail"></el-option>
                  <el-option key="hook" label="Hook通知" value="hook"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="描述" prop="description">
                <el-input v-model="update.params.description" autocomplete="off" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="10" v-if="['http', 'hook'].includes(update.params.mode)">
            <el-col :span="4">
              <el-form-item label="方法" prop="method">
                <el-select v-model="update.params.method" placeholder="方法">
                  <el-option key="get" label="GET" value="get"></el-option>
                  <el-option key="post" label="POST" value="post"></el-option>
                  <el-option key="put" label="PUT" value="put"></el-option>
                  <el-option key="patch" label="PATCH" value="patch"></el-option>
                  <el-option key="delete" label="DELETE" value="delete"></el-option>
                  <el-option key="head" label="HEAD" value="head"></el-option>
                  <el-option key="options" label="OPTIONS" value="options"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="20">
              <el-form-item label="URL" prop="mode">
                <el-input v-model="update.params.url" placeholder="请输入请求的URL地址" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="24" v-if="['shell', 'mail', 'http'].includes(update.params.mode)">
              <el-form-item label="内容" prop="content">
                <el-input v-model="update.params.content" placeholder="此处填写Shell命令、邮件通知地址或HTTP请求体（JSON字符串）"
                          autocomplete="off" @keyup.enter.native="submit('create')" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="update.dialog = false">取消</el-button>
          <el-button type="primary" @click="submit('update')">确定</el-button>
        </div>
      </el-dialog>
      <div class="panel-body" :class="classes">
        <el-table :data="tasks" style="width: 100%">
          <el-table-column type="expand">
            <template slot-scope="props">
              <el-form label-position="top" inline class="table-expand">
                <el-row :gutter="10">
                  <el-col v-if="props.row.mode === 'http'" :span="12">
                    <el-form-item label="请求方法">
                      <span>{{ props.row.method.toLocaleUpperCase() }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col v-if="props.row.mode === 'http'" :span="12">
                    <el-form-item label="URL">
                      <span>{{ props.row.url }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="24">
                    <el-form-item label="内容">
                      <pre class="task-pre"><code class="task-content">{{ props.row.content === '' ? props.row.method.toLocaleUpperCase() + ' ' + props.row.url : props.row.content }}</code></pre>
                    </el-form-item>
                  </el-col>
                </el-row>
              </el-form>
            </template>
          </el-table-column>
          <el-table-column prop="id" label="ID" width="300"></el-table-column>
          <el-table-column prop="name" label="名称" width="200"></el-table-column>
          <el-table-column prop="description" label="描述"></el-table-column>
          <el-table-column prop="mode" label="模式" width="60"></el-table-column>
          <el-table-column prop="created_at" label="创建于" width="155"></el-table-column>
          <el-table-column prop="option" label="操作" width="130">
            <template slot-scope="scope">
              <el-button size="mini" icon="el-icon-edit" circle
                         @click="handleUpdate(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-tickets" plain circle
                         @click="handleQueryLog(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-delete" type="danger" plain circle
                         @click="handleDelete(scope.$index, scope.row)"></el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="pagination">
          <el-pagination background layout="prev, pager, next" :current-page.sync="meta.page" :total="meta.total"
                         @current-change="changePage"></el-pagination>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import Vue from 'vue'
  import api from '../apis'

  export default {
    name: "Task",
    data() {
      return {
        classes: ['animated', 'fade-in', 'fast'],
        loading: false,
        params: {
          search: "",
          page: 1
        },
        create: {
          dialog: false,
          params: {
            name: '',
            url: '',
            mode: 'shell',
            method: 'post',
            description: '',
            content: ''
          },
          rules: {
            name: [
              {type: 'string', required: true, message: '请输入任务名称', trigger: 'blur'}
            ],
            mode: [
              {type: 'string', required: true, message: '请选择任务类型', trigger: 'change'}
            ],
            description: [
              {type: 'string', required: false, message: '请输入任务描述', trigger: 'blur'}
            ],
            content: [
              {type: 'string', required: true, message: '请输入任务内容', trigger: 'blur'}
            ]
          }
        },
        update: {
          id: null,
          index: null,
          dialog: false,
          params: {
            name: '',
            url: '',
            mode: '',
            method: '',
            description: '',
            content: ''
          },
          rules: {
            name: [
              {type: 'string', required: true, message: '请输入任务名称', trigger: 'blur'}
            ],
            mode: [
              {type: 'string', required: true, message: '请选择任务类型', trigger: 'change'}
            ],
            description: [
              {type: 'string', required: false, message: '请输入任务描述', trigger: 'blur'}
            ],
            content: [
              {type: 'string', required: false, message: '请输入任务内容', trigger: 'blur'}
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
      /**
       * 显示创建任务的表单
       */
      handleCreate() {
        this.create.dialog = true;
      },
      /**
       * 显示更新表单
       */
      handleUpdate(index, row) {
        this.update.id = row.id;
        this.update.index = index;
        this.update.params = {...row};
        this.update.dialog = true;
      },
      /**
       * 表单关闭事件触发逻辑
       * @param form 表单名称
       */
      handleClose(form) {
        switch (form) {
          case 'create':
            this.$refs.create.resetFields();
            this.create.dialog = false;
            break;
          case 'update':
            this.$refs.update.resetFields();
            this.update.dialog = false;
            break;
        }
      },
      /**
       * 提交表单
       * @param form 表单名称
       */
      submit(form) {
        switch (form) {
          case "create":
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
            break;
          case "update":
            this.$refs.update.validate((valid) => {
              if (valid) {
                api.task.update(this.update.params.id, this.update.params).then(res => {
                  this.handleClose('update');
                  this.tasks[this.update.index].name = res.data.name;
                  this.tasks[this.update.index].description = res.data.description;
                  this.tasks[this.update.index].content = res.data.content;
                  this.tasks[this.update.index].updated_at = res.data.updated_at;
                  this.$message.success(res.message);
                  this.update.index = null;
                }).catch(err => {
                  this.$message.warning(err.message);
                });
              } else {
                return false;
              }
            });
            break;
        }
      },

      handleQueryLog(index, row) {
        // TODO Redirect to log view
        window.console.log(index, row);
      },
      handleDelete(index, row) {
        this.$confirm('This operation will delete the task, whether to continue?', 'Alert', {
          confirmButtonText: 'Confirm',
          cancelButtonText: 'Cancel',
          type: 'warning'
        }).then(() => {
          api.task.delete(row.id).then(res => {
            Vue.delete(this.tasks, index);
            this.$message({
              type: 'success',
              message: res.message
            });
          }).catch(err => {
            this.$message.error(err.message)
          });
        }).catch(() => {
          this.$message({
            type: 'info',
            message: 'Operation canceled!'
          });
        });
      },
      handleClear() {
        // 判断是否有 Pipeline 页面跳转传入的参数
        if (this.$route.query.hasOwnProperty("id")) {
          // 如果有则替换路由
          this.$router.replace("/task");
        }
        this.fetchTasks();
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
      // 如果存在查询参数则
      if (this.$route.query.hasOwnProperty("id")) {
        this.params.search = this.$route.query.id;
      }
      this.fetchTasks();
    },
    beforeRouteEnter(to, from, next) {
      next(vm => {
        vm.classes = ['animated', 'fade-in', 'fast'];
      });
    },
    beforeRouteLeave(to, from, next) {
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
