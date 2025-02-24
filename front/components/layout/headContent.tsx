'use client'

import React from 'react';

interface HeadContentProps {
  title: string;
  des: string;
}

const HeadContent = ({title, des}: HeadContentProps) => {
  return (
    <>
      <title>{`${title} | DevPort`}</title>
      <meta name='description' content={des} />
    </>
  )
}

export default HeadContent;
