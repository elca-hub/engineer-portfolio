'use server'

import ProfileIntroductionSettingPresentation from '@/app/_containers/account-setting/profile/introduction/introductionPresentation'
import TextWithIcon from '@/components/ui/text/textWithIcon'
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
			<h1 className="text-3xl font-bold tracking-wide text-foreground">
				<TextWithIcon icon={<RiPencilLine />}>プロフィール設定</TextWithIcon>
			</h1>
			<ProfileIntroductionSettingPresentation user={userData}></ProfileIntroductionSettingPresentation>
		</>
	)
}
