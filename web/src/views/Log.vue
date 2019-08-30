<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header" :class="classes">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="16">
            </el-col>
            <el-col :span="8">
              <el-input placeholder="在这里输入要搜索的内容，按下回车进行搜索" v-model="params.search" @keyup.enter.native="fetchLogs" @clear="handleClear" clearable>
                <i slot="prefix" class="el-input__icon el-icon-search"></i>
              </el-input>
            </el-col>
          </el-row>
        </div>
      </div>
      <div class="panel-body" :class="classes">
        <el-tabs v-model="active" @tab-click="handleClick">
          <el-tab-pane label="流水线日志" name="pipeline">
            <el-table :data="logs" style="width: 100%" ref="pipeline" @expand-change="handleTableExpand">
              <el-table-column type="expand">
                <template slot-scope="props">
                  <el-timeline :reverse="reverse">
                    <el-timeline-item :timestamp="step.created_at" placement="top" :key="index" :type="step.status === 'finished' ? 'success' : 'danger'" size="large" :icon="step.status === 'finished' ? 'el-icon-success' : 'el-icon-error'" v-for="(step, index) in props.row.steps">
                      <el-card>
                        <h3>{{ step.task_name }}</h3>
                        <el-form label-position="top" inline class="table-expand">
                          <el-row :gutter="10">
                            <el-col :span="2">
                              <el-form-item label="ID">
                                <span>{{ step.id }}</span>
                              </el-form-item>
                            </el-col>
                            <el-col :span="6">
                              <el-form-item label="任务">
                                <router-link class="el-link el-link--default is-underline" :to="{path: '/task', query: {id: step.task_id}}"><span>{{ step.task_id }}</span></router-link>
                              </el-form-item>
                            </el-col>
                            <el-col :span="2">
                              <el-form-item label="类型">
                                <span>{{ step.mode }}</span>
                              </el-form-item>
                            </el-col>
                            <el-col :span="2">
                              <el-form-item label="超时">
                                <span>{{ step.timeout }}</span>
                              </el-form-item>
                            </el-col>
                            <el-col :span="2">
                              <el-form-item label="重试">
                                <span>{{ step.retries }}</span>
                              </el-form-item>
                            </el-col>
                            <el-col :span="2">
                              <el-form-item label="耗时">
                                <span>{{ step.duration }}</span>
                              </el-form-item>
                            </el-col>
                            <el-col :span="4">
                              <el-form-item label="开始于">
                                <span>{{ step.begin_with }}</span>
                              </el-form-item>
                            </el-col>
                            <el-col :span="4">
                              <el-form-item label="结束于">
                                <span>{{ step.finish_with }}</span>
                              </el-form-item>
                            </el-col>
                          </el-row>
                        </el-form>
                        <pre class="task-pre"><code class="task-content" :class="{ bash: step.mode === 'shell' }">{{ step.content }}</code></pre>
                        <collapse-view content="执行结果">
                          <pre class="task-pre"><code class="task-content bash" v-html="step.result"></code>
                          </pre>
                        </collapse-view>
                      </el-card>
                    </el-timeline-item>
                  </el-timeline>
                </template>
              </el-table-column>
              <el-table-column prop="pipeline_id" label="流水线" width="300">
                <template slot-scope="scope">
                  <router-link class="el-link el-link--default is-underline" :to="{path: '/pipeline', query: {id: scope.row.pipeline_id}}"><span>{{ scope.row.pipeline_id }}</span></router-link>
                </template>
              </el-table-column>
              <el-table-column label="节点" width="160">
                <template slot-scope="scope">
                  <router-link class="el-link el-link--default is-underline" :to="{path: '/node', query: {id: scope.row.node_id}}"><span>{{ scope.row.worker_name }}</span></router-link>
                </template>
              </el-table-column>
              <el-table-column label="结果">
                <template slot-scope="scope">
                  <el-tag v-if="scope.row.status === 0" size="small" type="danger">失败</el-tag>
                  <el-tag v-else size="small">成功</el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="duration" label="耗时(秒)"></el-table-column>
              <el-table-column prop="begin_with" label="开始于" width="160"></el-table-column>
              <el-table-column prop="finish_with" label="结束于" width="160"></el-table-column>
              <el-table-column prop="created_at" label="创建于" width="160"></el-table-column>
            </el-table>
            <div class="pagination">
              <el-pagination background layout="prev, pager, next" :current-page.sync="meta.page" :total="meta.total" @current-change="changePage"></el-pagination>
            </div>
          </el-tab-pane>
          <el-tab-pane label="用户操作日志" name="user">
            <el-table :data="logs" style="width: 100%">
              <el-table-column type="expand">
                <template slot-scope="props">
                  <json-viewer :copyable="true" style="background-color: #e6effb" v-if="props.row.result !== ''" :value="JSON.parse(props.row.result)"></json-viewer>
                  <pre v-else><div style="text-align: center; color: #909399">没有数据</div></pre>
                </template>
              </el-table-column>
              <el-table-column prop="user_id" label="用户" width="300"></el-table-column>
              <el-table-column prop="operation" label="操作"></el-table-column>
              <el-table-column prop="created_at" label="创建于" width="160"></el-table-column>
            </el-table>
            <div class="pagination">
              <el-pagination background layout="prev, pager, next" :current-page.sync="meta.page" :total="meta.total" @current-change="changePage"></el-pagination>
            </div>
          </el-tab-pane>
        </el-tabs>
      </div>
    </div>
  </div>
