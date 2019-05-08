<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header" :class="classes">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="16">
              <el-button type="primary" plain @click="createDialog = true">Create</el-button>
            </el-col>
            <el-col :span="8">
              <el-input placeholder="Search in here" v-model="params.search" @keyup.enter.native="fetchWorkers"><i slot="prefix" class="el-input__icon el-icon-search"></i></el-input>
            </el-col>
          </el-row>
        </div>
      </div>
      <el-dialog title="Create node" :visible.sync="createDialog" @close="createDialogClose" width="500px">
        <el-form :model="createParams" :rules="createRules" ref="create">
          <el-form-item label="Name" prop="name">
            <el-input v-model="createParams.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Description" prop="description">
            <el-input v-model="createParams.description" autocomplete="off" @keyup.enter.native="submitCreateForm"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="createDialog = false">Cancel</el-button>
          <el-button type="primary" @click="submitCreateForm">Confirm</el-button>
        </div>
      </el-dialog>
      <el-dialog title="Edit node" :visible.sync="editDialog" @close="editDialogClose" width="500px">
        <el-form :model="editParams" :rules="editRules" ref="edit">
          <el-form-item label="Name" prop="name">
            <el-input v-model="editParams.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Description" prop="remark">
            <el-input v-model="editParams.description" autocomplete="off" @keyup.enter.native="submitCreateForm"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="editDialog = false">Cancel</el-button>
          <el-button type="primary" @click="submitCreateForm">Confirm</el-button>
        </div>
      </el-dialog>
      <div class="panel-body" :class="classes">
        <el-table :data="nodes" style="width: 100%" v-loading="loading" empty-text="No more data">
          <el-table-column prop="id" label="ID" width="300"></el-table-column>
          <el-table-column prop="name" label="Name" width="200"></el-table-column>
          <el-table-column prop="host" label="Host" width="140"></el-table-column>
          <el-table-column prop="port" label="Port" width="80"></el-table-column>
          <el-table-column prop="status" label="Status" width="120"></el-table-column>
          <el-table-column prop="description" label="Description"></el-table-column>
          <el-table-column prop="option" label="Action" width="100">
            <template slot-scope="scope">
              <el-button size="mini" icon="el-icon-edit" circle @click="handleEdit(scope.$index, scope.row)"></el-button>
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
    name: 'Worker',
    data() {
      return {
        classes: ['animated', 'fade-in', 'fast'],
        loading: false,
        params: {
          search: '',
          page: 1
        },
        createDialog: false,
        createParams: {
          name: '',
          description: ''
        },
        createRules: {
          name: [
            {type: 'string', required: true, message: '请输入节点名称', trigger: 'blur'}
          ],
          description: [
            {type: 'string', required: false, message: '请输入节点名称', trigger: 'blur'}
          ]
        },
        editDialog: false,
        editParams: {
          name: '',
          description: ''
        },
        editRules: {
          name: [
            {type: 'string', required: true, message: '请输入节点名称', trigger: 'blur'}
          ],
          description: [
            {type: 'string', required: false, message: '请输入节点名称', trigger: 'blur'}
          ]
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
      createDialogClose() {
        this.createDialog = false;
        this.$refs.create.resetFields();
      },
      editDialogClose() {
        this.editDialog = false;
        this.$refs.edit.resetFields();
      },
      submitCreateForm() {
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
              this.createDialogClose();
            }).catch(err => {
              this.$message.warning(err.message);
            });
          } else {
            return false;
          }
        });
      },
      handleEdit(index, row) {
        this.editDialog = true;
        this.editParams.name = row.name;
        this.editParams.description = row.description;
      },
      // 删除节点
      handleDelete(index, row) {
        if (row.hasOwnProperty('id')) {
          api.node.delete({id: row.id}).then(res => {
            this.$message.success(res.message);
            this.fetchWorkers();
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
