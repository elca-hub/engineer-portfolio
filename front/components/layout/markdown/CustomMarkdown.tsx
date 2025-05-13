'use client'

import { defaultComponents } from '@/components/layout/markdown/defaultComponents'
import { CustomComponents } from '@/components/layout/markdown/types'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { AnimatePresence, motion } from 'framer-motion'
import { useEffect, useState } from 'react'
import { RiArrowLeftWideLine, RiArrowRightWideLine, RiBookLine } from 'react-icons/ri'
import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'

type Props = {
	children: string
	components?: CustomComponents
}

type SlideInfo = {
	title: string
	content: string
}

export default function CustomMarkdown({ children, components }: Props) {
	const [currentPage, setCurrentPage] = useState(0)
	const [slides, setSlides] = useState<SlideInfo[]>([])
	const [direction, setDirection] = useState<'left' | 'right'>('right')
	const [isFirstLoad, setIsFirstLoad] = useState(true)

	useEffect(() => {
		const sections = children.split('---')
		const processedSlides = sections.map((section) => {
			const lines = section.split('\n')
			const titleLine = lines.find((line) => line.startsWith('# '))
			const title = titleLine ? titleLine.replace('# ', '').trim() : '無題'
			return {
				title,
				content: section,
			}
		})
		setSlides(processedSlides)
		setCurrentPage(0)
		setIsFirstLoad(true)
	}, [children])

	useEffect(() => {
		if (isFirstLoad) {
			setTimeout(() => setIsFirstLoad(false), 0)
		}
	}, [isFirstLoad])

	const handlePrevPage = () => {
		if (currentPage > 0) {
			setDirection('left')
			setCurrentPage((prev) => prev - 1)
		}
	}

	const handleNextPage = () => {
		if (currentPage < slides.length - 1) {
			setDirection('right')
			setCurrentPage((prev) => prev + 1)
		}
	}

	const handleTitleClick = (index: number) => {
		if (index !== currentPage) {
			setDirection(index > currentPage ? 'right' : 'left')
			setCurrentPage(index)
		}
	}

	const variants = {
		enter: (direction: 'left' | 'right') => ({
			x: direction === 'right' ? 300 : -300,
			opacity: 0,
		}),
		center: {
			x: 0,
			opacity: 1,
		},
		exit: (direction: 'left' | 'right') => ({
			x: direction === 'right' ? -300 : 300,
			opacity: 0,
		}),
	}

	return slides.length > 1 ? (
		<div className="relative">
			<div className="mb-8 p-4 bg-gray-50 rounded-lg">
				<h2 className="text-lg font-bold mb-2">
					<TextWithIcon icon={<RiBookLine />}>目次</TextWithIcon>
				</h2>
				<ul className="space-y-1">
					{slides.map((slide, index) => (
						<li key={index}>
							<button
								onClick={() => handleTitleClick(index)}
								className={`text-left w-full px-2 py-1 rounded hover:bg-gray-200 transition-colors ${
									currentPage === index ? 'bg-gray-200 font-bold' : ''
								}`}
							>
								{index + 1}. {slide.title}
							</button>
						</li>
					))}
				</ul>
			</div>
			<div className="min-h-[50vh] relative overflow-x-hidden">
				<AnimatePresence custom={direction} mode="wait">
					<motion.div
						key={currentPage}
						custom={direction}
						variants={variants}
						initial={isFirstLoad ? 'center' : 'enter'}
						animate="center"
						exit={isFirstLoad ? 'center' : 'exit'}
						transition={{ duration: 0.4, ease: 'easeInOut' }}
						className="absolute w-full"
					>
						<ReactMarkdown
							remarkPlugins={[remarkGfm]}
							components={{
								...defaultComponents,
								...components,
							}}
						>
							{slides[currentPage]?.content || ''}
						</ReactMarkdown>
					</motion.div>
				</AnimatePresence>
			</div>
			<div className="flex justify-center gap-4 mt-8 items-center">
				<button
					onClick={handlePrevPage}
					disabled={currentPage === 0}
					className="h-fit p-2 rounded-full disabled:opacity-50 disabled:cursor-not-allowed hover:bg-primary transition-colors"
				>
					<RiArrowLeftWideLine />
				</button>
				<span className="text-subtext select-none">
					{currentPage + 1} / {slides.length}
				</span>
				<button
					onClick={handleNextPage}
					disabled={currentPage === slides.length - 1}
					className="h-fit p-2 rounded-full disabled:opacity-50 disabled:cursor-not-allowed hover:bg-primary transition-colors"
				>
					<RiArrowRightWideLine />
				</button>
			</div>
		</div>
	) : (
		<ReactMarkdown
			remarkPlugins={[remarkGfm]}
			components={{
				...defaultComponents,
				...components,
			}}
		>
			{slides[0]?.content || ''}
		</ReactMarkdown>
	)
}
