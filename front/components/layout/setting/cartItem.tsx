'use client'

import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import React, { useState } from 'react'
import { RiCheckLine, RiDeleteBinLine, RiEditLine, RiListCheck, RiMenu3Line } from 'react-icons/ri'

type cardItemProps = {
	items: {
		id: string
		title: string
		sort_index: number
		content: React.ReactNode
	}[]
	isFetching: boolean
	handleEdit: (id: string) => void
	handleDelete: (id: string) => void
	handleSort: (newOrder: { id: string; sort_index: number }[]) => void
	isSubmitting: boolean
	emptyItemMessage?: string
}

export default function CardItem(props: cardItemProps) {
	const [draggedItem, setDraggedItem] = useState<string | null>(null)
	const [sortedItems, setSortedItems] = useState(props.items)
	const [isSorting, setIsSorting] = useState(false)

	React.useEffect(() => {
		setSortedItems([...props.items])
	}, [props.items])

	const handleDragStart = (e: React.DragEvent, skillId: string) => {
		setDraggedItem(skillId)
		e.dataTransfer.effectAllowed = 'move'
		e.dataTransfer.setData('text/html', skillId)
	}

	const handleDragOver = (e: React.DragEvent) => {
		e.preventDefault()
		e.dataTransfer.dropEffect = 'move'
	}

	const handleDrop = (e: React.DragEvent, targetItemId: string) => {
		e.preventDefault()
		if (!draggedItem || draggedItem === targetItemId) return

		const draggedIndex = sortedItems.findIndex((item) => item.id === draggedItem)
		const targetIndex = sortedItems.findIndex((item) => item.id === targetItemId)

		if (draggedIndex === -1 || targetIndex === -1) return

		const newSortedSkills = [...sortedItems]
		const [draggedSkill] = newSortedSkills.splice(draggedIndex, 1)
		newSortedSkills.splice(targetIndex, 0, draggedSkill)

		setSortedItems(newSortedSkills)
		setDraggedItem(null)
		setIsSorting(true)
	}

	const handleConfirmSort = () => {
		// 差分を比較してsortIndexが変更されたアイテムのみ更新
		const updatedItems: {
			id: string
			sort_index: number
		}[] = []

		sortedItems.forEach((sortedItem, newIndex) => {
			const currentItem = props.items.find((item) => item.id === sortedItem.id)
			if (currentItem && currentItem.sort_index !== newIndex) {
				updatedItems.push({
					id: sortedItem.id,
					sort_index: newIndex,
				})
			}
		})

		// 更新が必要なスキルがない場合
		if (updatedItems.length === 0) {
			setIsSorting(false)
			return
		}

		props.handleSort(updatedItems)
		setIsSorting(false)
	}

	const handleCancelSort = () => {
		setSortedItems([...props.items])
		setIsSorting(false)
		setDraggedItem(null)
	}

	return (
		<div className="bg-white p-6 rounded-lg shadow-sm border">
			<div className="flex justify-between items-center mb-4">
				<h3 className="text-lg font-semibold">
					<TextWithIcon icon={<RiListCheck></RiListCheck>}>一覧</TextWithIcon>
				</h3>
				{isSorting && (
					<div className="flex gap-2">
						<DPButton colormode="primary" onPress={handleConfirmSort} isDisabled={props.isSubmitting}>
							<TextWithIcon icon={<RiCheckLine />}>確定</TextWithIcon>
						</DPButton>
						<DPButton colormode="secondary" onPress={handleCancelSort} isDisabled={props.isSubmitting}>
							キャンセル
						</DPButton>
					</div>
				)}
			</div>
			{props.isFetching ? (
				<div className="space-y-3">
					{[...Array(3)].map((_, index) => (
						<div key={index} className="flex items-center justify-between p-4 border rounded-lg animate-pulse">
							<div className="flex-1">
								<div className="h-5 bg-gray-200 rounded w-1/3 mb-2"></div>
								<div className="h-4 bg-gray-200 rounded w-1/4 mb-2"></div>
								<div className="w-full bg-gray-200 rounded-full h-2">
									<div className="bg-gray-300 h-2 rounded-full w-3/5"></div>
								</div>
							</div>
							<div className="flex gap-2 ml-4">
								<div className="w-8 h-8 bg-gray-200 rounded-md"></div>
								<div className="w-8 h-8 bg-gray-200 rounded-md"></div>
							</div>
						</div>
					))}
				</div>
			) : props.items.length === 0 ? (
				<div className="text-center py-8 text-gray-500">{props.emptyItemMessage || '何も登録されていません！上のフォームから登録しましょう！'}</div>
			) : (
				<div className="space-y-3">
					{sortedItems.map((item) => (
						<div
							key={item.id}
							className={`flex items-center p-4 border rounded-lg transition-all duration-200 ${
								draggedItem === item.id ? 'opacity-50 bg-blue-50' : 'hover:bg-gray-50'
							}`}
							draggable={true}
							onDragStart={(e) => handleDragStart(e, item.id)}
							onDragOver={handleDragOver}
							onDrop={(e) => handleDrop(e, item.id)}
						>
							<div className="mr-3 cursor-grab active:cursor-grabbing text-gray-400 hover:text-gray-600">
								<RiMenu3Line className="text-lg" />
							</div>
							<div className="flex-1">
								<div className="font-medium">{item.title}</div>
								{item.content}
							</div>
							<div className="flex gap-2 ml-4">
								<button
									onClick={() => props.handleEdit(item.id)}
									className="p-2 text-blue-600 hover:bg-blue-50 rounded-md transition-colors"
									disabled={props.isSubmitting}
									title="編集"
								>
									<RiEditLine />
								</button>
								<button
									onClick={() => props.handleDelete(item.id)}
									className="p-2 text-red-600 hover:bg-red-50 rounded-md transition-colors"
									disabled={props.isSubmitting}
									title="削除"
								>
									<RiDeleteBinLine />
								</button>
							</div>
						</div>
					))}
				</div>
			)}
		</div>
	)
}
