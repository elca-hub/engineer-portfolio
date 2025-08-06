'use client'

import { UserType } from '@/action/type/user'
import { CalloutContext } from '@/app/state'
import DatePickerField from '@/components/layout/input/datePickerField'
import InputField from '@/components/layout/input/inputField'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getSessionToken } from '@/lib/access'
import { CalendarDate, getLocalTimeZone, parseDate, today } from '@internationalized/date'
import { useRouter } from 'next/navigation'
import { useContext, useState } from 'react'
import { Controller, useForm } from 'react-hook-form'
import { RiAddLine, RiAwardLine, RiCalendarLine, RiDeleteBinLine, RiEditLine, RiSortAsc } from 'react-icons/ri'

type Props = {
	user: UserType
}

export type CertificationFormType = {
	certificationName: string
	organizationName: string
	acquiredDate: CalendarDate
	expirationDate?: CalendarDate
}

type Certification = {
	id: string
	name: string
	organizationName: string
	acquiredDate: string
	expirationDate?: string
}

export default function CertificationsSettingPresentation({ user }: Props) {
	const router = useRouter()
	const { callout, setCallout } = useContext(CalloutContext)

	const [certifications, setCertifications] = useState<Certification[]>([])
	const [editingCertification, setEditingCertification] = useState<Certification | null>(null)
	const [sortAscending, setSortAscending] = useState(true)
	const [isSubmitting, setIsSubmitting] = useState(false)

	const { control, handleSubmit, reset, watch } = useForm<CertificationFormType>({
		defaultValues: {
			certificationName: '',
			organizationName: '',
			acquiredDate: today(getLocalTimeZone()),
			expirationDate: undefined,
		},
	})

	const handleAddCertification = async (data: CertificationFormType) => {
		if (isSubmitting) return
		setIsSubmitting(true)

		try {
			const token = await getSessionToken()
			if (!token) {
				setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
				return
			}

			const newCertification: Certification = {
				id: Date.now().toString(),
				name: data.certificationName,
				organizationName: data.organizationName,
				acquiredDate: data.acquiredDate.toString(),
				expirationDate: data.expirationDate?.toString(),
			}

			setCertifications([...certifications, newCertification])
			reset({
				certificationName: '',
				organizationName: '',
				acquiredDate: today(getLocalTimeZone()),
				expirationDate: undefined,
			})
			setCallout([...callout, { content: '資格を追加しました', type: 'success' }])
		} catch (error) {
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

			const updatedCertifications = certifications.map((cert) =>
				cert.id === editingCertification.id
					? {
							...cert,
							name: data.certificationName,
							organizationName: data.organizationName,
							acquiredDate: data.acquiredDate.toString(),
							expirationDate: data.expirationDate?.toString(),
						}
					: cert,
			)

			setCertifications(updatedCertifications)
			setEditingCertification(null)
			reset({
				certificationName: '',
				organizationName: '',
				acquiredDate: today(getLocalTimeZone()),
				expirationDate: undefined,
			})
			setCallout([...callout, { content: '資格を更新しました', type: 'success' }])
		} catch (error) {
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

			const updatedCertifications = certifications.filter((cert) => cert.id !== certificationId)
			setCertifications(updatedCertifications)
			setCallout([...callout, { content: '資格を削除しました', type: 'success' }])
		} catch (error) {
			setCallout([...callout, { content: '資格の削除に失敗しました', type: 'error' }])
		} finally {
			setIsSubmitting(false)
		}
	}

	const handleEditCertification = (certification: Certification) => {
		setEditingCertification(certification)
		reset({
			certificationName: certification.name,
			organizationName: certification.organizationName,
			acquiredDate: parseDate(certification.acquiredDate),
			expirationDate: certification.expirationDate ? parseDate(certification.expirationDate) : undefined,
		})
	}

	const handleCancelEdit = () => {
		setEditingCertification(null)
		reset({
			certificationName: '',
			organizationName: '',
			acquiredDate: today(getLocalTimeZone()),
			expirationDate: undefined,
		})
	}

	const handleSortCertifications = () => {
		const sortedCertifications = [...certifications].sort((a, b) => {
			if (sortAscending) {
				return new Date(a.acquiredDate).getTime() - new Date(b.acquiredDate).getTime()
			} else {
				return new Date(b.acquiredDate).getTime() - new Date(a.acquiredDate).getTime()
			}
		})
		setCertifications(sortedCertifications)
		setSortAscending(!sortAscending)
	}

	const sortedCertifications = [...certifications].sort((a, b) => {
		if (sortAscending) {
			return new Date(b.acquiredDate).getTime() - new Date(a.acquiredDate).getTime()
		} else {
			return new Date(a.acquiredDate).getTime() - new Date(b.acquiredDate).getTime()
		}
	})

	const formatDate = (dateString: string) => {
		const date = new Date(dateString)
		return date.toLocaleDateString('ja-JP', {
			year: 'numeric',
			month: 'long',
			day: 'numeric',
		})
	}

	const isExpired = (expirationDate?: string) => {
		if (!expirationDate) return false
		return new Date(expirationDate) < new Date()
	}

	return (
		<div className="space-y-6">
			{/* 資格追加・編集フォーム */}
			<div className="bg-white p-6 rounded-lg shadow-sm border">
				<h3 className="text-lg font-semibold mb-4">{editingCertification ? '資格を編集' : '資格を追加'}</h3>
				<form onSubmit={handleSubmit(editingCertification ? handleUpdateCertification : handleAddCertification)}>
					<div className="grid grid-cols-1 md:grid-cols-2 gap-4">
						<Controller
							name="certificationName"
							control={control}
							rules={{
								required: '資格名を入力してください',
								maxLength: {
									value: 100,
									message: '資格名は100文字以内で入力してください',
								},
							}}
							render={({ field, fieldState }) => (
								<InputField
									title="資格名"
									type="text"
									field={field}
									fieldState={fieldState}
									isRequired
									helperText="資格名は100文字以内で入力してください"
									icon={<RiAwardLine />}
									autoComplete="off"
									placeholder="例: 基本情報技術者試験, TOEIC"
								/>
							)}
						/>

						<Controller
							name="organizationName"
							control={control}
							rules={{
								required: '発行機関を入力してください',
								maxLength: {
									value: 100,
									message: '発行機関は100文字以内で入力してください',
								},
							}}
							render={({ field, fieldState }) => (
								<InputField
									title="発行機関"
									type="text"
									field={field}
									fieldState={fieldState}
									isRequired
									helperText="発行機関は100文字以内で入力してください"
									icon={<RiAwardLine />}
									autoComplete="off"
									placeholder="例: IPA, ETS"
								/>
							)}
						/>

						<Controller
							name="acquiredDate"
							control={control}
							rules={{
								required: '取得日を入力してください',
								validate: (value) => value.compare(today(getLocalTimeZone())) <= 0 || '未来の日付は指定できません',
							}}
							render={({ field, fieldState }) => (
								<DatePickerField title="取得日" field={field} fieldState={fieldState} isRequired icon={<RiCalendarLine />} />
							)}
						/>

						<Controller
							name="expirationDate"
							control={control}
							rules={{
								validate: (value) => {
									if (!value) return true
									const acquiredDate = watch('acquiredDate')
									return value.compare(acquiredDate) > 0 || '有効期限は取得日より後の日付を指定してください'
								},
							}}
							render={({ field, fieldState }) => (
								<DatePickerField
									title="有効期限（任意）"
									field={field}
									fieldState={fieldState}
									icon={<RiCalendarLine />}
									helperText="永久資格の場合は空白にしてください"
								/>
							)}
						/>
					</div>

					<div className="flex gap-2 mt-4">
						<DPButton colormode="primary" type="submit" disabled={isSubmitting}>
							<TextWithIcon icon={editingCertification ? <RiEditLine /> : <RiAddLine />}>
								{editingCertification ? '更新する' : '追加する'}
							</TextWithIcon>
						</DPButton>
						{editingCertification && (
							<DPButton colormode="secondary" type="button" onClick={handleCancelEdit}>
								キャンセル
							</DPButton>
						)}
					</div>
				</form>
			</div>

			{/* 資格一覧 */}
			<div className="bg-white p-6 rounded-lg shadow-sm border">
				<div className="flex justify-between items-center mb-4">
					<h3 className="text-lg font-semibold">資格一覧</h3>
					{certifications.length > 0 && (
						<DPButton colormode="secondary" onClick={handleSortCertifications}>
							<TextWithIcon icon={<RiSortAsc />}>{sortAscending ? '新しい順' : '古い順'}</TextWithIcon>
						</DPButton>
					)}
				</div>

				{certifications.length === 0 ? (
					<div className="text-center py-8 text-gray-500">まだ資格が登録されていません。上のフォームから資格を追加してください。</div>
				) : (
					<div className="space-y-3">
						{sortedCertifications.map((certification) => (
							<div
								key={certification.id}
								className={`flex items-center justify-between p-4 border rounded-lg hover:bg-gray-50 ${
									isExpired(certification.expirationDate) ? 'border-red-200 bg-red-50' : ''
								}`}
							>
								<div className="flex-1">
									<div className="font-medium">{certification.name}</div>
									<div className="text-sm text-gray-600 mt-1">発行機関: {certification.organizationName}</div>
									<div className="text-sm text-gray-500 mt-1">取得日: {formatDate(certification.acquiredDate)}</div>
									{certification.expirationDate && (
										<div className={`text-sm mt-1 ${isExpired(certification.expirationDate) ? 'text-red-600 font-medium' : 'text-gray-500'}`}>
											有効期限: {formatDate(certification.expirationDate)}
											{isExpired(certification.expirationDate) && ' (期限切れ)'}
										</div>
									)}
									{!certification.expirationDate && <div className="text-sm text-green-600 mt-1">永久資格</div>}
								</div>
								<div className="flex gap-2 ml-4">
									<button
										onClick={() => handleEditCertification(certification)}
										className="p-2 text-blue-600 hover:bg-blue-50 rounded-md transition-colors"
										disabled={isSubmitting}
										title="編集"
									>
										<RiEditLine />
									</button>
									<button
										onClick={() => handleDeleteCertification(certification.id)}
										className="p-2 text-red-600 hover:bg-red-50 rounded-md transition-colors"
										disabled={isSubmitting}
										title="删除"
									>
										<RiDeleteBinLine />
									</button>
								</div>
							</div>
						))}
					</div>
				)}
			</div>
		</div>
	)
}
