'use client'

import { RiAddLine, RiLoginBoxLine } from 'react-icons/ri';
import { motion } from 'framer-motion';
import Link from 'next/link';
import { ButtonStyle } from '@/constants/tailwindConstant';
import TextWithIcon from '@/components/ui/text/textWithIcon';
import DPButton from '@/components/ui/button/button';

const LoginButtons = () => {
  return (
    <motion.div
      className='flex items-center justify-center md:gap-28 gap-10'
    >
      <Link href='/login'>
        <DPButton colormode='primary'>
          <TextWithIcon icon={<RiLoginBoxLine />}>
            ログイン
          </TextWithIcon>
        </DPButton>
      </Link>

      <Link href='/register'>
        <DPButton colormode='secondary'>
          <TextWithIcon icon={<RiAddLine />}>
            新規登録
          </TextWithIcon>
        </DPButton>
      </Link>
    </motion.div>
  );
};

export default LoginButtons;
