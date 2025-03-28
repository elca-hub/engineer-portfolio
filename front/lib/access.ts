'use server'

import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import { apiPrefix } from '@/constants/api'
import { getServerSession } from 'next-auth'
import 'server-only'

type LoginStatus = 'login' | 'not_login' | 'new_user' | 'error'

type RedirectStatus = {
	redirectPath: string
	isRedirect: boolean
}

export async function isLogin(): Promise<LoginStatus> {
	const session = await getServerSession(authOptions)

	if (!session) return 'not_login'

	const sessionUser = session.user

	if (!sessionUser) return 'not_login'

	const res = await fetch(`${apiPrefix}/is_exists?email=${sessionUser.email}`, {
		method: 'GET',
	})

	if (res.ok) {
		const json = (await res.json()) as { is_exists: boolean }

		return json.is_exists ? 'login' : 'new_user'
	} else {
		const errorJson = (await res.json()) as { errors: string[] }

		errorJson.errors.forEach((value) => console.error(value))

		return 'error'
	}
}

export async function handleAuthRedirect(nowPath: string): Promise<RedirectStatus> {
	const status = await isLogin()

	switch (status) {
		case 'login':
			return {
				redirectPath: '/dashboard',
				isRedirect: 'login' !== nowPath,
			}
		case 'new_user':
			return {
				redirectPath: '/register',
				isRedirect: 'register' !== nowPath,
			}
		case 'not_login':
			return {
				redirectPath: '/login',
				isRedirect: 'login' !== nowPath,
			}
		case 'error':
			return {
				redirectPath: '/login',
				isRedirect: 'login' !== nowPath,
			}
	}
}
