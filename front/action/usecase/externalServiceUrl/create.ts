'use server'

import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse } from '@/lib/api'

export default async function createExternalServiceUrl(
	token: string,
	serviceType: number,
	url: string,
): Promise<DPResponseData<{ external_service_url: ExternalServiceUrlType }>> {
	const apiRes = await fetch(`${apiPrefix}/auth/user/external-service-url/`, {
		method: 'POST',
		body: JSON.stringify({
			service_type: serviceType,
			url: url,
		}),
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	})

	return await NewDPResponse<{ external_service_url: ExternalServiceUrlType }>(apiRes)
}
