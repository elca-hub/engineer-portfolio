/**
 * @package
 */
'use client'

import { CalloutContext } from '@/app/state'
import PinCodeField from '@/components/layout/input/pinCodeField'
import DPButton from '@/components/ui/button/button'
import DPLink from '@/components/ui/text/link'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { useRouter } from 'next/navigation'
import { useContext, useEffect, useState } from 'react'
import { Controller, useForm } from 'react-hook-form'
import { RiErrorWarningLine, RiMailSendLine } from 'react-icons/ri'
import { confirmEmailApi } from './action'

type FormContent = {
	accessCode: string
}

export default function ConfirmEmailPresenter() {
	const { callout, setCallout } = useContext(CalloutContext)
	const router = useRouter()

	const { control, handleSubmit, watch } = useForm<FormContent>({
		defaultValues: {
			accessCode: '------',
		},
	})

	const [isSubmit, setIsSubmit] = useState(false)

	useEffect(() => {
		if (isSubmit) {
			const confirmFlow = async () => {
				const res = await confirmEmailApi(watch().accessCode)

				if (res.errors) {
					setCallout([...callout, { content: res.errors[0], type: 'error' }])

					if (res.errors[0] === '1日を超過したため、再度新規登録を行なってください') router.push('/register')
				} else {
					setCallout([...callout, { content: '認証に成功しました', type: 'success' }])
					router.push('/user')
				}
			}
			confirmFlow()
			setIsSubmit(false)
		}
	}, [isSubmit])

	return (
		<div className="flex h-screen flex-col items-center justify-center">
			<header className="mb-6">
				<TextWithIcon icon={<RiMailSendLine />} size="text-4xl">
					<h1 className="text-4xl font-bold tracking-widest text-foreground">確認コードの入力</h1>
				</TextWithIcon>
			</header>

			<main className="flex w-1/3 flex-col gap-4">
				<div className="text-center text-red-500">
					<TextWithIcon icon={<RiErrorWarningLine />} size="text-2xl">
						<h2>
							1日以内に入力しない場合、
							<span className="font-bold">新規登録からやり直し</span>
							となります。
						</h2>
					</TextWithIcon>
				</div>

				<form onSubmit={handleSubmit(() => setIsSubmit(true))}>
					<Controller
						name="accessCode"
						control={control}
						rules={{
							validate: (value) => !value.includes('-') || '確認コードを入力してください',
						}}
						render={({ field, fieldState }) => (
							<PinCodeField
								title="確認コード"
								field={field}
								fieldState={fieldState}
								length={6}
								icon={<RiMailSendLine />}
								helperText="メールアドレス宛に送信された確認コードを入力してください"
							></PinCodeField>
						)}
					></Controller>

					<hr className="border-gray-300" />

					<div className="mt-6 flex justify-center">
						<DPLink href="/resend-email">確認コードを再送信する</DPLink>
					</div>

					<div className="mt-6 flex justify-center">
						<DPButton colormode="primary" type="submit">
							<TextWithIcon icon={<RiMailSendLine />}>送信</TextWithIcon>
						</DPButton>
					</div>
				</form>
			</main>
		</div>
	)
}
