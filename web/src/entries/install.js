import Vue from 'vue'
import Install from '../views/Install.vue'
import '../plugins/axios.js'
import '../plugins/element.js'

Vue.config.productionTip = false;

new Vue({
  render: h => h(Install)
}).$mount('#app');
