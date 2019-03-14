<template>
  <div id="install">
    <div class="panel">
      <el-row>
        <el-col :span="12" :offset="6">
          <div class="panel-heading">
            <h1 class="title">Install Elastic Crontab System</h1>
          </div>
          <el-steps :active="step" finish-status="success" simple style="margin-top: 20px">
            <el-step title="环境检查" ></el-step>
            <el-step title="服务端口" ></el-step>
            <el-step title="数据库" ></el-step>
            <el-step title="用户信息" ></el-step>
            <el-step title="安装完成" ></el-step>
          </el-steps>
          <div class="panel-body">
            <div v-show="step === 0">
              <div class="system-info">
                <el-alert title="配置文件写入权限" type="success" :closable="false" description="如您运行时未指定配置文件路径，系统将使用 /etc/ects/ects.yaml 作为配置文件的默认路径" show-icon></el-alert>
              </div>
              <div class="footer">
                <el-button type="primary" plain @click="confirm">下一步</el-button>
              </div>
            </div>
            <div class="install-form-container" v-show="step === 1">
              <el-form :model="config.service" :rules="serviceRules" ref="ruleForm" label-width="100px" label-position="top">
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
                <el-button type="primary" plain @click="confirm">下一步</el-button>
              </div>
            </div>
            <div class="install-form-container" v-show="step === 2">
              <el-form :model="config.database" :rules="rules" ref="ruleForm" class="demo-ruleForm">
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
                  <el-form-item prop="user">
                    <el-input v-model="config.database.name" placeholder="数据库名称"></el-input>
                  </el-form-item>
                </el-col>
                <el-col class="line" :span="2"></el-col>
                <el-col :span="8">
                  <el-form-item prop="name">
                    <el-select v-model="config.database.char" placeholder="请选择字符类型">
                      <el-option v-for="char in chars" :key="char.value" :label="char.label" :value="char.value"></el-option>
                    </el-select>
                  </el-form-item>
                </el-col>
                <el-form-item prop="pass">
                  <el-input v-model="config.database.user" placeholder="数据库用户名"></el-input>
                </el-form-item>
                <el-form-item prop="name">
                  <el-input v-model="config.database.pass" placeholder="数据库密码"></el-input>
                </el-form-item>
              </el-form>
              <div class="footer">
                <el-button type="info" plain @click="back">上一步</el-button>
                <el-button type="primary" plain @click="confirm">下一步</el-button>
              </div>
            </div>
            <div class="install-form-container" v-show="step === 3">
              <el-form :model="config.user" :rules="rules" ref="ruleForm">
                <el-form-item prop="name">
                  <el-input v-model="config.user.name" placeholder="昵称"></el-input>
                </el-form-item>
                <el-form-item prop="region">
                  <el-input v-model="config.user.email" placeholder="邮箱"></el-input>
                </el-form-item>
                <el-form-item prop="region">
                  <el-input v-model="config.user.pass" placeholder="密码"></el-input>
                </el-form-item>
                <el-form-item prop="region">
                  <el-input v-model="config.user.pass" placeholder="确认密码"></el-input>
                </el-form-item>
              </el-form>
              <div class="footer">
                <el-button type="info" plain @click="back">上一步</el-button>
                <el-button type="primary" plain @click="confirm">确认安装</el-button>
              </div>
            </div>
            <div class="install-form-container" style="text-align: center" v-show="step === 4">
              <el-progress type="circle" :percentage="100" status="text">
                <i class="el-icon-check" style="color: #409EFF; font-size: 48px"></i>
              </el-progress>
              <h2 style="color: #8c939d; margin: 40px 0;">系统初始化完成</h2>
              <div class="footer">
                <el-button type="primary" @click="confirm">进入后台</el-button>
              </div>
            </div>
          </div>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Install",
    data() {
      return {
        step: 0,
        config: {
          service: {
            host: "0.0.0.0",
            port: 9701
          },
          database: {
            host: "localhost",
            port: 3306,
            user: "root",
            pass: "",
            name: "ects",
            char: ""
          },
          auth: {
            secret: "",
            ttl: 86400
          },
          user: {
            name: "",
            email: "",
            pass: ""
          }
        },
        serviceRules: {
          host: [
            { required: true, message: '请输入服务运行时绑定的IP', trigger: 'blur' }
          ],
          port: [
            { required: true, message: '请输入服务运行时监听的端口', trigger: 'blur' }
          ]
        },
        databaseRules: {
          host: [
            { required: true, message: '请输入服务MySQL主机地址', trigger: 'blur' }
          ],
          port: [
            { required: true, message: '请输入服务MySQL主机端口', trigger: 'blur' }
          ],
          user: [
            { required: true, message: '请输入服务MySQL用户名', trigger: 'blur' }
          ],
          pass: [
            { required: true, message: '请输入服务MySQL密码', trigger: 'blur' }
          ],
          name: [
            { required: true, message: '请输入服务MySQL数据库名称', trigger: 'blur' }
          ],
          char: [
            { required: true, message: '请选择数据库字符类型', trigger: 'blur' }
          ]
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
      back() {
        if (this.step >= 1) {
          this.step -= 1
        }
      },
      confirm() {
        if (this.step < 4) {
          this.step += 1
        } else {
          // 当步进数大于4时，表明已经安装成功，则跳转到后台页面
          window.location.href = "/"
        }
      },
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            alert('submit!');
          } else {
            window.console.log('error submit!!');
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
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
    height: 100%;
    text-align: center;
    background-color: #f0f2f5;
    display: flex;
    justify-content:center;
    align-items:Center;
    .panel {
      min-height: 600px;
      width: 100%;
      .panel-heading {
        padding: 40px;
        .title {
          font-size: 48px;
          color: #8c939d;
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
