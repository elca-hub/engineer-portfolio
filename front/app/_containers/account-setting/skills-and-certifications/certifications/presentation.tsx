'use client'

import { CertificationType } from '@/action/type/certification'
import { UserType } from '@/action/type/user'
import createCertification from '@/action/usecase/certifications/create'
import deleteCertification from '@/action/usecase/certifications/delete'
import fetchCertificationByUserId from '@/action/usecase/certifications/fetchByUserId'
import updateCertification from '@/action/usecase/certifications/update'
import updateCertificationSort from '@/action/usecase/certifications/updateSort'
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
import { RiAddLine, RiCalendarLine, RiEditLine, RiMedalLine, RiMessage2Line } from 'react-icons/ri'

type Props = {
	user: UserType
}

export type CertificationFormType = {
	certificationName: string
	comment: string
	year: CalendarDate
}

/**
 * @package
 */
export default function CertificationsSettingPresentation({ user }: Props) {
	const textareaRef = useRef<HTMLTextAreaElement>(null)

	const { callout, setCallout } = useContext(CalloutContext)

	const [currentCertifications, setCurrentCertifications] = useState<CertificationType[]>([])
	const [editingCertification, setEditingCertification] = useState<CertificationType | null>(null)
	const [isSubmitting, setIsSubmitting] = useState(false)
	const [isFetching, setIsFetching] = useState(false)

	const { control, handleSubmit, reset } = useForm<CertificationFormType>({
		defaultValues: {
			certificationName: '',
			comment: '',
			year: today('Asia/Tokyo'),
		},
	})

	useEffect(() => {
		updateCertificationList()
	}, [])

	const updateCertificationList = async () => {
		setIsFetching(true)
		try {
			const res = await fetchCertificationByUserId(user.user_id)

			if (res.errors) {
				for (const error of res.errors) {
					setCallout([...callout, { content: error, type: 'error' }])
				}
				return
			}

			const certifications = res.data?.certifications || []
			setCurrentCertifications(certifications)
		} catch (error) {
			console.error(error)
			setCallout([...callout, { content: '資格の取得に失敗しました', type: 'error' }])
		} finally {
			setIsFetching(false)
		}
	}

	const handleAddCertification = async (data: CertificationFormType) => {
		if (isSubmitting) return
		setIsSubmitting(true)

		try {
			const token = await getSessionToken()
			if (!token) {
				setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
				return
			}

			const newCertification: {
				name: string
				year: string
				comment: string
			} = {
				name: data.certificationName,
				year: data.year.toString(),
				comment: data.comment,
			}

			const res = await createCertification(token, newCertification)

			if (res.errors) {
				for (const error of res.errors) {
					setCallout([...callout, { content: error, type: 'error' }])
				}
				return
			}
			reset()
			setCallout([...callout, { content: '資格を追加しました', type: 'success' }])

			updateCertificationList()
		} catch (error) {
			console.error(error)
			setCallout([...callout, { content: '資格の追加に失敗しました', type: 'error' }])
		} finally {
			setIsSubmitting(false)
		}
	}

	const handleUpdateCertification = async (data: CertificationFormType) => {
		if (!editingCertification || isSubmitting) return
		setIsSubmitting(true)

		try {
			const token = await getSessionToken()
			if (!token) {
				setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
				return
			}

			const updateCertificationValue: CertificationType = {
				...editingCertification,
				name: data.certificationName,
				year: data.year.toString(),
				comment: data.comment,
			}

			const res = await updateCertification(token, updateCertificationValue)

			if (res.errors) {
				for (const error of res.errors) {
					setCallout([...callout, { content: error, type: 'error' }])
				}
				return
			}

			reset({
				certificationName: '',
				comment: '',
				year: today('Asia/Tokyo'),
			})

			setCallout([...callout, { content: '資格を更新しました', type: 'success' }])

			setEditingCertification(null)
			updateCertificationList()
		} catch (error) {
			console.error(error)
			setCallout([...callout, { content: '資格の更新に失敗しました', type: 'error' }])
		} finally {
			setIsSubmitting(false)
		}
	}

	const handleDeleteCertification = async (certificationId: string) => {
		if (isSubmitting) return
		setIsSubmitting(true)

		try {
			const token = await getSessionToken()
			if (!token) {
				setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
				return
			}

			const res = await deleteCertification(token, certificationId)

			if (res.errors) {
				for (const error of res.errors) {
					setCallout([...callout, { content: error, type: 'error' }])
				}
				return
			}
			setCallout([...callout, { content: '資格を削除しました', type: 'success' }])

			updateCertificationList()
		} catch (error) {
			console.error(error)
			setCallout([...callout, { content: '資格の削除に失敗しました', type: 'error' }])
		} finally {
			setIsSubmitting(false)
		}
	}

	const handleEditCertification = (id: string) => {
		const certification = currentCertifications.find((certification) => certification.id === id)

		if (!certification) return

		setEditingCertification(certification)

		reset({
			certificationName: certification.name,
			comment: certification.comment,
			year: parseDate(certification.year),
		})
	}

	const handleCancelEdit = () => {
		setEditingCertification(null)
		reset({
			certificationName: '',
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

			const res = await updateCertificationSort(token, { sort_list: newOrder })

			if (res.errors) {
				for (const error of res.errors) {
					setCallout([...callout, { content: error, type: 'error' }])
				}
				return
			}

			updateCertificationList()
		} catch {
			setCallout([...callout, { content: 'ソート順の更新に失敗しました', type: 'error' }])
		} finally {
			setIsSubmitting(false)
		}
	}

	return (
		<div className="space-y-6">
			<CardEdit
				editing={!!editingCertification}
				formContent={
					<form onSubmit={handleSubmit(editingCertification ? handleUpdateCertification : handleAddCertification)}>
						<div className="grid grid-cols-1 md:grid-cols-2 gap-4">
							<Controller
								name="certificationName"
								control={control}
								rules={{
									required: '資格名を入力してください',
									maxLength: {
										value: 50,
										message: '資格名は50文字以内で入力してください',
									},
								}}
								render={({ field, fieldState }) => (
									<InputField
										title="資格名"
										type="text"
										field={field}
										fieldState={fieldState}
										isRequired
										helperText="資格名は50文字以内で入力してください"
										icon={<RiMedalLine />}
										autoComplete="off"
										placeholder="例: 応用情報技術者試験"
									/>
								)}
							/>

							<Controller
								name="year"
								control={control}
								rules={{
									required: '資格取得年月日を入力してください',
									validate: (value) => value.compare(today(getLocalTimeZone())) < 0 || '未来の日付は指定できません',
								}}
								render={({ field, fieldState }) => (
									<DatePickerField
										helperText="資格を取得した年月日を入力してください"
										title="資格取得年月日"
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
										message: 'コメントは500文字以内で入力してください',
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
												placeholder="なぜその資格を取ったのか、どのような苦労があったのか自由に書いてみましょう！"
											/>
										}
									/>
								)}
							/>
						</div>

						<div className="flex gap-2 mt-4">
							<DPButton colormode="primary" type="submit" isDisabled={isSubmitting}>
								<TextWithIcon icon={editingCertification ? <RiEditLine /> : <RiAddLine />}>
									{editingCertification ? '更新する' : '追加する'}
								</TextWithIcon>
							</DPButton>
							{editingCertification && (
								<DPButton colormode="secondary" type="button" onPress={handleCancelEdit}>
									キャンセル
								</DPButton>
							)}
						</div>
					</form>
				}
			></CardEdit>

			<CardItem
				items={currentCertifications.map((skill) => {
					return {
						id: skill.id,
						sort_index: skill.sort_index,
						title: skill.name,
						content: (
							<>
								<div className="text-sm text-gray-500">取得(合格): {skill.year}</div>
								<div className="text-sm text-gray-600 mt-1">{skill.comment}</div>
							</>
						),
					}
				})}
				isFetching={isFetching}
				emptyItemMessage="資格が登録されていません。上のフォームから追加しましょう！"
				isSubmitting={isSubmitting}
				handleSort={(newOrder) => handleConfirmSort(newOrder)}
				handleDelete={(id) => handleDeleteCertification(id)}
				handleEdit={(id) => handleEditCertification(id)}
			></CardItem>
		</div>
	)
}
