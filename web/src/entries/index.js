import Vue from 'vue'
import App from '../App.vue'
import store from '../store'
import router from '../router'
import '../plugins/axios.js'
import '../plugins/element.js'

import Collaspe from '../components/Collapse'

Vue.config.productionTip = false;

Vue.component('collapse-view', Collaspe);

new Vue({
  store,
  router,
  render: h => h(App)
}).$mount('#app');
