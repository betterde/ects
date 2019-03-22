import Vue from 'vue'

export default {
  signin(params) {
    return Vue.axios.post('/api/auth/signin', params)
  },
  profile() {
    return Vue.axios.get('/api/account/profile');
  }
}
