<template>
  <div id="install">
    <div class="panel">
      <el-row>
        <el-col :span="14" :offset="5">
          <div class="panel-heading">
            <h1 class="title">Install Elastic Crontab System</h1>
          </div>
          <el-steps :active="step" finish-status="success" simple style="margin-top: 20px">
            <el-step title="ETCD"></el-step>
            <el-step title="JWT"></el-step>
            <el-step title="DB"></el-step>
            <el-step title="User"></el-step>
            <el-step title="Finish"></el-step>
          </el-steps>
          <div class="panel-body">
            <div class="install-form-container" v-show="step === 0">
              <el-form v-if="env === 'etcd'" :model="config.etcd" :rules="rules.etcd" ref="etcd" label-position="top">
                <el-row :gutter="20">
                  <el-col :span="24">
                    <el-form-item prop="endpoints" label="EndPoints">
                      <el-select v-model="config.etcd.endpoints" multiple filterable allow-create default-first-option
                                 placeholder="请输入 ETCD 的 EndPoints" style="width: 100%"></el-select>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item prop="service" label-width="80px" label="服务节点前缀">
                      <el-input v-model="config.etcd.service" placeholder="用于服务节点注册" clearable></el-input>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item prop="pipeline" label-width="80px" label="流水线前缀">
                      <el-input v-model="config.etcd.pipeline" placeholder="用户保存用户创建的流水线" clearable></el-input>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item prop="config" label-width="80px" label="服务配置前缀">
                      <el-input v-model="config.etcd.config" placeholder="用于保存服务器配置信息" clearable></el-input>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item prop="killer" label-width="80px" label="强杀指令前缀">
                      <el-input v-model="config.etcd.killer" placeholder="用于保存强杀正在执行的流水线ID" clearable></el-input>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item prop="locker" label-width="80px" label="分布式锁前缀">
                      <el-input v-model="config.etcd.locker" placeholder="用于抢占资源的分布式锁" clearable></el-input>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item prop="timeout" label-width="80px" label="超时时间">
                      <el-input v-model="config.etcd.timeout" placeholder="ETCD请求超时时间" clearable></el-input>
                    </el-form-item>
                  </el-col>
                </el-row>
              </el-form>
              <div class="footer">
                <el-button type="primary" plain @click="confirm('etcd')">下一步</el-button>
              </div>
            </div>
            <div class="install-form-container" v-show="step === 1">
              <el-form :model="config.auth" :rules="rules.auth" ref="auth" label-width="100px" label-position="top">
                <el-row :gutter="20">
                  <el-col :span="20">
                    <el-form-item prop="secret" label="Secret">
                      <el-input v-model="config.auth.secret" placeholder="Secret">
                        <el-tooltip slot="append" class="item" effect="dark" content="生成 Secret" placement="top">
                          <el-button icon="el-icon-document" @click="fetchJWTSecret"></el-button>
                        </el-tooltip>
                      </el-input>
                    </el-form-item>
                  </el-col>
                  <el-col :span="4">
                    <el-form-item prop="ttl" label="TTL">
                      <el-input v-model="config.auth.ttl" placeholder="TTL" clearable></el-input>
                    </el-form-item>
                  </el-col>
                </el-row>
              </el-form>
              <div class="footer">
                <el-button type="info" plain @click="back">上一步</el-button>
                <el-button type="primary" plain @click="confirm('auth')">下一步</el-button>
              </div>
            </div>
            <div class="install-form-container" v-show="step === 2">
              <el-form :model="config.database" :rules="rules.database" ref="database">
                <el-col :span="14">
                  <el-form-item prop="host" label="主机">
                    <el-input v-model="config.database.host" placeholder="主机地址" clearable></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="2" style="padding-top: 50px; text-align: center">:</el-col>
                <el-col :span="8">
                  <el-form-item prop="port" label="端口">
                    <el-input v-model="config.database.port" placeholder="端口号"></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="14">
                  <el-form-item prop="name" label="数据库">
                    <el-input v-model="config.database.name" placeholder="数据库名称" clearable></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="2" style="padding-top: 10px"></el-col>
                <el-col :span="8">
                  <el-form-item prop="char" label="字符集">
                    <el-select v-model="config.database.char" placeholder="数据库字符集">
                      <el-option v-for="char in chars" :key="char.value" :label="char.label"
                                 :value="char.value"></el-option>
                    </el-select>
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item prop="user" label="用户名">
                    <el-input v-model="config.database.user" placeholder="数据库用户名" clearable></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item prop="pass" label="密码">
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
              <el-form :model="config.user" :rules="rules.user" ref="user">
                <el-form-item prop="name">
                  <el-input v-model="config.user.name" placeholder="用户名" clearable></el-input>
                </el-form-item>
                <el-form-item prop="email">
                  <el-input v-model="config.user.email" placeholder="邮箱" clearable></el-input>
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
                <el-button type="primary" plain @click="confirm('user')">提交</el-button>
              </div>
            </div>
            <div class="install-form-container" style="text-align: center" v-show="step >= 4">
              <el-progress type="circle" :percentage="100" status="success">
                <i class="el-icon-check" style="color: #409EFF; font-size: 48px"></i>
              </el-progress>
              <h2 style="color: #8c939d; margin: 40px 0;">恭喜你，系统初始化成功！</h2>
              <div class="footer">
                <div class="command">
                  <code v-html="command"></code>
                </div>
                <p>现在复制上面的命令，并关闭 init 服务，在终端运行你复制的命令，就可以开始管理你的定时任务了。</p>
              </div>
            </div>
          </div>
        </el-col>
      </el-row>
      <el-row type="flex" align="bottom">
        <el-col :span="24">
          <footer class="copyright page-copyright-inverse text-center">
            <p class="version">ECTS : {{info.version}}</p>
            <p>Copyright © 2019 Betterde.Inc. All Rights Reserved.</p>
          </footer>
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
          callback(new Error('请输入密码'));
        } else if (value !== this.config.user.pass) {
          callback(new Error('输入的两次密码不匹配'));
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
        loading: false,
        config: {
          database: {
            host: "",
            port: 3306,
            user: "",
            pass: "",
            name: "ects",
            char: "utf8mb4"
          },
          auth: {
            secret: "",
            ttl: 86400
          },
          user: {
            name: "",
            email: "",
            pass: "",
            confirm: ""
          },
          etcd: {
            service: "/ects/nodes",
            pipeline: "/ects/pipelines",
            config: "/ects/config",
            killer: "/ects/killer",
            locker: "/ects/locker",
            timeout: 5,
            endpoints: ["localhost:2379"]
          }
        },
        rules: {
          database: {
            host: [
              {type: "string", required: true, message: '请输入数据库主机地址', trigger: 'blur'}
            ],
            port: [
              {type: "integer", required: true, message: '请输入数据库端口', trigger: 'blur'}
            ],
            user: [
              {type: "string", required: true, message: '请输入数据库用户名', trigger: 'blur'}
            ],
            pass: [
              {type: "string", required: true, message: '请输入数据库密码', trigger: 'blur'}
            ],
            name: [
              {type: "string", required: true, message: '请输入数据库名称', trigger: 'blur'}
            ],
            char: [
              {type: "string", required: true, message: '请选择数据库默认字符集', trigger: 'blur'}
            ]
          },
          auth: {
            secret: [
              {type: "string", required: true, message: '请输入 JWT Secret', trigger: 'blur'}
            ],
            ttl: [
              {type: "number", required: true, message: '请输入 JWT TTL', trigger: 'blur'}
            ]
          },
          user: {
            name: [
              {type: "string", required: true, message: '请输入你的用户名', trigger: 'blur'}
            ],
            email: [
              {type: 'email', message: '请输入你的邮箱地址', trigger: ['blur', 'change']}
            ],
            pass: [
              {validator: validatePass, trigger: 'blur'}
            ],
            confirm: [
              {validator: confirmPass, trigger: 'blur'}
            ]
          },
          etcd: {
            endpoints: [
              {type: "array", required: true, message: '请输入 ETCD 服务的 EndPoints', trigger: 'blur'}
            ],
            service: [
              {type: "string", required: true, message: '请输入用于服务注册和发现的前缀', trigger: 'blur'}
            ],
            pipeline: [
              {type: "string", required: true, message: '请输入用户保存需要执行流水线的前缀', trigger: 'blur'}
            ],
            killer: [
              {type: "string", required: true, message: '请输入需要保存强杀流水线的前缀', trigger: 'blur'}
            ],
            locker: [
              {type: "string", required: true, message: '请输入用于保存分布式锁的前缀', trigger: 'blur'}
            ],
            config: [
              {type: "string", required: true, message: '请输入用于保存服务配置信息的前缀', trigger: 'blur'}
            ],
            timeout: [
              {type: "number", required: true, message: '请设置 ETCD 请求的超时时间', trigger: 'blur'}
            ],
          }
        },
        chars: [
          {
            label: "utf8mb4",
            value: "utf8mb4"
          }
        ],
        command: "ects master --etcd"
      }
    },
    methods: {
      /**
       * Back
       */
      back() {
        if (this.step >= 1) {
          this.step -= 1;
        }
      },
      /**
       * Form validate
       */
      confirm(form) {
        if (form === 'index') {
          window.location.href = '/';
        } else {
          this.$refs[form].validate((valid) => {
            if (valid) {
              if (form === 'database') {
                api.system.database(this.config.database).then(res => {
                  if (res.data.exist) {
                    this.$confirm('数据库已经存在，如果继续则原有数据将被清空!', '警告', {
                      confirmButtonText: '继续',
                      cancelButtonText: '取消',
                      type: 'warning'
                    }).then(() => {
                      this.next();
                    }).catch(() => {
                      this.$message({
                        type: 'info',
                        message: '已经取消'
                      });
                    });
                  } else {
                    this.next();
                  }
                })
              } else {
                this.next();
              }
            } else {
              return false;
            }
          });
        }
      },
      /**
       * Next
       */
      next() {
        if (this.step < 3) {
          this.step += 1;
        } else {
          this.submit();
        }
      },
      /**
       * 提交配置信
       */
      submit() {
        api.system.initialize(this.config).then(res => {
          this.$message({
            message: res.message,
            type: 'success'
          });
          this.step += 2;
          this.generateCommand()
        }).catch(err => {
          this.$notify.error({
            title: "Error",
            message: err.data.message,
          });
        });
      },
      /**
       * 获取 JSON Web Token secret
       */
      fetchJWTSecret() {
        api.system.secret().then(res => {
          this.config.auth.secret = res.data
        }).catch(err => {
          this.$notify.error({
            title: "Error",
            message: err.message,
          })
        })
      },
      /**
       * 生成运行命令
       */
      generateCommand() {
        let endpoints = this.config.etcd.endpoints.join(',');
        this.command = `ects master --etcd=${endpoints} --name='master' --desc='master node' --config=${this.config.etcd.config}`;
      }
    },
    /**
     * 获取系统版本信息
     */
    mounted() {
      api.system.fetch().then(res => {
        this.info.version = res.data.version;
      });
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

        .el-progress__text {
          .el-icon-check {
            font-size: 48px;
          }
        }

        .system-info {
          margin-bottom: 30px;
        }

        .footer {
          text-align: center;
          margin: 20px 0;
          p {
            padding: 20px 20px 0 20px;
            color: #C0C4CC;
          }
        }

        .install-form-container {
          padding: 0 10%;

          .el-select {
            width: 100%;
          }

          .line {
            padding-top: 60px;
            text-align: center;
          }

          .el-input-number {
            width: auto;
          }
        }
      }

      .command {
        padding: 20px;
        font-size: 14px;
        border-radius: 4px;
        background-color: #fafafa;
      }

      .copyright {
        margin-top: 40px;
        font-size: 14px;
        color: #C0C4CC;

        p {
          margin: 10px 0;
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
