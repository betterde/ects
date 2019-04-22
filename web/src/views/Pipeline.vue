<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="16">
              <el-button type="primary" plain @click="create.dialog = true">Create</el-button>
            </el-col>
            <el-col :span="8">
              <el-input placeholder="Search in here" v-model="params.search"><i slot="prefix" class="el-input__icon el-icon-search"></i></el-input>
            </el-col>
          </el-row>
        </div>
      </div>
      <el-dialog title="Create pipeline" :visible.sync="create.dialog" @close="createDialogClose" width="40%">
        <el-form :model="create.params" :rules="create.rules" ref="create" label-position="top">
          <el-row :gutter="10">
           <el-col :span="24">
             <el-form-item label="Name" prop="name">
               <el-input v-model="create.params.name" autocomplete="off"></el-input>
             </el-form-item>
           </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="24">
              <el-form-item label="Spec" prop="spec">
                <el-popover v-model="cronPopover">
                  <cron-expression @change="changeCron" @close="cronPopover=false" i18n="en"></cron-expression>
                  <el-input slot="reference" @click="cronPopover=true" v-model="create.params.spec" placeholder="Please enter a cron expression"></el-input>
                </el-popover>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <collapse-view content="Crontab reference">
                <pre>
                  <code style="display: -webkit-box; height: 200px">
  *    *    *    *    *    *
  ┬    ┬    ┬    ┬    ┬    ┬
  │    │    │    │    │    |
  │    │    │    │    │    └ day of week (0 - 7) (0 or 7 is Sun)
  │    │    │    │    └───── month (1 - 12)
  │    │    │    └────────── day of month (1 - 31)
  │    │    └─────────────── hour (0 - 23)
  │    └──────────────────── minute (0 - 59)
  └───────────────────────── second (0 - 59, optional)
                  </code>
                </pre>
              </collapse-view>
            </el-col>
          </el-row>
          <el-row :gutter="10">
           <el-col :span="12">
             <el-form-item label="Finished" prop="finished">
               <el-select v-model="create.params.finished" placeholder="Please select a task" style="width: 100%" no-data-text="No more data">
                 <el-option v-for="task in create.tasks" :key="task.value" :label="task.label" :value="task.value"></el-option>
               </el-select>
             </el-form-item>
           </el-col>
           <el-col :span="12">
             <el-form-item label="Failed" prop="failed">
               <el-select v-model="create.params.failed" placeholder="Please select a task" style="width: 100%" no-data-text="No more data">
                 <el-option v-for="task in create.tasks" :key="task.value" :label="task.label" :value="task.value"></el-option>
               </el-select>
             </el-form-item>
           </el-col>
          </el-row>
          <el-form-item label="Description" prop="description">
            <el-input v-model="create.params.description" autocomplete="off" @keyup.enter.native="submitCreateForm"></el-input>
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
          <el-button type="primary" @click="submitCreateForm">Confirm</el-button>
        </div>
      </el-dialog>
      <el-dialog title="Edit pipeline" :visible.sync="edit.dialog" @close="editDialogClose" width="500px">
        <el-form :model="edit.params" :rules="edit.rules" ref="edit">
          <el-form-item label="Name" prop="name">
            <el-input v-model="edit.params.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Description" prop="remark">
            <el-input v-model="edit.params.description" autocomplete="off" @keyup.enter.native="submitCreateForm"></el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="edit.dialog = false">Cancel</el-button>
          <el-button type="primary" @click="submitCreateForm">Confirm</el-button>
        </div>
      </el-dialog>
      <div class="panel-body">
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
                <el-row :gutter="10">
                  <el-col :span="12">
                    <el-form-item label="Overlap">
                      <span>{{ props.row.overlap }}</span>
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
          <el-table-column prop="spec" label="Spec" width="100"></el-table-column>
          <el-table-column label="Status" width="100">
            <template slot-scope="scope">
              <el-tag v-if="scope.row.status === 0" size="medium" type="info">Disabled</el-tag>
              <el-tag v-else size="medium">Normal</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="description" label="Description"></el-table-column>
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
          <el-pagination background layout="prev, pager, next" :total="meta.total"></el-pagination>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import api from '../apis'
  import CronExpression from '../components/CronExpression'

  export default {
    name: "Pipeline",
    data() {
      return {
        loading: false,
        params: {
          search: ''
        },
        cronPopover: false,
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
              {type: 'string', required: true, message: 'Please enter a spec', trigger: 'change'}
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
          tasks: [],
          next_execution: ""
        },
        edit: {
          dialog: false,
          params: {},
          rules: {},
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
      changeCron(value) {
        this.create.params.spec = value;
      },
      handleEdit(index, row) {
        window.console.log(index,row);
      },
      handleDelete(index, row) {
        window.console.log(index,row);
      },
      fetchTasks() {
        this.loading = true;
        api.pipeline.fetch(this.params).then(res => {
          this.tasks = res.data;
          this.meta = res.meta;
        }).catch(err => {
          this.$message.warning(err.message)
        });
        this.loading = false;
      },
      createDialogClose() {},
      editDialogClose() {},
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
    },
    mounted() {
      this.fetchTasks();
    },
    components: {
      CronExpression
    }
  }
</script>

<style lang="scss">

</style>
