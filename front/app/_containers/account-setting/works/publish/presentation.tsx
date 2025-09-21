'use client'

import { WorkType } from '@/action/type/work'
import { UserType } from '@/action/type/user'
import HeadContent from '@/components/layout/headContent'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { RiEyeLine, RiFile4Line, RiFocus3Line, RiImageLine, RiLink, RiLockLine, RiMedal2Line, RiUploadLine } from 'react-icons/ri'
import { DropZone } from 'react-aria-components'
import { FileTrigger } from 'react-aria-components'
import Image from 'next/image'
import DPButton from '@/components/ui/button/button'
import { CalloutContext } from '@/app/state'
import { useContext, useEffect } from 'react'
import { getSessionToken } from '@/lib/access'
import uploadWorkThumbnail from '@/action/usecase/works/uploadWorkThumbnail'
import { useRouter } from 'next/navigation'
import SelectField from '@/components/layout/input/selectField'
import { PublishStatus } from '@/action/type/work'
import { Controller, useForm } from 'react-hook-form'
import updateWork from '@/action/usecase/works/updateWork'
import { HeaderButtonContext } from '@/app/account/setting/state'

type PublishWorkPresentationProps = {
	work: WorkType
	user: UserType
}

type PublishWorkFormType = {
	publishStatus: PublishStatus
}

