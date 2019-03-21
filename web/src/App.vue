<template>
  <div id="app">
    <component :is="current"></component>
  </div>
</template>

<script>
  import store from './store'
  import {mapState} from 'vuex'
  import Guest from './layouts/Guest'
  import Backend from './layouts/Backend'

  export default {
    name: 'app',
    data() {
      return {
        current: 'guest',
      }
    },
    methods: {},
    components: {
      guest: Guest,
      backend: Backend
    },
    computed: {
      ...mapState({
        profile: state => state.account.profile,
        access_token: state => state.account.access_token
      })
    },
    watch: {
      /**
       * 当access_token的state发生变化时
       * 说明用户登录信息验证成功，并且拿到了access_token
       * 此时调用fetchProfile 这个action 获取用户信息
       */
      access_token() {
      },
    },
    mounted() {
      // 如果用户
      if (this.profile.email) {
        this.current = 'backend'
      }
    }
  }
</script>

<style lang="scss">

</style>
