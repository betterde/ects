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
      <el-dialog title="Create pipeline" :visible.sync="create.dialog" @close="handleClose('create')" width="40%" :close-on-click-modal="false">
        <el-form :model="create.params" :rules="create.rules" ref="create" label-position="top">
          <el-row :gutter="10">
            <el-col :span="24">
              <el-form-item label="Name" prop="name">
                <el-input v-model="create.params.name" autocomplete="off" placeholder="Please enter a name"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="24">
              <el-form-item label="Spec" prop="spec">
                <el-input v-model="create.params.spec" placeholder="Please enter a cron expression"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <collapse-view content="Crontab reference">
                <pre><code style="display: -webkit-box; height: 200px">*    *    *    *    *    *
  ┬    ┬    ┬    ┬    ┬    ┬
  │    │    │    │    │    |
  │    │    │    │    │    └ day of week (0 - 7) (0 or 7 is Sun)
  │    │    │    │    └───── month (1 - 12)
  │    │    │    └────────── day of month (1 - 31)
  │    │    └─────────────── hour (0 - 23)
  │    └──────────────────── minute (0 - 59)
  └───────────────────────── second (0 - 59, optional)</code></pre>
              </collapse-view>
            </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="12">
              <el-form-item label="Finished" prop="finished">
                <el-select v-model="create.params.finished" placeholder="Please select a task" style="width: 100%" no-data-text="No more data">
                  <el-option v-for="task in tasks" :key="task.id" :label="task.name" :value="task.id"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="Failed" prop="failed">
                <el-select v-model="create.params.failed" placeholder="Please select a task" style="width: 100%" no-data-text="No more data">
                  <el-option v-for="task in tasks" :key="task.id" :label="task.name" :value="task.id"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
          <el-form-item label="Description" prop="description">
            <el-input v-model="create.params.description" autocomplete="off"></el-input>
          </el-form-item>
          <el-row :gutter="10">
            <el-col :span="12">
              <el-form-item label="Status" prop="status">
                <el-radio v-model="create.params.status" :label="0" border>Disable</el-radio>
                <el-radio v-model="create.params.status" :label="1" border>Enable</el-radio>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="Overlap" prop="overlap">
                <el-radio v-model="create.params.overlap" :label="0" border>Disable</el-radio>
                <el-radio v-model="create.params.overlap" :label="1" border>Enable</el-radio>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="create.dialog = false">Cancel</el-button>
          <el-button type="primary" @click="submit('create')">Confirm</el-button>
        </div>
      </el-dialog>
      <el-dialog title="Edit pipeline" :visible.sync="update.dialog" @close="handleClose('update')" width="40%" :close-on-click-modal="false">
        <el-form :model="update.params" :rules="update.rules" ref="edit" label-position="top">
          <el-row :gutter="10">
            <el-col :span="24">
              <el-form-item label="Name" prop="name">
                <el-input v-model="update.params.name" autocomplete="off"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="24">
              <el-form-item label="Spec" prop="spec">
                <el-input v-model="update.params.spec" placeholder="Please enter a cron expression"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <collapse-view content="Crontab reference">
                <pre><code style="display: -webkit-box; height: 200px">*    *    *    *    *    *
  ┬    ┬    ┬    ┬    ┬    ┬
  │    │    │    │    │    |
  │    │    │    │    │    └ day of week (0 - 7) (0 or 7 is Sun)
  │    │    │    │    └───── month (1 - 12)
  │    │    │    └────────── day of month (1 - 31)
  │    │    └─────────────── hour (0 - 23)
  │    └──────────────────── minute (0 - 59)
  └───────────────────────── second (0 - 59, optional)</code></pre>
              </collapse-view>
            </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="12">
              <el-form-item label="Finished" prop="finished">
                <el-select v-model="update.params.finished" placeholder="Please select a task" style="width: 100%" no-data-text="No more data">
                  <el-option v-for="task in tasks" :key="task.id" :label="task.name" :value="task.id"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="Failed" prop="failed">
                <el-select v-model="update.params.failed" placeholder="Please select a task" style="width: 100%" no-data-text="No more data">
                  <el-option v-for="task in tasks" :key="task.id" :label="task.name" :value="task.id"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
          <el-form-item label="Description" prop="description">
            <el-input v-model="update.params.description" autocomplete="off"></el-input>
          </el-form-item>
          <el-row :gutter="10">
            <el-col :span="12">
              <el-form-item label="Status" prop="status">
                <el-radio v-model="update.params.status" :label="0" border>Disable</el-radio>
                <el-radio v-model="update.params.status" :label="1" border>Enable</el-radio>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="Overlap" prop="overlap">
                <el-radio v-model="update.params.overlap" :label="0" border>Disable</el-radio>
                <el-radio v-model="update.params.overlap" :label="1" border>Enable</el-radio>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="update.dialog = false">Cancel</el-button>
          <el-button type="primary" @click="submit('update')">Confirm</el-button>
        </div>
      </el-dialog>
      <div class="panel-body" :class="classes">
        <el-table :data="pipelines" style="width: 100%" empty-text="No more data">
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
                    <el-form-item label="Overlap">
                      <span>{{ props.row.overlap === 1 }}</span>
                    </el-form-item>
                  </el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">
                    <el-form-item label="Finished">
                      <span>{{ props.row.finished }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="Failed">
                      <span>{{ props.row.failed }}</span>
                    </el-form-item>
                  </el-col>
                </el-row>
              </el-form>
            </template>
          </el-table-column>
          <el-table-column prop="name" label="Name" width="200"></el-table-column>
          <el-table-column prop="spec" label="Spec" width="100"></el-table-column>
          <el-table-column label="Status" width="100">
            <template slot-scope="scope">
              <el-tag v-if="scope.row.status === 0" size="medium" type="info">Disabled</el-tag>
              <el-tag v-else size="medium">Normal</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="description" label="Description"></el-table-column>
          <el-table-column prop="created_at" label="Created At" width="160"></el-table-column>
          <el-table-column prop="option" label="Action" width="130">
            <template slot-scope="scope">
              <el-button size="mini" icon="el-icon-edit" circle
                         @click="handleEdit(scope.$index, scope.row)"></el-button>
              <el-button size="mini" icon="el-icon-tickets" plain circle
                         @click="handleQueryLog(scope.row)"></el-button>
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
  import Vue from 'vue'
  import api from '../apis'
  import {mapState} from 'vuex'

  export default {
    name: "Pipeline",
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
            spec: '',
            description: '',
            status: 1,
            finished: '',
            failed: '',
            overlap: 1,
          },
          rules: {
            name: [
              {type: 'string', required: true, message: 'Please enter a name', trigger: 'blur'}
            ],
            spec: [
              {type: 'string', required: true, message: 'Please enter a spec', trigger: 'blur'}
            ],
            finished: [
              {type: 'string', required: false, message: 'Please select a task', trigger: 'change'}
            ],
            failed: [
              {type: 'string', required: false, message: 'Please select a task', trigger: 'change'}
            ],
            status: [
              {type: 'number', required: true, message: 'Please select pipeline status', trigger: 'change'}
            ],
            overlap: [
              {type: 'number', required: true, message: 'Please select pipeline overlap', trigger: 'change'}
            ]
          }
        },
        update: {
          id: null,
          index: null,
          dialog: false,
          params: {
            name: '',
            spec: '',
            description: '',
            status: 1,
            finished: '',
            failed: '',
            overlap: 1,
          },
          rules: {
            name: [
              {type: 'string', required: true, message: 'Please enter a name', trigger: 'blur'}
            ],
            spec: [
              {type: 'string', required: true, message: 'Please enter a spec', trigger: 'blur'}
            ],
            finished: [
              {type: 'string', required: true, message: 'Please select a task', trigger: 'change'}
            ],
            failed: [
              {type: 'string', required: true, message: 'Please select a task', trigger: 'change'}
            ],
            status: [
              {type: 'number', required: true, message: 'Please select pipeline status', trigger: 'change'}
            ],
            overlap: [
              {type: 'number', required: true, message: 'Please select pipeline overlap', trigger: 'change'}
            ]
          },
        },
        tasks: [],
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
       * Show create pipeline dialog
       */
      handleCreate() {
        this.create.dialog = true;
        this.fetchTasks();
      },
      /**
       * Show edit pipeline dialog
       */
      handleEdit(index, row) {
        this.fetchTasks();
        this.update.id = row.id;
        this.update.index = index;
        this.update.params = {...row};
        this.update.dialog = true;
      },
      /**
       * Get pipeline log
       */
      handleQueryLog(row) {
        this.$router.push({
          name: 'log',
          params: {
            pipeline_id: row.id
          }
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
          case 'edit':
            this.$refs.edit.resetFields();
            this.edit.dialog = false;
            this.update.id = null;
            this.update.index = null;
            break;
        }
      },
      submit(form) {
        switch (form) {
          case "create":
            this.$refs.create.validate((valid) => {
              if (valid) {
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
                return false;
              }
            });
            break;
          case "update":
            this.$refs.update.validate((valid) => {
              if (valid) {
                api.pipeline.update(this.update.id, this.update.params).then(res => {
                  Vue.set(this.pipelines[this.update.index], 'name', res.data.name);
                  Vue.set(this.pipelines[this.update.index], 'spec', res.data.spec);
                  Vue.set(this.pipelines[this.update.index], 'description', res.data.description);
                  Vue.set(this.pipelines[this.update.index], 'status', res.data.status);
                  Vue.set(this.pipelines[this.update.index], 'finished', res.data.finished);
                  Vue.set(this.pipelines[this.update.index], 'failed', res.data.failed);
                  Vue.set(this.pipelines[this.update.index], 'overlap', res.data.overlap);
                  this.handleClose('update');
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
       * Delete pipeline
       * @param index
       * @param row
       */
      handleDelete(index, row) {
        this.$confirm('This operation will delete the pipeline, whether to continue?', 'Alert', {
          confirmButtonText: 'Confirm',
          cancelButtonText: 'Cancel',
          type: 'warning'
        }).then(() => {
          api.pipeline.delete(row.id).then(res => {
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
       * Fetch task data
       */
      fetchTasks() {
        api.task.fetch(this.params).then(res => {
          this.tasks = res.data;
        }).catch(err => {
          this.$message.warning(err.message)
        });
      },
      /**
       * Fetch pipeline data
       */
      fetchPipelines() {
        this.loading = true;
        api.pipeline.fetch(this.params).then(res => {
          this.pipelines = res.data;
          this.meta = res.meta;
        }).catch(err => {
          this.$message.error(err.data.message);
        });
        this.loading = false;
      },
      /**
       * Change page
       */
      changePage(page) {
        this.meta.page = page;
        this.params.page = page;
        this.fetchPipelines();
      }
    },
    computed: {
      ...mapState({
        profile: state => state.account.profile,
      }),
    },
    mounted() {
      this.fetchPipelines();
    },
    /**
     * Modify the class name before entering the current component
     * @param to
     * @param from
     * @param next
     */
    beforeRouteEnter(to, from, next) {
      next(vm => {
        vm.classes = ['animated', 'fade-in', 'fast'];
      });
    },
    /**
     * Modify the class name before leaving the current component
     * @param to
     * @param from
     * @param next
     */
    beforeRouteLeave (to, from, next) {
      this.classes = ['animated', 'fade-out', 'faster'];
      // Wait for leave animation to finish
      setTimeout(next, 200);
    }
  }
</script>

<style lang="scss">
  .fade-enter-active, .fade-leave-active {
    transition: opacity .5s;
  }
  .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
    opacity: 0;
  }
</style>
