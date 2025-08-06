'use server'

import { CertificationType } from '@/action/type/certification'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse } from '@/lib/api'

type CreateCertificationParam = {
	name: string
	year: string
	comment: string
}

export default async function createCertification(
	token: string,
	params: CreateCertificationParam,
): Promise<
	DPResponseData<{
		certification: CertificationType
	}>
> {
	const apiRes = await fetch(`${apiPrefix}/auth/user/certifications/`, {
		method: 'POST',
		body: JSON.stringify(params),
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	return await NewDPResponse<{ certification: CertificationType }>(apiRes)
}
