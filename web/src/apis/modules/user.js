import Vue from 'vue'

export default {
  fetch(params) {
    return Vue.axios.get('/api/organization/user', {
      params: params
    })
  }
}
