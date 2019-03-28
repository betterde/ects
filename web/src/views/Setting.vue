<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-body">
        <el-tabs v-model="active" tab-position="left" class="setting-menu">
          <el-tab-pane label="基本设置" name="basic">
            <el-form :model="basic" :rules="rules" ref="basic" label-width="100px" class="demo-ruleForm">
              <el-col :span="18">
                <el-form-item label="系统名称" prop="name" style="text-align: left">
                  <el-input v-model="basic.name"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="用户注册" prop="type">
                  <el-switch v-model="basic.register"></el-switch>
                </el-form-item>
              </el-col>
              <el-col :span="18">
              </el-col>
              <el-form-item label="特殊资源" prop="resource">
                <el-radio-group v-model="basic.resource">
                  <el-radio label="线上品牌商赞助"></el-radio>
                  <el-radio label="线下场地免费"></el-radio>
                </el-radio-group>
              </el-form-item>
              <el-form-item label="活动形式" prop="desc">
                <el-input type="textarea" v-model="basic.desc"></el-input>
              </el-form-item>
            </el-form>
          </el-tab-pane>
          <el-tab-pane label="系统日志" name="log">
          </el-tab-pane>
          <el-tab-pane label="异常通知" name="notification">
            <el-form :model="notification" :rules="rules" ref="database">
              <el-col :span="18">
                <el-form-item prop="host">
                  <el-input v-model="notification.host" placeholder="主机地址"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item prop="port" label-width="0">
                  <el-input v-model="notification.port" placeholder="端口"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item prop="char">
                  <el-select v-model="notification.char" placeholder="请选择服务类型">
                    <el-option key="smtp" label="SMTP" value="smtp"></el-option>
                    <el-option key="pop3" label="POP3" value="pop3"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item prop="char">
                  <el-select v-model="notification.encryption" placeholder="请选择加密类型">
                    <el-option key="tls" label="TLS" value="tls"></el-option>
                    <el-option key="ssl" label="SSL" value="ssl"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item prop="user">
                  <el-input v-model="notification.user" placeholder="用户名"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item prop="pass">
                  <el-input v-model="notification.pass" placeholder="密码" show-password></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item prop="pass" style="text-align: right">
                  <el-button type="primary" plain>更新配置</el-button>
                </el-form-item>
              </el-col>
            </el-form>
          </el-tab-pane>
          <el-tab-pane label="版本升级" name="upgrade">
            <el-upload class="upgrade" drag action="https://jsonplaceholder.typicode.com/posts/" multiple>
              <i class="el-icon-upload"></i>
              <div class="el-upload__text">将最新的ECTS可执行文件拖到此处，或<em>点击上传</em></div>
            </el-upload>
          </el-tab-pane>
          <el-tab-pane label="系统信息" name="info">系统信息</el-tab-pane>
          <el-tab-pane label="获得帮助" name="service">获得帮助</el-tab-pane>
        </el-tabs>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Setting",
    data() {
      return {
        active: "basic",
        basic: {
          name: 'ECTS',
          register: "",
          region: '',
          date1: '',
          date2: '',
          delivery: false,
          type: [],
          resource: '',
          desc: ''
        },
        notification: {
          host: "smtp.mailtrap.io",
          port: 25,
          user: "",
          pass: "",
          name: "",
          char: "",
          encryption: ""
        },
        rules: {
          name: [
            { required: true, message: '请输入活动名称', trigger: 'blur' },
            { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
          ],
          region: [
            { required: true, message: '请选择活动区域', trigger: 'change' }
          ],
          date1: [
            { type: 'date', required: true, message: '请选择日期', trigger: 'change' }
          ],
          date2: [
            { type: 'date', required: true, message: '请选择时间', trigger: 'change' }
          ],
          type: [
            { type: 'array', required: true, message: '请至少选择一个活动性质', trigger: 'change' }
          ],
          resource: [
            { required: true, message: '请选择活动资源', trigger: 'change' }
          ],
          desc: [
            { required: true, message: '请填写活动形式', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            alert('submit!');
          } else {
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      }
    }
  }
</script>

<style lang="scss">
  .el-row {
    height: 100%;
    margin-bottom: 20px;
    &:last-child {
      margin-bottom: 0;
    }
  }
  .setting-menu {
    .el-tabs__item {
      color: #909399;
      height: 56px;
      width: 140px;
      line-height: 56px;
    }
    .el-tabs__item.is-left {
      text-align: center;
    }
  }
  #pane-notification {
    .el-select {
      width: 100%;
    }
  }

  .upgrade {
    height: 100%;
    .el-upload {
      height: 100%;
      width: 100%;
      .el-upload-dragger {
        width: 100%;
        height: 277px;
        .el-icon-upload {
          margin: 80px 0 16px;
        }
      }
    }
  }

  .el-col {
    height: 100%;
    border-radius: 4px;
  }
  .setting {
    height: 100%;
    min-height: 500px;
    background: #d3dce6;
  }
  .row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
  }
</style>
