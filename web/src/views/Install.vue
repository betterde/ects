<template>
  <div id="install">
    <div class="panel">
      <el-row>
        <el-col :span="12" :offset="6">
          <div class="panel-heading">
            <h1 class="title">Install Elastic Crontab System</h1>
          </div>
          <el-steps :active="1" finish-status="success" simple style="margin-top: 20px">
            <el-step title="数据库" ></el-step>
            <el-step title="系统" ></el-step>
            <el-step title="完成" ></el-step>
          </el-steps>
          <div class="panel-body">
            <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" label-position="top" class="demo-ruleForm">
              <el-form-item label="活动名称" prop="name">
                <el-input v-model="ruleForm.name"></el-input>
              </el-form-item>
              <el-form-item label="活动区域" prop="region">
                <el-select v-model="ruleForm.region" placeholder="请选择活动区域">
                  <el-option label="区域一" value="shanghai"></el-option>
                  <el-option label="区域二" value="beijing"></el-option>
                </el-select>
              </el-form-item>
              <el-form-item label="即时配送" prop="delivery">
                <el-switch v-model="ruleForm.delivery"></el-switch>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
                <el-button @click="resetForm('ruleForm')">重置</el-button>
              </el-form-item>
            </el-form>
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
        ruleForm: {
          name: '',
          region: '',
          date1: '',
          date2: '',
          delivery: false,
          type: [],
          resource: '',
          desc: ''
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
            console.log('error submit!!');
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
  #install {
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
