import React from 'react'
import { FieldError, Input, Label, Text, TextField, TextFieldProps } from 'react-aria-components'
import { ControllerFieldState, ControllerRenderProps, FieldPath, FieldValue, FieldValues } from 'react-hook-form'
import TextWithIcon from '../../ui/text/textWithIcon'

interface InputFieldProps extends TextFieldProps, React.RefAttributes<HTMLDivElement> {
	title: string
	field: ControllerRenderProps<FieldValue<FieldValues>, FieldPath<FieldValues>>
	fieldState: ControllerFieldState
	isRequired?: boolean
	helperText?: string
	icon?: React.ReactNode
}

const InputField = ({ ...props }: InputFieldProps) => {
	const titleNode = props.isDisabled ? (
		<>
			{props.title}
			<span className="text-red-500">*</span>
		</>
	) : (
		props.title
	)

	const labelNode = props.icon ? <TextWithIcon icon={props.icon}>{titleNode}</TextWithIcon> : titleNode

	return (
		<TextField {...props.field} {...props} className="my-4 flex flex-col gap-2">
			<Label className="text-gray-700">{labelNode}</Label>
			<Input className="rounded border border-subtext p-2 text-foreground transition duration-200 ease-in-out focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary" />
			<Text slot="description" className="text-sm text-subtext">
				{props.helperText}
			</Text>
			<FieldError className="text-red-500">{props.fieldState.error?.message}</FieldError>
		</TextField>
	)
}

export default InputField
