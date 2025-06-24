import CloseButton from '@/components/ui/button/closeButton'
import { motion } from 'framer-motion'

interface BaseSidebarProps {
	children: React.ReactNode
	isOpen: boolean
	onClose: () => void
	side: 'left' | 'right'
}

export default function BaseSidebar({ children, isOpen, onClose, side }: BaseSidebarProps) {
	const isLeft = side === 'left'

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
				className={`z-20 fixed top-0 pt-8 sm:pt-10 ${isLeft ? 'left-0 border-r rounded-r-lg' : 'right-0 border-l rounded-l-lg'} w-4/5 sm:w-2/3 md:w-1/2 lg:w-1/3 h-full bg-white border-border p-3 sm:p-4 shadow-lg shadow-black/10`}
				initial={{ x: isLeft ? '-100%' : '100%' }}
				animate={isOpen ? { x: 0 } : { x: isLeft ? '-100%' : '100%' }}
				transition={{ type: 'tween', duration: 0.2 }}
				style={{ pointerEvents: isOpen ? 'auto' : 'none' }}
			>
				<CloseButton onClick={onClose} className={`absolute top-3 sm:top-4 ${isLeft ? 'right-3 sm:right-4' : 'left-3 sm:left-4'}`} />
				<div className="flex flex-col gap-1 sm:gap-2">{children}</div>
			</motion.div>
		</>
	)
}
