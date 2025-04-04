'use server'

import HomePresentation from '@/app/_containers/home/presentation'
import DPHeader from '@/components/layout/header'
import { getAuthUser } from '@/lib/access'
import { UserType } from '@/lib/model/user'

export default async function HomeContainer() {
	let authUser: UserType | null = null

	const authRes = await getAuthUser()

	if (authRes.data) {
		authUser = authRes.data.user
	}

	return (
		<>
			<HomePresentation
				header={<DPHeader userIconName={authUser === null ? '' : authUser.icon_name} isLogin={!!authUser}></DPHeader>}
			></HomePresentation>
		</>
	)
}
