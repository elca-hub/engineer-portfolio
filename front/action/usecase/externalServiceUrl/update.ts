'use server'

import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import { apiPrefix } from '@/constants/constant'
import { NewDPResponse } from '@/lib/api'

export default async function updateExternalServiceUrl(
	token: string,
	externalServiceUrl: ExternalServiceUrlType,
): Promise<ExternalServiceUrlType | null> {
	const updateUser = await fetch(`${apiPrefix}/auth/user/external-service-url/${externalServiceUrl.id}`, {
		method: 'PUT',
		body: JSON.stringify({url: externalServiceUrl.url}),
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	const updateUserRes = await NewDPResponse<{ external_service_url: ExternalServiceUrlType }>(updateUser)

	if (updateUserRes.errors) {
		return null
	}

	return updateUserRes.data?.external_service_url ?? null
}
