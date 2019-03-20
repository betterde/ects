<template>
  <div class="loginview">
    <div class="panel login-panel">
      <div class="panel-header">
        <h1 class="panel-title">Elastic Crontab System</h1>
      </div>
      <div class="panel-body">
        <el-form :model="credentials" :rules="rules" ref="signin">
          <el-form-item prop="username">
            <el-input v-model="credentials.username" placeholder="账户"></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input type="password" v-model="credentials.password" @keyup.enter.native="submit('signin')" placeholder="密码" show-password></el-input>
          </el-form-item>
          <el-form-item class="login-button">
            <el-button type="primary" plain class="pull-right" style="width: 100%" @click="submit('signin')" :loading="loading">登录</el-button>
          </el-form-item>
          <div class="tips">
            <p>如果您忘记了密码，请点击这里</p>
          </div>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
  import store from '../store'
  import { mapActions, mapState } from 'vuex'

  export default {
    name: "SignIn",
    data () {
      return {
        loading: false,
        credentials: {
          username: '',
          password: ''
        },
        rules: {
          username: [
            { required: true, message: '请输入您的账户', trigger: 'blur' },
            { min: 6, max: 25, message: '长度在 6 到 25 个字符之间', trigger: 'blur' }
          ],
          password: [
            { required: true, message: '请输入您的密码', trigger: 'blur' },
            { min: 6, max: 25, message: '长度在 6 到 25 个字符之间', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      ...mapActions([
        'signIn',
        'fetchAccount'
      ]),
      submit(name) {
        this.$refs[name].validate((valid) => {
          if (valid) {
            this.loading = true;
            store.dispatch('signIn', this.credentials).then(res => {
              if (res.code === 200) {
                this.$message.success(res.message);
                setTimeout(() => {
                  this.$router.replace("/")
                }, 800);
              } else {
                this.$message.warning(res.message);
              }
              this.loading = false;
            });
          } else {
            return false;
          }
        });
      }
    },
    computed: {
      ...mapState ({
        access_token: state => state.account.access_token,
        user: state => state.account.user,
      }),
    },
    mounted() {

    }
  }
</script>

<style lang="scss">
  html, body {
    color: #76838f;
    margin: 0;
    #app {
      height: 100%;
    }
  }
  .loginview {
    background: #62a8ea;
    background-image: url(data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiA/Pgo8c3ZnIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgdmlld0JveD0iMCAwIDEgMSIgcHJlc2VydmVBc3BlY3RSYXRpbz0ibm9uZSI+CiAgPGxpbmVhckdyYWRpZW50IGlkPSJncmFkLXVjZ2ctZ2VuZXJhdGVkIiBncmFkaWVudFVuaXRzPSJ1c2VyU3BhY2VPblVzZSIgeDE9IjAlIiB5MT0iMCUiIHgyPSIwJSIgeTI9IjEwMCUiPgogICAgPHN0b3Agb2Zmc2V0PSIwJSIgc3RvcC1jb2xvcj0iIzYyYThlYSIgc3RvcC1vcGFjaXR5PSIxIi8+CiAgICA8c3RvcCBvZmZzZXQ9IjEwMCUiIHN0b3AtY29sb3I9IiMxNTcxYjEiIHN0b3Atb3BhY2l0eT0iMSIvPgogIDwvbGluZWFyR3JhZGllbnQ+CiAgPHJlY3QgeD0iMCIgeT0iMCIgd2lkdGg9IjEiIGhlaWdodD0iMSIgZmlsbD0idXJsKCNncmFkLXVjZ2ctZ2VuZXJhdGVkKSIgLz4KPC9zdmc+);
    background-image: -webkit-linear-gradient(top, #62a8ea 0, #3583ca 100%);
    background-image: -o-linear-gradient(top, #62a8ea 0, #3583ca 100%);
    background-image: -webkit-gradient(linear, left top, left bottom, from(#62a8ea), to(#3583ca));
    background-image: linear-gradient(to bottom, #62a8ea 0, #3583ca 100%);
    background-repeat: repeat-x;
    height: 100%;
    .panel-title {
      font-size: 28px;
      margin: 0 0 40px 0;
    }
    .footer-row {
      height: 100%;
      z-index: 0;
    }
  }
  .login-panel {
    background-color: #FFF;
    padding: 40px;
    border-radius: 4px;
    box-shadow: 0 2px 2px rgba(0, 0, 0, .05);
    width: 400px;
    display: inline-block;
    max-width: 100%;
    position: absolute;
    top: 50%;
    left: 50%;
    margin-left: -200px;
    transform: translateY(-50%);
    z-index: 10;
    .panel-header {
      text-align: center;
      .panel-title {
        color: #76838f;
      }
    }
    .panel-body {
      text-align: center;
      padding: 10px;
      .tips {
        font-size: 12px;
      }
    }
    .login-button {
      padding-top: 10px;
    }
  }
</style>
