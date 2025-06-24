'use server'

import ProfileIntroductionSettingPresentation from '@/app/_containers/account-setting/profile/introduction/introductionPresentation'
import SettingPageMainTitle from '@/components/ui/text/setting-page/mainTitle'
import { getAuthUser, handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'
import { RiPencilLine } from 'react-icons/ri'

export default async function ProfileIntroductionSettingContainer() {
	const authRedirectPath = await handleAuthRedirect('auth')
	if (authRedirectPath) redirect(authRedirectPath)

	const authUser = await getAuthUser()

	if (authUser.errors || !authUser.data) {
		redirect('/')
	}

	const userData = authUser.data.user

	return (
		<>
			<SettingPageMainTitle icon={<RiPencilLine />}>自己紹介設定</SettingPageMainTitle>
			<ProfileIntroductionSettingPresentation user={userData}></ProfileIntroductionSettingPresentation>
		</>
	)
}
