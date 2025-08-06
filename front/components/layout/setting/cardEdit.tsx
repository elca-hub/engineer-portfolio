'use client'

import TextWithIcon from '@/components/ui/text/textWithIcon'
import { RiAddLine, RiEditLine } from 'react-icons/ri'

type cardEditProps = {
	editigngSkill: boolean
	formContent: React.ReactNode
}

export default function CardEdit(props: cardEditProps) {
	return (
		<div className="bg-white p-6 rounded-lg shadow-sm border">
			<h3 className="text-lg font-semibold mb-4">
				{props.editigngSkill ? (
					<TextWithIcon icon={<RiEditLine></RiEditLine>}>編集</TextWithIcon>
				) : (
					<TextWithIcon icon={<RiAddLine></RiAddLine>}>追加</TextWithIcon>
				)}
			</h3>
			{props.formContent}
		</div>
	)
}
