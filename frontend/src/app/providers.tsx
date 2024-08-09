'use client';

import { RainbowKitProvider, darkTheme, getDefaultConfig } from '@rainbow-me/rainbowkit';
import '@rainbow-me/rainbowkit/styles.css';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { WagmiProvider } from 'wagmi';

import { GlobalProvider } from './GlobalContext';

const w3q = Object.freeze({
  id: 3334,
  name: 'Web3Q Galileo',
  nativeCurrency: {
    decimals: 18,
    name: 'W3Q',
    symbol: 'W3Q'
  },
  rpcUrls: {
    default: { http: ['https://galileo.web3q.io:8545'] }
  }
});

const config = getDefaultConfig({
  appName: 'Web3 App',
  chains: [w3q],
  projectId: process.env.NEXT_PUBLIC_WC_PROJECT_ID || 'df2b40017e56740f385b201c026e1247',
  ssr: true
});

const queryClient = new QueryClient();

export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <GlobalProvider>
      <WagmiProvider config={config}>
        <QueryClientProvider client={queryClient}>
          <RainbowKitProvider
            theme={darkTheme({
              accentColor: '#1d951b',
              accentColorForeground: 'white',
              borderRadius: 'medium',
              fontStack: 'system',
              overlayBlur: 'small'
            })}
          >
            {children}
          </RainbowKitProvider>
        </QueryClientProvider>
      </WagmiProvider>
    </GlobalProvider>
  );
}
