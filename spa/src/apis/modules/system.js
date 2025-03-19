import Vue from 'vue'

export default {
  /**
   * Fetch system info
   * @returns {*}
   */
  fetch() {
    return Vue.axios.get('/api/system/info');
  },
  /**
   * Get JWT secret
   * @returns {*}
   */
  secret() {
    return Vue.axios.post('/api/system/secret')
  },
  /**
   * Validate database name is exist
   * @returns {*}
   */
  database(params) {
    return Vue.axios.post('/api/system/database', params)
  },
  /**
   * Initialization system config
   * @param params
   * @returns {*}
   */
  initialize(params) {
    return Vue.axios.post('/api/system/initialization', params)
  }
}
