import React from 'react';
import { ControllerFieldState, ControllerRenderProps, FieldPath, FieldValue, FieldValues } from 'react-hook-form';
import TextWithIcon from '../../ui/text/textWithIcon';
import { FieldError, Input, Label, Text, TextField, TextFieldProps} from 'react-aria-components';

interface InputFieldProps extends TextFieldProps, React.RefAttributes<HTMLDivElement> {
  title: string;
  field: ControllerRenderProps<FieldValue<FieldValues>, FieldPath<FieldValues>>;
  fieldState: ControllerFieldState;
  isRequired?: boolean;
  helperText?: string;
  icon?: React.ReactNode;
}

const InputField = ({...props}: InputFieldProps) => {
  const titleNode = props.isDisabled ?
  (
    <>{props.title}<span className='text-red-500'>*</span></>
  ) : props.title;

  const labelNode = props.icon ? (
    <TextWithIcon icon={props.icon}>{titleNode}</TextWithIcon>
  ) : titleNode;

  return (
    <TextField
      {...props.field}
      {...props}
      className='flex flex-col gap-2 my-4'
    >
      <Label className="text-gray-700">{labelNode}</Label>
      <Input className='border border-subtext p-2 rounded focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary text-foreground transition duration-200 ease-in-out' />
      <Text slot="description" className='text-subtext text-sm'>{props.helperText}</Text>
      <FieldError className="text-red-500">{props.fieldState.error?.message}</FieldError>
    </TextField>
  )
};

export default InputField;
