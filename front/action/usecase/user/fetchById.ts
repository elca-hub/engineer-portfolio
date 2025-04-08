'use server'

import { UserType } from '@/action/type/user'
import { apiPrefix, defaultUserHeaderImage, defaultUserIcon } from '@/constants/constant'
import { NewDPResponse } from '@/lib/api'

export default async function fetchById(userId: string): Promise<UserType | null> {
	const fetchUser = await fetch(`${apiPrefix}/user/${userId}/`, {
		method: 'GET',
	})

	const fetchUserRes = await NewDPResponse<{ user: UserType }>(fetchUser)

	if (fetchUserRes.errors) {
		return null
	}

	const user = fetchUserRes.data?.user ?? null

	if (!user) {
		return null
	}

	if (user.icon_name === '') {
		user.icon_name = defaultUserIcon
	}

	console.log(user.header_icon_name)

	if (user.header_icon_name === '') {
		user.header_icon_name = defaultUserHeaderImage
	}

	return user
}
