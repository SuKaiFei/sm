import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import App from './App.vue'
import store from './store'
import router from './router'
import particles from 'particles.js/particles';
import '@/icons' // icon

Vue.config.productionTip = false
Vue.use(BootstrapVue)
Vue.use(particles)

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
