'use server'

import { UserType } from '@/action/type/user'
import HomePresentation from '@/app/_containers/home/presentation'
import DPHeader from '@/components/layout/header'
import { getAuthUser, handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'

export default async function HomeContainer() {
	const authRedirectPath = await handleAuthRedirect('public')
	if (authRedirectPath) redirect(authRedirectPath)

	let authUser: UserType | null = null

	const authRes = await getAuthUser()

	if (authRes.data) {
		authUser = authRes.data.user
	}

	return (
		<>
			<HomePresentation header={<DPHeader user={authUser ?? undefined} isLogin={!!authUser}></DPHeader>}></HomePresentation>
		</>
	)
}
