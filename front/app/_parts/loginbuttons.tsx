'use client'

import { RiAddLine, RiLoginBoxLine } from 'react-icons/ri';
import { motion } from 'framer-motion';
import Link from 'next/link';
import { ButtonStyle } from '@/constants/tailwindConstant';
import TextWithIcon from '@/components/ui/text/textWithIcon';

const LoginButtons = () => {
  return (
    <motion.div
      className='flex items-center justify-center md:gap-28 gap-10'
    >
      <Link className={ButtonStyle("primary")} href='/login'>
        <TextWithIcon icon={<RiLoginBoxLine />}>
          ログイン
        </TextWithIcon>
      </Link>

      <Link className={ButtonStyle("secondary")} href='/register'>
        <TextWithIcon icon={<RiAddLine />}>
          新規登録
        </TextWithIcon>
      </Link>
    </motion.div>
  );
};

export default LoginButtons;
