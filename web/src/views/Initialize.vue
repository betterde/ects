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
            <el-step title="Service"></el-step>
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
                                 placeholder="Endpoints" style="width: 100%"></el-select>
                    </el-form-item>
                  </el-col>
                  <el-col :span="6">
                    <el-form-item prop="service" label-width="80px" label="Service key">
                      <el-input v-model="config.etcd.service" placeholder="Service discover key"></el-input>
                    </el-form-item>
                  </el-col>
                  <el-col :span="6">
                    <el-form-item prop="pipeline" label-width="80px" label="Pipeline key">
                      <el-input v-model="config.etcd.pipeline" placeholder="Pipeline key"></el-input>
                    </el-form-item>
                  </el-col>
                  <el-col :span="6">
                    <el-form-item prop="config" label-width="80px" label="Config key">
                      <el-input v-model="config.etcd.config" placeholder="System config key"></el-input>
                    </el-form-item>
                  </el-col>
                  <el-col :span="6">
                    <el-form-item prop="timeout" label-width="80px" label="Timeout">
                      <el-input v-model="config.etcd.timeout" label="Timeout"></el-input>
                    </el-form-item>
                  </el-col>
                </el-row>
              </el-form>
              <div class="footer">
                <el-button type="primary" plain @click="confirm('etcd')">Fore</el-button>
              </div>
            </div>
            <div class="install-form-container" v-show="step === 1">
              <el-form :model="config.service" :rules="rules.service" ref="service" label-width="100px"
                       label-position="top">
                <el-row :gutter="10">
                  <el-col :span="14">
                    <el-form-item prop="host" label="Host">
                      <el-input v-model="config.service.host" placeholder="Host"></el-input>
                    </el-form-item>
                  </el-col>
                  <el-col class="line" :span="2">:</el-col>
                  <el-col :span="8">
                    <el-form-item prop="port" label="Port">
                      <el-input v-model="config.service.port" placeholder="Port"></el-input>
                    </el-form-item>
                  </el-col>
                </el-row>
              </el-form>
              <div class="footer">
                <el-button type="info" plain @click="back">Back</el-button>
                <el-button type="primary" plain @click="confirm('service')">Fore</el-button>
              </div>
            </div>
            <div class="install-form-container" v-show="step === 2">
              <el-form :model="config.auth" :rules="rules.auth" ref="auth" label-width="100px" label-position="top">
                <el-row :gutter="20">
                  <el-col :span="20">
                    <el-form-item prop="secret" label="Secret">
                      <el-input v-model="config.auth.secret" placeholder="Secret">
                        <el-tooltip slot="append" class="item" effect="dark" content="Generate secret" placement="top">
                          <el-button icon="el-icon-document" @click="fetchJWTSecret"></el-button>
                        </el-tooltip>
                      </el-input>
                    </el-form-item>
                  </el-col>
                  <el-col :span="4">
                    <el-form-item prop="ttl" label="TTL">
                      <el-input v-model="config.auth.ttl" placeholder="TTL"></el-input>
                    </el-form-item>
                  </el-col>
                </el-row>
              </el-form>
              <div class="footer">
                <el-button type="info" plain @click="back">Back</el-button>
                <el-button type="primary" plain @click="confirm('auth')">Fore</el-button>
              </div>
            </div>
            <div class="install-form-container" v-show="step === 3">
              <el-form :model="config.database" :rules="rules.database" ref="database">
                <el-col :span="14">
                  <el-form-item prop="host" label="Host">
                    <el-input v-model="config.database.host" placeholder="Host"></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="2" style="padding-top: 50px; text-align: center">:</el-col>
                <el-col :span="8">
                  <el-form-item prop="port" label="Port">
                    <el-input v-model="config.database.port" placeholder="Port"></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="14">
                  <el-form-item prop="name" label="Database">
                    <el-input v-model="config.database.name" placeholder="Database name"
                              @blur="fetchDatabaseIsExist"></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="2" style="padding-top: 10px"></el-col>
                <el-col :span="8">
                  <el-form-item prop="char" label="Charset">
                    <el-select v-model="config.database.char" placeholder="Database charset">
                      <el-option v-for="char in chars" :key="char.value" :label="char.label"
                                 :value="char.value"></el-option>
                    </el-select>
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item prop="user" label="Username">
                    <el-input v-model="config.database.user" placeholder="Database username"></el-input>
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item prop="pass" label="Password">
                    <el-input v-model="config.database.pass" placeholder="Database password" show-password></el-input>
                  </el-form-item>
                </el-col>
              </el-form>
              <div class="footer">
                <el-button type="info" plain @click="back">Back</el-button>
                <el-button type="primary" plain @click="confirm('database')">Fore</el-button>
              </div>
            </div>
            <div class="install-form-container" v-show="step === 4">
              <el-form :model="config.user" :rules="rules.user" ref="user">
                <el-form-item prop="name">
                  <el-input v-model="config.user.name" placeholder="Name"></el-input>
                </el-form-item>
                <el-form-item prop="email">
                  <el-input v-model="config.user.email" placeholder="Email"></el-input>
                </el-form-item>
                <el-form-item prop="pass">
                  <el-input v-model="config.user.pass" placeholder="Password" show-password></el-input>
                </el-form-item>
                <el-form-item prop="confirm">
                  <el-input v-model="config.user.confirm" placeholder="Confirm" show-password></el-input>
                </el-form-item>
              </el-form>
              <div class="footer">
                <el-button type="info" plain @click="back">Back</el-button>
                <el-button type="primary" plain @click="confirm('user')">Submit</el-button>
              </div>
            </div>
            <div class="install-form-container" style="text-align: center" v-show="step === 5">
              <el-progress type="circle" :percentage="100" status="text">
                <i class="el-icon-check" style="color: #409EFF; font-size: 48px"></i>
              </el-progress>
              <h2 style="color: #8c939d; margin: 40px 0;">System initialization complete</h2>
              <div class="footer">
                <div class="command">
                  <code v-html="command"></code>
                </div>
                <p>Copy this command to your terminal and execute it</p>
              </div>
            </div>
          </div>
        </el-col>
      </el-row>
      <el-row type="flex" align="bottom">
        <el-col :span="24">
          <footer class="copyright page-copyright-inverse text-center">
            <p class="version">ECTS Version: {{info.version}}</p>
            <p>Copyright © 2018 Betterde.Inc. All Rights Reserved.</p>
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
          callback(new Error('Please enter your password!'));
        } else {
          if (this.config.user.confirm !== '') {
            this.$refs.user.validateField('confirm');
          }
          callback();
        }
      };
      let confirmPass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('Please confirm your password!'));
        } else if (value !== this.config.user.pass) {
          callback(new Error('Password mismatch!'));
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
            service: "/ects/nodes",
            pipeline: "/ects/pipelines",
            config: "/ects/config",
            timeout: 5,
            endpoints: ["localhost:2379"]
          }
        },
        rules: {
          service: {
            host: [
              {type: "string", required: true, message: 'Please enter a host', trigger: 'blur'}
            ],
            port: [
              {type: "integer", required: true, message: 'Please enter a port', trigger: 'blur'}
            ]
          },
          database: {
            host: [
              {type: "string", required: true, message: 'Please enter a host', trigger: 'blur'}
            ],
            port: [
              {type: "integer", required: true, message: 'Please enter a port', trigger: 'blur'}
            ],
            user: [
              {type: "string", required: true, message: 'Please enter a username', trigger: 'blur'}
            ],
            pass: [
              {type: "string", required: true, message: 'Please enter a password', trigger: 'blur'}
            ],
            name: [
              {type: "string", required: true, message: 'Please enter a database name', trigger: 'blur'}
            ],
            char: [
              {type: "string", required: true, message: 'Please enter a charset', trigger: 'blur'}
            ]
          },
          auth: {
            secret: [
              {type: "string", required: true, message: 'Please enter a secret', trigger: 'blur'}
            ],
            ttl: [
              {type: "number", required: true, message: 'Please enter a ttl', trigger: 'blur'}
            ]
          },
          user: {
            name: [
              {type: "string", required: true, message: 'Please enter a user name', trigger: 'blur'}
            ],
            email: [
              {type: 'email', message: 'Please enter a valid email address', trigger: ['blur', 'change']}
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
              {type: "array", required: true, message: 'Please enter a endpoints', trigger: 'blur'}
            ],
            service: [
              {type: "string", required: true, message: 'Please enter a service discover key', trigger: 'blur'}
            ],
            pipeline: [
              {type: "string", required: true, message: 'Please enter a pipeline key', trigger: 'blur'}
            ],
            config: [
              {type: "string", required: true, message: 'Please enter a system config key', trigger: 'blur'}
            ],
            timeout: [
              {type: "number", required: true, message: 'Please enter a timeout', trigger: 'blur'}
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
              this.next();
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
        if (this.step < 4) {
          this.step += 1;
        } else {
          this.submit();
        }
      },
      /**
       * Submit system config
       */
      submit() {
        api.system.initialize(this.config).then(res => {
          this.$message({
            message: res.message,
            type: 'success'
          });
          this.step += 1;
          this.generateCommand()
        }).catch(err => {
          this.$notify.error({
            title: "Error",
            message: err.data.message,
          });
        });
      },
      /**
       * Fetch JSON Web Token secret
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
      fetchDatabaseIsExist() {
        api.system.database(this.config.database).then(res => {
          let exist = res.data;
          if (exist) {
            this.$notify.warning("The database already exists and will overwrite the original data if it continues!")
          }
        })
      },
      generateCommand() {
        let endpoints = this.config.etcd.endpoints.join(',');
        this.command = `ects master --etcd=${endpoints} --name='master' --desc='master node' --config=${this.config.etcd.config}`;
      }
    },
    /**
     * Fetch system version information
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