</template>

<script>
  import api from '../apis'
  import Vue from 'vue'
  import JsonViewer from 'vue-json-viewer'
  import hljs from 'highlight.js';
  import 'highlight.js/styles/github.css';
  Vue.use(JsonViewer);

  const highlightCode = () => {
    const elements = document.querySelectorAll('pre');
    elements.forEach((element) => {
      if (element.children.item(0).classList.contains("bash")) {
        hljs.highlightBlock(element)
      }
    })
  };

  export default {
    name: "Log",
    data() {
      return {
        search: "",
        active: 'pipeline',
        reverse: true,
        classes: ['animated', 'fade-in', 'fast'],
        loading: false,
        params: {
          field: "id",
          scene: "pipeline",
          search: "",
          page: 1
        },
        current: {
          id: null,
          index: null
        },
        logs: [],
        meta: {
          limit: 10,
          page: 1,
          total: 0
        }
      }
    },
    methods: {
      handleClick(tab, event) {
        this.params.scene = tab.name;
        this.params.search = "";
        this.params.page = 1;
        this.fetchLogs();
        window.console.log(tab, event);
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
          let instance = rows.pop();
          // 考虑到展开新行和关闭行都会触发 handleTableExpand 所以需要判断是否是当前行
          if (instance.id === row.id) {
            for (let index = 0; index < this.logs.length; index++) {
              if (this.logs[index].id === instance.id) {
                Vue.set(this.current, 'id', instance.id);
                Vue.set(this.current, 'index', index);
                this.params.field = 'pipeline_record_id';
                this.params.search = instance.id;
                this.params.scene = 'task';
                // 调用 API 获取流水线关联的任务
                api.log.fetch(this.params).then(res => {
                  // 将返回数据设置到对应流水线下的执行步骤属性中
                  Vue.set(this.logs[index], 'steps', res.data);
                  this.$nextTick(() => setTimeout(highlightCode, 0));
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
      fetchLogs () {
        this.loading = true;
        api.log.fetch(this.params).then(res => {
          this.logs = res.data;
          this.meta = res.meta;
        }).catch(err => {
          this.$message.warning(err.message)
        });
        this.loading = false;
      },
      handleClear() {
        // 判断是否有 Pipeline 页面跳转传入的参数
        if (this.$route.query.hasOwnProperty("id")) {
          // 如果有则替换路由
          this.$router.replace("/log");
        }
        this.fetchLogs();
      },
      changePage(page) {
        this.meta.page = page;
        this.params.page = page;
        this.fetchLogs();
      },
    },
    mounted() {
      // 如果存在查询参数则
      if (this.$route.query.hasOwnProperty("id")) {
        this.params.search = this.$route.query.id;
      }
      this.fetchLogs();
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
  .el-tabs__item {
    color: #909399;
  }
  .task-pre {
    width: 100%;
    margin: 15px 0 0 0;
    color: #5e6d82;
    padding: 10px 20px;
    background-color: #e6effb;
  }
</style>
