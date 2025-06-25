import BaseSidebar from './baseSidebar'

interface RightSidebarProps {
	children: React.ReactNode
	isOpen: boolean
	onClose: () => void
}

export default function RightSidebar({ children, isOpen, onClose }: RightSidebarProps) {
	return (
		<BaseSidebar side="right" isOpen={isOpen} onClose={onClose}>
			{children}
		</BaseSidebar>
	)
}
