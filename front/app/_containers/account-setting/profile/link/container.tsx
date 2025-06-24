'use server'

import fetchExternalServiceUrlByUserId from '@/action/usecase/externalServiceUrl/fetchByUserId'
import ProfileLinkSettingPresentation from '@/app/_containers/account-setting/profile/link/linkPreesntation'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getAuthUser, handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'
import { RiPencilLine } from 'react-icons/ri'

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
			<h1 className="text-3xl font-bold tracking-wide text-foreground">
				<TextWithIcon icon={<RiPencilLine />}>プロフィール設定</TextWithIcon>
			</h1>
			<ProfileLinkSettingPresentation externalServiceUrls={externalServiceUrls}></ProfileLinkSettingPresentation>
		</>
	)
}
