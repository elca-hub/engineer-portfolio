import React from 'react';
import { ControllerFieldState, ControllerRenderProps, FieldPath, FieldValue, FieldValues } from 'react-hook-form';
import TextWithIcon from '../../ui/text/textWithIcon';
import {Button, Calendar, CalendarCell, CalendarGrid, CalendarGridBody, CalendarGridHeader, CalendarHeaderCell, DateField, DateInput, DatePicker, DateSegment, Dialog, FieldError, Group, Heading, Input, Label, Popover, Text, TextField} from 'react-aria-components';
import { RiArrowDownFill, RiArrowDownSFill, RiArrowLeftSFill, RiArrowRightSFill, RiExpandUpDownFill } from 'react-icons/ri';

interface DatepickerFieldProps {
  title: string;
  field: ControllerRenderProps<FieldValue<FieldValues>, FieldPath<FieldValues>>;
  fieldState: ControllerFieldState;
  isRequired?: boolean;
  helperText?: string;
  autoFocus?: boolean;
  icon?: React.ReactNode;
}

const DatePickerField = ({title, field, fieldState, isRequired, helperText, autoFocus, icon}: DatepickerFieldProps) => {
  const titleNode = isRequired ?
  (
    <>{title}<span className='text-red-500'>*</span></>
  ) : title;

  const labelNode = icon ? (
    <TextWithIcon icon={icon}>{titleNode}</TextWithIcon>
  ) : titleNode;
  return (
    <DatePicker autoFocus={autoFocus} className="group flex flex-col gap-1" {...field} isRequired={isRequired} isInvalid={fieldState.error ? true : false}>
      <Label className="text-gray-700">{labelNode}</Label>
      <Group className="flex rounded border border-subtext bg-white transition pl-3 text-foreground focus-visible:border-primary focus:ring-1 focus-visible:ring-primary">
        <DateInput className="flex flex-1 py-2">
          {(segment) => <DateSegment className="px-0.5 tabular-nums outline-none rounded-sm focus:bg-sky-100 focus-text-white caret-transparent placeholder-shown:text-gray-500" segment={segment} />}
        </DateInput>
        <Button className="outline-none px-3 flex items-center text-gray-700 transition border-0 border-solid border-l border-l-gray-300 bg-transparent rounded-r pressed:bg-sky-100 focus-visible:ring-2 ring-primary">
          <RiExpandUpDownFill className="text-2xl" />
        </Button>
      </Group>
      <Text slot="description" className='text-subtext text-sm'>{helperText}</Text>
      <FieldError className="text-red-500">{fieldState.error?.message}</FieldError>
      <Popover>
        <Dialog className='p-6 text-gray-600'>
          <Calendar className="bg-white p-2 rounded-lg shadow-md border border-gray-200">
            <header className='flex justify-between items-center gap-1 pb-4 px-1 w-full'>
              <Button slot="previous"><RiArrowLeftSFill className="text-xl" /></Button>
              <Heading className="flex-1 font-semibold text-2xl ml-2" />
              <Button slot="next"><RiArrowRightSFill className="text-xl" /></Button>
            </header>
            <CalendarGrid className="border-spacing-1 border-separate">
              <CalendarGridHeader>
                {(day) => (
                  <CalendarHeaderCell className="text-xs text-gray-500 font-semibold">{day}</CalendarHeaderCell>
                )}
              </CalendarGridHeader>
              <CalendarGridBody>
                {(date) => (
                  <CalendarCell
                    date={date}
                    className="w-9 h-9 outline-none cursor-default rounded-full flex items-center justify-center outside-month:text-gray-300 hover:bg-gray-100 pressed:bg-gray-200 selected:bg-primary selected:text-white focus-visible:ring ring-primary ring-offset-2"
                  />
                )}
              </CalendarGridBody>
            </CalendarGrid>
          </Calendar>
        </Dialog>
      </Popover>
    </DatePicker>
  )
};

export default DatePickerField;
