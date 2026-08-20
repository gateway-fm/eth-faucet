// TEIZA design tokens (feeds CSS custom properties consumed by the TEIZA
// component override). Edit here to retune the TEIZA palette/typography.
export default {
  color: {
    primary: '#3F58FF', // buttons, badge dot
    accent: '#1D4AE1', // "Receive" heading, links, nav hover, copy icon
    text: '#1C1C1C',
    badgeBg: '#EAEDFF',
    addressBg: '#E7ECFB',
    cardBg: '#FFFFFF',
    inputBorder: '#EAEAEA',
    inputPlaceholder: '#909090',
    navText: '#1C1C1C',
    headerBorder: '#EDEDF0',
    gradientTop: '#FFFFFF',
    gradientBottom: '#E5EDFF',
  },
  font: {
    // DM Sans is vendored + self-hosted in /public/fonts (see App.svelte @font-face).
    family:
      "'DM Sans', system-ui, -apple-system, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif",
  },
  radius: {
    card: '24px',
    input: '10px',
    pill: '999px',
  },
};
