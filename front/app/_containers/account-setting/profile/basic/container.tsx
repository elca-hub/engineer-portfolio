'use server'

import ProfileBasicSettingPresentation from '@/app/_containers/account-setting/profile/basic/basicPresentation'
import SettingPageMainTitle from '@/components/ui/text/setting-page/mainTitle'
import { getAuthUser, handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'
import { RiUserLine } from 'react-icons/ri'

export default async function ProfileBasicSettingContainer() {
	const authRedirectPath = await handleAuthRedirect('auth')
	if (authRedirectPath) redirect(authRedirectPath)

	const authUser = await getAuthUser()

	if (authUser.errors || !authUser.data) {
		redirect('/')
	}

	const userData = authUser.data.user

	return (
		<>
			<SettingPageMainTitle icon={<RiUserLine />}>基本設定</SettingPageMainTitle>
			<ProfileBasicSettingPresentation user={userData}></ProfileBasicSettingPresentation>
		</>
	)
}
