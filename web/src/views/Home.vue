<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
  </div>
</template>

<script>
// @ is an alias to /src
import api from '../apis'

export default {
  name: 'home',
  mounted () {
    window.console.log(api);
    api.system.fetch();
  }
}
</script>
