'use server'

import { apiPrefix } from '@/constants/api'
import { cookieConfirmName } from '@/constants/cookieTokenName'
import { verifyJwt } from '@/lib/auth/jwt'
import { cookies } from 'next/headers'

type ErrorResponse = {
	status: number
	errors?: string[]
}

export const confirmEmailApi = async (accessCode: string): Promise<ErrorResponse> => {
	const cookie = await cookies()

	const emailToken = cookie.get(cookieConfirmName)

	const verify = await verifyJwt<{ email: string }>(emailToken?.value || '')

	if (verify === null) {
		return {
			status: 401,
			errors: ['1日を超過したため、再度新規登録を行なってください'],
		}
	}

	const res = await fetch(`${apiPrefix}/verification/email`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({
			email: verify.email,
			access_code: Number(accessCode),
		}),
		credentials: 'include',
		cache: 'no-cache',
	})

	if (res.ok) {
		cookie.delete(cookieConfirmName) // トークンを削除

		return { status: res.status }
	} else {
		const error = await res.json()

		return {
			status: res.status,
			errors: error.errors,
		}
	}
}
