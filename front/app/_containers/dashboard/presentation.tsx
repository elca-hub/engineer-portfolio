'use client'

import { CalloutContext } from '@/app/state'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { defaultUserIcon } from '@/constants/constant'
import { useContext, useEffect } from 'react'
import { Button, Disclosure, DisclosurePanel, Heading } from 'react-aria-components'
import { RiArrowRightSLine, RiSurveyLine } from 'react-icons/ri'

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

			<main>
				<div className="flex flex-col items-center gap-2">
					<div className="text-foreground text-4xl font-bold">
						<TextWithIcon icon={<RiSurveyLine></RiSurveyLine>}>ダッシュボード</TextWithIcon>
					</div>
					<h2 className="text-gray-600">ここでは{userInfo.userName}さんに関する情報を見ることができます</h2>
				</div>

				<div className="flex flex-col items-center justify-center gap-2">
					{userInfo.iconPath === defaultUserIcon && (
						<Disclosure>
							<Heading>
								<Button slot="trigger">
									<TextWithIcon icon={<RiArrowRightSLine></RiArrowRightSLine>}>自分好みのアイコンを設定しましょう！</TextWithIcon>
								</Button>
							</Heading>
							<DisclosurePanel>
								<p className="text-gray-600">アイコンは{userInfo.headerPath}から設定できます</p>
							</DisclosurePanel>
						</Disclosure>
					)}
				</div>
			</main>
		</div>
	)
}
