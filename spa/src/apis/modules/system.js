import Vue from 'vue'

export default {
  /**
   * Fetch system info
   * @returns {*}
   */
  fetch() {
    return Vue.axios.get('/api/initialize');
  },
  /**
   * Initialization system config
   * @param params
   * @returns {*}
   */
  initialize(params) {
    return Vue.axios.post('/api/initialize', params)
  },
  /**
   * Get JWT secret
   * @returns {*}
   */
  secret() {
    return Vue.axios.get('/api/initialize/secret')
  },
  /**
   * Validate database name is exist
   * @returns {*}
   */
  database(params) {
    return Vue.axios.get('/api/initialize/database', {
      params: params
    })
  }
}
