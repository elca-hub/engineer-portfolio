import Pin, { PinCodeState } from '@/components/ui/input/pin'
import React, { useEffect, useState } from 'react'
import { Label, Text, TextFieldProps } from 'react-aria-components'
import { ControllerFieldState, ControllerRenderProps, FieldPath, FieldValue, FieldValues } from 'react-hook-form'
import TextWithIcon from '../../ui/text/textWithIcon'

interface PinCodeFieldProps extends TextFieldProps, React.RefAttributes<HTMLDivElement> {
	title: string
	field: ControllerRenderProps<FieldValue<FieldValues>, FieldPath<FieldValues>>
	fieldState: ControllerFieldState
	helperText?: string
	icon?: React.ReactNode
	length: number
}

const PinCodeField = ({ ...props }: PinCodeFieldProps) => {
	const titleNode = props.isDisabled ? (
		<>
			{props.title}
			<span className="text-red-500">*</span>
		</>
	) : (
		props.title
	)

	const labelNode = props.icon ? <TextWithIcon icon={props.icon}>{titleNode}</TextWithIcon> : titleNode

	const defaultValue = Array.from({ length: props.length })
		.map(() => '-')
		.join('')

	const [accessCode, setAccessCode] = useState<string>(defaultValue)
	const [pin, setPin] = useState<PinCodeState>({ code: '', id: 0 })
	const [nowFocusId, setNowFocusId] = useState<number>(0)

	useEffect(() => {
		const nowAccessCode = accessCode.split('')
		nowAccessCode[pin.id] = pin.code === '' ? '-' : pin.code
		setAccessCode(nowAccessCode.join(''))

		if (pin.code === '') {
			if (pin.id > 0) {
				setNowFocusId(pin.id - 1)
			}
		} else {
			if (pin.id < props.length - 1) {
				setNowFocusId(pin.id + 1)
			}
		}
	}, [accessCode, pin, props.length])

	// pinコードが変更されたら、親コンポーネントのfieldにセットする
	useEffect(() => {
		props.field.onChange(accessCode)
	}, [accessCode, props.field])

	return (
		<div className="my-4 flex flex-col gap-2">
			<Label className="text-gray-700">{labelNode}</Label>
			<div className="flex justify-between gap-2">
				{Array.from({ length: props.length }).map((_, index) => (
					<Pin key={index} id={index} code={setPin} nowFocusId={nowFocusId}></Pin>
				))}
			</div>
			<Text slot="description" className="text-sm text-subtext">
				{props.helperText}
			</Text>
			<p className="mb-0 text-red-500">{props.fieldState.error?.message}</p>
		</div>
	)
}

export default PinCodeField
