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
import { TextArea } from 'react-aria-components'
import { Controller, useFieldArray, useForm } from 'react-hook-form'
import { RiEditLine, RiGithubLine, RiPriceTag3Line, RiArticleLine, RiEyeLine, RiFileTextLine, RiCheckboxCircleLine, RiLink } from 'react-icons/ri'
import updateWork from '@/action/usecase/works/updateWork'
import TagField from '@/components/layout/input/tagField'
import { WorkUrlType } from '@/action/type/workUrl'
import UrlListField from '@/components/layout/input/urlListField'

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

	const [lastUpdatedTime, setLastUpdatedTime] = useState<Date | null>(null)
	const [autoSaveTimeout, setAutoSaveTimeout] = useState<NodeJS.Timeout | null>(null)
	const autoSaveTimeoutRef = useRef<NodeJS.Timeout | null>(null)

	const { control, handleSubmit, watch, formState, reset, trigger } = useForm<WorkFormType>({
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

	return (
		<div>
			{/* 最終更新日時表示 */}
			{lastUpdatedTime && (
				<div className="flex justify-left gap-6 text-green-800 mb-4">
					<TextWithIcon icon={<RiCheckboxCircleLine />}>最終更新日時</TextWithIcon>
					{formatDateTime(lastUpdatedTime)}
				</div>
			)}

			<form onSubmit={handleSubmit((data) => handleUpdateWork(data, true))}>
				<div className="space-y-6">
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
								title="作品名を入力してください"
								type="text"
								field={field}
								fieldState={fieldState}
								helperText={`${field.value.length}/100`}
								icon={<RiFileTextLine />}
								autoComplete="on"
								placeholder="例: ポートフォリオサイト"
							/>
						)}
					/>

					<TagField
						title="タグ"
						helperText={`${fields.length}/${MAX_TAGS} タグ - Enterキーで追加、×ボタンで削除`}
						icon={<RiPriceTag3Line />}
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
					></TagField>

					<Controller
						name="githubRepositoryUrl"
						control={control}
						rules={{
							pattern: {
								value: /^https:\/\/github\.com\/.+/,
								message: 'GitHubのリポジトリURLを入力してください'
							},
						}}
						render={({ field, fieldState }) => (
							<InputField
								title="Githubリポジトリ"
								type="url"
								field={field}
								fieldState={fieldState}
								helperText="GitHubリポジトリのURLを入力してください"
								icon={<RiGithubLine />}
								autoComplete="on"
								placeholder="https://github.com/username/repository"
							/>
						)}
					/>

					<div className="space-y-4">
						<UrlListField
							title="URL"
							urls={urlFields}
							append={appendUrl}
							remove={removeUrl}
							maxUrls={5}
							validation={() => null}
							icon={<RiLink />}
							helperText={`${urlFields.length}/5 URL - Enterキーで追加、×ボタンで削除`}
						/>
					</div>

					<div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
						<div>
							<Controller
								name="content"
								control={control}
								rules={{
									required: '作品の紹介文を入力してください',
									maxLength: {
										value: 2000,
										message: '本文は2000文字以内で入力してください',
									},
								}}
								render={({ field, fieldState }) => (
									<InputField
										title="本文"
										field={field}
										fieldState={fieldState}
										isMultiline
										isRequired
										icon={<RiArticleLine />}
										rows={20}
										helperText="作品の紹介やこだわりポイントをアピールしてみましょう！markdownを利用して記述することができます。"
										customInput={
											<TextArea
												ref={textareaRef}
												rows={20}
												className="w-full rounded border border-subtext text-foreground p-2 transition duration-200 ease-in-out focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary data-[disabled]:text-subtext"
												placeholder="作品の紹介やこだわりポイントをアピールしてみましょう！"
											/>
										}
									/>
								)}
							/>
						</div>

						<div>
							<div className="flex items-center gap-2 mb-2">
								<RiEyeLine className="text-primary" />
								<h3 className="text-lg font-semibold">プレビュー</h3>
							</div>
							<p className="text-sm text-subtext mb-4">markdownを利用して記述することができます。</p>
							<div className="min-h-[400px] border border-subtext rounded p-4 bg-background">
								<CustomMarkdown>{contentValue || ''}</CustomMarkdown>
							</div>
						</div>
					</div>
				</div>

				<div className="flex gap-2 mt-6">
					<DPButton colormode="primary" type="submit" isDisabled={isSubmitting}>
						<TextWithIcon icon={<RiEditLine />}>更新する</TextWithIcon>
					</DPButton>
				</div>
			</form>
		</div>
	)
}
