import Vue from "vue"
import jquery from "jquery"
import { BootstrapVue, IconsPlugin } from "bootstrap-vue"
import axios from "axios"
import VueAxios from "vue-axios"
import App from "./App.vue"
import router from "./router"
import store from "./store"

// scss style
import "./assets/scss/index.scss"

// Install BootstrapVue
Vue.use(BootstrapVue, IconsPlugin)

// axios
Vue.use(axios, VueAxios)
Vue.prototype.$ = jquery

Vue.config.productionTip = false

new Vue({
    router,
    store,
    render: (h) => h(App),
}).$mount("#app")
