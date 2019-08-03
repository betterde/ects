import Vue from 'vue'

export default {
  /**
   * Fetch task list
   * @param parmas
   * @returns {*|AxiosInstance}
   */
  fetch(parmas) {
    return Vue.axios.get('/api/task', {
      params: parmas
    });
  },
  /**
   * Create task
   * @param params
   * @returns {*}
   */
  create(params) {
    return Vue.axios.post('/api/task', params);
  },
  /**
   * Update task attribute
   * @param id
   * @param params
   * @returns {AxiosPromise<any> | IDBRequest<IDBValidKey> | Promise<void>}
   */
  update(id, params) {
    return Vue.axios.put(`/api/task/${id}`, params);
  },
  /**
   * Delete task
   * @param id
   * @returns {*}
   */
  delete(id) {
    return Vue.axios.delete(`/api/task/${id}`);
  }
}
