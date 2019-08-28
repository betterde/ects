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
            <el-table :data="logs" style="width: 100%">
              <el-table-column type="expand">
                <template slot-scope="props">
                  <json-viewer :copyable="true" style="background-color: #e6effb" v-if="props.row.result !== ''" :value="JSON.parse(props.row.result)"></json-viewer>
                  <pre v-else><div style="text-align: center; color: #909399">没有数据</div></pre>
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
          <el-tab-pane label="任务日志" name="task">
            <el-table :data="logs" style="width: 100%">
              <el-table-column type="expand">
                <template slot-scope="props">
                  <json-viewer :copyable="true" style="background-color: #e6effb" v-if="props.row.result !== ''" :value="JSON.parse(props.row.result)"></json-viewer>
                  <pre v-else><div style="text-align: center; color: #909399">没有数据</div></pre>
                </template>
              </el-table-column>
              <el-table-column prop="pipeline_record_id" label="流水线" width="300"></el-table-column>
              <el-table-column prop="node_id" label="节点" width="300"></el-table-column>
              <el-table-column prop="task_id" label="任务" width="300"></el-table-column>
              <el-table-column prop="status" label="结果"></el-table-column>
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
  Vue.use(JsonViewer);

  export default {
    name: "Log",
    data() {
      return {
        search: "",
        active: 'pipeline',
        classes: ['animated', 'fade-in', 'fast'],
        loading: false,
        params: {
          scene: "pipeline",
          search: "",
          page: 1
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
      this.fetchLogs()
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
</style>
