'use server'

import { UserType } from '@/action/type/user'
import { apiPrefix } from '@/constants/constant'
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

	return user
}
