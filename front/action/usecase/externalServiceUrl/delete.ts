'use server'

import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse } from '@/lib/api'

export default async function deleteExternalServiceUrl(token: string, id: string): Promise<DPResponseData<null>> {
	const apiRes = await fetch(`${apiPrefix}/auth/user/external-service-url/${id}/`, {
		method: 'DELETE',
		headers: {
			Authorization: `Bearer ${token}`,
		},
	})

	return await NewDPResponse<null>(apiRes)
}
