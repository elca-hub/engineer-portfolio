"use client";

import React from "react";
import { useState } from "react";

export type calloutItemType = {
  content: string;
  type: 'info' | 'warn' | 'error' | 'success';
}

export const CalloutContext = React.createContext({
  callout: [] as calloutItemType[],
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  setCallout: (callout: calloutItemType[]) => {}
})

export const MainProvider = ({children}: {children: React.ReactNode}) => {
  const [callout, setCallout] = useState<calloutItemType[]>([]);

  return (
    <CalloutContext.Provider value={{callout, setCallout}}>
      {children}
    </CalloutContext.Provider>
  )
}
