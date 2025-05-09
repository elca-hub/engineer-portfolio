'use server'

import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import { UserType } from '@/action/type/user'
import fetchExternalServiceUrlByUserId from '@/action/usecase/externalServiceUrl/fetchByUserId'
import fetchBio from '@/action/usecase/user/fetchBio'
import fetchById from '@/action/usecase/user/fetchById'
import HeaderPresentation from '@/app/_containers/profile/headerPresentation'
import ProfilePresentation from '@/app/_containers/profile/presentation'
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
	let externalServiceUrls: ExternalServiceUrlType[] = []
	let bio: string | null = null

	if (fetchUser) {
		user = fetchUser
		bio = await fetchBio(user)

		const externalServiceUrlsRes = await fetchExternalServiceUrlByUserId(userId)

		if (externalServiceUrlsRes) {
			externalServiceUrls = externalServiceUrlsRes
		}
	}

	if (!user) redirect('/404')

	const authRes = await getAuthUser()

	if (authRes.data) {
		authUser = authRes.data.user
	}

	return (
		<>
			<ProfilePresentation
				header={<HeaderPresentation user={authUser ?? undefined} isLogin={!!authUser}></HeaderPresentation>}
				user={user}
				isAuthUser={authUser !== null && authUser.user_id === userId}
				externalServiceUrls={externalServiceUrls}
				bio={bio}
			></ProfilePresentation>
		</>
	)
}
