'use client'

import { defaultComponents } from '@/components/layout/markdown/defaultComponents'
import { CustomComponents } from '@/components/layout/markdown/types'
import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'

type Props = {
	children: string
	components?: CustomComponents
}

export default function CustomMarkdown({ children, components }: Props) {
	return (
		<ReactMarkdown
			remarkPlugins={[remarkGfm]}
			components={{
				...defaultComponents,
				...components,
			}}
		>
			{children}
		</ReactMarkdown>
	)
}
