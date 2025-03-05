import React from 'react';
import { ControllerFieldState, ControllerRenderProps, FieldPath, FieldValue, FieldValues } from 'react-hook-form';
import TextWithIcon from '../../ui/text/textWithIcon';
import {Button, Calendar, CalendarCell, CalendarGrid, CalendarGridBody, CalendarGridHeader, CalendarHeaderCell, DateField, DateInput, DatePicker, DateSegment, Dialog, FieldError, Group, Heading, Input, Label, Popover, Text, TextField} from 'react-aria-components';
import { RiArrowDownFill, RiArrowDownSFill, RiArrowLeftSFill, RiArrowRightSFill, RiExpandUpDownFill } from 'react-icons/ri';

interface InputFieldProps {
  title: string;
  type: string;
  field: ControllerRenderProps<FieldValue<FieldValues>, FieldPath<FieldValues>>;
  fieldState: ControllerFieldState;
  isRequired?: boolean;
  helperText?: string;
  autoFocus?: boolean;
  icon?: React.ReactNode;
}

const InputField = ({title, type, field, fieldState, isRequired, helperText, autoFocus, icon}: InputFieldProps) => {
  const titleNode = isRequired ?
  (
    <>{title}<span className='text-red-500'>*</span></>
  ) : title;

  const labelNode = icon ? (
    <TextWithIcon icon={icon}>{titleNode}</TextWithIcon>
  ) : titleNode;

  return (
    <TextField
      isInvalid={fieldState.error ? true : false}
      {...field}
      isRequired={isRequired}
      type={type}
      autoFocus={autoFocus}
      className='flex flex-col gap-2 my-4'
    >
      <Label className="text-gray-700">{labelNode}</Label>
      <Input className='border border-subtext p-2 rounded focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary text-foreground transition duration-200 ease-in-out' />
      <Text slot="description" className='text-subtext text-sm'>{helperText}</Text>
      <FieldError className="text-red-500">{fieldState.error?.message}</FieldError>
    </TextField>
  )
};

export default InputField;
