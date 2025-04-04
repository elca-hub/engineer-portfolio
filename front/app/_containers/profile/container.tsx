'use server'

import ProfilePresentation from '@/app/_containers/profile/presentation'
import HeadContent from '@/components/layout/headContent'
import DPHeader from '@/components/layout/header'
import { apiPrefix } from '@/constants/constant'
import { getAuthUser } from '@/lib/access'
import { NewDPResponse } from '@/lib/api'
import { UserType } from '@/lib/model/user'
import { redirect } from 'next/navigation'

type Props = {
	userId: string
}

export default async function ProfileContainer({ userId }: Props) {
	const fetchUser = await fetch(`${apiPrefix}/user/${userId}/`, {
		method: 'GET',
	})

	const fetchUserRes = await NewDPResponse<{ user: UserType }>(fetchUser)

	let authUser: UserType | null = null
	let user: UserType | null = null

	if (fetchUserRes.data) {
		user = fetchUserRes.data.user
	}

	if (!user) redirect('/404')

	const authRes = await getAuthUser()

	if (authRes.data) {
		authUser = authRes.data.user
	}

	return (
		<>
			<HeadContent
				title="ダッシュボード"
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<ProfilePresentation
				header={<DPHeader userIconName={authUser === null ? '' : authUser.icon_name} isLogin={!!authUser}></DPHeader>}
				user={user}
				isAuthUser={authUser !== null && authUser.user_id === userId}
			></ProfilePresentation>
		</>
	)
}
