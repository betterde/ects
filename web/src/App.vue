<template>
  <div id="app">
    <el-container>
      <el-header>
        <el-row :gutter="20">
          <el-col :span="4">
            <div class="grid-content">
              <h2 style="line-height: 36px; color: #8c939d">Elastic Crontab System</h2>
            </div>
          </el-col>
          <el-col :span="16">
            <el-menu :default-active="activeIndex" class="el-menu-nav" mode="horizontal" @select="handleSelect">
              <el-menu-item index="1">系统概览</el-menu-item>
              <el-menu-item index="2">任务管理</el-menu-item>
              <el-menu-item index="3">节点管理</el-menu-item>
              <el-menu-item index="4">团队管理</el-menu-item>
              <el-menu-item index="5">人员管理</el-menu-item>
              <el-submenu index="6">
                <template slot="title">系统管理</template>
                <el-menu-item index="6-1">系统日志</el-menu-item>
                <el-menu-item index="6-2">任务日志</el-menu-item>
                <el-menu-item index="6-3">通知管理</el-menu-item>
              </el-submenu>
            </el-menu>
          </el-col>
          <el-col :span="4">
            <el-dropdown>
              <span class="el-dropdown-link">
                <div class="avatar grid-content"></div>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item>消息通知</el-dropdown-item>
                <el-dropdown-item>个人信息</el-dropdown-item>
                <el-dropdown-item>退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </el-col>
        </el-row>
      </el-header>
      <el-main>
        <el-row :gutter="20">
          <el-col :span="18" :offset="3"><div class="main-content"></div></el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<script>
  export default {
    name: 'app',
    data() {
      return {
        activeIndex: '1',
      }
    },
    methods: {
      handleSelect(key, keyPath) {
        window.console.log(key, keyPath);
      }
    }
  }
</script>

<style lang="scss">
  #app {
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
    height: 100%;
    text-align: center;
    background-color: #f0f2f5;
    .el-header {
      background-color: #FFFFFF;
      box-shadow: 0 1px 4px rgba(0, 21, 41, .08);
      .el-dropdown {
        float: right;
      }
      .avatar {
        height: 36px;
        width: 36px;
        background-color: #8c939d;
        float: right;
      }
    }

    .el-menu-nav {
      display: table;
      margin: auto;
      align-items: center;
      text-align: center;

      &.el-menu--horizontal {
        border: none;
      }
    }

    .el-container {
      height: 100%;
      min-height: 100%;
      .grid-content {
        margin: 12px 0;
        border-radius: 4px;
        min-height: 36px;
        text-align: left;
      }
    }

    .el-main {
      height: auto;
      display: block;
      .el-row {
        height: 100%;
      }
      .el-col {
        height: 100%;
      }
    }
  }

  .el-col {
    border-radius: 4px;
  }

  .main-content {
    height: 100%;
    border-radius: 4px;
    background-color: #FFFFFF;
    box-shadow: 1px 1px 3px rgba(0, 21, 41, .08);
  }
</style>
