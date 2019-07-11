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
  },
  /**
   * Create pipeline
   * @param params
   * @returns {*}
   */
  create(params) {
    return Vue.axios.post('/api/pipeline', params);
  },
  /**
   * Update pipeline attribute
   * @param id
   * @param params
   * @returns {AxiosPromise<any> | IDBRequest<IDBValidKey> | Promise<void>}
   */
  update(id, params) {
    return Vue.axios.put(`/api/pipeline/${id}`, params);
  },
  /**
   * Delete pipeline
   * @param id
   * @returns {*}
   */
  delete(id) {
    return Vue.axios.delete(`/api/pipeline/${id}`);
  }
}
