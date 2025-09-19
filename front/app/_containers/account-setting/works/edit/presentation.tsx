'use client'

import { WorkType } from '@/action/type/work'
import { UserType } from '@/action/type/user'
import { CalloutContext } from '@/app/state'
import InputField from '@/components/layout/input/inputField'
import CustomMarkdown from '@/components/layout/markdown/CustomMarkdown'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getSessionToken } from '@/lib/access'
import { useContext, useEffect, useRef, useState, useCallback } from 'react'
import { DropZone, FileTrigger, TextArea, Tabs, TabList, Tab, TabPanel } from 'react-aria-components'
import { Controller, useFieldArray, useForm } from 'react-hook-form'
import { RiGithubLine, RiArticleLine, RiEyeLine, RiCheckboxCircleLine, RiLink, RiImageAddLine, RiMoreLine, RiUpload2Line, RiSave2Line } from 'react-icons/ri'
import updateWork from '@/action/usecase/works/updateWork'
import TagField from '@/components/layout/input/tagField'
import { WorkUrlType } from '@/action/type/workUrl'
import UrlListField from '@/components/layout/input/urlListField'
import uploadWorksImage from '@/action/usecase/works/uploadWorksImage'
import HeadContent from '@/components/layout/headContent'
import { HeaderButtonContext } from '@/app/account/setting/state'
import Link from 'next/link'

type EditWorkPresentationProps = {
	work: WorkType
	user: UserType
}

export type WorkFormType = {
	title: string
	content: string
	githubRepositoryUrl: string
	tags: { name: string }[]
	externalServiceUrls: WorkUrlType[]
}

const MAX_TAGS = 5
const MAX_TAG_NAME_LENGTH = 50

