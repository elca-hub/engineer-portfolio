'use server'

import CertificationsSettingPresentation from '@/app/_containers/account-setting/skills-and-certifications/certifications/presentation'
import SettingPageMainTitle from '@/components/ui/text/setting-page/mainTitle'
import { getAuthUser, handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'
import { RiMedalLine } from 'react-icons/ri'

export default async function CertificationsSettingContainer() {
	const authRedirectPath = await handleAuthRedirect('auth')
	if (authRedirectPath) redirect(authRedirectPath)

	const authUser = await getAuthUser()

	if (authUser.errors || !authUser.data) {
		redirect('/')
	}

	const userData = authUser.data.user

	return (
		<>
			<SettingPageMainTitle icon={<RiMedalLine />}>資格の管理</SettingPageMainTitle>
			<CertificationsSettingPresentation user={userData}></CertificationsSettingPresentation>
		</>
	)
}
