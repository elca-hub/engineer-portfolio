'use server'

import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse, validateUUID } from '@/lib/api'

export default async function deleteSkill(token: string, skillId: string): Promise<DPResponseData<null>> {
	if (!validateUUID(skillId)) throw new Error('IDが不正です。')

	const apiRes = await fetch(`${apiPrefix}/auth/user/skills/${skillId}/`, {
		method: 'DELETE',
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	return await NewDPResponse<null>(apiRes)
}
