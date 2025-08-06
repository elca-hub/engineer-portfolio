'use server'

import { SkillType } from '@/action/type/skill'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse, validateUserId } from '@/lib/api'

export default async function fetchByUserId(userId: string): Promise<
	DPResponseData<{
		skills: SkillType[]
	}>
> {
	if (!validateUserId(userId)) throw new Error('ユーザIDが異なります。再度ログインしてください')

	const fetchUser = await fetch(`${apiPrefix}/user/${userId}/skills/`, {
		method: 'GET',
	})
	
	return await NewDPResponse<{ skills: SkillType[] }>(fetchUser)
}
