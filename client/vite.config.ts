import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

const srcPath = path.resolve(__dirname, 'src').replace(/\\/g, '/')

export default defineConfig({
  server: {
    proxy: {
      '/api': {
        target: 'http://h25s-04.trap.show',
        changeOrigin: true,
      },
    },
  },
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: `@use "${srcPath}/styles/common" as *;`,
      },
    },
  },
  plugins: [vue()],
})
