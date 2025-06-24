'use server'

import fetchExternalServiceUrlByUserId from '@/action/usecase/externalServiceUrl/fetchByUserId'
import ProfileSettingPresentation from '@/app/_containers/account-setting/profile/presentation'
import { getAuthUser, handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'

export default async function ProfileSettingContainer() {
	const authRedirectPath = await handleAuthRedirect('auth')
	if (authRedirectPath) redirect(authRedirectPath)

	const authUser = await getAuthUser()

	if (authUser.errors || !authUser.data) {
		redirect('/')
	}

	const userData = authUser.data.user

	const externalServiceUrls = await fetchExternalServiceUrlByUserId(userData.user_id)

	if (!externalServiceUrls) {
		redirect('/404')
	}

	return (
		<>
			<ProfileSettingPresentation user={userData} externalServiceUrls={externalServiceUrls}></ProfileSettingPresentation>
		</>
	)
}
