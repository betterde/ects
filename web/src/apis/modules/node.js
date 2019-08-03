import Vue from 'vue'

export default {
  /**
   * Fetch node list
   * @param params
   * @returns {*}
   */
  fetch(params) {
    return Vue.axios.get('/api/node', {
      params: params
    });
  },
  /**
   * Create node
   * @param params
   * @returns {*}
   */
  create(params) {
    return Vue.axios.post('/api/node', params);
  },
  /**
   * Update node attributes
   * @param id
   * @param params
   * @returns {*}
   */
  update(id, params) {
    return Vue.axios.put(`/api/node/${id}`, params)
  },
  /**
   * Delete node
   * @param id
   * @returns {*}
   */
  delete(id) {
    return Vue.axios.delete(`/api/node/${id}`);
  }
}
