'use client'

import { CalloutContext } from '@/app/state'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import Image from 'next/image'
import { useContext, useEffect } from 'react'
import { Button, Dialog, DialogTrigger, DropZone, FileTrigger, Heading, Modal } from 'react-aria-components'
import { RiCloseLine, RiImageAddLine, RiImageLine } from 'react-icons/ri'

type Props = {
	header: React.ReactNode
	userInfo: {
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
	isAuthUser: boolean
}

/**
 * @package
 */
export default function ProfilePresentation({ header, userInfo, isAuthUser }: Props) {
	const { callout, setCallout } = useContext(CalloutContext)

	useEffect(() => {
		if (userInfo.isDone) {
			if (!userInfo.isFetch) {
				setCallout([...callout, { content: 'ユーザ情報の取得に失敗しました', type: 'error' }])
			}
		}
	}, [userInfo.isDone, userInfo.isFetch])

	return (
		<div className="flex flex-col h-screen">
			{header}
			<div className="relative">
				<div className="relative w-full h-60">
					{userInfo.headerPath ? (
						<Image src={userInfo.headerPath} alt="header" className="object-cover rounded-lg" />
					) : (
						<div className="w-full h-60 bg-gray-200" />
					)}
				</div>
				<div className="z-2 absolute -bottom-1/2 left-16">
					<div className="flex items-end gap-4">
						<Image src={userInfo.iconPath} alt="icon" width="171" height="171" className="w-40 h-40 border-2 border-white shadow-md rounded-xl" />
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
											<div className="mx-auto mb-4 relative w-[120px] h-[120px]">
												<DropZone className="relative rounded-xl data-[drop-target]:ring-2 data-[drop-target]:ring-primary data-[drop-target]:ring-offset-2 data-[drop-target]:ring-offset-background">
													<Image src={userInfo.iconPath} alt="icon" width="120" height="120" className="rounded-xl object-cover brightness-50" />
													<FileTrigger>
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
