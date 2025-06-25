import BaseSidebar from './baseSidebar'

interface LeftSidebarProps {
	children: React.ReactNode
	isOpen: boolean
	onClose: () => void
}

export default function LeftSidebar({ children, isOpen, onClose }: LeftSidebarProps) {
	return (
		<BaseSidebar side="left" isOpen={isOpen} onClose={onClose}>
			{children}
		</BaseSidebar>
	)
}
