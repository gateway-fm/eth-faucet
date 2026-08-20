<script>
  import 'animate.css';
  import 'bulma/css/bulma.css';
  import Faucet from './Faucet.svelte';
  import { brand, brandCssVars } from './lib/brand.js';

  // Brand favicon wins; fall back to the VITE_FAVICON_PATH env for older deployments.
  const faviconPath =
    brand.assets.favicon || import.meta.env.VITE_FAVICON_PATH || '/favicon.png';

  // Active brand's design tokens, injected as CSS custom properties on :root
  // during <head> parse so there is no theme flash before the app mounts.
  const rootVars = `:root {\n  ${brandCssVars()}\n}`;
</script>

<svelte:head>
  <link
    href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"
    rel="stylesheet"
  />
  <link rel="icon" href={faviconPath} />
  <!-- eslint-disable-next-line svelte/no-at-html-tags -->
  {@html `<style>${rootVars}</style>`}
</svelte:head>

<Faucet />

<style>
  /* DM Sans (SIL Open Font License) — vendored in /public/fonts, self-hosted for offline/embedded use */
  @font-face {
    font-family: 'DM Sans';
    font-style: normal;
    font-weight: 100 1000;
    font-display: swap;
    src: url('/fonts/dmsans-latin-ext.woff2') format('woff2');
    unicode-range: U+0100-02BA, U+02BD-02C5, U+02C7-02CC, U+02CE-02D7,
      U+02DD-02FF, U+0304, U+0308, U+0329, U+1D00-1DBF, U+1E00-1E9F,
      U+1EF2-1EFF, U+2020, U+20A0-20AB, U+20AD-20C0, U+2113, U+2C60-2C7F,
      U+A720-A7FF;
  }
  @font-face {
    font-family: 'DM Sans';
    font-style: normal;
    font-weight: 100 1000;
    font-display: swap;
    src: url('/fonts/dmsans-latin.woff2') format('woff2');
    unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6,
      U+02DA, U+02DC, U+0304, U+0308, U+0329, U+2000-206F, U+20AC, U+2122,
      U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
  }

  :global(html),
  :global(body),
  :global(button),
  :global(input),
  :global(textarea),
  :global(.title),
  :global(.subtitle) {
    font-family: var(
      --fc-font-family,
      'DM Sans',
      system-ui,
      -apple-system,
      'Segoe UI',
      Roboto,
      Helvetica,
      Arial,
      sans-serif
    );
  }
</style>
