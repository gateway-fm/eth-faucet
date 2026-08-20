<script>
  import CopyButton from './CopyButton.svelte';
  import Navigation from './Navigation.svelte';
  import tokenIcon from '../assets/tokenIcon.png';

  export let faucetInfo;
  export let input;
  export let handleRequest;
  export let gweiToEth;

  $: paidCustomer = faucetInfo.paid_customer;

  const openMessageWindow = (subject, email) => {
    const emailSupport = 'support+presto@gateway.fm';
    const mailtoLink = `mailto:${emailSupport}?subject=${encodeURIComponent(subject)}`;

    window.location.href = mailtoLink;
  };

  function autoResize(event) {
    const textarea = event.target;
    textarea.style.height = 'auto'; // Reset height
    textarea.style.height = `${textarea.scrollHeight}px`; // Set new height
  }
</script>

<main>
  <section class="hero is-info is-fullheight">
    <div class="hero-head">
      <nav class="navbar">
        <div class="header-container" class:paid={paidCustomer}>
          <div class="navbar-brand">
            <a class="navbar-item" href={faucetInfo.logo_link}>
              <span class="icon icon-brand">
                <img src={faucetInfo.logo_url} alt="logo" />
              </span>
            </a>
          </div>
          <div class="navbar-desktop nav-center">
            {#if !paidCustomer}
              <Navigation />
            {/if}
          </div>
          <div class="navbar-end">
            {#if !paidCustomer}
              <a
                class="navbar-desktop"
                href="https://presto.gateway.fm/onboarding"
                target="_blank"
                rel="noopener noreferrer"
              >
                <button class="button is-primary is-rounded"> Deploy rollup </button></a
              >
            {/if}
          </div>
          <div class="navbar-mobile">
            {#if !paidCustomer}
              <Navigation />
            {/if}
          </div>
        </div>
      </nav>
    </div>

    <div class="hero-body">
      <div class="container has-text-centered">
        <div class="centered-column">
          <div class="network">
            <span class="dot"></span>
            <div>{faucetInfo.network}</div>
          </div>
          {#if tokenIcon}
            <div class="token-icon">
              <img src={tokenIcon} alt="token" />
            </div>
          {/if}
          <div class="title">
            <span class="receive">Receive</span>
            <span class="gas-token">
              {gweiToEth(faucetInfo.payout)}
              {faucetInfo.symbol}
            </span>
          </div>
          <div id="hcaptcha" data-size="invisible"></div>
          <div class="card">
            <div>
              <div class="subtitle">
                <div class="serving-label">Serving from</div>
                <div class="address-from">
                  <span class="address-text">{faucetInfo.account}</span>
                  <CopyButton text={faucetInfo.account} />
                </div>
              </div>
            </div>
            <div class="field is-grouped">
              <div class="control is-expanded">
                <textarea
                  bind:value={input}
                  class="input"
                  rows="1"
                  on:input={autoResize}
                  placeholder="Enter your address"
                ></textarea>
              </div>
              <div class="control">
                <button
                  on:click={() => handleRequest(input)}
                  class="button is-primary is-rounded"
                >
                  Request
                </button>
                {#if !paidCustomer}
                  <div class="box-offer">
                    Claim {gweiToEth(faucetInfo.payout)}
                    {faucetInfo.symbol} test token for development.<br />
                    If you need additional tokens for extensive testing, please
                    <!-- svelte-ignore a11y-invalid-attribute -->
                    <a
                      class="link"
                      href="#"
                      role="button"
                      on:click={() =>
                        openMessageWindow('Additional tokens request')}
                    >
                      contact support
                    </a>
                  </div>
                {/if}
              </div>
            </div>
          </div>
          {#if !paidCustomer}
            <div class="deploy-btn-mobile">
              <a
                class="deploy-link-mobile"
                href="https://presto.gateway.fm/onboarding"
                target="_blank"
                rel="noopener noreferrer"
              >
                <button class="button is-primary is-rounded">
                  Deploy rollup
                </button></a
              >
            </div>
          {/if}
          <div class="box-logo">
            Powered by <a
              class="navbar-item"
              href={faucetInfo.logo_link}
              target="_blank"><img src={faucetInfo.logo_url} alt="logo" /></a
            >
          </div>
        </div>
      </div>
    </div>
  </section>
</main>

<style>
  .deploy-btn-mobile {
    display: none;
  }
  .input {
    resize: none;
    overflow: hidden;
    width: 100%;
    height: auto;
    line-height: 24px;
    background-color: var(--fc-color-card-bg, #ffffff);
    border: 1px solid var(--fc-color-input-border, #eaeaea);
    border-radius: var(--fc-radius-input, 10px);
    padding: 16px 20px;
    font-size: 16px;
    color: var(--fc-color-text, #1c1c1c);
    box-shadow: none;
  }
  .input::placeholder {
    color: var(--fc-color-input-placeholder, #909090);
  }
  .input:focus {
    border-color: var(--fc-color-primary, #3f58ff);
    box-shadow: 0 0 0 3px rgba(63, 88, 255, 0.12);
  }
  .header-container {
    display: flex;
    width: 100%;
    padding-inline: 32px;
    justify-content: space-between;
    align-items: center;
    gap: 16px;
  }
  .nav-center {
    flex: 1;
    display: flex;
    justify-content: center;
  }
  /* Paid customer: header is only the logo — center it */
  .header-container.paid {
    justify-content: center;
  }
  .header-container.paid .nav-center,
  .header-container.paid .navbar-end,
  .header-container.paid .navbar-mobile {
    display: none;
  }
  .hero {
    padding-top: 0;
  }
  /* Top-align the body just below the header instead of vertically centering it */
  .hero.is-fullheight .hero-body {
    align-items: flex-start;
    padding-top: 40px;
  }
  .hero-head {
    background-color: #ffffff;
    border-bottom: 1px solid var(--fc-color-header-border, #ededf0);
    padding: 16px 0;
  }
  .box-logo {
    display: flex;
    justify-content: center;
    align-items: center;
    color: var(--fc-color-text, #1c1c1c);
    margin-top: 32px;
    gap: 8px;
    font-size: 14px;
    font-weight: 700;
  }
  .box-logo img {
    height: 20px;
  }
  .network {
    display: inline-flex;
    justify-content: center;
    align-items: center;
    gap: 8px;
    font-size: 15px;
    font-weight: 700;
    padding: 8px 16px;
    background-color: var(--fc-color-badge-bg, #eaedff);
    border-radius: var(--fc-radius-pill, 999px);
    color: var(--fc-color-accent, #1d4ae1);
  }
  .dot {
    width: 8px;
    height: 8px;
    border-radius: 50%;
    background-color: var(--fc-color-accent, #1d4ae1);
  }
  .token-icon {
    margin-top: 24px;
    margin-bottom: 8px;
  }
  .token-icon img {
    width: 68px;
    height: 68px;
    display: block;
  }
  .link {
    color: var(--fc-color-accent, #1d4ae1) !important;
    text-decoration: underline;
    cursor: pointer;
  }

  .button.is-primary {
    display: inline-flex;
    justify-content: center;
    align-items: center;
    width: auto;
    background-color: var(--fc-color-primary, #3f58ff);
    border: none;
    color: #ffffff;
    font-weight: 600;
    padding: 12px 28px;
    gap: 12px;
    height: auto;
  }

  .button:hover {
    opacity: 0.9;
  }

  .box-offer {
    font-weight: 400;
    font-size: 14px;
    line-height: 22px;
    letter-spacing: 0px;
    color: var(--fc-color-text, #1c1c1c);
    text-align: center;
  }
  .address-from {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 12px;
    padding: 12px 16px;
    border-radius: var(--fc-radius-input, 10px);
    background-color: var(--fc-color-address-bg, #e7ecfb);
    flex: 1;
    min-width: 0;
  }
  .address-text {
    flex: 1;
    min-width: 0;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    text-align: left;
    color: var(--fc-color-text, #1c1c1c);
    font-size: 16px;
  }
  .field {
    display: flex;
    flex-direction: column;
    gap: 24px;
  }

  .gas-token {
    color: var(--fc-color-text, #1c1c1c);
  }

  .card {
    display: flex;
    flex-direction: column;
    min-width: 100%;
    gap: 24px;
    box-shadow: 0px 12px 32px rgba(20, 40, 120, 0.08);
    border-radius: var(--fc-radius-card, 24px);
    padding: 40px;
    background-color: var(--fc-color-card-bg, #ffffff);
    color: var(--fc-color-text, #1c1c1c);
  }

  .title {
    display: inline-flex;
    align-items: baseline;
    color: var(--fc-color-text, #1c1c1c);
    gap: 20px;
    font-weight: 500;
    font-size: 84px;
    line-height: 92px;
    letter-spacing: -1px;
    margin: 0;
  }
  .receive {
    color: var(--fc-color-accent, #1d4ae1);
  }
  .hero.is-info {
    background: linear-gradient(
      180deg,
      var(--fc-color-gradient-top, #ffffff) 0%,
      var(--fc-color-gradient-bottom, #e5edff) 100%
    );
    color: var(--fc-color-text, #1c1c1c);
  }

  .serving-label {
    color: var(--fc-color-text, #1c1c1c);
    flex-shrink: 0;
    white-space: nowrap;
  }

  .subtitle {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 24px;
    font-size: 16px;
    font-weight: 400;
    letter-spacing: 0px;
    color: var(--fc-color-text, #1c1c1c);
  }

  .navbar-mobile {
    display: none;
  }

  .hero .subtitle {
    line-height: 1.5;
  }
  .centered-column {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
    max-width: 880px;
    margin: 0 auto;
  }

  .control {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 20px;
  }
  .control:not(.is-expanded) {
    align-items: center;
  }

  .navbar-item {
    color: var(--fc-color-accent, #1d4ae1) !important;
    padding: 0;
  }
  .navbar-item:hover {
    background-color: transparent !important;
    cursor: pointer;
  }

  .icon {
    width: 16px;
    height: 16px;
  }

  .icon-brand {
    width: auto;
    height: 34px;
  }
  .icon-brand img {
    height: 34px;
    width: auto;
    display: block;
  }

  @media (max-width: 992px) {
    .subtitle {
      font-size: 14px;
    }

    .navbar-mobile {
      display: block;
    }
  }

  @media (max-width: 768px) {
    .subtitle {
      flex-direction: column;
      align-items: stretch;
      font-size: 16px;
      gap: 12px;
    }

    .title {
      font-size: 40px;
      line-height: 48px;
      gap: 12px;
    }
    .navbar-desktop {
      display: none;
    }

    .nav-center {
      display: none;
    }

    .hero-body {
      display: flex;
      justify-content: center;
      align-items: flex-start;
      padding: 2rem 1.5rem;
    }
    .hero-body .container {
      width: 100%;
      max-width: 100%;
    }
    .card {
      min-width: 0;
      width: 100%;
    }
    .deploy-link-mobile {
      display: flex;
      width: 80%;
    }
    .deploy-link-mobile .button.is-primary {
      width: 100%;
    }

    .deploy-btn-mobile {
      display: flex;
      justify-content: center;
      margin-top: 16px;
      min-width: 100%;
    }
    .card {
      padding: 24px;
    }
  }
</style>
