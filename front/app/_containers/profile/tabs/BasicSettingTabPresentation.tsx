import { UserType } from '@/action/type/user'
import userUpdate from '@/action/usecase/user/update'
import { CalloutContext } from '@/app/state'
import DatePickerField from '@/components/layout/input/datePickerField'
import InputField from '@/components/layout/input/inputField'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getSessionToken } from '@/lib/access'
import { CalendarDate, getLocalTimeZone, parseDate, today } from '@internationalized/date'
import { useRouter } from 'next/navigation'
import { useContext, useEffect, useState } from 'react'
import { Controller, useForm } from 'react-hook-form'
import { RiBriefcaseLine, RiBuilding2Line, RiCake2Line, RiMapPinLine, RiUserLine } from 'react-icons/ri'

type Props = {
	user: UserType
}

export type UserUpdateFormType = {
	name: string
	birthday: CalendarDate
	organizationName: string
	occupationName: string
	place: string
}

export default function BasicSettingTabPresentation({ user }: Props) {
	const router = useRouter()

	const { callout, setCallout } = useContext(CalloutContext)

	const { control, handleSubmit, watch } = useForm<UserUpdateFormType>({
		defaultValues: {
			name: user.name,
			birthday: parseDate(user.birthday),
			organizationName: user.organization_name,
			occupationName: user.occupation_name,
			place: user.place,
		},
	})
	const [isSubmit, setIsSubmit] = useState(false)

	useEffect(() => {
		if (isSubmit) {
			const name = watch('name')
			const birthday = watch('birthday')
			const organizationName = watch('organizationName')
			const occupationName = watch('occupationName')
			const place = watch('place')

			const updateFlow = async () => {
				const token = await getSessionToken()
				if (!token) {
					setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
					return
				}

				user.name = name
				user.birthday = birthday.toString()
				user.organization_name = organizationName
				user.occupation_name = occupationName
				user.place = place

				const res = await userUpdate(token, {
					userData: user,
				})
				if (res.errors) {
					for (const error of res.errors) {
						setCallout([...callout, { content: error, type: 'error' }])
					}
				}

				if (res.data) {
					setCallout([...callout, { content: '変更しました', type: 'success' }])
					router.refresh()
					return
				}
			}

			updateFlow()

			setIsSubmit(false)
		}
	}, [isSubmit])

	return (
		<form onSubmit={handleSubmit(() => setIsSubmit(true))}>
			<Controller
				name="name"
				control={control}
				rules={{
					required: 'ユーザ名を入力してください',
					max: {
						value: 50,
						message: 'ユーザ名は50文字以内で入力してください',
					},
				}}
				render={({ field, fieldState }) => (
					<InputField
						title="ユーザ名"
						type="text"
						field={field}
						fieldState={fieldState}
						isRequired
						helperText="ユーザ名は50文字以下で入力してください。特殊記号は使用できません。"
						icon={<RiUserLine />}
						autoComplete="off"
						popoverContent="本名を入力する必要はありません。自分の個性的な名前をつけましょう！"
						placeholder={user.name}
					></InputField>
				)}
			></Controller>

			<Controller
				name="birthday"
				control={control}
				rules={{
					required: '生年月日が未入力です',
					validate: (value) => value.compare(today(getLocalTimeZone())) < 0 || '未来の日付は指定できません',
				}}
				render={({ field, fieldState }) => (
					<DatePickerField title="生年月日" field={field} fieldState={fieldState} isRequired icon={<RiCake2Line />}></DatePickerField>
				)}
			></Controller>

			<Controller
				name="organizationName"
				control={control}
				rules={{
					max: {
						value: 50,
						message: '組織名は50文字以内で入力してください',
					},
				}}
				render={({ field, fieldState }) => (
					<InputField
						title="組織・企業名"
						type="text"
						field={field}
						fieldState={fieldState}
						helperText="組織名は50文字以内で入力してください"
						icon={<RiBuilding2Line />}
						autoComplete="off"
						placeholder={user.organization_name || 'あなたの所属している組織や企業を紹介しましょう！'}
					></InputField>
				)}
			></Controller>

			<Controller
				name="occupationName"
				control={control}
				rules={{
					max: {
						value: 50,
						message: '職業は50文字以内で入力してください',
					},
				}}
				render={({ field, fieldState }) => (
					<InputField
						title="職業"
						type="text"
						field={field}
						fieldState={fieldState}
						helperText="職業は50文字以内で入力してください"
						icon={<RiBriefcaseLine />}
						autoComplete="off"
						placeholder={user.occupation_name || 'あなたの職業を紹介しましょう！'}
					></InputField>
				)}
			></Controller>

			<Controller
				name="place"
				control={control}
				rules={{ max: { value: 50, message: '場所は50文字以内で入力してください' } }}
				render={({ field, fieldState }) => (
					<InputField
						title="場所"
						type="text"
						field={field}
						fieldState={fieldState}
						helperText="場所は50文字以内で入力してください"
						icon={<RiMapPinLine />}
						autoComplete="off"
						placeholder={user.place || 'あなたの住んでいる場所を紹介しましょう！'}
					></InputField>
				)}
			></Controller>

			<div className="flex justify-center mt-4">
				<DPButton colormode="primary" type="submit">
					<TextWithIcon icon={<RiUserLine />}>変更する</TextWithIcon>
				</DPButton>
			</div>
		</form>
	)
}
