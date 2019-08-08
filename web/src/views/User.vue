<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header" :class="classes">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="8">
              <el-input placeholder="在这里搜索" v-model="params.search"><i slot="prefix" class="el-input__icon el-icon-search"></i></el-input>
            </el-col>
            <el-col :span="16" style="text-align: right">
              <el-button type="primary" plain @click="handleCreate">创建</el-button>
            </el-col>
          </el-row>
        </div>
      </div>
      <el-dialog title="Create user" :visible.sync="create.dialog" @close="handleClose('create')" width="600px" :close-on-click-modal="false">
        <el-form :model="create.params" :rules="create.rules" ref="create" label-position="top">
          <el-row :gutter="10">
           <el-col :span="18">
             <el-form-item label="Name" prop="name">
               <el-input v-model="create.params.name" autocomplete="off"></el-input>
             </el-form-item>
           </el-col>
           <el-col :span="6">
             <el-form-item label="Manager" prop="manager">
               <el-select v-model="create.params.manager" placeholder="请选择">
                 <el-option label="Yes" :value="true"></el-option>
                 <el-option label="No" :value="false"></el-option>
               </el-select>
             </el-form-item>
           </el-col>
          </el-row>
          <el-form-item label="Email" prop="email">
            <el-input v-model="create.params.email" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Password" prop="pass">
            <el-input v-model="create.params.pass" placeholder="Password" show-password></el-input>
          </el-form-item>
          <el-form-item label="Confirm" prop="confirm">
            <el-input v-model="create.params.confirm" placeholder="Confirm" show-password></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="create.dialog = false">Cancel</el-button>
          <el-button type="primary" @click="submit('create')">Confirm</el-button>
        </div>
      </el-dialog>
      <el-dialog title="Update user" :visible.sync="update.dialog" @close="handleClose('update')" width="600px" :close-on-click-modal="false">
        <el-form :model="update.params" :rules="update.rules" ref="update" label-position="top">
          <el-row :gutter="10">
            <el-col :span="18">
              <el-form-item label="Name" prop="name">
                <el-input v-model="update.params.name" autocomplete="off"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="Manager" prop="manager">
                <el-select v-model="update.params.manager" placeholder="请选择">
                  <el-option label="Yes" :value="true"></el-option>
                  <el-option label="No" :value="false"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
          <el-form-item label="Email" prop="email">
            <el-input v-model="update.params.email" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Password" prop="pass">
            <el-input v-model="update.params.pass" placeholder="Password" show-password></el-input>
          </el-form-item>
          <el-form-item label="Confirm" prop="confirm">
            <el-input v-model="update.params.confirm" placeholder="Confirm" show-password></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="update.dialog = false">Cancel</el-button>
          <el-button type="primary" @click="submit('update')">Confirm</el-button>
        </div>
      </el-dialog>
      <div class="panel-body" :class="classes">
        <el-table :data="users" style="width: 100%" empty-text="No more data">
          <el-table-column type="expand">
            <template slot-scope="props">
              <el-form label-position="top" inline class="table-expand">
                <el-row :gutter="10">
                  <el-col :span="12">
                    <el-form-item label="创建于">
                      <span>{{ props.row.created_at }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="更新于">
                      <span>{{ props.row.updated_at }}</span>
                    </el-form-item>
                  </el-col>
                </el-row>
              </el-form>
            </template>
          </el-table-column>
          <el-table-column prop="id" label="ID" width="300"></el-table-column>
          <el-table-column prop="name" label="姓名" width="200"></el-table-column>
          <el-table-column prop="email" label="邮箱"></el-table-column>
          <el-table-column label="管理员" width="100">
            <template slot-scope="scope">
              <el-tag v-if="scope.row.manager === false" size="medium" type="info">NO</el-tag>
              <el-tag v-else size="medium">Yes</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="option" label="操作" width="130">
            <template slot-scope="scope">
              <el-button size="mini" icon="el-icon-edit" circle @click="handleUpdate(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-tickets" plain circle @click="handleDelete(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-delete" type="danger" plain circle @click="handleDelete(scope.$index, scope.row)"></el-button>
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
  import Vue from 'vue'
  import api from '../apis'

  export default {
    name: "User",
    data() {
      let validateCreatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('Please enter your password!'));
        } else {
          if (this.create.params.confirm !== '') {
            this.$refs.create.validateField('confirm');
          }
          callback();
        }
      };
      let confirmCreatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('Please confirm your password!'));
        } else if (value !== this.create.params.pass) {
          callback(new Error('Password mismatch!'));
        } else {
          callback();
        }
      };
      let validateUpdatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('Please enter your password!'));
        } else {
          if (this.update.params.confirm !== '') {
            this.$refs.update.validateField('confirm');
          }
          callback();
        }
      };
      let confirmUpdatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('Please confirm your password!'));
        } else if (value !== this.update.params.pass) {
          callback(new Error('Password mismatch!'));
        } else {
          callback();
        }
      };
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
            name: "",
            email: "",
            pass: "",
            confirm: "",
            manager: false,
          },
          rules: {
            name: [
              {type: 'string', required: true, message: 'Please enter a name', trigger: 'blur'}
            ],
            email: [
              {type: 'string', required: true, message: 'Please enter a email', trigger: 'blur'}
            ],
            pass: [
              {type: 'string', required: true, validator: validateCreatePass, trigger: 'blur'}
            ],
            confirm: [
              {type: 'string', required: true, validator: confirmCreatePass, trigger: 'blur'}
            ],
            manager: [
              {type: 'boolean', required: true, message: 'Please select user role', trigger: 'change'}
            ]
          }
        },
        update: {
          index: null,
          dialog: false,
          params: {
            name: "",
            email: "",
            pass: "",
            confirm: "",
            manager: false,
          },
          rules: {
            name: [
              {type: 'string', required: true, message: 'Please enter a name', trigger: 'blur'}
            ],
            email: [
              {type: 'string', required: true, message: 'Please enter a email', trigger: 'blur'}
            ],
            pass: [
              {type: 'string', required: true, validator: validateUpdatePass, trigger: 'blur'}
            ],
            confirm: [
              {type: 'string', required: true, validator: confirmUpdatePass, trigger: 'blur'}
            ],
            manager: [
              {type: 'boolean', required: true, message: 'Please select user role', trigger: 'change'}
            ]
          }
        },
        users: [],
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
      handleUpdate(index, row) {
        this.update.id = row.id;
        this.update.index = index;
        this.update.params = {...row};
        this.update.dialog = true;
      },
      submit(form) {
        switch (form) {
          case "create":
            this.$refs.create.validate((valid) => {
              if (valid) {
                api.user.create(this.create.params).then(res => {
                  this.meta.total += 1;
                  // 判断是否需要跳转到最后一页
                  if (this.meta.total > (this.meta.limit * this.meta.page)) {
                    this.changePage(Math.ceil(this.meta.total / this.meta.limit));
                  } else {
                    // 如果不需要跳转则直接将数据追加到当前列表，减少API请求
                    this.users.push(res.data);
                  }
                  this.handleClose(form);
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
                api.user.update(this.update.id, this.update.params).then(res => {
                  Vue.set(this.users[this.update.index], 'name', res.data.name);
                  Vue.set(this.users[this.update.index], 'email', res.data.email);
                  Vue.set(this.users[this.update.index], 'manager', res.data.manager);
                  Vue.set(this.users[this.update.index], 'updated_at', res.data.updated_at);
                  Vue.set(this.users[this.update.index], 'deleted_at', res.data.deleted_at);
                  this.handleClose(form);
                  this.$message.success(res.message);
                }).catch(err => {
                  this.$message.warning(err.message);
                });
              } else {
                return false;
              }
            });
            break
        }
      },
      /**
       * Delete user
       * @param index
       * @param row
       */
      handleDelete(index, row) {
        this.$confirm('This operation will delete the user, whether to continue?', 'Alert', {
          confirmButtonText: 'Confirm',
          cancelButtonText: 'Cancel',
          type: 'warning'
        }).then(() => {
          api.user.delete(row.id).then(res => {
            Vue.delete(this.users, index);
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
          case 'update':
            this.$refs.update.resetFields();
            this.update.dialog = false;
            break;
        }
      },
      /**
       * Fetch user list
       */
      fetchUsers() {
        this.loading = true;
        api.user.fetch(this.params).then(res => {
          this.users = res.data;
          this.meta = res.meta;
        }).catch(err => {
          this.$message.warning(err.message)
        });
        this.loading = false;
      },
      changePage(page) {
        this.meta.page = page;
        this.params.page = page;
        this.fetchUsers();
      },
    },
    mounted() {
      this.fetchUsers()
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
