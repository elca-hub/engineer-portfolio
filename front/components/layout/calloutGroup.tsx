"use client";

import { CalloutContext } from "@/app/state";
import { AnimatePresence, motion } from "motion/react";
import { JSX, useContext, useEffect } from "react";
import { RiAlarmWarningLine, RiErrorWarningLine, RiInfoI, RiThumbUpLine } from "react-icons/ri";

export default function CalloutGroup() {
  const {callout, setCallout} = useContext(CalloutContext);

  const convertType = (type: 'info' | 'warn' | 'error' | 'success'): {
    bgColor: string;
    textColor: string;
    icon: JSX.Element;
  } => {
    switch (type) {
      case 'info':
        return {
          bgColor: 'bg-blue-100',
          textColor: 'text-blue-500',
          icon: <RiInfoI />
        }
      case 'warn':
        return {
          bgColor: 'bg-yellow-100',
          textColor: 'text-yellow-500',
          icon: <RiAlarmWarningLine />
        }
      case 'error':
        return {
          bgColor: 'bg-red-100',
          textColor: 'text-red-500',
          icon: <RiErrorWarningLine />
        }
      case 'success':
        return {
          bgColor: 'bg-green-100',
          textColor: 'text-green-700',
          icon: <RiThumbUpLine />
        }
    }
  }

  useEffect(() => {
    if (callout.length > 0) {
      const timer = setTimeout(() => {
        const remove = callout.slice(1);
        setCallout(remove);
      }, 5000);

      return () => clearTimeout(timer);
    }
  }, [callout, setCallout]);

  return (
    <div className="fixed right-0 top-0 z-50 p-4">
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
            <div className={`${convertType(item.type).bgColor} ${convertType(item.type).textColor} flex items-center gap-2 rounded-md p-3`}>
              {convertType(item.type).icon}
              <p>{item.content}</p>
            </div>
          </motion.div>
        ))}
      </AnimatePresence>
    </div>
  )
}
