'use server'

import { SkillType } from '@/action/type/skill'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse } from '@/lib/api'

export default async function fetchByUserId(userId: string): Promise<
	DPResponseData<{
		skills: SkillType[]
	}>
> {
	const fetchUser = await fetch(`${apiPrefix}/user/${userId}/skills/`, {
		method: 'GET',
	})

	console.log(fetchUser)

	return await NewDPResponse<{ skills: SkillType[] }>(fetchUser)
}
