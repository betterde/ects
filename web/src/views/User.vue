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
              <el-input placeholder="Search in here" v-model="params.search"><i slot="prefix" class="el-input__icon el-icon-search"></i></el-input>
            </el-col>
          </el-row>
        </div>
      </div>
      <el-dialog title="Create user" :visible.sync="create.dialog" @close="handleClose('create')" width="600px" :close-on-click-modal="false">
        <el-form :model="create.params" :rules="create.rules" ref="create" label-position="top">

        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="create.dialog = false">Cancel</el-button>
          <el-button type="primary" @click="submitCreateForm">Confirm</el-button>
        </div>
      </el-dialog>
      <div class="panel-body" :class="classes">
        <el-table :data="users" style="width: 100%">
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
                    <el-form-item label="Create at">
                      <span>{{ props.row.created_at }}</span>
                    </el-form-item>
                  </el-col>
                </el-row>
              </el-form>
            </template>
          </el-table-column>
          <el-table-column prop="name" label="Name" width="200"></el-table-column>
          <el-table-column prop="email" label="Email" width="300"></el-table-column>
          <el-table-column label="Manager" width="100">
            <template slot-scope="scope">
              <el-tag v-if="scope.row.manager === 0" size="medium" type="info">NO</el-tag>
              <el-tag v-else size="medium">Yes</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="Status">
            <template slot-scope="scope">
              <el-tag v-if="scope.row.deleted_at === ''" size="medium">Normal</el-tag>
              <el-tag v-else size="medium" type="info">Disabled</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="option" label="Action" width="130">
            <template slot-scope="scope">
              <el-button size="mini" icon="el-icon-edit" circle @click="handleEdit(scope.$index, scope.row)"></el-button>
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
  import api from '../apis'

  export default {
    name: "User",
    data() {
      return {
        classes: ['animated', 'fade-in', 'fast'],
        loading: false,
        params: {
          search: ""
        },
        create: {
          dialog: false,
          params: {},
          rules: {}
        },
        edit: {
          dialog: false,
          params: {},
          rules: {}
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
      submitCreateForm() {
        this.$refs.create.validate((valid) => {
          if (valid) {
            if (this.profile.team_id.length === 36) {
              this.create.params.team_id = this.profile.team_id;
              api.pipeline.create(this.create.params).then(res => {
                this.meta.total += 1;
                // 判断是否需要跳转到最后一页
                if (this.meta.total > (this.meta.limit * this.meta.page)) {
                  this.changePage(Math.ceil(this.meta.total / this.meta.limit));
                } else {
                  // 如果不需要跳转则直接将数据追加到当前列表，减少API请求
                  this.pipelines.push(res.data);
                }
                this.handleClose('create');
                this.$message.success(res.message);
              }).catch(err => {
                this.$message.warning(err.message);
              });
            } else {
              this.$message.warning('You can\'t create it until you join a team');
            }
          } else {
            return false;
          }
        });
      },
      handleEdit(index, row) {
        window.console.log(index, row);
      },
      handleDelete(index, row) {
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
