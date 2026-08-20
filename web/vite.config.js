import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  define: {
    // Brand is selected at build time. Passed via `VITE_BRAND=<id> yarn build`
    // (see Dockerfile). Defaults to the Gateway "base" brand.
    'import.meta.env.VITE_BRAND': JSON.stringify(process.env.VITE_BRAND || 'base'),
  },
})
