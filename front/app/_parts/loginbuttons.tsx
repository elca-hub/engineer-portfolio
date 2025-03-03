'use client'

import { RiAddLine, RiLoginBoxLine } from 'react-icons/ri';
import { motion } from 'framer-motion';
import ButtonIcon from '@/components/ui/button/buttonIcon';

const LoginButtons = () => {
  return (
    <motion.div
      className='flex items-center justify-center md:gap-28 gap-10'
    >
      <ButtonIcon mode="secondary" label="ログイン">
        <RiLoginBoxLine />
      </ButtonIcon>

      <ButtonIcon mode="primary" label="新規登録">
        <RiAddLine />
      </ButtonIcon>
    </motion.div>
  );
};

export default LoginButtons;
