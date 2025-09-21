'use client'

import React, { useState } from 'react'
import Image from 'next/image'
import { WorkType } from '@/action/type/work'
import DPButton from '@/components/ui/button/button'
import { RiExternalLinkLine, RiImageLine } from 'react-icons/ri'

interface WorkCardItemProps {
	work: WorkType
	maxDescriptionLength?: number
	onButtonClick?: (work: WorkType) => void
	buttonText?: string
}

export default function WorkCardItem({
	work,
	maxDescriptionLength = 100,
	onButtonClick,
	buttonText = '詳細'
}: WorkCardItemProps) {
	const [imageError, setImageError] = useState(false)

	const truncateText = (text: string, maxLength: number): string => {
		if (text.length <= maxLength) return text
		return text.slice(0, maxLength) + '...'
	}

	const handleButtonClick = () => {
		if (onButtonClick) {
			onButtonClick(work)
		}
	}

	return (
		<div className="bg-white rounded-lg shadow-sm border hover:shadow-md transition-shadow duration-300 overflow-hidden">
			{/* Thumbnail Image Section */}
			<div className="relative w-full h-48 bg-gray-100">
				{work.thumbnail_image_url && !imageError ? (
					<Image
						src={work.thumbnail_image_url}
						alt={work.title}
						fill
						className="object-cover"
						onError={() => setImageError(true)}
					/>
				) : (
					<div className="w-full h-full flex items-center justify-center text-gray-400">
						<RiImageLine className="text-4xl" />
					</div>
				)}
			</div>

			{/* Content Section */}
			<div className="p-4">
				{/* Title */}
				<h3 className="text-lg font-semibold text-gray-900 mb-2 line-clamp-2">
					{work.title}
				</h3>

				{/* Description */}
				<p className="text-gray-600 text-sm mb-4 leading-relaxed">
					{truncateText(work.content, maxDescriptionLength)}
				</p>

				{/* Button */}
				<div className="flex justify-end">
					<DPButton
						colormode="primary"
						buttonSize="small"
						onPress={handleButtonClick}
						className="flex items-center gap-1"
					>
						{buttonText}
						<RiExternalLinkLine className="text-sm" />
					</DPButton>
				</div>
			</div>
		</div>
	)
}