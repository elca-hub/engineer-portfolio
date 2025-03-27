'use client'

import TextWithIcon from '@/components/ui/text/textWithIcon'
import { RiSurveyLine } from 'react-icons/ri'

type Props = {
	header: React.ReactNode
	userName: string
}

/**
 * @package
 */
export default function DashboardPresentation({ header, userName }: Props) {
	return (
		<div className="flex flex-col h-screen gap-2">
			{header}

			<main>
				<div className="flex flex-col items-center gap-2">
					<div className="text-foreground text-4xl font-bold">
						<TextWithIcon icon={<RiSurveyLine></RiSurveyLine>}>ダッシュボード</TextWithIcon>
					</div>
					<h2 className="text-gray-600">ここでは{userName}さんに関する情報を見ることができます</h2>
				</div>
			</main>
		</div>
	)
}
