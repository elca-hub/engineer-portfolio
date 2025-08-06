'use server'

import { CertificationType } from '@/action/type/certification'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse, validateUUID } from '@/lib/api'

export default async function updateCertification(
	token: string,
	params: CertificationType,
): Promise<
	DPResponseData<{
		certification: CertificationType
	}>
> {
	if (!validateUUID(params.id)) throw new Error('IDが不正です。')

	const apiRes = await fetch(`${apiPrefix}/auth/user/certifications/${params.id}/`, {
		method: 'PUT',
		body: JSON.stringify(params),
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	return await NewDPResponse<{ certification: CertificationType }>(apiRes)
}
