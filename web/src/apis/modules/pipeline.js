import Vue from 'vue'

export default {
  /**
   * 获取流水线列表
   * @param params
   * @returns {*}
   */
  fetch(params) {
    return Vue.axios.get('/api/pipeline', {
      params: params
    });
  },
  /**
   * 创建流水线
   * @param params
   * @returns {*}
   */
  create(params) {
    return Vue.axios.post('/api/pipeline', params);
  },
  /**
   * 更新流水线
   * @param id
   * @param params
   * @returns {AxiosPromise<any> | IDBRequest<IDBValidKey> | Promise<void>}
   */
  update(id, params) {
    return Vue.axios.put(`/api/pipeline/${id}`, params);
  },
  /**
   * 删除流水线
   * @param id
   * @returns {*}
   */
  delete(id) {
    return Vue.axios.delete(`/api/pipeline/${id}`);
  },
  /**
   * 获取流水线关联的任务
   * @param id
   * @returns {AxiosPromise<any>}
   */
  fetchTasks(id) {
    return Vue.axios.get(`/api/pipeline/tasks`, {
      params: {
        pipeline_id: id
      }
    });
  }
}
