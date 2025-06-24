'use client'

import { UserType } from '@/action/type/user'
import { CalloutContext } from '@/app/state'
import DBSidebar from '@/components/layout/sidebar/sidebar'
import DPButton from '@/components/ui/button/button'
import UserImage from '@/components/ui/image/userImage'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { logoutFlow } from '@/lib/access'
import { signIn, signOut } from 'next-auth/react'
import { Pacifico } from 'next/font/google'
import Image from 'next/image'
import { useContext, useEffect, useState } from 'react'
import { Button, Input, Label, Link, Menu, MenuItem, MenuItemProps, MenuTrigger, Popover, SearchField } from 'react-aria-components'
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

	return (
		<header className="z-50 flex items-center justify-between px-4 py-2 bg-white border-b border-gray-200 shadow-sm">
			{sidebar ? (
				<div className="flex items-center gap-4">
					<button className="p-2 rounded shadow hover:bg-gray-300 transition-all duration-300 border" onClick={() => setSidebarOpen(true)}>
						<RiMenuLine className="size-5" />
					</button>
					<DBSidebar isOpen={isSidebarOpen} onClose={() => setSidebarOpen(false)}>
						{sidebar}
					</DBSidebar>
					<Link href="/" className="outline-none">
						<Image src="/logo.webp" alt="logo" width="40" height="40" priority />
					</Link>
				</div>
			) : (
				<>
					<Link href="/" className="outline-none">
						<div className="flex items-center gap-4">
							<Image src="/logo.webp" alt="logo" width={40} height={40} priority />
							<p className={`${pacifico.className} hidden md:block md:text-2xl`}>
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

			<div className="flex items-center gap-4">
				{isLogin && user ? (
					<>
						{children ? (
							children
						) : user ? (
							<Link href={`/${user.user_id}/profile`} className="outline-none">
								<DPButton colormode="primary">
									<TextWithIcon icon={<RiProfileLine />}>プロフィールへ</TextWithIcon>
								</DPButton>
							</Link>
						) : (
							<></>
						)}
						<MenuTrigger>
							<Button
								aria-label="menu"
								className="rounded-xl inline-flex items-center justify-center text-white bg-transparent border-none hover:scale-95 transition-all cursor-pointer outline-none focus-visible:ring-2 focus-visible:ring-primary"
							>
								<UserImage imageType="icon" fileName={user.icon_name} className="w-12 h-12 rounded-xl object-cover" width={47} height={47} priority />
							</Button>
							<Popover className="outline-hidden overflow-auto bg-background p-2 rounded-lg bg-white shadow-lg ring-2 ring-primary entering:animate-in entering:fade-in entering:placement-bottom:slide-in-from-top-1 entering:placement-top:slide-in-from-bottom-1 exiting:animate-out exiting:fade-out exiting:placement-bottom:slide-out-to-top-1 exiting:placement-top:slide-out-to-bottom-1 fill-mode-forwards origin-top-left">
								<Menu className="outline-none">
									{user && (
										<MyMenuItem id="user-setting">
											<Link href="/account/setting">
												<TextWithIcon icon={<RiUserLine />}>ユーザ設定</TextWithIcon>
											</Link>
										</MyMenuItem>
									)}
									<MyMenuItem id="signout" onAction={() => setIsLogout(true)}>
										<TextWithIcon icon={<RiLockLine />}>ログアウト</TextWithIcon>
									</MyMenuItem>
								</Menu>
							</Popover>
						</MenuTrigger>
					</>
				) : (
					<DPButton colormode="primary" onPress={() => signIn('google')}>
						<TextWithIcon icon={<RiGoogleFill />}>ログイン</TextWithIcon>
					</DPButton>
				)}
			</div>
		</header>
	)
}

function MyMenuItem(props: MenuItemProps) {
	return (
		<MenuItem
			{...props}
			className="cursor-pointer group flex w-full items-center rounded-md px-3 py-2 box-border outline-none cursor-default text-foreground focus:bg-primary focus:text-white"
		/>
	)
}
