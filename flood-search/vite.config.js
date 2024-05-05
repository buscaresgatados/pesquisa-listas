import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { NuxtUIVueComponentResolver } from 'nuxt-ui-vue'
import Components from 'unplugin-vue-components/vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [
        NuxtUIVueComponentResolver()
      ]
    })
  ],
})
