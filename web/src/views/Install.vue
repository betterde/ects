<template>
  <div id="install">
    <div class="panel">
      <el-row>
        <el-col :span="12" :offset="6">
          <div class="panel-heading">
            <h1 class="title">Install Elastic Crontab System</h1>
            <p class="version">Version: {{info.version}}</p>
          </div>
          <el-steps :active="step" finish-status="success" simple style="margin-top: 20px">
            <el-step title="Env"></el-step>
            <el-step title="Service"></el-step>
            <el-step title="DB"></el-step>
            <el-step title="Manager"></el-step>
            <el-step title="Finish"></el-step>
          </el-steps>
          <div class="panel-body">
            <div v-show="step === 0">
              <div class="system-info">
                <div class="env-select" style="text-align: center; margin: 0 0 20px 0">
                  <el-radio v-model="env" label="etcd" border>ETCD</el-radio>
                  <el-radio v-model="env" label="yaml" border>YAML</el-radio>
                </div>
                <el-form v-if="env === 'etcd'" :model="config.etcd" :rules="serviceRules" ref="service" label-position="left">
                  <el-col :span="24">
                    <el-form-item prop="host">
                      <el-select v-model="config.etcd.endpoints" multiple filterable allow-create default-first-option placeholder="请输入ETCD的Endpoints" style="width: 100%"></el-select>
                    </el-form-item>
                  </el-col>
                </el-form>

                <el-alert v-if="env === 'yaml'" title="配置文件写入权限" :type="info.permission" :closable="false"
                          description="如您运行时未指定配置文件路径，系统将使用 /etc/ects/ects.yaml 作为配置文件的默认路径" show-icon></el-alert>
              </div>
              <div class="footer">
                <el-button type="primary" plain @click="confirm('permission')" :disabled="info.permission === 'error'">
                  下一步
                </el-button>
              </div>
            </div>
            <div class="install-form-container" v-show="step === 1">
              <el-form :model="config.service" :rules="serviceRules" ref="service" label-width="100px"
                       label-position="top">
                <el-col :span="14">
                  <el-form-item prop="host">
                    <el-input v-model="config.service.host" placeholder="服务器地址"></el-input>
                  </el-form-item>
                </el-col>
                <el-col class="line" :span="2">:</el-col>
                <el-col :span="8">
                  <el-form-item prop="port">
                    <el-input v-model="config.service.port" placeholder="监听端口"></el-input>
                  </el-form-item>
                </el-col>
              </el-form>
              <div class="footer">
                <el-button type="info" plain @click="back">上一步</el-button>
                <el-button type="primary" plain @click="confirm('service')">下一步</el-button>
              </div>
            </div>
            <div class="install-form-container" v-show="step === 2">
              <el-form :model="config.database" :rules="databaseRules" ref="database">
                <el-col :span="14">
                  <el-form-item prop="host">
                    <el-input v-model="config.database.host" placeholder="数据库主机地址"></el-input>
                  </el-form-item>
                </el-col>
                <el-col class="line" :span="2">:</el-col>
                <el-col :span="8">
                  <el-form-item prop="port">
                    <el-input v-model="config.database.port" placeholder="数据库端口"></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="14">
                  <el-form-item prop="name">
                    <el-input v-model="config.database.name" placeholder="数据库名称"></el-input>
                  </el-form-item>
                </el-col>
                <el-col class="line" :span="2"></el-col>
                <el-col :span="8">
                  <el-form-item prop="char">
                    <el-select v-model="config.database.char" placeholder="请选择字符类型">
                      <el-option v-for="char in chars" :key="char.value" :label="char.label"
                                 :value="char.value"></el-option>
                    </el-select>
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item prop="user">
                    <el-input v-model="config.database.user" placeholder="数据库用户名"></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item prop="pass">
                    <el-input v-model="config.database.pass" placeholder="数据库密码" show-password></el-input>
                  </el-form-item>
                </el-col>
              </el-form>
              <div class="footer">
                <el-button type="info" plain @click="back">上一步</el-button>
                <el-button type="primary" plain @click="confirm('database')">下一步</el-button>
              </div>
            </div>
            <div class="install-form-container" v-show="step === 3">
              <el-form :model="config.user" :rules="userRules" ref="user">
                <el-form-item prop="name">
                  <el-input v-model="config.user.name" placeholder="昵称"></el-input>
                </el-form-item>
                <el-form-item prop="email">
                  <el-input v-model="config.user.email" placeholder="邮箱"></el-input>
                </el-form-item>
                <el-form-item prop="pass">
                  <el-input v-model="config.user.pass" placeholder="密码" show-password></el-input>
                </el-form-item>
                <el-form-item prop="confirm">
                  <el-input v-model="config.user.confirm" placeholder="确认密码" show-password></el-input>
                </el-form-item>
              </el-form>
              <div class="footer">
                <el-button type="info" plain @click="back">上一步</el-button>
                <el-button type="primary" plain @click="confirm('user')">确认安装</el-button>
              </div>
            </div>
            <div class="install-form-container" style="text-align: center" v-show="step === 4">
              <el-progress type="circle" :percentage="100" status="text">
                <i class="el-icon-check" style="color: #409EFF; font-size: 48px"></i>
              </el-progress>
              <h2 style="color: #8c939d; margin: 40px 0;">系统初始化完成</h2>
              <div class="footer">
                <el-button type="primary" @click="confirm('index')">进入后台</el-button>
              </div>
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
    name: "Install",
    data() {
      let validatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入密码'));
        } else {
          if (this.config.user.confirm !== '') {
            this.$refs.user.validateField('confirm');
          }
          callback();
        }
      };
      let confirmPass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请再次输入密码'));
        } else if (value !== this.config.user.pass) {
          callback(new Error('两次输入密码不一致!'));
        } else {
          callback();
        }
      };
      return {
        info: {
          version: "",
          permission: "success"
        },
        step: 0,
        env: "etcd",
        config: {
          service: {
            host: "0.0.0.0",
            port: 9701
          },
          database: {
            host: "localhost",
            port: 3306,
            user: "root",
            pass: "George@1994",
            name: "ects",
            char: "utf8mb4"
          },
          auth: {
            secret: "",
            ttl: 86400
          },
          user: {
            name: "George",
            email: "george@betterde.com",
            pass: "George@1994",
            confirm: "George@1994"
          },
          etcd: {
            endpoints: []
          }
        },
        serviceRules: {
          host: [
            {type: "string", required: true, message: '请输入服务运行时绑定的IP', trigger: 'blur'}
          ],
          port: [
            {type: "integer", required: true, message: '请输入服务运行时监听的端口', trigger: 'blur'}
          ]
        },
        databaseRules: {
          host: [
            {type: "string", required: true, message: '请输入服务MySQL主机地址', trigger: 'blur'}
          ],
          port: [
            {type: "integer", required: true, message: '请输入服务MySQL主机端口', trigger: 'blur'}
          ],
          user: [
            {type: "string", required: true, message: '请输入服务MySQL用户名', trigger: 'blur'}
          ],
          pass: [
            {type: "string", required: true, message: '请输入服务MySQL密码', trigger: 'blur'}
          ],
          name: [
            {type: "string", required: true, message: '请输入服务MySQL数据库名称', trigger: 'blur'}
          ],
          char: [
            {type: "string", required: true, message: '请选择数据库字符类型', trigger: 'blur'}
          ]
        },
        userRules: {
          name: [
            {type: "string", required: true, message: '请输入管理员姓名', trigger: 'blur'}
          ],
          email: [
            {type: "string", required: true, message: '请输入管理员姓名', trigger: 'blur'},
            {type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change']}
          ],
          pass: [
            {validator: validatePass, trigger: 'blur'}
          ],
          confirm: [
            {validator: confirmPass, trigger: 'blur'}
          ]
        },
        chars: [
          {
            label: "utf8mb4",
            value: "utf8mb4"
          }
        ]
      }
    },
    methods: {
      /**
       * 后退
       */
      back() {
        if (this.step >= 1) {
          this.step -= 1;
        }
      },
      /**
       * 表单验证
       */
      confirm(form) {
        if (form === 'permission') {
          this.next();
        } else if (form === 'index') {
          window.location.href = '/';
        } else {
          this.$refs[form].validate((valid) => {
            if (valid) {
              this.next();
            } else {
              return false;
            }
          });
        }
      },
      /**
       * 进入下一步
       */
      next() {
        if (this.step < 3) {
          this.step += 1;
        } else {
          this.submit();
        }
      },
      /**
       * 提交配置信息
       */
      submit() {
        api.system.install(this.config).then(res => {
          if (res.code !== 200) {
            this.$notify.error({
              title: "错误",
              message: res.message,
            });
          } else {
            this.$message({
              message: '恭喜你，系统安装成功',
              type: 'success'
            });
            this.step += 1;
            localStorage.setItem("access_token", res.data.access_token);
          }
        });
      }
    },
    /**
     * 获取权限信息
     */
    mounted() {
      api.system.fetch().then(res => {
        this.info.version = res.version;
        this.info.permission = res.permission ? 'success' : 'error';
      }).catch(err => {
        window.console.log(err);
      })
    }
  }
