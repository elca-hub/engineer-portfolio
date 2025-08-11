import React from 'react'
import { Button, Dialog, DialogTrigger, FieldError, Input, Label, Popover, Text, TextField, TextFieldProps } from 'react-aria-components'
import { ControllerFieldState, ControllerRenderProps, FieldPath, FieldValue, FieldValues } from 'react-hook-form'
import { RiInfoI } from 'react-icons/ri'
import TextWithIcon from '../../ui/text/textWithIcon'

interface InputFieldProps extends TextFieldProps, React.RefAttributes<HTMLDivElement> {
	title?: string
	field: ControllerRenderProps<FieldValue<FieldValues>, FieldPath<FieldValues>>
	fieldState: ControllerFieldState
	isRequired?: boolean
	helperText?: string
	icon?: React.ReactNode
	popoverContent?: string
	isMultiline?: boolean
	rows?: number
	customInput?: React.ReactNode
	isLoading?: boolean
	placeholder?: string
	inputSize?: 'big'
}

const InputField = ({ ...props }: InputFieldProps) => {
	const titleNode = props.isRequired ? (
		<span>
			{props.title}
			<span className="text-red-500">*</span>
		</span>
	) : (
		props.title
	)
	const labelNode = props.icon ? <TextWithIcon icon={props.icon}>{titleNode}</TextWithIcon> : titleNode

	const labelPopOver = props.popoverContent ? (
		<>
			<DialogTrigger>
				<div className="flex flex-row items-center gap-4">
					{labelNode}
					<Button className="outline-primary rounded-full ring-1 ring-foreground text-foreground">
						<RiInfoI></RiInfoI>
					</Button>
				</div>
				<Popover
					placement="top"
					className="bg-white ring-1 ring-primary p-2 rounded entering:animate-in entering:fade-in entering:placement-bottom:slide-in-from-top-1 entering:placement-top:slide-in-from-bottom-1 exiting:animate-out exiting:fade-out exiting:placement-bottom:slide-out-to-top-1 exiting:placement-top:slide-out-to-bottom-1 fill-mode-forwards origin-top-left"
				>
					<Dialog className="text-foreground outline-none">
						<p>{props.popoverContent}</p>
					</Dialog>
				</Popover>
			</DialogTrigger>
		</>
	) : (
		labelNode
	)

	let normalInputClassName = 'rounded border border-subtext p-2 text-foreground transition duration-200 ease-in-out focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary'

	switch (props.inputSize) {
		case 'big':
			normalInputClassName += ' text-3xl font-medium'
			break
	}

	return props.isLoading ? (
		<div className="animate-pulse">
			<div className="h-4 bg-gray-200 rounded w-3/4 mb-2"></div>
			<div className="h-4 bg-gray-200 rounded w-1/2 mb-2"></div>
			<div className="h-4 bg-gray-200 rounded w-2/3"></div>
		</div>
	) : (
		<TextField {...props.field} {...props} className={`my-4 flex flex-col gap-2 ${props.className}`}>
			<Label className="text-gray-700">{labelPopOver}</Label>
			{props.customInput ? (
				props.customInput
			) : (
				<Input
					placeholder={props.placeholder}
					className={normalInputClassName}
				/>
			)}
			<Text slot="description" className="text-sm text-subtext">
				{props.helperText}
			</Text>
			<FieldError className="text-red-500">{props.fieldState.error?.message}</FieldError>
		</TextField>
	)
}

export default InputField
