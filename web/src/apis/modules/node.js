import Vue from 'vue'

export default {
  /**
   * 获取节点列表
   * @param params
   * @returns {*}
   */
  fetch(params) {
    return Vue.axios.get('/api/worker', {
      params: params
    });
  },
  /**
   * 创建节点
   * @param params
   * @returns {*}
   */
  create(params) {
    return Vue.axios.post('/api/worker', params);
  },
  /**
   * 删除节点
   * @param params
   * @returns {*}
   */
  delete(params) {
    return Vue.axios.delete(`/api/worker/${params.id}`);
  }
}
