import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  server: {
    // Dev-only: proxy the faucet API to a running instance so the UI can be
    // previewed with real data, e.g. FAUCET_PROXY_TARGET=https://<faucet-host>.
    proxy: process.env.FAUCET_PROXY_TARGET
      ? {
          '/api': {
            changeOrigin: true,
            secure: true,
            target: process.env.FAUCET_PROXY_TARGET,
          },
        }
      : undefined,
  },
})
