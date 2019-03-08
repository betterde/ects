import Vue from 'vue'
import api from '../../apis'
import * as types from '../types'

export default {
  state: {
    installed: false,
    version: ""
  },
  mutations: {
    SET_SYSTEM: (state, data) => {
      Vue.set(state, 'installed', data.installed);
      Vue.set(state, 'version', data.version);
    }
  },
  actions: {
    fetchSystem({ commit }) {
      api.system.fetch().then(response => {
        commit(types.SET_SYSTEM, response.data)
      });
    }
  }
}
