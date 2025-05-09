'use server'

import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse } from '@/lib/api'

export default async function uploadBio(
	token: string,
	bio: string,
): Promise<
	DPResponseData<{
		bio: string
	}>
> {
	const updateBio = await fetch(`${apiPrefix}/auth/user/bio/`, {
		method: 'POST',
		body: JSON.stringify({ bio: bio }),
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	return await NewDPResponse<{ bio: string }>(updateBio)
}
