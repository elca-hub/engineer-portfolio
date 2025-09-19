'use server'

import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse, validateUUID } from '@/lib/api'

export default async function uploadWorkThumbnail(token: string, workId: string, image: File): Promise<
  DPResponseData<{ image_url: string }>
> {

	if (!validateUUID(workId)) throw new Error('作品IDが不正です')

  const formData = new FormData()
	formData.append('work_thumbnail', image)

	return await NewDPResponse<{ image_url: string }>(
		await fetch(`${apiPrefix}/auth/user/works/${workId}/thumbnail`, {
			method: 'POST',
			body: formData,
			headers: {
				Authorization: `Bearer ${token}`,
			},
		}),
	)
}
