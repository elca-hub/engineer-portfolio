'use client'

import { CalloutContext } from '@/app/state'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { defaultUserIcon } from '@/constants/constant'
import { useContext, useEffect } from 'react'
import { RiInfoI, RiSurveyLine } from 'react-icons/ri'

type Props = {
	header: React.ReactNode
	userInfo: {
		userName: string
		iconPath: string
		headerPath: string
		bioPath: string
		skillsLength: number
		externalLinksLength: number
		isFetch: boolean
		isDone: boolean
	}
}

/**
 * @package
 */
export default function DashboardPresentation({ header, userInfo }: Props) {
	const { callout, setCallout } = useContext(CalloutContext)

	useEffect(() => {
		if (userInfo.isDone) {
			if (!userInfo.isFetch) {
				setCallout([...callout, { content: 'ユーザ情報の取得に失敗しました', type: 'error' }])
			}
		}
	}, [userInfo.isDone, userInfo.isFetch])

	return (
		<div className="flex flex-col h-screen gap-2">
			{header}

			<main className="px-10">
				<div className="flex flex-col items-center gap-2">
					<div className="text-foreground text-4xl font-bold">
						<TextWithIcon icon={<RiSurveyLine></RiSurveyLine>}>ダッシュボード</TextWithIcon>
					</div>
					<h2 className="text-gray-600">ここでは{userInfo.userName}さんに関する情報を見ることができます</h2>
				</div>

				<div className="flex flex-col items-center justify-center gap-2 mt-4">
					{userInfo.iconPath === defaultUserIcon && (
						<div className="bg-blue-50 rounded-lg border-2 border-blue-100 py-2 px-4 w-full">
							<p className="text-gray-800 font-medium text-lg mb-1">
								<TextWithIcon icon={<RiInfoI></RiInfoI>}>アイコンが設定されていません</TextWithIcon>
							</p>
							<p className="text-gray-600 text-sm">アイコンを設定して自分がどんな人か一目でわかるようにしましょう！</p>
						</div>
					)}
				</div>
			</main>
		</div>
	)
}
