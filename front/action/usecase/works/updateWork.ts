'use server'

import { WorkType } from '@/action/type/work'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse, validateUserId, validateUUID } from '@/lib/api'

export default async function updateWork(token: string, userId: string, work: WorkType, isDeleteImage: boolean): Promise<
  DPResponseData<{ work: WorkType }>
> {
	if (!validateUserId(userId)) throw new Error('ユーザIDが異なります。再度ログインしてください')

	if (!validateUUID(work.id)) throw new Error('作品IDが不正です')

	return await NewDPResponse<{ work: WorkType }>(
		await fetch(`${apiPrefix}/auth/user/works/${work.id}?is_delete_image=${isDeleteImage}`, {
			method: 'PUT',
			body: JSON.stringify({ work: work }),
			headers: {
				Authorization: `Bearer ${token}`,
				'Content-Type': 'application/json',
			},
		}),
	)
}
