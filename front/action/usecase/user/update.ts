'use server'

import { UserType } from '@/action/type/user'
import { apiPrefix } from '@/constants/constant'
import { NewDPResponse } from '@/lib/api'

export default async function userUpdate(token: string, { userData }: { userData?: Partial<UserType> }): Promise<UserType | null> {
	const updateUser = await fetch(`${apiPrefix}/auth/user/`, {
		method: 'PUT',
		body: JSON.stringify({ user: userData }),
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	const updateUserRes = await NewDPResponse<{ user: UserType }>(updateUser)

	if (updateUserRes.errors) {
		return null
	}

	return updateUserRes.data?.user ?? null
}
