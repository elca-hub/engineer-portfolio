'use client'

import { UserType } from '@/action/type/user'
import userUpdate from '@/action/usecase/user/update'
import { CalloutContext } from '@/app/state'
import DatePickerField from '@/components/layout/input/datePickerField'
import InputField from '@/components/layout/input/inputField'
import DPModal from '@/components/layout/modal'
import DPButton from '@/components/ui/button/button'
import UserImage from '@/components/ui/image/userImage'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getSessionToken } from '@/lib/access'
import { CalendarDate, getLocalTimeZone, parseDate, today } from '@internationalized/date'
import { useRouter } from 'next/navigation'
import { useContext, useEffect, useState } from 'react'
import { Button, DialogTrigger, DropZone, FileTrigger } from 'react-aria-components'
import { Controller, useForm } from 'react-hook-form'
import { RiCake2Line, RiImageAddLine, RiUserLine } from 'react-icons/ri'

type Props = {
	header: React.ReactNode
	user: UserType
	isAuthUser: boolean
}

export type UserUpdateFormType = {
	name: string
	birthday: CalendarDate
}

function UserIconComponent({
	user,
	type,
	setIconFile,
}: {
	user: UserType
	type: 'icon' | 'header'
	setIconFile: (file: { file: File; type: 'icon' | 'header' }) => void
}) {
	const iconClass = type === 'icon' ? 'size-[6rem] md:size-40' : 'w-full md:h-40 h-[8rem]'

	const iconOnlyClass = type === 'icon' ? 'absolute -bottom-10 left-4 z-10' : 'relative'

	return (
		<div className={`${iconOnlyClass}`}>
			<DropZone
				onDrop={(e) => {
					const targetFile = e.items[0]
					if (targetFile.kind !== 'file') return
					targetFile.getFile().then((f) => {
						setIconFile({ file: f, type: type })
					})
				}}
				className={`mx-auto relative rounded-xl data-[drop-target]:ring-2 data-[drop-target]:ring-primary data-[drop-target]:ring-offset-2 data-[drop-target]:ring-offset-background`}
			>
				<UserImage
					imageType={type}
					fileName={type === 'icon' ? user.icon_name : user.header_icon_name}
					width={type === 'icon' ? 500 : 1000}
					height={type === 'icon' ? 500 : 1000}
					className={`${iconClass} ${type === 'icon' ? 'border-2 border-white shadow-md rounded-xl object-cover' : ''} rounded-xl object-cover brightness-50`}
				/>
				<FileTrigger
					onSelect={(e) => {
						if (!e) return
						const file = e.item(0)
						if (file === null) return
						setIconFile({ file: file, type: type })
					}}
					acceptedFileTypes={['image/*']}
				>
					<Button
						className={`absolute top-1/2 ${type === 'icon' ? 'left-1/2 -translate-x-1/2' : 'right-1/3'} -translate-y-1/2 rounded-full bg-white/80 p-3 shadow-md hover:scale-95 transition-all duration-200`}
					>
						<RiImageAddLine className="text-xl md:text-2xl" />
					</Button>
				</FileTrigger>
			</DropZone>
		</div>
	)
}

/**
 * @package
 */
export default function ProfilePresentation({ header, user, isAuthUser }: Props) {
	const router = useRouter()
	const { callout, setCallout } = useContext(CalloutContext)

	const [iconFile, setIconFile] = useState<{ file: File; type: 'icon' | 'header' }>()

	const { control, handleSubmit, watch } = useForm<UserUpdateFormType>({
		defaultValues: {
			name: user.name,
			birthday: parseDate(user.birthday),
		},
	})

	const [isSubmit, setIsSubmit] = useState(false)

	useEffect(() => {
		if (isSubmit) {
			const name = watch('name')
			const birthday = watch('birthday')
			const updateFlow = async () => {
				const token = await getSessionToken()
				if (!token) {
					setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
					return
				}

				const res = await userUpdate(token, { userData: { name, birthday: birthday.toString() } })
				if (res === null) {
					setCallout([...callout, { content: '変更に失敗しました', type: 'error' }])
					return
				}

				setCallout([...callout, { content: '変更しました', type: 'info' }])
				router.refresh()
			}

			updateFlow()

			setIsSubmit(false)
		}
	}, [isSubmit])

	useEffect(() => {
		if (iconFile) {
			const { file, type } = iconFile
			if (file.type !== 'image/png' && file.type !== 'image/jpeg') {
				setCallout([...callout, { content: 'PNGまたはJPEG形式の画像を選択してください', type: 'error' }])
				setIconFile(undefined)
				return
			}

			if (file.size > 1024 * 1024 * 50) {
				setCallout([...callout, { content: '画像のサイズは50MB以下にしてください', type: 'error' }])
				setIconFile(undefined)
				return
			}

			const iconUploadFlow = async () => {
				const token = await getSessionToken()
				if (!token) {
					setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
					return
				}

				let res: UserType | null = null

				if (type === 'icon') {
					res = await userUpdate(token, { icon: file })
				} else {
					res = await userUpdate(token, { headerIcon: file })
				}

				if (res === null) {
					setCallout([...callout, { content: 'アップロードに失敗しました', type: 'error' }])
					return
				}

				setCallout([...callout, { content: '変更しました', type: 'info' }])
				router.refresh()
			}

			iconUploadFlow()

			setIconFile(undefined)
		}
	}, [iconFile])

	return (
		<div className="flex flex-col h-screen">
			{header}
			<div className="relative">
				<div className="relative w-full h-60">
					<UserImage
						priority={true}
						imageType="header"
						fileName={user.header_icon_name}
						className="object-cover w-full h-full"
						width={1000}
						height={1000}
					/>
				</div>
				<div className="z-2 absolute -bottom-1/4 md:-bottom-1/2 left-2 md:left-16">
					<div className="flex items-end gap-4">
						<UserImage
							imageType="icon"
							fileName={user.icon_name}
							width="171"
							height="171"
							className="size-[120px] md:size-[170px] border-2 border-white shadow-md rounded-xl object-cover"
							priority
						/>
						<div className="flex flex-col mb-1 bg-white/70 rounded-lg p-2 md:bg-transparent">
							{isAuthUser && (
								<DialogTrigger>
									<DPButton colormode="primary">
										<p className="text-sm">
											<TextWithIcon icon={<RiUserLine />}>プロフィール編集</TextWithIcon>
										</p>
									</DPButton>
									<DPModal header={{ title: 'プロフィール編集', icon: <RiUserLine /> }}>
										<div className="relative mb-14">
											<UserIconComponent user={user} type="icon" setIconFile={setIconFile} />
											<UserIconComponent user={user} type="header" setIconFile={setIconFile} />
										</div>

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

											<div className="flex justify-center mt-4">
												<DPButton colormode="primary" type="submit">
													<TextWithIcon icon={<RiUserLine />}>変更する</TextWithIcon>
												</DPButton>
											</div>
										</form>
									</DPModal>
								</DialogTrigger>
							)}
							<h1 className="text-3xl font-bold tracking-wide text-foreground">{user.name}</h1>
							<p className="text-lg font-medium text-gray-600">@{user.user_id}</p>
						</div>
					</div>
				</div>
			</div>
			<main className="px-10 mt-2"></main>
		</div>
	)
}
