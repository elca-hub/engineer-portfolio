'use server'

import ProfileBasicSettingPresentation from '@/app/_containers/account-setting/profile/basic/basicPresentation'
import TextWithIcon from '@/components/ui/text/textWithIcon'
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
			<h1 className="text-3xl font-bold tracking-wide text-foreground">
				<TextWithIcon icon={<RiUserLine />}>基本設定</TextWithIcon>
			</h1>
			<ProfileBasicSettingPresentation user={userData}></ProfileBasicSettingPresentation>
		</>
	)
}
