import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import axios from 'axios'

import './styles/index.scss'
import 'bootstrap/dist/css/bootstrap.min.css'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faTrash } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

const baseURL = window.location.origin + '/v1'

// const baseURL = 'http://localhost:3000/v1'

const axiosInstance = axios.create({
  baseURL
})

library.add(faTrash)

const app = createApp(App)
app.component('font-awesome-icon', FontAwesomeIcon)
app.provide('$axios', axiosInstance)

app.use(createPinia())
app.use(router)

app.mount('#app')
