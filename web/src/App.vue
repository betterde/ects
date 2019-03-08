<template>
  <div id="app">
    <img src="./assets/logo.png">
    <div>
      <p>
        If Element is successfully added to this project, you'll see an
        <code v-text="'<el-button>'"></code>
        below
      </p>
      <el-button>el-button</el-button>
    </div>
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import api from './apis'

export default {
  name: 'app',
  components: {
    HelloWorld
  },
  data() {
    return {
      system: {
        installed: false,
        version: ""
      }
    }
  },
  mounted() {
    api.system.fetch().then(res => {
      this.system = res.data
    });
    if (this.system.installed) {

    } else {
      this.$router.push('/install')
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
