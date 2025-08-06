'use server'

import { CertificationType } from '@/action/type/certification'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse, validateUserId } from '@/lib/api'

export default async function fetchCertificationByUserId(userId: string): Promise<
	DPResponseData<{
		certifications: CertificationType[]
	}>
> {
	if (!validateUserId(userId)) throw new Error('ユーザIDが異なります。再度ログインしてください')

	const fetchUser = await fetch(`${apiPrefix}/user/${userId}/certifications/`, {
		method: 'GET',
	})

	return await NewDPResponse<{ certifications: CertificationType[] }>(fetchUser)
}
