import Vue from 'vue'

export default {
  /**
   * 获取节点概览
   * @returns {AxiosPromise<any>}
   */
  fetchNodes() {
    return Vue.axios.get('/api/dashboard/nodes');
  },
  /**
   * 获取正在调度的流水线数量
   * @returns {AxiosPromise<any>}
   */
  fetchPipelines() {
    return Vue.axios.get('/api/dashboard/pipelines');
  },
  fetchFailtures() {
    return Vue.axios.get('/api/dashboard/failtures');
  }
}
