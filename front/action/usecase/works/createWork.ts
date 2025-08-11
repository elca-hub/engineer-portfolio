'use server'

import { WorkType } from '@/action/type/work'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse } from '@/lib/api'

export default async function createWork(
	token: string,
): Promise<
	DPResponseData<{
		work: WorkType
	}>
> {
	const apiRes = await fetch(`${apiPrefix}/auth/user/works/`, {
		method: 'POST',
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	return await NewDPResponse<{ work: WorkType }>(apiRes)
}
