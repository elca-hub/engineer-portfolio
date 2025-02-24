import { ColorMode } from '@/types/colorMode';
import React from 'react';

interface ButtonProps {
  mode: ColorMode;
  children: React.ReactNode;
}

const baseStyle = 'px-4 py-2 rounded font-bold text-lg hover:opacity-80 hover:scale-[0.98] transition-all duration-300';

const Button = ({ mode, children }: ButtonProps) => {
  return (
    <button
      className={`${baseStyle} ${mode === 'primary' ? 'bg-primary text-foreground' : 'bg-secondary text-black'}`}
    >
      {children}
    </button>
  );
};

export default Button;