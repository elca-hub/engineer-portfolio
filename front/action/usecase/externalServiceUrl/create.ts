'use server'

import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import { apiPrefix } from '@/constants/constant'
import { NewDPResponse } from '@/lib/api'

export default async function createExternalServiceUrl(
	token: string,
	{ externalServiceUrl }: { externalServiceUrl?: Omit<ExternalServiceUrlType, 'id'> },
): Promise<ExternalServiceUrlType | null> {
	const apiRes = await fetch(`${apiPrefix}/auth/user/external-service-url/`, {
		method: 'POST',
		body: JSON.stringify({ external_service_url: externalServiceUrl }),
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	const res = await NewDPResponse<{ external_service_url: ExternalServiceUrlType }>(apiRes)

	if (res.errors) {
		return null
	}

	return res.data?.external_service_url ?? null
}
