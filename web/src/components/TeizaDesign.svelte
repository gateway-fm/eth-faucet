<script>
  // TEIZA co-branded faucet, specced from the "Gateway Makeover x TEIZA"
  // design: DM Sans, ink #1C1C1C, text accent #1D4AE1, CTA fill #3F58FF,
  // white-to-blue page gradient, "Powered by Gateway.fm" footer.
  import gatewayLogo from './../assets/gatewayfm-ink-logo.svg';
  import teizaLogo from './../assets/teiza-logo.svg';
  import teizaMark from './../assets/teiza-mark.svg';

  export let faucetInfo;
  export let handleRequest;
  export let gweiToEth;

  let address = '';
  let requesting = false;
  let copied = false;

  $: payout = gweiToEth(faucetInfo.payout);

  const navLinks = [
    { title: 'Rollup', url: 'https://gateway.fm/presto' },
    { title: 'RPC', url: 'https://gateway.fm/rpc' },
    { title: 'Blog', url: 'https://gateway.fm/blog' },
    { title: 'About', url: 'https://gateway.fm/about' },
  ];

  async function onRequest() {
    if (!address.trim() || requesting) return;
    requesting = true;
    try {
      await handleRequest(address.trim());
    } finally {
      requesting = false;
    }
  }

  let copiedTimer;

  async function onCopy() {
    try {
      await navigator.clipboard.writeText(faucetInfo.account);
    } catch {
      // Clipboard API needs a secure context; fall back to a hidden textarea.
      const textarea = document.createElement('textarea');
      textarea.value = faucetInfo.account;
      textarea.setAttribute('readonly', '');
      textarea.style.position = 'absolute';
      textarea.style.left = '-9999px';
      document.body.appendChild(textarea);
      textarea.select();
      document.execCommand('copy');
      document.body.removeChild(textarea);
    }
    copied = true;
    clearTimeout(copiedTimer);
    copiedTimer = setTimeout(() => (copied = false), 1600);
  }
</script>

