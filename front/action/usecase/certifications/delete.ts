'use server'

import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse, validateUUID } from '@/lib/api'

export default async function deleteCertification(token: string, certificationId: string): Promise<DPResponseData<null>> {
	if (!validateUUID(certificationId)) throw new Error('IDが不正です。')

	const apiRes = await fetch(`${apiPrefix}/auth/user/certifications/${certificationId}/`, {
		method: 'DELETE',
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	return await NewDPResponse<null>(apiRes)
}
