'use client'

type cardEditProps = {
	editigngSkill: boolean
	formContent: React.ReactNode
}

export default function CardEdit(props: cardEditProps) {
	return (
		<div className="bg-white p-6 rounded-lg shadow-sm border">
			<h3 className="text-lg font-semibold mb-4">{props.editigngSkill ? 'スキルを編集' : 'スキルを追加'}</h3>
			{props.formContent}
		</div>
	)
}
