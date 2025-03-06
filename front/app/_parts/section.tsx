'use client'

import Image from 'next/image';
import React from 'react';
import { motion } from 'framer-motion';
import { BudouXText } from '@/components/ui/text/budouxText';

export interface ButtonProps {
  title: string;
  contents: string;
  imagePath: string;
  imageAlt: string | null;
  imageWidth: number;
  imageHeight: number;
}

const duration = 1.0;

const HomeSection = ({ title, contents, imagePath, imageAlt, imageWidth, imageHeight}: ButtonProps) => {
  return (
    <section className='mx-10 flex h-screen flex-col items-center justify-center gap-10 md:flex-row md:gap-20'>
      <div>
        <motion.h3
          className='mb-4 text-4xl font-bold text-foreground md:mb-8'
          initial={{ y: 50, opacity: 0 }}
          whileInView={{
            y: 0,
            opacity: 1,
            transition: {
              duration: duration
            }
          }}
          viewport={{ once: true }}
        >
          {title}
        </motion.h3>
        <motion.div
          className='whitespace-pre-line text-lg text-subtext md:text-2xl'
          style={{ whiteSpace: 'aut-phrase' }}
          initial={{ y: -10, opacity: 0 }}
          whileInView={{
            y: 0,
            opacity: 1,
            transition: {
              duration: duration
            }
          }}
          viewport={{ once: true }}
        >
          <BudouXText text={contents}></BudouXText>
        </motion.div>
      </div>
      <motion.div
        initial={{ x: -50, opacity: 0 }}
        whileInView={{
          x: 0,
          opacity: 1,
          transition: {
            duration: duration
          }
        }}
        viewport={{ once: true }}
        className='md:max-w-[50%]'
      >
        <Image src={imagePath} alt={imageAlt ?? ''} width={imageWidth} height={imageHeight} loading='lazy' />
      </motion.div>
    </section>
  );
};

export default HomeSection;