import Vue from 'vue'
import store from '../../store'

export default {
  signin(params) {
    return Vue.axios.post('/api/auth/signin', params)
  },
  profile() {
    window.console.log(store.state.account.access_token);
    return Vue.axios.get('/api/account/profile');
  }
}
