import { brands, DEFAULT_BRAND } from '../brands/index.js';

// Brand is chosen at BUILD time via VITE_BRAND (see vite.config.js / Dockerfile).
const requested = import.meta.env.VITE_BRAND || DEFAULT_BRAND;

if (!brands[requested]) {
  console.warn(
    `[brand] unknown VITE_BRAND "${requested}", falling back to "${DEFAULT_BRAND}"`,
  );
}

export const brand = brands[requested] ?? brands[DEFAULT_BRAND];

const kebab = (s) => s.replace(/[A-Z]/g, (m) => '-' + m.toLowerCase());

// Flatten the nested tokens object into a CSS custom-property block, e.g.
// { color: { primary: '#3F58FF' }, radius: { card: '24px' } }
//   → "--fc-color-primary: #3F58FF; --fc-radius-card: 24px;"
export function brandCssVars(tokens = brand.tokens) {
  const lines = [];
  for (const [group, entries] of Object.entries(tokens)) {
    for (const [key, value] of Object.entries(entries)) {
      lines.push(`--fc-${kebab(group)}-${kebab(key)}: ${value};`);
    }
  }
  return lines.join('\n  ');
}
