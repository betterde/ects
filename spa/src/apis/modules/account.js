import Vue from 'vue'

export default {
  /**
   * User sign in
   * @param params
   * @returns {AxiosPromise<any>}
   */
  signin(params) {
    return Vue.axios.post('/api/auth/signin', params)
  },
  /**
   * Fetch user profile
   * @returns {AxiosPromise<any>}
   */
  profile() {
    return Vue.axios.get('/api/account/profile');
  }
}
