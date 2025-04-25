'use server'

import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse } from '@/lib/api'

export default async function updateExternalServiceUrl(
	token: string,
	externalServiceUrl: ExternalServiceUrlType,
): Promise<DPResponseData<{external_service_url: ExternalServiceUrlType}>> {
	const updateUser = await fetch(`${apiPrefix}/auth/user/external-service-url/${externalServiceUrl.id}/`, {
		method: 'PUT',
		body: JSON.stringify({url: externalServiceUrl.url}),
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	return await NewDPResponse<{ external_service_url: ExternalServiceUrlType }>(updateUser)
}
