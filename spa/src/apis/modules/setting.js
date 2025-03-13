import Vue from 'vue'

export default {
  /**
   * 获取通知
   * @returns {AxiosPromise<any>}
   */
  fetchNotification() {
    return Vue.axios.get('/api/setting/notification');
  },
  /**
   * 更新通知设置
   * @param params
   * @returns {AxiosPromise<any>}
   */
  updateNotification(params) {
    return Vue.axios.put('/api/setting/notification', params)
  },
  /**
   * 发送测试邮件
   * @param params
   * @returns {AxiosPromise<any>}
   */
  sendMail(params) {
    return Vue.axios.post('/api/setting/mail', params)
  }
}
