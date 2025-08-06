'use server'

import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse } from '@/lib/api'

type updateSortParams = {
	sort_list: {
		id: string
		sort_index: number
	}[]
}

export default async function updateSkillsSort(token: string, params: updateSortParams): Promise<DPResponseData<null>> {
	const apiRes = await fetch(`${apiPrefix}/auth/user/skills/sort`, {
		method: 'PUT',
		body: JSON.stringify(params),
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	return await NewDPResponse<null>(apiRes)
}
