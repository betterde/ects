<template>
  <div class="main-content">
    <div class="panel">
      <div class="panel-header" :class="classes">
        <div class="panel-tools">
          <el-row :gutter="20">
            <el-col :span="8">
              <el-input placeholder="在这里输入要搜索的内容，按下回车进行搜索" clearable>
                <i slot="prefix" class="el-input__icon el-icon-search"></i>
              </el-input>
            </el-col>
            <el-col :span="16" style="text-align: right">
            </el-col>
          </el-row>
        </div>
      </div>
      <div class="panel-body" :class="classes">
        <el-tabs v-model="active" tab-position="left" class="setting-menu">
          <el-tab-pane label="系统通知" name="notification">
            <el-form :model="notification" :rules="rules" ref="setting">
              <el-col :span="24">
                <el-form-item prop="url">
                  <el-input v-model="notification.url" placeholder="后台URL用于邮件内链接跳转到ECTS后台，如: https://ects.betterde.com"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="10">
                <el-form-item prop="host">
                  <el-input v-model="notification.host" placeholder="主机地址"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="3">
                <el-form-item prop="port" label-width="0">
                  <el-input v-model.number="notification.port" placeholder="端口"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item prop="char">
                  <el-select v-model="notification.protocol" placeholder="请选择协议">
                    <el-option key="smtp" label="SMTP" value="smtp"></el-option>
                    <el-option key="pop3" label="POP3" value="pop3"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="5">
                <el-form-item prop="char">
                  <el-select v-model="notification.encryption" placeholder="请选择加密类型">
                    <el-option key="tls" label="TLS" value="tls"></el-option>
                    <el-option key="ssl" label="SSL" value="ssl"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item prop="user">
                  <el-input v-model="notification.user" autocomplete="off" placeholder="用户名"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item prop="pass">
                  <el-input v-model="notification.pass" autocomplete="off" placeholder="密码" show-password></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item prop="pass" style="text-align: right">
                  <el-popover placement="top" width="300" v-model="visible">
                    <el-form :model="sendmail.params" :rules="sendmail.rules" ref="sendmail">
                      <el-form-item prop="email">
                        <el-input v-model="sendmail.params.email" placeholder="请输入接收测试邮件的邮箱"></el-input>
                      </el-form-item>
                    </el-form>
                    <div style="text-align: right; margin: 0">
                      <el-button size="mini" type="text" @click="closePopover">取消</el-button>
                      <el-button type="primary" size="mini" @click="send">确定</el-button>
                    </div>
                    <el-button type="info" slot="reference" plain style="margin-right: 20px">发送测试邮件</el-button>
                  </el-popover>
                  <el-button type="primary" plain @click="submit">保存</el-button>
                </el-form-item>
              </el-col>
            </el-form>
          </el-tab-pane>
          <el-tab-pane label="系统升级" name="upgrade">
            <el-upload class="upgrade" drag action="https://jsonplaceholder.typicode.com/posts/" multiple>
              <i class="el-icon-upload"></i>
              <div class="el-upload__text">请将最新版的ECTS二进制文件拖拽到这里或 <em>点击上传</em></div>
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
  import api from '../apis'

  export default {
    name: "Setting",
    data() {
      return {
        sendmail: {
          params: {
            email: ''
          },
          rules: {
            email: [
              { required: true, message: '请输入邮件地址', trigger: 'blur' },
              { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }
            ]
          }
        },
        visible: false,
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
          url: '',
          host: "",
          port: 25,
          user: '',
          pass: '',
          name: '',
          protocol: '',
          encryption: ''
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
      submit() {
        this.$refs.setting.validate((valid) => {
          if (valid) {
            api.setting.updateNotification(this.notification).then(res => {
              this.notification = res.data;
              this.$message.success({
                offset: 95,
                message: res.message
              })
            }).catch(err => {
              this.$message.error({
                offset: 95,
                message: err.message
              })
            })
          } else {
            return false;
          }
        });
      },
      send() {
        api.setting.sendMail(this.sendmail.params).then(res => {
          this.$message.success({
            offset: 95,
            message: res.message
          });
          this.$refs.sendmail.resetFields();
          this.visible = false;
        }).catch(err => {
          this.$message.error({
            offset: 95,
            message: err.message
          })
        })
      },
      closePopover() {
        this.$refs.sendmail.resetFields();
        this.visible = false;
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      }
    },
    mounted() {
      api.setting.fetchNotification().then(res => {
        this.notification = res.data;
      }).catch(err => {
        this.$message.error({
          offset: 95,
          message: err.message
        })
      })
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
