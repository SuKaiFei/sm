import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import App from './App.vue'
import store from './store'
import router from './router'
import particles from 'particles.js/particles';
import vuelidate from 'vuelidate'

import '@/icons' // icon
import '@/permission' // permission control
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'

Vue.config.productionTip = false
Vue.use(mavonEditor)
Vue.use(BootstrapVue)
Vue.use(particles)

new Vue({
	router,
	store,
	render: h => h(App),
	validations: vuelidate
}).$mount('#app')
