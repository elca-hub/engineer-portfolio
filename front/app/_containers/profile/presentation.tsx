'use client'

import { CalloutContext } from '@/app/state'
import Image from 'next/image'
import { useContext, useEffect } from 'react'

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
export default function ProfilePresentation({ header, userInfo }: Props) {
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
						<div className="flex flex-col gap-2 mb-2">
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
