import { getLocalTimeZone, today } from '@internationalized/date';
import React from 'react';
import { Button, Calendar, CalendarCell, CalendarGrid, CalendarGridBody, CalendarGridHeader, CalendarHeaderCell, DateInput, DatePicker, DateSegment, Dialog, FieldError, Group, Heading, I18nProvider, Label, Popover, Text } from 'react-aria-components';
import { ControllerFieldState, ControllerRenderProps, FieldPath, FieldValue, FieldValues } from 'react-hook-form';
import { RiArrowLeftSFill, RiArrowRightSFill, RiCalendar2Line } from 'react-icons/ri';
import TextWithIcon from '../../ui/text/textWithIcon';

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
    <I18nProvider locale='ja'>
      <DatePicker maxValue={today(getLocalTimeZone())} autoFocus={autoFocus} className="group flex flex-col gap-1" {...field} isRequired={isRequired} isInvalid={fieldState.error ? true : false}>
        <Label className="text-gray-700">{labelNode}</Label>
        <Group className="flex rounded border border-subtext bg-white pl-3 text-foreground transition focus:ring-1 focus-visible:border-primary focus-visible:ring-primary">
          <DateInput className="flex flex-1 py-2">
            {(segment) => <DateSegment className="focus-text-white rounded-sm px-0.5 tabular-nums caret-transparent outline-none placeholder-shown:text-gray-500 focus:bg-sky-100" segment={segment} />}
          </DateInput>
          <Button className="flex items-center rounded-r border-0 border-l border-solid border-l-gray-300 bg-transparent px-3 text-gray-700 outline-none ring-primary transition focus-visible:ring-2 pressed:bg-sky-100">
            <RiCalendar2Line className="text-2xl" />
          </Button>
        </Group>
        <Text slot="description" className='text-sm text-subtext'>{helperText}</Text>
        <FieldError className="text-red-500">{fieldState.error?.message}</FieldError>
        <Popover>
          <Dialog className='p-6 text-gray-600'>
            <Calendar className="rounded-lg border border-gray-200 bg-white p-2 shadow-md">
              <header className='flex w-full items-center justify-between gap-1 px-1 pb-4'>
                <Button slot="previous"><RiArrowLeftSFill className="text-xl" /></Button>
                <Heading className="ml-2 flex-1 text-2xl font-semibold" />
                <Button slot="next"><RiArrowRightSFill className="text-xl" /></Button>
              </header>
              <CalendarGrid className="border-separate border-spacing-1">
                <CalendarGridHeader>
                  {(day) => (
                    <CalendarHeaderCell className="text-xs font-semibold text-gray-500">{day}</CalendarHeaderCell>
                  )}
                </CalendarGridHeader>
                <CalendarGridBody>
                  {(date) => (
                    <CalendarCell
                      date={date}
                      className="flex size-9 cursor-default items-center justify-center rounded-full outline-none ring-primary ring-offset-2 outside-month:text-gray-300 hover:bg-gray-100 focus-visible:ring pressed:bg-gray-200 selected:bg-primary selected:text-white disabled:text-gray-300"
                    />
                  )}
                </CalendarGridBody>
              </CalendarGrid>
            </Calendar>
          </Dialog>
        </Popover>
      </DatePicker>
    </I18nProvider>
    
  )
};

export default DatePickerField;
