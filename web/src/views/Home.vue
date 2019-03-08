<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js"/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import api from '../apis'

export default {
  name: 'home',
  components: {
    HelloWorld
  },
  mounted () {
    window.console.log(api);
    api.system.fetch();
  }
}
</script>
