'use server'

import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse } from '@/lib/api'

export default async function uploadBioImage(
	token: string,
	image: File,
): Promise<
	DPResponseData<{
		image_url: string
	}>
> {
	const formData = new FormData()
	formData.append('bio_image', image)
	const updateBio = await fetch(`${apiPrefix}/auth/user/bio/image`, {
		method: 'POST',
		body: formData,
		headers: {
			Authorization: `Bearer ${token}`,
		},
	})

	return await NewDPResponse<{ image_url: string }>(updateBio)
}
