'use client';

import { ConnectButton } from '@rainbow-me/rainbowkit';
import Image from 'next/image';
import CustomGrid from '@/components/CustomGrid';

import { Wave } from '@/components/Wave';

export default function Home() {
  return (
    <main className="flex min-h-screen w-full flex-col items-center justify-between">
      <div className="w-full flex-grow bg-white">
        <div className="flex justify-between px-20 py-10">
          <Image alt="logo" height={40} src="/images/logo.svg" width={40} />
          <ConnectButton />
        </div>
        <div className="my-4 flex w-full items-center justify-center text-3xl text-green-600">
          Web3 Sounds
        </div>
        <div className="w-fullr flex items-center justify-center text-3xl">
          <h1 className="text-green-600">永久记录每一个人的声音</h1>
          {/* <Image alt="logo" height={40} src="/images/soundWave.svg" width={40} /> */}
          <Wave />
        </div>
      </div>
      <div className="h-96 w-full justify-center bg-[url('/images/background.jpeg')] bg-cover bg-center">
        <CustomGrid />
        <a
          className="mt-16 flex h-10 items-center justify-center text-xl text-purple-400 underline hover:text-2xl"
          href="/uploader"
        >
          记录你的声音
        </a>
      </div>
      <div className="flex h-40 w-full items-center justify-center bg-green-600 text-white">
        <Image alt="github" height={24} src="/images/github.svg" width={24} />
        <div className="ml-4">链接声音，一起探索未来世界计算机</div>
      </div>
    </main>
  );
}
