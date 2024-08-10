import type { Config } from 'tailwindcss';

const config: Config = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}'
  ],
  plugins: [],
  theme: {
    extend: {
      animation: {
        wave: 'wave 0.66s linear infinite',
        wave2: 'wave 0.8s linear infinite',
        wave3: 'wave 0.7s linear infinite',
        wave4: 'wave 0.5s linear infinite',
        wave5: 'wave 0.9s linear infinite',
        wave6: 'wave 1.2s linear infinite'
      },
      backgroundImage: {
        'gradient-conic': 'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))',
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))'
      },
      keyframes: {
        wave: {
          '0%': { height: '8px' },
          '100%': { height: '12px' },
          '50%': { height: '32px' }
        }
      }
    }
  }
};

export default config;
