'use client';

import { ConnectButton } from '@rainbow-me/rainbowkit';
import Image from 'next/image';

export default function Header() {
  return (
    <div className="flex justify-between px-20 py-10">
      <Image alt="logo" height={40} src="/images/logo.svg" width={40} />
      <ConnectButton />
    </div>
  );
}
