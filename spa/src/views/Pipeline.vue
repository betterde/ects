<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header" :class="classes">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="8">
              <el-input placeholder="在这里输入要搜索的内容，按下回车进行搜索" v-model="params.search" @keyup.enter.native="fetchPipelines" @clear="handleClear" clearable>
                <i slot="prefix" class="el-input__icon el-icon-search"></i>
              </el-input>
            </el-col>
            <el-col :span="16" style="text-align: right">
              <el-button type="primary" plain @click="handleCreate">创建</el-button>
            </el-col>
          </el-row>
        </div>
      </div>
      <el-dialog title="创建流水线" :visible.sync="create.dialog" @close="handleClose('create')" width="40%" :close-on-click-modal="false">
        <el-form :model="create.params" :rules="create.rules" ref="create" label-position="top">
          <el-row :gutter="10">
            <el-col :span="24">
              <el-form-item label="名称" prop="name">
                <el-input v-model="create.params.name" autocomplete="off" placeholder="请输入名称"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="24">
              <el-form-item label="表达式" prop="spec">
                <el-input v-model="create.params.spec" placeholder="请输入表达式"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <collapse-view content="表达式参考">
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
              <el-form-item label="成功" prop="finished">
                <el-select v-model="create.params.finished" placeholder="请选择一个当流水线成功时触发任务" style="width: 100%">
                  <el-option v-for="task in tasks" :key="task.id" :label="task.name" :value="task.id"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="失败" prop="failed">
                <el-select v-model="create.params.failed" placeholder="请选择一个流水线失败时触发任务" style="width: 100%">
                  <el-option v-for="task in tasks" :key="task.id" :label="task.name" :value="task.id"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
          <el-form-item label="描述" prop="description">
            <el-input v-model="create.params.description" autocomplete="off" placeholder="请输入流水线的描述信息"></el-input>
          </el-form-item>
          <el-row :gutter="10">
            <el-col :span="12">
              <el-form-item label="状态" prop="status">
                <el-radio v-model="create.params.status" :label="0" border>禁用</el-radio>
                <el-radio v-model="create.params.status" :label="1" border>启用</el-radio>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="周期并行" prop="overlap">
                <el-radio v-model="create.params.overlap" :label="0" border>禁用</el-radio>
                <el-radio v-model="create.params.overlap" :label="1" border>启用</el-radio>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="create.dialog = false">取消</el-button>
          <el-button type="primary" @click="submit('create')">确定</el-button>
        </div>
      </el-dialog>
      <el-dialog title="编辑流水线" :visible.sync="update.dialog" @close="handleClose('update')" width="40%" :close-on-click-modal="false">
        <el-form :model="update.params" :rules="update.rules" ref="update" label-position="top">
          <el-row :gutter="10">
            <el-col :span="24">
              <el-form-item label="名称" prop="name">
                <el-input v-model="update.params.name" autocomplete="off" placeholder="请输入名称"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="10">
            <el-col :span="24">
              <el-form-item label="表达式" prop="spec">
                <el-input v-model="update.params.spec" placeholder="请输入表达式"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <collapse-view content="表达式参考">
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
              <el-form-item label="成功" prop="finished">
                <el-select v-model="update.params.finished" placeholder="请选择一个当流水线成功时触发任务" style="width: 100%">
                  <el-option v-for="task in tasks" :key="task.id" :label="task.name" :value="task.id"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="失败" prop="failed">
                <el-select v-model="update.params.failed" placeholder="请选择一个流水线失败时触发任务" style="width: 100%">
                  <el-option v-for="task in tasks" :key="task.id" :label="task.name" :value="task.id"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
          <el-form-item label="描述" prop="description">
            <el-input v-model="update.params.description" autocomplete="off"></el-input>
          </el-form-item>
          <el-row :gutter="10">
            <el-col :span="12">
              <el-form-item label="状态" prop="status">
                <el-radio v-model="update.params.status" :label="0" border>禁用</el-radio>
                <el-radio v-model="update.params.status" :label="1" border>启用</el-radio>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="周期并行" prop="overlap">
                <el-radio v-model="update.params.overlap" :label="0" border>禁用</el-radio>
                <el-radio v-model="update.params.overlap" :label="1" border>启用</el-radio>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="update.dialog = false">取消</el-button>
          <el-button type="primary" @click="submit('update')">确定</el-button>
        </div>
      </el-dialog>
      <el-dialog title="添加任务" :visible.sync="bind.dialog" @close="handleClose('bind')" width="600px" :close-on-click-modal="false">
        <el-form :model="bind.params" :rules="bind.rules" ref="bind" label-position="top">
          <el-row :gutter="10">
            <el-col :span="24">
              <el-form-item label="任务" prop="task_id">
                <el-select v-model="bind.params.task_id" placeholder="请选择任务" style="width: 100%">
                  <el-option v-for="task in tasks" :key="task.id" :label="task.name" :value="task.id" :disabled="task.disabled">
                    <span style="float: left">{{ task.name }}</span>
                    <span v-if="task.disabled === true" style="float: right; color: #8492a6; font-size: 13px">该任务已添加</span>
                  </el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="用户">
                <el-input v-model="bind.params.user" placeholder="执行命令的用户"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="超时" prop="timeout">
                <el-input v-model.number="bind.params.timeout" placeholder="超时时间"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="间隔" prop="interval">
                <el-input v-model.number="bind.params.interval" placeholder="间隔时间"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="重试" prop="retries">
                <el-input v-model.number="bind.params.retries" placeholder="失败重试次数"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="工作目录">
                <el-input v-model="bind.params.directory" placeholder="任务执行时的工作目录"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="环境变量">
                <el-input v-model="bind.params.environment" placeholder="任务执行时的环境变量"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="bind.dialog = false">取消</el-button>
          <el-button type="primary" @click="submit('bind')">确定</el-button>
        </div>
      </el-dialog>
      <el-dialog title="编辑任务" :visible.sync="modify.dialog" @close="handleClose('modify')" width="800px" :close-on-click-modal="false">
        <el-form :model="modify.params" :rules="modify.rules" ref="modify" label-position="top">
          <el-row :gutter="10">
            <el-col :span="24">
              <el-form-item label="任务">
                <el-select v-model="modify.params.task_id" placeholder="请选择任务" style="width: 100%">
                  <el-option v-for="task in tasks" :key="task.id" :label="task.name" :value="task.id"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="用户">
                <el-input v-model="modify.params.user" placeholder="执行命令的用户"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="超时">
                <el-input v-model.number="modify.params.timeout" placeholder="超时时间"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="间隔">
                <el-input v-model.number="modify.params.interval" placeholder="间隔时间"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="重试">
                <el-input v-model.number="modify.params.retries" placeholder="失败重试次数"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="工作目录">
                <el-input v-model="modify.params.directory" placeholder="任务执行时的工作目录"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="环境变量">
                <el-input v-model="modify.params.environment" placeholder="任务执行时的环境变量"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="modify.dialog = false">取消</el-button>
          <el-button type="primary" @click="submit('modify')">确定</el-button>
        </div>
      </el-dialog>
      <div class="panel-body" :class="classes">
        <el-table :data="pipelines" style="width: 100%" ref="pipeline" @expand-change="handleTableExpand">
          <el-table-column type="expand">
            <template slot-scope="props">
              <el-form label-position="top" inline class="table-expand">
                <el-row :gutter="10">
                  <el-col :span="8">
                    <el-form-item label="ID">
                      <span>{{ props.row.id }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="8">
                    <el-form-item label="成功">
                      <router-link class="el-link el-link--default is-underline" v-if="props.row.finished !== ''" :to="{path: '/task', query: {id: props.row.finished}}"><span>{{ props.row.finished }}</span></router-link>
                      <span v-else>未设置</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="8">
                    <el-form-item label="失败">
                      <router-link class="el-link el-link--default is-underline" v-if="props.row.failed !== ''" :to="{path: '/task', query: {id: props.row.failed}}"><span>{{ props.row.failed }}</span></router-link>
                      <span v-else>未设置</span>
                    </el-form-item>
                  </el-col>
                </el-row>
              </el-form>
              <el-divider>以下是流水线关联的任务可以拖动设置执行顺序</el-divider>
              <div style="text-align: center; width: 100%"></div>
              <el-table :data="props.row.steps === null ? [] : props.row.steps" row-key="step" ref="task" style="width: 100%" class="tasks-table">
                <el-table-column prop="step" label="#" width="50"></el-table-column>
                <el-table-column label="名称" width="180">
                  <template slot-scope="props">
                    <router-link class="el-link el-link--default is-underline" :to="{path: '/task', query: {id: props.row.task.id}}"><span>{{ props.row.task.name }}</span></router-link>
                  </template>
                </el-table-column>
                <el-table-column prop="timeout" label="超时" width="50"></el-table-column>
                <el-table-column prop="interval" label="间隔" width="50"></el-table-column>
                <el-table-column prop="retries" label="重试" width="50"></el-table-column>
                <el-table-column prop="user" label="用户" width="100"></el-table-column>
                <el-table-column prop="directory" label="工作目录" width="300">
                  <template slot-scope="props">
                    {{ props.row.directory === "" ? "未设置" : props.row.directory }}
                  </template>
                </el-table-column>
                <el-table-column prop="environment" label="环境变量">
                  <template slot-scope="props">
                    {{ props.row.environment === "" ? "未设置" : props.row.environment }}
                  </template>
                </el-table-column>
                <el-table-column prop="option" label="操作" width="100">
                  <template slot-scope="scope">
                    <el-tooltip class="item" effect="dark" content="编辑关系" placement="top">
                      <el-button size="mini" icon="el-icon-edit" circle
                                 @click="handleModify(scope.$index, scope.row)"></el-button>
                    </el-tooltip>
                    <el-tooltip class="item" effect="dark" content="移除绑定" placement="top">
                      <el-button size="mini" icon="el-icon-delete" type="danger" plain circle
                                 @click="handleRemove(scope.$index, scope.row)"></el-button>
                    </el-tooltip>
                  </template>
                </el-table-column>
              </el-table>
            </template>
          </el-table-column>
          <el-table-column prop="name" label="名称" width="200"></el-table-column>
          <el-table-column prop="spec" label="表达式" width="130"></el-table-column>
          <el-table-column label="状态" width="70">
            <template slot-scope="scope">
              <el-tag v-if="scope.row.status === 0" size="small" type="info">禁用</el-tag>
              <el-tag v-else size="small">正常</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="重复" width="60">
            <template slot-scope="scope">
              <el-tag v-if="scope.row.overlap === 0" size="small" type="info">否</el-tag>
              <el-tag v-else size="small">是</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="description" label="描述"></el-table-column>
          <el-table-column prop="created_at" label="创建于" width="160"></el-table-column>
          <el-table-column prop="option" label="操作" width="200">
            <template slot-scope="scope">
              <el-tooltip class="item" effect="dark" content="编辑" placement="top">
                <el-button size="mini" icon="el-icon-edit" circle
                           @click="handleEdit(scope.$index, scope.row)"></el-button>
              </el-tooltip>
              <el-tooltip class="item" effect="dark" content="添加任务" placement="top">
                <el-button size="mini" icon="el-icon-plus" circle
                           @click="handleBind(scope.$index, scope.row)"></el-button>
              </el-tooltip>
              <el-tooltip class="item" effect="dark" content="查询日志" placement="top">
                <el-button size="mini" icon="el-icon-tickets" plain circle
                           @click="handleDetail(scope.row)"></el-button>
              </el-tooltip>
              <el-tooltip class="item" effect="dark" content="同步到节点" placement="top">
                <el-button size="mini" icon="el-icon-refresh" plain circle
                           @click="handleSync(scope.row)"></el-button>
              </el-tooltip>
              <el-tooltip class="item" effect="dark" content="删除" placement="top">
                <el-button size="mini" icon="el-icon-delete" type="danger" plain circle
                           @click="handleDelete(scope.$index, scope.row)"></el-button>
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
  import Sortable from 'sortablejs'

  export default {
    name: "Pipeline",
    data() {
      return {
        classes: ['animated', 'fade-in', 'fast'],
        loading: false,
        params: {
          scene: 'table',
          search: '',
          page: 1
        },
        create: {
          dialog: false,
          params: {
            name: '',
            spec: '',
            description: '',
            status: 0,
            finished: '',
            failed: '',
            overlap: 1,
          },
          rules: {
            name: [
              {type: 'string', required: true, message: '请输入名称', trigger: 'blur'}
            ],
            spec: [
              {type: 'string', required: true, message: '请输入一个 Cron 表达式', trigger: 'blur'}
            ],
            finished: [
              {type: 'string', required: false, message: '请选择一个流水线执行成功后执行的任务', trigger: 'change'}
            ],
            failed: [
              {type: 'string', required: false, message: '请选择一个流水线执行失败后执行的任务', trigger: 'change'}
            ],
            status: [
              {type: 'number', required: true, message: '请选择状态', trigger: 'change'}
            ],
            overlap: [
              {type: 'number', required: true, message: '请选择是否允许重叠', trigger: 'change'}
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
              {type: 'string', required: true, message: '请输入名称', trigger: 'blur'}
            ],
            spec: [
              {type: 'string', required: true, message: '请输入一个 Cron 表达式', trigger: 'blur'}
            ],
            finished: [
              {type: 'string', required: false, message: '请选择一个流水线执行成功后执行的任务', trigger: 'change'}
            ],
            failed: [
              {type: 'string', required: false, message: '请选择一个流水线执行失败后执行的任务', trigger: 'change'}
            ],
            status: [
              {type: 'number', required: true, message: '请选择状态', trigger: 'change'}
            ],
            overlap: [
              {type: 'number', required: true, message: '请选择是否允许重叠', trigger: 'change'}
            ]
          },
        },
        current: {
          id: null,
          index: null
        },
        bind: {
          index: null,
          dialog: false,
          params: {
            pipeline_id: "",
            task_id: "",
            step: 1,
            timeout: 0,
            interval: 0,
            retries: 0,
            directory: "",
            user: "",
            environment: "",
            dependence: "strong",
          },
          rules:{
            task_id: [
              {type: 'string', required: true, message: '请选择一个任务', trigger: 'change'}
            ],
            timeout: [
              {type: 'integer', required: true, message: '请输入超时时间', trigger: 'blur'}
            ],
            interval: [
              {type: 'integer', required: true, message: '请输入间隔时间', trigger: 'blur'}
            ],
            retries: [
              {type: 'integer', required: true, message: '请输入重试次数', trigger: 'blur'}
            ],
            directory: [
              {type: 'string', required: false, message: '请输入工作目录', trigger: 'blur'}
            ],
            user: [
              {type: 'string', required: false, message: '请输入执行用户', trigger: 'blur'}
            ],
            environment: [
              {type: 'string', required: false, message: '请输入环境变量', trigger: 'blur'}
            ],
            dependence: [
              {type: 'string', required: false, message: '请选择是否依赖', trigger: 'blur'}
            ],
          },
        },
        modify: {
          id: null,
          index: null,
          dialog: false,
          params: {
            pipeline_id: "",
            task_id: "",
            step: 1,
            timeout: 0,
            interval: 0,
            retries: 0,
            directory: "",
            user: "",
            environment: "",
            dependence: "strong",
          },
          rules:{}
        },
        tasks: [],
        pipelines: [],
        meta: {
          limit: 10,
          page: 1,
          total: 0
        },
        sortable: null
      }
    },
    methods: {
      /**
       * 显示创建流水线的表单
       */
      handleCreate() {
        this.create.dialog = true;
        this.fetchTasks();
      },
      /**
       * 显示编辑流水线的表单
       */
      handleEdit(index, row) {
        this.fetchTasks();
        this.update.id = row.id;
        this.update.index = index;
        this.update.params = {...row};
        this.update.dialog = true;
      },
      /**
       * 显示绑定任务表单
       */
      handleBind(index, row) {
        api.pipeline.fetchTasks(row.id).then(res => {
          Vue.set(this.pipelines[index], 'steps', res.data);
          this.tasks.forEach(task => {
            if (row.steps.length === 0) {
              Vue.set(task, "disabled", false);
            } else {
              row.steps.forEach(step => {
                if (task.id === step.task_id) {
                  Vue.set(task, "disabled", true);
                }
              });
            }
          });
        }).catch(err => {
          this.$message.error(err.message);
        });
        this.bind.index = index;
        this.bind.params.pipeline_id = row.id;
        this.bind.dialog = true;
      },
      /**
       * 显示编辑关联详情表单
       */
      handleModify(index, row) {
        this.modify.id = row.id;
        this.modify.index = index;
        this.modify.params = {...row};
        this.modify.dialog = true;
      },
      /**
       * 获取流水线日志
       */
      handleDetail(row) {
        this.$router.push({
          name: 'pipeline_detail',
          params: {
            id: row.id
          }
        });
      },
      /**
       * 表单控制逻辑
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
            this.bind.dialog = false;
            this.bind.params.pipeline_id = null;
            break;
          case 'modify':
            this.$refs.modify.resetFields();
            this.modify.dialog = false;
            this.modify.id = null;
            this.modify.index = null;
            break;
        }
      },
      /**
       * 提交表单
       * @param form string
       */
      submit(form) {
        switch (form) {
          case 'create':
            this.$refs.create.validate((valid) => {
              if (valid) {
                api.pipeline.create(this.create.params).then(res => {
                  this.meta.total += 1;
                  // 判断是否需要跳转到最后一页
                  if (this.meta.total > (this.meta.limit * this.meta.page)) {
                    this.changePage(Math.ceil(this.meta.total / this.meta.limit));
                  } else {
                    // 如果不需要跳转则直接将数据追加到当前列表，减少API请求
                    this.pipelines.unshift(res.data);
                  }
                  this.handleClose(form);
                  this.$message.success({
                    offset: 95,
                    message: res.message
                  });
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
                api.pipeline.update(this.update.id, this.update.params).then(res => {
                  Vue.set(this.pipelines[this.update.index], 'name', res.data.name);
                  Vue.set(this.pipelines[this.update.index], 'spec', res.data.spec);
                  Vue.set(this.pipelines[this.update.index], 'description', res.data.description);
                  Vue.set(this.pipelines[this.update.index], 'status', res.data.status);
                  Vue.set(this.pipelines[this.update.index], 'finished', res.data.finished);
                  Vue.set(this.pipelines[this.update.index], 'failed', res.data.failed);
                  Vue.set(this.pipelines[this.update.index], 'overlap', res.data.overlap);
                  this.handleClose(form);
                  this.$message({
                    type: 'success',
                    offset: 95,
                    message: res.message
                  });
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
          case "bind":
            this.$refs.bind.validate((valid) => {
              if (valid) {
                api.pipeline.bindTask(this.bind.params).then(res => {
                  let steps = [];
                  if (this.pipelines[this.bind.index].steps instanceof Array) {
                    steps = this.pipelines[this.bind.index].steps;
                  }
                  steps.push(res.data);
                  Vue.set(this.pipelines[this.bind.index], 'steps', steps);
                  this.handleClose(form);
                  this.$message.success({
                    offset: 95,
                    message: res.message
                  });
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
          case 'modify':
            this.$refs.modify.validate((valid) => {
              if (valid) {
                api.pipeline.modifyRelation(this.modify.id, this.modify.params).then(res => {
                  Vue.set(this.pipelines[this.current.index].steps, this.modify.index, res.data);
                  this.handleClose(form);
                  this.$message.success({
                    offset: 95,
                    message: res.message
                  });
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
        }
      },
      /**
       * 删除流水线
       * @param index
       * @param row
       */
      handleDelete(index, row) {
        this.$confirm('此操作将删除流水线，是否继续', '警告', {
          confirmButtonText: '继续',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          api.pipeline.delete(row.id).then(res => {
            Vue.delete(this.pipelines, index);
            this.$message({
              type: 'success',
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
      /**
       * 解绑任务
       * @param index
       * @param row
       */
      handleRemove(index, row) {
        window.console.log(row);
        this.$confirm('此操作将解绑任务，是否继续', '警告', {
          confirmButtonText: '继续',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          api.pipeline.unbindTask(row.id).then(res => {
            Vue.delete(this.pipelines, index);
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
      /**
       * 点击站看行的处理逻辑
       * @param row
       * @param rows
       */
      handleTableExpand(row, rows) {
        // 如果当前展开超过一行
        if (rows.length > 1) {
          // 则遍历行，进行关闭，保证同时只有一个行展开
          rows.forEach(item => {
            if (row.id !== item.id) {
              this.$refs.pipeline.toggleRowExpansion(item, false);
            }
          });
        }

        // 判断当前展开的行如果大于0
        if (rows.length > 0) {
          // 最后一个行为当前行
          let intance = rows.pop();
          // 考虑到展开新行和关闭行都会触发 handleTableExpand 所以需要判断是否是当前行
          if (intance.id === row.id) {
            for (let index = 0; index < this.pipelines.length; index++) {
              if (this.pipelines[index].id === intance.id) {
                Vue.set(this.current, 'id', intance.id);
                Vue.set(this.current, 'index', index);
                // 调用 API 获取流水线关联的任务
                api.pipeline.fetchTasks(intance.id).then(res => {
                  // 将返回数据设置到对应流水线下的执行步骤属性中
                  Vue.set(this.pipelines[index], 'steps', res.data);
                  // 只有当关联任务超过一条时才开启拖拽排序功能
                  if (res.data instanceof Array && res.data.length > 1) {
                    this.rowDrop();
                  } else {
                    if (this.sortable !== null) {
                      this.sortable.option("disabled", true);
                    }
                  }
                }).catch(err => {
                  this.$message.error({
                    offset: 95,
                    message: err.message
                  });
                });
              }
            }
          }
        }
      },
      /**
       * 清空搜索框时触发
       */
      handleClear() {
        // 判断是否有 Pipeline 页面跳转传入的参数
        if (this.$route.query.hasOwnProperty("id")) {
          // 如果有则替换路由
          this.$router.replace("/pipeline");
        }
        this.fetchPipelines();
      },
      /**
       * 同步流水线到 ETCD
       */
      handleSync(row) {
        api.pipeline.sync(row.id).then(res => {
          this.$message.success({
            offset: 95,
            message: res.message
          });
          window.console.log(res.data)
        }).catch(err => {
          this.$message.error({
            offset: 95,
            message: err.message
          });
        });
      },
      /**
       * 获取任务列表
       */
      fetchTasks() {
        api.task.fetch(this.params).then(res => {
          Vue.set(this, "tasks", res.data);
        }).catch(err => {
          this.loading = false;
          if (err.hasOwnProperty('message')) {
            this.$message.error(err.message);
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
        this.loading = true;
        api.pipeline.fetch(this.params).then(res => {
          this.pipelines = res.data;
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
       * 分页跳转时触发
       */
      changePage(page) {
        this.meta.page = page;
        this.params.page = page;
        this.fetchPipelines();
      },
      /**
       * 为流水线关联的任务表格开启排序功能
       */
      rowDrop() {
        let tbody = document.querySelector('.el-table__body-wrapper .tasks-table tbody');
        let _this = this;
        this.sortable = Sortable.create(tbody, {
          onEnd({ newIndex, oldIndex }) {
            api.pipeline.sort({
              pipeline_id: _this.current.id,
              origin: oldIndex,
              current: newIndex
            }).then(res => {
              Vue.set(_this.pipelines[_this.current.index], 'steps', res.data);
              _this.$message.success({
                offset: 95,
                message: '排序成功'
              })
            }).catch(err => {
              _this.$message.error({
                offset: 95,
                message: err.message
              });
              return false;
            });
            let currRow = _this.pipelines[_this.current.index].steps.splice(oldIndex, 1)[0];
            _this.pipelines[_this.current.index].steps.splice(newIndex, 0, currRow);
          }
        })
      }
    },
    mounted() {
      // 如果存在查询参数则
      if (this.$route.query.hasOwnProperty("id")) {
        this.params.search = this.$route.query.id;
      }
      // 先获取流水线数据
      this.fetchPipelines();
      // 再获取用于选择任务的下拉列表数据
      this.fetchTasks();
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
      // 等待动画效果完成再离开
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
