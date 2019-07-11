import Vue from 'vue'

export default {
  /**
   * 获取节点列表
   * @param params
   * @returns {*}
   */
  fetch(params) {
    return Vue.axios.get('/api/node', {
      params: params
    });
  },
  /**
   * 创建节点
   * @param params
   * @returns {*}
   */
  create(params) {
    return Vue.axios.post('/api/node', params);
  },
  /**
   * 删除节点
   * @param params
   * @returns {*}
   */
  delete(params) {
    return Vue.axios.delete(`/api/node/${params.id}`);
  }
}
