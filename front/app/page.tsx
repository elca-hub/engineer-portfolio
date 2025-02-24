import { Metadata } from 'next';
import { setMetadata } from '@/utils/setMetadata';
import HomeClient from '@/clients/home';


const Home = () => {
  return (
    <HomeClient />
  );
}

export const metadata: Metadata = setMetadata('Home', 'DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。');

export default Home;
