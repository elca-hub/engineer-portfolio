import CloseButton from '@/components/ui/button/closeButton'
import { motion } from 'framer-motion'

interface DBSidebarProps {
	children: React.ReactNode
	isOpen: boolean
	onClose: () => void
}

export default function DBSidebar({ children, isOpen, onClose }: DBSidebarProps) {
	return (
		<>
			<motion.div
				className="z-10 fixed inset-0 bg-black/50 backdrop-blur-sm"
				initial={{ opacity: 0 }}
				animate={isOpen ? { opacity: 1 } : { opacity: 0 }}
				transition={{ type: 'tween', duration: 0.2 }}
				style={{ pointerEvents: isOpen ? 'auto' : 'none' }}
				onClick={onClose}
			/>
			<motion.div
				className="z-20 fixed top-0 pt-10 left-0 w-1/3 h-full bg-white border-r border-border p-4 shadow-lg shadow-black/10 rounded-r-lg"
				initial={{ x: '-100%' }}
				animate={isOpen ? { x: 0 } : { x: '-100%' }}
				transition={{ type: 'tween', duration: 0.2 }}
				style={{ pointerEvents: isOpen ? 'auto' : 'none' }}
			>
				<CloseButton onClick={onClose} className="absolute top-4 right-4" />
				<div className="flex flex-col gap-2">{children}</div>
			</motion.div>
		</>
	)
}
