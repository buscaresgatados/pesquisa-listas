import { createApp } from 'vue'
import './style.css'
import nuxtLabsTheme from 'nuxt-ui-vue/dist/theme/nuxtlabsTheme'

import { createUI } from 'nuxt-ui-vue'

import App from './App.vue'

const app = createApp(App)

const UI = createUI({
  registerComponents: false,
})

app.use(UI, nuxtLabsTheme)

app.mount('#app')