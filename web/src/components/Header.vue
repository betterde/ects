<template>
  <el-header>
    <el-row :gutter="20">
      <el-col :span="4">
        <div class="grid-content">
          <h2 style="line-height: 36px; color: #8c939d">ECTS</h2>
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
</template>

<script>
    export default {
        name: "Header"
    }
</script>

<style scoped>

</style>
