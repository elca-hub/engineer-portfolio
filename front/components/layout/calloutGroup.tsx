"use client";

import { CalloutContext, calloutItemType } from "@/app/state";
import { Callout, Flex } from "@radix-ui/themes";
import { AnimatePresence, motion } from "motion/react";
import { useContext, useEffect, useState } from "react";
import { RiAlarmWarningLine, RiErrorWarningLine, RiInfoCardLine, RiInfoI, RiThumbUpLine } from "react-icons/ri";

export default function CalloutGroup() {
  const {callout, setCallout} = useContext(CalloutContext);

  const convertType = (type: 'info' | 'warn' | 'error' | 'success') => {
    switch (type) {
      case 'info':
        return 'blue';
      case 'warn':
        return 'yellow';
      case 'error':
        return 'red';
      case 'success':
        return 'green';
    }
  }

  const convertIcon = (type: 'info' | 'warn' | 'error' | 'success') => {
    switch (type) {
      case 'info':
        return <RiInfoI />;
      case 'warn':
        return <RiAlarmWarningLine />;
      case 'error':
        return <RiErrorWarningLine />;
      case 'success':
        return <RiThumbUpLine />;
    }
  }

  useEffect(() => {
    if (callout.length > 0) {
      const timer = setTimeout(() => {
        const remove = callout.slice(1);
        setCallout(remove);
      }, 5000)

      return () => clearTimeout(timer);
    }
  }, [callout]);

  return (
    <div className="fixed top-0 right-0 p-4 z-50">
      <AnimatePresence>
        {callout.map((item, index) => (
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.3 }}
            className="mb-2"
            exit={{ opacity: 0, y: 20 }}
            key={index}
          >
            <Callout.Root color={convertType(item.type)}>
              <Callout.Icon>
                {convertIcon(item.type)}
              </Callout.Icon>
              <Callout.Text>
                {item.content}
              </Callout.Text>
            </Callout.Root>
          </motion.div>
        ))}
      </AnimatePresence>
    </div>
  )
}
