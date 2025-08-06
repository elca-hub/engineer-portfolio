'use server'

import SkillsSettingPresentation from '@/app/_containers/account-setting/skills-and-certifications/skills/presentation'
import SettingPageMainTitle from '@/components/ui/text/setting-page/mainTitle'
import { getAuthUser, handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'
import { RiMedal2Line } from 'react-icons/ri'

export default async function SkillsSettingContainer() {
	const authRedirectPath = await handleAuthRedirect('auth')
	if (authRedirectPath) redirect(authRedirectPath)

	const authUser = await getAuthUser()

	if (authUser.errors || !authUser.data) {
		redirect('/')
	}

	const userData = authUser.data.user

	return (
		<>
			<SettingPageMainTitle icon={<RiMedal2Line />}>スキルの管理</SettingPageMainTitle>
			<SkillsSettingPresentation user={userData}></SkillsSettingPresentation>
		</>
	)
}
