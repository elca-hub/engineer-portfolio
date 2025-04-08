'use server'

import { UserType } from '@/action/type/user'
import fetchById from '@/action/usecase/user/fetchById'
import ProfilePresentation from '@/app/_containers/profile/presentation'
import DPHeader from '@/components/layout/header'
import { getAuthUser, handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'

type Props = {
	userId: string
}

export default async function ProfileContainer({ userId }: Props) {
	const authRedirectPath = await handleAuthRedirect('public')
	if (authRedirectPath) redirect(authRedirectPath)

	const fetchUser = await fetchById(userId)

	let authUser: UserType | null = null
	let user: UserType | null = null

	if (fetchUser) {
		user = fetchUser
	}

	if (!user) redirect('/404')

	const authRes = await getAuthUser()

	if (authRes.data) {
		authUser = authRes.data.user
	}

	return (
		<>
			<ProfilePresentation
				header={<DPHeader userIconName={authUser === null ? '' : authUser.icon_name} isLogin={!!authUser}></DPHeader>}
				user={user}
				isAuthUser={authUser !== null && authUser.user_id === userId}
			></ProfilePresentation>
		</>
	)
}
