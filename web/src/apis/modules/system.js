import Vue from 'vue'

export default {
  // 获取系统信息
  fetch () {
    return Vue.axios.get("/api/install");
  }
}
