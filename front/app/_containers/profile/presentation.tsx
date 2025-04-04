'use client'

import uploadIconApi from '@/app/_containers/profile/action'
import { CalloutContext } from '@/app/state'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import Image from 'next/image'
import { useRouter } from 'next/navigation'
import { useContext, useEffect, useState } from 'react'
import { Button, Dialog, DialogTrigger, DropZone, FileTrigger, Heading, Modal } from 'react-aria-components'
import { RiCloseLine, RiImageAddLine, RiImageLine } from 'react-icons/ri'

type userInfoType = {
	userName: string
	userId: string
	iconPath: string
	headerPath: string
	bioPath: string
	skillsLength: number
	externalLinksLength: number
	isFetch: boolean
	isDone: boolean
}

type Props = {
	header: React.ReactNode
	userInfo: userInfoType
	isAuthUser: boolean
}

/**
 * @package
 */
export default function ProfilePresentation({ header, userInfo, isAuthUser }: Props) {
	const router = useRouter()
	const { callout, setCallout } = useContext(CalloutContext)

	const [iconFile, setIconFile] = useState<File>()

	useEffect(() => {
		if (userInfo.isDone) {
			if (!userInfo.isFetch) {
				setCallout([...callout, { content: 'ユーザ情報の取得に失敗しました', type: 'error' }])
			}
		}
	}, [userInfo.isDone, userInfo.isFetch])

	useEffect(() => {
		if (iconFile) {
			if (iconFile.type !== 'image/png' && iconFile.type !== 'image/jpeg') {
				setCallout([...callout, { content: 'PNGまたはJPEG形式の画像を選択してください', type: 'error' }])
				setIconFile(undefined)
				return
			}

			if (iconFile.size > 1024 * 1024 * 50) {
				setCallout([...callout, { content: '画像のサイズは50MB以下にしてください', type: 'error' }])
				setIconFile(undefined)
				return
			}

			const iconUploadFlow = async () => {
				const formData = new FormData()
				formData.append('icon', iconFile)

				const res = await uploadIconApi(formData)

				if (res.errors) {
					setCallout([...callout, { content: 'アップロードに失敗しました', type: 'error' }])
					return
				}

				setCallout([...callout, { content: 'アイコンを変更しました', type: 'info' }])
				router.refresh()
			}

			iconUploadFlow()

			setIconFile(undefined)
		}
	}, [iconFile])

	return (
		<div className="flex flex-col h-screen">
			{header}
			<div className="relative">
				<div className="relative w-full h-60">
					{userInfo.headerPath ? (
						<Image src={userInfo.headerPath} alt="header" className="object-fit rounded-lg" />
					) : (
						<div className="w-full h-60 bg-gray-200" />
					)}
				</div>
				<div className="z-2 absolute -bottom-1/2 left-16">
					<div className="flex items-end gap-4">
						<Image
							src={userInfo.iconPath}
							alt="icon"
							width="171"
							height="171"
							className="w-40 h-40 border-2 border-white shadow-md rounded-xl object-cover"
							priority
						/>
						<div className="flex flex-col mb-1">
							{isAuthUser && (
								<DialogTrigger>
									<DPButton colormode="primary">
										<p className="text-sm">
											<TextWithIcon icon={<RiImageLine />}>画像を編集</TextWithIcon>
										</p>
									</DPButton>
									<Modal className="fixed inset-0 z-50 flex items-center justify-center bg-black/30  entering:animate-in entering:fade-in entering:duration-200 exiting:animate-out exiting:fade-out exiting:duration-200">
										<Dialog className="outline-none bg-white rounded-lg shadow-lg p-6 w-1/2 relative">
											<Button slot="close" className="absolute top-4 right-4">
												<RiCloseLine className="w-8 h-8 hover:scale-90 transition-all duration-200" />
											</Button>
											<Heading slot="title" className="text-2xl font-bold text-foreground mb-4">
												<TextWithIcon icon={<RiImageLine />}>画像を編集</TextWithIcon>
											</Heading>

											<h3 className="text-xl font-medium text-foreground mb-2">アイコン</h3>
											<div className="mx-auto mb-4 relative w-fit h-fit">
												<DropZone
													onDrop={(e) => {
														const targetFile = e.items[0]
														if (targetFile.kind !== 'file') return
														targetFile.getFile().then((f) => {
															setIconFile(f)
														})
													}}
													className="relative rounded-xl data-[drop-target]:ring-2 data-[drop-target]:ring-primary data-[drop-target]:ring-offset-2 data-[drop-target]:ring-offset-background"
												>
													<Image
														src={userInfo.iconPath}
														alt="icon"
														width={250}
														height={250}
														className="w-20 h-20 rounded-xl object-cover brightness-50"
													/>
													<FileTrigger
														onSelect={(e) => {
															if (!e) return
															const file = e.item(0)
															if (file === null) return
															setIconFile(file)
														}}
														acceptedFileTypes={['image/*']}
													>
														<Button className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 rounded-full bg-white/80 p-2 shadow-md hover:scale-95 transition-all duration-200">
															<RiImageAddLine className="text-lg" />
														</Button>
													</FileTrigger>
												</DropZone>
											</div>
										</Dialog>
									</Modal>
								</DialogTrigger>
							)}
							<h1 className="text-3xl font-bold tracking-wide text-foreground">{userInfo.userName}</h1>
							<p className="text-lg font-medium text-gray-600">@{userInfo.userId}</p>
						</div>
					</div>
				</div>
			</div>
			<main className="px-10 mt-2"></main>
		</div>
	)
}
