import Vue from 'vue'

export default {
  /**
   * Fetch log list
   * @param params
   * @returns {AxiosPromise<any>}
   */
  fetch(params) {
    return Vue.axios.get('/api/log', {
      params: params
    })
  }
}
