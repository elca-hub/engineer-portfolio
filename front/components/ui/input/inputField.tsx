import React, { useEffect, useState } from 'react';
import { ControllerFieldState, ControllerRenderProps, UseFormRegisterReturn } from 'react-hook-form';

interface InputFieldProps {
  title: string;
  type: string;
  field: ControllerRenderProps<any, any>;
  fieldState: ControllerFieldState;
  isRequired?: boolean;
}

const InputField = ({title, type, field, fieldState, isRequired}: InputFieldProps) => {
  return (
    <div className='flex flex-col gap-2 my-2'>
      <label className='text-subtext'>
        {title}{isRequired && <span className='text-red-500'>*</span>}
      </label>
      <input
        type={type}
        {...field}
        className='border border-subtext p-2 rounded'
      />
      {fieldState.error && <span className='text-red-500'>{fieldState.error.message}</span>}
    </div>
  )
};

export default InputField;