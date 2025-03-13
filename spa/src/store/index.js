import Vue from 'vue'
import Vuex from 'vuex'
import system from './modules/system'
import account from './modules/account'

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    system: system,
    account: account
  },
  strict: process.env.MIX_NODE_ENV !== 'production'
})
