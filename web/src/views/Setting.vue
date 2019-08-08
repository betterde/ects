<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-body" :class="classes">
        <el-tabs v-model="active" tab-position="left" class="setting-menu">
          <el-tab-pane label="系统通知" name="notification">
            <el-form :model="notification" :rules="rules" ref="database">
              <el-col :span="10">
                <el-form-item prop="host">
                  <el-input v-model="notification.host" placeholder="Host"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="2">
                <el-form-item prop="port" label-width="0">
                  <el-input v-model="notification.port" placeholder="Port"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item prop="char">
                  <el-select v-model="notification.char" placeholder="Service protocol">
                    <el-option key="smtp" label="SMTP" value="smtp"></el-option>
                    <el-option key="pop3" label="POP3" value="pop3"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item prop="char">
                  <el-select v-model="notification.encryption" placeholder="Encryption type">
                    <el-option key="tls" label="TLS" value="tls"></el-option>
                    <el-option key="ssl" label="SSL" value="ssl"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item prop="user">
                  <el-input v-model="notification.user" placeholder="User"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item prop="pass">
                  <el-input v-model="notification.pass" placeholder="Password" show-password></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item prop="pass" style="text-align: right">
                  <el-button type="primary" plain>Upgrade</el-button>
                </el-form-item>
              </el-col>
            </el-form>
          </el-tab-pane>
          <el-tab-pane label="系统升级" name="upgrade">
            <el-upload class="upgrade" drag action="https://jsonplaceholder.typicode.com/posts/" multiple>
              <i class="el-icon-upload"></i>
              <div class="el-upload__text">Drag the latest ECTS binary executable file here，or <em>click to upload</em></div>
            </el-upload>
          </el-tab-pane>
          <el-tab-pane label="系统信息" name="info">
            <div class="information">
              <h1 class="information-title">
                系统信息
              </h1>
            </div>
          </el-tab-pane>
          <el-tab-pane label="系统服务" name="service">
            <div class="service-info">
              <h1 class="service-info-title">
                获得帮助
              </h1>
            </div>
          </el-tab-pane>
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
        classes: ['animated', 'fade-in', 'fast'],
        active: "notification",
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
        height: 220px;
        .el-icon-upload {
          margin: 80px 0 16px;
        }
      }
    }
  }
  .information {
    height: 220px;
    .information-title {
      color: #8c939d;
      text-align: center;
    }
  }
  .service-info {
    height: 220px;
    .service-info-title {
      color: #8c939d;
      text-align: center;
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
