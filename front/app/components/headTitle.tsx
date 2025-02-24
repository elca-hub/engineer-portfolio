import React from 'react';
import Head from 'next/head';

interface HeadTitleProps {
  title: string;
  des: string;
}

const HeadTitle = ({title, des}: HeadTitleProps) => {
  return (
    <Head>
      <title>{title} | DevPort</title>
      <meta name="description" content={des} />
    </Head>
  )
}

export default HeadTitle;
