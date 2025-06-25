'use server'

import fetchExternalServiceUrlByUserId from '@/action/usecase/externalServiceUrl/fetchByUserId'
import ProfileLinkSettingPresentation from '@/app/_containers/account-setting/profile/link/linkPresentation'
import SettingPageMainTitle from '@/components/ui/text/setting-page/mainTitle'
import { getAuthUser, handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'
import { RiLink } from 'react-icons/ri'

export default async function ProfileLinkSettingContainer() {
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
			<SettingPageMainTitle icon={<RiLink />}>リンク設定</SettingPageMainTitle>
			<ProfileLinkSettingPresentation externalServiceUrls={externalServiceUrls}></ProfileLinkSettingPresentation>
		</>
	)
}
