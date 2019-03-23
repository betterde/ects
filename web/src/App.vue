<template>
  <div id="app">
    <component :is="layout.current"></component>
  </div>
</template>

<script>
  import {mapState} from 'vuex'
  import Guest from './layouts/Guest'
  import Backend from './layouts/Backend'

  export default {
    name: 'app',
    data() {
      return {
      }
    },
    methods: {},
    components: {
      guest: Guest,
      backend: Backend
    },
    computed: {
      ...mapState({
        layout: state => state.system.layout,
        profile: state => state.account.profile,
        access_token: state => state.account.access_token
      })
    },
    watch: {
      profile(value) {
        if (value.id === "") {
          this.$store.commit("SET_LAYOUT_CURRENT", 'guest');
        } else {
          this.$store.commit("SET_LAYOUT_CURRENT", 'backend');
        }
      },
      /**
       * 当access_token的state发生变化时
       * 说明用户登录信息验证成功，并且拿到了access_token
       * 此时调用fetchProfile 这个action 获取用户信息
       */
      access_token(value) {
        if (value == null) {
          window.console.log(value);
        }
      },
      '$route' (to, from) {
        if (from.path === '/' && to.name === 'notfound') {
          this.$store.commit("SET_LAYOUT_CURRENT", 'guest');
        }

        if (from.path === '/' && to.name !== 'notfound') {
          this.$store.commit("SET_LAYOUT_CURRENT", 'backend');
        }

        if (from.name === 'notfound') {
          this.$store.commit('SET_MENU_ACTIVE', to.path);
        }

        this.$store.commit('SET_MENU_ACTIVE', to.path);
      }
    }
  }
</script>

<style lang="scss">

</style>
