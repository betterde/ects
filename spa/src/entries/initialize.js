import Vue from 'vue'
import Initialize from '../views/Initialize.vue'
import '../plugins/axios.js'
import '../plugins/element.js'

Vue.config.productionTip = false;

new Vue({
  render: h => h(Initialize)
}).$mount('#app');
