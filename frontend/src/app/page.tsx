import Image from 'next/image';

import { Button } from '@/components/Button';

export default function Home() {
  return (
    <main className="flex min-h-screen w-full flex-col items-center justify-between">
      <div className="w-full flex-grow bg-white">
        <div className="flex justify-between px-20 py-10">
          <Image alt="logo" height={24} src="/images/logo.svg" width={24} />
          <Button className="bg-white">链接钱包</Button>
        </div>
        <div className="align-center flex justify-center text-3xl">
          <h1 className="text-purple-400">永久记录每一个人的声音</h1>
        </div>
      </div>
      <div className="flex h-96 w-full bg-[url('/images/background.jpeg')] bg-cover bg-center">
        {/* TODO 色块儿 */}
      </div>
      <div className="flex h-40 w-full items-center justify-center bg-purple-400 text-white">
        <Image alt="github" height={24} src="/images/github.svg" width={24} />
        <div className="ml-4">链接声音，一起探索未来世界计算机</div>
      </div>
    </main>
  );
}
