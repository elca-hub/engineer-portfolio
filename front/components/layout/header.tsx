'use client'

import { CalloutContext } from '@/app/state'
import { Pacifico } from 'next/font/google'
import Image from 'next/image'
import { useContext, useEffect, useState } from 'react'
import { Input, Label, SearchField } from 'react-aria-components'
import { Controller, useForm } from 'react-hook-form'
import { RiUserSearchLine } from 'react-icons/ri'

type Props = {
	children?: React.ReactNode
	userIconPath: string
}

const pacifico = Pacifico({
	weight: '400',
	subsets: ['latin'],
})

export type SearchUserFormContent = {
	name: string
}

export default function DPHeader({ children, userIconPath }: Props) {
	const { callout, setCallout } = useContext(CalloutContext)

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

	const [isSubmit, setIsSubmit] = useState(false)

	useEffect(() => {
		if (isSubmit) {
			const searchUser = async () => {
				console.log('searchUser')
				setCallout([...callout, { type: 'info', content: watch().name }])
			}

			searchUser()

			setIsSubmit(false)
		}
	}, [isSubmit])

	return (
		<header className="flex items-center justify-between px-4 py-2 bg-white border-b border-gray-200 shadow-sm">
			<div className="flex items-center gap-4">
				<Image src="/logo.webp" alt="logo" width={40} height={40} />
				<p className={`${pacifico.className} text-2xl`}>
					<span className="text-primary">Dev</span>
					<span className="text-secondary">Port</span>
				</p>
			</div>

			<div className="h-full">
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
							<SearchField className="relative">
								<Label className="sr-only">ユーザを検索</Label>
								<RiUserSearchLine className="absolute top-1/2 -translate-y-1/2 text-xl text-gray-600 left-2" />
								<Input
									type="text"
									{...field}
									placeholder="ユーザを検索"
									className="w-full p-2 pl-9 bg-gray-200 text-gray-600 rounded placeholder-gray-600 transition duration-200 ease-in-out focus:border-primary focus:outline-none focus:ring-2 focus:ring-primary"
								/>
							</SearchField>
						)}
					></Controller>
				</form>
			</div>

			<div className="flex items-center gap-4">
				{children}
				<Image src={userIconPath} alt="userIcon" width={46} height={46} />
			</div>
		</header>
	)
}
