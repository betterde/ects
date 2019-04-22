import Vue from 'vue'

export default {
  /**
   * Fetch pipeline list
   * @param params
   * @returns {*}
   */
  fetch(params) {
    return Vue.axios.get('/api/pipeline', {
      params: params
    });
  }
}