export default function EditWorkPresentation({ work, user }: EditWorkPresentationProps) {
	const textareaRef = useRef<HTMLTextAreaElement>(null)
	const { callout, setCallout } = useContext(CalloutContext)

	const [isSubmitting, setIsSubmitting] = useState(false)

	/* 自動保存関連 */
	const [lastUpdatedTime, setLastUpdatedTime] = useState<Date | null>(null)
	const [autoSaveTimeout, setAutoSaveTimeout] = useState<NodeJS.Timeout | null>(null)
	const autoSaveTimeoutRef = useRef<NodeJS.Timeout | null>(null)

	/* 画像アップロード関連 */
	const [imageFile, setImageFile] = useState<File | null>(null)
	const [isUploadImage, setIsUploadImage] = useState(false)

	/* プレビューの高さ調整関連 */
	const [contentHeight, setContentHeight] = useState<number>(400)

	/* ヘッダーボタン関連 */
	const contextObj = useContext(HeaderButtonContext)
	const setHeaderButton = contextObj === undefined ? null : contextObj.setHeaderButton

	const { control, handleSubmit, watch, formState, reset, trigger, setValue } = useForm<WorkFormType>({
		defaultValues: {
			title: work.title || '',
			content: work.content || '',
			githubRepositoryUrl: work.github_repository_url || '',
			externalServiceUrls: work.external_service_urls || [],
			tags: work.tags || [{ name: '' }],
		},
	})

	const { fields, append, remove } = useFieldArray({
		control,
		name: 'tags'
	})

	const { fields: urlFields, append: appendUrl, remove: removeUrl } = useFieldArray({
		control,
		name: 'externalServiceUrls'
	})

	const contentValue = watch('content')

	const convertFormTypeToWorkType = (data: WorkFormType): WorkType => {
		return {
			id: work.id,
			title: data.title,
			content: data.content,
			tags: data.tags.map(tag => ({ name: tag.name })),
			github_repository_url: data.githubRepositoryUrl,
			external_service_urls: data.externalServiceUrls,
			is_draft: work.is_draft,
			thumbnail_image_url: work.thumbnail_image_url,
		}
	}

	// work全体を更新するAPI
	const updateWorkHandler = async (data: WorkType, isDeleteImage: boolean = false) => {
		const token = await getSessionToken()
		if (!token) {
			setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
			return
		}

		const res = await updateWork(token, user.user_id, data, isDeleteImage)
		if (res.errors) {
			setCallout([...callout, { content: '作品の更新に失敗しました', type: 'error' }])
			return
		}

		console.log(data)

		setLastUpdatedTime(new Date())
	}

	// フォームが変更されたら自動保存を予約
	useEffect(() => {
		if (formState.isDirty) {
			const isValid = async () => {
				const isValid = await trigger()
				if (!isValid) return

				scheduleAutoSave(convertFormTypeToWorkType(watch()))
			}
			isValid()

			reset(
				watch(),
				{
					keepValues: true,
				}
			)
		}
	}, [formState.isDirty])

	// 自動保存処理
	const scheduleAutoSave = useCallback((updatedWork: WorkType) => {
		// 既存のタイムアウトをクリア
		if (autoSaveTimeoutRef.current) {
			clearTimeout(autoSaveTimeoutRef.current)
		}

		// 5秒後に自動保存を実行
		const timeoutId = setTimeout(() => {
			updateWorkHandler(updatedWork, false)
		}, 5000)

		autoSaveTimeoutRef.current = timeoutId
		setAutoSaveTimeout(timeoutId)
	}, [])

	// 初期化時の設定
	useEffect(() => {
		setLastUpdatedTime(new Date())
	}, [work])

	// コンポーネントがアンマウントされる際にタイムアウトをクリーンアップ
	useEffect(() => {
		return () => {
			if (autoSaveTimeout) {
				clearTimeout(autoSaveTimeout)
			}
		}
	}, [autoSaveTimeout])

	useEffect(() => {
		scheduleAutoSave(convertFormTypeToWorkType(watch()))
	}, [watch('content')])

	const handleUpdateWork = async (data: WorkFormType, isPushSubmit: boolean = false) => {
		if (isSubmitting || !(formState.isDirty || isPushSubmit)) return

		setIsSubmitting(true)
		console.log(data)
		await updateWorkHandler(convertFormTypeToWorkType(data), isPushSubmit)
		setIsSubmitting(false)

		reset(data, {
			keepValues: true,
		})
	}

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

				const res = await uploadWorksImage(token, work.id, imageFile)

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
						const currentValue = watch('content')
						const newValue = currentValue.substring(0, start) + `\n![image](${res.data.image_url})\n` + currentValue.substring(end)
						setValue('content', newValue)
					} else {
						setValue('content', `${watch('content')}\n![image](${res.data.image_url})`)
					}
				}
			}
			uploadFlow()
			setImageFile(null)
			setIsUploadImage(false)
		}
	}, [imageFile])

	// 画面サイズに応じた高さ調整
	const calculateContentHeight = useCallback(() => {
		const contentContainer = document.getElementById('content-container')
		if (contentContainer) {
			const screenHeight = window.innerHeight
			const containerPosition = contentContainer.getBoundingClientRect()
			// 下部のマージンを考慮して少し余裕を持たせる
			const calculatedHeight = screenHeight - containerPosition.top - 140
			setContentHeight(Math.max(calculatedHeight, 300)) // 最小高さを300pxに設定
		}
	}, [])

	useEffect(() => {
		calculateContentHeight()
	}, [calculateContentHeight])

	// リサイズイベントのリスナー
	useEffect(() => {
		const handleResize = () => {
			calculateContentHeight()
		}

		window.addEventListener('resize', handleResize)
		return () => window.removeEventListener('resize', handleResize)
	}, [calculateContentHeight])

	// 日時フォーマット関数
	const formatDateTime = (date: Date) => {
		return new Intl.DateTimeFormat('ja-JP', {
			year: 'numeric',
			month: '2-digit',
			day: '2-digit',
			hour: '2-digit',
			minute: '2-digit',
			second: '2-digit'
		}).format(date)
	}

	/* ヘッダーボタン関連 */
	useEffect(() => {
		if (setHeaderButton) {
			setHeaderButton(
				<div className="flex gap-2">
					<DPButton colormode="mono" onPress={() => handleUpdateWork(watch(), false)}>
						<TextWithIcon icon={<RiSave2Line />}>下書き保存</TextWithIcon>
					</DPButton>
					<Link href={`/account/setting/works/${work.id}/publish`}>
						<DPButton colormode="primary">
							<TextWithIcon icon={<RiUpload2Line />}>公開範囲の設定へ</TextWithIcon>
						</DPButton>
					</Link>
				</div>
			)
		}
	}, [setHeaderButton])

	return (
		<>
			<HeadContent title={`「${work.title}」の編集`} des={work.content} />
			<div className="h-full flex flex-col">
				{/* 最終更新日時表示 */}
				{lastUpdatedTime && (
					<div className="flex justify-left gap-6 text-green-800">
						<TextWithIcon icon={<RiCheckboxCircleLine />}>最終更新日時</TextWithIcon>
						{formatDateTime(lastUpdatedTime)}
					</div>
				)}

				<form className="h-full" onSubmit={handleSubmit((data) => handleUpdateWork(data, true))}>
					<div className="h-full flex flex-col">
						<Controller
							name="title"
							control={control}
							rules={{
								maxLength: {
									value: 100,
									message: '作品名は100文字以内で入力してください',
								},
							}}
							render={({ field, fieldState }) => (
								<InputField
									type="text"
									field={field}
									fieldState={fieldState}
									helperText={`${field.value.length}/100`}
									placeholder="作品名を入力してください"
									autoComplete="on"
									inputSize="big"
								/>
							)}
						/>

						<Tabs className="h-full">
							<TabList className="flex border-b border-subtext mb-4">
								<Tab id="content" className="px-4 py-2 cursor-pointer data-[selected]:border-b-2 data-[selected]:border-primary data-[selected]:text-primary hover:text-primary transition-colors">
									<TextWithIcon icon={<RiArticleLine />}>本文</TextWithIcon>
								</Tab>
								<Tab id="option" className="px-4 py-2 cursor-pointer data-[selected]:border-b-2 data-[selected]:border-primary data-[selected]:text-primary hover:text-primary transition-colors">
									<TextWithIcon icon={<RiMoreLine />}>その他</TextWithIcon>
								</Tab>
							</TabList>

							<TabPanel className="focus:outline-none" id="content">
								<DropZone
									onDrop={(e) => {
										const targetFile = e.items[0]
										if (targetFile.kind !== 'file') return
										targetFile.getFile().then((f) => {
											setImageFile(f)
										})
									}}
									className={`z-10 mb-4`}
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
										<DPButton colormode="primary" buttonSize="small">
											<TextWithIcon icon={<RiImageAddLine />}>画像</TextWithIcon>
										</DPButton>
									</FileTrigger>
								</DropZone>
								<div id="content-container" className="grid grid-cols-1 lg:grid-cols-2 gap-6">
									<div>
										<Controller
											name="content"
											control={control}
											render={({ field, fieldState }) => (
												<InputField
													title="本文"
													field={field}
													fieldState={fieldState}
													isMultiline
													icon={<RiArticleLine />}
													isDisabled={isUploadImage}
													isLoading={isSubmitting}
													customInput={
														<DropZone
															onDrop={(e) => {
																const targetFile = e.items[0]
																if (targetFile.kind !== 'file') return
																targetFile.getFile().then((f) => {
																	setImageFile(f)
																})
															}}
															className="w-full"
														>
															<TextArea
																ref={textareaRef}
																className="w-full rounded border border-subtext text-foreground p-2 transition duration-200 ease-in-out focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary data-[disabled]:text-subtext resize-none"
																style={{ height: `${contentHeight}px` }}
																placeholder="作品の魅力を伝えられるように、とびっきりの内容を書きましょう！！"
																onKeyDown={(e) => {
																	if (e.key === 'Tab') {
																		e.preventDefault()
																		const textarea = e.currentTarget
																		const start = textarea.selectionStart
																		const end = textarea.selectionEnd
																		const value = textarea.value
																		textarea.value = value.substring(0, start) + '\t' + value.substring(end)
																		textarea.selectionStart = textarea.selectionEnd = start + 1
																		if (typeof field?.onChange === 'function') {
																			field.onChange(textarea.value)
																		}
																	}
																}}
															/>
														</DropZone>
													}
													className="my-0"
												/>
											)}
										/>
									</div>

									<div className="my-4">
										<div className="flex items-center gap-2 mb-2">
											<TextWithIcon icon={<RiEyeLine />}>プレビュー</TextWithIcon>
										</div>
										<div id="preview-content" className="border border-subtext rounded p-4 bg-background overflow-auto" style={{ height: `${contentHeight}px` }}>
											<CustomMarkdown>{contentValue || ''}</CustomMarkdown>
										</div>
									</div>
								</div>
							</TabPanel>

							<TabPanel className="focus:outline-none" id="option">
								<div className="space-y-4">
									<TagField
										helperText={`${fields.length}/${MAX_TAGS} タグ - Enterキーで追加、×ボタンで削除`}
										tags={fields}
										maxTags={MAX_TAGS}
										append={append}
										remove={remove}
										validation={
											(tags) => {
												if (!tags) return null
												if (tags.length > MAX_TAGS) return `タグは最大${MAX_TAGS}つまでです`
												if (tags.some(tag => tag.name.length > MAX_TAG_NAME_LENGTH)) return `タグは${MAX_TAG_NAME_LENGTH}文字以内で入力してください`
												return null
											}
										}
									/>

									<Controller
										name="githubRepositoryUrl"
										control={control}
										rules={{
											pattern: {
												value: /^(https:\/\/github\.com\/.+)?$/,
												message: 'GitHubのリポジトリURLを入力してください'
											},
										}}
										render={({ field, fieldState }) => (
											<InputField
												title="Githubリポジトリ（オプション）"
												type="url"
												field={field}
												fieldState={fieldState}
												icon={<RiGithubLine />}
												autoComplete="on"
												placeholder="https://github.com/username/repository"
											/>
										)}
									/>

									<UrlListField
										title="URL（オプション）"
										urls={urlFields}
										append={appendUrl}
										remove={removeUrl}
										maxUrls={5}
										validation={() => null}
										icon={<RiLink />}
										helperText={`${urlFields.length}/5 URL - Enterキーで追加、×ボタンで削除`}
									/>
								</div>
							</TabPanel>
						</Tabs>
					</div>
				</form>
			</div>
		</>
	)
}
