'use server'

import HeadContent from '@/components/layout/headContent'
import { getAuthUser, handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'
import HeaderButtonProvider from './state'
import SettingHeader from './settingHeader'

type Props = {
	children: React.ReactNode
}

export default async function SettingLayout({ children }: Props) {
	const authRedirectPath = await handleAuthRedirect('auth')
	if (authRedirectPath) redirect(authRedirectPath)

	const user = await getAuthUser()

	if (user.errors || !user.data) {
		redirect('/')
	}

	return (
		<>
			<HeadContent
				title={'プロフィール編集'}
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<HeaderButtonProvider>
				<div className="min-h-screen">
					<SettingHeader user={user.data.user} />
					<div className="flex flex-row">
						<main
							className="w-full bg-white border-r border-border p-4 flex flex-col gap-2"
							style={{ minHeight: 'calc(100vh - var(--header-height, 64px))' }}
						>
							{children}
						</main>
					</div>
				</div>
			</HeaderButtonProvider>
		</>
	)
}
