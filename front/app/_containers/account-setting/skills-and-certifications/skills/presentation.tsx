'use client'

import { SkillType } from '@/action/type/skill'
import { UserType } from '@/action/type/user'
import createSkill from '@/action/usecase/skills/create'
import deleteSkill from '@/action/usecase/skills/delete'
import fetchByUserId from '@/action/usecase/skills/fetchByUserId'
import updateSkill from '@/action/usecase/skills/update'
import updateSkillsSort from '@/action/usecase/skills/updateSort'
import { CalloutContext } from '@/app/state'
import DatePickerField from '@/components/layout/input/datePickerField'
import InputField from '@/components/layout/input/inputField'
import CardEdit from '@/components/layout/setting/cardEdit'
import CardItem from '@/components/layout/setting/cartItem'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getSessionToken } from '@/lib/access'
import { CalendarDate, getLocalTimeZone, parseDate, today } from '@internationalized/date'
import { useContext, useEffect, useRef, useState } from 'react'
import { TextArea } from 'react-aria-components'
import { Controller, useForm } from 'react-hook-form'
import { RiAddLine, RiCalendarLine, RiEditLine, RiMedal2Line, RiMessage2Line } from 'react-icons/ri'

type Props = {
	user: UserType
}

export type SkillFormType = {
	skillName: string
	comment: string
	year: CalendarDate
}

