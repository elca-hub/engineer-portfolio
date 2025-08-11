'use client'

import { registerApi } from '@/app/_containers/register/action'
import { CalloutContext } from '@/app/state'
import DatePickerField from '@/components/layout/input/datePickerField'
import InputField from '@/components/layout/input/inputField'
import DPModal from '@/components/layout/modal'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { loginFlow } from '@/lib/access'
import { CalendarDate, getLocalTimeZone, today } from '@internationalized/date'
import { useRouter } from 'next/navigation'
import { useContext, useEffect, useState } from 'react'
import { DialogTrigger } from 'react-aria-components'
import { Controller, useForm } from 'react-hook-form'
import { RiAlarmWarningLine, RiCake2Line, RiCheckLine, RiIdCardLine, RiUserAddLine, RiUserLine } from 'react-icons/ri'

export type RegisterFormContent = {
	name: string
	userId: string
	birthday: CalendarDate
}

/**
 * @package
 */
export default function RegisterPresentation() {
	const { callout, setCallout } = useContext(CalloutContext)

	const router = useRouter()

	const { control, handleSubmit, watch, formState } = useForm<RegisterFormContent>({
		defaultValues: {
			name: '',
			userId: '',
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
					userId: watch().userId,
				})

				if (res.errors) {
					for (const error of res.errors) {
						setCallout([...callout, { type: 'error', content: error }])
					}

					return
				}

				const login = await loginFlow(res.data?.email || '')

				if (login.errors) {
					for (const error of login.errors) {
						setCallout([...callout, { type: 'error', content: error }])
					}

					return
				}

				setCallout([...callout, { type: 'success', content: '登録が完了しました' }])
				router.replace(`/${watch().userId}/profile`)
				return
			}

			registerFlow()

			setIsSubmit(false)
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
								icon={<RiUserLine />}
								autoComplete="off"
								autoFocus
								popoverContent="本名を入力する必要はありません。自分の個性的な名前をつけましょう！"
							></InputField>
						)}
					></Controller>

					<Controller
						name="userId"
						control={control}
						rules={{
							required: 'ユーザIDを入力してください',
							max: {
								value: 50,
								message: 'IDは50文字以内で入力してください',
							},
						}}
						render={({ field, fieldState }) => (
							<InputField
								title="ユーザID"
								type="text"
								field={field}
								fieldState={fieldState}
								isRequired
								helperText="IDは50文字以下で入力してください。特殊記号は使用できません。"
								icon={<RiIdCardLine />}
								autoComplete="off"
								popoverContent="プロフィールの共有、閲覧などに使用されます。"
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

					<DialogTrigger>
						<div className="flex justify-center mt-6">
							<DPButton colormode="primary" isDisabled={!formState.isValid}>
								<TextWithIcon icon={<RiUserAddLine />}>新規登録</TextWithIcon>
							</DPButton>
						</div>
						<DPModal header={{ title: '注意', icon: <RiAlarmWarningLine /> }}>
							<p>以下のIDは今後変更することができません。</p>
							<p className="text-xl text-yellow-600">@{watch().userId}</p>
							<p>このIDでよろしいですか？</p>

							<div className="mt-6 flex justify-center">
								<DPButton colormode="secondary" type="submit" onPress={() => setIsSubmit(true)}>
									<TextWithIcon icon={<RiCheckLine />}>登録</TextWithIcon>
								</DPButton>
							</div>
						</DPModal>
					</DialogTrigger>
				</form>
			</main>
		</div>
	)
}
