<template>
  <div class="main-content">
    <div class="panel">
      <el-row :gutter="20">
       <el-col :span="6">
         <div class="card" :class="classes">
           <div class="card_header">
             <h4>主要节点</h4>
           </div>
           <div class="card_body">
             <h1 class="number">{{ nodes.master }}</h1>
           </div>
           <div class="card_footer">
             <h4>在线数量</h4>
           </div>
         </div>
       </el-col>
       <el-col :span="6">
         <div class="card" :class="classes">
           <div class="card_header">
             <h4>工作节点</h4>
           </div>
           <div class="card_body">
             <h1 class="number">{{ nodes.worker }}</h1>
           </div>
           <div class="card_footer">
             <h4>在线数量</h4>
           </div>
         </div>
       </el-col>
       <el-col :span="6">
         <div class="card" :class="classes">
           <div class="card_header">
             <h4>流水线</h4>
           </div>
           <div class="card_body">
             <h1 class="number">{{ pipelines.total }}</h1>
           </div>
           <div class="card_footer">
             <h4>调度中</h4>
           </div>
         </div>
       </el-col>
       <el-col :span="6">
         <div class="card" :class="classes">
           <div class="card_header">
             <h4>定时任务</h4>
           </div>
           <div class="card_body">
             <h1 class="number">{{ failtures.total }}</h1>
           </div>
           <div class="card_footer">
             <h4>失败次数</h4>
           </div>
         </div>
       </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
  import api from '../apis'
  export default {
    name: "Dashboard",
    data() {
      return {
        classes: ['animated', 'fade-in', 'fast'],
        nodes: {
          master: 0,
          worker: 0
        },
        pipelines: {
          total: 0
        },
        failtures: {
          total: 0
        }
      }
    },
    mounted() {
      api.dashboard.fetchNodes().then(res => {
        this.nodes = res.data;
      }).catch(err => {
        this.$message.error({
          offset: 95,
          message: err.message
        })
      });
      api.dashboard.fetchPipelines().then(res => {
        this.pipelines.total = res.data;
      }).catch(err => {
        this.$message.error({
          offset: 95,
          message: err.message
        })
      });
      api.dashboard.fetchFailtures().then(res => {
        this.failtures.total = res.data;
      }).catch(err => {
        this.$message.error({
          offset: 95,
          message: err.message
        })
      });
    }
  }
</script>

<style lang="scss">
  .card {
    width: 100%;
    padding: 40px;
    color: #909399;
    background: white;
    border-radius: 4px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    .card_header {
      width: 100%;
      text-align: center;
    }
    .card_body {
      text-align: center;
      .number {
        color: #5f6268;
        font-size: 100px;
      }
    }
    .card_footer {
      text-align: center;
    }
  }
</style>
