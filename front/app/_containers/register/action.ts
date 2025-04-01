'use server'
import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData, NewDPResponse } from '@/lib/api'
import { getServerSession } from 'next-auth'

export async function registerApi(formContent: { name: string; birthday: string }): Promise<DPResponseData<{ email: string }>> {
	const userSession = await getServerSession(authOptions)

	if (!userSession) {
		return { errors: ['session not found'] }
	}

	const sessionUserData = userSession.user

	if (!sessionUserData) {
		return { errors: ['user not found'] }
	}

	const res = await fetch(`${apiPrefix}/register`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({
			birthday: formContent.birthday.toString(),
			email: sessionUserData.email,
			name: formContent.name,
		}),
	})

	return await NewDPResponse<{ email: string }>(res)
}
