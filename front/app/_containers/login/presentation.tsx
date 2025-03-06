'use client'

import { loginApi } from '@/app/_containers/login/action'
import { CalloutContext } from '@/app/state'
import InputField from '@/components/layout/input/inputField'
import DPButton from '@/components/ui/button/button'
import DPLink from '@/components/ui/text/link'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { useRouter } from 'next/navigation'
import { useContext, useEffect, useState } from 'react'
import { Controller, useForm, ValidationRule } from 'react-hook-form'
import { RiLockLine, RiMailLine } from 'react-icons/ri'

type FormContent = {
	email: string
	password: string
}

/**
 * @package
 */
export default function UserLoginPresentation() {
	const { callout, setCallout } = useContext(CalloutContext)
	const router = useRouter()

	const { control, handleSubmit, reset, watch } = useForm<FormContent>({
		defaultValues: {
			email: '',
			password: '',
		},
	})

	const [isSubmit, setIsSubmit] = useState(false)

	useEffect(() => {
		if (isSubmit) {
			const loginFlow = async () => {
				const res = await loginApi(watch().email, watch().password)

				if (res.errors) {
					if (res.errors[0] === 'メールアドレスの認証が完了していません') {
						setCallout([...callout, { content: res.errors[0], type: 'info' }])
						router.push('/confirm-email')
					} else {
						setCallout([...callout, { content: res.errors[0], type: 'error' }])
					}
				} else {
					setCallout([...callout, { content: 'ログインに成功しました', type: 'success' }])
					router.push('/')
				}
			}
			loginFlow()
			reset()
			setIsSubmit(false)
		}
	}, [isSubmit])

	const passwordValidationRule: ValidationRule<RegExp> = {
		value: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[.+\-[\]*~_#:?]).{8,64}$/,
		message: 'パスワードは英数字をそれぞれ1文字以上含む8文字以上で入力してください',
	}

	return (
		<div className="flex h-screen flex-col items-center justify-center">
			<header className="mb-6">
				<TextWithIcon icon={<RiLockLine />} size="text-4xl">
					<h1 className="text-4xl font-bold tracking-widest text-foreground">ログイン</h1>
				</TextWithIcon>
			</header>

			<main className="flex w-1/3 flex-col gap-4">
				<form onSubmit={handleSubmit(() => setIsSubmit(true))}>
					<Controller
						name="email"
						control={control}
						rules={{ required: 'メールアドレスが未入力です' }}
						render={({ field, fieldState }) => (
							<InputField
								title="メールアドレス"
								type="email"
								field={field}
								fieldState={fieldState}
								isRequired
								autoFocus
								icon={<RiMailLine />}
							></InputField>
						)}
					></Controller>
					<Controller
						name="password"
						control={control}
						rules={{
							required: 'パスワードが未入力です',
							pattern: passwordValidationRule,
						}}
						render={({ field, fieldState }) => (
							<InputField
								title="パスワード"
								type="password"
								field={field}
								fieldState={fieldState}
								isRequired
								helperText="英数字をそれぞれ1文字以上含む8文字以上で入力してください"
								icon={<RiLockLine />}
							></InputField>
						)}
					></Controller>

					<div className="mt-6 flex justify-center">
						<DPButton colormode="primary" type="submit">
							<TextWithIcon icon={<RiLockLine />}>ログイン</TextWithIcon>
						</DPButton>
					</div>
				</form>

				<div className="flex justify-center">
					<span>
						アカウントをお持ちでない方は
						<DPLink href="/register">新規登録</DPLink>
						から！
					</span>
				</div>
			</main>
		</div>
	)
}
