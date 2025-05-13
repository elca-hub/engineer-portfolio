import { UserType } from '@/action/type/user'
import fetchBio from '@/action/usecase/user/fetchBio'
import uploadBio from '@/action/usecase/user/uploadBio'
import uploadBioImage from '@/action/usecase/user/uploadBioImage'
import { CalloutContext } from '@/app/state'
import InputField from '@/components/layout/input/inputField'
import CustomMarkdown from '@/components/layout/markdown/CustomMarkdown'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getSessionToken } from '@/lib/access'
import { useRouter } from 'next/navigation'
import { useContext, useEffect, useRef, useState } from 'react'
import { Button, DropZone, FileTrigger, Tab, TabList, TabPanel, Tabs, TextArea } from 'react-aria-components'
import { Controller, useForm } from 'react-hook-form'
import { RiEyeLine, RiImageAddLine, RiPencilLine, RiUserLine } from 'react-icons/ri'
type Props = {
	user: UserType
}

export type UserUpdateFormType = {
	bio: string
}

export default function BioSettingTabPresentation({ user }: Props) {
	const router = useRouter()

	const { callout, setCallout } = useContext(CalloutContext)
	const [isLoading, setIsLoading] = useState(true)

	// デバウンス処理用
	const [debouncedValue, setDebouncedValue] = useState('')

	/* 画像アップロード関連 */
	const [imageFile, setImageFile] = useState<File | null>(null)
	const [isUploadImage, setIsUploadImage] = useState(false)

	const { control, handleSubmit, setValue, watch } = useForm<UserUpdateFormType>({
		defaultValues: {
			bio: '',
		},
	})
	const [isSubmit, setIsSubmit] = useState(false)

	const textareaRef = useRef<HTMLTextAreaElement>(null)

	useEffect(() => {
		const timer = setTimeout(() => {
			setDebouncedValue(watch('bio'))
		}, 1000)
		return () => clearTimeout(timer)
	}, [watch('bio')])

	useEffect(() => {
		const fb = async () => {
			setIsLoading(true)
			const bio = await fetchBio(user)
			setValue('bio', bio === null ? '' : bio)
			setIsLoading(false)
		}

		fb()
	}, [])

	useEffect(() => {
		if (isSubmit || debouncedValue !== '') {
			const bio = watch('bio')
			const updateFlow = async (isSubmit: boolean) => {
				const token = await getSessionToken()
				if (!token) {
					setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
					return
				}

				// const beforeVal = await fetchBio(user)
				// if (beforeVal === bio) {
				// 	if (isSubmit) {
				// 		router.refresh()
				// 		setCallout([...callout, { content: '変更しました', type: 'info' }])
				// 	}
				// 	return
				// }

				const res = await uploadBio(token, bio, isSubmit)
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
			updateFlow(isSubmit)
			setIsSubmit(false)
			setDebouncedValue('')
		}
	}, [isSubmit, debouncedValue])

	/* 画像アップロード関連 */
	useEffect(() => {
		if (imageFile) {
			setIsUploadImage(true)
			const uploadFlow = async () => {
				const token = await getSessionToken()
				if (!token) {
					setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
					return
				}

				const res = await uploadBioImage(token, imageFile)

				if (res.errors) {
					for (const error of res.errors) {
						setCallout([...callout, { content: error, type: 'error' }])
					}
				}

				if (res.data) {
					const textarea = textareaRef.current
					if (textarea) {
						const start = textarea.selectionStart
						const end = textarea.selectionEnd
						const currentValue = watch('bio')
						const newValue = currentValue.substring(0, start) + `\n![image](${res.data.image_url})\n` + currentValue.substring(end)
						setValue('bio', newValue)
					} else {
						setValue('bio', `${watch('bio')}\n\n![image](${res.data.image_url})`)
					}
				}
			}
			uploadFlow()
			setImageFile(null)
			setIsUploadImage(false)
		}
	}, [imageFile])
	const tabItemClassName = `cursor-pointer text-lg text-subtext outline-border outline-primary data-[selected]:border-b data-[selected]:border-primary data-[selected]:font-medium data-[selected]:text-foreground`

	return (
		<>
			<Tabs>
				<TabList aria-label="自己紹介" className="flex justify-start gap-x-4 mb-4">
					<Tab id="edit" className={tabItemClassName}>
						<TextWithIcon icon={<RiPencilLine />}>編集</TextWithIcon>
					</Tab>
					<Tab className={`${tabItemClassName} font-medium text-foreground`} isDisabled>
						<DropZone
							onDrop={(e) => {
								const targetFile = e.items[0]
								if (targetFile.kind !== 'file') return
								targetFile.getFile().then((f) => {
									setImageFile(f)
								})
							}}
							className={`z-10`}
						>
							<FileTrigger
								onSelect={(e) => {
									if (!e) return
									const file = e.item(0)
									if (file === null) return
									setImageFile(file)
								}}
								acceptedFileTypes={['image/*']}
							>
								<Button className="z-10">
									<TextWithIcon icon={<RiImageAddLine />}>画像</TextWithIcon>
								</Button>
							</FileTrigger>
						</DropZone>
					</Tab>
					<Tab id="preview" className={tabItemClassName}>
						<TextWithIcon icon={<RiEyeLine />}>プレビュー</TextWithIcon>
					</Tab>
				</TabList>
				<TabPanel id="edit">
					<form onSubmit={handleSubmit(() => setIsSubmit(true))}>
						{isLoading ? (
							<div className="animate-pulse">
								<div className="h-4 bg-gray-200 rounded w-1/4 mb-2"></div>
								<div className="h-32 bg-gray-200 rounded w-full"></div>
							</div>
						) : (
							<Controller
								name="bio"
								control={control}
								rules={{
									max: {
										value: 5000,
										message: '自己紹介文は5000文字以内で入力してください',
									},
								}}
								render={({ field, fieldState }) => (
									<InputField
										title="自己紹介文"
										field={field}
										fieldState={fieldState}
										isMultiline
										helperText="markdownによる入力が可能です。1000文字以内で入力してください。"
										icon={<RiUserLine />}
										autoComplete="on"
										popoverContent="入力していない状態でも内容は自動的に保存されます。"
										rows={10}
										isDisabled={isUploadImage}
										isLoading={isLoading}
										customInput={
											<TextArea
												ref={textareaRef}
												rows={10}
												className="w-full rounded border border-subtext text-foreground p-2 transition duration-200 ease-in-out focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary data-[disabled]:text-subtext"
												placeholder="自分の魅力を伝えられるように、とびっきりの内容を書きましょう！！"
												onKeyDown={(e) => {
													if (e.key === 'Tab') {
														e.preventDefault()
														const textarea = e.currentTarget
														const start = textarea.selectionStart
														const end = textarea.selectionEnd
														const value = textarea.value
														textarea.value = value.substring(0, start) + '\t' + value.substring(end)
														textarea.selectionStart = textarea.selectionEnd = start + 1
														// React Hook Formの値も更新
														if (typeof field?.onChange === 'function') {
															field.onChange(textarea.value)
														}
													}
												}}
											/>
										}
									></InputField>
								)}
							></Controller>
						)}
						<div className="flex justify-center mt-4">
							<DPButton colormode="primary" type="submit">
								<TextWithIcon icon={<RiPencilLine />}>変更する</TextWithIcon>
							</DPButton>
						</div>
					</form>
				</TabPanel>

				<TabPanel id="preview">
					<div className="bg-gray-100 p-4 rounded-lg">
						{isLoading ? (
							<div className="animate-pulse">
								<div className="h-4 bg-gray-200 rounded w-3/4 mb-2"></div>
								<div className="h-4 bg-gray-200 rounded w-1/2 mb-2"></div>
								<div className="h-4 bg-gray-200 rounded w-2/3"></div>
							</div>
						) : (
							<CustomMarkdown>{watch('bio')}</CustomMarkdown>
						)}
					</div>
				</TabPanel>
			</Tabs>
		</>
	)
}
