'use server'

import { SkillType } from '@/action/type/skill'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse } from '@/lib/api'

export default async function updateSkill(
	token: string,
	params: SkillType,
): Promise<
	DPResponseData<{
		skill: SkillType
	}>
> {
	const apiRes = await fetch(`${apiPrefix}/auth/user/skills/${params.id}/`, {
		method: 'PUT',
		body: JSON.stringify(params),
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	return await NewDPResponse<{ skill: SkillType }>(apiRes)
}
