import tokens from './tokens.js';
import Redesign from './components/Redesign.svelte';
import favicon from './assets/tokenIcon.png';

// TEIZA brand. Its Figma restructures the redesign enough (centered CTA, circle
// token-mark, badge dot, DM Sans, spacing) that it ships a full component
// override rather than tokens alone. Deployment identity (network name, symbol,
// logo, logo link, paid_customer) still comes from the backend via /api/info.
export default {
  id: 'teiza',
  tokens,
  assets: {
    favicon,
  },
  overrides: {
    redesign: Redesign, // replaces the shared redesign screen for this brand
  },
};
