import Vue from 'vue'

export default {
  fetch(params) {
    return Vue.axios.get('/api/log', {
      params: params
    })
  }
}
