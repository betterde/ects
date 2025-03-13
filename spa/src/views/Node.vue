<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header" :class="classes">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="8">
              <el-input placeholder="在这里输入要搜索的内容，按下回车进行搜索" v-model="params.search" @keyup.enter.native="fetchWorkers" @clear="handleClear" clearable>
                <i slot="prefix" class="el-input__icon el-icon-search"></i>
              </el-input>
            </el-col>
            <el-col :span="16" style="text-align: right">
              <el-button type="primary" plain @click="handleCreate">创建</el-button>
            </el-col>
          </el-row>
        </div>
      </div>
      <el-dialog title="创建节点" :visible.sync="create.dialog" @close="handleClose('create')" width="500px" :close-on-click-modal="false">
        <el-form :model="create.params" :rules="create.rules" ref="create">
          <el-form-item label="名称" prop="name">
            <el-input v-model="create.params.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="描述" prop="description">
            <el-input v-model="create.params.description" autocomplete="off" @keyup.enter.native="submit('create')"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="create.dialog = false">Cancel</el-button>
          <el-button type="primary" @click="submit('create')">Confirm</el-button>
        </div>
      </el-dialog>
      <el-dialog title="编辑节点" :visible.sync="update.dialog" @close="handleClose('update')" width="500px" :close-on-click-modal="false">
        <el-form :model="update.params" :rules="update.rules" ref="update">
          <el-form-item label="名称" prop="name">
            <el-input v-model="update.params.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="描述" prop="remark">
            <el-input v-model="update.params.description" autocomplete="off" @keyup.enter.native="submit('update')"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="update.dialog = false">Cancel</el-button>
          <el-button type="primary" @click="submit('update')">Confirm</el-button>
        </div>
      </el-dialog>
      <el-dialog title="关联流水线" :visible.sync="bind.dialog" @close="handleClose('bind')" width="500px" :close-on-click-modal="false">
        <el-form :model="bind.params" :rules="bind.rules" ref="bind" label-position="top">
          <el-form-item label="流水线" prop="pipeline_id">
            <el-select v-model="bind.params.pipeline_id" placeholder="请选择流水线" style="width: 100%">
              <el-option v-for="pipeline in pipelines" :key="pipeline.id" :label="pipeline.name" :value="pipeline.id" :disabled="pipeline.disabled">
                <span style="float: left">{{ pipeline.name }}</span>
                <span v-if="pipeline.disabled === true" style="float: right; color: #8492a6; font-size: 13px">该任务已添加</span>
              </el-option>
            </el-select>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="bind.dialog = false">取消</el-button>
          <el-button type="primary" @click="submit('bind')">确定</el-button>
        </div>
      </el-dialog>
      <div class="panel-body" :class="classes">
        <el-table :data="nodes" style="width: 100%" v-loading="loading" ref="nodes">
          <el-table-column type="expand">
            <template slot-scope="props">
              <el-form label-position="top" inline class="table-expand">

              </el-form>
              <el-divider>以下是节点关联的流水线</el-divider>
              <div style="text-align: center; width: 100%"></div>
              <el-table :data="props.row.pipelines === null ? [] : props.row.pipelines" row-key="step" ref="task" style="width: 100%" class="tasks-table">
                <el-table-column prop="pipeline.name" label="名称">
                  <template slot-scope="scope">
                    <router-link class="el-link el-link--default is-underline" :to="{path: '/pipeline', query: {id: scope.row.pipeline_id}}"><span>{{ scope.row.pipeline.name }}</span></router-link>
                  </template>
                </el-table-column>
                <el-table-column prop="pipeline.spec" label="表达式"></el-table-column>
                <el-table-column prop="pipeline.description" label="描述"></el-table-column>
                <el-table-column prop="created_at" label="关联于"></el-table-column>
                <el-table-column prop="option" label="操作" width="130">
                  <template slot-scope="scope">
                    <el-tooltip class="item" effect="dark" content="解绑流水线" placement="top">
                      <el-button size="mini" :disabled="scope.row.status === 'online'" icon="el-icon-delete" type="danger"
                                 plain circle @click="handleRemove(props.$index, scope.$index, scope.row)"></el-button>
                    </el-tooltip>
                  </template>
                </el-table-column>
              </el-table>
            </template>
          </el-table-column>
          <el-table-column prop="id" label="ID" width="300"></el-table-column>
          <el-table-column prop="name" label="名称"></el-table-column>
          <el-table-column prop="host" label="主机" width="140"></el-table-column>
          <el-table-column prop="port" label="端口" width="80"></el-table-column>
          <el-table-column prop="status" label="状态" width="80"></el-table-column>
          <el-table-column prop="version" label="版本" width="80"></el-table-column>
          <el-table-column prop="mode" label="类型" width="80"></el-table-column>
          <el-table-column prop="option" label="操作" width="130">
            <template slot-scope="scope">
              <el-tooltip class="item" effect="dark" content="编辑节点" placement="top">
                <el-button size="mini" icon="el-icon-edit" circle @click="handleUpdate(scope.$index, scope.row)"></el-button>
              </el-tooltip>
              <el-tooltip class="item" effect="dark" content="关联流水线" placement="top">
                <el-button size="mini" icon="el-icon-plus" :disabled="scope.row.mode === 'master'" circle @click="handleBind(scope.$index, scope.row)"></el-button>
              </el-tooltip>
              <el-tooltip class="item" effect="dark" content="删除节点" placement="top">
                <el-button size="mini" :disabled="scope.row.status === 'online'" icon="el-icon-delete" type="danger"
                           plain circle @click="handleDelete(scope.$index, scope.row)"></el-button>
              </el-tooltip>
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
          page: 1,
        },
        create: {
          dialog: false,
          params: {
            name: '',
            description: ''
          },
          rules: {
            name: [
              {type: 'string', required: true, message: '请输入节点名称', trigger: 'blur'}
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
              {type: 'string', required: true, message: '请输入节点名称', trigger: 'blur'}
            ]
          }
        },
        bind: {
          index: null,
          dialog: false,
          params: {
            node_id: null,
            pipeline_id: '',
          },
          rules: {
            pipeline_id: [
              {type: 'string', required: true, message: '请选择流水线', trigger: 'change'}
            ]
          }
        },
        nodes: [],
        pipelines: [],
        meta: {
          limit: 10,
          page: 1,
          total: 0
        }
      }
    },
    methods: {
      /**
       * 显示创建节点表单
       */
      handleCreate() {
        this.create.dialog = true;
      },
      /**
       * 显示编辑节点表单
       */
      handleUpdate(index, row) {
        this.update.id = row.id;
        this.update.index = index;
        this.update.params = {...row};
        this.update.dialog = true;
      },
      /**
       * 显示关联流水线表单
       */
      handleBind(index, row) {
        this.bind.index = index;
        this.bind.params.node_id = row.id;
        api.pipeline.fetch({scene: 'selector'}).then(res => {
          res.data.forEach(pipeline => {
            if (row.pipelines.length === 0) {
              pipeline.disabled = false;
            } else {
              row.pipelines.forEach(relation => {
                if (pipeline.id === relation.pipeline_id) {
                  pipeline.disabled = true;
                }
              });
            }
          });
          this.pipelines = res.data;
        }).catch(err => {
          this.$message.error({
            offset: 95,
            message: err.message
          });
        });
        this.bind.dialog = true;
      },
      submit(form) {
        switch (form) {
          case 'create':
            this.$refs.create.validate((valid) => {
              if (valid) {
                api.node.create(this.create.params).then(res => {
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
                  this.$message.error({
                    offset: 95,
                    message: err.message
                  });
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
                  this.$message.error({
                    offset: 95,
                    message: err.message
                  });
                });
              } else {
                return false;
              }
            });
            break;
          case 'bind':
            this.$refs.bind.validate((valid) => {
              if (valid) {
                api.node.bindPipeline(this.bind.params).then(res => {
                  let pipelines = this.nodes[this.bind.index].pipelines;
                  pipelines.push(res.data);
                  Vue.set(this.nodes[this.bind.index], 'pipelines', pipelines);
                  this.handleClose(form);
                  this.$message.success(res.message);
                }).catch(err => {
                  this.$message.error({
                    offset: 95,
                    message: err.message
                  });
                });
              } else {
                return false;
              }
            });
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
          case 'bind':
            this.$refs.bind.resetFields();
            this.bind.index = null;
            this.bind.dialog = false;
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
            this.$message.error({
              offset: 95,
              message: err.message
            });
          });
        }
      },
      /**
       * 解绑流水线
       * @param nindex
       * @param rindex
       * @param row
       */
      handleRemove(nindex, rindex, row) {
        this.$confirm('此操作将解绑流水线，是否继续', '警告', {
          confirmButtonText: '继续',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          api.node.unbindPipeline(row.id).then(res => {
            Vue.delete(this.nodes[nindex].pipelines, rindex);
            this.$message.success({
              offset: 95,
              message: res.message
            });
          }).catch(err => {
            this.$message.error({
              offset: 95,
              message: err.message
            })
          });
        }).catch(() => {
          this.$message.info({
            offset: 95,
            message: '操作已取消'
          });
        });
      },
      changePage(page) {
        this.meta.page = page;
        this.params.page = page;
        this.fetchWorkers();
      },
      /**
       * 获取节点信息
       */
      fetchWorkers() {
        this.loading = true;
        api.node.fetch(this.params).then(res => {
          this.nodes = res.data;
          this.meta = res.meta;
          this.loading = false;
        }).catch(err => {
          this.loading = false;
          if (err.hasOwnProperty('message')) {
            this.$message.error({
              offset: 95,
              message: err.message
            });
          }

          if (typeof err === 'string' || err instanceof String) {
            this.$message.error({
              offset: 95,
              message: err
            });
          }
        });
      },
      /**
       * 获取流水线列表
       */
      fetchPipelines() {
        api.pipeline.fetch({scene: 'selector'}).then(res => {
          this.pipelines = res.data;
        }).catch(err => {
          this.$message.error({
            offset: 95,
            message: err.message
          });
        });
      },
      /**
       * 清空搜索框时触发
       */
      handleClear() {
        // 判断是否有 Pipeline 页面跳转传入的参数
        if (this.$route.query.hasOwnProperty("id")) {
          // 如果有则替换路由
          this.$router.replace("/node");
        }
        this.fetchPipelines();
      },
    },
    mounted() {
      // 如果存在查询参数则
      if (this.$route.query.hasOwnProperty("id")) {
        this.params.search = this.$route.query.id;
      }
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
