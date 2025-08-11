'use client';
import React, { createContext, ReactNode, useState } from 'react';

export const HeaderButtonContext = createContext<ProviderContext | undefined>(
  undefined
);

type ProviderContext = {
  headerButton: React.ReactNode | null;
  setHeaderButton: React.Dispatch<React.SetStateAction<React.ReactNode | null>>;
};

export function HeaderButtonProvider({ children }: { children: ReactNode }) {
  const [headerButton, setHeaderButton] = useState<React.ReactNode | null>(null);

  const stateObject = {
    headerButton,
    setHeaderButton,
  };

  return (
    <HeaderButtonContext.Provider value={stateObject}>
      {children}
    </HeaderButtonContext.Provider>
  );
}
export default HeaderButtonProvider;
