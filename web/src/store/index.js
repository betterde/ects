import Vue from 'vue'
import Vuex from 'vuex'
import system from './modules/system'

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    system: system
  },
  strict: process.env.MIX_NODE_ENV !== 'production'
})
