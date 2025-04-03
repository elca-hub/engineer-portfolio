'use server'

import HeaderPresentation from '@/app/_containers/profile/headerPresentation'
import ProfilePresentation from '@/app/_containers/profile/presentation'
import HeadContent from '@/components/layout/headContent'
import { apiPrefix, defaultUserIcon } from '@/constants/constant'
import { getSessionToken, handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'

export default async function ProfileContainer() {
	const res = await handleAuthRedirect('/profile')

	if (res.isRedirect) redirect(res.redirectPath)

	const token = await getSessionToken()
	if (!token) {
		redirect('/login')
	}

	const fetchUserInfo = await fetch(`${apiPrefix}/auth/user`, {
		method: 'GET',
		headers: {
			Authorization: `Bearer ${token}`,
		},
	})

	const userInfoData = await fetchUserInfo.json()

	let userInfo = {
		userName: '',
		iconPath: '',
		headerPath: '',
		bioPath: '',
		skillsLength: 0,
		externalLinksLength: 0,
		isFetch: false,
		isDone: false,
	}

	if (fetchUserInfo.ok) {
		userInfo = {
			userName: userInfoData.name,
			iconPath: userInfoData.icon_path === '' ? defaultUserIcon : userInfoData.icon_path,
			headerPath: userInfoData.header_path,
			bioPath: userInfoData.bio_path,
			skillsLength: userInfoData.skills.length,
			externalLinksLength: userInfoData.external_service_url.length,
			isFetch: true,
			isDone: true,
		}
	} else {
		userInfo = {
			userName: '',
			iconPath: defaultUserIcon,
			headerPath: '',
			bioPath: '',
			skillsLength: 0,
			externalLinksLength: 0,
			isFetch: false,
			isDone: true,
		}
	}

	return (
		<>
			<HeadContent
				title="ダッシュボード"
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<ProfilePresentation
				header={<HeaderPresentation userIconPath={userInfo.iconPath}></HeaderPresentation>}
				userInfo={userInfo}
			></ProfilePresentation>
		</>
	)
}
