'use server'

import HeaderPresentation from '@/app/_containers/profile/headerPresentation'
import ProfilePresentation from '@/app/_containers/profile/presentation'
import HeadContent from '@/components/layout/headContent'
import { apiPrefix, defaultUserIcon } from '@/constants/constant'
import { getSessionToken } from '@/lib/access'

type Props = {
	userId: string
}

export default async function ProfileContainer({ userId }: Props) {
	const fetchUserInfo = await fetch(`${apiPrefix}/user/${userId}/`, {
		method: 'GET',
	})

	let authUserId: string
	const token = await getSessionToken()

	if (token) {
		const authUserInfo = await fetch(`${apiPrefix}/auth/user`, {
			method: 'GET',
			headers: {
				Authorization: `Bearer ${token}`,
			},
		})

		if (authUserInfo.ok) {
			const authUserInfoData = await authUserInfo.json()
			authUserId = authUserInfoData.user_id
		} else {
			authUserId = ''
		}
	} else {
		authUserId = ''
	}

	const userInfoData = await fetchUserInfo.json()

	let userInfo = {
		userName: '',
		userId: userId,
		iconName: '',
		headerPath: '',
		bioPath: '',
		skillsLength: 0,
		externalLinksLength: 0,
		isFetch: false,
		isDone: false,
	}

	if (fetchUserInfo.ok) {
		userInfo = {
			userName: userInfoData.user.name,
			userId: userInfoData.user.user_id,
			iconName: userInfoData.user.icon_name,
			headerPath: userInfoData.user.header_path,
			bioPath: userInfoData.user.bio_path,
			skillsLength: userInfoData.user.skills.length,
			externalLinksLength: userInfoData.user.external_service_url.length,
			isFetch: true,
			isDone: true,
		}
	} else {
		userInfo = {
			userName: '',
			userId: userId,
			iconName: defaultUserIcon,
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
				header={<HeaderPresentation userIconName={userInfo.iconName} isLogin={!!token}></HeaderPresentation>}
				userInfo={userInfo}
				isAuthUser={authUserId === userId}
			></ProfilePresentation>
		</>
	)
}
