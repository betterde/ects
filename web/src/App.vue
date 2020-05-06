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
      })
    },
    watch: {
      profile(value) {
        if (value.id !== '') {
          this.$store.commit('SET_LAYOUT_CURRENT', 'backend');
        }
      },
      '$route' (to) {
        if (to.meta.hasOwnProperty('requiresAuth') === false || to.meta.requiresAuth === false) {
          this.$store.commit('SET_LAYOUT_CURRENT', 'guest');
        } else {
          this.$store.commit('SET_LAYOUT_CURRENT', 'backend');
        }
      }
    }
  }
</script>

<style lang="scss">

</style>
