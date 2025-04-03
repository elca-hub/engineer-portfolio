'use server'

import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import { apiPrefix } from '@/constants/constant'
import { DPResponseData } from '@/lib/api'
import { getServerSession } from 'next-auth'
import { cookies } from 'next/headers'
import 'server-only'

type LoginStatus = 'login' | 'not_login' | 'new_user' | 'error' | 'cookie_expired'

type RedirectStatus = {
	redirectPath: string
	isRedirect: boolean
}

export async function getSessionToken(): Promise<string | null> {
	const cookie = await cookies()
	const dpSession = cookie.get('devport_api_token')

	if (!dpSession || dpSession.value === '') {
		return null
	}
	return dpSession.value
}

export async function logoutFlow(): Promise<DPResponseData<{ isSuccess: boolean }>> {
	const token = await getSessionToken()

	const res = await fetch(`${apiPrefix}/auth/user/logout`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${token}`,
		},
		body: JSON.stringify({
			token: token,
		}),
	})

	if (!res.ok) {
		const error = await res.json()
		console.error(error)

		return {
			errors: error.errors,
		}
	}

	const cookieStore = await cookies()

	cookieStore.delete('devport_api_token')

	return {
		data: { isSuccess: true },
	}
}

export async function loginFlow(email: string): Promise<DPResponseData<{ isSuccess: boolean; id: string }>> {
	const res = await fetch(`${apiPrefix}/login`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({
			email,
		}),
	})

	if (!res.ok) {
		const error = await res.json()
		console.error(error)

		return {
			errors: error.errors,
		}
	}
	const data = await res.json()
	const cookieStore = await cookies()
	cookieStore.set('devport_api_token', data.token, {
		maxAge: 60 * 60 * 24,
		httpOnly: true,
	})

	return {
		data: { isSuccess: true, id: data.user_id },
	}
}

export async function isLogin(): Promise<{ status: LoginStatus; userId?: string }> {
	const session = await getServerSession(authOptions)

	if (!session) return { status: 'not_login' }

	const sessionUser = session.user

	if (!sessionUser) return { status: 'not_login' }

	const res = await fetch(`${apiPrefix}/is_exists?email=${sessionUser.email}`, {
		method: 'GET',
	})

	if (res.ok) {
		const json = (await res.json()) as { is_exists: boolean; user_id: string }

		if (!json.is_exists) return { status: 'new_user' }

		const cookieStore = await cookies()
		const dpSession = cookieStore.get('devport_api_token')
		return !dpSession || dpSession.value === '' ? { status: 'cookie_expired' } : { status: 'login', userId: json.user_id }
	} else {
		const errorJson = (await res.json()) as { errors: string[] }

		errorJson.errors.forEach((value) => console.error(value))

		return { status: 'error' }
	}
}

export async function handleAuthRedirect(nowPath: string): Promise<RedirectStatus> {
	const login = await isLogin()

	switch (login.status) {
		case 'login':
			if (login.userId) {
				return {
					redirectPath: `/${login.userId}/profile`,
					isRedirect: '/login' === nowPath || '/register' === nowPath || '/auth-cookie' === nowPath,
				}
			} else {
				return {
					redirectPath: '/login',
					isRedirect: '/login' !== nowPath,
				}
			}
		case 'new_user':
			return {
				redirectPath: '/register',
				isRedirect: '/register' !== nowPath,
			}
		case 'not_login':
			return {
				redirectPath: '/login',
				isRedirect: '/login' !== nowPath,
			}
		case 'error':
			return {
				redirectPath: '/login',
				isRedirect: '/login' !== nowPath,
			}
		case 'cookie_expired':
			return {
				redirectPath: '/auth-cookie',
				isRedirect: nowPath === '/login',
			}
	}
}
