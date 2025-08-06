'use server'

import { SkillType } from '@/action/type/skill'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse } from '@/lib/api'

type CreateSkillParams = {
	name: string
	year: string
	comment: string
}

export default async function createSkill(
	token: string,
	params: CreateSkillParams,
): Promise<
	DPResponseData<{
		skill: SkillType
	}>
> {
	const apiRes = await fetch(`${apiPrefix}/auth/user/skills/`, {
		method: 'POST',
		body: JSON.stringify(params),
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	return await NewDPResponse<{ skill: SkillType }>(apiRes)
}
