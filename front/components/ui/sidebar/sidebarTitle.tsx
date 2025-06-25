import TextWithIcon from '@/components/ui/text/textWithIcon'
import { ReactNode } from 'react'

interface SidebarTitleProps {
	children: ReactNode
	icon?: ReactNode
}

export default function SidebarTitle({ children, icon }: SidebarTitleProps) {
	return (
		<h3 className="text-2xl font-medium mb-3">
			<TextWithIcon icon={icon}>{children}</TextWithIcon>
		</h3>
	)
}
