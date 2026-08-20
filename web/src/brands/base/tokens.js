// Gateway (default) design tokens — the original "new-design" (purple).
// The base brand uses the shared, un-tokenized redesign component verbatim, so
// only `font.family` is actually consumed today (via App.svelte's global rule);
// the colors are kept here for reference / future token-only brands.
export default {
  color: {
    primary: '#8950FA', // buttons
    accent: '#8950FA', // "Receive", links, nav hover
    text: '#161718',
    badgeBg: '#DCF0FD',
    addressBg: '#EEE8FF',
    cardBg: '#FFFFFF',
    inputBorder: '#DBDBDB',
    inputPlaceholder: '#7A7A7A',
    navText: '#676E73',
    headerBorder: 'transparent',
    gradientTop: '#FFFFFF',
    gradientBottom: '#FFFFFF',
  },
  font: {
    // Bulma default sans-serif stack (what the original design rendered in).
    family:
      'BlinkMacSystemFont, -apple-system, "Segoe UI", "Roboto", "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue", "Helvetica", "Arial", sans-serif',
  },
  radius: {
    card: '19px',
    input: '9999px',
    pill: '9999px',
  },
};