export default function PublishWorkPresentation({ work, user }: PublishWorkPresentationProps) {
	const { callout, setCallout } = useContext(CalloutContext)

	const contextObj = useContext(HeaderButtonContext)
	const setHeaderButton = contextObj === undefined ? null : contextObj.setHeaderButton


	const router = useRouter()

	const { control, watch } = useForm<PublishWorkFormType>({
		defaultValues: {
			publishStatus: work.publish_status,
		},
	})

	const host = process.env.NEXT_PUBLIC_S3_IMAGE_HOST
	let imagePath: string
	if (work.thumbnail_image_url) {
		imagePath = `${host}/work_thumbnails/${user.user_id}/${work.id}/thumbnail/${work.thumbnail_image_url}`
	} else {
		imagePath = `/work_default_thumbnail.png`
	}

	const handleUploadThumbnail = (file: File) => {
		if (file.type !== 'image/png' && file.type !== 'image/jpeg') {
			setCallout([...callout, { content: 'PNGまたはJPEG形式の画像を選択してください', type: 'error' }])
			return
		}

		if (file.size > 1024 * 1024 * 500) {
			setCallout([...callout, { content: '画像のサイズは500MB以下にしてください', type: 'error' }])
			return
		}

		const uploadFlow = async () => {
			const token = await getSessionToken()
			if (!token) {
				setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
				return
			}

			const response = await uploadWorkThumbnail(token, work.id, file)
			if (response.errors) {
				for (const error of response.errors) {
					setCallout([...callout, { content: error, type: 'error' }])
				}
				return
			}

			if (response.data) {
				setCallout([...callout, { content: 'サムネイル画像をアップロードしました', type: 'success' }])
				router.refresh()
			}
		}

		uploadFlow()
	}

	const handleUpdatePublishStatus = async (data: PublishWorkFormType) => {
		const token = await getSessionToken()
		if (!token) {
			setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
			return
		}

		const uploadFlow = async () => {
			const response = await updateWork(token, user.user_id, {
				...work,
				publish_status: data.publishStatus,
			}, true)
			if (response.errors) {
				for (const error of response.errors) {
					setCallout([...callout, { content: error, type: 'error' }])
				}
			}

			if (response.data) {
				setCallout([...callout, { content: '公開範囲の設定に成功しました', type: 'success' }])

				router.push(`/account/setting/works/edit/${work.id}`)
			}
		}

		uploadFlow()
	}

	/* ヘッダーボタン関連 */
	useEffect(() => {
		if (setHeaderButton) {
			setHeaderButton(
				<div className="flex gap-2">
					<DPButton colormode="primary" onPress={() => handleUpdatePublishStatus(watch())}>
						<TextWithIcon icon={<RiUploadLine />}>投稿！</TextWithIcon>
					</DPButton>
				</div>
			)
		}
	}, [setHeaderButton])

	return (
		<>
			<HeadContent title="公開範囲の設定" des="公開範囲の設定" />
			<div className="h-full flex flex-col gap-8 mx-[140px] my-4">
				<div>
					<SectionComponent
					icon={<RiFocus3Line />}
					title="公開範囲の設定" />
					<form>
						<div className="grid grid-cols-1 md:grid-cols-2 gap-4">
							<Controller
								name="publishStatus"
								control={control}
								rules={{
									required: '公開範囲の設定をしてください',
								}}
								render={({ field, fieldState }) => (
									<SelectField
										title="公開範囲の設定"
										field={field}
										fieldState={fieldState}
										isRequired
										helperText="公開範囲の設定をしてください"
										icon={<RiMedal2Line />}
										options={[
											{ label: '非公開', value: 'private', icon: <RiLockLine /> },
											{ label: '限定公開', value: 'limited', icon: <RiLink /> },
											{ label: '全体公開', value: 'public', icon: <RiEyeLine /> }
										]}
										defaultSelectedKey={work.publish_status === 'draft' ? 'private' : work.publish_status}
									/>
								)}
							/>

						</div>
					</form>
				</div>

				<div className="flex flex-row gap-4 justify-between">
					<div className="flex flex-col gap-4 mb-4">
						<SectionComponent
						icon={<RiImageLine />}
						title="サムネイル画像設定"/>

						<div className="text-subtext text-lg leading-relaxed">
							<p>
								サムネイルを設定することで、<br />
								より分かりやすく作品を紹介することができます。<br />
								画像サイズは500MB以下にしてください。
							</p>
						</div>
					</div>
					<div>
						<DropZone
							onDrop={(e) => {
								const targetFile = e.items[0]
								if (targetFile.kind !== 'file') return
								targetFile.getFile().then((f) => {
									handleUploadThumbnail(f)
								})
							}}
							className={`mx-auto relative rounded-xl data-[drop-target]:ring-2 data-[drop-target]:ring-primary data-[drop-target]:ring-offset-2 data-[drop-target]:ring-offset-background`}
						>
							<div className="bg-gray-50 rounded-xl py-4 px-10 border border-subtext border-dashed flex flex-col gap-6 items-center justify-center max-w-[300px] h-full">
								<h3 className="text-lg font-medium text-subtext">
									<TextWithIcon icon={<RiUploadLine />}>画像をアップロード</TextWithIcon>
								</h3>
								<div>
									<p className="text-subtext text-sm text-center mb-1">
										ドラッグ&ドロップ <br />
										もしくは
									</p>
									<FileTrigger
										onSelect={(e) => {
											if (!e) return
											const file = e.item(0)
											if (file === null) return
											handleUploadThumbnail(file)
										}}
										acceptedFileTypes={['image/*']}
									>
										<DPButton colormode="mono" buttonSize="small">
											<TextWithIcon icon={<RiFile4Line />}>ファイルを選択</TextWithIcon>
										</DPButton>
									</FileTrigger>
								</div>
							</div>
						</DropZone>
					</div>
				</div>

				<div>
					<SectionComponent
					icon={<RiEyeLine />}
					title="サムネイルのプレビュー" />

					<div className="flex justify-center">
						<Image src={imagePath} alt="サムネイル" width={300} height={300} />
					</div>
				</div>
			</div>
		</>
	)
}

function SectionComponent({icon, title}: {icon: React.ReactNode, title: string}) {
	return (
		<div className="flex flex-col gap-4 mb-4">
			<h2 className="text-3xl font-bold">
				<TextWithIcon icon={icon}>{title}</TextWithIcon>
			</h2>
		</div>
	)
}
