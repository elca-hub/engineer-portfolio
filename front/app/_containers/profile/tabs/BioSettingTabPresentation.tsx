import { UserType } from '@/action/type/user'
import fetchBio from '@/action/usecase/user/fetchBio'
import uploadBio from '@/action/usecase/user/uploadBio'
import { CalloutContext } from '@/app/state'
import InputField from '@/components/layout/input/inputField'
import CustomMarkdown from '@/components/layout/markdown/CustomMarkdown'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getSessionToken } from '@/lib/access'
import { useRouter } from 'next/navigation'
import { useContext, useEffect, useState } from 'react'
import { Tab, TabList, TabPanel, Tabs } from 'react-aria-components'
import { Controller, useForm } from 'react-hook-form'
import { RiEyeLine, RiPencilLine, RiUserLine } from 'react-icons/ri'
type Props = {
	user: UserType
}

export type UserUpdateFormType = {
	bio: string
}

export default function BioSettingTabPresentation({ user }: Props) {
	const router = useRouter()

	const { callout, setCallout } = useContext(CalloutContext)

	// デバウンス処理用
	const [debouncedValue, setDebouncedValue] = useState('')

	const { control, handleSubmit, setValue, watch } = useForm<UserUpdateFormType>({
		defaultValues: {
			bio: '',
		},
	})
	const [isSubmit, setIsSubmit] = useState(false)

	useEffect(() => {
		const timer = setTimeout(() => {
			setDebouncedValue(watch('bio'))
		}, 1000)
		return () => clearTimeout(timer)
	}, [watch('bio')])

	useEffect(() => {
		const fb = async () => {
			const bio = await fetchBio(user)
			setValue('bio', bio === null ? '' : bio)
		}

		fb()
	}, [])

	useEffect(() => {
		if (isSubmit || debouncedValue !== '') {
			const bio = watch('bio')
			const updateFlow = async () => {
				const token = await getSessionToken()
				if (!token) {
					setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
					return
				}

				const beforeVal = await fetchBio(user)
				if (beforeVal === bio) return

				const res = await uploadBio(token, bio)

				if (res.errors) {
					for (const error of res.errors) {
						setCallout([...callout, { content: error, type: 'error' }])
					}
					return
				}

				if (isSubmit) {
					router.refresh()
					setCallout([...callout, { content: '変更しました', type: 'info' }])
				}
			}
			updateFlow()
			setIsSubmit(false)
		}
	}, [isSubmit, debouncedValue])

	const tabItemClassName = `cursor-pointer text-lg text-subtext outline-border outline-primary data-[selected]:border-b data-[selected]:border-primary data-[selected]:font-medium data-[selected]:text-foreground`

	return (
		<>
			<Tabs>
				<TabList aria-label="自己紹介" className="flex justify-start gap-x-4 mb-4">
					<Tab id="edit" className={tabItemClassName}>
						<TextWithIcon icon={<RiPencilLine />}>編集</TextWithIcon>
					</Tab>
					<Tab id="preview" className={tabItemClassName}>
						<TextWithIcon icon={<RiEyeLine />}>プレビュー</TextWithIcon>
					</Tab>
				</TabList>
				<TabPanel id="edit">
					<form onSubmit={handleSubmit(() => setIsSubmit(true))}>
						<Controller
							name="bio"
							control={control}
							rules={{
								max: {
									value: 1000,
									message: '自己紹介文は1000文字以内で入力してください',
								},
							}}
							render={({ field, fieldState }) => (
								<InputField
									title="自己紹介文"
									field={field}
									fieldState={fieldState}
									isMultiline
									helperText="markdownによる入力が可能です。"
									icon={<RiUserLine />}
									autoComplete="on"
									popoverContent="自己紹介文は1000文字以内で入力してください。"
									rows={10}
								></InputField>
							)}
						></Controller>
						<div className="flex justify-center mt-4">
							<DPButton colormode="primary" type="submit">
								<TextWithIcon icon={<RiPencilLine />}>変更する</TextWithIcon>
							</DPButton>
						</div>
					</form>
				</TabPanel>
				<TabPanel id="preview">
					<div className="bg-gray-100 p-4 rounded-lg">
						<CustomMarkdown>{watch('bio')}</CustomMarkdown>
					</div>
				</TabPanel>
			</Tabs>
		</>
	)
}
