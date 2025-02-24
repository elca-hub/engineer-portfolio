import React from 'react';
import { ColorMode } from '@/types/colorMode';
import Button from './button';

interface ButtonIconProps {
  mode: ColorMode;
  label: string;
  children: React.ReactNode;
}

const ButtonIcon = ({mode, label, children}: ButtonIconProps) => {
  return (
    <Button mode={mode}>
      <div className='flex items-center justify-center gap-2'>
        {children}
        {label}
      </div>
    </Button>
  )
}

export default ButtonIcon;