<main class="tz-root">
  <header class="tz-header">
    <a class="tz-logo" href="https://teiza.com/" target="_blank" rel="noopener noreferrer">
      <img src={teizaLogo} alt="TEIZA" />
    </a>
    <nav class="tz-nav">
      {#each navLinks as link}
        <a href={link.url} target="_blank" rel="noopener noreferrer">{link.title}</a>
      {/each}
    </nav>
    <a
      class="tz-deploy"
      href="https://gateway.fm/presto/"
      target="_blank"
      rel="noopener noreferrer">Deploy rollup</a
    >
  </header>

  <section class="tz-content">
    <span class="tz-badge"><span class="tz-badge-dot"></span>{faucetInfo.network}</span>
    <img class="tz-mark" src={teizaMark} alt="" />
    <h1 class="tz-heading">
      <span class="tz-heading-accent">Receive</span>
      {payout}
      {faucetInfo.symbol}
    </h1>

    <div class="tz-card">
      <div class="tz-serving-row">
        <span class="tz-serving-label">Serving from</span>
        <div class="tz-serving-chip">
          <span class="tz-serving-address">{faucetInfo.account}</span>
          {#if copied}
            <span class="tz-copied-label">Copied!</span>
          {/if}
          <button
            aria-label="Copy serving address"
            class="tz-copy-button"
            title="Copy address"
            on:click={onCopy}
          >
            {#if copied}
              <svg class="tz-copy-icon" width="18" height="18" viewBox="0 0 24 24" fill="none">
                <path
                  d="M5 12.5l4.5 4.5L19 7.5"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
              </svg>
            {:else}
              <svg class="tz-copy-icon" width="18" height="18" viewBox="0 0 24 24" fill="none">
                <rect
                  x="4"
                  y="4"
                  width="12.5"
                  height="12.5"
                  rx="1.5"
                  stroke="currentColor"
                  stroke-width="1.5"
                />
                <path
                  d="M8.5 20h10A1.5 1.5 0 0 0 20 18.5v-10"
                  stroke="currentColor"
                  stroke-width="1.5"
                  stroke-linecap="round"
                />
              </svg>
            {/if}
          </button>
        </div>
      </div>

      <input
        class="tz-address-input"
        placeholder="Enter your address"
        spellcheck="false"
        bind:value={address}
      />

      <div class="tz-button-row">
        <button class="tz-request" disabled={!address.trim() || requesting} on:click={onRequest}>
          {requesting ? 'Requesting…' : 'Request'}
        </button>
      </div>

      <p class="tz-helper">
        Claim {payout}
        {faucetInfo.symbol} test token for development.
        <br />
        If you need additional tokens for extensive testing, please
        <a
          class="tz-helper-link"
          href="https://gateway.fm/"
          target="_blank"
          rel="noopener noreferrer">contact support</a
        >.
      </p>
    </div>

    <div class="tz-powered">
      Powered by
      <a href="https://gateway.fm/" target="_blank" rel="noopener noreferrer">
        <img src={gatewayLogo} alt="gateway.fm" />
      </a>
    </div>
  </section>

  <div id="hcaptcha" data-size="invisible"></div>
</main>

<style>
  @font-face {
    font-family: 'DM Sans';
    font-style: normal;
    font-weight: 400 700;
    src: url('/fonts/dm-sans/DMSans-latin.woff2') format('woff2');
    unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+0304,
      U+0308, U+0329, U+2000-206F, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
  }
  @font-face {
    font-family: 'DM Sans';
    font-style: normal;
    font-weight: 400 700;
    src: url('/fonts/dm-sans/DMSans-latin-ext.woff2') format('woff2');
    unicode-range: U+0100-02BA, U+02BD-02C5, U+02C7-02CC, U+02CE-02D7, U+02DD-02FF, U+0304, U+0308,
      U+0329, U+1D00-1DBF, U+1E00-1E9F, U+1EF2-1EFF, U+2020, U+20A0-20AB, U+20AD-20C0, U+2113,
      U+2C60-2C7F, U+A720-A7FF;
  }

  .tz-root {
    background: linear-gradient(180deg, #ffffff 0%, #fdfcff 50%, #e2ebff 100%);
    color: #1c1c1c;
    display: flex;
    flex-direction: column;
    font-family: 'DM Sans', system-ui, sans-serif;
    min-height: 100vh;
    width: 100%;
  }

  .tz-header {
    align-items: center;
    background: #ffffff;
    border-bottom: 1px solid rgba(28, 28, 28, 0.08);
    display: flex;
    justify-content: space-between;
    padding: 24px 40px;
    position: relative;
  }
  .tz-logo img {
    display: block;
    height: 45px;
  }
  .tz-nav {
    display: flex;
    gap: 24px;
    left: 50%;
    position: absolute;
    transform: translateX(-50%);
  }
  .tz-nav a {
    color: #1c1c1c;
    font-size: 14px;
    text-decoration: none;
    transition: color 0.3s;
  }
  .tz-nav a:hover {
    color: #1d4ae1;
  }
  .tz-deploy {
    align-items: center;
    background: #3f58ff;
    border-radius: 16px;
    color: #ffffff;
    display: inline-flex;
    font-size: 16px;
    height: 38px;
    padding: 0 20px;
    text-decoration: none;
  }

  .tz-content {
    align-items: center;
    display: flex;
    flex: 0 0 auto;
    flex-direction: column;
    padding: 30px 16px 48px;
    width: 100%;
  }
  .tz-badge {
    align-items: center;
    background: rgba(63, 88, 255, 0.1);
    border-radius: 16px;
    color: #1d4ae1;
    display: inline-flex;
    font-size: 16px;
    font-weight: 700;
    gap: 8px;
    height: 38px;
    padding: 0 16px;
  }
  .tz-badge-dot {
    background: #1d4ae1;
    border-radius: 50%;
    height: 8px;
    width: 8px;
  }
  .tz-mark {
    height: 60px;
    margin: 20px 0 14px;
    width: 60px;
  }
  .tz-heading {
    color: #1c1c1c;
    font-size: 80px;
    font-weight: 500;
    line-height: 1.1;
    margin: 0 0 36px;
    text-align: center;
  }
  .tz-heading-accent {
    color: #1d4ae1;
  }

  .tz-card {
    background: #ffffff;
    border-radius: 30px;
    box-shadow:
      8px 10px 20px rgba(46, 48, 63, 0.1),
      9px 8px 30px rgba(30, 69, 128, 0.05);
    display: flex;
    flex-direction: column;
    gap: 24px;
    max-width: 885px;
    padding: 32px 60px;
    width: 100%;
  }
  .tz-serving-row {
    align-items: center;
    display: flex;
    gap: 16px;
    width: 100%;
  }
  .tz-serving-label {
    font-size: 16px;
    font-weight: 500;
    white-space: nowrap;
  }
  .tz-serving-chip {
    align-items: center;
    background: rgba(29, 74, 225, 0.1);
    border-radius: 8px;
    display: flex;
    flex: 1;
    gap: 8px;
    height: 41px;
    justify-content: space-between;
    min-width: 0;
    padding: 0 12px;
  }
  .tz-serving-address {
    /* flex + min-width:0 so the address ellipsizes instead of forcing the
       page wider than the viewport on narrow screens. */
    flex: 1;
    font-size: 16px;
    font-weight: 500;
    min-width: 0;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
  .tz-copy-button {
    align-items: center;
    background: transparent;
    border: none;
    color: #1d4ae1;
    cursor: pointer;
    display: flex;
    padding: 0;
  }
  .tz-copy-button:active .tz-copy-icon {
    transform: scale(0.88);
  }
  .tz-copy-icon {
    animation: tz-pop 180ms ease-out;
    display: block;
    transition: transform 120ms ease-out;
  }
  .tz-copied-label {
    animation: tz-fade-in 160ms ease-out;
    color: #1d4ae1;
    font-size: 14px;
    font-weight: 500;
    margin-left: auto;
    white-space: nowrap;
  }
  @keyframes tz-pop {
    from {
      opacity: 0;
      transform: scale(0.6);
    }
    to {
      opacity: 1;
      transform: scale(1);
    }
  }
  @keyframes tz-fade-in {
    from {
      opacity: 0;
      transform: translateX(4px);
    }
    to {
      opacity: 1;
      transform: translateX(0);
    }
  }

  .tz-address-input {
    background: #ffffff;
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 8px;
    box-shadow: none;
    color: #1c1c1c;
    font-family: 'DM Sans', system-ui, sans-serif;
    font-size: 16px;
    font-weight: 500;
    height: 43px;
    padding: 0 16px;
    width: 100%;
  }
  .tz-address-input::placeholder {
    color: rgba(28, 28, 28, 0.5);
  }
  .tz-address-input:focus {
    border-color: #1d4ae1;
    outline: none;
  }

  .tz-button-row {
    display: flex;
    justify-content: center;
  }
  .tz-request {
    align-items: center;
    background: #3f58ff;
    border: none;
    border-radius: 16px;
    color: #ffffff;
    cursor: pointer;
    display: inline-flex;
    font-family: 'DM Sans', system-ui, sans-serif;
    font-size: 16px;
    height: 38px;
    justify-content: center;
    padding: 0 32px;
  }
  .tz-request:disabled {
    background: rgba(28, 28, 28, 0.2);
    cursor: default;
  }

  .tz-helper {
    color: #1c1c1c;
    font-size: 16px;
    font-weight: 500;
    line-height: 22px;
    margin: 0;
    text-align: center;
  }
  .tz-helper-link {
    color: inherit;
    text-decoration: underline;
  }

  .tz-powered {
    align-items: center;
    color: #1c1c1c;
    display: flex;
    gap: 12px;
    justify-content: center;
    margin-top: 32px;
  }
  .tz-powered img {
    display: block;
    height: 20px;
  }

  #hcaptcha {
    display: none;
  }

  /* Toasts are appended to <body> by bulma-toast, so they need global rules.
     Scoped to .tz-toast (see notify() in Faucet.svelte) so the other frontend
     types keep Bulma's styling. */
  :global(.notification.tz-toast) {
    --tz-toast-shadow:
      8px 10px 20px rgba(46, 48, 63, 0.1), 9px 8px 30px rgba(30, 69, 128, 0.05);
    align-items: center;
    background: #ffffff;
    border: 1px solid rgba(28, 28, 28, 0.1);
    border-radius: 16px;
    box-shadow: var(--tz-toast-shadow);
    color: #1c1c1c;
    display: flex;
    font-family: 'DM Sans', system-ui, sans-serif;
    font-size: 16px;
    font-weight: 500;
    gap: 12px;
    line-height: 22px;
    margin-bottom: 24px;
    max-width: min(620px, calc(100vw - 32px));
    padding: 14px 20px;
    text-align: left;
    /* Messages carry tx hashes: unbreakable tokens must wrap, not overflow. */
    min-width: 0;
    overflow-wrap: anywhere;
    word-break: break-word;
  }
  /* Semantic status accent (inset so it follows the rounded corners), kept
     separate from the brand blue. */
  :global(.notification.tz-toast.is-success) {
    box-shadow: inset 4px 0 0 0 #12b76a, var(--tz-toast-shadow);
  }
  :global(.notification.tz-toast.is-warning),
  :global(.notification.tz-toast.is-danger) {
    box-shadow: inset 4px 0 0 0 #e17e26, var(--tz-toast-shadow);
  }
  :global(.notification.tz-toast .delete) {
    background: rgba(28, 28, 28, 0.06);
    flex: 0 0 auto;
    height: 20px;
    margin: 0;
    max-height: 20px;
    max-width: 20px;
    min-height: 20px;
    min-width: 20px;
    order: 2;
    /* relative, not static: Bulma draws the cross with absolutely positioned
       pseudo-elements sized off the containing block, so a static button would
       stretch them across the page. */
    position: relative;
    right: auto;
    top: auto;
    transition: background 150ms ease-out;
    width: 20px;
  }
  :global(.notification.tz-toast .delete:hover) {
    background: rgba(28, 28, 28, 0.14);
  }
  /* Bulma draws the cross with white pseudo-elements; on a white toast they
     need to be ink. */
  :global(.notification.tz-toast .delete::before),
  :global(.notification.tz-toast .delete::after) {
    background-color: #1c1c1c;
  }

  @media (max-width: 900px) {
    .tz-nav {
      display: none;
    }
    .tz-heading {
      font-size: 40px;
    }
    .tz-card {
      padding: 24px;
    }
    .tz-serving-row {
      align-items: flex-start;
      flex-direction: column;
      gap: 8px;
    }
    .tz-serving-chip {
      width: 100%;
    }
  }
</style>
