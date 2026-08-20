import tokens from './tokens.js';

// Gateway default brand: uses the shared redesign component as-is (no override),
// and the env-configured favicon (VITE_FAVICON_PATH) — i.e. original behaviour.
export default {
  id: 'base',
  tokens,
  assets: {}, // no brand favicon → App falls back to VITE_FAVICON_PATH
  overrides: {}, // no component override → shared redesign is used
};
