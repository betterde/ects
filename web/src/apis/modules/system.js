import Vue from 'vue'

export default {
  /**
   * 获取系统信息
   * @returns {*}
   */
  fetch() {
    return Vue.axios.get('/api/install');
  },
  /**
   * 创建安装配置文件
   * @param params
   * @returns {*}
   */
  install(params) {
    return Vue.axios.post('/api/install', params)
  }
}
