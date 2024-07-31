import Image from 'next/image';

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <Image src="/profile.png" width={500} height={500} alt="Picture of the author" />
      <div>home</div>
    </main>
  );
}
