'use server'

import { isNewUser } from '@/app/_containers/dashboard/action'
import HeaderPresentation from '@/app/_containers/dashboard/headerPresentation'
import DashboardPresentation from '@/app/_containers/dashboard/presentation'
import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import HeadContent from '@/components/layout/headContent'
import { getServerSession } from 'next-auth'
import { redirect } from 'next/navigation'

export default async function DashboardContainer() {
	const session = await getServerSession(authOptions)

	if (!session) {
		redirect('/login')
	}

	const sessionUser = session.user

	if (!sessionUser) {
		redirect('/login')
	}

	const isNew = await isNewUser(sessionUser.email ?? '')

	if (isNew) {
		return (
			<>
				<HeadContent
					title="新規登録"
					des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
				/>
			</>
		)
	} else {
		return (
			<>
				<HeadContent
					title="ダッシュボード"
					des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
				/>
				<DashboardPresentation header={<HeaderPresentation userIconPath="/dummy.png"></HeaderPresentation>} userName="test"></DashboardPresentation>
			</>
		)
	}
}
