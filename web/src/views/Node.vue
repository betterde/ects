<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header" :class="classes">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="8">
              <el-input placeholder="在这里搜索" v-model="params.search" @keyup.enter.native="fetchWorkers"><i slot="prefix" class="el-input__icon el-icon-search"></i></el-input>
            </el-col>
            <el-col :span="16" style="text-align: right">
              <el-button type="primary" plain @click="handleCreate">创建</el-button>
            </el-col>
          </el-row>
        </div>
      </div>
      <el-dialog title="Create node" :visible.sync="create.dialog" @close="handleClose('create')" width="500px" :close-on-click-modal="false">
        <el-form :model="create.params" :rules="create.rules" ref="create">
          <el-form-item label="Name" prop="name">
            <el-input v-model="create.params.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Description" prop="description">
            <el-input v-model="create.params.description" autocomplete="off" @keyup.enter.native="submit('create')"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="create.dialog = false">Cancel</el-button>
          <el-button type="primary" @click="submit('create')">Confirm</el-button>
        </div>
      </el-dialog>
      <el-dialog title="Edit node" :visible.sync="update.dialog" @close="handleClose('update')" width="500px" :close-on-click-modal="false">
        <el-form :model="update.params" :rules="update.rules" ref="update">
          <el-form-item label="Name" prop="name">
            <el-input v-model="update.params.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Description" prop="remark">
            <el-input v-model="update.params.description" autocomplete="off" @keyup.enter.native="submit('update')"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="update.dialog = false">Cancel</el-button>
          <el-button type="primary" @click="submit('update')">Confirm</el-button>
        </div>
      </el-dialog>
      <div class="panel-body" :class="classes">
        <el-table :data="nodes" style="width: 100%" v-loading="loading">
          <el-table-column prop="id" label="ID" width="300"></el-table-column>
          <el-table-column prop="name" label="名称"></el-table-column>
          <el-table-column prop="host" label="主机" width="140"></el-table-column>
          <el-table-column prop="port" label="端口" width="80"></el-table-column>
          <el-table-column prop="status" label="状态" width="80"></el-table-column>
          <el-table-column prop="version" label="版本" width="80"></el-table-column>
          <el-table-column prop="mode" label="类型" width="80"></el-table-column>
          <el-table-column prop="option" label="操作" width="100">
            <template slot-scope="scope">
              <el-button size="mini" icon="el-icon-edit" circle @click="handleUpdate(scope.$index, scope.row)"></el-button>
              <el-button size="mini" :disabled="scope.row.status === 'online'" icon="el-icon-delete" type="danger" plain circle @click="handleDelete(scope.$index, scope.row)"></el-button>
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
    name: 'Worker',
    data() {
      return {
        classes: ['animated', 'fade-in', 'fast'],
        loading: false,
        params: {
          search: '',
          page: 1
        },
        create: {
          dialog: false,
          params: {
            name: '',
            description: ''
          },
          rules: {
            name: [
              {type: 'string', required: true, message: 'Please enter a name', trigger: 'blur'}
            ],
            description: [
              {type: 'string', required: false, message: 'Please enter a description', trigger: 'blur'}
            ]
          }
        },
        update: {
          id: null,
          index: null,
          dialog: false,
          params: {
            name: '',
            description: ''
          },
          rules: {
            name: [
              {type: 'string', required: true, message: 'Please enter a name', trigger: 'blur'}
            ],
            description: [
              {type: 'string', required: false, message: 'Please enter a description', trigger: 'blur'}
            ]
          }
        },
        nodes: [],
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
          case 'create':
            this.$refs.create.validate((valid) => {
              if (valid) {
                api.node.create(this.createParams).then(res => {
                  this.meta.total += 1;
                  // 判断是否需要跳转到最后一页
                  if (this.meta.total > (this.meta.limit * this.meta.page)) {
                    this.changePage(Math.ceil(this.meta.total / this.meta.limit));
                  } else {
                    // 如果不需要跳转则直接将数据追加到当前列表，减少API请求
                    this.nodes.push(res.data);
                  }
                  this.$message.success(res.message);
                  this.handleClose(form);
                }).catch(err => {
                  this.$message.warning(err.message);
                });
              } else {
                return false;
              }
            });
            break;
          case 'update':
            this.$refs.update.validate((valid) => {
              if (valid) {
                api.node.update(this.update.id, this.update.params).then(res => {
                  Vue.set(this.nodes[this.update.index], 'name', res.data.name);
                  Vue.set(this.nodes[this.update.index], 'description', res.data.description);
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
        }
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
            this.update.id = null;
            this.update.index = null;
            break;
        }
      },
      /**
       * Delete handler
       * @param index
       * @param row
       */
      handleDelete(index, row) {
        if (row.hasOwnProperty('id')) {
          api.node.delete(row.id).then(res => {
            Vue.delete(this.nodes, index);
            this.$message.success(res.message);
          }).catch(err => {
            this.$message.warning(err.message);
          });
        }
      },
      changePage(page) {
        this.meta.page = page;
        this.params.page = page;
        this.fetchWorkers();
      },
      fetchWorkers() {
        this.loading = true;
        api.node.fetch(this.params).then(res => {
          this.nodes = res.data;
          this.meta = res.meta;
        }).catch(err => {
          this.$message.warning(err.message)
        });
        this.loading = false;
      }
    },
    mounted() {
      this.fetchWorkers();
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
