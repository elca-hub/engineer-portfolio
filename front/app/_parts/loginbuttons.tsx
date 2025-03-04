'use client'

import { RiAddLine, RiLoginBoxLine } from 'react-icons/ri';
import { motion } from 'framer-motion';
import ButtonIcon from '@/components/ui/button/buttonIcon';
import Link from 'next/link';

const LoginButtons = () => {
  return (
    <motion.div
      className='flex items-center justify-center md:gap-28 gap-10'
    >
      <Link href='/login' passHref>
        <ButtonIcon mode="secondary" label="ログイン" href='/login'>
          <RiLoginBoxLine />
        </ButtonIcon>
      </Link>

      <Link href='/register' passHref>
        <ButtonIcon mode="primary" label="新規登録" href='/register'>
          <RiAddLine />
        </ButtonIcon>
      </Link>
    </motion.div>
  );
};

export default LoginButtons;
