import React from 'react';
import { ColorMode } from '@/types/colorMode';
import Button from './button';

interface ButtonIconProps {
  mode: ColorMode;
  label: string;
  children: React.ReactNode;
  href?: string;
}

const ButtonIcon = ({mode, label, href, children}: ButtonIconProps) => {
  return (
    <Button mode={mode} href={href}>
      <div className='flex items-center justify-center gap-2'>
        {children}
        {label}
      </div>
    </Button>
  )
}

export default ButtonIcon;
