import { createApp } from 'vue'
import './style.css'
import nuxtLabsTheme from 'nuxt-ui-vue/dist/theme/nuxtlabsTheme'
import { createUI } from 'nuxt-ui-vue'
import App from './App.vue'
import { UInput } from 'nuxt-ui-vue';
import router from './routes';

const app = createApp(App)
app.use(UInput)

const UI = createUI({
  registerComponents: false,
})

app.use(router);
app.use(UI, nuxtLabsTheme)

app.mount('#app')