export default function SkillsSettingPresentation({ user }: Props) {
	const textareaRef = useRef<HTMLTextAreaElement>(null)

	const { callout, setCallout } = useContext(CalloutContext)

	const [currentSkills, setCurrentSkills] = useState<SkillType[]>([])
	const [editingSkill, setEditingSkill] = useState<SkillType | null>(null)
	const [isSubmitting, setIsSubmitting] = useState(false)
	const [isFetching, setIsFetching] = useState(false) // スキル一覧の取得

	const { control, handleSubmit, reset } = useForm<SkillFormType>({
		defaultValues: {
			skillName: '',
			comment: '',
			year: today('Asia/Tokyo'),
		},
	})

	useEffect(() => {
		updateSkillsList()
	}, [])

	const updateSkillsList = async () => {
		setIsFetching(true)
		try {
			const res = await fetchByUserId(user.user_id)

			if (res.errors) {
				for (const error of res.errors) {
					setCallout([...callout, { content: error, type: 'error' }])
				}
				return
			}

			const skills = res.data?.skills || []
			setCurrentSkills(skills)
		} catch (error) {
			console.error(error)
			setCallout([...callout, { content: 'スキルの取得に失敗しました', type: 'error' }])
		} finally {
			setIsFetching(false)
		}
	}

	const handleAddSkill = async (data: SkillFormType) => {
		if (isSubmitting) return
		setIsSubmitting(true)

		try {
			const token = await getSessionToken()
			if (!token) {
				setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
				return
			}

			const newSkill: {
				name: string
				year: string
				comment: string
			} = {
				name: data.skillName,
				year: data.year.toString(),
				comment: data.comment,
			}

			const res = await createSkill(token, newSkill)

			if (res.errors) {
				for (const error of res.errors) {
					setCallout([...callout, { content: error, type: 'error' }])
				}
				return
			}
			reset()
			setCallout([...callout, { content: 'スキルを追加しました', type: 'success' }])

			updateSkillsList()
		} catch (error) {
			console.error(error)
			setCallout([...callout, { content: 'スキルの追加に失敗しました', type: 'error' }])
		} finally {
			setIsSubmitting(false)
		}
	}

	const handleUpdateSkill = async (data: SkillFormType) => {
		if (!editingSkill || isSubmitting) return
		setIsSubmitting(true)

		try {
			const token = await getSessionToken()
			if (!token) {
				setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
				return
			}

			const updatedSkill: SkillType = {
				...editingSkill,
				name: data.skillName,
				year: data.year.toString(),
				comment: data.comment,
			}

			const res = await updateSkill(token, updatedSkill)

			if (res.errors) {
				for (const error of res.errors) {
					setCallout([...callout, { content: error, type: 'error' }])
				}
				return
			}
			reset({
				skillName: '',
				comment: '',
				year: today('Asia/Tokyo'),
			})
			setCallout([...callout, { content: 'スキルを更新しました', type: 'success' }])

			setEditingSkill(null)
			updateSkillsList()
		} catch (error) {
			console.error(error)
			setCallout([...callout, { content: 'スキルの更新に失敗しました', type: 'error' }])
		} finally {
			setIsSubmitting(false)
		}
	}

	const handleDeleteSkill = async (skillId: string) => {
		if (isSubmitting) return
		setIsSubmitting(true)

		try {
			const token = await getSessionToken()
			if (!token) {
				setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
				return
			}

			const res = await deleteSkill(token, skillId)

			if (res.errors) {
				for (const error of res.errors) {
					setCallout([...callout, { content: error, type: 'error' }])
				}
				return
			}
			setCallout([...callout, { content: 'スキルを削除しました', type: 'success' }])

			updateSkillsList()
		} catch (error) {
			console.error(error)
			setCallout([...callout, { content: 'スキルの削除に失敗しました', type: 'error' }])
		} finally {
			setIsSubmitting(false)
		}
	}

	const handleEditSkill = (id: string) => {
		const skill = currentSkills.find((skill) => skill.id === id)

		if (!skill) return

		setEditingSkill(skill)

		reset({
			skillName: skill.name,
			comment: skill.comment,
			year: parseDate(skill.year),
		})
	}

	const handleCancelEdit = () => {
		setEditingSkill(null)
		reset({
			skillName: '',
			comment: '',
			year: today('Asia/Tokyo'),
		})
	}

	const handleConfirmSort = async (
		newOrder: {
			id: string
			sort_index: number
		}[],
	) => {
		if (isSubmitting) return
		setIsSubmitting(true)

		try {
			const token = await getSessionToken()
			if (!token) {
				setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
				return
			}

			const res = await updateSkillsSort(token, { sort_list: newOrder })

			if (res.errors) {
				for (const error of res.errors) {
					setCallout([...callout, { content: error, type: 'error' }])
				}
				return
			}

			setCallout([...callout, { content: `ソート順の更新が成功しました`, type: 'success' }])
			updateSkillsList()
		} catch {
			setCallout([...callout, { content: 'ソート順の更新に失敗しました', type: 'error' }])
		} finally {
			setIsSubmitting(false)
		}
	}

	return (
		<div className="space-y-6">
			{/* スキル追加・編集フォーム */}
			<CardEdit
				editigngSkill={!!editingSkill}
				formContent={
					<form onSubmit={handleSubmit(editingSkill ? handleUpdateSkill : handleAddSkill)}>
						<div className="grid grid-cols-1 md:grid-cols-2 gap-4">
							<Controller
								name="skillName"
								control={control}
								rules={{
									required: 'スキル名を入力してください',
									maxLength: {
										value: 50,
										message: 'スキル名は50文字以内で入力してください',
									},
								}}
								render={({ field, fieldState }) => (
									<InputField
										title="スキル名"
										type="text"
										field={field}
										fieldState={fieldState}
										isRequired
										helperText="スキル名は50文字以内で入力してください"
										icon={<RiMedal2Line />}
										autoComplete="off"
										placeholder="例: React, TypeScript, Python"
									/>
								)}
							/>

							<Controller
								name="year"
								control={control}
								rules={{
									required: '学習開始年月日を入力してください',
									validate: (value) => value.compare(today(getLocalTimeZone())) < 0 || '未来の日付は指定できません',
								}}
								render={({ field, fieldState }) => (
									<DatePickerField
										helperText="学習を始めた年月日を入力してください"
										title="学習開始年月日を入力してください"
										field={field}
										fieldState={fieldState}
										isRequired
										icon={<RiCalendarLine />}
									></DatePickerField>
								)}
							></Controller>

							<Controller
								name="comment"
								control={control}
								rules={{
									maxLength: {
										value: 500,
										message: 'コメントは5000文字以内で入力してください',
									},
								}}
								render={({ field, fieldState }) => (
									<InputField
										title="コメント"
										field={field}
										fieldState={fieldState}
										isMultiline
										icon={<RiMessage2Line />}
										rows={10}
										helperText={'コメントは500文字以内で入力してください。'}
										customInput={
											<TextArea
												ref={textareaRef}
												rows={10}
												className="w-full rounded border border-subtext text-foreground p-2 transition duration-200 ease-in-out focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary data-[disabled]:text-subtext"
												placeholder="なぜそのスキルを保有しているのか、どのように活用しているのか書いてみましょう！"
											/>
										}
									/>
								)}
							/>
						</div>

						<div className="flex gap-2 mt-4">
							<DPButton colormode="primary" type="submit" isDisabled={isSubmitting}>
								<TextWithIcon icon={editingSkill ? <RiEditLine /> : <RiAddLine />}>{editingSkill ? '更新する' : '追加する'}</TextWithIcon>
							</DPButton>
							{editingSkill && (
								<DPButton colormode="secondary" type="button" onPress={handleCancelEdit}>
									キャンセル
								</DPButton>
							)}
						</div>
					</form>
				}
			></CardEdit>

			{/* スキル一覧 */}
			<CardItem
				items={currentSkills.map((skill) => {
					return {
						id: skill.id,
						sort_index: skill.sort_index,
						title: skill.name,
						content: (
							<>
								<div className="text-sm text-gray-500">学習開始: {skill.year}</div>
								<div className="text-sm text-gray-600 mt-1">{skill.comment}</div>
							</>
						),
					}
				})}
				isFetching={isFetching}
				emptyItemMessage="スキルが登録されていません。上のフォームから追加しましょう！"
				isSubmitting={isSubmitting}
				handleSort={(newOrder) => handleConfirmSort(newOrder)}
				handleDelete={(id) => handleDeleteSkill(id)}
				handleEdit={(id) => handleEditSkill(id)}
			></CardItem>
		</div>
	)
}