</script>

<style lang="scss">
  .transition-box {
    margin-bottom: 10px;
    width: 200px;
    height: 100px;
    border-radius: 4px;
    background-color: #409EFF;
    text-align: center;
    color: #fff;
    padding: 40px 20px;
    box-sizing: border-box;
    margin-right: 20px;
  }

  #install {
    font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
    height: 100%;
    text-align: center;
    background-color: #f0f2f5;
    display: flex;
    justify-content: center;
    align-items: Center;

    .panel {
      min-height: 600px;
      width: 100%;

      .panel-heading {
        padding: 40px;

        .title {
          font-size: 48px;
          color: #8c939d;
        }

        .version {
          margin-top: 30px;
          color: #C0C4CC
        }

      }

      .panel-body {
        text-align: left;
        width: 100%;
        margin-top: 40px;
        border-radius: 4px;
        padding: 40px 40px 18px 40px;
        background-color: #FFFFFF;
        box-shadow: 1px 1px 3px rgba(0, 21, 41, .08);

        .system-info {
          margin-bottom: 30px;
        }

        .footer {
          text-align: center;
          margin: 20px 0;
        }

        .install-form-container {
          padding: 0 10%;

          .el-select {
            width: 100%;
          }

          .line {
            padding: 10px 0;
            text-align: center;
          }

        }
      }
    }

    .main-content {
      min-height: 200px;
      width: 100%;
      border-radius: 4px;
      background-color: #FFFFFF;
    }

    .nav-step {
      padding: 10px;
    }

  }
</style>
