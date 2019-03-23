import Vue from 'vue'
import api from '../../apis'
import * as types from '../types'

export default {
  state: {
    layout: {
      current: "guest"
    },
    menu: {
      active: "/"
    },
    installed: false,
    version: ""
  },
  mutations: {
    SET_SYSTEM: (state, data) => {
      Vue.set(state, 'installed', data.installed);
      Vue.set(state, 'version', data.version);
    },
    SET_MENU_ACTIVE: (state, data) => {
      Vue.set(state.menu, 'active', data)
    },
    SET_LAYOUT_CURRENT: (state, data) => {
      Vue.set(state.layout, 'current', data)
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
