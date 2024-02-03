import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    // 别名路径配置项 实际的路径转换 @ -> src
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
