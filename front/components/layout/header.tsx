'use client'

import { UserType } from '@/action/type/user'
import { CalloutContext } from '@/app/state'
import LeftSidebar from '@/components/layout/sidebar/leftSidebar'
import RightSidebar from '@/components/layout/sidebar/rightSidebar'
import DPButton from '@/components/ui/button/button'
import UserImage from '@/components/ui/image/userImage'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { logoutFlow } from '@/lib/access'
import { signIn, signOut } from 'next-auth/react'
import { Pacifico } from 'next/font/google'
import Image from 'next/image'
import { useContext, useEffect, useState } from 'react'
import { Input, Label, Link, SearchField } from 'react-aria-components'
import { Controller, useForm } from 'react-hook-form'
import { RiGoogleFill, RiLockLine, RiMenuLine, RiProfileLine, RiUserLine, RiUserSearchLine } from 'react-icons/ri'

type Props = {
	children?: React.ReactNode
	user?: UserType
	isLogin: boolean
	sidebar?: React.ReactNode
}

const pacifico = Pacifico({
	weight: '400',
	subsets: ['latin'],
})

export type SearchUserFormContent = {
	name: string
}

export default function DPHeader({ children, user, isLogin, sidebar }: Props) {
	const { callout, setCallout } = useContext(CalloutContext)
	const [isLogout, setIsLogout] = useState(false)

	const {
		control,
		handleSubmit,
		watch,
		formState: { errors },
	} = useForm<SearchUserFormContent>({
		defaultValues: {
			name: '',
		},
	})

	// errorをwatch
	useEffect(() => {
		if (errors.name && errors.name.message) {
			setCallout([...callout, { type: 'error', content: errors.name.message }])
		}
	}, [errors])

	useEffect(() => {
		if (isLogout) {
			const logout = async () => {
				await logoutFlow()
			}

			logout()
			signOut()
			setIsLogout(false)
		}
	}, [isLogout])

	const [isSubmit, setIsSubmit] = useState(false)

	useEffect(() => {
		if (isSubmit) {
			const searchUser = async () => {
				setCallout([...callout, { type: 'info', content: watch().name }])
			}

			searchUser()

			setIsSubmit(false)
		}
	}, [isSubmit])

	const [isSidebarOpen, setSidebarOpen] = useState(false)
	const [isRightSidebarOpen, setRightSidebarOpen] = useState(false)

	return (
		<header className="z-50 flex items-center justify-between px-2 sm:px-4 py-2 bg-white border-b border-gray-200 shadow-sm">
			{sidebar ? (
				<div className="flex items-center gap-2 sm:gap-4">
					<button className="p-2 rounded shadow hover:bg-gray-300 transition-all duration-300 border" onClick={() => setSidebarOpen(true)}>
						<RiMenuLine className="size-4 sm:size-5" />
					</button>
					<LeftSidebar isOpen={isSidebarOpen} onClose={() => setSidebarOpen(false)}>
						{sidebar}
					</LeftSidebar>
					<Link href="/" className="outline-none">
						<Image src="/logo.webp" alt="logo" width={32} height={32} className="sm:w-10 sm:h-10" priority />
					</Link>
				</div>
			) : (
				<>
					<Link href="/" className="outline-none">
						<div className="flex items-center gap-2 sm:gap-4">
							<Image src="/logo.webp" alt="logo" width={32} height={32} className="sm:w-10 sm:h-10" priority />
							<p className={`${pacifico.className} hidden sm:block text-xl sm:text-2xl`}>
								<span className="text-primary">Dev</span>
								<span className="text-secondary">Port</span>
							</p>
						</div>
					</Link>
				</>
			)}

			<div className="h-full hidden md:block">
				<form onSubmit={handleSubmit(() => setIsSubmit(true))} className="h-full">
					<Controller
						name="name"
						control={control}
						rules={{
							required: '検索するユーザ名を入力してください',
							max: {
								value: 50,
								message: 'ユーザ名は50文字以内で入力してください',
							},
						}}
						render={({ field }) => (
							<SearchField className="relative h-full">
								<Label className="sr-only">ユーザを検索</Label>
								<RiUserSearchLine className="absolute top-1/2 -translate-y-1/2 text-xl text-gray-600 left-2" />
								<Input
									type="text"
									{...field}
									placeholder="ユーザを検索"
									className="h-full w-full p-2 pl-9 bg-gray-200 text-gray-600 rounded placeholder-gray-600 transition duration-200 ease-in-out focus:border-primary focus:outline-none focus:ring-2 focus:ring-primary"
								/>
							</SearchField>
						)}
					></Controller>
				</form>
			</div>

			<div className="flex items-center gap-1 sm:gap-2 md:gap-4">
				{isLogin && user ? (
					<>
						<div className="hidden sm:block">
							{children ? (
								children
							) : user ? (
								<Link href={`/${user.user_id}/profile`} className="outline-none">
									<DPButton colormode="primary">
										<TextWithIcon icon={<RiUserLine />}>プロフィールへ</TextWithIcon>
									</DPButton>
								</Link>
							) : (
								<></>
							)}
						</div>
						<button
							aria-label="user menu"
							className="rounded-xl inline-flex items-center justify-center text-white bg-transparent border-none hover:scale-95 transition-all cursor-pointer outline-none focus-visible:ring-2 focus-visible:ring-primary"
							onClick={() => setRightSidebarOpen(true)}
						>
							<UserImage
								imageType="icon"
								fileName={user.icon_name}
								className="w-8 h-8 sm:w-10 sm:h-10 md:w-12 md:h-12 rounded-xl object-cover"
								width={40}
								height={40}
								priority
							/>
						</button>
						<RightSidebar isOpen={isRightSidebarOpen} onClose={() => setRightSidebarOpen(false)}>
							<UserSidebarContent user={user} onLogout={() => setIsLogout(true)} />
						</RightSidebar>
					</>
				) : (
					<DPButton colormode="primary" onPress={() => signIn('google')}>
						<TextWithIcon icon={<RiGoogleFill />}>
							<span className="hidden sm:inline">ログイン</span>
							<span className="sm:hidden">ログイン</span>
						</TextWithIcon>
					</DPButton>
				)}
			</div>
		</header>
	)
}

function UserSidebarContent({ user, onLogout }: { user: UserType; onLogout: () => void }) {
	return (
		<>
			<div className="flex items-center gap-3 p-4 border-b border-gray-200">
				<UserImage imageType="icon" fileName={user.icon_name} className="w-12 h-12 rounded-xl object-cover" width={48} height={48} priority />
				<div>
					<p className="font-medium text-foreground">{user.name}</p>
					<p className="text-sm text-gray-600">@{user.user_id}</p>
				</div>
			</div>
			<div className="flex flex-col gap-2 p-2">
				<Link href={`/${user.user_id}/profile`} className="outline-none">
					<div className="flex items-center gap-3 p-3 rounded-md hover:bg-gray-100 transition-colors">
						<RiProfileLine className="size-5 text-gray-600" />
						<span className="text-foreground">プロフィール</span>
					</div>
				</Link>
				<Link href="/account/setting" className="outline-none">
					<div className="flex items-center gap-3 p-3 rounded-md hover:bg-gray-100 transition-colors">
						<RiUserLine className="size-5 text-gray-600" />
						<span className="text-foreground">ユーザ設定</span>
					</div>
				</Link>
				<button onClick={onLogout} className="flex items-center gap-3 p-3 rounded-md hover:bg-gray-100 transition-colors text-left w-full">
					<RiLockLine className="size-5 text-gray-600" />
					<span className="text-foreground">ログアウト</span>
				</button>
			</div>
		</>
	)
}
