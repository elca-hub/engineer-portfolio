'use server'

import HeaderPresentation from '@/app/_containers/dashboard/headerPresentation'
import DashboardPresentation from '@/app/_containers/dashboard/presentation'
import HeadContent from '@/components/layout/headContent'
import { handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'

export default async function DashboardContainer() {
	const res = await handleAuthRedirect('/dashboard')

	if (res.isRedirect) redirect(res.redirectPath)

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
