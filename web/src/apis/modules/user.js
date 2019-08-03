import Vue from 'vue'

export default {
  /**
   * Fetch user list
   * @param params
   * @returns {AxiosPromise<any>}
   */
  fetch(params) {
    return Vue.axios.get('/api/user', {
      params: params
    })
  },
  /**
   * Create user
   * @param params
   * @returns {AxiosPromise<any>}
   */
  create(params) {
    return Vue.axios.post('/api/user', params)
  },
  /**
   * Update user attributes
   * @param id
   * @param params
   * @returns {AxiosPromise<any>}
   */
  update(id, params) {
    return Vue.axios.put(`/api/user/${id}`, params)
  },
  /**
   * Delete user
   * @param id
   * @returns {AxiosPromise}
   */
  delete(id) {
    return Vue.axios.delete(`/api/user/${id}`)
  }
}
