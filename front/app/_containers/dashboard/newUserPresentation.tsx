'use client'

import { registerApi } from '@/app/_containers/dashboard/action'
import { CalloutContext } from '@/app/state'
import DatePickerField from '@/components/layout/input/datePickerField'
import InputField from '@/components/layout/input/inputField'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { CalendarDate, getLocalTimeZone, today } from '@internationalized/date'
import { useRouter } from 'next/navigation'
import { useContext, useEffect, useState } from 'react'
import { Controller, useForm } from 'react-hook-form'
import { RiCake2Line, RiIdCardLine, RiUserAddLine } from 'react-icons/ri'

export type RegisterFormContent = {
	name: string
	birthday: CalendarDate
}

/**
 * @package
 */
export default function NewUserPresentation() {
	const { callout, setCallout } = useContext(CalloutContext)

	const router = useRouter()

	const { control, handleSubmit, watch } = useForm<RegisterFormContent>({
		defaultValues: {
			name: '',
			birthday: today('Asia/Tokyo'),
		},
	})

	const [isSubmit, setIsSubmit] = useState(false)

	useEffect(() => {
		if (isSubmit) {
			const registerFlow = async () => {
				const res = await registerApi({
					name: watch().name,
					birthday: watch().birthday.toString(),
				})

				if (res.errors) {
					for (const error of res.errors) {
						setCallout([...callout, { type: 'error', content: error }])
					}
				} else {
					setCallout([...callout, { type: 'success', content: 'ユーザ登録が完了しました' }])
					router.push('/dashboard')
				}
			}

			registerFlow()
		}
	}, [isSubmit])

	return (
		<div className="flex h-screen flex-col items-center justify-center">
			<header className="mb-6">
				<TextWithIcon icon={<RiUserAddLine />} size="text-4xl">
					<h1 className="text-4xl font-bold tracking-widest text-foreground">新規登録</h1>
				</TextWithIcon>
			</header>

			<main className="flex w-1/3 flex-col gap-4">
				<form onSubmit={handleSubmit(() => setIsSubmit(true))}>
					<Controller
						name="name"
						control={control}
						rules={{
							required: 'ユーザ名を入力してください',
							max: {
								value: 50,
								message: 'ユーザ名は50文字以内で入力してください',
							},
						}}
						render={({ field, fieldState }) => (
							<InputField
								title="ユーザ名"
								type="text"
								field={field}
								fieldState={fieldState}
								isRequired
								helperText="ユーザ名は50文字以下で入力してください。特殊記号は使用できません。"
								icon={<RiIdCardLine />}
								autoComplete="off"
								autoFocus
							></InputField>
						)}
					></Controller>

					<Controller
						name="birthday"
						control={control}
						rules={{
							required: '生年月日が未入力です',
							validate: (value) => value.compare(today(getLocalTimeZone())) < 0 || '未来の日付は指定できません',
						}}
						render={({ field, fieldState }) => (
							<DatePickerField title="生年月日" field={field} fieldState={fieldState} isRequired icon={<RiCake2Line />}></DatePickerField>
						)}
					></Controller>

					<div className="mt-6 flex justify-center">
						<DPButton colormode="primary" type="submit">
							<TextWithIcon icon={<RiUserAddLine />}>新規登録</TextWithIcon>
						</DPButton>
					</div>
				</form>
			</main>
		</div>
	)
}
