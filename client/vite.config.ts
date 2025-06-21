import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

const srcPath = path.resolve(__dirname, 'src').replace(/\\/g, '/')

export default defineConfig({
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: `@use "${srcPath}/styles/common" as *;`,
      },
    },
  },
  plugins: [vue()],
})
