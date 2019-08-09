<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header" :class="classes">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="16">
            </el-col>
            <el-col :span="8">
              <el-input placeholder="在这里搜索" v-model="params.search"><i slot="prefix" class="el-input__icon el-icon-search"></i></el-input>
            </el-col>
          </el-row>
        </div>
      </div>
      <div class="panel-body" :class="classes">
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
        classes: ['animated', 'fade-in', 'fast'],
        params: {
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
      changePage(page) {
        this.meta.page = page;
        this.params.page = page;
        this.fetchLogs();
      },
    },
    mounted() {
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

</style>
