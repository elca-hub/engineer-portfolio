import React, { useEffect, useState } from 'react';
import { ControllerFieldState, ControllerRenderProps, UseFormRegisterReturn } from 'react-hook-form';
import TextWithIcon from '../text/textWithIcon';
import { RiQuestionLine } from 'react-icons/ri';

interface InputFieldProps {
  title: string;
  type: string;
  field: ControllerRenderProps<any, any>;
  fieldState: ControllerFieldState;
  isRequired?: boolean;
  helperText?: string;
  autoFocus?: boolean;
}

const InputField = ({title, type, field, fieldState, isRequired, helperText, autoFocus}: InputFieldProps) => {
  return (
    <div className='flex flex-col gap-2 my-4'>
      <label className='text-subtext'>
        {title}{isRequired && <span className='text-red-500'>*</span>}
      </label>
      <input
        type={type}
        {...field}
        className='border border-subtext p-2 rounded focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary text-foreground transition duration-200 ease-in-out '
        autoFocus={autoFocus}
      />
      {fieldState.error && <span className='text-red-500'>{fieldState.error.message}</span>}

      {helperText &&
        <div className="text-subtext text-sm">
          <TextWithIcon icon={<RiQuestionLine />}>
            {helperText}
          </TextWithIcon>
        </div>
      }
    </div>
  )
};

export default InputField;