import Vue from 'vue'

export default {
  /**
   * 获取任务列表
   * @param parmas
   * @returns {*|AxiosInstance}
   */
  fetch(parmas) {
    return Vue.axios.get('/api/task', {
      params: parmas
    });
  },
  /**
   * 创建任务
   * @param params
   * @returns {*}
   */
  create(params) {
    return Vue.axios.post('/api/task', params);
  },
  /**
   * 删除任务
   * @param params
   * @returns {*}
   */
  delete(params) {
    return Vue.axios.delete(`/api/task/${params.id}`);
  }
}
