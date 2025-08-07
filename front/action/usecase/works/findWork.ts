'use server'

import { WorkType } from '@/action/type/work'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse, validateUserId, validateUUID } from '@/lib/api'

export default async function findWork(userId: string, workId: string): Promise<
  DPResponseData<{ work: WorkType }>
> {
	if (!validateUserId(userId)) throw new Error('ユーザIDが異なります。再度ログインしてください')

	if (!validateUUID(workId)) throw new Error('作品IDが不正です')

	return await NewDPResponse<{ work: WorkType }>(
		await fetch(`${apiPrefix}/user/${userId}/works/${workId}`, {
			method: 'GET',
		}),
	)
}
