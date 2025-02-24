'use client'

import Image from 'next/image';
import React from 'react';
import { motion } from 'framer-motion';
import { BudouXText } from '@/components/bundouxText';

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
    <section className='h-screen mx-10 flex items-center justify-center gap-10 md:gap-20 md:flex-row flex-col'>
      <div>
        <motion.h3
          className='text-4xl md:mb-8 mb-4 text-foreground font-bold'
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
          className='md:text-2xl text-lg text-subtext whitespace-pre-line'
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
        <Image src={imagePath} alt={imageAlt ?? ''} width={imageWidth} height={imageHeight} />
      </motion.div>
    </section>
  );
};

export default HomeSection;