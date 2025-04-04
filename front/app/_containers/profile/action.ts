'use server'

import { apiPrefix } from '@/constants/constant'
import { getSessionToken } from '@/lib/access'
import { DPResponseData, NewDPResponse } from '@/lib/api'

export default async function uploadIconApi(formData: FormData): Promise<DPResponseData<{ icon_path: string }>> {
	const token = await getSessionToken()
	const res = await fetch(`${apiPrefix}/auth/user/update`, {
		method: 'POST',
		body: formData,
		headers: {
			Accept: 'application/json',
			Authorization: `Bearer ${token}`,
		},
	})

	return await NewDPResponse<{ icon_path: string }>(res)
}
