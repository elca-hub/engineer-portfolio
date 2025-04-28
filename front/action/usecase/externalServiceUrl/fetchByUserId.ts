'use server'

import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import { apiPrefix } from '@/constants/constant'
import { NewDPResponse } from '@/lib/api'

export default async function fetchExternalServiceUrlByUserId(userId: string): Promise<ExternalServiceUrlType[] | null> {
	const fetchUser = await fetch(`${apiPrefix}/user/${userId}/external-service-url/`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
		},
	})

	const fetchUserRes = await NewDPResponse<{ external_service_urls: ExternalServiceUrlType[] }>(fetchUser)

	if (fetchUserRes.errors) {
		return null
	}

	const externalServiceUrl = fetchUserRes.data?.external_service_urls ?? null

	if (!externalServiceUrl) {
		return null
	}

	return externalServiceUrl
}
