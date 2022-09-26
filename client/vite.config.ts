import vue from '@vitejs/plugin-vue'
import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/api': {
        target: "",
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '/'),
      }
    }
  }
})
