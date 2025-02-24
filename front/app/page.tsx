import { Metadata } from 'next';
import { setMetadata } from './utils/setMetadata';
import Image from 'next/image';
import { Pacifico } from 'next/font/google';
export const metadata: Metadata = setMetadata('Home', 'DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。');

const pacifico = Pacifico({
  weight: '400',
  subsets: ['latin'],
});

export default function Home() {
  return (
    <div>
      <header className='flex items-center justify-center gap-16 py-16 bg-gradient-to-b from-lightblue to-background'>
        <Image src="/logo.webp" alt="DevPort" width={200} height={200} />
        <div className='text-center'>
          <h1 className={`${pacifico.className} text-6xl mb-6`}>
            <span className='text-primary'>Dev</span><span className='text-secondary'>Port</span>
          </h1>
          <p className={`${pacifico.className} text-subtext text-2xl`}>
            <span className='text-primary'>Dev</span>eloper <span className='text-secondary'>Port</span>folio, Redefined.
          </p>
        </div>
      </header>
      <div>
        
      </div>
    </div>
  );
}
